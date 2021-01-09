# golang-github-app
GitHub Application Implemented in Golang
_____  

This repository was created to serve as an entry point to a fully operational GitHub App.  
  
GitHub applications have two parts:  
  
1. A GitHub application created on github.com on a GitHub organization that subscribes to certain events and has certain 
  permissions for repositories that the application is installed on.  
2. An API that receives events for the configured GitHub application in part 1 and responds to those events in some way, 
  however you like. The API authenticates each incoming request with the webhook secret token generated in part 1. 
  The API can be created in any language; programmer’s choice.

It is implemented in Go, using some of the following modules:

* [Logrus](https://github.com/sirupsen/logrus) (logging)
* [Viper](https://github.com/spf13/viper) (configuration)
* [go-github](https://github.com/google/go-github) (v3 REST API client)
* [githubv4]() (v4 GraphQL API client)


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
- [GitHub Demo Days youtube video](https://www.youtube.com/watch?v=iaBEWB1As0k) 



## Roadmap

I will initially develop this repository as a first approach to a future internal application at my work. So code
will likely be implemented in the order that I need to approach new features.

# Versioning

In general, go-github follows [semver](https://semver.org/).  GitHub makes no guarantee about the stability of preview functionality,
so neither do I consider it a stable part of the artifact.

# License

This library is distributed under the Apache License found in the LICENSE file.  