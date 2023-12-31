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
	"os"
)


// KYCApiService KYCApi service
type KYCApiService service

type KYCApiCheckImageQualityRequest struct {
	ctx context.Context
	ApiService *KYCApiService
	referenceId string
	consent bool
	consentPurpose string
	image **os.File
	qualityParameter *string
	imageUrl *string
}

func (r *KYCApiCheckImageQualityRequest) Image(image *os.File) *KYCApiCheckImageQualityRequest {
	r.image = &image
	return r
}

func (r *KYCApiCheckImageQualityRequest) QualityParameter(qualityParameter string) *KYCApiCheckImageQualityRequest {
	r.qualityParameter = &qualityParameter
	return r
}

func (r *KYCApiCheckImageQualityRequest) ImageUrl(imageUrl string) *KYCApiCheckImageQualityRequest {
	r.imageUrl = &imageUrl
	return r
}

func (r KYCApiCheckImageQualityRequest) Execute() (*CheckImageQualityResponse, *http.Response, error) {
	return r.ApiService.CheckImageQualityExecute(r)
}

/*
CheckImageQuality Image Quality Check

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param referenceId
 @param consent
 @param consentPurpose
 @return KYCApiCheckImageQualityRequest
*/
func (a *KYCApiService) CheckImageQuality(referenceId string, consent bool, consentPurpose string) KYCApiCheckImageQualityRequest {
	return KYCApiCheckImageQualityRequest{
		ApiService: a,
		ctx: a.client.cfg.Context,
		referenceId: referenceId,
		consent: consent,
		consentPurpose: consentPurpose,
	}
}

// Execute executes the request
//  @return CheckImageQualityResponse
func (a *KYCApiService) CheckImageQualityExecute(r KYCApiCheckImageQualityRequest) (*CheckImageQualityResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CheckImageQualityResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "KYCApiService.CheckImageQuality")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2/kyc/forensics/image_quality"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if strlen(r.referenceId) < 1 {
		return localVarReturnValue, nil, reportError("referenceId must have at least 1 elements")
	}
	if strlen(r.consentPurpose) < 1 {
		return localVarReturnValue, nil, reportError("consentPurpose must have at least 1 elements")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"multipart/form-data"}

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
	localVarFormParams.Add("reference_id", parameterToString(r.referenceId, ""))
	localVarFormParams.Add("consent", parameterToString(r.consent, ""))
	localVarFormParams.Add("consent_purpose", parameterToString(r.consentPurpose, ""))
	var imageLocalVarFormFileName string
	var imageLocalVarFileName     string
	var imageLocalVarFileBytes    []byte

	imageLocalVarFormFileName = "image"

	var imageLocalVarFile *os.File
	if r.image != nil {
		imageLocalVarFile = *r.image
	}
	if imageLocalVarFile != nil {
		fbs, _ := ioutil.ReadAll(imageLocalVarFile)
		imageLocalVarFileBytes = fbs
		imageLocalVarFileName = imageLocalVarFile.Name()
		imageLocalVarFile.Close()
	}
	formFiles = append(formFiles, formFile{fileBytes: imageLocalVarFileBytes, fileName: imageLocalVarFileName, formFileName: imageLocalVarFormFileName})
	if r.qualityParameter != nil {
		localVarFormParams.Add("quality_parameter", parameterToString(*r.qualityParameter, ""))
	}
	if r.imageUrl != nil {
		localVarFormParams.Add("image_url", parameterToString(*r.imageUrl, ""))
	}
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
			var v CheckImageQuality400Response
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

type KYCApiCheckPhotocopyRequest struct {
	ctx context.Context
	ApiService *KYCApiService
	referenceId string
	consent bool
	consentPurpose string
	image **os.File
	imageUrl *string
}

func (r *KYCApiCheckPhotocopyRequest) Image(image *os.File) *KYCApiCheckPhotocopyRequest {
	r.image = &image
	return r
}

func (r *KYCApiCheckPhotocopyRequest) ImageUrl(imageUrl string) *KYCApiCheckPhotocopyRequest {
	r.imageUrl = &imageUrl
	return r
}

func (r KYCApiCheckPhotocopyRequest) Execute() (*CheckPhotocopyResponse, *http.Response, error) {
	return r.ApiService.CheckPhotocopyExecute(r)
}

