package network

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// VpnLinkConnectionsClient is the network Client
type VpnLinkConnectionsClient struct {
	BaseClient
}

// NewVpnLinkConnectionsClient creates an instance of the VpnLinkConnectionsClient client.
func NewVpnLinkConnectionsClient(subscriptionID string) VpnLinkConnectionsClient {
	return NewVpnLinkConnectionsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewVpnLinkConnectionsClientWithBaseURI creates an instance of the VpnLinkConnectionsClient client using a custom
// endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure
// stack).
func NewVpnLinkConnectionsClientWithBaseURI(baseURI string, subscriptionID string) VpnLinkConnectionsClient {
	return VpnLinkConnectionsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// GetIkeSas lists IKE Security Associations for Vpn Site Link Connection in the specified resource group.
// Parameters:
// resourceGroupName - the name of the resource group.
// gatewayName - the name of the gateway.
// connectionName - the name of the vpn connection.
// linkConnectionName - the name of the vpn link connection.
func (client VpnLinkConnectionsClient) GetIkeSas(ctx context.Context, resourceGroupName string, gatewayName string, connectionName string, linkConnectionName string) (result VpnLinkConnectionsGetIkeSasFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/VpnLinkConnectionsClient.GetIkeSas")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetIkeSasPreparer(ctx, resourceGroupName, gatewayName, connectionName, linkConnectionName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.VpnLinkConnectionsClient", "GetIkeSas", nil, "Failure preparing request")
		return
	}

	result, err = client.GetIkeSasSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.VpnLinkConnectionsClient", "GetIkeSas", result.Response(), "Failure sending request")
		return
	}

	return
}

