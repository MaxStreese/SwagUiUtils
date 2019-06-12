package main

import (
	"gopkg.in/src-d/go-billy.v4/memfs"
	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/storage/memory"
	"io"
	"log"
	"os"
)

func main() {
	fs := memfs.New()

	storer := memory.NewStorage()

	_, err := git.Clone(storer, fs, &git.CloneOptions{
		URL:      "https://github.com/swagger-api/swagger-ui.git",
		Progress: os.Stdout,
	})
	if err != nil {
		log.Fatal(err)
	}

	// Prints the content of the CHANGELOG file from the cloned repository
	changelog, err := fs.Open(".gitignore")
	if err != nil {
		log.Fatal(err)
	}

	io.Copy(os.Stdout, changelog)
}

func fatalLogOnErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

/*
- clone repository
- checkout tag
- loop files to create their byte representation
- save byte representation
*/
