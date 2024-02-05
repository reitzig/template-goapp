<!-- editorconfig-checker-disable -->
[![license](https://img.shields.io/github/license/reitzig/template-goapp.svg)](https://github.com/reitzig/template-goapp/blob/master/LICENSE)
[![release](https://img.shields.io/github/release/reitzig/template-goapp.svg)](https://github.com/reitzig/template-goapp/releases/latest)
[![GitHub release date](https://img.shields.io/github/release-date/reitzig/template-goapp.svg)](https://github.com/reitzig/template-goapp/releases)
[![Tests](https://github.com/reitzig/template-goapp/actions/workflows/test.yaml/badge.svg)](https://github.com/reitzig/template-goapp/actions/workflows/test.yaml)
<!-- editorconfig-checker-enable -->

# Template: Go application

Just how I set up my go apps,
which I most certainly do not claim to be best practice or anything.
Comments welcome!

## Features

- [jdx/mise](https://github.com/jdx/mise)
  for managing tools and environment variables
- [.editorconfig](https://editorconfig.org/)
  for sane collaboration
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

- include a test with coverage
- set up logging
- stub for cobra + viper

## Instantiating the template

- Install mise, changie, lefthook, editorconfig-checker
- "Use this template" from GitHub web UI
- `rm .changes/unreleased/* .changes/v*.md && changie merge`
- Replace all occurrences of `[reitzig/]template-goapp` with `[your-account/]your-repo`
- Rewrite README, replace LICENSE
