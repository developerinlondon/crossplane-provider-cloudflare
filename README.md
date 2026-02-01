# Crossplane Provider Cloudflare

A [Crossplane](https://crossplane.io/) provider for Cloudflare built using [Upjet](https://github.com/crossplane/upjet) code generation tools. Exposes 207 Cloudflare resources as Kubernetes CRDs.

## Stack

| Component | Version |
|-----------|---------|
| Upjet | v2.2.0 |
| crossplane-runtime | v2.1.0 |
| Cloudflare TF Provider | 5.16.0 |
| Terraform | 1.5.7 (MPL licensed) |
| Go | 1.25.6 |

## Building and Publishing

### Understanding Docker Images vs xpkg

Crossplane does NOT use raw Docker images. It uses **xpkg** (Crossplane Package) format.

```
┌─────────────────────────────────────────────────────────────────┐
│                        WRONG APPROACH                           │
│  Docker image → push to ghcr.io → Crossplane ERROR             │
│                        "can't find package.yaml"                │
└─────────────────────────────────────────────────────────────────┘

┌─────────────────────────────────────────────────────────────────┐
│                       CORRECT APPROACH                          │
│  1. Build Docker image (runtime binary)                         │
│  2. Build xpkg (bundles: image ref + CRDs + metadata)          │
│  3. Push xpkg to ghcr.io → Crossplane works                    │
└─────────────────────────────────────────────────────────────────┘
```

**What's in each:**
- **Docker image**: Go binary (the provider controller)
- **xpkg**: OCI artifact containing CRDs + `crossplane.yaml` + reference to Docker image

Both are OCI images, but xpkg has Crossplane-specific structure.

### Build Steps

```bash
# 1. Generate code (if needed)
make generate

# 2. Build runtime Docker image
make build.all

# 3. Build xpkg (embeds runtime image + CRDs)
mkdir -p _output/xpkg/linux_amd64
crank xpkg build \
  --package-root ./package \
  --embed-runtime-image ghcr.io/developerinlondon/crossplane-provider-cloudflare:v0.1.0 \
  -o _output/xpkg/linux_amd64/provider-cloudflare.xpkg

# 4. Push xpkg to registry
crank xpkg push \
  --package-files _output/xpkg/linux_amd64/provider-cloudflare.xpkg \
  ghcr.io/developerinlondon/crossplane-provider-cloudflare:v0.1.2
```

### Quick Rebuild (CRDs only)

If you only changed CRDs (no Go code changes), skip the Docker image rebuild:

```bash
# Regenerate CRDs
make generate

# Copy to package directory (if needed)
cp test-crds/*.yaml package/crds/

# Build and push xpkg only
crank xpkg build \
  --package-root ./package \
  --embed-runtime-image ghcr.io/developerinlondon/crossplane-provider-cloudflare:v0.1.0 \
  -o _output/xpkg/linux_amd64/provider-cloudflare.xpkg

crank xpkg push \
  --package-files _output/xpkg/linux_amd64/provider-cloudflare.xpkg \
  ghcr.io/developerinlondon/crossplane-provider-cloudflare:v0.1.2
```

## Known Issues

### CRD Naming Collision (Snippet vs Snippets)

Cloudflare has two similar resources:
- `cloudflare_snippet` (singular) → Kind: `Snippet`
- `cloudflare_snippets` (plural) → Kind: `Snippets`

Both default to plural name `snippets`, causing CRD file collision. Fixed by adding `path=cfsnippet` to the Snippet type:

```go
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,cloudflare},path=cfsnippet
type Snippet struct {
```

## Developing

Run code-generation pipeline:
```console
make generate
```

Run against a Kubernetes cluster:
```console
make run
```

Build binary:
```console
make build
```

## Installation

```yaml
apiVersion: pkg.crossplane.io/v1
kind: Provider
metadata:
  name: provider-cloudflare
spec:
  package: ghcr.io/developerinlondon/crossplane-provider-cloudflare:v0.1.2
```

## Links

- [GitHub](https://github.com/developerinlondon/crossplane-provider-cloudflare)
- [Upjet Documentation](https://github.com/crossplane/upjet/blob/main/docs/generating-a-provider.md)
