package main

import (
	"fmt"
	"go/parser"
	"go/token"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

type FileObject struct {
	path        string
	archtecture string
}

// var CategorizedFiles map[string][]FileObject = make(map[string][]FileObject)
var CategorizedFiles map[string][]string = make(map[string][]string)

const SOURCE_DIR = "../FrameworkPrototype"

func main() {
	// Define the directory containing Go source files
	var filesObjects []FileObject

	err := filepath.Walk(SOURCE_DIR, func(path string, fileInfo os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if fileInfo.IsDir() {
			return nil // Skip directories
		}
		if !strings.HasSuffix(fileInfo.Name(), ".go") {
			return nil // Skip non-Go files
		}
		var fileTags []string

		// Read the file
		//filePath := SOURCE_DIR + "/" + fileInfo.Name()
		src, err := os.ReadFile(path)
		if err != nil {
			fmt.Println("Error reading file:", err)
		}

		// Parse the source file
		fset := token.NewFileSet()
		file, err := parser.ParseFile(fset, fileInfo.Name(), src, parser.ParseComments)
		if err != nil {
			fmt.Println("Error parsing file:", err)
		}

		// Check the build tags in the file comments
		for _, commentGroup := range file.Comments {
			for _, comment := range commentGroup.List {
				if strings.HasPrefix(comment.Text, "// +build") ||
					strings.HasPrefix(comment.Text, "//go:build") ||
					strings.HasPrefix(comment.Text, "// go:build") ||
					strings.HasPrefix(comment.Text, "//+build") {
					// Extract the build tags
					bTags := strings.Fields(comment.Text)[1:]
					buildTags := removeItem(bTags, "+build")
					fileTags = append(fileTags, buildTags...)
				}
			}
		}

		processedFileTags := removeDuplicates(fileTags)
		p, err := filepath.Abs(path)
		if err != nil {
			panic(err)
		}
		if hasRequiredTag(processedFileTags, "wasm") {
			fobj := FileObject{path: p, archtecture: "wasm"}
			filesObjects = append(filesObjects, fobj)
		} else if hasRequiredTag(processedFileTags, "amd64") {
			fobj := FileObject{path: p, archtecture: "amd64"}
			filesObjects = append(filesObjects, fobj)
		}

		return nil

	})

	// List all files in the directory

	if err != nil {
		panic(err)
	}

	for _, fl := range filesObjects {
		switch fl.archtecture {
		case "amd64":
			CategorizedFiles["amd64"] = append(CategorizedFiles["amd64"], fl.path)
		case "wasm":
			CategorizedFiles["wasm"] = append(CategorizedFiles["wasm"], fl.path)
		}
	}

	fmt.Println(CategorizedFiles)

	for k, v := range CategorizedFiles {
		switch k {
		case "amd64":
			os.Setenv("GOARCH", "amd64")
			os.Setenv("GOOS", "linux")
			args := []string{"build", "-o", "output"}
			args = append(args, v...)
			fmt.Println(args)
			cmd := exec.Command("go", args...)
			err := cmd.Run()

			if err != nil {
				fmt.Println(err.Error())
			}
		case "wasm":
			os.Setenv("GOARCH", "wasm")
			os.Setenv("GOOS", "js")
			args := []string{"build", "-o", "output.wasm"}
			args = append(args, v...)
			fmt.Println(args)
			cmd := exec.Command("go", args...)
			err := cmd.Run()

			if err != nil {
				fmt.Println(err.Error())
			}
		}
	}

}

func hasRequiredTag(tags []string, requiredTag string) bool {
	for _, tag := range tags {
		if tag == requiredTag {
			return true
		}
	}
	return false
}

func removeItem(slice []string, item string) []string {
	var result []string
	for _, s := range slice {
		if s != item {
			result = append(result, s)
		}
	}
	return result
}

func removeDuplicates(slice []string) []string {
	encountered := map[string]bool{}
	result := []string{}

	for _, v := range slice {
		if encountered[v] == false {
			encountered[v] = true
			result = append(result, v)
		}
	}

	return result
}
