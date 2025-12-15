package main

import (
	"flag"
	"fmt"
	"io/fs"
	"os"
	"os/exec"
	"path/filepath"
)

type FilesDestination struct {
	path   string
	dryRun bool
	// TODO handle indexing files
}

func newFilesDestination(path string) *FilesDestination {
	destination := FilesDestination{path: path}
	destination.dryRun = true
	return &destination

}

func (filesDestination *FilesDestination) copyToDestination(fileName string, filePath string) error {
	copyCommand := exec.Command(copyBinaryPath, filePath, filesDestination.path)
	fmt.Printf("Copying %s to %s\n", filePath, filesDestination.path)

	if filesDestination.dryRun == true {
		return nil
	}

	err := copyCommand.Run()

	if err != nil {
		panic(err)
	}

	fmt.Println("File Copied Sucessfully")
	return nil
}

const tempDirectory string = "temp"
const unzipBinaryPath string = "/usr/bin/unzip"
const copyBinaryPath string = "/usr/bin/cp"

var filesDestination *FilesDestination

func visit(path string, info fs.FileInfo, visitError error) error {
	if visitError != nil {
		panic(visitError)
	}
	if info.IsDir() {
		fmt.Println("Skipping folder...")
		return nil
	}

	filesDestination.copyToDestination(info.Name(), path)
	return nil
}

func main() {

	var filePointer = flag.String("file", "", "Zip file containing images")
	var fileNamePointer = flag.String("filename", "", "Filename to be used as prefix")
	var filesDestinationPointer = flag.String("fileDestination", "", "Destination to copy files to")
	flag.Parse()

	if *filePointer == "" {

		fmt.Println("ERROR: File path is required")
		flag.PrintDefaults()
		os.Exit(1)
	}
	if *fileNamePointer == "" {

		fmt.Println("ERROR: Filename is required")
		flag.PrintDefaults()
		os.Exit(1)
	}

	if *filesDestinationPointer == "" {

		fmt.Println("ERROR: Files destination is required")
		flag.PrintDefaults()
		os.Exit(1)
	}
	filesDestination = newFilesDestination(*filesDestinationPointer)
	command := exec.Command(unzipBinaryPath, *filePointer, "-d", tempDirectory)

	fmt.Println("Unzipping...")

	commandError := command.Run()
	if commandError != nil {
		panic(commandError)
	}
	defer os.RemoveAll(tempDirectory)

	readError := filepath.Walk(tempDirectory, visit)

	if readError != nil {
		panic(readError)
	}
}