// GetIkeSasPreparer prepares the GetIkeSas request.
func (client VpnLinkConnectionsClient) GetIkeSasPreparer(ctx context.Context, resourceGroupName string, gatewayName string, connectionName string, linkConnectionName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"connectionName":     autorest.Encode("path", connectionName),
		"gatewayName":        autorest.Encode("path", gatewayName),
		"linkConnectionName": autorest.Encode("path", linkConnectionName),
		"resourceGroupName":  autorest.Encode("path", resourceGroupName),
		"subscriptionId":     autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2021-03-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/vpnGateways/{gatewayName}/vpnConnections/{connectionName}/vpnLinkConnections/{linkConnectionName}/getikesas", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetIkeSasSender sends the GetIkeSas request. The method will close the
// http.Response Body if it receives an error.
func (client VpnLinkConnectionsClient) GetIkeSasSender(req *http.Request) (future VpnLinkConnectionsGetIkeSasFuture, err error) {
	var resp *http.Response
	future.FutureAPI = &azure.Future{}
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	var azf azure.Future
	azf, err = azure.NewFutureFromResponse(resp)
	future.FutureAPI = &azf
	future.Result = future.result
	return
}

// GetIkeSasResponder handles the response to the GetIkeSas request. The method always
// closes the http.Response Body.
func (client VpnLinkConnectionsClient) GetIkeSasResponder(resp *http.Response) (result String, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result.Value),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByVpnConnection retrieves all vpn site link connections for a particular virtual wan vpn gateway vpn connection.
// Parameters:
// resourceGroupName - the resource group name of the vpn gateway.
// gatewayName - the name of the gateway.
// connectionName - the name of the vpn connection.
func (client VpnLinkConnectionsClient) ListByVpnConnection(ctx context.Context, resourceGroupName string, gatewayName string, connectionName string) (result ListVpnSiteLinkConnectionsResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/VpnLinkConnectionsClient.ListByVpnConnection")
		defer func() {
			sc := -1
			if result.lvslcr.Response.Response != nil {
				sc = result.lvslcr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listByVpnConnectionNextResults
	req, err := client.ListByVpnConnectionPreparer(ctx, resourceGroupName, gatewayName, connectionName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.VpnLinkConnectionsClient", "ListByVpnConnection", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByVpnConnectionSender(req)
	if err != nil {
		result.lvslcr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "network.VpnLinkConnectionsClient", "ListByVpnConnection", resp, "Failure sending request")
		return
	}

	result.lvslcr, err = client.ListByVpnConnectionResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.VpnLinkConnectionsClient", "ListByVpnConnection", resp, "Failure responding to request")
		return
	}
	if result.lvslcr.hasNextLink() && result.lvslcr.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListByVpnConnectionPreparer prepares the ListByVpnConnection request.
func (client VpnLinkConnectionsClient) ListByVpnConnectionPreparer(ctx context.Context, resourceGroupName string, gatewayName string, connectionName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"connectionName":    autorest.Encode("path", connectionName),
		"gatewayName":       autorest.Encode("path", gatewayName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2021-03-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/vpnGateways/{gatewayName}/vpnConnections/{connectionName}/vpnLinkConnections", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByVpnConnectionSender sends the ListByVpnConnection request. The method will close the
// http.Response Body if it receives an error.
func (client VpnLinkConnectionsClient) ListByVpnConnectionSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListByVpnConnectionResponder handles the response to the ListByVpnConnection request. The method always
// closes the http.Response Body.
func (client VpnLinkConnectionsClient) ListByVpnConnectionResponder(resp *http.Response) (result ListVpnSiteLinkConnectionsResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listByVpnConnectionNextResults retrieves the next set of results, if any.
func (client VpnLinkConnectionsClient) listByVpnConnectionNextResults(ctx context.Context, lastResults ListVpnSiteLinkConnectionsResult) (result ListVpnSiteLinkConnectionsResult, err error) {
	req, err := lastResults.listVpnSiteLinkConnectionsResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "network.VpnLinkConnectionsClient", "listByVpnConnectionNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByVpnConnectionSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "network.VpnLinkConnectionsClient", "listByVpnConnectionNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByVpnConnectionResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.VpnLinkConnectionsClient", "listByVpnConnectionNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListByVpnConnectionComplete enumerates all values, automatically crossing page boundaries as required.
func (client VpnLinkConnectionsClient) ListByVpnConnectionComplete(ctx context.Context, resourceGroupName string, gatewayName string, connectionName string) (result ListVpnSiteLinkConnectionsResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/VpnLinkConnectionsClient.ListByVpnConnection")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListByVpnConnection(ctx, resourceGroupName, gatewayName, connectionName)
	return
}

// ResetConnection resets the VpnLink connection specified.
// Parameters:
// resourceGroupName - the name of the resource group.
// gatewayName - the name of the gateway.
// connectionName - the name of the vpn connection.
// linkConnectionName - the name of the vpn link connection.
func (client VpnLinkConnectionsClient) ResetConnection(ctx context.Context, resourceGroupName string, gatewayName string, connectionName string, linkConnectionName string) (result VpnLinkConnectionsResetConnectionFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/VpnLinkConnectionsClient.ResetConnection")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.ResetConnectionPreparer(ctx, resourceGroupName, gatewayName, connectionName, linkConnectionName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.VpnLinkConnectionsClient", "ResetConnection", nil, "Failure preparing request")
		return
	}

	result, err = client.ResetConnectionSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.VpnLinkConnectionsClient", "ResetConnection", result.Response(), "Failure sending request")
		return
	}

	return
}

// ResetConnectionPreparer prepares the ResetConnection request.
func (client VpnLinkConnectionsClient) ResetConnectionPreparer(ctx context.Context, resourceGroupName string, gatewayName string, connectionName string, linkConnectionName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"connectionName":     autorest.Encode("path", connectionName),
		"gatewayName":        autorest.Encode("path", gatewayName),
		"linkConnectionName": autorest.Encode("path", linkConnectionName),
		"resourceGroupName":  autorest.Encode("path", resourceGroupName),
		"subscriptionId":     autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2021-03-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/vpnGateways/{gatewayName}/vpnConnections/{connectionName}/vpnLinkConnections/{linkConnectionName}/resetconnection", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ResetConnectionSender sends the ResetConnection request. The method will close the
// http.Response Body if it receives an error.
func (client VpnLinkConnectionsClient) ResetConnectionSender(req *http.Request) (future VpnLinkConnectionsResetConnectionFuture, err error) {
	var resp *http.Response
	future.FutureAPI = &azure.Future{}
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	var azf azure.Future
	azf, err = azure.NewFutureFromResponse(resp)
	future.FutureAPI = &azf
	future.Result = future.result
	return
}

// ResetConnectionResponder handles the response to the ResetConnection request. The method always
// closes the http.Response Body.
func (client VpnLinkConnectionsClient) ResetConnectionResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByClosing())
	result.Response = resp
	return
}
