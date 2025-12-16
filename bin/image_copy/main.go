package main

import (
	"flag"
	"fmt"
	"io/fs"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
)

type FilesDestination struct {
	path             string
	fileName         string
	realRun          bool
	currentFileIndex int
}

func newFilesDestination(path string, fileName string, realRun bool) *FilesDestination {
	destination := FilesDestination{path: path, fileName: fileName, realRun: realRun}
	destination.currentFileIndex = 1
	return &destination

}

func (filesDestination *FilesDestination) copyToDestination(fileName string, filePath string) error {
	newFileName := fmt.Sprintf("%s-%s%s", filesDestination.fileName, strconv.Itoa(filesDestination.currentFileIndex), filepath.Ext(filePath))
	newFilePath := fmt.Sprintf("%s/%s", filesDestination.path, newFileName)
	copyCommand := exec.Command(copyBinaryPath, filePath, newFilePath)
	fmt.Printf("Copying %s to %s\n", filePath, newFilePath)

	if filesDestination.realRun == false {
		return nil
	}

	err := copyCommand.Run()

	if err != nil {
		panic(err)
	}

	filesDestination.currentFileIndex++
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
	var realRunPointer = flag.Bool("realRun", false, "Set to true to actually modify files, false to check files to be changed")
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

	filesDestination = newFilesDestination(*filesDestinationPointer, *fileNamePointer, *realRunPointer)
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

	if filesDestination.realRun == false {
		fmt.Println("\nThis was a dry run, use the --realRun flag if you are happy with the current output")
	}
}
