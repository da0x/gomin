//
//   pagination/offsetlimit.go
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

type OffsetLimit struct {
	Offset, Limit int
}

func (ol *OffsetLimit) Next() {
	ol.Offset += ol.Limit
}

func (ol *OffsetLimit) Values() url.Values {
	v := url.Values{}
	v.Add("offset", strconv.Itoa(ol.Offset))
	v.Add("limit", strconv.Itoa(ol.Limit))
	return v
}

func NewOffsetLimit(limit string) OffsetLimit {
	l, err := strconv.Atoi(limit)
	if err != nil {
		log.Fatalf("pagination: error converting offset limit: %v", err)
	}
	return OffsetLimit{0, l}
}
