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


// Performs the force merge operation on one or more indices.
package forcemerge

import (
	gobytes "bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
)

const (
	indexMask = iota + 1
)

// ErrBuildPath is returned in case of missing parameters within the build of the request.
var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type Forcemerge struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	buf *gobytes.Buffer

	paramSet int

	index string
}

// NewForcemerge type alias for index.
type NewForcemerge func() *Forcemerge

// NewForcemergeFunc returns a new instance of Forcemerge with the provided transport.
// Used in the index of the library this allows to retrieve every apis in once place.
func NewForcemergeFunc(tp elastictransport.Interface) NewForcemerge {
	return func() *Forcemerge {
		n := New(tp)

		return n
	}
}

// Performs the force merge operation on one or more indices.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/master/indices-forcemerge.html
func New(tp elastictransport.Interface) *Forcemerge {
	r := &Forcemerge{
		transport: tp,
		values:    make(url.Values),
		headers:   make(http.Header),
		buf:       gobytes.NewBuffer(nil),
	}

	return r
}

// HttpRequest returns the http.Request object built from the
// given parameters.
func (r *Forcemerge) HttpRequest(ctx context.Context) (*http.Request, error) {
	var path strings.Builder
	var method string
	var req *http.Request

	var err error

	r.path.Scheme = "http"

	switch {
	case r.paramSet == 0:
		path.WriteString("/")
		path.WriteString("_forcemerge")

		method = http.MethodPost
	case r.paramSet == indexMask:
		path.WriteString("/")
		path.WriteString(url.PathEscape(r.index))
		path.WriteString("/")
		path.WriteString("_forcemerge")

		method = http.MethodPost
	}

	r.path.Path = path.String()
	r.path.RawQuery = r.values.Encode()

	if r.path.Path == "" {
		return nil, ErrBuildPath
	}

	if ctx != nil {
		req, err = http.NewRequestWithContext(ctx, method, r.path.String(), r.buf)
	} else {
		req, err = http.NewRequest(method, r.path.String(), r.buf)
	}

	req.Header = r.headers.Clone()

	if req.Header.Get("Accept") == "" {
		req.Header.Set("Accept", "application/vnd.elasticsearch+json;compatible-with=8")
	}

	if err != nil {
		return req, fmt.Errorf("could not build http.Request: %w", err)
	}

	return req, nil
}

// Do runs the http.Request through the provided transport.
func (r Forcemerge) Do(ctx context.Context) (*http.Response, error) {
	req, err := r.HttpRequest(ctx)
	if err != nil {
		return nil, err
	}

	res, err := r.transport.Perform(req)
	if err != nil {
		return nil, fmt.Errorf("an error happened during the Forcemerge query execution: %w", err)
	}

	return res, nil
}

// IsSuccess allows to run a query with a context and retrieve the result as a boolean.
// This only exists for endpoints without a request payload and allows for quick control flow.
func (r Forcemerge) IsSuccess(ctx context.Context) (bool, error) {
	res, err := r.Do(ctx)

	if err != nil {
		return false, err
	}
	io.Copy(ioutil.Discard, res.Body)
	err = res.Body.Close()
	if err != nil {
		return false, err
	}

	if res.StatusCode >= 200 && res.StatusCode < 300 {
		return true, nil
	}

	return false, nil
}

// Header set a key, value pair in the Forcemerge headers map.
func (r *Forcemerge) Header(key, value string) *Forcemerge {
	r.headers.Set(key, value)

	return r
}

// Index A comma-separated list of index names; use `_all` or empty string to perform
// the operation on all indices
// API Name: index
func (r *Forcemerge) Index(v string) *Forcemerge {
	r.paramSet |= indexMask
	r.index = v

	return r
}

// AllowNoIndices Whether to ignore if a wildcard indices expression resolves into no concrete
// indices. (This includes `_all` string or when no indices have been specified)
// API name: allow_no_indices
func (r *Forcemerge) AllowNoIndices(b bool) *Forcemerge {
	r.values.Set("allow_no_indices", strconv.FormatBool(b))

	return r
}

// ExpandWildcards Whether to expand wildcard expression to concrete indices that are open,
// closed or both.
// API name: expand_wildcards
func (r *Forcemerge) ExpandWildcards(value string) *Forcemerge {
	r.values.Set("expand_wildcards", value)

	return r
}

// Flush Specify whether the index should be flushed after performing the operation
// (default: true)
// API name: flush
func (r *Forcemerge) Flush(b bool) *Forcemerge {
	r.values.Set("flush", strconv.FormatBool(b))

	return r
}

// IgnoreUnavailable Whether specified concrete indices should be ignored when unavailable
// (missing or closed)
// API name: ignore_unavailable
func (r *Forcemerge) IgnoreUnavailable(b bool) *Forcemerge {
	r.values.Set("ignore_unavailable", strconv.FormatBool(b))

	return r
}

// MaxNumSegments The number of segments the index should be merged into (default: dynamic)
// API name: max_num_segments
func (r *Forcemerge) MaxNumSegments(value string) *Forcemerge {
	r.values.Set("max_num_segments", value)

	return r
}

// OnlyExpungeDeletes Specify whether the operation should only expunge deleted documents
// API name: only_expunge_deletes
func (r *Forcemerge) OnlyExpungeDeletes(b bool) *Forcemerge {
	r.values.Set("only_expunge_deletes", strconv.FormatBool(b))

	return r
}

// WaitForCompletion Should the request wait until the force merge is completed.
// API name: wait_for_completion
func (r *Forcemerge) WaitForCompletion(b bool) *Forcemerge {
	r.values.Set("wait_for_completion", strconv.FormatBool(b))

	return r
}
