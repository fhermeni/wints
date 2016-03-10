# Wints

A web application to manage final internships at Polytech.

## Requirements

- A postgreSQL database
- [Golang](https://golang.org/)
- [GoDep](https://github.com/tools/godep) to handle the dependencies
- [node.js](nodejs.org) to build the frontend

## Installation

go get github.com/pierrre/gotestcover
go get github.com/tools/godep
	
In your `$GOPATH`:

- `git clone git@scm-oasis.inria.fr:fhermeni/wints.git src/scm-oasis.inria.fr/fhermeni/wints` to get the source
- `cd src/scm-oasis.inria.fr/fhermeni/wints/wintsd; godep restore` to restore the dependencies.
- `go install scm-oasis.inria.fr/fhermeni/wints/wintsd` to build the executable.

## Configuration
- `wintsd --generate-config > wints.conf` to generate a blank config
- customize the configuration file as you need 
- `wintsd --install` to generate the database tables
- `wintsd --test` to check if everything is ok

## Running
- `wintsd` launches the daemon. For test purposes, it is preferable to launch it with the `--fakeMailer` option to prevent to send mails (they will be printed on stdout instead)

## Benchmarks, tests

go test -x -v -tags=integration -bench BenchmarkInternships -cpuprofile=cpu.prof


## Developement
```
assets
  |- dist
  |	   |- css/wints.min.css
  |	   |- js/wints.min.js
  |	   |- html/*.html
  | - src
  		|- js
  		|- css
  		| -hbs
  		\- html
```

- concat and minify css
- concat and minify javascript
- minify html
- precompile handlebars