/*
CheckPhotocopy Photocopy Check

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param referenceId 
 @param consent 
 @param consentPurpose 
 @return KYCApiCheckPhotocopyRequest
*/
func (a *KYCApiService) CheckPhotocopy(referenceId string, consent bool, consentPurpose string) KYCApiCheckPhotocopyRequest {
	return KYCApiCheckPhotocopyRequest{
		ApiService: a,
		ctx: a.client.cfg.Context,
		referenceId: referenceId,
		consent: consent,
		consentPurpose: consentPurpose,
	}
}

// Execute executes the request
//  @return CheckPhotocopyResponse
func (a *KYCApiService) CheckPhotocopyExecute(r KYCApiCheckPhotocopyRequest) (*CheckPhotocopyResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CheckPhotocopyResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "KYCApiService.CheckPhotocopy")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2/kyc/forensics/photocopy_check"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if strlen(r.referenceId) < 1 {
		return localVarReturnValue, nil, reportError("referenceId must have at least 1 elements")
	}
	if strlen(r.consentPurpose) < 1 {
		return localVarReturnValue, nil, reportError("consentPurpose must have at least 1 elements")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"multipart/form-data"}

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
	localVarFormParams.Add("reference_id", parameterToString(r.referenceId, ""))
	localVarFormParams.Add("consent", parameterToString(r.consent, ""))
	localVarFormParams.Add("consent_purpose", parameterToString(r.consentPurpose, ""))
	var imageLocalVarFormFileName string
	var imageLocalVarFileName     string
	var imageLocalVarFileBytes    []byte

	imageLocalVarFormFileName = "image"

	var imageLocalVarFile *os.File
	if r.image != nil {
		imageLocalVarFile = *r.image
	}
	if imageLocalVarFile != nil {
		fbs, _ := ioutil.ReadAll(imageLocalVarFile)
		imageLocalVarFileBytes = fbs
		imageLocalVarFileName = imageLocalVarFile.Name()
		imageLocalVarFile.Close()
	}
	formFiles = append(formFiles, formFile{fileBytes: imageLocalVarFileBytes, fileName: imageLocalVarFileName, formFileName: imageLocalVarFormFileName})
	if r.imageUrl != nil {
		localVarFormParams.Add("image_url", parameterToString(*r.imageUrl, ""))
	}
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
			var v CheckPhotocopy400Response
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

type KYCApiCheckVideoLivenessRequest struct {
	ctx context.Context
	ApiService *KYCApiService
	referenceId string
	consent string
	consentPurpose string
	video **os.File
	videoUrl *string
}

func (r *KYCApiCheckVideoLivenessRequest) Video(video *os.File) *KYCApiCheckVideoLivenessRequest {
	r.video = &video
	return r
}

func (r *KYCApiCheckVideoLivenessRequest) VideoUrl(videoUrl string) *KYCApiCheckVideoLivenessRequest {
	r.videoUrl = &videoUrl
	return r
}

func (r KYCApiCheckVideoLivenessRequest) Execute() (*CheckVideoLivenessResponse, *http.Response, error) {
	return r.ApiService.CheckVideoLivenessExecute(r)
}

/*
CheckVideoLiveness Liveness Check

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param referenceId 
 @param consent 
 @param consentPurpose 
 @return KYCApiCheckVideoLivenessRequest
*/
func (a *KYCApiService) CheckVideoLiveness(referenceId string, consent string, consentPurpose string) KYCApiCheckVideoLivenessRequest {
	return KYCApiCheckVideoLivenessRequest{
		ApiService: a,
		ctx: a.client.cfg.Context,
		referenceId: referenceId,
		consent: consent,
		consentPurpose: consentPurpose,
	}
}

// Execute executes the request
//  @return CheckVideoLivenessResponse
func (a *KYCApiService) CheckVideoLivenessExecute(r KYCApiCheckVideoLivenessRequest) (*CheckVideoLivenessResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CheckVideoLivenessResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "KYCApiService.CheckVideoLiveness")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2/kyc/forensics/video_liveness"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if strlen(r.referenceId) < 1 {
		return localVarReturnValue, nil, reportError("referenceId must have at least 1 elements")
	}
	if strlen(r.consent) < 1 {
		return localVarReturnValue, nil, reportError("consent must have at least 1 elements")
	}
	if strlen(r.consentPurpose) < 1 {
		return localVarReturnValue, nil, reportError("consentPurpose must have at least 1 elements")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"multipart/form-data"}

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
	localVarFormParams.Add("reference_id", parameterToString(r.referenceId, ""))
	localVarFormParams.Add("consent", parameterToString(r.consent, ""))
	localVarFormParams.Add("consent_purpose", parameterToString(r.consentPurpose, ""))
	var videoLocalVarFormFileName string
	var videoLocalVarFileName     string
	var videoLocalVarFileBytes    []byte

	videoLocalVarFormFileName = "video"

	var videoLocalVarFile *os.File
	if r.video != nil {
		videoLocalVarFile = *r.video
	}
	if videoLocalVarFile != nil {
		fbs, _ := ioutil.ReadAll(videoLocalVarFile)
		videoLocalVarFileBytes = fbs
		videoLocalVarFileName = videoLocalVarFile.Name()
		videoLocalVarFile.Close()
	}
	formFiles = append(formFiles, formFile{fileBytes: videoLocalVarFileBytes, fileName: videoLocalVarFileName, formFileName: videoLocalVarFormFileName})
	if r.videoUrl != nil {
		localVarFormParams.Add("video_url", parameterToString(*r.videoUrl, ""))
	}
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
			var v CheckVideoLiveness400Response
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

