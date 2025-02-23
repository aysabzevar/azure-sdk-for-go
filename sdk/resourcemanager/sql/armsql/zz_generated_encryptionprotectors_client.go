//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armsql

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

// EncryptionProtectorsClient contains the methods for the EncryptionProtectors group.
// Don't use this type directly, use NewEncryptionProtectorsClient() instead.
type EncryptionProtectorsClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewEncryptionProtectorsClient creates a new instance of EncryptionProtectorsClient with the specified values.
func NewEncryptionProtectorsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *EncryptionProtectorsClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Host) == 0 {
		cp.Host = arm.AzurePublicCloud
	}
	return &EncryptionProtectorsClient{subscriptionID: subscriptionID, ep: string(cp.Host), pl: armruntime.NewPipeline(module, version, credential, &cp)}
}

// BeginCreateOrUpdate - Updates an existing encryption protector.
// If the operation fails it returns a generic error.
func (client *EncryptionProtectorsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, serverName string, encryptionProtectorName EncryptionProtectorName, parameters EncryptionProtector, options *EncryptionProtectorsBeginCreateOrUpdateOptions) (EncryptionProtectorsCreateOrUpdatePollerResponse, error) {
	resp, err := client.createOrUpdate(ctx, resourceGroupName, serverName, encryptionProtectorName, parameters, options)
	if err != nil {
		return EncryptionProtectorsCreateOrUpdatePollerResponse{}, err
	}
	result := EncryptionProtectorsCreateOrUpdatePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("EncryptionProtectorsClient.CreateOrUpdate", "", resp, client.pl, client.createOrUpdateHandleError)
	if err != nil {
		return EncryptionProtectorsCreateOrUpdatePollerResponse{}, err
	}
	result.Poller = &EncryptionProtectorsCreateOrUpdatePoller{
		pt: pt,
	}
	return result, nil
}

