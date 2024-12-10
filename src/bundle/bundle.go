package main

import (
	"fmt"
	"github.com/shurcooL/vfsgen"
	"log"
	"net/http"
	"os"
)

func main() {
	// Assets contains project assets.
	dir, _ := os.Getwd()
	fmt.Println("bundle v0.1", dir)

	var Assets http.FileSystem = http.Dir("assets/")
	err := vfsgen.Generate(Assets, vfsgen.Options{
		PackageName: "main",
		BuildTags:   "!dev",
	})
	if err != nil {
		log.Fatalln(err)
	}
}
