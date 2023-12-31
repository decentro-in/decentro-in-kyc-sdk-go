# APIClient.AadhaarXmlApi

All URIs are relative to *https://in.staging.decentro.tech*

Method | Path | Description
------------- | ------------- | -------------
[**GenerateOtp**](AadhaarXmlApi.md#GenerateOtp) | **Post** /v2/kyc/aadhaar_connect/otp | Generate OTP
[**InitiateSession**](AadhaarXmlApi.md#InitiateSession) | **Post** /v2/kyc/aadhaar_connect | Initiate Session
[**ReloadCaptcha**](AadhaarXmlApi.md#ReloadCaptcha) | **Post** /v2/kyc/aadhaar_connect/captcha/reload | Reload Captcha
[**ValidateOtp**](AadhaarXmlApi.md#ValidateOtp) | **Post** /v2/kyc/aadhaar_connect/otp/validate | Validate OTP



## GenerateOtp

Generate OTP

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

    
    aadhaarxmlGenerateOtpRequest := *decentroinkyc.NewAadhaarxmlGenerateOtpRequest()
    aadhaarxmlGenerateOtpRequest.SetReferenceId("ABCDEF12345")
    aadhaarxmlGenerateOtpRequest.SetConsent(true)
    aadhaarxmlGenerateOtpRequest.SetPurpose("For Aadhaar Verification")
    aadhaarxmlGenerateOtpRequest.SetInitiationTransactionId("D009XXXXXXXXXXXXXXXXXX")
    aadhaarxmlGenerateOtpRequest.SetAadhaarNumber("3243XXXXXXXX")
    aadhaarxmlGenerateOtpRequest.SetCaptcha("xxxxx")
    
    request := client.AadhaarXmlApi.GenerateOtp(
        aadhaarxmlGenerateOtpRequest,
    )
    
    resp, httpRes, err := request.Execute()

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AadhaarXmlApi.GenerateOtp``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", httpRes)
    }
    // response from `GenerateOtp`: AadhaarxmlGenerateOtpResponse
    fmt.Fprintf(os.Stdout, "Response from `AadhaarXmlApi.GenerateOtp`: %v\n", resp)
    fmt.Fprintf(os.Stdout, "Response from `AadhaarxmlGenerateOtpResponse.GenerateOtp.DecentroTxnId`: %v\n", *resp.DecentroTxnId)
    fmt.Fprintf(os.Stdout, "Response from `AadhaarxmlGenerateOtpResponse.GenerateOtp.Status`: %v\n", *resp.Status)
    fmt.Fprintf(os.Stdout, "Response from `AadhaarxmlGenerateOtpResponse.GenerateOtp.ResponseCode`: %v\n", *resp.ResponseCode)
    fmt.Fprintf(os.Stdout, "Response from `AadhaarxmlGenerateOtpResponse.GenerateOtp.Message`: %v\n", *resp.Message)
    fmt.Fprintf(os.Stdout, "Response from `AadhaarxmlGenerateOtpResponse.GenerateOtp.ResponseKey`: %v\n", *resp.ResponseKey)
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

    
    aadhaarxmlInitiateSessionRequest := *decentroinkyc.NewAadhaarxmlInitiateSessionRequest()
    aadhaarxmlInitiateSessionRequest.SetReferenceId("ABCDEF12345")
    aadhaarxmlInitiateSessionRequest.SetConsent(true)
    aadhaarxmlInitiateSessionRequest.SetPurpose("For Aadhaar verification")
    
    request := client.AadhaarXmlApi.InitiateSession(
        aadhaarxmlInitiateSessionRequest,
    )
    
    resp, httpRes, err := request.Execute()

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AadhaarXmlApi.InitiateSession``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", httpRes)
    }
    // response from `InitiateSession`: AadhaarxmlInitiateSessionResponse
    fmt.Fprintf(os.Stdout, "Response from `AadhaarXmlApi.InitiateSession`: %v\n", resp)
    fmt.Fprintf(os.Stdout, "Response from `AadhaarxmlInitiateSessionResponse.InitiateSession.DecentroTxnId`: %v\n", *resp.DecentroTxnId)
    fmt.Fprintf(os.Stdout, "Response from `AadhaarxmlInitiateSessionResponse.InitiateSession.Status`: %v\n", *resp.Status)
    fmt.Fprintf(os.Stdout, "Response from `AadhaarxmlInitiateSessionResponse.InitiateSession.ResponseCode`: %v\n", *resp.ResponseCode)
    fmt.Fprintf(os.Stdout, "Response from `AadhaarxmlInitiateSessionResponse.InitiateSession.Message`: %v\n", *resp.Message)
    fmt.Fprintf(os.Stdout, "Response from `AadhaarxmlInitiateSessionResponse.InitiateSession.Data`: %v\n", *resp.Data)
    fmt.Fprintf(os.Stdout, "Response from `AadhaarxmlInitiateSessionResponse.InitiateSession.ResponseKey`: %v\n", *resp.ResponseKey)
}
```

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReloadCaptcha

Reload Captcha

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

    
    aadhaarxmlReloadCaptchaRequest := *decentroinkyc.NewAadhaarxmlReloadCaptchaRequest()
    aadhaarxmlReloadCaptchaRequest.SetReferenceId("ABCDEF12345")
    aadhaarxmlReloadCaptchaRequest.SetConsent(true)
    aadhaarxmlReloadCaptchaRequest.SetPurpose("For Aadhaar Verification")
    aadhaarxmlReloadCaptchaRequest.SetInitiationTransactionId("D009XXXXXXXXXXXXXXXXXX")
    
    request := client.AadhaarXmlApi.ReloadCaptcha(
        aadhaarxmlReloadCaptchaRequest,
    )
    
    resp, httpRes, err := request.Execute()

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AadhaarXmlApi.ReloadCaptcha``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", httpRes)
    }
    // response from `ReloadCaptcha`: AadhaarxmlReloadCaptchaResponse
    fmt.Fprintf(os.Stdout, "Response from `AadhaarXmlApi.ReloadCaptcha`: %v\n", resp)
    fmt.Fprintf(os.Stdout, "Response from `AadhaarxmlReloadCaptchaResponse.ReloadCaptcha.DecentroTxnId`: %v\n", *resp.DecentroTxnId)
    fmt.Fprintf(os.Stdout, "Response from `AadhaarxmlReloadCaptchaResponse.ReloadCaptcha.Status`: %v\n", *resp.Status)
    fmt.Fprintf(os.Stdout, "Response from `AadhaarxmlReloadCaptchaResponse.ReloadCaptcha.ResponseCode`: %v\n", *resp.ResponseCode)
    fmt.Fprintf(os.Stdout, "Response from `AadhaarxmlReloadCaptchaResponse.ReloadCaptcha.Message`: %v\n", *resp.Message)
    fmt.Fprintf(os.Stdout, "Response from `AadhaarxmlReloadCaptchaResponse.ReloadCaptcha.Data`: %v\n", *resp.Data)
    fmt.Fprintf(os.Stdout, "Response from `AadhaarxmlReloadCaptchaResponse.ReloadCaptcha.ResponseKey`: %v\n", *resp.ResponseKey)
}
```

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ValidateOtp

Validate OTP

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

    
    aadhaarxmlValidateOtpRequest := *decentroinkyc.NewAadhaarxmlValidateOtpRequest()
    aadhaarxmlValidateOtpRequest.SetReferenceId("ABCDEF12345")
    aadhaarxmlValidateOtpRequest.SetConsent(true)
    aadhaarxmlValidateOtpRequest.SetPurpose("For Aadhaar Verification")
    aadhaarxmlValidateOtpRequest.SetInitiationTransactionId("D009XXXXXXXXXXXXXXXXXX")
    aadhaarxmlValidateOtpRequest.SetGeneratePdf(false)
    aadhaarxmlValidateOtpRequest.SetGenerateXml(false)
    aadhaarxmlValidateOtpRequest.SetShareCode(false)
    aadhaarxmlValidateOtpRequest.SetOtp("XXXXXX")
    
    request := client.AadhaarXmlApi.ValidateOtp(
        aadhaarxmlValidateOtpRequest,
    )
    
    resp, httpRes, err := request.Execute()

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AadhaarXmlApi.ValidateOtp``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", httpRes)
    }
    // response from `ValidateOtp`: AadhaarxmlValidateOtpResponse
    fmt.Fprintf(os.Stdout, "Response from `AadhaarXmlApi.ValidateOtp`: %v\n", resp)
    fmt.Fprintf(os.Stdout, "Response from `AadhaarxmlValidateOtpResponse.ValidateOtp.DecentroTxnId`: %v\n", *resp.DecentroTxnId)
    fmt.Fprintf(os.Stdout, "Response from `AadhaarxmlValidateOtpResponse.ValidateOtp.Status`: %v\n", *resp.Status)
    fmt.Fprintf(os.Stdout, "Response from `AadhaarxmlValidateOtpResponse.ValidateOtp.ResponseCode`: %v\n", *resp.ResponseCode)
    fmt.Fprintf(os.Stdout, "Response from `AadhaarxmlValidateOtpResponse.ValidateOtp.Message`: %v\n", *resp.Message)
    fmt.Fprintf(os.Stdout, "Response from `AadhaarxmlValidateOtpResponse.ValidateOtp.Data`: %v\n", *resp.Data)
    fmt.Fprintf(os.Stdout, "Response from `AadhaarxmlValidateOtpResponse.ValidateOtp.ResponseKey`: %v\n", *resp.ResponseKey)
}
```

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

