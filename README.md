# ProxyFilter

💣 Filter Bad Or Slow Proxies Out From Big Lists 📃

## Getting Started

Following these instructions would get you a copy of this project up and running on your local machine for development and testing purposes. See deployment for notes on how to deploy the project on a live system.

### Prerequisites

Project Requirements

```
Golang - Version 1.11.1 and above
Visual Studio Code - Version 1.30.0 and above
Proxy List - List to be in LF (Line Feed) form, not CRLF (Carriage Return, Line Feed)
```

### Installing

Follow these steps if you need a development environment running

```
Golang - https://golang.org/dl
Visual Studio Code - https://code.visualstudio.com/download
```

Steps:
1. Run `git clone https://github.com/Etosticity/proxyfilter`
2. Inside the Git Clone directory, run `code .`
3. Finally, run `go run main.go`

## Full Deployment

Follow these steps if you need binaries to be fully installed into your local machine

Steps:
1. Run `git clone https://github.com/Etosticity/proxyfilter`
2. Inside the Git Cloned directory, run `go clean`
3. Then, run `go install`
4. Finally, run `proxyfilter`

## Portable Deployment

Follow these steps if you don't want binaries to be installed into your local machine

Steps:
1. Run `git clone https://github.com/Etosticity/proxyfilter`
2. Inside the Git Cloned directory, run `go clean`
3. Then, run `go build -v`
4. Depending on your Operating System, run `./proxyfilter` or `proxyfilter.exe` accordingly

## Release History
* v1.4.5-alpha
  * Fixed issue [issue #1](https://github.com/Etosticity/proxyfilter/issues/1)
  * Added more documentation
  * Added a new system resource usage limiter
* v1.3.4-alpha
  * Simplified command line flags checking
* v1.3.3-alpha
  * Added README.md
* v1.3.2-alpha
  * Added documentation
  * Added return statements
  * Removed filtered proxy list
* v1.1.0
  * Added LICENSE
  * Added .gitignore
  * Added main tool
  * Added an example proxy list
  * Added [@godacity_](https://www.instagram.com/godacity_)'s original Python tool

## Built With
* [Go](https://golang.org) - The Programming Language Used
* [Git](https://git-scm.com) - The Source Control Manager Used
* [Visual Studio Code](https://code.visualstudio.com) - The Integrated Development Environment Used

## Versioning

We use [SemVer](http://semver.org) for versioning. For the versions available, see the [tags in this repository](https://github.com/Etosticity/proxyfilter/tags).

## Authors

* **@godacity_** - *Original Tool Developer* - [Instagram](https://www.instagram.com/godacity_)
* **Etosticity Rammington** - *Project Owner* - [Instagram](https://www.instagram.com/etosticity)

Also, see the list of [contributors](https://github.com/Etosticity/proxyfilter/graphs/contributors) who participated in this project.

## License

**ProxyFilter** is licensed under AGPL-3.0 - see [LICENSE](https://github.com/Etosticity/proxyfilter/blob/master/LICENSE) file for more details

## Acknowledgments

- [@godacity_](https://www.instagram.com/godacity_) - For allowing his [original code](https://github.com/Etosticity/proxyfilter/blob/master/_%40godacity/proxy_checker.py) in Python to be ported over into Golang. 