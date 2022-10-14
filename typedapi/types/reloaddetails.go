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


package types

// ReloadDetails type.
//
// https://github.com/elastic/elasticsearch-specification/blob/9b556a1c9fd30159115d6c15226d0cac53a1d1a7/specification/indices/reload_search_analyzers/types.ts#L20-L24
type ReloadDetails struct {
	Index             string   `json:"index"`
	ReloadedAnalyzers []string `json:"reloaded_analyzers"`
	ReloadedNodeIds   []string `json:"reloaded_node_ids"`
}

// ReloadDetailsBuilder holds ReloadDetails struct and provides a builder API.
type ReloadDetailsBuilder struct {
	v *ReloadDetails
}

// NewReloadDetails provides a builder for the ReloadDetails struct.
func NewReloadDetailsBuilder() *ReloadDetailsBuilder {
	r := ReloadDetailsBuilder{
		&ReloadDetails{},
	}

	return &r
}

// Build finalize the chain and returns the ReloadDetails struct
func (rb *ReloadDetailsBuilder) Build() ReloadDetails {
	return *rb.v
}

func (rb *ReloadDetailsBuilder) Index(index string) *ReloadDetailsBuilder {
	rb.v.Index = index
	return rb
}

func (rb *ReloadDetailsBuilder) ReloadedAnalyzers(reloaded_analyzers ...string) *ReloadDetailsBuilder {
	rb.v.ReloadedAnalyzers = reloaded_analyzers
	return rb
}

func (rb *ReloadDetailsBuilder) ReloadedNodeIds(reloaded_node_ids ...string) *ReloadDetailsBuilder {
	rb.v.ReloadedNodeIds = reloaded_node_ids
	return rb
}
