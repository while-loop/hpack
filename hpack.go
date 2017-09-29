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

type HeaderField struct {
	name      string
	value     string
	sensitive bool
}

// http://httpwg.org/specs/rfc7541.html#calculating.table.size
func (h *HeaderField) Size() uint {
	return uint(len(h.name) + len(h.value) + 32)
}

func (h *HeaderField) String() string {
	return fmt.Sprintf("sens: %v, name: %s, value: %s", h.sensitive, h.name, h.value)
}

type HeaderList []HeaderField

type Decoder struct {
}

type Encoder struct {
}
