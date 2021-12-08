package main

import (
	"archive/tar"
	"compress/gzip"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"

	//search "GoPractice/search"
)

func main() {

	//grab the file name from the commandline args
	
	parentFile, err := os.Open(os.Args[1])
	if err != nil {
		panic("Bundle is missing as first argument: ")
	}
	defer parentFile.Close()
	childDir := extractInitalBundle(parentFile)
	extractChildBundles(childDir)
	//search.Demo()

}

//takes a file and sets up a tar/gzip reader then untars and creates a directory with the child tars
func extractInitalBundle(tarBundle *os.File) string {
	var splitDirectory []string
	gzf, err := gzip.NewReader(tarBundle)
	if err != nil {
		log.Fatal(err)
	}

	tarReader := tar.NewReader(gzf)

	for {
		hdr, err := tarReader.Next()
		if err == io.EOF {
			break // End of archive
		}
		if err != nil {
			log.Fatal(err)
		}

		var text = hdr.Name
		splitDirectory = strings.Split(text, "/")

		if _, err := os.Stat(splitDirectory[0]); os.IsNotExist(err) {
			// path/to/whatever does not exist
			os.Mkdir(splitDirectory[0], 0755)
		}
		f, err := os.Create(text)
		if _, err := io.Copy(f, tarReader); err != nil {

		}
	}
	return splitDirectory[0]
}

func extractChildBundles(childDir string) {
	var files []string
	var removedSuffixName = ""
	os.Mkdir("unzippedbundles", 0755)
	err := filepath.Walk(childDir, func(path string, info os.FileInfo, err error) error {
		files = append(files, path)
		return nil
	})
	if err != nil {
		panic(err)
	}

	for _, file := range files {

		//first entry will be the dir so we skip
		if file == "bundles" {
			continue
		}

		currentFile, err := os.Open(file)
		if err != nil {
			log.Fatal(err)
		}
		defer currentFile.Close()

		gzf, err := gzip.NewReader(currentFile)
		if err != nil {
			log.Fatal(err)
		}

		tarReader := tar.NewReader(gzf)

		for {
			hdr, err := tarReader.Next()
			if err == io.EOF {
				break // End of archive
			}
			if err != nil {
				log.Fatal(err)
			}
			var newFileName = "unzipped" + file

			if strings.Contains(file, "gz") {
				newFileName = strings.TrimSuffix(strings.TrimSuffix(newFileName, ".gz"), ".tar")
				removedSuffixName = newFileName + "-" + hdr.Name
			}
			if strings.Contains(removedSuffixName, "/") {
				removedSuffixName = strings.Replace(removedSuffixName, "/", "", -1)
			}

			f, err := os.Create(removedSuffixName)
			if _, err := io.Copy(f, tarReader); err != nil {

			}

		}
	}
}
