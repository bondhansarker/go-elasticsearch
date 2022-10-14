// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.


// Code generated from the elasticsearch-specification DO NOT EDIT.
// https://github.com/elastic/elasticsearch-specification/tree/9b556a1c9fd30159115d6c15226d0cac53a1d1a7


// Package delimitedpayloadencoding
package delimitedpayloadencoding

import "strings"

// https://github.com/elastic/elasticsearch-specification/blob/9b556a1c9fd30159115d6c15226d0cac53a1d1a7/specification/_types/analysis/token_filters.ts#L61-L65
type DelimitedPayloadEncoding struct {
	name string
}

var (
	Int = DelimitedPayloadEncoding{"int"}

	Float = DelimitedPayloadEncoding{"float"}

	Identity = DelimitedPayloadEncoding{"identity"}
)

func (d DelimitedPayloadEncoding) MarshalText() (text []byte, err error) {
	return []byte(d.String()), nil
}

func (d *DelimitedPayloadEncoding) UnmarshalText(text []byte) error {
	switch strings.ToLower(string(text)) {

	case "int":
		*d = Int
	case "float":
		*d = Float
	case "identity":
		*d = Identity
	default:
		*d = DelimitedPayloadEncoding{string(text)}
	}

	return nil
}

func (d DelimitedPayloadEncoding) String() string {
	return d.name
}
