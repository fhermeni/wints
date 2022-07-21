#!/bin/sh
#Script to launch the development environment

ARG="run"
if [ $# -eq 1 ]; then
	ARG=$1
fi

case ${ARG} in
run)
	echo "==== godoc listening at :6060 ===="
	godoc -http=:6060&
	gulp assets watch&
	go run -ldflags "-X main.Version=SNAPSHOT" main.go --fake-mailer
	wait
	;;
install)
	npm install --save-dev  handlebars\
							gulp-handlebars\
							gulp-wrap\
							gulp-declare\
							gulp-concat\
							gulp-uglify\
							gulp-rename\
							gulp-clean-css\
							gulp-htmlmin\
							gulp-livereload\
							gulp-order\
							gulp-util\
							merge-stream\

	go get -u golang.org/x/tools/cmd/cover
	go get -u github.com/tools/godep
	go get -u github.com/pierrre/gotestcover
	go get -u github.com/kisielk/errcheck
	go get -u golang.org/x/text/encoding/charmap
	go get -u github.com/maruel/panicparse/cmd/pp
	go get -u github.com/stathat/go
	;;
*)
	echo "usage $0 (run|install)"
	;;
esac