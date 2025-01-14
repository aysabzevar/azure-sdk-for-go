//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armdevtestlabs

import (
	"context"
	"errors"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// EnvironmentsClient contains the methods for the Environments group.
// Don't use this type directly, use NewEnvironmentsClient() instead.
type EnvironmentsClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewEnvironmentsClient creates a new instance of EnvironmentsClient with the specified values.
func NewEnvironmentsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *EnvironmentsClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Host) == 0 {
		cp.Host = arm.AzurePublicCloud
	}
	return &EnvironmentsClient{subscriptionID: subscriptionID, ep: string(cp.Host), pl: armruntime.NewPipeline(module, version, credential, &cp)}
}

// BeginCreateOrUpdate - Create or replace an existing environment. This operation can take a while to complete.
// If the operation fails it returns the *CloudError error type.
func (client *EnvironmentsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, labName string, userName string, name string, dtlEnvironment DtlEnvironment, options *EnvironmentsBeginCreateOrUpdateOptions) (EnvironmentsCreateOrUpdatePollerResponse, error) {
	resp, err := client.createOrUpdate(ctx, resourceGroupName, labName, userName, name, dtlEnvironment, options)
	if err != nil {
		return EnvironmentsCreateOrUpdatePollerResponse{}, err
	}
	result := EnvironmentsCreateOrUpdatePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("EnvironmentsClient.CreateOrUpdate", "", resp, client.pl, client.createOrUpdateHandleError)
	if err != nil {
		return EnvironmentsCreateOrUpdatePollerResponse{}, err
	}
	result.Poller = &EnvironmentsCreateOrUpdatePoller{
		pt: pt,
	}
	return result, nil
}

