// Copyright 2017 Anthony Alves
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package hpack

import "fmt"

var (
	staticTable        = newStaticTable()
	tableOverflowError = fmt.Errorf("needed space exceeds table allowed max size")
)

type pair struct {
	first  interface{}
	second interface{}
}

type table struct {
	vals     map[string]uint
	pairVals map[pair]uint
	fields   []HeaderField

	size           uint // bytes
	maxSize        uint // bytes
	allowedMaxSize uint // bytes [0, n]
}

func newTable(allowedMaxSize uint) *table {
	return &table{
		vals:           map[string]uint{},
		pairVals:       map[pair]uint{},
		allowedMaxSize: allowedMaxSize,
		maxSize:        allowedMaxSize / 2,
	}
}

func newStaticTable() *table {
	t := newTable(1024)
	for _, headerField := range staticHeaders {
		t.addField(headerField)
	}
	return t
}

func (t *table) addField(field HeaderField) error {
	if err := t.evict(field.Size()); err != nil {
		return err
	}

	id := uint(len(t.vals) + 1)
	t.vals[field.name] = id
	t.pairVals[f2p(field)] = id
	t.size += field.Size()

	return nil
}

// recursive evict method
func (t *table) evict(neededSpace uint) error {
	if neededSpace > t.allowedMaxSize {
		return tableOverflowError
	}

	if t.size+neededSpace < t.maxSize || len(t.fields) <= 0 {
		return nil
	}

	// pop the first header field
	headerField := t.fields[0]

	// remove the field from table
	copy(t.fields, t.fields[1:])
	t.size -= headerField.Size()
	delete(t.pairVals, f2p(headerField))
	delete(t.vals, headerField.name)

	return t.evict(neededSpace)
}

func (t *table) search(h HeaderField) {

}

// HeaderField to pair
func f2p(h HeaderField) pair {
	return pair{h.name, h.value}
}

var (
	staticHeaders = []HeaderField{
		{":authority", "", false},
		{":method", "GET", false},
		{":method", "POST", false},
		{":path", "/", false},
		{":path", "index.html", false},
		{":scheme", "http", false},
		{":scheme", "https", false},
		{":status", "200", false},
		{":status", "204", false},
		{":status", "206", false},
		{":status", "304", false},
		{":status", "400", false},
		{":status", "404", false},
		{":status", "500", false},
		{"accept-charset", "", false},
		{"accept-encoding", "gzip, deflate", false},
		{"accept-language", "", false},
		{"accept-ranges", "", false},
		{"accept", "", false},
		{"access-control-allow-origin", "", false},
		{"age", "", false},
		{"allow", "", false},
		{"authorization", "", false},
		{"cache-control", "", false},
		{"content-disposition", "", false},
		{"content-encoding", "", false},
		{"content-language", "", false},
		{"content-length", "", false},
		{"content-location", "", false},
		{"content-range", "", false},
		{"content-type", "", false},
		{"cookie", "", false},
		{"date", "", false},
		{"etag", "", false},
		{"expect", "", false},
		{"expires", "", false},
		{"from", "", false},
		{"host", "", false},
		{"if-match", "", false},
		{"if-modified-since", "", false},
		{"if-none-match", "", false},
		{"if-range", "", false},
		{"if-unmodified-since", "", false},
		{"last-modified", "", false},
		{"link", "", false},
		{"location", "", false},
		{"max-forwards", "", false},
		{"proxy-authenticate", "", false},
		{"proxy-authorization", "", false},
		{"range", "", false},
		{"referer", "", false},
		{"refresh", "", false},
		{"retry-after", "", false},
		{"server", "", false},
		{"set-cookie", "", false},
		{"strict-transport-security", "", false},
		{"transfer-encoding", "", false},
		{"user-agent", "", false},
		{"vary", "", false},
		{"via", "", false},
		{"www-authenticate", "", false},
	}
)
