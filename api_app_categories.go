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

// AppCategoriesApiService AppCategoriesApi service
type AppCategoriesApiService service

type ApiAppCategoriesGetCollectionRequest struct {
	ctx _context.Context
	ApiService *AppCategoriesApiService
	filterPlatforms *[]string
	existsParent *[]string
	fieldsAppCategories *[]string
	limit *int32
	include *[]string
	limitSubcategories *int32
}

func (r ApiAppCategoriesGetCollectionRequest) FilterPlatforms(filterPlatforms []string) ApiAppCategoriesGetCollectionRequest {
	r.filterPlatforms = &filterPlatforms
	return r
}
func (r ApiAppCategoriesGetCollectionRequest) ExistsParent(existsParent []string) ApiAppCategoriesGetCollectionRequest {
	r.existsParent = &existsParent
	return r
}
func (r ApiAppCategoriesGetCollectionRequest) FieldsAppCategories(fieldsAppCategories []string) ApiAppCategoriesGetCollectionRequest {
	r.fieldsAppCategories = &fieldsAppCategories
	return r
}
func (r ApiAppCategoriesGetCollectionRequest) Limit(limit int32) ApiAppCategoriesGetCollectionRequest {
	r.limit = &limit
	return r
}
func (r ApiAppCategoriesGetCollectionRequest) Include(include []string) ApiAppCategoriesGetCollectionRequest {
	r.include = &include
	return r
}
func (r ApiAppCategoriesGetCollectionRequest) LimitSubcategories(limitSubcategories int32) ApiAppCategoriesGetCollectionRequest {
	r.limitSubcategories = &limitSubcategories
	return r
}

func (r ApiAppCategoriesGetCollectionRequest) Execute() (AppCategoriesResponse, *_nethttp.Response, error) {
	return r.ApiService.AppCategoriesGetCollectionExecute(r)
}

/*
 * AppCategoriesGetCollection Method for AppCategoriesGetCollection
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiAppCategoriesGetCollectionRequest
 */
func (a *AppCategoriesApiService) AppCategoriesGetCollection(ctx _context.Context) ApiAppCategoriesGetCollectionRequest {
	return ApiAppCategoriesGetCollectionRequest{
		ApiService: a,
		ctx: ctx,
	}
}

/*
 * Execute executes the request
 * @return AppCategoriesResponse
 */
