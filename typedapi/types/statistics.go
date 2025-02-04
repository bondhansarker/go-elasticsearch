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

// Statistics type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4ca0cc05d3ae3fa06c2cd7be91905b656a474334/specification/slm/_types/SnapshotLifecycle.ts#L51-L74
type Statistics struct {
	Policy                        *string   `json:"policy,omitempty"`
	RetentionDeletionTime         *Duration `json:"retention_deletion_time,omitempty"`
	RetentionDeletionTimeMillis   *int64    `json:"retention_deletion_time_millis,omitempty"`
	RetentionFailed               *int64    `json:"retention_failed,omitempty"`
	RetentionRuns                 *int64    `json:"retention_runs,omitempty"`
	RetentionTimedOut             *int64    `json:"retention_timed_out,omitempty"`
	TotalSnapshotDeletionFailures *int64    `json:"total_snapshot_deletion_failures,omitempty"`
	TotalSnapshotsDeleted         *int64    `json:"total_snapshots_deleted,omitempty"`
	TotalSnapshotsFailed          *int64    `json:"total_snapshots_failed,omitempty"`
	TotalSnapshotsTaken           *int64    `json:"total_snapshots_taken,omitempty"`
}

// NewStatistics returns a Statistics.
func NewStatistics() *Statistics {
	r := &Statistics{}

	return r
}
