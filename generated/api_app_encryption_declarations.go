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
	"strings"
)

// Linger please
var (
	_ _context.Context
)

// AppEncryptionDeclarationsApiService AppEncryptionDeclarationsApi service
type AppEncryptionDeclarationsApiService service

// AppEncryptionDeclarationsApiAppEncryptionDeclarationsAppGetToOneRelatedOpts Optional parameters for the method 'AppEncryptionDeclarationsAppGetToOneRelated'
type AppEncryptionDeclarationsApiAppEncryptionDeclarationsAppGetToOneRelatedOpts struct {
	FieldsApps optional.Interface
}

/*
AppEncryptionDeclarationsAppGetToOneRelated Method for AppEncryptionDeclarationsAppGetToOneRelated
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param id the id of the requested resource
 * @param optional nil or *AppEncryptionDeclarationsApiAppEncryptionDeclarationsAppGetToOneRelatedOpts - Optional Parameters:
 * @param "FieldsApps" (optional.Interface of []string) -  the fields to include for returned resources of type apps
@return AppResponse
*/
func (a *AppEncryptionDeclarationsApiService) AppEncryptionDeclarationsAppGetToOneRelated(ctx _context.Context, id string, localVarOptionals *AppEncryptionDeclarationsApiAppEncryptionDeclarationsAppGetToOneRelatedOpts) (AppResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  AppResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/v1/appEncryptionDeclarations/{id}/app"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", _neturl.PathEscape(parameterToString(id, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if localVarOptionals != nil && localVarOptionals.FieldsApps.IsSet() {
		localVarQueryParams.Add("fields[apps]", parameterToString(localVarOptionals.FieldsApps.Value(), "csv"))
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

/*
AppEncryptionDeclarationsBuildsCreateToManyRelationship Method for AppEncryptionDeclarationsBuildsCreateToManyRelationship
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param id the id of the requested resource
 * @param appEncryptionDeclarationBuildsLinkagesRequest List of related linkages
*/
func (a *AppEncryptionDeclarationsApiService) AppEncryptionDeclarationsBuildsCreateToManyRelationship(ctx _context.Context, id string, appEncryptionDeclarationBuildsLinkagesRequest AppEncryptionDeclarationBuildsLinkagesRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/v1/appEncryptionDeclarations/{id}/relationships/builds"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", _neturl.PathEscape(parameterToString(id, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

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
	localVarPostBody = &appEncryptionDeclarationBuildsLinkagesRequest
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(r)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 409 {
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

// AppEncryptionDeclarationsApiAppEncryptionDeclarationsGetCollectionOpts Optional parameters for the method 'AppEncryptionDeclarationsGetCollection'
type AppEncryptionDeclarationsApiAppEncryptionDeclarationsGetCollectionOpts struct {
	FilterPlatform                  optional.Interface
	FilterApp                       optional.Interface
	FilterBuilds                    optional.Interface
	FieldsAppEncryptionDeclarations optional.Interface
	Limit                           optional.Int32
	Include                         optional.Interface
	FieldsApps                      optional.Interface
}

/*
AppEncryptionDeclarationsGetCollection Method for AppEncryptionDeclarationsGetCollection
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param optional nil or *AppEncryptionDeclarationsApiAppEncryptionDeclarationsGetCollectionOpts - Optional Parameters:
 * @param "FilterPlatform" (optional.Interface of []string) -  filter by attribute 'platform'
 * @param "FilterApp" (optional.Interface of []string) -  filter by id(s) of related 'app'
 * @param "FilterBuilds" (optional.Interface of []string) -  filter by id(s) of related 'builds'
 * @param "FieldsAppEncryptionDeclarations" (optional.Interface of []string) -  the fields to include for returned resources of type appEncryptionDeclarations
 * @param "Limit" (optional.Int32) -  maximum resources per page
 * @param "Include" (optional.Interface of []string) -  comma-separated list of relationships to include
 * @param "FieldsApps" (optional.Interface of []string) -  the fields to include for returned resources of type apps
@return AppEncryptionDeclarationsResponse
*/
func (a *AppEncryptionDeclarationsApiService) AppEncryptionDeclarationsGetCollection(ctx _context.Context, localVarOptionals *AppEncryptionDeclarationsApiAppEncryptionDeclarationsGetCollectionOpts) (AppEncryptionDeclarationsResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  AppEncryptionDeclarationsResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/v1/appEncryptionDeclarations"
	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if localVarOptionals != nil && localVarOptionals.FilterPlatform.IsSet() {
		localVarQueryParams.Add("filter[platform]", parameterToString(localVarOptionals.FilterPlatform.Value(), "csv"))
	}
	if localVarOptionals != nil && localVarOptionals.FilterApp.IsSet() {
		localVarQueryParams.Add("filter[app]", parameterToString(localVarOptionals.FilterApp.Value(), "csv"))
	}
	if localVarOptionals != nil && localVarOptionals.FilterBuilds.IsSet() {
		localVarQueryParams.Add("filter[builds]", parameterToString(localVarOptionals.FilterBuilds.Value(), "csv"))
	}
	if localVarOptionals != nil && localVarOptionals.FieldsAppEncryptionDeclarations.IsSet() {
		localVarQueryParams.Add("fields[appEncryptionDeclarations]", parameterToString(localVarOptionals.FieldsAppEncryptionDeclarations.Value(), "csv"))
	}
	if localVarOptionals != nil && localVarOptionals.Limit.IsSet() {
		localVarQueryParams.Add("limit", parameterToString(localVarOptionals.Limit.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Include.IsSet() {
		localVarQueryParams.Add("include", parameterToString(localVarOptionals.Include.Value(), "csv"))
	}
	if localVarOptionals != nil && localVarOptionals.FieldsApps.IsSet() {
		localVarQueryParams.Add("fields[apps]", parameterToString(localVarOptionals.FieldsApps.Value(), "csv"))
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

// AppEncryptionDeclarationsApiAppEncryptionDeclarationsGetInstanceOpts Optional parameters for the method 'AppEncryptionDeclarationsGetInstance'
type AppEncryptionDeclarationsApiAppEncryptionDeclarationsGetInstanceOpts struct {
	FieldsAppEncryptionDeclarations optional.Interface
	Include                         optional.Interface
	FieldsApps                      optional.Interface
}

/*
AppEncryptionDeclarationsGetInstance Method for AppEncryptionDeclarationsGetInstance
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param id the id of the requested resource
 * @param optional nil or *AppEncryptionDeclarationsApiAppEncryptionDeclarationsGetInstanceOpts - Optional Parameters:
 * @param "FieldsAppEncryptionDeclarations" (optional.Interface of []string) -  the fields to include for returned resources of type appEncryptionDeclarations
 * @param "Include" (optional.Interface of []string) -  comma-separated list of relationships to include
 * @param "FieldsApps" (optional.Interface of []string) -  the fields to include for returned resources of type apps
@return AppEncryptionDeclarationResponse
*/
func (a *AppEncryptionDeclarationsApiService) AppEncryptionDeclarationsGetInstance(ctx _context.Context, id string, localVarOptionals *AppEncryptionDeclarationsApiAppEncryptionDeclarationsGetInstanceOpts) (AppEncryptionDeclarationResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  AppEncryptionDeclarationResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/v1/appEncryptionDeclarations/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", _neturl.PathEscape(parameterToString(id, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if localVarOptionals != nil && localVarOptionals.FieldsAppEncryptionDeclarations.IsSet() {
		localVarQueryParams.Add("fields[appEncryptionDeclarations]", parameterToString(localVarOptionals.FieldsAppEncryptionDeclarations.Value(), "csv"))
	}
	if localVarOptionals != nil && localVarOptionals.Include.IsSet() {
		localVarQueryParams.Add("include", parameterToString(localVarOptionals.Include.Value(), "csv"))
	}
	if localVarOptionals != nil && localVarOptionals.FieldsApps.IsSet() {
		localVarQueryParams.Add("fields[apps]", parameterToString(localVarOptionals.FieldsApps.Value(), "csv"))
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
