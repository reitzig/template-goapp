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

### Tooling 

- [jdx/mise](https://github.com/jdx/mise)
  for managing tools and environment variables
- [.editorconfig](https://editorconfig.org/)
  for sane collaboration
- [cocogitto/cocogitto](https://github.com/cocogitto/cocogitto)
  for maintaining an orderly Git history
- [evilmartians/lefthook](https://github.com/evilmartians/lefthook)
  for pre-commit hooks
- [miniscruff/changie](https://github.com/miniscruff/changie)
  for changelogs
- [goreleaser/goreleaser](https://github.com/goreleaser/goreleaser)
  for release assembly
- GitHub Actions for testing and releasing
- Dependabot config for golang and GitHub Actions

### Limitations

- [Trigger Release](.github/workflows/trigger-release.yaml)
  creates a release draft without a tag;
  _publishing_ the draft will create the tag on `main` 
  -- don't wait around!

### TODO

- set up logging
- include an end-to-end test with coverage
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
