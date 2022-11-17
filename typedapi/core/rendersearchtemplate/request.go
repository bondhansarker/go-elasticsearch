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
// https://github.com/elastic/elasticsearch-specification/tree/555082f38110f65b60d470107d211fc354a5c55a


package rendersearchtemplate

import (
	"encoding/json"
	"fmt"
)

// Request holds the request body struct for the package rendersearchtemplate
//
// https://github.com/elastic/elasticsearch-specification/blob/555082f38110f65b60d470107d211fc354a5c55a/specification/_global/render_search_template/RenderSearchTemplateRequest.ts#L25-L39
type Request struct {
	File   *string                `json:"file,omitempty"`
	Params map[string]interface{} `json:"params,omitempty"`
	Source *string                `json:"source,omitempty"`
}

// NewRequest returns a Request
func NewRequest() *Request {
	r := &Request{
		Params: make(map[string]interface{}, 0),
	}
	return r
}

// FromJSON allows to load an arbitrary json into the request structure
func (rb *Request) FromJSON(data string) (*Request, error) {
	var req Request
	err := json.Unmarshal([]byte(data), &req)

	if err != nil {
		return nil, fmt.Errorf("could not deserialise json into Rendersearchtemplate request: %w", err)
	}

	return &req, nil
}