// CreateOrUpdate - Updates an existing encryption protector.
// If the operation fails it returns a generic error.
func (client *EncryptionProtectorsClient) createOrUpdate(ctx context.Context, resourceGroupName string, serverName string, encryptionProtectorName EncryptionProtectorName, parameters EncryptionProtector, options *EncryptionProtectorsBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, serverName, encryptionProtectorName, parameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, client.createOrUpdateHandleError(resp)
	}
	return resp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *EncryptionProtectorsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, serverName string, encryptionProtectorName EncryptionProtectorName, parameters EncryptionProtector, options *EncryptionProtectorsBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/encryptionProtector/{encryptionProtectorName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serverName == "" {
		return nil, errors.New("parameter serverName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serverName}", url.PathEscape(serverName))
	if encryptionProtectorName == "" {
		return nil, errors.New("parameter encryptionProtectorName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{encryptionProtectorName}", url.PathEscape(string(encryptionProtectorName)))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// createOrUpdateHandleError handles the CreateOrUpdate error response.
func (client *EncryptionProtectorsClient) createOrUpdateHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}

// Get - Gets a server encryption protector.
// If the operation fails it returns a generic error.
func (client *EncryptionProtectorsClient) Get(ctx context.Context, resourceGroupName string, serverName string, encryptionProtectorName EncryptionProtectorName, options *EncryptionProtectorsGetOptions) (EncryptionProtectorsGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, serverName, encryptionProtectorName, options)
	if err != nil {
		return EncryptionProtectorsGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return EncryptionProtectorsGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return EncryptionProtectorsGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *EncryptionProtectorsClient) getCreateRequest(ctx context.Context, resourceGroupName string, serverName string, encryptionProtectorName EncryptionProtectorName, options *EncryptionProtectorsGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/encryptionProtector/{encryptionProtectorName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serverName == "" {
		return nil, errors.New("parameter serverName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serverName}", url.PathEscape(serverName))
	if encryptionProtectorName == "" {
		return nil, errors.New("parameter encryptionProtectorName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{encryptionProtectorName}", url.PathEscape(string(encryptionProtectorName)))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *EncryptionProtectorsClient) getHandleResponse(resp *http.Response) (EncryptionProtectorsGetResponse, error) {
	result := EncryptionProtectorsGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.EncryptionProtector); err != nil {
		return EncryptionProtectorsGetResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *EncryptionProtectorsClient) getHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}

// ListByServer - Gets a list of server encryption protectors
// If the operation fails it returns a generic error.
func (client *EncryptionProtectorsClient) ListByServer(resourceGroupName string, serverName string, options *EncryptionProtectorsListByServerOptions) *EncryptionProtectorsListByServerPager {
	return &EncryptionProtectorsListByServerPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listByServerCreateRequest(ctx, resourceGroupName, serverName, options)
		},
		advancer: func(ctx context.Context, resp EncryptionProtectorsListByServerResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.EncryptionProtectorListResult.NextLink)
		},
	}
}

// listByServerCreateRequest creates the ListByServer request.
func (client *EncryptionProtectorsClient) listByServerCreateRequest(ctx context.Context, resourceGroupName string, serverName string, options *EncryptionProtectorsListByServerOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/encryptionProtector"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serverName == "" {
		return nil, errors.New("parameter serverName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serverName}", url.PathEscape(serverName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByServerHandleResponse handles the ListByServer response.
func (client *EncryptionProtectorsClient) listByServerHandleResponse(resp *http.Response) (EncryptionProtectorsListByServerResponse, error) {
	result := EncryptionProtectorsListByServerResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.EncryptionProtectorListResult); err != nil {
		return EncryptionProtectorsListByServerResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// listByServerHandleError handles the ListByServer error response.
func (client *EncryptionProtectorsClient) listByServerHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}

// BeginRevalidate - Revalidates an existing encryption protector.
// If the operation fails it returns a generic error.
func (client *EncryptionProtectorsClient) BeginRevalidate(ctx context.Context, resourceGroupName string, serverName string, encryptionProtectorName EncryptionProtectorName, options *EncryptionProtectorsBeginRevalidateOptions) (EncryptionProtectorsRevalidatePollerResponse, error) {
	resp, err := client.revalidate(ctx, resourceGroupName, serverName, encryptionProtectorName, options)
	if err != nil {
		return EncryptionProtectorsRevalidatePollerResponse{}, err
	}
	result := EncryptionProtectorsRevalidatePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("EncryptionProtectorsClient.Revalidate", "", resp, client.pl, client.revalidateHandleError)
	if err != nil {
		return EncryptionProtectorsRevalidatePollerResponse{}, err
	}
	result.Poller = &EncryptionProtectorsRevalidatePoller{
		pt: pt,
	}
	return result, nil
}

// Revalidate - Revalidates an existing encryption protector.
// If the operation fails it returns a generic error.
func (client *EncryptionProtectorsClient) revalidate(ctx context.Context, resourceGroupName string, serverName string, encryptionProtectorName EncryptionProtectorName, options *EncryptionProtectorsBeginRevalidateOptions) (*http.Response, error) {
	req, err := client.revalidateCreateRequest(ctx, resourceGroupName, serverName, encryptionProtectorName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, client.revalidateHandleError(resp)
	}
	return resp, nil
}

// revalidateCreateRequest creates the Revalidate request.
func (client *EncryptionProtectorsClient) revalidateCreateRequest(ctx context.Context, resourceGroupName string, serverName string, encryptionProtectorName EncryptionProtectorName, options *EncryptionProtectorsBeginRevalidateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/encryptionProtector/{encryptionProtectorName}/revalidate"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serverName == "" {
		return nil, errors.New("parameter serverName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serverName}", url.PathEscape(serverName))
	if encryptionProtectorName == "" {
		return nil, errors.New("parameter encryptionProtectorName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{encryptionProtectorName}", url.PathEscape(string(encryptionProtectorName)))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	return req, nil
}

// revalidateHandleError handles the Revalidate error response.
func (client *EncryptionProtectorsClient) revalidateHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}
