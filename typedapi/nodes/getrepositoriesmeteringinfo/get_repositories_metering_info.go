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


// Returns cluster repositories metering information.
package getrepositoriesmeteringinfo

import (
	gobytes "bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
)

const (
	nodeidMask = iota + 1
)

// ErrBuildPath is returned in case of missing parameters within the build of the request.
var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type GetRepositoriesMeteringInfo struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	buf *gobytes.Buffer

	paramSet int

	nodeid string
}

// NewGetRepositoriesMeteringInfo type alias for index.
type NewGetRepositoriesMeteringInfo func(nodeid string) *GetRepositoriesMeteringInfo

// NewGetRepositoriesMeteringInfoFunc returns a new instance of GetRepositoriesMeteringInfo with the provided transport.
// Used in the index of the library this allows to retrieve every apis in once place.
func NewGetRepositoriesMeteringInfoFunc(tp elastictransport.Interface) NewGetRepositoriesMeteringInfo {
	return func(nodeid string) *GetRepositoriesMeteringInfo {
		n := New(tp)

		n.NodeId(nodeid)

		return n
	}
}

// Returns cluster repositories metering information.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/get-repositories-metering-api.html
func New(tp elastictransport.Interface) *GetRepositoriesMeteringInfo {
	r := &GetRepositoriesMeteringInfo{
		transport: tp,
		values:    make(url.Values),
		headers:   make(http.Header),
		buf:       gobytes.NewBuffer(nil),
	}

	return r
}

// HttpRequest returns the http.Request object built from the
// given parameters.
func (r *GetRepositoriesMeteringInfo) HttpRequest(ctx context.Context) (*http.Request, error) {
	var path strings.Builder
	var method string
	var req *http.Request

	var err error

	r.path.Scheme = "http"

	switch {
	case r.paramSet == nodeidMask:
		path.WriteString("/")
		path.WriteString("_nodes")
		path.WriteString("/")
		path.WriteString(url.PathEscape(r.nodeid))
		path.WriteString("/")
		path.WriteString("_repositories_metering")

		method = http.MethodGet
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
func (r GetRepositoriesMeteringInfo) Do(ctx context.Context) (*http.Response, error) {
	req, err := r.HttpRequest(ctx)
	if err != nil {
		return nil, err
	}

	res, err := r.transport.Perform(req)
	if err != nil {
		return nil, fmt.Errorf("an error happened during the GetRepositoriesMeteringInfo query execution: %w", err)
	}

	return res, nil
}

// IsSuccess allows to run a query with a context and retrieve the result as a boolean.
// This only exists for endpoints without a request payload and allows for quick control flow.
func (r GetRepositoriesMeteringInfo) IsSuccess(ctx context.Context) (bool, error) {
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

// Header set a key, value pair in the GetRepositoriesMeteringInfo headers map.
func (r *GetRepositoriesMeteringInfo) Header(key, value string) *GetRepositoriesMeteringInfo {
	r.headers.Set(key, value)

	return r
}

// NodeId Comma-separated list of node IDs or names used to limit returned information.
// All the nodes selective options are explained
// [here](https://www.elastic.co/guide/en/elasticsearch/reference/current/cluster.html#cluster-nodes).
// API Name: nodeid
func (r *GetRepositoriesMeteringInfo) NodeId(v string) *GetRepositoriesMeteringInfo {
	r.paramSet |= nodeidMask
	r.nodeid = v

	return r
}
