# APIClient.KYCApi

All URIs are relative to *https://in.staging.decentro.tech*

Method | Path | Description
------------- | ------------- | -------------
[**CheckImageQuality**](KYCApi.md#CheckImageQuality) | **Post** /v2/kyc/forensics/image_quality | Image Quality Check
[**CheckPhotocopy**](KYCApi.md#CheckPhotocopy) | **Post** /v2/kyc/forensics/photocopy_check | Photocopy Check
[**CheckVideoLiveness**](KYCApi.md#CheckVideoLiveness) | **Post** /v2/kyc/forensics/video_liveness | Liveness Check
[**ClassifyDocument**](KYCApi.md#ClassifyDocument) | **Post** /v2/kyc/document_classification | ID Classification
[**ExtractText**](KYCApi.md#ExtractText) | **Post** /kyc/scan_extract/ocr | Scan &amp; Extract
[**MaskAadhaarUid**](KYCApi.md#MaskAadhaarUid) | **Post** /v2/kyc/identities/mask_aadhaar_uid | Aadhaar Masking
[**MatchFace**](KYCApi.md#MatchFace) | **Post** /v2/kyc/forensics/face_match | Face Match
[**Validate**](KYCApi.md#Validate) | **Post** /kyc/public_registry/validate | Validate



## CheckImageQuality

Image Quality Check

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

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CheckPhotocopy

Photocopy Check

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

    request := client.KYCApi.CheckPhotocopy(
        "referenceId_example",
        true,
        "consentPurpose_example",
    )
    request.Image(os.NewFile(1234, "some_file"))
    request.ImageUrl("imageUrl_example")
    
    resp, httpRes, err := request.Execute()

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KYCApi.CheckPhotocopy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CheckPhotocopy`: CheckPhotocopyResponse
    fmt.Fprintf(os.Stdout, "Response from `KYCApi.CheckPhotocopy`: %v\n", resp)
}
```

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CheckVideoLiveness

Liveness Check

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

    request := client.KYCApi.CheckVideoLiveness(
        "referenceId_example",
        "consent_example",
        "consentPurpose_example",
    )
    request.Video(os.NewFile(1234, "some_file"))
    request.VideoUrl("videoUrl_example")
    
    resp, httpRes, err := request.Execute()

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KYCApi.CheckVideoLiveness``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CheckVideoLiveness`: CheckVideoLivenessResponse
    fmt.Fprintf(os.Stdout, "Response from `KYCApi.CheckVideoLiveness`: %v\n", resp)
}
```

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ClassifyDocument

ID Classification

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

    request := client.KYCApi.ClassifyDocument(
        "referenceId_example",
        "documentType_example",
        true,
        "consentPurpose_example",
    )
    request.Document(os.NewFile(1234, "some_file"))
    request.DocumentUrl("documentUrl_example")
    
    resp, httpRes, err := request.Execute()

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KYCApi.ClassifyDocument``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ClassifyDocument`: ClassifyDocumentResponse
    fmt.Fprintf(os.Stdout, "Response from `KYCApi.ClassifyDocument`: %v\n", resp)
}
```

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExtractText

Scan & Extract

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

    request := client.KYCApi.ExtractText(
        "referenceId_example",
        "documentType_example",
        "consent_example",
        "consentPurpose_example",
        56,
    )
    request.Document(os.NewFile(1234, "some_file"))
    request.DocumentUrl("documentUrl_example")
    request.DocumentBack(os.NewFile(1234, "some_file"))
    request.DocumentBackUrl("documentBackUrl_example")
    
    resp, httpRes, err := request.Execute()

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KYCApi.ExtractText``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExtractText`: ExtractTextResponse
    fmt.Fprintf(os.Stdout, "Response from `KYCApi.ExtractText`: %v\n", resp)
}
```

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MaskAadhaarUid

Aadhaar Masking

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

    request := client.KYCApi.MaskAadhaarUid(
        "referenceId_example",
        true,
        "consentPurpose_example",
    )
    request.Image(os.NewFile(1234, "some_file"))
    request.ImageUrl("imageUrl_example")
    
    resp, httpRes, err := request.Execute()

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KYCApi.MaskAadhaarUid``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MaskAadhaarUid`: MaskAadhaarUidResponse
    fmt.Fprintf(os.Stdout, "Response from `KYCApi.MaskAadhaarUid`: %v\n", resp)
}
```

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MatchFace

Face Match

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

    request := client.KYCApi.MatchFace(
        "referenceId_example",
        "consent_example",
        "consentPurpose_example",
    )
    request.Image1(os.NewFile(1234, "some_file"))
    request.Image2(os.NewFile(1234, "some_file"))
    request.Threshold(56)
    request.Image1Url("image1Url_example")
    request.Image2Url("image2Url_example")
    
    resp, httpRes, err := request.Execute()

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KYCApi.MatchFace``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MatchFace`: MatchFaceResponse
    fmt.Fprintf(os.Stdout, "Response from `KYCApi.MatchFace`: %v\n", resp)
}
```

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Validate

Validate

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

    
    validateRequest := *decentroinkyc.NewValidateRequest(
        "ABCDEF12345",
        "PAN",
        "BNZPM2501F",
        "Y",
        "For Testing Purpose Only",
    )
    validateRequest.SetDob("1997-04-21")
    validateRequest.SetName("Radha Ramalinga")
    
    request := client.KYCApi.Validate(
        validateRequest,
    )
    
    resp, httpRes, err := request.Execute()

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KYCApi.Validate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Validate`: ValidateResponse
    fmt.Fprintf(os.Stdout, "Response from `KYCApi.Validate`: %v\n", resp)
}
```

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