func (a *AppCategoriesApiService) AppCategoriesGetCollectionExecute(r ApiAppCategoriesGetCollectionRequest) (AppCategoriesResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  AppCategoriesResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AppCategoriesApiService.AppCategoriesGetCollection")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/appCategories"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.filterPlatforms != nil {
		localVarQueryParams.Add("filter[platforms]", parameterToString(*r.filterPlatforms, "csv"))
	}
	if r.existsParent != nil {
		localVarQueryParams.Add("exists[parent]", parameterToString(*r.existsParent, "csv"))
	}
	if r.fieldsAppCategories != nil {
		localVarQueryParams.Add("fields[appCategories]", parameterToString(*r.fieldsAppCategories, "csv"))
	}
	if r.limit != nil {
		localVarQueryParams.Add("limit", parameterToString(*r.limit, ""))
	}
	if r.include != nil {
		localVarQueryParams.Add("include", parameterToString(*r.include, "csv"))
	}
	if r.limitSubcategories != nil {
		localVarQueryParams.Add("limit[subcategories]", parameterToString(*r.limitSubcategories, ""))
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

type ApiAppCategoriesGetInstanceRequest struct {
	ctx _context.Context
	ApiService *AppCategoriesApiService
	id string
	fieldsAppCategories *[]string
	include *[]string
	limitSubcategories *int32
}

func (r ApiAppCategoriesGetInstanceRequest) FieldsAppCategories(fieldsAppCategories []string) ApiAppCategoriesGetInstanceRequest {
	r.fieldsAppCategories = &fieldsAppCategories
	return r
}
func (r ApiAppCategoriesGetInstanceRequest) Include(include []string) ApiAppCategoriesGetInstanceRequest {
	r.include = &include
	return r
}
func (r ApiAppCategoriesGetInstanceRequest) LimitSubcategories(limitSubcategories int32) ApiAppCategoriesGetInstanceRequest {
	r.limitSubcategories = &limitSubcategories
	return r
}

func (r ApiAppCategoriesGetInstanceRequest) Execute() (AppCategoryResponse, *_nethttp.Response, error) {
	return r.ApiService.AppCategoriesGetInstanceExecute(r)
}

/*
 * AppCategoriesGetInstance Method for AppCategoriesGetInstance
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param id the id of the requested resource
 * @return ApiAppCategoriesGetInstanceRequest
 */
func (a *AppCategoriesApiService) AppCategoriesGetInstance(ctx _context.Context, id string) ApiAppCategoriesGetInstanceRequest {
	return ApiAppCategoriesGetInstanceRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

/*
 * Execute executes the request
 * @return AppCategoryResponse
 */
func (a *AppCategoriesApiService) AppCategoriesGetInstanceExecute(r ApiAppCategoriesGetInstanceRequest) (AppCategoryResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  AppCategoryResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AppCategoriesApiService.AppCategoriesGetInstance")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/appCategories/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", _neturl.PathEscape(parameterToString(r.id, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.fieldsAppCategories != nil {
		localVarQueryParams.Add("fields[appCategories]", parameterToString(*r.fieldsAppCategories, "csv"))
	}
	if r.include != nil {
		localVarQueryParams.Add("include", parameterToString(*r.include, "csv"))
	}
	if r.limitSubcategories != nil {
		localVarQueryParams.Add("limit[subcategories]", parameterToString(*r.limitSubcategories, ""))
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

type ApiAppCategoriesParentGetToOneRelatedRequest struct {
	ctx _context.Context
	ApiService *AppCategoriesApiService
	id string
	fieldsAppCategories *[]string
}

func (r ApiAppCategoriesParentGetToOneRelatedRequest) FieldsAppCategories(fieldsAppCategories []string) ApiAppCategoriesParentGetToOneRelatedRequest {
	r.fieldsAppCategories = &fieldsAppCategories
	return r
}

func (r ApiAppCategoriesParentGetToOneRelatedRequest) Execute() (AppCategoryResponse, *_nethttp.Response, error) {
	return r.ApiService.AppCategoriesParentGetToOneRelatedExecute(r)
}

/*
 * AppCategoriesParentGetToOneRelated Method for AppCategoriesParentGetToOneRelated
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param id the id of the requested resource
 * @return ApiAppCategoriesParentGetToOneRelatedRequest
 */
func (a *AppCategoriesApiService) AppCategoriesParentGetToOneRelated(ctx _context.Context, id string) ApiAppCategoriesParentGetToOneRelatedRequest {
	return ApiAppCategoriesParentGetToOneRelatedRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

/*
 * Execute executes the request
 * @return AppCategoryResponse
 */
func (a *AppCategoriesApiService) AppCategoriesParentGetToOneRelatedExecute(r ApiAppCategoriesParentGetToOneRelatedRequest) (AppCategoryResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  AppCategoryResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AppCategoriesApiService.AppCategoriesParentGetToOneRelated")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/appCategories/{id}/parent"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", _neturl.PathEscape(parameterToString(r.id, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.fieldsAppCategories != nil {
		localVarQueryParams.Add("fields[appCategories]", parameterToString(*r.fieldsAppCategories, "csv"))
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

type ApiAppCategoriesSubcategoriesGetToManyRelatedRequest struct {
	ctx _context.Context
	ApiService *AppCategoriesApiService
	id string
	fieldsAppCategories *[]string
	limit *int32
}

func (r ApiAppCategoriesSubcategoriesGetToManyRelatedRequest) FieldsAppCategories(fieldsAppCategories []string) ApiAppCategoriesSubcategoriesGetToManyRelatedRequest {
	r.fieldsAppCategories = &fieldsAppCategories
	return r
}
func (r ApiAppCategoriesSubcategoriesGetToManyRelatedRequest) Limit(limit int32) ApiAppCategoriesSubcategoriesGetToManyRelatedRequest {
	r.limit = &limit
	return r
}

func (r ApiAppCategoriesSubcategoriesGetToManyRelatedRequest) Execute() (AppCategoriesResponse, *_nethttp.Response, error) {
	return r.ApiService.AppCategoriesSubcategoriesGetToManyRelatedExecute(r)
}

/*
 * AppCategoriesSubcategoriesGetToManyRelated Method for AppCategoriesSubcategoriesGetToManyRelated
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param id the id of the requested resource
 * @return ApiAppCategoriesSubcategoriesGetToManyRelatedRequest
 */
func (a *AppCategoriesApiService) AppCategoriesSubcategoriesGetToManyRelated(ctx _context.Context, id string) ApiAppCategoriesSubcategoriesGetToManyRelatedRequest {
	return ApiAppCategoriesSubcategoriesGetToManyRelatedRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

/*
 * Execute executes the request
 * @return AppCategoriesResponse
 */
func (a *AppCategoriesApiService) AppCategoriesSubcategoriesGetToManyRelatedExecute(r ApiAppCategoriesSubcategoriesGetToManyRelatedRequest) (AppCategoriesResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  AppCategoriesResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AppCategoriesApiService.AppCategoriesSubcategoriesGetToManyRelated")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/appCategories/{id}/subcategories"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", _neturl.PathEscape(parameterToString(r.id, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.fieldsAppCategories != nil {
		localVarQueryParams.Add("fields[appCategories]", parameterToString(*r.fieldsAppCategories, "csv"))
	}
	if r.limit != nil {
		localVarQueryParams.Add("limit", parameterToString(*r.limit, ""))
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
