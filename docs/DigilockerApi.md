# APIClient.DigilockerApi

All URIs are relative to *https://in.staging.decentro.tech*

Method | Path | Description
------------- | ------------- | -------------
[**DownloadFile**](DigilockerApi.md#DownloadFile) | **Post** /v2/kyc/sso/digilocker/{initial_decentro_txn_id}/file/download | Download File
[**EAadhaar**](DigilockerApi.md#EAadhaar) | **Post** /v2/kyc/sso/digilocker/{initial_decentro_txn_id}/eaadhaar | eAadhaar
[**FileData**](DigilockerApi.md#FileData) | **Post** /v2/kyc/sso/digilocker/{initial_decentro_txn_id}/file | File Data
[**GenerateAccessToken**](DigilockerApi.md#GenerateAccessToken) | **Post** /v2/kyc/sso/digilocker/{initial_decentro_txn_id}/token | Access Token
[**InitiateSession**](DigilockerApi.md#InitiateSession) | **Post** /v2/kyc/sso/digilocker/session | Initiate Session
[**IssuedFiles**](DigilockerApi.md#IssuedFiles) | **Post** /v2/kyc/sso/digilocker/{initial_decentro_txn_id}/files/issued | Issued Files
[**PullFile**](DigilockerApi.md#PullFile) | **Post** /v2/kyc/sso/digilocker/{initial_decentro_txn_id}/file/pull | Pull File
[**RevokeToken**](DigilockerApi.md#RevokeToken) | **Delete** /v2/kyc/sso/digilocker/session/{initial_decentro_txn_id} | Revoke Token



## DownloadFile

Download File

### Example

```go
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

    
    digilockerDownloadFileRequest := *decentroinkyc.NewDigilockerDownloadFileRequest()
    digilockerDownloadFileRequest.SetFileUrn("in.gov.transport-RVCER-KA05JS9471")
    digilockerDownloadFileRequest.SetConsent(true)
    digilockerDownloadFileRequest.SetPurpose("for bank account verification")
    digilockerDownloadFileRequest.SetReferenceId("ABCDEF123456")
    
    request := client.DigilockerApi.DownloadFile(
        "initialDecentroTxnId_example",
        digilockerDownloadFileRequest,
    )
    
    resp, httpRes, err := request.Execute()

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DigilockerApi.DownloadFile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", httpRes)
    }
    // response from `DownloadFile`: DigilockerDownloadFileResponse
    fmt.Fprintf(os.Stdout, "Response from `DigilockerApi.DownloadFile`: %v\n", resp)
    fmt.Fprintf(os.Stdout, "Response from `DigilockerDownloadFileResponse.DownloadFile.DecentroTxnId`: %v\n", *resp.DecentroTxnId)
    fmt.Fprintf(os.Stdout, "Response from `DigilockerDownloadFileResponse.DownloadFile.Status`: %v\n", *resp.Status)
    fmt.Fprintf(os.Stdout, "Response from `DigilockerDownloadFileResponse.DownloadFile.ResponseCode`: %v\n", *resp.ResponseCode)
    fmt.Fprintf(os.Stdout, "Response from `DigilockerDownloadFileResponse.DownloadFile.Message`: %v\n", *resp.Message)
    fmt.Fprintf(os.Stdout, "Response from `DigilockerDownloadFileResponse.DownloadFile.Data`: %v\n", *resp.Data)
    fmt.Fprintf(os.Stdout, "Response from `DigilockerDownloadFileResponse.DownloadFile.ResponseKey`: %v\n", *resp.ResponseKey)
}
```

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EAadhaar

eAadhaar

### Example

```go
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

    
    digilockerEAadhaarRequest := *decentroinkyc.NewDigilockerEAadhaarRequest()
    digilockerEAadhaarRequest.SetConsent(true)
    digilockerEAadhaarRequest.SetPurpose("for bank account verification")
    digilockerEAadhaarRequest.SetReferenceId("ABCDEF123456")
    digilockerEAadhaarRequest.SetGenerateXml(true)
    digilockerEAadhaarRequest.SetGeneratePdf(true)
    
    request := client.DigilockerApi.EAadhaar(
        "initialDecentroTxnId_example",
        digilockerEAadhaarRequest,
    )
    
    resp, httpRes, err := request.Execute()

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DigilockerApi.EAadhaar``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", httpRes)
    }
    // response from `EAadhaar`: DigilockerEAadhaarResponse
    fmt.Fprintf(os.Stdout, "Response from `DigilockerApi.EAadhaar`: %v\n", resp)
    fmt.Fprintf(os.Stdout, "Response from `DigilockerEAadhaarResponse.EAadhaar.DecentroTxnId`: %v\n", *resp.DecentroTxnId)
    fmt.Fprintf(os.Stdout, "Response from `DigilockerEAadhaarResponse.EAadhaar.Status`: %v\n", *resp.Status)
    fmt.Fprintf(os.Stdout, "Response from `DigilockerEAadhaarResponse.EAadhaar.ResponseCode`: %v\n", *resp.ResponseCode)
    fmt.Fprintf(os.Stdout, "Response from `DigilockerEAadhaarResponse.EAadhaar.Message`: %v\n", *resp.Message)
    fmt.Fprintf(os.Stdout, "Response from `DigilockerEAadhaarResponse.EAadhaar.Data`: %v\n", *resp.Data)
    fmt.Fprintf(os.Stdout, "Response from `DigilockerEAadhaarResponse.EAadhaar.ResponseKey`: %v\n", *resp.ResponseKey)
}
```

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FileData

File Data

### Example

```go
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

    
    digilockerFileDataRequest := *decentroinkyc.NewDigilockerFileDataRequest()
    digilockerFileDataRequest.SetFileUrn("in.gov.pan-PANCR-XXXXXXXXXX")
    digilockerFileDataRequest.SetConsent(true)
    digilockerFileDataRequest.SetPurpose("for bank account verification")
    digilockerFileDataRequest.SetReferenceId("ABCDEF123456")
    
    request := client.DigilockerApi.FileData(
        "initialDecentroTxnId_example",
        digilockerFileDataRequest,
    )
    
    resp, httpRes, err := request.Execute()

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DigilockerApi.FileData``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", httpRes)
    }
    // response from `FileData`: DigilockerFileDataResponse
    fmt.Fprintf(os.Stdout, "Response from `DigilockerApi.FileData`: %v\n", resp)
    fmt.Fprintf(os.Stdout, "Response from `DigilockerFileDataResponse.FileData.DecentroTxnId`: %v\n", *resp.DecentroTxnId)
    fmt.Fprintf(os.Stdout, "Response from `DigilockerFileDataResponse.FileData.Status`: %v\n", *resp.Status)
    fmt.Fprintf(os.Stdout, "Response from `DigilockerFileDataResponse.FileData.ResponseCode`: %v\n", *resp.ResponseCode)
    fmt.Fprintf(os.Stdout, "Response from `DigilockerFileDataResponse.FileData.Message`: %v\n", *resp.Message)
    fmt.Fprintf(os.Stdout, "Response from `DigilockerFileDataResponse.FileData.Data`: %v\n", *resp.Data)
    fmt.Fprintf(os.Stdout, "Response from `DigilockerFileDataResponse.FileData.ResponseKey`: %v\n", *resp.ResponseKey)
}
```

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GenerateAccessToken

Access Token

### Example

```go
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

    
    digilockerGenerateAccessTokenRequest := *decentroinkyc.NewDigilockerGenerateAccessTokenRequest()
    digilockerGenerateAccessTokenRequest.SetCode("XXXXXXXXXXXXXXXXXXXXXXXXXX")
    digilockerGenerateAccessTokenRequest.SetConsent(true)
    digilockerGenerateAccessTokenRequest.SetPurpose("for bank account verification")
    digilockerGenerateAccessTokenRequest.SetReferenceId("ABCDEF123456")
    
    request := client.DigilockerApi.GenerateAccessToken(
        "initialDecentroTxnId_example",
        digilockerGenerateAccessTokenRequest,
    )
    
    resp, httpRes, err := request.Execute()

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DigilockerApi.GenerateAccessToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", httpRes)
    }
    // response from `GenerateAccessToken`: DigilockerGenerateAccessTokenResponse
    fmt.Fprintf(os.Stdout, "Response from `DigilockerApi.GenerateAccessToken`: %v\n", resp)
    fmt.Fprintf(os.Stdout, "Response from `DigilockerGenerateAccessTokenResponse.GenerateAccessToken.DecentroTxnId`: %v\n", *resp.DecentroTxnId)
    fmt.Fprintf(os.Stdout, "Response from `DigilockerGenerateAccessTokenResponse.GenerateAccessToken.Status`: %v\n", *resp.Status)
    fmt.Fprintf(os.Stdout, "Response from `DigilockerGenerateAccessTokenResponse.GenerateAccessToken.ResponseCode`: %v\n", *resp.ResponseCode)
    fmt.Fprintf(os.Stdout, "Response from `DigilockerGenerateAccessTokenResponse.GenerateAccessToken.Message`: %v\n", *resp.Message)
    fmt.Fprintf(os.Stdout, "Response from `DigilockerGenerateAccessTokenResponse.GenerateAccessToken.Data`: %v\n", *resp.Data)
    fmt.Fprintf(os.Stdout, "Response from `DigilockerGenerateAccessTokenResponse.GenerateAccessToken.ResponseKey`: %v\n", *resp.ResponseKey)
}
```

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InitiateSession

Initiate Session

### Example

```go
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

    
    digilockerInitiateSessionRequest := *decentroinkyc.NewDigilockerInitiateSessionRequest()
    digilockerInitiateSessionRequest.SetConsent(true)
    digilockerInitiateSessionRequest.SetPurpose("for bank account verification")
    digilockerInitiateSessionRequest.SetReferenceId("ABCDEF123456")
    digilockerInitiateSessionRequest.SetRedirectUrl("<Redirect URL for DigiLocker>")
    
    request := client.DigilockerApi.InitiateSession(
        digilockerInitiateSessionRequest,
    )
    
    resp, httpRes, err := request.Execute()

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DigilockerApi.InitiateSession``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", httpRes)
    }
    // response from `InitiateSession`: DigilockerInitiateSessionResponse
    fmt.Fprintf(os.Stdout, "Response from `DigilockerApi.InitiateSession`: %v\n", resp)
    fmt.Fprintf(os.Stdout, "Response from `DigilockerInitiateSessionResponse.InitiateSession.DecentroTxnId`: %v\n", *resp.DecentroTxnId)
    fmt.Fprintf(os.Stdout, "Response from `DigilockerInitiateSessionResponse.InitiateSession.Status`: %v\n", *resp.Status)
    fmt.Fprintf(os.Stdout, "Response from `DigilockerInitiateSessionResponse.InitiateSession.ResponseCode`: %v\n", *resp.ResponseCode)
    fmt.Fprintf(os.Stdout, "Response from `DigilockerInitiateSessionResponse.InitiateSession.Message`: %v\n", *resp.Message)
    fmt.Fprintf(os.Stdout, "Response from `DigilockerInitiateSessionResponse.InitiateSession.Data`: %v\n", *resp.Data)
    fmt.Fprintf(os.Stdout, "Response from `DigilockerInitiateSessionResponse.InitiateSession.ResponseKey`: %v\n", *resp.ResponseKey)
}
```

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IssuedFiles

Issued Files

### Example

```go
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

    
    digilockerIssuedFilesRequest := *decentroinkyc.NewDigilockerIssuedFilesRequest()
    digilockerIssuedFilesRequest.SetConsent(true)
    digilockerIssuedFilesRequest.SetPurpose("for bank account verification")
    digilockerIssuedFilesRequest.SetReferenceId("ABCDEF123456")
    
    request := client.DigilockerApi.IssuedFiles(
        "initialDecentroTxnId_example",
        digilockerIssuedFilesRequest,
    )
    
    resp, httpRes, err := request.Execute()

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DigilockerApi.IssuedFiles``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", httpRes)
    }
    // response from `IssuedFiles`: DigilockerIssuedFilesResponse
    fmt.Fprintf(os.Stdout, "Response from `DigilockerApi.IssuedFiles`: %v\n", resp)
    fmt.Fprintf(os.Stdout, "Response from `DigilockerIssuedFilesResponse.IssuedFiles.DecentroTxnId`: %v\n", *resp.DecentroTxnId)
    fmt.Fprintf(os.Stdout, "Response from `DigilockerIssuedFilesResponse.IssuedFiles.Status`: %v\n", *resp.Status)
    fmt.Fprintf(os.Stdout, "Response from `DigilockerIssuedFilesResponse.IssuedFiles.ResponseCode`: %v\n", *resp.ResponseCode)
    fmt.Fprintf(os.Stdout, "Response from `DigilockerIssuedFilesResponse.IssuedFiles.Message`: %v\n", *resp.Message)
    fmt.Fprintf(os.Stdout, "Response from `DigilockerIssuedFilesResponse.IssuedFiles.Data`: %v\n", *resp.Data)
    fmt.Fprintf(os.Stdout, "Response from `DigilockerIssuedFilesResponse.IssuedFiles.ResponseKey`: %v\n", *resp.ResponseKey)
}
```

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PullFile

Pull File

### Example

```go
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

    documentTypeRelatedParameters := *decentroinkyc.NewDigilockerPullFileRequestDocumentTypeRelatedParameters()
    documentTypeRelatedParameters.SetIdNumber("ABCDE1234F")
    documentTypeRelatedParameters.SetFullName("Manish Gupta")
    
    digilockerPullFileRequest := *decentroinkyc.NewDigilockerPullFileRequest()
    digilockerPullFileRequest.SetConsent(true)
    digilockerPullFileRequest.SetPurpose("for bank account purpose")
    digilockerPullFileRequest.SetReferenceId("0619a700-4739-4ffc-9bdc-feaca40108ba")
    digilockerPullFileRequest.SetDocumentType("PAN")
    digilockerPullFileRequest.SetDocumentTypeRelatedParameters(documentTypeRelatedParameters)
    
    request := client.DigilockerApi.PullFile(
        "initialDecentroTxnId_example",
        digilockerPullFileRequest,
    )
    
    resp, httpRes, err := request.Execute()

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DigilockerApi.PullFile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", httpRes)
    }
    // response from `PullFile`: DigilockerPullFileResponse
    fmt.Fprintf(os.Stdout, "Response from `DigilockerApi.PullFile`: %v\n", resp)
    fmt.Fprintf(os.Stdout, "Response from `DigilockerPullFileResponse.PullFile.DecentroTxnId`: %v\n", *resp.DecentroTxnId)
    fmt.Fprintf(os.Stdout, "Response from `DigilockerPullFileResponse.PullFile.Status`: %v\n", *resp.Status)
    fmt.Fprintf(os.Stdout, "Response from `DigilockerPullFileResponse.PullFile.ResponseCode`: %v\n", *resp.ResponseCode)
    fmt.Fprintf(os.Stdout, "Response from `DigilockerPullFileResponse.PullFile.Message`: %v\n", *resp.Message)
    fmt.Fprintf(os.Stdout, "Response from `DigilockerPullFileResponse.PullFile.ResponseKey`: %v\n", *resp.ResponseKey)
}
```

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RevokeToken

Revoke Token

### Example

```go
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

    request := client.DigilockerApi.RevokeToken(
        "initialDecentroTxnId_example",
    )
    
    resp, httpRes, err := request.Execute()

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DigilockerApi.RevokeToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", httpRes)
    }
    // response from `RevokeToken`: DigilockerRevokeTokenResponse
    fmt.Fprintf(os.Stdout, "Response from `DigilockerApi.RevokeToken`: %v\n", resp)
    fmt.Fprintf(os.Stdout, "Response from `DigilockerRevokeTokenResponse.RevokeToken.DecentroTxnId`: %v\n", *resp.DecentroTxnId)
    fmt.Fprintf(os.Stdout, "Response from `DigilockerRevokeTokenResponse.RevokeToken.Status`: %v\n", *resp.Status)
    fmt.Fprintf(os.Stdout, "Response from `DigilockerRevokeTokenResponse.RevokeToken.ResponseCode`: %v\n", *resp.ResponseCode)
    fmt.Fprintf(os.Stdout, "Response from `DigilockerRevokeTokenResponse.RevokeToken.Message`: %v\n", *resp.Message)
    fmt.Fprintf(os.Stdout, "Response from `DigilockerRevokeTokenResponse.RevokeToken.Data`: %v\n", *resp.Data)
    fmt.Fprintf(os.Stdout, "Response from `DigilockerRevokeTokenResponse.RevokeToken.ResponseKey`: %v\n", *resp.ResponseKey)
}
```

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

