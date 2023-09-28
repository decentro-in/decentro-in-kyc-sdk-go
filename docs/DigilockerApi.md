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
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DownloadFile`: string
    fmt.Fprintf(os.Stdout, "Response from `DigilockerApi.DownloadFile`: %v\n", resp)
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
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EAadhaar`: string
    fmt.Fprintf(os.Stdout, "Response from `DigilockerApi.EAadhaar`: %v\n", resp)
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
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FileData`: string
    fmt.Fprintf(os.Stdout, "Response from `DigilockerApi.FileData`: %v\n", resp)
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
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GenerateAccessToken`: string
    fmt.Fprintf(os.Stdout, "Response from `DigilockerApi.GenerateAccessToken`: %v\n", resp)
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
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `InitiateSession`: string
    fmt.Fprintf(os.Stdout, "Response from `DigilockerApi.InitiateSession`: %v\n", resp)
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
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IssuedFiles`: string
    fmt.Fprintf(os.Stdout, "Response from `DigilockerApi.IssuedFiles`: %v\n", resp)
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
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PullFile`: string
    fmt.Fprintf(os.Stdout, "Response from `DigilockerApi.PullFile`: %v\n", resp)
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

    request := client.DigilockerApi.RevokeToken(
        "initialDecentroTxnId_example",
    )
    
    resp, httpRes, err := request.Execute()

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DigilockerApi.RevokeToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RevokeToken`: string
    fmt.Fprintf(os.Stdout, "Response from `DigilockerApi.RevokeToken`: %v\n", resp)
}
```

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

