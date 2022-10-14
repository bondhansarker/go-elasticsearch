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


// Stops an existing, started rollup job.
package stopjob

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
	idMask = iota + 1
)

// ErrBuildPath is returned in case of missing parameters within the build of the request.
var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type StopJob struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	buf *gobytes.Buffer

	paramSet int

	id string
}

// NewStopJob type alias for index.
type NewStopJob func(id string) *StopJob

// NewStopJobFunc returns a new instance of StopJob with the provided transport.
// Used in the index of the library this allows to retrieve every apis in once place.
func NewStopJobFunc(tp elastictransport.Interface) NewStopJob {
	return func(id string) *StopJob {
		n := New(tp)

		n.Id(id)

		return n
	}
}

// Stops an existing, started rollup job.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/master/rollup-stop-job.html
func New(tp elastictransport.Interface) *StopJob {
	r := &StopJob{
		transport: tp,
		values:    make(url.Values),
		headers:   make(http.Header),
		buf:       gobytes.NewBuffer(nil),
	}

	return r
}

// HttpRequest returns the http.Request object built from the
// given parameters.
func (r *StopJob) HttpRequest(ctx context.Context) (*http.Request, error) {
	var path strings.Builder
	var method string
	var req *http.Request

	var err error

	r.path.Scheme = "http"

	switch {
	case r.paramSet == idMask:
		path.WriteString("/")
		path.WriteString("_rollup")
		path.WriteString("/")
		path.WriteString("job")
		path.WriteString("/")
		path.WriteString(url.PathEscape(r.id))
		path.WriteString("/")
		path.WriteString("_stop")

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
func (r StopJob) Do(ctx context.Context) (*http.Response, error) {
	req, err := r.HttpRequest(ctx)
	if err != nil {
		return nil, err
	}

	res, err := r.transport.Perform(req)
	if err != nil {
		return nil, fmt.Errorf("an error happened during the StopJob query execution: %w", err)
	}

	return res, nil
}

// IsSuccess allows to run a query with a context and retrieve the result as a boolean.
// This only exists for endpoints without a request payload and allows for quick control flow.
func (r StopJob) IsSuccess(ctx context.Context) (bool, error) {
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

// Header set a key, value pair in the StopJob headers map.
func (r *StopJob) Header(key, value string) *StopJob {
	r.headers.Set(key, value)

	return r
}

// Id The ID of the job to stop
// API Name: id
func (r *StopJob) Id(v string) *StopJob {
	r.paramSet |= idMask
	r.id = v

	return r
}

// Timeout Block for (at maximum) the specified duration while waiting for the job to
// stop.  Defaults to 30s.
// API name: timeout
func (r *StopJob) Timeout(value string) *StopJob {
	r.values.Set("timeout", value)

	return r
}

// WaitForCompletion True if the API should block until the job has fully stopped, false if should
// be executed async. Defaults to false.
// API name: wait_for_completion
func (r *StopJob) WaitForCompletion(b bool) *StopJob {
	r.values.Set("wait_for_completion", strconv.FormatBool(b))

	return r
}
