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
// https://github.com/elastic/elasticsearch-specification/tree/4ca0cc05d3ae3fa06c2cd7be91905b656a474334


package types

import (
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/snowballlanguage"
)

// SnowballAnalyzer type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4ca0cc05d3ae3fa06c2cd7be91905b656a474334/specification/_types/analysis/analyzers.ts#L88-L93
type SnowballAnalyzer struct {
	Language  snowballlanguage.SnowballLanguage `json:"language"`
	Stopwords []string                          `json:"stopwords,omitempty"`
	Type      string                            `json:"type,omitempty"`
	Version   *string                           `json:"version,omitempty"`
}

// NewSnowballAnalyzer returns a SnowballAnalyzer.
func NewSnowballAnalyzer() *SnowballAnalyzer {
	r := &SnowballAnalyzer{}

	r.Type = "snowball"

	return r
}
