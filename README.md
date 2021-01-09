# golang-github-app
GitHub Application Implemented in Golang
_____  

This repository was created to serve as an entry point to a fully operational GitHub App.

It is implemented in Go, using some of the following modules:

* [Logrus](https://github.com/sirupsen/logrus) (logging)

## Official materials

- [Getting started with apps](https://docs.github.com/en/free-pro-team@latest/developers/apps/getting-started-with-apps)
- [Endpoints available for GitHub Apps](https://docs.github.com/en/free-pro-team@latest/rest/overview/endpoints-available-for-github-apps)
- [Guides](https://docs.github.com/en/free-pro-team@latest/developers/apps/guides)


## Cool Golang Libraries & boilerplates

- [go-github](https://github.com/google/go-github) is a Go client library for accessing the GitHub API v3 from Google.
- [ghinstallation](https://github.com/bradleyfalzon/ghinstallation) is designed to provide automatic authentication for
  https://github.com/google/go-github or your own HTTP client.
- [go-githubapp](https://github.com/palantir/go-githubapp) is a library for building GitHub Apps and other services that
  handle GitHub webhooks.
- [go-github-app-boilerplate](https://github.com/sharkySharks/go-github-app-boilerplate) is a boilerplate for creating a
  GitHub App in Golang.

## Published articles

- [GitHub Apps in Golang](https://medium.com/@sharkysharks/github-apps-in-golang-1809bb4efb40)
- [Automating Github with Golang](https://www.x-cellent.com/blog/automating-github-with-golang-building-your-own-github-bot/)

## Roadmap

I will initially develop this repository as a first approach to a future internal application at my work. So code
will likely be implemented in the order that I need to approach new features.

# Versioning

In general, go-github follows [semver](https://semver.org/).  GitHub makes no guarantee about the stability of preview functionality,
so neither do I consider it a stable part of the artifact.

# License

This library is distributed under the Apache License found in the LICENSE file.  