//
//   expire/after.go
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

package expire

import (
	"fmt"
	"os"
	"time"
)

func after(expirationDateStr string) {
	expirationDate, err := time.Parse("02/01/2006", expirationDateStr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error parsing expiration date: %s\n", err)
		os.Exit(1)
	}
	currentDate := time.Now()
	if currentDate.After(expirationDate) {
		fmt.Fprintln(os.Stderr, "Error: This program has expired.")
		os.Exit(1)
	}
}
