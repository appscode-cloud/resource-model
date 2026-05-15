# AGENTS.md

This file provides guidance to coding agents (e.g. Claude Code, claude.ai/code) when working with code in this repository.

## Repository purpose

Go module `go.bytebuilders.dev/resource-model` — the canonical Go types for **ACE cluster/cloud/config resources** consumed by `b3` and other AppsCode platform services. Library plus a small set of CLIs (`cmd/`) that generate / inspect the model.

## Architecture

- `apis/`:
  - `cluster/` — cluster spec types (the dominant import surface; this is what credential managers like `aws-credential-manager` consume).
  - `cloud/` — cloud-provider metadata types.
  - `config/` — runtime config types.
- `crds/` — generated CRD YAMLs.
- `cmd/` — small utility CLIs.
- `pkg/` — library helpers.
- `config/` — Kubebuilder Kustomize bases.
- `data/` — embedded data (cloud catalogs, etc.).
- `hack/`, `Makefile` — codegen / build harness.
- `third_party/` — vendored third-party assets.
- `vendor/` — checked-in Go deps.

## Common commands

- `make ci` — full CI pipeline.
- `make gen` — regenerate clientset + manifests after API type changes.
- `make manifests` — regenerate CRDs only.
- `make clientset` — regenerate client code.
- `make fmt`, `make lint`, `make unit-tests` / `make test` — standard.
- `make verify` — codegen + module-tidy verification.
- `make add-license` / `make check-license` — manage license headers.

## Conventions

- Module path is `go.bytebuilders.dev/resource-model` (vanity URL); imports must use that.
- License: `LICENSE`. Sign off commits (`git commit -s`).
- This is a **shared library** — every exported type is API. Downstream operators (especially credential managers and `b3`) pin this dep. Changes ripple widely.
- Do not hand-edit `zz_generated.*.go` or `crds/*.yaml` — change `apis/<group>/<version>/*_types.go` and re-run `make gen`.
- Vendor directory is checked in; keep `go mod tidy && go mod vendor` clean.
