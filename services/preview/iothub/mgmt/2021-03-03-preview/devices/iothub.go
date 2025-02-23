package devices

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/validation"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// IotHubClient is the use this API to manage the IoT hubs in your Azure subscription.
type IotHubClient struct {
	BaseClient
}

// NewIotHubClient creates an instance of the IotHubClient client.
func NewIotHubClient(subscriptionID string) IotHubClient {
	return NewIotHubClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewIotHubClientWithBaseURI creates an instance of the IotHubClient client using a custom endpoint.  Use this when
// interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewIotHubClientWithBaseURI(baseURI string, subscriptionID string) IotHubClient {
	return IotHubClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// ManualFailover manually initiate a failover for the IoT Hub to its secondary region. To learn more, see
// https://aka.ms/manualfailover
// Parameters:
// iotHubName - name of the IoT hub to failover
// failoverInput - region to failover to. Must be the Azure paired region. Get the value from the secondary
// location in the locations property. To learn more, see https://aka.ms/manualfailover/region
// resourceGroupName - name of the resource group containing the IoT hub resource
func (client IotHubClient) ManualFailover(ctx context.Context, iotHubName string, failoverInput FailoverInput, resourceGroupName string) (result IotHubManualFailoverFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/IotHubClient.ManualFailover")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: failoverInput,
			Constraints: []validation.Constraint{{Target: "failoverInput.FailoverRegion", Name: validation.Null, Rule: true, Chain: nil}}}}); err != nil {
		return result, validation.NewError("devices.IotHubClient", "ManualFailover", err.Error())
	}

	req, err := client.ManualFailoverPreparer(ctx, iotHubName, failoverInput, resourceGroupName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "devices.IotHubClient", "ManualFailover", nil, "Failure preparing request")
		return
	}

	result, err = client.ManualFailoverSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "devices.IotHubClient", "ManualFailover", result.Response(), "Failure sending request")
		return
	}

	return
}

// ManualFailoverPreparer prepares the ManualFailover request.
func (client IotHubClient) ManualFailoverPreparer(ctx context.Context, iotHubName string, failoverInput FailoverInput, resourceGroupName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"iotHubName":        autorest.Encode("path", iotHubName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2021-03-03-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Devices/IotHubs/{iotHubName}/failover", pathParameters),
		autorest.WithJSON(failoverInput),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ManualFailoverSender sends the ManualFailover request. The method will close the
// http.Response Body if it receives an error.
func (client IotHubClient) ManualFailoverSender(req *http.Request) (future IotHubManualFailoverFuture, err error) {
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

// ManualFailoverResponder handles the response to the ManualFailover request. The method always
// closes the http.Response Body.
func (client IotHubClient) ManualFailoverResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByClosing())
	result.Response = resp
	return
}
