package permissions

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

// The generator enforces these invariants at generation time; asserting them
// here keeps them true for anyone reading or hand-editing the generated file,
// and fails the unit-test job as well as the generate job.

func TestRoutePermissionsAreWellFormed(t *testing.T) {
	require.NotEmpty(t, RoutePermissions, "generation did not run")

	for route, perms := range RoutePermissions {
		method, path, ok := strings.Cut(route, " ")
		require.True(t, ok && method != "" && strings.HasPrefix(path, "/"), "route key %q is not %q", route, "METHOD /path")
		require.NotContains(t, path, "{", "%s: path uses OpenAPI templating, not gin's :id syntax", route)
		require.NotEmpty(t, perms, "%s: has no permissions", route)

		for _, perm := range perms {
			parts := strings.Split(perm, ":")
			require.Len(t, parts, 3, "%s: permission %q is not crec:<resource>:<action>", route, perm)
			require.Equal(t, "crec", parts[0], "%s: permission %q is not in the crec namespace", route, perm)
			for i, part := range parts {
				require.NotEmpty(t, part, "%s: permission %q has an empty segment %d", route, perm, i+1)
			}
		}
	}
}

func TestExemptRoutesAreNotAlsoPermissioned(t *testing.T) {
	for route := range ExemptRoutes {
		perms, ok := RoutePermissions[route]
		require.False(t, ok, "%s is exempt but also requires %q", route, perms)
	}
}

func TestRequiredPermissionsLookup(t *testing.T) {
	// A route that must exist, to catch the map being keyed in an unexpected shape.
	perms, ok := RequiredPermissions("POST", "/wallets")
	require.True(t, ok, "POST /wallets not found in RoutePermissions")
	require.Equal(t, []string{"crec:wallet:create"}, perms)

	// Unknown routes must report false so callers deny by default.
	_, ok = RequiredPermissions("POST", "/does-not-exist")
	require.False(t, ok, "unknown route reported a permission")

	require.True(t, IsExempt("GET", "/health-check"), "health check should be exempt")
	require.False(t, IsExempt("POST", "/wallets"), "POST /wallets should not be exempt")
}

func TestRequiredPermissionsLookupUsesGinPathSyntax(t *testing.T) {
	// Templated routes must be keyed with gin's :id syntax, matching what
	// gin.Context.FullPath() returns, not OpenAPI's {id}.
	perms, ok := RequiredPermissions("PATCH", "/channels/:channel_id/operations/:operation_id")
	require.True(t, ok, "gin-style route not found in RoutePermissions")
	require.Equal(t, []string{"crec:operation:sign"}, perms)

	_, ok = RequiredPermissions("PATCH", "/channels/{channel_id}/operations/{operation_id}")
	require.False(t, ok, "OpenAPI-style path should not resolve")
}

func TestBodyDiscriminatedEndpointsListMultiplePermissions(t *testing.T) {
	// Body-discriminated PATCH endpoints list both update and archive permissions;
	// the middleware checks at least one, the controller makes the final selection.
	perms, ok := RequiredPermissions("PATCH", "/wallets/:wallet_id")
	require.True(t, ok)
	require.Contains(t, perms, "crec:wallet:update")
	require.Contains(t, perms, "crec:wallet:archive")

	perms, ok = RequiredPermissions("PATCH", "/channels/:channel_id")
	require.True(t, ok)
	require.Contains(t, perms, "crec:channel:update")
	require.Contains(t, perms, "crec:channel:archive")

	perms, ok = RequiredPermissions("PATCH", "/channels/:channel_id/watchers/:watcher_id")
	require.True(t, ok)
	require.Contains(t, perms, "crec:watcher:update")
	require.Contains(t, perms, "crec:watcher:archive")
}
