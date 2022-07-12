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
// https://github.com/elastic/elasticsearch-specification/tree/1b56d7e58f5c59f05d1641c6d6a8117c5e01d741

package types

// PhraseSuggest type.
//
// https://github.com/elastic/elasticsearch-specification/blob/1b56d7e58f5c59f05d1641c6d6a8117c5e01d741/specification/_global/search/_types/suggester.ts#L55-L60
type PhraseSuggest struct {
	Length  int                 `json:"length"`
	Offset  int                 `json:"offset"`
	Options PhraseSuggestOption `json:"options"`
	Text    string              `json:"text"`
}

// PhraseSuggestBuilder holds PhraseSuggest struct and provides a builder API.
type PhraseSuggestBuilder struct {
	v *PhraseSuggest
}

// NewPhraseSuggest provides a builder for the PhraseSuggest struct.
func NewPhraseSuggestBuilder() *PhraseSuggestBuilder {
	r := PhraseSuggestBuilder{
		&PhraseSuggest{},
	}

	return &r
}

// Build finalize the chain and returns the PhraseSuggest struct
func (rb *PhraseSuggestBuilder) Build() PhraseSuggest {
	return *rb.v
}

func (rb *PhraseSuggestBuilder) Length(length int) *PhraseSuggestBuilder {
	rb.v.Length = length
	return rb
}

func (rb *PhraseSuggestBuilder) Offset(offset int) *PhraseSuggestBuilder {
	rb.v.Offset = offset
	return rb
}

func (rb *PhraseSuggestBuilder) Options(options *PhraseSuggestOptionBuilder) *PhraseSuggestBuilder {
	v := options.Build()
	rb.v.Options = v
	return rb
}

func (rb *PhraseSuggestBuilder) Text(text string) *PhraseSuggestBuilder {
	rb.v.Text = text
	return rb
}