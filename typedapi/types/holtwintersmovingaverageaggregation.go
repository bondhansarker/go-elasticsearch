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
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/gappolicy"
)

// HoltWintersMovingAverageAggregation type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4ca0cc05d3ae3fa06c2cd7be91905b656a474334/specification/_types/aggregations/pipeline.ts#L222-L225
type HoltWintersMovingAverageAggregation struct {
	// BucketsPath Path to the buckets that contain one set of values to correlate.
	BucketsPath *string                  `json:"buckets_path,omitempty"`
	Format      *string                  `json:"format,omitempty"`
	GapPolicy   *gappolicy.GapPolicy     `json:"gap_policy,omitempty"`
	Meta        map[string]interface{}   `json:"meta,omitempty"`
	Minimize    *bool                    `json:"minimize,omitempty"`
	Model       string                   `json:"model,omitempty"`
	Name        *string                  `json:"name,omitempty"`
	Predict     *int                     `json:"predict,omitempty"`
	Settings    HoltWintersModelSettings `json:"settings"`
	Window      *int                     `json:"window,omitempty"`
}

// NewHoltWintersMovingAverageAggregation returns a HoltWintersMovingAverageAggregation.
func NewHoltWintersMovingAverageAggregation() *HoltWintersMovingAverageAggregation {
	r := &HoltWintersMovingAverageAggregation{}

	r.Model = "holt_winters"

	return r
}