type KYCApiClassifyDocumentRequest struct {
	ctx context.Context
	ApiService *KYCApiService
	referenceId string
	documentType string
	consent bool
	consentPurpose string
	document **os.File
	documentUrl *string
}

func (r *KYCApiClassifyDocumentRequest) Document(document *os.File) *KYCApiClassifyDocumentRequest {
	r.document = &document
	return r
}

func (r *KYCApiClassifyDocumentRequest) DocumentUrl(documentUrl string) *KYCApiClassifyDocumentRequest {
	r.documentUrl = &documentUrl
	return r
}

func (r KYCApiClassifyDocumentRequest) Execute() (*ClassifyDocumentResponse, *http.Response, error) {
	return r.ApiService.ClassifyDocumentExecute(r)
}

/*
ClassifyDocument ID Classification

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param referenceId 
 @param documentType 
 @param consent 
 @param consentPurpose 
 @return KYCApiClassifyDocumentRequest
*/
func (a *KYCApiService) ClassifyDocument(referenceId string, documentType string, consent bool, consentPurpose string) KYCApiClassifyDocumentRequest {
	return KYCApiClassifyDocumentRequest{
		ApiService: a,
		ctx: a.client.cfg.Context,
		referenceId: referenceId,
		documentType: documentType,
		consent: consent,
		consentPurpose: consentPurpose,
	}
}

// Execute executes the request
//  @return ClassifyDocumentResponse
func (a *KYCApiService) ClassifyDocumentExecute(r KYCApiClassifyDocumentRequest) (*ClassifyDocumentResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ClassifyDocumentResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "KYCApiService.ClassifyDocument")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2/kyc/document_classification"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if strlen(r.referenceId) < 1 {
		return localVarReturnValue, nil, reportError("referenceId must have at least 1 elements")
	}
	if strlen(r.documentType) < 1 {
		return localVarReturnValue, nil, reportError("documentType must have at least 1 elements")
	}
	if strlen(r.consentPurpose) < 1 {
		return localVarReturnValue, nil, reportError("consentPurpose must have at least 1 elements")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"multipart/form-data"}

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
	localVarFormParams.Add("reference_id", parameterToString(r.referenceId, ""))
	localVarFormParams.Add("document_type", parameterToString(r.documentType, ""))
	localVarFormParams.Add("consent", parameterToString(r.consent, ""))
	localVarFormParams.Add("consent_purpose", parameterToString(r.consentPurpose, ""))
	var documentLocalVarFormFileName string
	var documentLocalVarFileName     string
	var documentLocalVarFileBytes    []byte

	documentLocalVarFormFileName = "document"

	var documentLocalVarFile *os.File
	if r.document != nil {
		documentLocalVarFile = *r.document
	}
	if documentLocalVarFile != nil {
		fbs, _ := ioutil.ReadAll(documentLocalVarFile)
		documentLocalVarFileBytes = fbs
		documentLocalVarFileName = documentLocalVarFile.Name()
		documentLocalVarFile.Close()
	}
	formFiles = append(formFiles, formFile{fileBytes: documentLocalVarFileBytes, fileName: documentLocalVarFileName, formFileName: documentLocalVarFormFileName})
	if r.documentUrl != nil {
		localVarFormParams.Add("document_url", parameterToString(*r.documentUrl, ""))
	}
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
			var v ClassifyDocument400Response
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

type KYCApiExtractTextRequest struct {
	ctx context.Context
	ApiService *KYCApiService
	referenceId string
	documentType string
	consent string
	consentPurpose string
	kycValidate int32
	document **os.File
	documentUrl *string
	documentBack **os.File
	documentBackUrl *string
}

