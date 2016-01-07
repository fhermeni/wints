package main

import (
	"log"
	"net/http"

	"github.com/shurcooL/vfsgen"
)

func main() {
	var fs http.FileSystem = http.Dir("assets")

	err := vfsgen.Generate(fs, vfsgen.Options{PackageName: "httpd", Filename: "httpd/assets.go"})
	if err != nil {
		log.Fatalln(err)
	}
}
