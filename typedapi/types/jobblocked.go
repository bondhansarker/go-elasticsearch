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

import (
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/jobblockedreason"
)

// JobBlocked type.
//
// https://github.com/elastic/elasticsearch-specification/blob/9b556a1c9fd30159115d6c15226d0cac53a1d1a7/specification/ml/_types/Job.ts#L169-L172
type JobBlocked struct {
	Reason jobblockedreason.JobBlockedReason `json:"reason"`
	TaskId *TaskId                           `json:"task_id,omitempty"`
}

// JobBlockedBuilder holds JobBlocked struct and provides a builder API.
type JobBlockedBuilder struct {
	v *JobBlocked
}

// NewJobBlocked provides a builder for the JobBlocked struct.
func NewJobBlockedBuilder() *JobBlockedBuilder {
	r := JobBlockedBuilder{
		&JobBlocked{},
	}

	return &r
}

// Build finalize the chain and returns the JobBlocked struct
func (rb *JobBlockedBuilder) Build() JobBlocked {
	return *rb.v
}

func (rb *JobBlockedBuilder) Reason(reason jobblockedreason.JobBlockedReason) *JobBlockedBuilder {
	rb.v.Reason = reason
	return rb
}

func (rb *JobBlockedBuilder) TaskId(taskid *TaskIdBuilder) *JobBlockedBuilder {
	v := taskid.Build()
	rb.v.TaskId = &v
	return rb
}
