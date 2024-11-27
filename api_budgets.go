/*
MX Platform API

The MX Platform API is a powerful, fully-featured API designed to make aggregating and enhancing financial data easy and reliable. It can seamlessly connect your app or website to tens of thousands of financial institutions.

API version: 0.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package mxplatformgo

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)


// BudgetsAPIService BudgetsAPI service
type BudgetsAPIService service

type ApiUsersUserGuidBudgetsBudgetGuidDeleteRequest struct {
	ctx context.Context
	ApiService *BudgetsAPIService
	userGuid string
	budgetGuid string
}

func (r ApiUsersUserGuidBudgetsBudgetGuidDeleteRequest) Execute() (*http.Response, error) {
	return r.ApiService.UsersUserGuidBudgetsBudgetGuidDeleteExecute(r)
}

/*
UsersUserGuidBudgetsBudgetGuidDelete Delete a budget

Delete a budget.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param userGuid The unique identifier for the budget. Defined by MX.
 @param budgetGuid The unique identifier for the budget. Defined by MX.
 @return ApiUsersUserGuidBudgetsBudgetGuidDeleteRequest
*/
func (a *BudgetsAPIService) UsersUserGuidBudgetsBudgetGuidDelete(ctx context.Context, userGuid string, budgetGuid string) ApiUsersUserGuidBudgetsBudgetGuidDeleteRequest {
	return ApiUsersUserGuidBudgetsBudgetGuidDeleteRequest{
		ApiService: a,
		ctx: ctx,
		userGuid: userGuid,
		budgetGuid: budgetGuid,
	}
}