func (r *KYCApiExtractTextRequest) Document(document *os.File) *KYCApiExtractTextRequest {
	r.document = &document
	return r
}

func (r *KYCApiExtractTextRequest) DocumentUrl(documentUrl string) *KYCApiExtractTextRequest {
	r.documentUrl = &documentUrl
	return r
}

func (r *KYCApiExtractTextRequest) DocumentBack(documentBack *os.File) *KYCApiExtractTextRequest {
	r.documentBack = &documentBack
	return r
}

func (r *KYCApiExtractTextRequest) DocumentBackUrl(documentBackUrl string) *KYCApiExtractTextRequest {
	r.documentBackUrl = &documentBackUrl
	return r
}

func (r KYCApiExtractTextRequest) Execute() (*ExtractTextResponse, *http.Response, error) {
	return r.ApiService.ExtractTextExecute(r)
}

/*
ExtractText Scan & Extract

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param referenceId 
 @param documentType 
 @param consent 
 @param consentPurpose 
 @param kycValidate 
 @return KYCApiExtractTextRequest
*/
func (a *KYCApiService) ExtractText(referenceId string, documentType string, consent string, consentPurpose string, kycValidate int32) KYCApiExtractTextRequest {
	return KYCApiExtractTextRequest{
		ApiService: a,
		ctx: a.client.cfg.Context,
		referenceId: referenceId,
		documentType: documentType,
		consent: consent,
		consentPurpose: consentPurpose,
		kycValidate: kycValidate,
	}
}

// Execute executes the request
//  @return ExtractTextResponse
func (a *KYCApiService) ExtractTextExecute(r KYCApiExtractTextRequest) (*ExtractTextResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ExtractTextResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "KYCApiService.ExtractText")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/kyc/scan_extract/ocr"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if strlen(r.referenceId) < 1 {
		return localVarReturnValue, nil, reportError("referenceId must have at least 1 elements")
	}
	if strlen(r.documentType) < 1 {
		return localVarReturnValue, nil, reportError("documentType must have at least 1 elements")
	}
	if strlen(r.consent) < 1 {
		return localVarReturnValue, nil, reportError("consent must have at least 1 elements")
	}
	if strlen(r.consentPurpose) < 1 {
		return localVarReturnValue, nil, reportError("consentPurpose must have at least 1 elements")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"multipart/form-data"}

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
	localVarFormParams.Add("reference_id", parameterToString(r.referenceId, ""))
	localVarFormParams.Add("document_type", parameterToString(r.documentType, ""))
	localVarFormParams.Add("consent", parameterToString(r.consent, ""))
	localVarFormParams.Add("consent_purpose", parameterToString(r.consentPurpose, ""))
	localVarFormParams.Add("kyc_validate", parameterToString(r.kycValidate, ""))
	var documentLocalVarFormFileName string
	var documentLocalVarFileName     string
	var documentLocalVarFileBytes    []byte

	documentLocalVarFormFileName = "document"

	var documentLocalVarFile *os.File
	if r.document != nil {
		documentLocalVarFile = *r.document
	}
	if documentLocalVarFile != nil {
		fbs, _ := ioutil.ReadAll(documentLocalVarFile)
		documentLocalVarFileBytes = fbs
		documentLocalVarFileName = documentLocalVarFile.Name()
		documentLocalVarFile.Close()
	}
	formFiles = append(formFiles, formFile{fileBytes: documentLocalVarFileBytes, fileName: documentLocalVarFileName, formFileName: documentLocalVarFormFileName})
	if r.documentUrl != nil {
		localVarFormParams.Add("document_url", parameterToString(*r.documentUrl, ""))
	}
	var documentBackLocalVarFormFileName string
	var documentBackLocalVarFileName     string
	var documentBackLocalVarFileBytes    []byte

	documentBackLocalVarFormFileName = "document_back"

	var documentBackLocalVarFile *os.File
	if r.documentBack != nil {
		documentBackLocalVarFile = *r.documentBack
	}
	if documentBackLocalVarFile != nil {
		fbs, _ := ioutil.ReadAll(documentBackLocalVarFile)
		documentBackLocalVarFileBytes = fbs
		documentBackLocalVarFileName = documentBackLocalVarFile.Name()
		documentBackLocalVarFile.Close()
	}
	formFiles = append(formFiles, formFile{fileBytes: documentBackLocalVarFileBytes, fileName: documentBackLocalVarFileName, formFileName: documentBackLocalVarFormFileName})
	if r.documentBackUrl != nil {
		localVarFormParams.Add("document_back_url", parameterToString(*r.documentBackUrl, ""))
	}
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
			var v ExtractText400Response
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

