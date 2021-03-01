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
	"os"
)

// Linger please
var (
	_ _context.Context
)

// FinanceReportsApiService FinanceReportsApi service
type FinanceReportsApiService service

type ApiFinanceReportsGetCollectionRequest struct {
	ctx _context.Context
	ApiService *FinanceReportsApiService
	filterRegionCode *[]string
	filterReportDate *[]string
	filterReportType *[]string
	filterVendorNumber *[]string
}

func (r ApiFinanceReportsGetCollectionRequest) FilterRegionCode(filterRegionCode []string) ApiFinanceReportsGetCollectionRequest {
	r.filterRegionCode = &filterRegionCode
	return r
}
func (r ApiFinanceReportsGetCollectionRequest) FilterReportDate(filterReportDate []string) ApiFinanceReportsGetCollectionRequest {
	r.filterReportDate = &filterReportDate
	return r
}
func (r ApiFinanceReportsGetCollectionRequest) FilterReportType(filterReportType []string) ApiFinanceReportsGetCollectionRequest {
	r.filterReportType = &filterReportType
	return r
}
func (r ApiFinanceReportsGetCollectionRequest) FilterVendorNumber(filterVendorNumber []string) ApiFinanceReportsGetCollectionRequest {
	r.filterVendorNumber = &filterVendorNumber
	return r
}

func (r ApiFinanceReportsGetCollectionRequest) Execute() (*os.File, *_nethttp.Response, error) {
	return r.ApiService.FinanceReportsGetCollectionExecute(r)
}

/*
 * FinanceReportsGetCollection Method for FinanceReportsGetCollection
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiFinanceReportsGetCollectionRequest
 */
func (a *FinanceReportsApiService) FinanceReportsGetCollection(ctx _context.Context) ApiFinanceReportsGetCollectionRequest {
	return ApiFinanceReportsGetCollectionRequest{
		ApiService: a,
		ctx: ctx,
	}
}

/*
 * Execute executes the request
 * @return *os.File
 */
func (a *FinanceReportsApiService) FinanceReportsGetCollectionExecute(r ApiFinanceReportsGetCollectionRequest) (*os.File, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  *os.File
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FinanceReportsApiService.FinanceReportsGetCollection")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/financeReports"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.filterRegionCode == nil {
		return localVarReturnValue, nil, reportError("filterRegionCode is required and must be specified")
	}
	if r.filterReportDate == nil {
		return localVarReturnValue, nil, reportError("filterReportDate is required and must be specified")
	}
	if r.filterReportType == nil {
		return localVarReturnValue, nil, reportError("filterReportType is required and must be specified")
	}
	if r.filterVendorNumber == nil {
		return localVarReturnValue, nil, reportError("filterVendorNumber is required and must be specified")
	}

	localVarQueryParams.Add("filter[regionCode]", parameterToString(*r.filterRegionCode, "csv"))
	localVarQueryParams.Add("filter[reportDate]", parameterToString(*r.filterReportDate, "csv"))
	localVarQueryParams.Add("filter[reportType]", parameterToString(*r.filterReportType, "csv"))
	localVarQueryParams.Add("filter[vendorNumber]", parameterToString(*r.filterVendorNumber, "csv"))
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json", "gzip"}

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