// CreateOrUpdate - Create or replace an existing environment. This operation can take a while to complete.
// If the operation fails it returns the *CloudError error type.
func (client *EnvironmentsClient) createOrUpdate(ctx context.Context, resourceGroupName string, labName string, userName string, name string, dtlEnvironment DtlEnvironment, options *EnvironmentsBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, labName, userName, name, dtlEnvironment, options)
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
func (client *EnvironmentsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, labName string, userName string, name string, dtlEnvironment DtlEnvironment, options *EnvironmentsBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/users/{userName}/environments/{name}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if labName == "" {
		return nil, errors.New("parameter labName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{labName}", url.PathEscape(labName))
	if userName == "" {
		return nil, errors.New("parameter userName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{userName}", url.PathEscape(userName))
	if name == "" {
		return nil, errors.New("parameter name cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{name}", url.PathEscape(name))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-09-15")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, dtlEnvironment)
}

// createOrUpdateHandleError handles the CreateOrUpdate error response.
func (client *EnvironmentsClient) createOrUpdateHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// BeginDelete - Delete environment. This operation can take a while to complete.
// If the operation fails it returns the *CloudError error type.
func (client *EnvironmentsClient) BeginDelete(ctx context.Context, resourceGroupName string, labName string, userName string, name string, options *EnvironmentsBeginDeleteOptions) (EnvironmentsDeletePollerResponse, error) {
	resp, err := client.deleteOperation(ctx, resourceGroupName, labName, userName, name, options)
	if err != nil {
		return EnvironmentsDeletePollerResponse{}, err
	}
	result := EnvironmentsDeletePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("EnvironmentsClient.Delete", "", resp, client.pl, client.deleteHandleError)
	if err != nil {
		return EnvironmentsDeletePollerResponse{}, err
	}
	result.Poller = &EnvironmentsDeletePoller{
		pt: pt,
	}
	return result, nil
}

// Delete - Delete environment. This operation can take a while to complete.
// If the operation fails it returns the *CloudError error type.
func (client *EnvironmentsClient) deleteOperation(ctx context.Context, resourceGroupName string, labName string, userName string, name string, options *EnvironmentsBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, labName, userName, name, options)
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
func (client *EnvironmentsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, labName string, userName string, name string, options *EnvironmentsBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/users/{userName}/environments/{name}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if labName == "" {
		return nil, errors.New("parameter labName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{labName}", url.PathEscape(labName))
	if userName == "" {
		return nil, errors.New("parameter userName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{userName}", url.PathEscape(userName))
	if name == "" {
		return nil, errors.New("parameter name cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{name}", url.PathEscape(name))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-09-15")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// deleteHandleError handles the Delete error response.
func (client *EnvironmentsClient) deleteHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// Get - Get environment.
// If the operation fails it returns the *CloudError error type.
func (client *EnvironmentsClient) Get(ctx context.Context, resourceGroupName string, labName string, userName string, name string, options *EnvironmentsGetOptions) (EnvironmentsGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, labName, userName, name, options)
	if err != nil {
		return EnvironmentsGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return EnvironmentsGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return EnvironmentsGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *EnvironmentsClient) getCreateRequest(ctx context.Context, resourceGroupName string, labName string, userName string, name string, options *EnvironmentsGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/users/{userName}/environments/{name}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if labName == "" {
		return nil, errors.New("parameter labName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{labName}", url.PathEscape(labName))
	if userName == "" {
		return nil, errors.New("parameter userName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{userName}", url.PathEscape(userName))
	if name == "" {
		return nil, errors.New("parameter name cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{name}", url.PathEscape(name))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Expand != nil {
		reqQP.Set("$expand", *options.Expand)
	}
	reqQP.Set("api-version", "2018-09-15")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *EnvironmentsClient) getHandleResponse(resp *http.Response) (EnvironmentsGetResponse, error) {
	result := EnvironmentsGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.DtlEnvironment); err != nil {
		return EnvironmentsGetResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *EnvironmentsClient) getHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// List - List environments in a given user profile.
// If the operation fails it returns the *CloudError error type.
func (client *EnvironmentsClient) List(resourceGroupName string, labName string, userName string, options *EnvironmentsListOptions) *EnvironmentsListPager {
	return &EnvironmentsListPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listCreateRequest(ctx, resourceGroupName, labName, userName, options)
		},
		advancer: func(ctx context.Context, resp EnvironmentsListResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.DtlEnvironmentList.NextLink)
		},
	}
}

// listCreateRequest creates the List request.
func (client *EnvironmentsClient) listCreateRequest(ctx context.Context, resourceGroupName string, labName string, userName string, options *EnvironmentsListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/users/{userName}/environments"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if labName == "" {
		return nil, errors.New("parameter labName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{labName}", url.PathEscape(labName))
	if userName == "" {
		return nil, errors.New("parameter userName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{userName}", url.PathEscape(userName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Expand != nil {
		reqQP.Set("$expand", *options.Expand)
	}
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	if options != nil && options.Orderby != nil {
		reqQP.Set("$orderby", *options.Orderby)
	}
	reqQP.Set("api-version", "2018-09-15")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *EnvironmentsClient) listHandleResponse(resp *http.Response) (EnvironmentsListResponse, error) {
	result := EnvironmentsListResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.DtlEnvironmentList); err != nil {
		return EnvironmentsListResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// listHandleError handles the List error response.
func (client *EnvironmentsClient) listHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// Update - Allows modifying tags of environments. All other properties will be ignored.
// If the operation fails it returns the *CloudError error type.
func (client *EnvironmentsClient) Update(ctx context.Context, resourceGroupName string, labName string, userName string, name string, dtlEnvironment DtlEnvironmentFragment, options *EnvironmentsUpdateOptions) (EnvironmentsUpdateResponse, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, labName, userName, name, dtlEnvironment, options)
	if err != nil {
		return EnvironmentsUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return EnvironmentsUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return EnvironmentsUpdateResponse{}, client.updateHandleError(resp)
	}
	return client.updateHandleResponse(resp)
}

// updateCreateRequest creates the Update request.
func (client *EnvironmentsClient) updateCreateRequest(ctx context.Context, resourceGroupName string, labName string, userName string, name string, dtlEnvironment DtlEnvironmentFragment, options *EnvironmentsUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/users/{userName}/environments/{name}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if labName == "" {
		return nil, errors.New("parameter labName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{labName}", url.PathEscape(labName))
	if userName == "" {
		return nil, errors.New("parameter userName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{userName}", url.PathEscape(userName))
	if name == "" {
		return nil, errors.New("parameter name cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{name}", url.PathEscape(name))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-09-15")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, dtlEnvironment)
}

// updateHandleResponse handles the Update response.
func (client *EnvironmentsClient) updateHandleResponse(resp *http.Response) (EnvironmentsUpdateResponse, error) {
	result := EnvironmentsUpdateResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.DtlEnvironment); err != nil {
		return EnvironmentsUpdateResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// updateHandleError handles the Update error response.
func (client *EnvironmentsClient) updateHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}
