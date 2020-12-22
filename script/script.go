//
//   script/script.go
//   mini-go
//
//   Copyright 2020 Daher Alfawares
//
//   Licensed under the Apache License, Version 2.0 (the "License");
//   you may not use this file except in compliance with the License
//   You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2
//
//   Unless required by applicable law or agreed to in writing,
//   distributed under the License is distributed on an "AS IS" BASIS
//   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied
//   See the License for the specific language governing permissions
//   limitations under the License

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
	err := writeToFile(name, "#!/bin/sh\nset -e\nset -o pipefail\n"+script)
	if err != nil {
		log.Println("script.Execute() failed to create file:", name, err)
		return
	}
	err = execute(name)
	if err != nil {
		log.Println("script.Execute() error running:", name, err)
	}
	err = os.Remove(name)
	if err != nil {
		log.Println("script.deleteFile() failed to remove file:", name, err)
	}
}
