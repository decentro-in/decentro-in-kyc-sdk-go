# decentroinkyc - Decentro's Go SDK

KYC & Onboarding

For more information, please visit [https://decentro.tech](https://decentro.tech)

## Installation

Add to your project:

```shell
go get github.com/decentro-in/decentro-in-kyc-sdk-go
```

## Getting Started

```golang
package main

import (
    "fmt"
    "os"
    decentroinkyc "github.com/decentro-in/decentro-in-kyc-sdk-go"
)

func main() {
    configuration := decentroinkyc.NewConfiguration()
    configuration.SetClientId("CLIENT_ID")
    configuration.SetClientSecret("CLIENT_SECRET")
    configuration.SetModuleSecret("MODULE_SECRET")
    client := decentroinkyc.NewAPIClient(configuration)

    request := client.KYCApi.CheckImageQuality(
        "referenceId_example",
        true,
        "consentPurpose_example",
    )
    request.Image(os.NewFile(1234, "some_file"))
    request.QualityParameter("qualityParameter_example")
    request.ImageUrl("imageUrl_example")
    
    resp, httpRes, err := request.Execute()

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KYCApi.CheckImageQuality``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", httpRes)
    }
    // response from `CheckImageQuality`: CheckImageQualityResponse
    fmt.Fprintf(os.Stdout, "Response from `KYCApi.CheckImageQuality`: %v\n", resp)
    fmt.Fprintf(os.Stdout, "Response from `CheckImageQualityResponse.CheckImageQuality.DecentroTxnId`: %v\n", *resp.DecentroTxnId)
    fmt.Fprintf(os.Stdout, "Response from `CheckImageQualityResponse.CheckImageQuality.Status`: %v\n", *resp.Status)
    fmt.Fprintf(os.Stdout, "Response from `CheckImageQualityResponse.CheckImageQuality.ResponseCode`: %v\n", *resp.ResponseCode)
    fmt.Fprintf(os.Stdout, "Response from `CheckImageQualityResponse.CheckImageQuality.Message`: %v\n", *resp.Message)
    fmt.Fprintf(os.Stdout, "Response from `CheckImageQualityResponse.CheckImageQuality.Data`: %v\n", *resp.Data)
}

```

## Documentation for API Endpoints

All URIs are relative to *https://in.staging.decentro.tech*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*KYCApi* | [**CheckImageQuality**](docs/KYCApi.md#checkimagequality) | **Post** /v2/kyc/forensics/image_quality | Image Quality Check
*KYCApi* | [**CheckPhotocopy**](docs/KYCApi.md#checkphotocopy) | **Post** /v2/kyc/forensics/photocopy_check | Photocopy Check
*KYCApi* | [**CheckVideoLiveness**](docs/KYCApi.md#checkvideoliveness) | **Post** /v2/kyc/forensics/video_liveness | Liveness Check
*KYCApi* | [**ClassifyDocument**](docs/KYCApi.md#classifydocument) | **Post** /v2/kyc/document_classification | ID Classification
*KYCApi* | [**ExtractText**](docs/KYCApi.md#extracttext) | **Post** /kyc/scan_extract/ocr | Scan &amp; Extract
*KYCApi* | [**MaskAadhaarUid**](docs/KYCApi.md#maskaadhaaruid) | **Post** /v2/kyc/identities/mask_aadhaar_uid | Aadhaar Masking
*KYCApi* | [**MatchFace**](docs/KYCApi.md#matchface) | **Post** /v2/kyc/forensics/face_match | Face Match
*KYCApi* | [**Validate**](docs/KYCApi.md#validate) | **Post** /kyc/public_registry/validate | Validate
*AadhaarXmlApi* | [**GenerateOtp**](docs/AadhaarXmlApi.md#generateotp) | **Post** /v2/kyc/aadhaar_connect/otp | Generate OTP
*AadhaarXmlApi* | [**InitiateSession**](docs/AadhaarXmlApi.md#initiatesession) | **Post** /v2/kyc/aadhaar_connect | Initiate Session
*AadhaarXmlApi* | [**ReloadCaptcha**](docs/AadhaarXmlApi.md#reloadcaptcha) | **Post** /v2/kyc/aadhaar_connect/captcha/reload | Reload Captcha
*AadhaarXmlApi* | [**ValidateOtp**](docs/AadhaarXmlApi.md#validateotp) | **Post** /v2/kyc/aadhaar_connect/otp/validate | Validate OTP
*DigilockerApi* | [**DownloadFile**](docs/DigilockerApi.md#downloadfile) | **Post** /v2/kyc/sso/digilocker/{initial_decentro_txn_id}/file/download | Download File
*DigilockerApi* | [**EAadhaar**](docs/DigilockerApi.md#eaadhaar) | **Post** /v2/kyc/sso/digilocker/{initial_decentro_txn_id}/eaadhaar | eAadhaar
*DigilockerApi* | [**FileData**](docs/DigilockerApi.md#filedata) | **Post** /v2/kyc/sso/digilocker/{initial_decentro_txn_id}/file | File Data
*DigilockerApi* | [**GenerateAccessToken**](docs/DigilockerApi.md#generateaccesstoken) | **Post** /v2/kyc/sso/digilocker/{initial_decentro_txn_id}/token | Access Token
*DigilockerApi* | [**InitiateSession**](docs/DigilockerApi.md#initiatesession) | **Post** /v2/kyc/sso/digilocker/session | Initiate Session
*DigilockerApi* | [**IssuedFiles**](docs/DigilockerApi.md#issuedfiles) | **Post** /v2/kyc/sso/digilocker/{initial_decentro_txn_id}/files/issued | Issued Files
*DigilockerApi* | [**PullFile**](docs/DigilockerApi.md#pullfile) | **Post** /v2/kyc/sso/digilocker/{initial_decentro_txn_id}/file/pull | Pull File
*DigilockerApi* | [**RevokeToken**](docs/DigilockerApi.md#revoketoken) | **Delete** /v2/kyc/sso/digilocker/session/{initial_decentro_txn_id} | Revoke Token
*UpgradedAadhaarXmlApi* | [**GenerateOtp**](docs/UpgradedAadhaarXmlApi.md#generateotp) | **Post** /v2/kyc/aadhaar/otp | Generate OTP
*UpgradedAadhaarXmlApi* | [**ValidateOtp**](docs/UpgradedAadhaarXmlApi.md#validateotp) | **Post** /v2/kyc/aadhaar/otp/validate | Validate OTP


## Documentation For Models

 - [AadhaarxmlGenerateOtp400Response](docs/AadhaarxmlGenerateOtp400Response.md)
 - [AadhaarxmlGenerateOtpRequest](docs/AadhaarxmlGenerateOtpRequest.md)
 - [AadhaarxmlGenerateOtpResponse](docs/AadhaarxmlGenerateOtpResponse.md)
 - [AadhaarxmlInitiateSession400Response](docs/AadhaarxmlInitiateSession400Response.md)
 - [AadhaarxmlInitiateSession400ResponseData](docs/AadhaarxmlInitiateSession400ResponseData.md)
 - [AadhaarxmlInitiateSessionRequest](docs/AadhaarxmlInitiateSessionRequest.md)
 - [AadhaarxmlInitiateSessionResponse](docs/AadhaarxmlInitiateSessionResponse.md)
 - [AadhaarxmlInitiateSessionResponseData](docs/AadhaarxmlInitiateSessionResponseData.md)
 - [AadhaarxmlReloadCaptcha400Response](docs/AadhaarxmlReloadCaptcha400Response.md)
 - [AadhaarxmlReloadCaptcha400ResponseData](docs/AadhaarxmlReloadCaptcha400ResponseData.md)
 - [AadhaarxmlReloadCaptchaRequest](docs/AadhaarxmlReloadCaptchaRequest.md)
 - [AadhaarxmlReloadCaptchaResponse](docs/AadhaarxmlReloadCaptchaResponse.md)
 - [AadhaarxmlReloadCaptchaResponseData](docs/AadhaarxmlReloadCaptchaResponseData.md)
 - [AadhaarxmlValidateOtp400Response](docs/AadhaarxmlValidateOtp400Response.md)
 - [AadhaarxmlValidateOtp400ResponseData](docs/AadhaarxmlValidateOtp400ResponseData.md)
 - [AadhaarxmlValidateOtp400ResponseDataAadhaarFile](docs/AadhaarxmlValidateOtp400ResponseDataAadhaarFile.md)
 - [AadhaarxmlValidateOtp400ResponseDataProofOfAddress](docs/AadhaarxmlValidateOtp400ResponseDataProofOfAddress.md)
 - [AadhaarxmlValidateOtp400ResponseDataProofOfIdentity](docs/AadhaarxmlValidateOtp400ResponseDataProofOfIdentity.md)
 - [AadhaarxmlValidateOtpRequest](docs/AadhaarxmlValidateOtpRequest.md)
 - [AadhaarxmlValidateOtpResponse](docs/AadhaarxmlValidateOtpResponse.md)
 - [AadhaarxmlValidateOtpResponseData](docs/AadhaarxmlValidateOtpResponseData.md)
 - [AadhaarxmlValidateOtpResponseDataAadhaarFile](docs/AadhaarxmlValidateOtpResponseDataAadhaarFile.md)
 - [AadhaarxmlValidateOtpResponseDataProofOfAddress](docs/AadhaarxmlValidateOtpResponseDataProofOfAddress.md)
 - [AadhaarxmlValidateOtpResponseDataProofOfIdentity](docs/AadhaarxmlValidateOtpResponseDataProofOfIdentity.md)
 - [CheckImageQuality400Response](docs/CheckImageQuality400Response.md)
 - [CheckImageQualityRequest](docs/CheckImageQualityRequest.md)
 - [CheckImageQualityResponse](docs/CheckImageQualityResponse.md)
 - [CheckImageQualityResponseData](docs/CheckImageQualityResponseData.md)
 - [CheckImageQualityResponseDataImageQuality](docs/CheckImageQualityResponseDataImageQuality.md)
 - [CheckImageQualityResponseDataImageQualityQualityScores](docs/CheckImageQualityResponseDataImageQualityQualityScores.md)
 - [CheckImageQualityResponseDataImageQualityQualityScoresBrightness](docs/CheckImageQualityResponseDataImageQualityQualityScoresBrightness.md)
 - [CheckImageQualityResponseDataImageQualityQualityScoresCompressionQuality](docs/CheckImageQualityResponseDataImageQualityQualityScoresCompressionQuality.md)
 - [CheckImageQualityResponseDataImageQualityQualityScoresSharpness](docs/CheckImageQualityResponseDataImageQualityQualityScoresSharpness.md)
 - [CheckImageQualityResponseDataImageQualityQualityScoresTextQuality](docs/CheckImageQualityResponseDataImageQualityQualityScoresTextQuality.md)
 - [CheckPhotocopy400Response](docs/CheckPhotocopy400Response.md)
 - [CheckPhotocopyRequest](docs/CheckPhotocopyRequest.md)
 - [CheckPhotocopyResponse](docs/CheckPhotocopyResponse.md)
 - [CheckPhotocopyResponseData](docs/CheckPhotocopyResponseData.md)
 - [CheckVideoLiveness400Response](docs/CheckVideoLiveness400Response.md)
 - [CheckVideoLivenessRequest](docs/CheckVideoLivenessRequest.md)
 - [CheckVideoLivenessResponse](docs/CheckVideoLivenessResponse.md)
 - [CheckVideoLivenessResponseData](docs/CheckVideoLivenessResponseData.md)
 - [ClassifyDocument400Response](docs/ClassifyDocument400Response.md)
 - [ClassifyDocumentRequest](docs/ClassifyDocumentRequest.md)
 - [ClassifyDocumentResponse](docs/ClassifyDocumentResponse.md)
 - [ClassifyDocumentResponseData](docs/ClassifyDocumentResponseData.md)
 - [DigilockerDownloadFile400Response](docs/DigilockerDownloadFile400Response.md)
 - [DigilockerDownloadFile400ResponseData](docs/DigilockerDownloadFile400ResponseData.md)
 - [DigilockerDownloadFileRequest](docs/DigilockerDownloadFileRequest.md)
 - [DigilockerDownloadFileResponse](docs/DigilockerDownloadFileResponse.md)
 - [DigilockerDownloadFileResponseData](docs/DigilockerDownloadFileResponseData.md)
 - [DigilockerEAadhaar400Response](docs/DigilockerEAadhaar400Response.md)
 - [DigilockerEAadhaarRequest](docs/DigilockerEAadhaarRequest.md)
 - [DigilockerEAadhaarResponse](docs/DigilockerEAadhaarResponse.md)
 - [DigilockerEAadhaarResponseData](docs/DigilockerEAadhaarResponseData.md)
 - [DigilockerEAadhaarResponseDataProofOfAddress](docs/DigilockerEAadhaarResponseDataProofOfAddress.md)
 - [DigilockerEAadhaarResponseDataProofOfIdentity](docs/DigilockerEAadhaarResponseDataProofOfIdentity.md)
 - [DigilockerFileData400Response](docs/DigilockerFileData400Response.md)
 - [DigilockerFileData400ResponseData](docs/DigilockerFileData400ResponseData.md)
 - [DigilockerFileDataRequest](docs/DigilockerFileDataRequest.md)
 - [DigilockerFileDataResponse](docs/DigilockerFileDataResponse.md)
 - [DigilockerFileDataResponseData](docs/DigilockerFileDataResponseData.md)
 - [DigilockerGenerateAccessToken400Response](docs/DigilockerGenerateAccessToken400Response.md)
 - [DigilockerGenerateAccessTokenRequest](docs/DigilockerGenerateAccessTokenRequest.md)
 - [DigilockerGenerateAccessTokenResponse](docs/DigilockerGenerateAccessTokenResponse.md)
 - [DigilockerGenerateAccessTokenResponseData](docs/DigilockerGenerateAccessTokenResponseData.md)
 - [DigilockerInitiateSession400Response](docs/DigilockerInitiateSession400Response.md)
 - [DigilockerInitiateSessionRequest](docs/DigilockerInitiateSessionRequest.md)
 - [DigilockerInitiateSessionResponse](docs/DigilockerInitiateSessionResponse.md)
 - [DigilockerInitiateSessionResponseData](docs/DigilockerInitiateSessionResponseData.md)
 - [DigilockerIssuedFiles400Response](docs/DigilockerIssuedFiles400Response.md)
 - [DigilockerIssuedFiles400ResponseDataInner](docs/DigilockerIssuedFiles400ResponseDataInner.md)
 - [DigilockerIssuedFilesRequest](docs/DigilockerIssuedFilesRequest.md)
 - [DigilockerIssuedFilesResponse](docs/DigilockerIssuedFilesResponse.md)
 - [DigilockerIssuedFilesResponseDataInner](docs/DigilockerIssuedFilesResponseDataInner.md)
 - [DigilockerPullFile400Response](docs/DigilockerPullFile400Response.md)
 - [DigilockerPullFileRequest](docs/DigilockerPullFileRequest.md)
 - [DigilockerPullFileRequestDocumentTypeRelatedParameters](docs/DigilockerPullFileRequestDocumentTypeRelatedParameters.md)
 - [DigilockerPullFileResponse](docs/DigilockerPullFileResponse.md)
 - [DigilockerRevokeToken400Response](docs/DigilockerRevokeToken400Response.md)
 - [DigilockerRevokeToken400ResponseData](docs/DigilockerRevokeToken400ResponseData.md)
 - [DigilockerRevokeTokenResponse](docs/DigilockerRevokeTokenResponse.md)
 - [DigilockerRevokeTokenResponseData](docs/DigilockerRevokeTokenResponseData.md)
 - [ExtractText400Response](docs/ExtractText400Response.md)
 - [ExtractText400ResponseError](docs/ExtractText400ResponseError.md)
 - [ExtractTextRequest](docs/ExtractTextRequest.md)
 - [ExtractTextResponse](docs/ExtractTextResponse.md)
 - [ExtractTextResponseOcrResult](docs/ExtractTextResponseOcrResult.md)
 - [MaskAadhaarRequest](docs/MaskAadhaarRequest.md)
 - [MaskAadhaarUid400Response](docs/MaskAadhaarUid400Response.md)
 - [MaskAadhaarUidResponse](docs/MaskAadhaarUidResponse.md)
 - [MaskAadhaarUidResponseData](docs/MaskAadhaarUidResponseData.md)
 - [MatchFace400Response](docs/MatchFace400Response.md)
 - [MatchFaceRequest](docs/MatchFaceRequest.md)
 - [MatchFaceResponse](docs/MatchFaceResponse.md)
 - [MatchFaceResponseData](docs/MatchFaceResponseData.md)
 - [UpgradedaadhaarxmlGenerateOtp400Response](docs/UpgradedaadhaarxmlGenerateOtp400Response.md)
 - [UpgradedaadhaarxmlGenerateOtpRequest](docs/UpgradedaadhaarxmlGenerateOtpRequest.md)
 - [UpgradedaadhaarxmlGenerateOtpResponse](docs/UpgradedaadhaarxmlGenerateOtpResponse.md)
 - [UpgradedaadhaarxmlValidateOtp400Response](docs/UpgradedaadhaarxmlValidateOtp400Response.md)
 - [UpgradedaadhaarxmlValidateOtp400ResponseData](docs/UpgradedaadhaarxmlValidateOtp400ResponseData.md)
 - [UpgradedaadhaarxmlValidateOtp400ResponseDataProofOfAddress](docs/UpgradedaadhaarxmlValidateOtp400ResponseDataProofOfAddress.md)
 - [UpgradedaadhaarxmlValidateOtp400ResponseDataProofOfIdentity](docs/UpgradedaadhaarxmlValidateOtp400ResponseDataProofOfIdentity.md)
 - [UpgradedaadhaarxmlValidateOtpRequest](docs/UpgradedaadhaarxmlValidateOtpRequest.md)
 - [UpgradedaadhaarxmlValidateOtpResponse](docs/UpgradedaadhaarxmlValidateOtpResponse.md)
 - [UpgradedaadhaarxmlValidateOtpResponseData](docs/UpgradedaadhaarxmlValidateOtpResponseData.md)
 - [UpgradedaadhaarxmlValidateOtpResponseDataProofOfAddress](docs/UpgradedaadhaarxmlValidateOtpResponseDataProofOfAddress.md)
 - [UpgradedaadhaarxmlValidateOtpResponseDataProofOfIdentity](docs/UpgradedaadhaarxmlValidateOtpResponseDataProofOfIdentity.md)
 - [Validate400Response](docs/Validate400Response.md)
 - [Validate400ResponseError](docs/Validate400ResponseError.md)
 - [ValidateRequest](docs/ValidateRequest.md)
 - [ValidateResponse](docs/ValidateResponse.md)
 - [ValidateResponseError](docs/ValidateResponseError.md)
 - [ValidateResponseKycResult](docs/ValidateResponseKycResult.md)
 - [ValidateResponseKycResultAddress](docs/ValidateResponseKycResultAddress.md)
 - [ValidateResponseKycResultAddressState](docs/ValidateResponseKycResultAddressState.md)
 - [ValidateResponseKycResultAddressesInner](docs/ValidateResponseKycResultAddressesInner.md)
 - [ValidateResponseKycResultAllClassOfVehicleInner](docs/ValidateResponseKycResultAllClassOfVehicleInner.md)
 - [ValidateResponseKycResultBusinessDetailsInner](docs/ValidateResponseKycResultBusinessDetailsInner.md)
 - [ValidateResponseKycResultCompanyMasterData](docs/ValidateResponseKycResultCompanyMasterData.md)
 - [ValidateResponseKycResultDirectorData](docs/ValidateResponseKycResultDirectorData.md)
 - [ValidateResponseKycResultDirectors](docs/ValidateResponseKycResultDirectors.md)
 - [ValidateResponseKycResultDirectorsOneOfInner](docs/ValidateResponseKycResultDirectorsOneOfInner.md)
 - [ValidateResponseKycResultFilingStatusInner](docs/ValidateResponseKycResultFilingStatusInner.md)
 - [ValidateResponseKycResultForeignCompanyMasterData](docs/ValidateResponseKycResultForeignCompanyMasterData.md)
 - [ValidateResponseKycResultLlpMasterData](docs/ValidateResponseKycResultLlpMasterData.md)
 - [ValidateResponseKycResultPollingBoothDetails](docs/ValidateResponseKycResultPollingBoothDetails.md)
 - [ValidateResponseKycResultPremissesAddress](docs/ValidateResponseKycResultPremissesAddress.md)
 - [ValidateResponseKycResultPrimaryBusinessContact](docs/ValidateResponseKycResultPrimaryBusinessContact.md)
 - [ValidateResponseKycResultPrimaryBusinessContactMobileNumber](docs/ValidateResponseKycResultPrimaryBusinessContactMobileNumber.md)
 - [ValidateResponseKycResultProductsInner](docs/ValidateResponseKycResultProductsInner.md)
