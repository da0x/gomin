//
//   pagination/pagination_test.go
//   golang
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

package pagination

import "testing"
import "net/url"

func TestOffsetLimit(t *testing.T) {
	o := OffsetLimit{0, 10}
	o.Next()
	println(o.Values().Encode())
	o.Next()
	println(o.Values().Encode())
	o.Next()
	println(o.Values().Encode())
	o.Next()
	println(o.Values().Encode())
	o.Next()
	println(o.Values().Encode())
}

func TestOffsetLimitUpdate(t *testing.T) {
	o := OffsetLimit{0, 10}
	u, e := url.Parse("http://golang.com")
	if e != nil {
		panic(e)
	}
	println(o.Update(u).String())
	o.Next()
	println(o.Update(u).String())
	o.Next()
	println(o.Update(u).String())
	o.Next()
	println(o.Update(u).String())
}

func TestPageSize(t *testing.T) {
	o := PageSize{0, 100}
	o.Next()
	println(o.Values().Encode())
	o.Next()
	println(o.Values().Encode())
	o.Next()
	println(o.Values().Encode())
	o.Next()
	println(o.Values().Encode())
	o.Next()
	println(o.Values().Encode())
}

func TestPageSizeUpdate(t *testing.T) {
	o := PageSize{0, 10}
	u, e := url.Parse("http://golang.com/?hello=world")
	if e != nil {
		panic(e)
	}
	println(o.Update(u).String())
	o.Next()
	println(o.Update(u).String())
	o.Next()
	println(o.Update(u).String())
	o.Next()
	println(o.Update(u).String())
}
