# APIClient.UpgradedAadhaarXmlApi

All URIs are relative to *https://in.staging.decentro.tech*

Method | Path | Description
------------- | ------------- | -------------
[**GenerateOtp**](UpgradedAadhaarXmlApi.md#GenerateOtp) | **Post** /v2/kyc/aadhaar/otp | Generate OTP
[**ValidateOtp**](UpgradedAadhaarXmlApi.md#ValidateOtp) | **Post** /v2/kyc/aadhaar/otp/validate | Validate OTP



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

    
    upgradedaadhaarxmlGenerateOtpRequest := *decentroinkyc.NewUpgradedaadhaarxmlGenerateOtpRequest()
    upgradedaadhaarxmlGenerateOtpRequest.SetReferenceId("0f2ef482-5802-42aa-bbf2-fbfb770188d4")
    upgradedaadhaarxmlGenerateOtpRequest.SetConsent(true)
    upgradedaadhaarxmlGenerateOtpRequest.SetPurpose("for bank account verification")
    upgradedaadhaarxmlGenerateOtpRequest.SetAadhaarNumber("898989898989")
    
    request := client.UpgradedAadhaarXmlApi.GenerateOtp(
        upgradedaadhaarxmlGenerateOtpRequest,
    )
    
    resp, httpRes, err := request.Execute()

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UpgradedAadhaarXmlApi.GenerateOtp``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", httpRes)
    }
    // response from `GenerateOtp`: UpgradedaadhaarxmlGenerateOtpResponse
    fmt.Fprintf(os.Stdout, "Response from `UpgradedAadhaarXmlApi.GenerateOtp`: %v\n", resp)
    fmt.Fprintf(os.Stdout, "Response from `UpgradedaadhaarxmlGenerateOtpResponse.GenerateOtp.DecentroTxnId`: %v\n", *resp.DecentroTxnId)
    fmt.Fprintf(os.Stdout, "Response from `UpgradedaadhaarxmlGenerateOtpResponse.GenerateOtp.Status`: %v\n", *resp.Status)
    fmt.Fprintf(os.Stdout, "Response from `UpgradedaadhaarxmlGenerateOtpResponse.GenerateOtp.ResponseCode`: %v\n", *resp.ResponseCode)
    fmt.Fprintf(os.Stdout, "Response from `UpgradedaadhaarxmlGenerateOtpResponse.GenerateOtp.Message`: %v\n", *resp.Message)
    fmt.Fprintf(os.Stdout, "Response from `UpgradedaadhaarxmlGenerateOtpResponse.GenerateOtp.ResponseKey`: %v\n", *resp.ResponseKey)
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
    client := decentroinkyc.NewAPIClient(configuration)

    
    upgradedaadhaarxmlValidateOtpRequest := *decentroinkyc.NewUpgradedaadhaarxmlValidateOtpRequest()
    upgradedaadhaarxmlValidateOtpRequest.SetReferenceId("683edd8c-c76a-4dd9-9275-c79b82f9ee36")
    upgradedaadhaarxmlValidateOtpRequest.SetConsent(true)
    upgradedaadhaarxmlValidateOtpRequest.SetPurpose("for bank account verification")
    upgradedaadhaarxmlValidateOtpRequest.SetInitiationTransactionId("DXXXXXXXXXXXXXXXXXXXX")
    upgradedaadhaarxmlValidateOtpRequest.SetOtp("136668")
    upgradedaadhaarxmlValidateOtpRequest.SetShareCode("1111")
    upgradedaadhaarxmlValidateOtpRequest.SetGeneratePdf(false)
    upgradedaadhaarxmlValidateOtpRequest.SetGenerateXml(false)
    
    request := client.UpgradedAadhaarXmlApi.ValidateOtp(
        upgradedaadhaarxmlValidateOtpRequest,
    )
    
    resp, httpRes, err := request.Execute()

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UpgradedAadhaarXmlApi.ValidateOtp``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", httpRes)
    }
    // response from `ValidateOtp`: UpgradedaadhaarxmlValidateOtpResponse
    fmt.Fprintf(os.Stdout, "Response from `UpgradedAadhaarXmlApi.ValidateOtp`: %v\n", resp)
    fmt.Fprintf(os.Stdout, "Response from `UpgradedaadhaarxmlValidateOtpResponse.ValidateOtp.DecentroTxnId`: %v\n", *resp.DecentroTxnId)
    fmt.Fprintf(os.Stdout, "Response from `UpgradedaadhaarxmlValidateOtpResponse.ValidateOtp.Status`: %v\n", *resp.Status)
    fmt.Fprintf(os.Stdout, "Response from `UpgradedaadhaarxmlValidateOtpResponse.ValidateOtp.ResponseCode`: %v\n", *resp.ResponseCode)
    fmt.Fprintf(os.Stdout, "Response from `UpgradedaadhaarxmlValidateOtpResponse.ValidateOtp.Message`: %v\n", *resp.Message)
    fmt.Fprintf(os.Stdout, "Response from `UpgradedaadhaarxmlValidateOtpResponse.ValidateOtp.Data`: %v\n", *resp.Data)
    fmt.Fprintf(os.Stdout, "Response from `UpgradedaadhaarxmlValidateOtpResponse.ValidateOtp.ResponseKey`: %v\n", *resp.ResponseKey)
}
```

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

