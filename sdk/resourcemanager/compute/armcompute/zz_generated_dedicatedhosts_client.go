//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armcompute

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// DedicatedHostsClient contains the methods for the DedicatedHosts group.
// Don't use this type directly, use NewDedicatedHostsClient() instead.
type DedicatedHostsClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewDedicatedHostsClient creates a new instance of DedicatedHostsClient with the specified values.
func NewDedicatedHostsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *DedicatedHostsClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Host) == 0 {
		cp.Host = arm.AzurePublicCloud
	}
	return &DedicatedHostsClient{subscriptionID: subscriptionID, ep: string(cp.Host), pl: armruntime.NewPipeline(module, version, credential, &cp)}
}

// BeginCreateOrUpdate - Create or update a dedicated host .
// If the operation fails it returns a generic error.
func (client *DedicatedHostsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, hostGroupName string, hostName string, parameters DedicatedHost, options *DedicatedHostsBeginCreateOrUpdateOptions) (DedicatedHostsCreateOrUpdatePollerResponse, error) {
	resp, err := client.createOrUpdate(ctx, resourceGroupName, hostGroupName, hostName, parameters, options)
	if err != nil {
		return DedicatedHostsCreateOrUpdatePollerResponse{}, err
	}
	result := DedicatedHostsCreateOrUpdatePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("DedicatedHostsClient.CreateOrUpdate", "", resp, client.pl, client.createOrUpdateHandleError)
	if err != nil {
		return DedicatedHostsCreateOrUpdatePollerResponse{}, err
	}
	result.Poller = &DedicatedHostsCreateOrUpdatePoller{
		pt: pt,
	}
	return result, nil
}