type KYCApiMaskAadhaarUidRequest struct {
	ctx context.Context
	ApiService *KYCApiService
	referenceId string
	consent bool
	consentPurpose string
	image **os.File
	imageUrl *string
}

func (r *KYCApiMaskAadhaarUidRequest) Image(image *os.File) *KYCApiMaskAadhaarUidRequest {
	r.image = &image
	return r
}

func (r *KYCApiMaskAadhaarUidRequest) ImageUrl(imageUrl string) *KYCApiMaskAadhaarUidRequest {
	r.imageUrl = &imageUrl
	return r
}

func (r KYCApiMaskAadhaarUidRequest) Execute() (*MaskAadhaarUidResponse, *http.Response, error) {
	return r.ApiService.MaskAadhaarUidExecute(r)
}

/*
MaskAadhaarUid Aadhaar Masking

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param referenceId
 @param consent
 @param consentPurpose
 @return KYCApiMaskAadhaarUidRequest
*/
func (a *KYCApiService) MaskAadhaarUid(referenceId string, consent bool, consentPurpose string) KYCApiMaskAadhaarUidRequest {
	return KYCApiMaskAadhaarUidRequest{
		ApiService: a,
		ctx: a.client.cfg.Context,
		referenceId: referenceId,
		consent: consent,
		consentPurpose: consentPurpose,
	}
}

// Execute executes the request
//  @return MaskAadhaarUidResponse
func (a *KYCApiService) MaskAadhaarUidExecute(r KYCApiMaskAadhaarUidRequest) (*MaskAadhaarUidResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *MaskAadhaarUidResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "KYCApiService.MaskAadhaarUid")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2/kyc/identities/mask_aadhaar_uid"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if strlen(r.referenceId) < 1 {
		return localVarReturnValue, nil, reportError("referenceId must have at least 1 elements")
	}
	if strlen(r.consentPurpose) < 1 {
		return localVarReturnValue, nil, reportError("consentPurpose must have at least 1 elements")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"multipart/form-data"}

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
	localVarFormParams.Add("reference_id", parameterToString(r.referenceId, ""))
	localVarFormParams.Add("consent", parameterToString(r.consent, ""))
	localVarFormParams.Add("consent_purpose", parameterToString(r.consentPurpose, ""))
	var imageLocalVarFormFileName string
	var imageLocalVarFileName     string
	var imageLocalVarFileBytes    []byte

	imageLocalVarFormFileName = "image"

	var imageLocalVarFile *os.File
	if r.image != nil {
		imageLocalVarFile = *r.image
	}
	if imageLocalVarFile != nil {
		fbs, _ := ioutil.ReadAll(imageLocalVarFile)
		imageLocalVarFileBytes = fbs
		imageLocalVarFileName = imageLocalVarFile.Name()
		imageLocalVarFile.Close()
	}
	formFiles = append(formFiles, formFile{fileBytes: imageLocalVarFileBytes, fileName: imageLocalVarFileName, formFileName: imageLocalVarFormFileName})
	if r.imageUrl != nil {
		localVarFormParams.Add("image_url", parameterToString(*r.imageUrl, ""))
	}
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
			var v MaskAadhaarUid400Response
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

type KYCApiMatchFaceRequest struct {
	ctx context.Context
	ApiService *KYCApiService
	referenceId string
	consent string
	consentPurpose string
	image1 **os.File
	image2 **os.File
	threshold *int32
	image1Url *string
	image2Url *string
}

func (r *KYCApiMatchFaceRequest) Image1(image1 *os.File) *KYCApiMatchFaceRequest {
	r.image1 = &image1
	return r
}

func (r *KYCApiMatchFaceRequest) Image2(image2 *os.File) *KYCApiMatchFaceRequest {
	r.image2 = &image2
	return r
}

func (r *KYCApiMatchFaceRequest) Threshold(threshold int32) *KYCApiMatchFaceRequest {
	r.threshold = &threshold
	return r
}

func (r *KYCApiMatchFaceRequest) Image1Url(image1Url string) *KYCApiMatchFaceRequest {
	r.image1Url = &image1Url
	return r
}

func (r *KYCApiMatchFaceRequest) Image2Url(image2Url string) *KYCApiMatchFaceRequest {
	r.image2Url = &image2Url
	return r
}

