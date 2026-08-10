// Command gen-permissions derives the route -> RBAC permissions map from the
// OpenAPI spec and writes it as Go source.
//
// Every operation in the spec must declare the permission(s) it requires via
// the x-permissions extension, e.g.
//
//	post:
//	  x-permissions: [crec:wallet:create]
//
// An operation may list more than one permission, in which case a caller must
// hold at least one of them (OR, not AND) — list more than one when the
// endpoint is reachable via distinct, independently-grantable permissions,
// not to require several at once.
//
// Operations that intentionally require no permission (pre-auth infrastructure
// such as the health check) must be listed in exemptRoutes below, and should
// declare `x-permissions: []` in the spec to document the exemption (vacuum's
// x-permissions-required rule treats a missing key and an empty list
// differently — only the former is an error). Anything that is neither
// annotated nor exempt fails generation, so a new endpoint cannot ship
// unguarded by accident.
package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go/format"
	"os"
	"regexp"
	"sort"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
)

const permissionExtension = "x-permissions"

// exemptRoutes are operations that deliberately require no permission because
// they are served before an auth context exists. Adding to this list is a
// security-relevant change and should be reviewed as such.
var exemptRoutes = map[string]struct{}{
	"GET /health-check": {},
}

func main() {
	if len(os.Args) != 4 {
		fatalf("usage: gen-permissions <spec.yaml> <output.go> <package>")
	}
	specPath, outPath, pkg := os.Args[1], os.Args[2], os.Args[3]

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	doc, err := loader.LoadFromFile(specPath)
	if err != nil {
		fatalf("loading %s: %v", specPath, err)
	}

	perms := map[string][]string{}
	var problems []string

	for _, path := range sortedPaths(doc) {
		item := doc.Paths.Value(path)
		for method, op := range item.Operations() {
			// Keyed in gin's route syntax (:id), not OpenAPI's ({id}), since the
			// only consumer is gin's own c.FullPath() in the permission middleware.
			route := method + " " + toGinPath(path)
			if _, exempt := exemptRoutes[route]; exempt {
				if raw, annotated := op.Extensions[permissionExtension]; annotated {
					// An explicit x-permissions: [] documents the exemption and is
					// fine; a non-empty value is a genuine conflict with exemptRoutes.
					values, err := parsePermissions(raw)
					if err != nil || len(values) > 0 {
						problems = append(problems, fmt.Sprintf(
							"%s is in exemptRoutes but also declares a non-empty %s — remove one", route, permissionExtension))
					}
				}
				continue
			}

			raw, ok := op.Extensions[permissionExtension]
			if !ok {
				problems = append(problems, fmt.Sprintf(
					"%s is missing %s (add it, or add the route to exemptRoutes)", route, permissionExtension))
				continue
			}

			values, err := parsePermissions(raw)
			if err != nil {
				problems = append(problems, fmt.Sprintf("%s: %s: %v", route, permissionExtension, err))
				continue
			}
			if len(values) == 0 {
				problems = append(problems, fmt.Sprintf("%s: %s is empty", route, permissionExtension))
				continue
			}
			var bad bool
			for _, v := range values {
				if err := validatePermission(v); err != nil {
					problems = append(problems, fmt.Sprintf("%s: %v", route, err))
					bad = true
				}
			}
			if bad {
				continue
			}
			perms[route] = values
		}
	}

	if len(problems) > 0 {
		sort.Strings(problems)
		fmt.Fprintf(os.Stderr, "gen-permissions: %d problem(s) in %s:\n", len(problems), specPath)
		for _, p := range problems {
			fmt.Fprintf(os.Stderr, "  - %s\n", p)
		}
		os.Exit(1)
	}

	src, err := render(pkg, specPath, perms)
	if err != nil {
		fatalf("rendering output: %v", err)
	}
	if err := os.WriteFile(outPath, src, 0o644); err != nil {
		fatalf("writing %s: %v", outPath, err)
	}
	fmt.Printf("gen-permissions: wrote %s (%d routes, %d exempt)\n", outPath, len(perms), len(exemptRoutes))
}