// CreateOrUpdate - Create or update a dedicated host .
// If the operation fails it returns a generic error.
func (client *DedicatedHostsClient) createOrUpdate(ctx context.Context, resourceGroupName string, hostGroupName string, hostName string, parameters DedicatedHost, options *DedicatedHostsBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, hostGroupName, hostName, parameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return nil, client.createOrUpdateHandleError(resp)
	}
	return resp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *DedicatedHostsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, hostGroupName string, hostName string, parameters DedicatedHost, options *DedicatedHostsBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/hostGroups/{hostGroupName}/hosts/{hostName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if hostGroupName == "" {
		return nil, errors.New("parameter hostGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{hostGroupName}", url.PathEscape(hostGroupName))
	if hostName == "" {
		return nil, errors.New("parameter hostName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{hostName}", url.PathEscape(hostName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-07-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// createOrUpdateHandleError handles the CreateOrUpdate error response.
func (client *DedicatedHostsClient) createOrUpdateHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}

// BeginDelete - Delete a dedicated host.
// If the operation fails it returns a generic error.
func (client *DedicatedHostsClient) BeginDelete(ctx context.Context, resourceGroupName string, hostGroupName string, hostName string, options *DedicatedHostsBeginDeleteOptions) (DedicatedHostsDeletePollerResponse, error) {
	resp, err := client.deleteOperation(ctx, resourceGroupName, hostGroupName, hostName, options)
	if err != nil {
		return DedicatedHostsDeletePollerResponse{}, err
	}
	result := DedicatedHostsDeletePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("DedicatedHostsClient.Delete", "", resp, client.pl, client.deleteHandleError)
	if err != nil {
		return DedicatedHostsDeletePollerResponse{}, err
	}
	result.Poller = &DedicatedHostsDeletePoller{
		pt: pt,
	}
	return result, nil
}

// Delete - Delete a dedicated host.
// If the operation fails it returns a generic error.
func (client *DedicatedHostsClient) deleteOperation(ctx context.Context, resourceGroupName string, hostGroupName string, hostName string, options *DedicatedHostsBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, hostGroupName, hostName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, client.deleteHandleError(resp)
	}
	return resp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *DedicatedHostsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, hostGroupName string, hostName string, options *DedicatedHostsBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/hostGroups/{hostGroupName}/hosts/{hostName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if hostGroupName == "" {
		return nil, errors.New("parameter hostGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{hostGroupName}", url.PathEscape(hostGroupName))
	if hostName == "" {
		return nil, errors.New("parameter hostName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{hostName}", url.PathEscape(hostName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-07-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	return req, nil
}

// deleteHandleError handles the Delete error response.
func (client *DedicatedHostsClient) deleteHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}

// Get - Retrieves information about a dedicated host.
// If the operation fails it returns a generic error.
func (client *DedicatedHostsClient) Get(ctx context.Context, resourceGroupName string, hostGroupName string, hostName string, options *DedicatedHostsGetOptions) (DedicatedHostsGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, hostGroupName, hostName, options)
	if err != nil {
		return DedicatedHostsGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return DedicatedHostsGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return DedicatedHostsGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *DedicatedHostsClient) getCreateRequest(ctx context.Context, resourceGroupName string, hostGroupName string, hostName string, options *DedicatedHostsGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/hostGroups/{hostGroupName}/hosts/{hostName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if hostGroupName == "" {
		return nil, errors.New("parameter hostGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{hostGroupName}", url.PathEscape(hostGroupName))
	if hostName == "" {
		return nil, errors.New("parameter hostName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{hostName}", url.PathEscape(hostName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Expand != nil {
		reqQP.Set("$expand", string(*options.Expand))
	}
	reqQP.Set("api-version", "2021-07-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *DedicatedHostsClient) getHandleResponse(resp *http.Response) (DedicatedHostsGetResponse, error) {
	result := DedicatedHostsGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.DedicatedHost); err != nil {
		return DedicatedHostsGetResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *DedicatedHostsClient) getHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}

// ListByHostGroup - Lists all of the dedicated hosts in the specified dedicated host group. Use the nextLink property in the response to get the next page
// of dedicated hosts.
// If the operation fails it returns a generic error.
func (client *DedicatedHostsClient) ListByHostGroup(resourceGroupName string, hostGroupName string, options *DedicatedHostsListByHostGroupOptions) *DedicatedHostsListByHostGroupPager {
	return &DedicatedHostsListByHostGroupPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listByHostGroupCreateRequest(ctx, resourceGroupName, hostGroupName, options)
		},
		advancer: func(ctx context.Context, resp DedicatedHostsListByHostGroupResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.DedicatedHostListResult.NextLink)
		},
	}
}

// listByHostGroupCreateRequest creates the ListByHostGroup request.
func (client *DedicatedHostsClient) listByHostGroupCreateRequest(ctx context.Context, resourceGroupName string, hostGroupName string, options *DedicatedHostsListByHostGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/hostGroups/{hostGroupName}/hosts"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if hostGroupName == "" {
		return nil, errors.New("parameter hostGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{hostGroupName}", url.PathEscape(hostGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-07-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByHostGroupHandleResponse handles the ListByHostGroup response.
func (client *DedicatedHostsClient) listByHostGroupHandleResponse(resp *http.Response) (DedicatedHostsListByHostGroupResponse, error) {
	result := DedicatedHostsListByHostGroupResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.DedicatedHostListResult); err != nil {
		return DedicatedHostsListByHostGroupResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// listByHostGroupHandleError handles the ListByHostGroup error response.
func (client *DedicatedHostsClient) listByHostGroupHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}

// BeginUpdate - Update an dedicated host .
// If the operation fails it returns a generic error.
func (client *DedicatedHostsClient) BeginUpdate(ctx context.Context, resourceGroupName string, hostGroupName string, hostName string, parameters DedicatedHostUpdate, options *DedicatedHostsBeginUpdateOptions) (DedicatedHostsUpdatePollerResponse, error) {
	resp, err := client.update(ctx, resourceGroupName, hostGroupName, hostName, parameters, options)
	if err != nil {
		return DedicatedHostsUpdatePollerResponse{}, err
	}
	result := DedicatedHostsUpdatePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("DedicatedHostsClient.Update", "", resp, client.pl, client.updateHandleError)
	if err != nil {
		return DedicatedHostsUpdatePollerResponse{}, err
	}
	result.Poller = &DedicatedHostsUpdatePoller{
		pt: pt,
	}
	return result, nil
}

// Update - Update an dedicated host .
// If the operation fails it returns a generic error.
func (client *DedicatedHostsClient) update(ctx context.Context, resourceGroupName string, hostGroupName string, hostName string, parameters DedicatedHostUpdate, options *DedicatedHostsBeginUpdateOptions) (*http.Response, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, hostGroupName, hostName, parameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return nil, client.updateHandleError(resp)
	}
	return resp, nil
}

// updateCreateRequest creates the Update request.
func (client *DedicatedHostsClient) updateCreateRequest(ctx context.Context, resourceGroupName string, hostGroupName string, hostName string, parameters DedicatedHostUpdate, options *DedicatedHostsBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/hostGroups/{hostGroupName}/hosts/{hostName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if hostGroupName == "" {
		return nil, errors.New("parameter hostGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{hostGroupName}", url.PathEscape(hostGroupName))
	if hostName == "" {
		return nil, errors.New("parameter hostName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{hostName}", url.PathEscape(hostName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-07-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// updateHandleError handles the Update error response.
func (client *DedicatedHostsClient) updateHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}
