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


// Package include
package include

import "strings"

// https://github.com/elastic/elasticsearch-specification/blob/9b556a1c9fd30159115d6c15226d0cac53a1d1a7/specification/ml/_types/Include.ts#L20-L42
type Include struct {
	name string
}

var (
	Definition = Include{"definition"}

	Featureimportancebaseline = Include{"feature_importance_baseline"}

	Hyperparameters = Include{"hyperparameters"}

	Totalfeatureimportance = Include{"total_feature_importance"}
)

func (i Include) MarshalText() (text []byte, err error) {
	return []byte(i.String()), nil
}

func (i *Include) UnmarshalText(text []byte) error {
	switch strings.ToLower(string(text)) {

	case "definition":
		*i = Definition
	case "feature_importance_baseline":
		*i = Featureimportancebaseline
	case "hyperparameters":
		*i = Hyperparameters
	case "total_feature_importance":
		*i = Totalfeatureimportance
	default:
		*i = Include{string(text)}
	}

	return nil
}

func (i Include) String() string {
	return i.name
}
