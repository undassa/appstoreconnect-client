/*
 * App Store Connect API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.2
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package appstoreconnect

import (
	"bytes"
	_context "context"
	_ioutil "io/ioutil"
	_nethttp "net/http"
	_neturl "net/url"
	"strings"
)

// Linger please
var (
	_ _context.Context
)

// BetaAppReviewDetailsApiService BetaAppReviewDetailsApi service
type BetaAppReviewDetailsApiService service

type ApiBetaAppReviewDetailsAppGetToOneRelatedRequest struct {
	ctx _context.Context
	ApiService *BetaAppReviewDetailsApiService
	id string
	fieldsApps *[]string
}

func (r ApiBetaAppReviewDetailsAppGetToOneRelatedRequest) FieldsApps(fieldsApps []string) ApiBetaAppReviewDetailsAppGetToOneRelatedRequest {
	r.fieldsApps = &fieldsApps
	return r
}

func (r ApiBetaAppReviewDetailsAppGetToOneRelatedRequest) Execute() (AppResponse, *_nethttp.Response, error) {
	return r.ApiService.BetaAppReviewDetailsAppGetToOneRelatedExecute(r)
}

/*
 * BetaAppReviewDetailsAppGetToOneRelated Method for BetaAppReviewDetailsAppGetToOneRelated
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param id the id of the requested resource
 * @return ApiBetaAppReviewDetailsAppGetToOneRelatedRequest
 */
