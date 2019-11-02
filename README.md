# github-api-cli

<img alt="travis build status" src="https://img.shields.io/travis/thamaraiselvam/github-api-cli?style=flat-square"> <img alt="Codecov code coverage" src="https://img.shields.io/codecov/c/github/thamaraiselvam/github-api-cli?style=flat-square">
<img alt="current tag" src="https://img.shields.io/github/v/tag/thamaraiselvam/github-api-cli.svg?sort=semver&style=flat-square">

github-api-cli provides an executable called gac, that can be used to access all of GitHub’s public API functionality from your command-line

github-api-cli is written in [cobra](https://github.com/spf13/cobra)

## Usage

### Install

```
brew tap thamaraiselvam/stable
brew install github-api-cli
```

### Available Commands

```
github-api-cli info <username>
github-api-cli list followers <username>
github-api-cli list following <username>
```

### Development

Run

```sh
$ git clone https://github.com/thamaraiselvam/git-api-cli.git
$ cd git-api-cli
$ make setup
$ make build
```
