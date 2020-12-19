package script

import (
	"github.com/google/uuid"
	"io"
	"log"
	"os"
	"os/exec"
)

// writeToFile will print any string of text to a file safely by
// checking for errors and syncing at the end.
func writeToFile(filename string, data string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.WriteString(file, data)
	if err != nil {
		return err
	}
	return file.Sync()
}

// deleteFile deletes the given path from the system.
// it will panic upon failure
func deleteFile(filename string) {
	err := os.Remove(filename)
	if err != nil {
		log.Fatalln("script.deleteFile() failed to remove file:", filename, err)
	}
}

func execute(filename string) error {
	cmd := exec.Command("sh", filename)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

// Execute saves the input string into a file
// runs the file as a shell script then
// it deletes the file
// all execution output of the script will be
// visible in your terminal
func Execute(script string) {
	name := uuid.New().String() + ".sh"
	err := writeToFile(name, "#!/bin/sh\nset -e\n set -o pipefail\n"+script)
	if err != nil {
		log.Fatalln("script.Execute() failed to create file:", name, err)
	}
	defer deleteFile(name)
	err = execute(name)
	if err != nil {
		log.Fatalln("script.Execute() error running:", name, err)
	}
}
