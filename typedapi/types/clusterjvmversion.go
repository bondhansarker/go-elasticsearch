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

// ClusterJvmVersion type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4ca0cc05d3ae3fa06c2cd7be91905b656a474334/specification/cluster/stats/types.ts#L161-L169
type ClusterJvmVersion struct {
	BundledJdk      bool   `json:"bundled_jdk"`
	Count           int    `json:"count"`
	UsingBundledJdk bool   `json:"using_bundled_jdk"`
	Version         string `json:"version"`
	VmName          string `json:"vm_name"`
	VmVendor        string `json:"vm_vendor"`
	VmVersion       string `json:"vm_version"`
}

// NewClusterJvmVersion returns a ClusterJvmVersion.
func NewClusterJvmVersion() *ClusterJvmVersion {
	r := &ClusterJvmVersion{}

	return r
}
