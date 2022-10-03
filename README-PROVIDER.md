# VRA Resource Provider

The VRA Resource Provider lets you manage [VRA](https://www.vmware.com/products/vrealize-automation.html) resources.

## Installing

This package is available for several languages/platforms:

### Node.js (JavaScript/TypeScript)

To use from JavaScript or TypeScript in Node.js, install using either `npm`:

```bash
npm install @pulumiverse/vra
```

or `yarn`:

```bash
yarn add @pulumiverse/vra
```

### Python

To use from Python, install using `pip`:

```bash
pip install pulumi_vra
```

### Go

To use from Go, use `go get` to grab the latest version of the library:

```bash
go get github.com/pulumiverse/pulumi-vra/sdk/go/...
```

### .NET

To use from .NET, install using `dotnet add package`:

```bash
dotnet add package Pulumi.Vra
```

## Configuration

The following configuration points are available for the `vra` provider:

- none
- `vra:accessToken` (environment: `n/a`) - the access token for `vra`
- `vra:refreshToken` (environment: `n/a`) - the refresh token for `vra`
- `vra:url` (environment: `n/a`) - the url of the `vra` instance
- `vra:insecure` (environment: `n/a`) - the option to ignore the `vra` instance certificate
- `vra:reauthorizeTimeout` (environment: `n/a`) - how often to reauthorize against the `vra` instance

## Reference

For detailed reference documentation, please visit [the Pulumi registry](https://www.pulumi.com/registry/packages/foo/api-docs/).