// parsePermissions accepts the shapes the OpenAPI loader may produce for an
// extension value: a JSON array, a Go slice, or a bare string.
func parsePermissions(raw any) ([]string, error) {
	switch v := raw.(type) {
	case []string:
		return v, nil
	case string:
		return []string{v}, nil
	case []any:
		out := make([]string, 0, len(v))
		for _, e := range v {
			s, ok := e.(string)
			if !ok {
				return nil, fmt.Errorf("expected strings, got %T", e)
			}
			out = append(out, s)
		}
		return out, nil
	case json.RawMessage:
		var out []string
		if err := json.Unmarshal(v, &out); err != nil {
			var single string
			if err2 := json.Unmarshal(v, &single); err2 == nil {
				return []string{single}, nil
			}
			return nil, err
		}
		return out, nil
	default:
		return nil, fmt.Errorf("unsupported value type %T", raw)
	}
}

// validatePermission enforces the crec:<resource>:<action> shape so a
// typo becomes a build failure rather than a permanent 403 at runtime.
func validatePermission(p string) error {
	parts := strings.Split(p, ":")
	if len(parts) != 3 {
		return fmt.Errorf("permission %q must have the form crec:<resource>:<action>", p)
	}
	if parts[0] != "crec" {
		return fmt.Errorf("permission %q must start with the crec namespace", p)
	}
	for i, part := range parts {
		if part == "" {
			return fmt.Errorf("permission %q has an empty segment %d", p, i+1)
		}
	}
	return nil
}

var openAPIPathParam = regexp.MustCompile(`\{([^}]+)\}`)

// toGinPath converts an OpenAPI templated path ("/foo/{id}") to gin's route
// syntax ("/foo/:id"), matching what c.FullPath() returns at request time.
func toGinPath(path string) string {
	return openAPIPathParam.ReplaceAllString(path, ":$1")
}

func sortedPaths(doc *openapi3.T) []string {
	paths := make([]string, 0, doc.Paths.Len())
	for path := range doc.Paths.Map() {
		paths = append(paths, path)
	}
	sort.Strings(paths)
	return paths
}

func render(pkg, specPath string, perms map[string][]string) ([]byte, error) {
	routes := make([]string, 0, len(perms))
	for route := range perms {
		routes = append(routes, route)
	}
	sort.Strings(routes)

	exempt := make([]string, 0, len(exemptRoutes))
	for route := range exemptRoutes {
		exempt = append(exempt, route)
	}
	sort.Strings(exempt)

	var b bytes.Buffer
	fmt.Fprintf(&b, "// Code generated by cmd/gen-permissions from %s. DO NOT EDIT.\n\n", specPath)
	fmt.Fprintf(&b, "package %s\n\n", pkg)
	b.WriteString("// RoutePermissions maps \"METHOD <gin route pattern>\" to the crec:* permissions\n")
	b.WriteString("// the route requires; a caller must hold at least one of them (OR, not AND).\n")
	b.WriteString("// Paths use gin's :id syntax, matching gin.Context.FullPath(), not OpenAPI's\n")
	b.WriteString("// {id}.\n")
	b.WriteString("var RoutePermissions = map[string][]string{\n")
	for _, route := range routes {
		fmt.Fprintf(&b, "\t%q: {", route)
		for _, p := range perms[route] {
			fmt.Fprintf(&b, "%q, ", p)
		}
		b.WriteString("},\n")
	}
	b.WriteString("}\n\n")
	b.WriteString("// ExemptRoutes require no permission because they are served before an auth\n")
	b.WriteString("// context exists.\n")
	b.WriteString("var ExemptRoutes = map[string]struct{}{\n")
	for _, route := range exempt {
		fmt.Fprintf(&b, "\t%q: {},\n", route)
	}
	b.WriteString("}\n\n")
	b.WriteString("// RequiredPermissions returns the permissions a route requires; the caller must\n")
	b.WriteString("// hold at least one of them (OR, not AND). The second result is false when the\n")
	b.WriteString("// route is exempt or unknown; callers must treat unknown routes as denied.\n")
	b.WriteString("func RequiredPermissions(method, path string) ([]string, bool) {\n")
	b.WriteString("\tperms, ok := RoutePermissions[method+\" \"+path]\n")
	b.WriteString("\treturn perms, ok\n")
	b.WriteString("}\n\n")
	b.WriteString("// IsExempt reports whether a route is exempt from permission checks.\n")
	b.WriteString("func IsExempt(method, path string) bool {\n")
	b.WriteString("\t_, ok := ExemptRoutes[method+\" \"+path]\n")
	b.WriteString("\treturn ok\n")
	b.WriteString("}\n")

	return format.Source(b.Bytes())
}

func fatalf(format string, args ...any) {
	fmt.Fprintf(os.Stderr, "gen-permissions: "+format+"\n", args...)
	os.Exit(1)
}
