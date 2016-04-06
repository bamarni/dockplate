# dockplate

Templating for Dockerfiles.

## Why?

It is common for Dockerfiles to share a similar base with slight modifications.
It happens when applying the same changes from different base images, or tuning images for production / development.

For those cases it might be useful to have a simple template which can be used
to generate different "flavours" of the same Dockerfile.

## Installation

### Prebuilt binaries

Prebuilt binaries are available in the [releases](https://github.com/bamarni/dockplate/releases).

### From source (requires Go)

    go get github.com/bamarni/dockplate

## Usage

```
# ./Dockerfile.template
FROM debian:{{.Env.BASE_TAG}}

# Common steps
# ...

{{if eq .Env.ENVIRONMENT "prod"}}
# Production tweaks
# ...
{{else if eq .Env.ENVIRONMENT "dev"}}
# Development tweaks
# ...
{{end}}

RUN ...

CMD ...
```

*The template parser is based on golang's `text/template`,
cf. https://golang.org/pkg/text/template/ for full reference.*

You can then generate Dockerfiles like this :

    BASE_TAG=jessie ENVIRONMENT=prod dockplate < Dockerfile.template > jessie-prod/Dockerfile

