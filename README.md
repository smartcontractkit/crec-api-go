# CREC-API :: Golang Bindings

To regenerate the Golang bindings for the CREC API and models, run:

```bash
make tools && make generate
```

# Adding/Updating a Service

1. Copy ABI file into `services/{service}/abi`

2. Add `abigen` command to `Makefile`

```
abigen --abi services/{service}/abi/{contract}.abi.json --pkg {unique_package_name} --out services/{service}/gen/{unique_package_name}/{unique_package_name}.gen.go
## be sure to create the new directory under the service's `gen` directory
mkdir -p services/{service}/gen/{unique_package_name}
```

3. Run `make generate`

4. Create schema file (json) in `services/{service}/schema` that maps all contract events to the `verifiableEvent` model, which can be found an existing schemas.
