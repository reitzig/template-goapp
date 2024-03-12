<!-- editorconfig-checker-disable -->
[![license](https://img.shields.io/github/license/reitzig/template-goapp.svg)](https://github.com/reitzig/template-goapp/blob/master/LICENSE)
[![release](https://img.shields.io/github/release/reitzig/template-goapp.svg)](https://github.com/reitzig/template-goapp/releases/latest)
[![GitHub release date](https://img.shields.io/github/release-date/reitzig/template-goapp.svg)](https://github.com/reitzig/template-goapp/releases)
[![Tests](https://github.com/reitzig/template-goapp/actions/workflows/test.yaml/badge.svg)](https://github.com/reitzig/template-goapp/actions/workflows/test.yaml)
[![Go Report Card](https://goreportcard.com/badge/github.com/reitzig/template-goapp)](https://goreportcard.com/report/github.com/reitzig/template-goapp)
<!-- editorconfig-checker-enable -->

# Template: Go application

Just how I set up my go apps,
which I most certainly do not claim to be best practice or anything.
Comments welcome!

## Features

* [spf13/cobra](https://github.com/spf13/cobra)
  for implementing the command-line interface
* Acceptance tests with Cucumber

### Testing

- We include a [test workflow](.github/workflows/test.yaml) that
  runs unit and acceptance tests with coverage.
- Run unit tests through `go test`, as per usual.
- Run acceptance tests through [test/Dockerfile](test/Dockerfile);
  a run configuration for IDEA/GoLand is 
  [included](.run/Run%20Acceptance%20Tests.run.xml).
- Run acceptance tests with coverage like so:
  ```bash
  docker rm template-goapp-test
  rm -f test/coverage/*
  
  docker build -f test/Dockerfile -t template-goapp-test . \
    && docker run -it --name template-goapp-test template-goapp-test \
    && docker cp template-goapp-test:/home/test/app/coverage test/
  go tool covdata percent -i=test/coverage
  ```
  <!-- TODO: Is there a nicer way? -->

### Tooling 

- [jdx/mise](https://github.com/jdx/mise)
  for managing tools and environment variables
- [.editorconfig](https://editorconfig.org/)
  for consistent formatting
- [golangci/golangci-lint](https://github.com/golangci/golangci-lint)
  for maintaining orderly Go code
- [rubocop/rubocop](https://github.com/rubocop/rubocop)
  for maintaining orderly Ruby code
- [cocogitto/cocogitto](https://github.com/cocogitto/cocogitto)
  for maintaining an orderly Git history
- [evilmartians/lefthook](https://github.com/evilmartians/lefthook)
  for pre-commit hooks
- [miniscruff/changie](https://github.com/miniscruff/changie)
  for creating changelogs
- [goreleaser/goreleaser](https://github.com/goreleaser/goreleaser)
  for release assembly
- GitHub Actions for testing and releasing
- Dependabot for keeping dependencies up to date

### Limitations

- [Trigger Release](.github/workflows/trigger-release.yaml)
  creates a release draft without a tag;
  _publishing_ the draft will create the tag on `main` 
  -- don't wait around!

### TODO

- set up logging
- stub for viper
- CI: run tests before releasing

## Instantiating the template

- Install mise, changie, cocogitto, lefthook, editorconfig-checker
  <!-- TODO: Tool for that? -->
- "Use this template" from GitHub web UI and clone
- Run `lefthook install`
- `rm .changes/unreleased/* .changes/v*.md && changie merge`
- Replace all occurrences of `[reitzig/]template-goapp` with `[your-account/]your-app`
- Rewrite README, replace LICENSE
