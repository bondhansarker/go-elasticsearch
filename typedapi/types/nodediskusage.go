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

// NodeDiskUsage type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4ca0cc05d3ae3fa06c2cd7be91905b656a474334/specification/cluster/allocation_explain/types.ts#L56-L60
type NodeDiskUsage struct {
	LeastAvailable DiskUsage `json:"least_available"`
	MostAvailable  DiskUsage `json:"most_available"`
	NodeName       string    `json:"node_name"`
}

// NewNodeDiskUsage returns a NodeDiskUsage.
func NewNodeDiskUsage() *NodeDiskUsage {
	r := &NodeDiskUsage{}

	return r
}
