// Copyright 2023 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go_gapic. DO NOT EDIT.

package compute

import (
	"context"
	"fmt"
	"io/ioutil"
	"math"
	"net/http"
	"net/url"

	computepb "cloud.google.com/go/compute/apiv1/computepb"
	gax "github.com/googleapis/gax-go/v2"
	"google.golang.org/api/googleapi"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
	"google.golang.org/api/option/internaloption"
	httptransport "google.golang.org/api/transport/http"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

var newRegionsClientHook clientHook

// RegionsCallOptions contains the retry settings for each method of RegionsClient.
type RegionsCallOptions struct {
	Get  []gax.CallOption
	List []gax.CallOption
}

func defaultRegionsRESTCallOptions() *RegionsCallOptions {
	return &RegionsCallOptions{
		Get:  []gax.CallOption{},
		List: []gax.CallOption{},
	}
}

// internalRegionsClient is an interface that defines the methods available from Google Compute Engine API.
type internalRegionsClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	Get(context.Context, *computepb.GetRegionRequest, ...gax.CallOption) (*computepb.Region, error)
	List(context.Context, *computepb.ListRegionsRequest, ...gax.CallOption) *RegionIterator
}

// RegionsClient is a client for interacting with Google Compute Engine API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// The Regions API.
type RegionsClient struct {
	// The internal transport-dependent client.
	internalClient internalRegionsClient

	// The call options for this service.
	CallOptions *RegionsCallOptions
}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *RegionsClient) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *RegionsClient) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *RegionsClient) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// Get returns the specified Region resource. To decrease latency for this method, you can optionally omit any unneeded information from the response by using a field mask. This practice is especially recommended for unused quota information (the quotas field). To exclude one or more fields, set your request’s fields query parameter to only include the fields you need. For example, to only include the id and selfLink fields, add the query parameter ?fields=id,selfLink to your request.
func (c *RegionsClient) Get(ctx context.Context, req *computepb.GetRegionRequest, opts ...gax.CallOption) (*computepb.Region, error) {
	return c.internalClient.Get(ctx, req, opts...)
}

// List retrieves the list of region resources available to the specified project. To decrease latency for this method, you can optionally omit any unneeded information from the response by using a field mask. This practice is especially recommended for unused quota information (the items.quotas field). To exclude one or more fields, set your request’s fields query parameter to only include the fields you need. For example, to only include the id and selfLink fields, add the query parameter ?fields=id,selfLink to your request.
func (c *RegionsClient) List(ctx context.Context, req *computepb.ListRegionsRequest, opts ...gax.CallOption) *RegionIterator {
	return c.internalClient.List(ctx, req, opts...)
}

// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type regionsRESTClient struct {
	// The http endpoint to connect to.
	endpoint string

	// The http client.
	httpClient *http.Client

	// The x-goog-* metadata to be sent with each request.
	xGoogMetadata metadata.MD

	// Points back to the CallOptions field of the containing RegionsClient
	CallOptions **RegionsCallOptions
}

// NewRegionsRESTClient creates a new regions rest client.
//
// The Regions API.
func NewRegionsRESTClient(ctx context.Context, opts ...option.ClientOption) (*RegionsClient, error) {
	clientOpts := append(defaultRegionsRESTClientOptions(), opts...)
	httpClient, endpoint, err := httptransport.NewClient(ctx, clientOpts...)
	if err != nil {
		return nil, err
	}

	callOpts := defaultRegionsRESTCallOptions()
	c := &regionsRESTClient{
		endpoint:    endpoint,
		httpClient:  httpClient,
		CallOptions: &callOpts,
	}
	c.setGoogleClientInfo()

	return &RegionsClient{internalClient: c, CallOptions: callOpts}, nil
}

func defaultRegionsRESTClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("https://compute.googleapis.com"),
		internaloption.WithDefaultMTLSEndpoint("https://compute.mtls.googleapis.com"),
		internaloption.WithDefaultAudience("https://compute.googleapis.com/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
	}
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *regionsRESTClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", versionGo()}, keyval...)
	kv = append(kv, "gapic", getVersionClient(), "gax", gax.Version, "rest", "UNKNOWN")
	c.xGoogMetadata = metadata.Pairs("x-goog-api-client", gax.XGoogHeader(kv...))
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *regionsRESTClient) Close() error {
	// Replace httpClient with nil to force cleanup.
	c.httpClient = nil
	return nil
}

// Connection returns a connection to the API service.
//
// Deprecated: This method always returns nil.
func (c *regionsRESTClient) Connection() *grpc.ClientConn {
	return nil
}

// Get returns the specified Region resource. To decrease latency for this method, you can optionally omit any unneeded information from the response by using a field mask. This practice is especially recommended for unused quota information (the quotas field). To exclude one or more fields, set your request’s fields query parameter to only include the fields you need. For example, to only include the id and selfLink fields, add the query parameter ?fields=id,selfLink to your request.
func (c *regionsRESTClient) Get(ctx context.Context, req *computepb.GetRegionRequest, opts ...gax.CallOption) (*computepb.Region, error) {
	baseUrl, err := url.Parse(c.endpoint)
	if err != nil {
		return nil, err
	}
	baseUrl.Path += fmt.Sprintf("/compute/v1/projects/%v/regions/%v", req.GetProject(), req.GetRegion())

	// Build HTTP headers from client and context metadata.
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v&%s=%v", "project", url.QueryEscape(req.GetProject()), "region", url.QueryEscape(req.GetRegion())))

	headers := buildHeaders(ctx, c.xGoogMetadata, md, metadata.Pairs("Content-Type", "application/json"))
	opts = append((*c.CallOptions).Get[0:len((*c.CallOptions).Get):len((*c.CallOptions).Get)], opts...)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	resp := &computepb.Region{}
	e := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		if settings.Path != "" {
			baseUrl.Path = settings.Path
		}
		httpReq, err := http.NewRequest("GET", baseUrl.String(), nil)
		if err != nil {
			return err
		}
		httpReq = httpReq.WithContext(ctx)
		httpReq.Header = headers

		httpRsp, err := c.httpClient.Do(httpReq)
		if err != nil {
			return err
		}
		defer httpRsp.Body.Close()

		if err = googleapi.CheckResponse(httpRsp); err != nil {
			return err
		}

		buf, err := ioutil.ReadAll(httpRsp.Body)
		if err != nil {
			return err
		}

		if err := unm.Unmarshal(buf, resp); err != nil {
			return maybeUnknownEnum(err)
		}

		return nil
	}, opts...)
	if e != nil {
		return nil, e
	}
	return resp, nil
}

