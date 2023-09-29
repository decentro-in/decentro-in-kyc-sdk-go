/*
decentro-in-kyc

KYC & Onboarding

API version: 1.0.0
Contact: admin@decentro.tech
*/


package decentroinkyc

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
)


// UpgradedAadhaarXmlApiService UpgradedAadhaarXmlApi service
type UpgradedAadhaarXmlApiService service

type UpgradedAadhaarXmlApiGenerateOtpRequest struct {
	ctx context.Context
	ApiService *UpgradedAadhaarXmlApiService
	upgradedaadhaarxmlGenerateOtpRequest UpgradedaadhaarxmlGenerateOtpRequest
}

func (r UpgradedAadhaarXmlApiGenerateOtpRequest) Execute() (*UpgradedaadhaarxmlGenerateOtpResponse, *http.Response, error) {
	return r.ApiService.GenerateOtpExecute(r)
}

/*
GenerateOtp Generate OTP

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param upgradedaadhaarxmlGenerateOtpRequest
 @return UpgradedAadhaarXmlApiGenerateOtpRequest
*/
func (a *UpgradedAadhaarXmlApiService) GenerateOtp(upgradedaadhaarxmlGenerateOtpRequest UpgradedaadhaarxmlGenerateOtpRequest) UpgradedAadhaarXmlApiGenerateOtpRequest {
	return UpgradedAadhaarXmlApiGenerateOtpRequest{
		ApiService: a,
		ctx: a.client.cfg.Context,
		upgradedaadhaarxmlGenerateOtpRequest: upgradedaadhaarxmlGenerateOtpRequest,
	}
}

// Execute executes the request
//  @return UpgradedaadhaarxmlGenerateOtpResponse
func (a *UpgradedAadhaarXmlApiService) GenerateOtpExecute(r UpgradedAadhaarXmlApiGenerateOtpRequest) (*UpgradedaadhaarxmlGenerateOtpResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *UpgradedaadhaarxmlGenerateOtpResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UpgradedAadhaarXmlApiService.GenerateOtp")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2/kyc/aadhaar/otp"

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
	localVarPostBody = r.upgradedaadhaarxmlGenerateOtpRequest
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["client_id"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["client_id"] = key
			}
		}
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["client_secret"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["client_secret"] = key
			}
		}
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["module_secret"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["module_secret"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v UpgradedaadhaarxmlGenerateOtp400Response
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

type UpgradedAadhaarXmlApiValidateOtpRequest struct {
	ctx context.Context
	ApiService *UpgradedAadhaarXmlApiService
	upgradedaadhaarxmlValidateOtpRequest UpgradedaadhaarxmlValidateOtpRequest
}

func (r UpgradedAadhaarXmlApiValidateOtpRequest) Execute() (*UpgradedaadhaarxmlValidateOtpResponse, *http.Response, error) {
	return r.ApiService.ValidateOtpExecute(r)
}

/*
ValidateOtp Validate OTP

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param upgradedaadhaarxmlValidateOtpRequest
 @return UpgradedAadhaarXmlApiValidateOtpRequest
*/
func (a *UpgradedAadhaarXmlApiService) ValidateOtp(upgradedaadhaarxmlValidateOtpRequest UpgradedaadhaarxmlValidateOtpRequest) UpgradedAadhaarXmlApiValidateOtpRequest {
	return UpgradedAadhaarXmlApiValidateOtpRequest{
		ApiService: a,
		ctx: a.client.cfg.Context,
		upgradedaadhaarxmlValidateOtpRequest: upgradedaadhaarxmlValidateOtpRequest,
	}
}

// Execute executes the request
//  @return UpgradedaadhaarxmlValidateOtpResponse
func (a *UpgradedAadhaarXmlApiService) ValidateOtpExecute(r UpgradedAadhaarXmlApiValidateOtpRequest) (*UpgradedaadhaarxmlValidateOtpResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *UpgradedaadhaarxmlValidateOtpResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UpgradedAadhaarXmlApiService.ValidateOtp")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2/kyc/aadhaar/otp/validate"

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
	localVarPostBody = r.upgradedaadhaarxmlValidateOtpRequest
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v UpgradedaadhaarxmlValidateOtp400Response
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
