/*
 * App Store Connect API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.2
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package appstoreopenapi

import (
	_context "context"
	"github.com/antihax/optional"
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

// SalesReportsApiSalesReportsGetCollectionOpts Optional parameters for the method 'SalesReportsGetCollection'
type SalesReportsApiSalesReportsGetCollectionOpts struct {
	FilterReportDate optional.Interface
	FilterVersion    optional.Interface
}

/*
SalesReportsGetCollection Method for SalesReportsGetCollection
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param filterFrequency filter by attribute 'frequency'
 * @param filterReportSubType filter by attribute 'reportSubType'
 * @param filterReportType filter by attribute 'reportType'
 * @param filterVendorNumber filter by attribute 'vendorNumber'
 * @param optional nil or *SalesReportsApiSalesReportsGetCollectionOpts - Optional Parameters:
 * @param "FilterReportDate" (optional.Interface of []string) -  filter by attribute 'reportDate'
 * @param "FilterVersion" (optional.Interface of []string) -  filter by attribute 'version'
@return *os.File
*/
func (a *SalesReportsApiService) SalesReportsGetCollection(ctx _context.Context, filterFrequency []string, filterReportSubType []string, filterReportType []string, filterVendorNumber []string, localVarOptionals *SalesReportsApiSalesReportsGetCollectionOpts) (*os.File, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  *os.File
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/v1/salesReports"
	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	localVarQueryParams.Add("filter[frequency]", parameterToString(filterFrequency, "csv"))
	if localVarOptionals != nil && localVarOptionals.FilterReportDate.IsSet() {
		localVarQueryParams.Add("filter[reportDate]", parameterToString(localVarOptionals.FilterReportDate.Value(), "csv"))
	}
	localVarQueryParams.Add("filter[reportSubType]", parameterToString(filterReportSubType, "csv"))
	localVarQueryParams.Add("filter[reportType]", parameterToString(filterReportType, "csv"))
	localVarQueryParams.Add("filter[vendorNumber]", parameterToString(filterVendorNumber, "csv"))
	if localVarOptionals != nil && localVarOptionals.FilterVersion.IsSet() {
		localVarQueryParams.Add("filter[version]", parameterToString(localVarOptionals.FilterVersion.Value(), "csv"))
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
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(r)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
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