// Execute executes the request
func (a *BudgetsAPIService) UsersUserGuidBudgetsBudgetGuidDeleteExecute(r ApiUsersUserGuidBudgetsBudgetGuidDeleteRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BudgetsAPIService.UsersUserGuidBudgetsBudgetGuidDelete")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/users/{user_guid}/budgets/{budget_guid}"
	localVarPath = strings.Replace(localVarPath, "{"+"user_guid"+"}", url.PathEscape(parameterValueToString(r.userGuid, "userGuid")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"budget_guid"+"}", url.PathEscape(parameterValueToString(r.budgetGuid, "budgetGuid")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiUsersUserGuidBudgetsBudgetGuidGetRequest struct {
	ctx context.Context
	ApiService *BudgetsAPIService
	budgetGuid string
	userGuid string
}

func (r ApiUsersUserGuidBudgetsBudgetGuidGetRequest) Execute() (*BudgetResponseBody, *http.Response, error) {
	return r.ApiService.UsersUserGuidBudgetsBudgetGuidGetExecute(r)
}

/*
UsersUserGuidBudgetsBudgetGuidGet Read a specific budget

Read a specific budget.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param budgetGuid The unique identifier for the budget. Defined by MX.
 @param userGuid The unique identifier for the budget. Defined by MX.
 @return ApiUsersUserGuidBudgetsBudgetGuidGetRequest
*/
func (a *BudgetsAPIService) UsersUserGuidBudgetsBudgetGuidGet(ctx context.Context, budgetGuid string, userGuid string) ApiUsersUserGuidBudgetsBudgetGuidGetRequest {
	return ApiUsersUserGuidBudgetsBudgetGuidGetRequest{
		ApiService: a,
		ctx: ctx,
		budgetGuid: budgetGuid,
		userGuid: userGuid,
	}
}

// Execute executes the request
//  @return BudgetResponseBody
func (a *BudgetsAPIService) UsersUserGuidBudgetsBudgetGuidGetExecute(r ApiUsersUserGuidBudgetsBudgetGuidGetRequest) (*BudgetResponseBody, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *BudgetResponseBody
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BudgetsAPIService.UsersUserGuidBudgetsBudgetGuidGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/users/{user_guid}/budgets/{budget_guid}"
	localVarPath = strings.Replace(localVarPath, "{"+"budget_guid"+"}", url.PathEscape(parameterValueToString(r.budgetGuid, "budgetGuid")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"user_guid"+"}", url.PathEscape(parameterValueToString(r.userGuid, "userGuid")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiUsersUserGuidBudgetsBudgetGuidPutRequest struct {
	ctx context.Context
	ApiService *BudgetsAPIService
	userGuid string
	budgetGuid string
	budgetUpdateRequestBody *BudgetUpdateRequestBody
}

func (r ApiUsersUserGuidBudgetsBudgetGuidPutRequest) BudgetUpdateRequestBody(budgetUpdateRequestBody BudgetUpdateRequestBody) ApiUsersUserGuidBudgetsBudgetGuidPutRequest {
	r.budgetUpdateRequestBody = &budgetUpdateRequestBody
	return r
}

func (r ApiUsersUserGuidBudgetsBudgetGuidPutRequest) Execute() (*BudgetResponseBody, *http.Response, error) {
	return r.ApiService.UsersUserGuidBudgetsBudgetGuidPutExecute(r)
}

/*
UsersUserGuidBudgetsBudgetGuidPut Update a specific budget

Update a specific budget.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param userGuid The unique identifier for the budget. Defined by MX.
 @param budgetGuid The unique identifier for the budget. Defined by MX.
 @return ApiUsersUserGuidBudgetsBudgetGuidPutRequest
*/
func (a *BudgetsAPIService) UsersUserGuidBudgetsBudgetGuidPut(ctx context.Context, userGuid string, budgetGuid string) ApiUsersUserGuidBudgetsBudgetGuidPutRequest {
	return ApiUsersUserGuidBudgetsBudgetGuidPutRequest{
		ApiService: a,
		ctx: ctx,
		userGuid: userGuid,
		budgetGuid: budgetGuid,
	}
}

// Execute executes the request
//  @return BudgetResponseBody
func (a *BudgetsAPIService) UsersUserGuidBudgetsBudgetGuidPutExecute(r ApiUsersUserGuidBudgetsBudgetGuidPutRequest) (*BudgetResponseBody, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *BudgetResponseBody
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BudgetsAPIService.UsersUserGuidBudgetsBudgetGuidPut")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/users/{user_guid}/budgets/{budget_guid}"
	localVarPath = strings.Replace(localVarPath, "{"+"user_guid"+"}", url.PathEscape(parameterValueToString(r.userGuid, "userGuid")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"budget_guid"+"}", url.PathEscape(parameterValueToString(r.budgetGuid, "budgetGuid")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.budgetUpdateRequestBody
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiUsersUserGuidBudgetsGeneratePostRequest struct {
	ctx context.Context
	ApiService *BudgetsAPIService
	userGuid string
}

func (r ApiUsersUserGuidBudgetsGeneratePostRequest) Execute() (*BudgetResponseBody, *http.Response, error) {
	return r.ApiService.UsersUserGuidBudgetsGeneratePostExecute(r)
}

/*
UsersUserGuidBudgetsGeneratePost Auto-generate budgets

This endpoint will automatically create budgets for several categories based on existing transactions; these budgets are returned as an array. Specifically, budgets will only be generated if the `user` has at least one `transaction` in a given category during each of the two previous calendar months. For example, if the request is made on March 6, and there is at least one "Bills & Utilities" `transaction` in both January and February, a budget will be generated for "Bills & Utilities." If there are two "Bills & Utilities" transactions in February but none in January, no budget will be generated for that category. If budgets already exist for particular categories, new budgets will be generated and returned based on the available transactions. If one or more budgets remain unchanged, they will nevertheless be returned in the response. If no transaction data for the `user` meet the above criteria, a `422 Unprocessable Entity` error will be returned with status code 4221 along with the message, `There aren't enough transactions to automatically create any budgets`.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param userGuid The unique identifier for the user. Defined by MX.
 @return ApiUsersUserGuidBudgetsGeneratePostRequest
*/
func (a *BudgetsAPIService) UsersUserGuidBudgetsGeneratePost(ctx context.Context, userGuid string) ApiUsersUserGuidBudgetsGeneratePostRequest {
	return ApiUsersUserGuidBudgetsGeneratePostRequest{
		ApiService: a,
		ctx: ctx,
		userGuid: userGuid,
	}
}

// Execute executes the request
//  @return BudgetResponseBody
func (a *BudgetsAPIService) UsersUserGuidBudgetsGeneratePostExecute(r ApiUsersUserGuidBudgetsGeneratePostRequest) (*BudgetResponseBody, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *BudgetResponseBody
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BudgetsAPIService.UsersUserGuidBudgetsGeneratePost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/users/{user_guid}/budgets/generate"
	localVarPath = strings.Replace(localVarPath, "{"+"user_guid"+"}", url.PathEscape(parameterValueToString(r.userGuid, "userGuid")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiUsersUserGuidBudgetsGetRequest struct {
	ctx context.Context
	ApiService *BudgetsAPIService
	userGuid string
}

func (r ApiUsersUserGuidBudgetsGetRequest) Execute() (*BudgetResponseBody, *http.Response, error) {
	return r.ApiService.UsersUserGuidBudgetsGetExecute(r)
}

/*
UsersUserGuidBudgetsGet List all budgets

List all budgets

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param userGuid The unique identifier for the user. Defined by MX.
 @return ApiUsersUserGuidBudgetsGetRequest
*/
func (a *BudgetsAPIService) UsersUserGuidBudgetsGet(ctx context.Context, userGuid string) ApiUsersUserGuidBudgetsGetRequest {
	return ApiUsersUserGuidBudgetsGetRequest{
		ApiService: a,
		ctx: ctx,
		userGuid: userGuid,
	}
}

// Execute executes the request
//  @return BudgetResponseBody
func (a *BudgetsAPIService) UsersUserGuidBudgetsGetExecute(r ApiUsersUserGuidBudgetsGetRequest) (*BudgetResponseBody, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *BudgetResponseBody
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BudgetsAPIService.UsersUserGuidBudgetsGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/users/{user_guid}/budgets"
	localVarPath = strings.Replace(localVarPath, "{"+"user_guid"+"}", url.PathEscape(parameterValueToString(r.userGuid, "userGuid")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiUsersUserGuidBudgetsPostRequest struct {
	ctx context.Context
	ApiService *BudgetsAPIService
	userGuid string
	budgetCreateRequestBody *BudgetCreateRequestBody
}

func (r ApiUsersUserGuidBudgetsPostRequest) BudgetCreateRequestBody(budgetCreateRequestBody BudgetCreateRequestBody) ApiUsersUserGuidBudgetsPostRequest {
	r.budgetCreateRequestBody = &budgetCreateRequestBody
	return r
}

func (r ApiUsersUserGuidBudgetsPostRequest) Execute() (*BudgetResponseBody, *http.Response, error) {
	return r.ApiService.UsersUserGuidBudgetsPostExecute(r)
}

/*
UsersUserGuidBudgetsPost Create a budget

Create a budget. This endpoint accepts the optional `MX-Skip-Webhook` header and `skip_webhook` parameter. You cannot create a duplicate budget. For example, if you attempt to create a budget for "Gas", but that budget already exist, the request will fail. You can retrieve a list of all existing categories by using the List Categories endpoint.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param userGuid The unique identifier for the user. Defined by MX.
 @return ApiUsersUserGuidBudgetsPostRequest
*/
func (a *BudgetsAPIService) UsersUserGuidBudgetsPost(ctx context.Context, userGuid string) ApiUsersUserGuidBudgetsPostRequest {
	return ApiUsersUserGuidBudgetsPostRequest{
		ApiService: a,
		ctx: ctx,
		userGuid: userGuid,
	}
}

// Execute executes the request
//  @return BudgetResponseBody
func (a *BudgetsAPIService) UsersUserGuidBudgetsPostExecute(r ApiUsersUserGuidBudgetsPostRequest) (*BudgetResponseBody, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *BudgetResponseBody
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BudgetsAPIService.UsersUserGuidBudgetsPost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/users/{user_guid}/budgets"
	localVarPath = strings.Replace(localVarPath, "{"+"user_guid"+"}", url.PathEscape(parameterValueToString(r.userGuid, "userGuid")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.budgetCreateRequestBody == nil {
		return localVarReturnValue, nil, reportError("budgetCreateRequestBody is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.budgetCreateRequestBody
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}