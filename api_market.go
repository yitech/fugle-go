/*
FastAPI

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
)


// MarketAPIService MarketAPI service
type MarketAPIService service

type ApiGetCandlesApiV1HistoricalCandlesGetRequest struct {
	ctx context.Context
	ApiService *MarketAPIService
	symbol *string
	fromDate *string
	toDate *string
	resolution *string
}

func (r ApiGetCandlesApiV1HistoricalCandlesGetRequest) Symbol(symbol string) ApiGetCandlesApiV1HistoricalCandlesGetRequest {
	r.symbol = &symbol
	return r
}

func (r ApiGetCandlesApiV1HistoricalCandlesGetRequest) FromDate(fromDate string) ApiGetCandlesApiV1HistoricalCandlesGetRequest {
	r.fromDate = &fromDate
	return r
}

func (r ApiGetCandlesApiV1HistoricalCandlesGetRequest) ToDate(toDate string) ApiGetCandlesApiV1HistoricalCandlesGetRequest {
	r.toDate = &toDate
	return r
}

func (r ApiGetCandlesApiV1HistoricalCandlesGetRequest) Resolution(resolution string) ApiGetCandlesApiV1HistoricalCandlesGetRequest {
	r.resolution = &resolution
	return r
}

func (r ApiGetCandlesApiV1HistoricalCandlesGetRequest) Execute() (*KLinesResponse, *http.Response, error) {
	return r.ApiService.GetCandlesApiV1HistoricalCandlesGetExecute(r)
}

/*
GetCandlesApiV1HistoricalCandlesGet Get Candles

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetCandlesApiV1HistoricalCandlesGetRequest
*/
func (a *MarketAPIService) GetCandlesApiV1HistoricalCandlesGet(ctx context.Context) ApiGetCandlesApiV1HistoricalCandlesGetRequest {
	return ApiGetCandlesApiV1HistoricalCandlesGetRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return KLinesResponse
func (a *MarketAPIService) GetCandlesApiV1HistoricalCandlesGetExecute(r ApiGetCandlesApiV1HistoricalCandlesGetRequest) (*KLinesResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *KLinesResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MarketAPIService.GetCandlesApiV1HistoricalCandlesGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/historical/candles"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.symbol == nil {
		return localVarReturnValue, nil, reportError("symbol is required and must be specified")
	}
	if r.fromDate == nil {
		return localVarReturnValue, nil, reportError("fromDate is required and must be specified")
	}
	if r.toDate == nil {
		return localVarReturnValue, nil, reportError("toDate is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "symbol", r.symbol, "form", "")
	parameterAddToHeaderOrQuery(localVarQueryParams, "from_date", r.fromDate, "form", "")
	parameterAddToHeaderOrQuery(localVarQueryParams, "to_date", r.toDate, "form", "")
	if r.resolution != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "resolution", r.resolution, "form", "")
	} else {
		var defaultValue string = "D"
		r.resolution = &defaultValue
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
		if localVarHTTPResponse.StatusCode == 422 {
			var v HTTPValidationError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
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

type ApiGetQuoteApiV1IntradayQuoteGetRequest struct {
	ctx context.Context
	ApiService *MarketAPIService
	symbol *string
	type_ *string
}

func (r ApiGetQuoteApiV1IntradayQuoteGetRequest) Symbol(symbol string) ApiGetQuoteApiV1IntradayQuoteGetRequest {
	r.symbol = &symbol
	return r
}

func (r ApiGetQuoteApiV1IntradayQuoteGetRequest) Type_(type_ string) ApiGetQuoteApiV1IntradayQuoteGetRequest {
	r.type_ = &type_
	return r
}

func (r ApiGetQuoteApiV1IntradayQuoteGetRequest) Execute() (*QuoteResponse, *http.Response, error) {
	return r.ApiService.GetQuoteApiV1IntradayQuoteGetExecute(r)
}

/*
GetQuoteApiV1IntradayQuoteGet Get Quote

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetQuoteApiV1IntradayQuoteGetRequest
*/
func (a *MarketAPIService) GetQuoteApiV1IntradayQuoteGet(ctx context.Context) ApiGetQuoteApiV1IntradayQuoteGetRequest {
	return ApiGetQuoteApiV1IntradayQuoteGetRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return QuoteResponse
func (a *MarketAPIService) GetQuoteApiV1IntradayQuoteGetExecute(r ApiGetQuoteApiV1IntradayQuoteGetRequest) (*QuoteResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *QuoteResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MarketAPIService.GetQuoteApiV1IntradayQuoteGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/intraday/quote"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.symbol == nil {
		return localVarReturnValue, nil, reportError("symbol is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "symbol", r.symbol, "form", "")
	if r.type_ != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "type", r.type_, "form", "")
	} else {
		var defaultValue string = "EQUITY"
		r.type_ = &defaultValue
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
		if localVarHTTPResponse.StatusCode == 422 {
			var v HTTPValidationError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
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