func (a *BetaAppReviewDetailsApiService) BetaAppReviewDetailsAppGetToOneRelated(ctx _context.Context, id string) ApiBetaAppReviewDetailsAppGetToOneRelatedRequest {
	return ApiBetaAppReviewDetailsAppGetToOneRelatedRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

/*
 * Execute executes the request
 * @return AppResponse
 */
func (a *BetaAppReviewDetailsApiService) BetaAppReviewDetailsAppGetToOneRelatedExecute(r ApiBetaAppReviewDetailsAppGetToOneRelatedRequest) (AppResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  AppResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BetaAppReviewDetailsApiService.BetaAppReviewDetailsAppGetToOneRelated")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/betaAppReviewDetails/{id}/app"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", _neturl.PathEscape(parameterToString(r.id, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.fieldsApps != nil {
		localVarQueryParams.Add("fields[apps]", parameterToString(*r.fieldsApps, "csv"))
	}
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
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiBetaAppReviewDetailsGetCollectionRequest struct {
	ctx _context.Context
	ApiService *BetaAppReviewDetailsApiService
	filterApp *[]string
	fieldsBetaAppReviewDetails *[]string
	limit *int32
	include *[]string
	fieldsApps *[]string
}

func (r ApiBetaAppReviewDetailsGetCollectionRequest) FilterApp(filterApp []string) ApiBetaAppReviewDetailsGetCollectionRequest {
	r.filterApp = &filterApp
	return r
}
func (r ApiBetaAppReviewDetailsGetCollectionRequest) FieldsBetaAppReviewDetails(fieldsBetaAppReviewDetails []string) ApiBetaAppReviewDetailsGetCollectionRequest {
	r.fieldsBetaAppReviewDetails = &fieldsBetaAppReviewDetails
	return r
}
func (r ApiBetaAppReviewDetailsGetCollectionRequest) Limit(limit int32) ApiBetaAppReviewDetailsGetCollectionRequest {
	r.limit = &limit
	return r
}
func (r ApiBetaAppReviewDetailsGetCollectionRequest) Include(include []string) ApiBetaAppReviewDetailsGetCollectionRequest {
	r.include = &include
	return r
}
func (r ApiBetaAppReviewDetailsGetCollectionRequest) FieldsApps(fieldsApps []string) ApiBetaAppReviewDetailsGetCollectionRequest {
	r.fieldsApps = &fieldsApps
	return r
}

func (r ApiBetaAppReviewDetailsGetCollectionRequest) Execute() (BetaAppReviewDetailsResponse, *_nethttp.Response, error) {
	return r.ApiService.BetaAppReviewDetailsGetCollectionExecute(r)
}

/*
 * BetaAppReviewDetailsGetCollection Method for BetaAppReviewDetailsGetCollection
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiBetaAppReviewDetailsGetCollectionRequest
 */
func (a *BetaAppReviewDetailsApiService) BetaAppReviewDetailsGetCollection(ctx _context.Context) ApiBetaAppReviewDetailsGetCollectionRequest {
	return ApiBetaAppReviewDetailsGetCollectionRequest{
		ApiService: a,
		ctx: ctx,
	}
}

/*
 * Execute executes the request
 * @return BetaAppReviewDetailsResponse
 */
func (a *BetaAppReviewDetailsApiService) BetaAppReviewDetailsGetCollectionExecute(r ApiBetaAppReviewDetailsGetCollectionRequest) (BetaAppReviewDetailsResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  BetaAppReviewDetailsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BetaAppReviewDetailsApiService.BetaAppReviewDetailsGetCollection")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/betaAppReviewDetails"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.filterApp == nil {
		return localVarReturnValue, nil, reportError("filterApp is required and must be specified")
	}

	localVarQueryParams.Add("filter[app]", parameterToString(*r.filterApp, "csv"))
	if r.fieldsBetaAppReviewDetails != nil {
		localVarQueryParams.Add("fields[betaAppReviewDetails]", parameterToString(*r.fieldsBetaAppReviewDetails, "csv"))
	}
	if r.limit != nil {
		localVarQueryParams.Add("limit", parameterToString(*r.limit, ""))
	}
	if r.include != nil {
		localVarQueryParams.Add("include", parameterToString(*r.include, "csv"))
	}
	if r.fieldsApps != nil {
		localVarQueryParams.Add("fields[apps]", parameterToString(*r.fieldsApps, "csv"))
	}
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
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiBetaAppReviewDetailsGetInstanceRequest struct {
	ctx _context.Context
	ApiService *BetaAppReviewDetailsApiService
	id string
	fieldsBetaAppReviewDetails *[]string
	include *[]string
	fieldsApps *[]string
}

func (r ApiBetaAppReviewDetailsGetInstanceRequest) FieldsBetaAppReviewDetails(fieldsBetaAppReviewDetails []string) ApiBetaAppReviewDetailsGetInstanceRequest {
	r.fieldsBetaAppReviewDetails = &fieldsBetaAppReviewDetails
	return r
}
func (r ApiBetaAppReviewDetailsGetInstanceRequest) Include(include []string) ApiBetaAppReviewDetailsGetInstanceRequest {
	r.include = &include
	return r
}
func (r ApiBetaAppReviewDetailsGetInstanceRequest) FieldsApps(fieldsApps []string) ApiBetaAppReviewDetailsGetInstanceRequest {
	r.fieldsApps = &fieldsApps
	return r
}

func (r ApiBetaAppReviewDetailsGetInstanceRequest) Execute() (BetaAppReviewDetailResponse, *_nethttp.Response, error) {
	return r.ApiService.BetaAppReviewDetailsGetInstanceExecute(r)
}

/*
 * BetaAppReviewDetailsGetInstance Method for BetaAppReviewDetailsGetInstance
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param id the id of the requested resource
 * @return ApiBetaAppReviewDetailsGetInstanceRequest
 */
func (a *BetaAppReviewDetailsApiService) BetaAppReviewDetailsGetInstance(ctx _context.Context, id string) ApiBetaAppReviewDetailsGetInstanceRequest {
	return ApiBetaAppReviewDetailsGetInstanceRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

/*
 * Execute executes the request
 * @return BetaAppReviewDetailResponse
 */
func (a *BetaAppReviewDetailsApiService) BetaAppReviewDetailsGetInstanceExecute(r ApiBetaAppReviewDetailsGetInstanceRequest) (BetaAppReviewDetailResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  BetaAppReviewDetailResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BetaAppReviewDetailsApiService.BetaAppReviewDetailsGetInstance")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/betaAppReviewDetails/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", _neturl.PathEscape(parameterToString(r.id, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.fieldsBetaAppReviewDetails != nil {
		localVarQueryParams.Add("fields[betaAppReviewDetails]", parameterToString(*r.fieldsBetaAppReviewDetails, "csv"))
	}
	if r.include != nil {
		localVarQueryParams.Add("include", parameterToString(*r.include, "csv"))
	}
	if r.fieldsApps != nil {
		localVarQueryParams.Add("fields[apps]", parameterToString(*r.fieldsApps, "csv"))
	}
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
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiBetaAppReviewDetailsUpdateInstanceRequest struct {
	ctx _context.Context
	ApiService *BetaAppReviewDetailsApiService
	id string
	betaAppReviewDetailUpdateRequest *BetaAppReviewDetailUpdateRequest
}

func (r ApiBetaAppReviewDetailsUpdateInstanceRequest) BetaAppReviewDetailUpdateRequest(betaAppReviewDetailUpdateRequest BetaAppReviewDetailUpdateRequest) ApiBetaAppReviewDetailsUpdateInstanceRequest {
	r.betaAppReviewDetailUpdateRequest = &betaAppReviewDetailUpdateRequest
	return r
}

func (r ApiBetaAppReviewDetailsUpdateInstanceRequest) Execute() (BetaAppReviewDetailResponse, *_nethttp.Response, error) {
	return r.ApiService.BetaAppReviewDetailsUpdateInstanceExecute(r)
}

/*
 * BetaAppReviewDetailsUpdateInstance Method for BetaAppReviewDetailsUpdateInstance
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param id the id of the requested resource
 * @return ApiBetaAppReviewDetailsUpdateInstanceRequest
 */
func (a *BetaAppReviewDetailsApiService) BetaAppReviewDetailsUpdateInstance(ctx _context.Context, id string) ApiBetaAppReviewDetailsUpdateInstanceRequest {
	return ApiBetaAppReviewDetailsUpdateInstanceRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

/*
 * Execute executes the request
 * @return BetaAppReviewDetailResponse
 */
func (a *BetaAppReviewDetailsApiService) BetaAppReviewDetailsUpdateInstanceExecute(r ApiBetaAppReviewDetailsUpdateInstanceRequest) (BetaAppReviewDetailResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPatch
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  BetaAppReviewDetailResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BetaAppReviewDetailsApiService.BetaAppReviewDetailsUpdateInstance")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/betaAppReviewDetails/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", _neturl.PathEscape(parameterToString(r.id, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.betaAppReviewDetailUpdateRequest == nil {
		return localVarReturnValue, nil, reportError("betaAppReviewDetailUpdateRequest is required and must be specified")
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
	localVarPostBody = r.betaAppReviewDetailUpdateRequest
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 409 {
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}