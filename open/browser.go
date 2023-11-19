//
//   open/browser.go
//   mini-go
//
//   Copyright 2023 Daher Alfawares
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

package open

import (
	"fmt"
	"runtime"

	"github.com/skratchdot/open-golang/open"
)

func browser(url string) error {
	var err error
	switch runtime.GOOS {
	case "darwin":
		err = open.RunWith("open", url)
	case "windows":
		cmd := "cmd /c start " + url
		err = open.RunWith(cmd, "")
	default:
		err = fmt.Errorf("Unsupported platform")
	}
	return err
}
