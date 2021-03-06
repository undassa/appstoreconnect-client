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

// SalesReportsApiService SalesReportsApi service
type SalesReportsApiService service

type ApiSalesReportsGetCollectionRequest struct {
	ctx _context.Context
	ApiService *SalesReportsApiService
	filterFrequency *[]string
	filterReportSubType *[]string
	filterReportType *[]string
	filterVendorNumber *[]string
	filterReportDate *[]string
	filterVersion *[]string
}

func (r ApiSalesReportsGetCollectionRequest) FilterFrequency(filterFrequency []string) ApiSalesReportsGetCollectionRequest {
	r.filterFrequency = &filterFrequency
	return r
}
func (r ApiSalesReportsGetCollectionRequest) FilterReportSubType(filterReportSubType []string) ApiSalesReportsGetCollectionRequest {
	r.filterReportSubType = &filterReportSubType
	return r
}
func (r ApiSalesReportsGetCollectionRequest) FilterReportType(filterReportType []string) ApiSalesReportsGetCollectionRequest {
	r.filterReportType = &filterReportType
	return r
}
func (r ApiSalesReportsGetCollectionRequest) FilterVendorNumber(filterVendorNumber []string) ApiSalesReportsGetCollectionRequest {
	r.filterVendorNumber = &filterVendorNumber
	return r
}
func (r ApiSalesReportsGetCollectionRequest) FilterReportDate(filterReportDate []string) ApiSalesReportsGetCollectionRequest {
	r.filterReportDate = &filterReportDate
	return r
}
func (r ApiSalesReportsGetCollectionRequest) FilterVersion(filterVersion []string) ApiSalesReportsGetCollectionRequest {
	r.filterVersion = &filterVersion
	return r
}

func (r ApiSalesReportsGetCollectionRequest) Execute() (*os.File, *_nethttp.Response, error) {
	return r.ApiService.SalesReportsGetCollectionExecute(r)
}

/*
 * SalesReportsGetCollection Method for SalesReportsGetCollection
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiSalesReportsGetCollectionRequest
 */
func (a *SalesReportsApiService) SalesReportsGetCollection(ctx _context.Context) ApiSalesReportsGetCollectionRequest {
	return ApiSalesReportsGetCollectionRequest{
		ApiService: a,
		ctx: ctx,
	}
}

/*
 * Execute executes the request
 * @return *os.File
 */
func (a *SalesReportsApiService) SalesReportsGetCollectionExecute(r ApiSalesReportsGetCollectionRequest) (*os.File, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  *os.File
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SalesReportsApiService.SalesReportsGetCollection")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/salesReports"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.filterFrequency == nil {
		return localVarReturnValue, nil, reportError("filterFrequency is required and must be specified")
	}
	if r.filterReportSubType == nil {
		return localVarReturnValue, nil, reportError("filterReportSubType is required and must be specified")
	}
	if r.filterReportType == nil {
		return localVarReturnValue, nil, reportError("filterReportType is required and must be specified")
	}
	if r.filterVendorNumber == nil {
		return localVarReturnValue, nil, reportError("filterVendorNumber is required and must be specified")
	}

	localVarQueryParams.Add("filter[frequency]", parameterToString(*r.filterFrequency, "csv"))
	if r.filterReportDate != nil {
		localVarQueryParams.Add("filter[reportDate]", parameterToString(*r.filterReportDate, "csv"))
	}
	localVarQueryParams.Add("filter[reportSubType]", parameterToString(*r.filterReportSubType, "csv"))
	localVarQueryParams.Add("filter[reportType]", parameterToString(*r.filterReportType, "csv"))
	localVarQueryParams.Add("filter[vendorNumber]", parameterToString(*r.filterVendorNumber, "csv"))
	if r.filterVersion != nil {
		localVarQueryParams.Add("filter[version]", parameterToString(*r.filterVersion, "csv"))
	}
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
