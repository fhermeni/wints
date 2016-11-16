# Wints

A web application to manage final internships at Polytech.

## Requirements

- A postgreSQL database
- [Golang](https://golang.org/)
- [GoDep](https://github.com/tools/godep) to handle the dependencies
- [node.js](https://nodejs.org) to build the frontend
- [gulp](http://gulpjs.com/) to manage installation workflow

## Installation
```Shell
go get github.com/pierrre/gotestcover  
go get github.com/tools/godep  
```

In your `$GOPATH`:
```Shell
git clone https://github.com/fhermeni/wints.git src/github.com/fhermeni/wints # to get the source
cd src/github.com/fhermeni/wints/; godep restore # to restore the dependencies.
go install github.com/fhermeni/wints/wintsd # to build the executable.
```

## Usage of ./bin/wints:  
```Shell
  -conf string  
        Wints configuration file (default "wints.conf")  
  -fake-mailer  
        Do not send emails. Print them out stdout  
  -install-db  
        install the database  
  -new-root string  
        Invite a root user  
```
## Running
- `wints` launches the daemon. For test purposes, it is preferable to launch it with the `--fakeMailer` option to prevent to send mails (they will be logged into logs/mailer... instead)

## Benchmarks, tests

go test -x -v -tags=integration -bench BenchmarkInternships -cpuprofile=cpu.prof


## Developement

to run in development mode: `./dev.sh` --> server on https://localhost:8999