func (r KYCApiMatchFaceRequest) Execute() (*MatchFaceResponse, *http.Response, error) {
	return r.ApiService.MatchFaceExecute(r)
}

/*
MatchFace Face Match

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param referenceId 
 @param consent 
 @param consentPurpose 
 @return KYCApiMatchFaceRequest
*/
func (a *KYCApiService) MatchFace(referenceId string, consent string, consentPurpose string) KYCApiMatchFaceRequest {
	return KYCApiMatchFaceRequest{
		ApiService: a,
		ctx: a.client.cfg.Context,
		referenceId: referenceId,
		consent: consent,
		consentPurpose: consentPurpose,
	}
}

// Execute executes the request
//  @return MatchFaceResponse
func (a *KYCApiService) MatchFaceExecute(r KYCApiMatchFaceRequest) (*MatchFaceResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *MatchFaceResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "KYCApiService.MatchFace")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2/kyc/forensics/face_match"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if strlen(r.referenceId) < 1 {
		return localVarReturnValue, nil, reportError("referenceId must have at least 1 elements")
	}
	if strlen(r.consent) < 1 {
		return localVarReturnValue, nil, reportError("consent must have at least 1 elements")
	}
	if strlen(r.consentPurpose) < 1 {
		return localVarReturnValue, nil, reportError("consentPurpose must have at least 1 elements")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"multipart/form-data"}

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
	localVarFormParams.Add("reference_id", parameterToString(r.referenceId, ""))
	localVarFormParams.Add("consent", parameterToString(r.consent, ""))
	localVarFormParams.Add("consent_purpose", parameterToString(r.consentPurpose, ""))
	var image1LocalVarFormFileName string
	var image1LocalVarFileName     string
	var image1LocalVarFileBytes    []byte

	image1LocalVarFormFileName = "image1"

	var image1LocalVarFile *os.File
	if r.image1 != nil {
		image1LocalVarFile = *r.image1
	}
	if image1LocalVarFile != nil {
		fbs, _ := ioutil.ReadAll(image1LocalVarFile)
		image1LocalVarFileBytes = fbs
		image1LocalVarFileName = image1LocalVarFile.Name()
		image1LocalVarFile.Close()
	}
	formFiles = append(formFiles, formFile{fileBytes: image1LocalVarFileBytes, fileName: image1LocalVarFileName, formFileName: image1LocalVarFormFileName})
	var image2LocalVarFormFileName string
	var image2LocalVarFileName     string
	var image2LocalVarFileBytes    []byte

	image2LocalVarFormFileName = "image2"

	var image2LocalVarFile *os.File
	if r.image2 != nil {
		image2LocalVarFile = *r.image2
	}
	if image2LocalVarFile != nil {
		fbs, _ := ioutil.ReadAll(image2LocalVarFile)
		image2LocalVarFileBytes = fbs
		image2LocalVarFileName = image2LocalVarFile.Name()
		image2LocalVarFile.Close()
	}
	formFiles = append(formFiles, formFile{fileBytes: image2LocalVarFileBytes, fileName: image2LocalVarFileName, formFileName: image2LocalVarFormFileName})
	if r.threshold != nil {
		localVarFormParams.Add("threshold", parameterToString(*r.threshold, ""))
	}
	if r.image1Url != nil {
		localVarFormParams.Add("image1_url", parameterToString(*r.image1Url, ""))
	}
	if r.image2Url != nil {
		localVarFormParams.Add("image2_url", parameterToString(*r.image2Url, ""))
	}
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
			var v MatchFace400Response
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

type KYCApiValidateRequest struct {
	ctx context.Context
	ApiService *KYCApiService
	validateRequest ValidateRequest
}

func (r KYCApiValidateRequest) Execute() (*ValidateResponse, *http.Response, error) {
	return r.ApiService.ValidateExecute(r)
}

/*
Validate Validate

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param validateRequest
 @return KYCApiValidateRequest
*/
func (a *KYCApiService) Validate(validateRequest ValidateRequest) KYCApiValidateRequest {
	return KYCApiValidateRequest{
		ApiService: a,
		ctx: a.client.cfg.Context,
		validateRequest: validateRequest,
	}
}

// Execute executes the request
//  @return ValidateResponse
func (a *KYCApiService) ValidateExecute(r KYCApiValidateRequest) (*ValidateResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ValidateResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "KYCApiService.Validate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/kyc/public_registry/validate"

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
	localVarPostBody = r.validateRequest
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
			var v Validate400Response
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