// List retrieves the list of region resources available to the specified project. To decrease latency for this method, you can optionally omit any unneeded information from the response by using a field mask. This practice is especially recommended for unused quota information (the items.quotas field). To exclude one or more fields, set your request’s fields query parameter to only include the fields you need. For example, to only include the id and selfLink fields, add the query parameter ?fields=id,selfLink to your request.
func (c *regionsRESTClient) List(ctx context.Context, req *computepb.ListRegionsRequest, opts ...gax.CallOption) *RegionIterator {
	it := &RegionIterator{}
	req = proto.Clone(req).(*computepb.ListRegionsRequest)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	it.InternalFetch = func(pageSize int, pageToken string) ([]*computepb.Region, string, error) {
		resp := &computepb.RegionList{}
		if pageToken != "" {
			req.PageToken = proto.String(pageToken)
		}
		if pageSize > math.MaxInt32 {
			req.MaxResults = proto.Uint32(math.MaxInt32)
		} else if pageSize != 0 {
			req.MaxResults = proto.Uint32(uint32(pageSize))
		}
		baseUrl, err := url.Parse(c.endpoint)
		if err != nil {
			return nil, "", err
		}
		baseUrl.Path += fmt.Sprintf("/compute/v1/projects/%v/regions", req.GetProject())

		params := url.Values{}
		if req != nil && req.Filter != nil {
			params.Add("filter", fmt.Sprintf("%v", req.GetFilter()))
		}
		if req != nil && req.MaxResults != nil {
			params.Add("maxResults", fmt.Sprintf("%v", req.GetMaxResults()))
		}
		if req != nil && req.OrderBy != nil {
			params.Add("orderBy", fmt.Sprintf("%v", req.GetOrderBy()))
		}
		if req != nil && req.PageToken != nil {
			params.Add("pageToken", fmt.Sprintf("%v", req.GetPageToken()))
		}
		if req != nil && req.ReturnPartialSuccess != nil {
			params.Add("returnPartialSuccess", fmt.Sprintf("%v", req.GetReturnPartialSuccess()))
		}

		baseUrl.RawQuery = params.Encode()

		// Build HTTP headers from client and context metadata.
		headers := buildHeaders(ctx, c.xGoogMetadata, metadata.Pairs("Content-Type", "application/json"))
		e := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
			if settings.Path != "" {
				baseUrl.Path = settings.Path
			}
			httpReq, err := http.NewRequest("GET", baseUrl.String(), nil)
			if err != nil {
				return err
			}
			httpReq.Header = headers

			httpRsp, err := c.httpClient.Do(httpReq)
			if err != nil {
				return err
			}
			defer httpRsp.Body.Close()

			if err = googleapi.CheckResponse(httpRsp); err != nil {
				return err
			}

			buf, err := ioutil.ReadAll(httpRsp.Body)
			if err != nil {
				return err
			}

			if err := unm.Unmarshal(buf, resp); err != nil {
				return maybeUnknownEnum(err)
			}

			return nil
		}, opts...)
		if e != nil {
			return nil, "", e
		}
		it.Response = resp
		return resp.GetItems(), resp.GetNextPageToken(), nil
	}

	fetch := func(pageSize int, pageToken string) (string, error) {
		items, nextPageToken, err := it.InternalFetch(pageSize, pageToken)
		if err != nil {
			return "", err
		}
		it.items = append(it.items, items...)
		return nextPageToken, nil
	}

	it.pageInfo, it.nextFunc = iterator.NewPageInfo(fetch, it.bufLen, it.takeBuf)
	it.pageInfo.MaxSize = int(req.GetMaxResults())
	it.pageInfo.Token = req.GetPageToken()

	return it
}

// RegionIterator manages a stream of *computepb.Region.
type RegionIterator struct {
	items    []*computepb.Region
	pageInfo *iterator.PageInfo
	nextFunc func() error

	// Response is the raw response for the current page.
	// It must be cast to the RPC response type.
	// Calling Next() or InternalFetch() updates this value.
	Response interface{}

	// InternalFetch is for use by the Google Cloud Libraries only.
	// It is not part of the stable interface of this package.
	//
	// InternalFetch returns results from a single call to the underlying RPC.
	// The number of results is no greater than pageSize.
	// If there are no more results, nextPageToken is empty and err is nil.
	InternalFetch func(pageSize int, pageToken string) (results []*computepb.Region, nextPageToken string, err error)
}

// PageInfo supports pagination. See the google.golang.org/api/iterator package for details.
func (it *RegionIterator) PageInfo() *iterator.PageInfo {
	return it.pageInfo
}

// Next returns the next result. Its second return value is iterator.Done if there are no more
// results. Once Next returns Done, all subsequent calls will return Done.
func (it *RegionIterator) Next() (*computepb.Region, error) {
	var item *computepb.Region
	if err := it.nextFunc(); err != nil {
		return item, err
	}
	item = it.items[0]
	it.items = it.items[1:]
	return item, nil
}

func (it *RegionIterator) bufLen() int {
	return len(it.items)
}

func (it *RegionIterator) takeBuf() interface{} {
	b := it.items
	it.items = nil
	return b
}