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
    "context"
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
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CheckImageQuality`: CheckImageQualityResponse
    fmt.Fprintf(os.Stdout, "Response from `KYCApi.CheckImageQuality`: %v\n", resp)
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

 - [AadhaarxmlGenerateOtpRequest](docs/AadhaarxmlGenerateOtpRequest.md)
 - [AadhaarxmlInitiateSessionRequest](docs/AadhaarxmlInitiateSessionRequest.md)
 - [AadhaarxmlReloadCaptchaRequest](docs/AadhaarxmlReloadCaptchaRequest.md)
 - [AadhaarxmlValidateOtpRequest](docs/AadhaarxmlValidateOtpRequest.md)
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
 - [DigilockerDownloadFileRequest](docs/DigilockerDownloadFileRequest.md)
 - [DigilockerEAadhaarRequest](docs/DigilockerEAadhaarRequest.md)
 - [DigilockerFileDataRequest](docs/DigilockerFileDataRequest.md)
 - [DigilockerGenerateAccessTokenRequest](docs/DigilockerGenerateAccessTokenRequest.md)
 - [DigilockerInitiateSessionRequest](docs/DigilockerInitiateSessionRequest.md)
 - [DigilockerIssuedFilesRequest](docs/DigilockerIssuedFilesRequest.md)
 - [DigilockerPullFileRequest](docs/DigilockerPullFileRequest.md)
 - [DigilockerPullFileRequestDocumentTypeRelatedParameters](docs/DigilockerPullFileRequestDocumentTypeRelatedParameters.md)
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
 - [UpgradedaadhaarxmlGenerateOtpRequest](docs/UpgradedaadhaarxmlGenerateOtpRequest.md)
 - [UpgradedaadhaarxmlValidateOtpRequest](docs/UpgradedaadhaarxmlValidateOtpRequest.md)
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
