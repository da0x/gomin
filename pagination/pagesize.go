//
//   pagination/pagesize.go
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

import (
	"log"
	"net/url"
	"strconv"
)

type PageSize struct {
	Page, Size int
}

func (ol *PageSize) Next() {
	ol.Page += 1
}

func (ol *PageSize) Values() url.Values {
	v := url.Values{}
	v.Set("size", strconv.Itoa(ol.Size))
	v.Set("page", strconv.Itoa(ol.Page))
	return v
}

func (ol *PageSize) Update(u *url.URL) *url.URL {
	v := u.Query()
	v.Set("size", strconv.Itoa(ol.Size))
	v.Set("page", strconv.Itoa(ol.Page))
	u.RawQuery = v.Encode()
	return u
}

func NewPageSize(size string) PageSize {
	s, err := strconv.Atoi(size)
	if err != nil {
		log.Fatalf("pagination: error converting page size: %v", err)
	}
	return PageSize{0, s}
}
