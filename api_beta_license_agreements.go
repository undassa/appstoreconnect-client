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

// BetaLicenseAgreementsApiService BetaLicenseAgreementsApi service
type BetaLicenseAgreementsApiService service

type ApiBetaLicenseAgreementsAppGetToOneRelatedRequest struct {
	ctx _context.Context
	ApiService *BetaLicenseAgreementsApiService
	id string
	fieldsApps *[]string
}

func (r ApiBetaLicenseAgreementsAppGetToOneRelatedRequest) FieldsApps(fieldsApps []string) ApiBetaLicenseAgreementsAppGetToOneRelatedRequest {
	r.fieldsApps = &fieldsApps
	return r
}

func (r ApiBetaLicenseAgreementsAppGetToOneRelatedRequest) Execute() (AppResponse, *_nethttp.Response, error) {
	return r.ApiService.BetaLicenseAgreementsAppGetToOneRelatedExecute(r)
}

/*
 * BetaLicenseAgreementsAppGetToOneRelated Method for BetaLicenseAgreementsAppGetToOneRelated
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param id the id of the requested resource
 * @return ApiBetaLicenseAgreementsAppGetToOneRelatedRequest
 */
func (a *BetaLicenseAgreementsApiService) BetaLicenseAgreementsAppGetToOneRelated(ctx _context.Context, id string) ApiBetaLicenseAgreementsAppGetToOneRelatedRequest {
	return ApiBetaLicenseAgreementsAppGetToOneRelatedRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

/*
 * Execute executes the request
 * @return AppResponse
 */
func (a *BetaLicenseAgreementsApiService) BetaLicenseAgreementsAppGetToOneRelatedExecute(r ApiBetaLicenseAgreementsAppGetToOneRelatedRequest) (AppResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  AppResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BetaLicenseAgreementsApiService.BetaLicenseAgreementsAppGetToOneRelated")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/betaLicenseAgreements/{id}/app"
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

type ApiBetaLicenseAgreementsGetCollectionRequest struct {
	ctx _context.Context
	ApiService *BetaLicenseAgreementsApiService
	filterApp *[]string
	fieldsBetaLicenseAgreements *[]string
	limit *int32
	include *[]string
	fieldsApps *[]string
}

func (r ApiBetaLicenseAgreementsGetCollectionRequest) FilterApp(filterApp []string) ApiBetaLicenseAgreementsGetCollectionRequest {
	r.filterApp = &filterApp
	return r
}
func (r ApiBetaLicenseAgreementsGetCollectionRequest) FieldsBetaLicenseAgreements(fieldsBetaLicenseAgreements []string) ApiBetaLicenseAgreementsGetCollectionRequest {
	r.fieldsBetaLicenseAgreements = &fieldsBetaLicenseAgreements
	return r
}
func (r ApiBetaLicenseAgreementsGetCollectionRequest) Limit(limit int32) ApiBetaLicenseAgreementsGetCollectionRequest {
	r.limit = &limit
	return r
}
func (r ApiBetaLicenseAgreementsGetCollectionRequest) Include(include []string) ApiBetaLicenseAgreementsGetCollectionRequest {
	r.include = &include
	return r
}
func (r ApiBetaLicenseAgreementsGetCollectionRequest) FieldsApps(fieldsApps []string) ApiBetaLicenseAgreementsGetCollectionRequest {
	r.fieldsApps = &fieldsApps
	return r
}

func (r ApiBetaLicenseAgreementsGetCollectionRequest) Execute() (BetaLicenseAgreementsResponse, *_nethttp.Response, error) {
	return r.ApiService.BetaLicenseAgreementsGetCollectionExecute(r)
}

/*
 * BetaLicenseAgreementsGetCollection Method for BetaLicenseAgreementsGetCollection
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiBetaLicenseAgreementsGetCollectionRequest
 */
func (a *BetaLicenseAgreementsApiService) BetaLicenseAgreementsGetCollection(ctx _context.Context) ApiBetaLicenseAgreementsGetCollectionRequest {
	return ApiBetaLicenseAgreementsGetCollectionRequest{
		ApiService: a,
		ctx: ctx,
	}
}

/*
 * Execute executes the request
 * @return BetaLicenseAgreementsResponse
 */
func (a *BetaLicenseAgreementsApiService) BetaLicenseAgreementsGetCollectionExecute(r ApiBetaLicenseAgreementsGetCollectionRequest) (BetaLicenseAgreementsResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  BetaLicenseAgreementsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BetaLicenseAgreementsApiService.BetaLicenseAgreementsGetCollection")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/betaLicenseAgreements"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.filterApp != nil {
		localVarQueryParams.Add("filter[app]", parameterToString(*r.filterApp, "csv"))
	}
	if r.fieldsBetaLicenseAgreements != nil {
		localVarQueryParams.Add("fields[betaLicenseAgreements]", parameterToString(*r.fieldsBetaLicenseAgreements, "csv"))
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

type ApiBetaLicenseAgreementsGetInstanceRequest struct {
	ctx _context.Context
	ApiService *BetaLicenseAgreementsApiService
	id string
	fieldsBetaLicenseAgreements *[]string
	include *[]string
	fieldsApps *[]string
}

func (r ApiBetaLicenseAgreementsGetInstanceRequest) FieldsBetaLicenseAgreements(fieldsBetaLicenseAgreements []string) ApiBetaLicenseAgreementsGetInstanceRequest {
	r.fieldsBetaLicenseAgreements = &fieldsBetaLicenseAgreements
	return r
}
func (r ApiBetaLicenseAgreementsGetInstanceRequest) Include(include []string) ApiBetaLicenseAgreementsGetInstanceRequest {
	r.include = &include
	return r
}
func (r ApiBetaLicenseAgreementsGetInstanceRequest) FieldsApps(fieldsApps []string) ApiBetaLicenseAgreementsGetInstanceRequest {
	r.fieldsApps = &fieldsApps
	return r
}

func (r ApiBetaLicenseAgreementsGetInstanceRequest) Execute() (BetaLicenseAgreementResponse, *_nethttp.Response, error) {
	return r.ApiService.BetaLicenseAgreementsGetInstanceExecute(r)
}

/*
 * BetaLicenseAgreementsGetInstance Method for BetaLicenseAgreementsGetInstance
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param id the id of the requested resource
 * @return ApiBetaLicenseAgreementsGetInstanceRequest
 */
func (a *BetaLicenseAgreementsApiService) BetaLicenseAgreementsGetInstance(ctx _context.Context, id string) ApiBetaLicenseAgreementsGetInstanceRequest {
	return ApiBetaLicenseAgreementsGetInstanceRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

/*
 * Execute executes the request
 * @return BetaLicenseAgreementResponse
 */
func (a *BetaLicenseAgreementsApiService) BetaLicenseAgreementsGetInstanceExecute(r ApiBetaLicenseAgreementsGetInstanceRequest) (BetaLicenseAgreementResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  BetaLicenseAgreementResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BetaLicenseAgreementsApiService.BetaLicenseAgreementsGetInstance")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/betaLicenseAgreements/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", _neturl.PathEscape(parameterToString(r.id, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.fieldsBetaLicenseAgreements != nil {
		localVarQueryParams.Add("fields[betaLicenseAgreements]", parameterToString(*r.fieldsBetaLicenseAgreements, "csv"))
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

type ApiBetaLicenseAgreementsUpdateInstanceRequest struct {
	ctx _context.Context
	ApiService *BetaLicenseAgreementsApiService
	id string
	betaLicenseAgreementUpdateRequest *BetaLicenseAgreementUpdateRequest
}

func (r ApiBetaLicenseAgreementsUpdateInstanceRequest) BetaLicenseAgreementUpdateRequest(betaLicenseAgreementUpdateRequest BetaLicenseAgreementUpdateRequest) ApiBetaLicenseAgreementsUpdateInstanceRequest {
	r.betaLicenseAgreementUpdateRequest = &betaLicenseAgreementUpdateRequest
	return r
}

func (r ApiBetaLicenseAgreementsUpdateInstanceRequest) Execute() (BetaLicenseAgreementResponse, *_nethttp.Response, error) {
	return r.ApiService.BetaLicenseAgreementsUpdateInstanceExecute(r)
}

/*
 * BetaLicenseAgreementsUpdateInstance Method for BetaLicenseAgreementsUpdateInstance
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param id the id of the requested resource
 * @return ApiBetaLicenseAgreementsUpdateInstanceRequest
 */
func (a *BetaLicenseAgreementsApiService) BetaLicenseAgreementsUpdateInstance(ctx _context.Context, id string) ApiBetaLicenseAgreementsUpdateInstanceRequest {
	return ApiBetaLicenseAgreementsUpdateInstanceRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

/*
 * Execute executes the request
 * @return BetaLicenseAgreementResponse
 */
func (a *BetaLicenseAgreementsApiService) BetaLicenseAgreementsUpdateInstanceExecute(r ApiBetaLicenseAgreementsUpdateInstanceRequest) (BetaLicenseAgreementResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPatch
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  BetaLicenseAgreementResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BetaLicenseAgreementsApiService.BetaLicenseAgreementsUpdateInstance")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/betaLicenseAgreements/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", _neturl.PathEscape(parameterToString(r.id, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.betaLicenseAgreementUpdateRequest == nil {
		return localVarReturnValue, nil, reportError("betaLicenseAgreementUpdateRequest is required and must be specified")
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
	localVarPostBody = r.betaLicenseAgreementUpdateRequest
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
