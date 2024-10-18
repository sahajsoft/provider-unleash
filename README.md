# Provider Unleash

`provider-unleash` is a [Crossplane](https://crossplane.io/) provider that
is built using [Upjet](https://github.com/crossplane/upjet) code
generation tools and exposes XRM-conformant managed resources for the
Unleash API.

## Getting Started

Install the provider by using the following command after changing the image tag
to the [latest release](https://marketplace.upbound.io/providers/sahajsoft/provider-unleash):
```
up ctp provider install sahajsoft/provider-unleash:v0.1.0
```

Alternatively, you can use declarative installation:
```
cat <<EOF | kubectl apply -f -
apiVersion: pkg.crossplane.io/v1
kind: Provider
metadata:
  name: provider-unleash
spec:
  package: sahajsoft/provider-unleash:v0.1.0
EOF
```

Notice that in this example Provider resource is referencing ControllerConfig with debug enabled.

You can see the API reference [here](https://doc.crds.dev/github.com/sahajsoft/provider-unleash).

## Developing

Run code-generation pipeline:
```console
go run cmd/generator/main.go "$PWD"
```

Run against a Kubernetes cluster:

```console
make run
```

Build, push, and install:

```console
make all
```

Build binary:

```console
make build
```

## Testing

- Install unleash server on the cluster or configure the provider to use an existing server.
- Modify `examples/provider/secrets.yaml` with the correct values.

Apply the manifest:

```
k apply -f examples/providerconfig/secrets.yaml
k apply -f examples/providerconfig/providerconfig.yaml
```

Apply the example token manifest:

```
k apply -f examples/token/token.yaml
```

## Report a Bug

For filing bugs, suggesting improvements, or requesting new features, please
open an [issue](https://github.com/sahajsoft/provider-unleash/issues).

## Examples

You can find provider config and token configuration examples in the [examples](./examples) directory.

## Notes

This was automatically generated using [Upjet](https://github.com/crossplane/upjet/blob/main/docs/generating-a-provider.md)
