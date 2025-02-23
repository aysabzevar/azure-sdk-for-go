//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armresourcehealth

import (
	"context"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"strings"
)

// ChildResourcesClient contains the methods for the ChildResources group.
// Don't use this type directly, use NewChildResourcesClient() instead.
type ChildResourcesClient struct {
	ep string
	pl runtime.Pipeline
}

// NewChildResourcesClient creates a new instance of ChildResourcesClient with the specified values.
func NewChildResourcesClient(credential azcore.TokenCredential, options *arm.ClientOptions) *ChildResourcesClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Host) == 0 {
		cp.Host = arm.AzurePublicCloud
	}
	return &ChildResourcesClient{ep: string(cp.Host), pl: armruntime.NewPipeline(module, version, credential, &cp)}
}

// List - Lists the all the children and its current health status for a parent resource. Use the nextLink property in the response to get the next page
// of children current health
// If the operation fails it returns the *ErrorResponse error type.
func (client *ChildResourcesClient) List(resourceURI string, options *ChildResourcesListOptions) *ChildResourcesListPager {
	return &ChildResourcesListPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listCreateRequest(ctx, resourceURI, options)
		},
		advancer: func(ctx context.Context, resp ChildResourcesListResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.AvailabilityStatusListResult.NextLink)
		},
	}
}

// listCreateRequest creates the List request.
func (client *ChildResourcesClient) listCreateRequest(ctx context.Context, resourceURI string, options *ChildResourcesListOptions) (*policy.Request, error) {
	urlPath := "/{resourceUri}/providers/Microsoft.ResourceHealth/childResources"
	urlPath = strings.ReplaceAll(urlPath, "{resourceUri}", resourceURI)
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2017-07-01")
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	if options != nil && options.Expand != nil {
		reqQP.Set("$expand", *options.Expand)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *ChildResourcesClient) listHandleResponse(resp *http.Response) (ChildResourcesListResponse, error) {
	result := ChildResourcesListResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.AvailabilityStatusListResult); err != nil {
		return ChildResourcesListResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// listHandleError handles the List error response.
func (client *ChildResourcesClient) listHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}
