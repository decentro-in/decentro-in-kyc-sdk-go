/*
decentro-in-kyc

Testing AadhaarXmlApiService

*/


package decentroinkyc

import (
    "testing"
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/require"
    decentroinkyc "github.com/decentro-in/decentro-in-kyc-sdk-go"
)

func Test_decentroinkyc_AadhaarXmlApiService(t *testing.T) {

    configuration := decentroinkyc.NewConfiguration()
    configuration.SetHost("http://127.0.0.1:4010")
    configuration.SetClientId("CLIENT_ID")
    configuration.SetClientSecret("CLIENT_SECRET")
    configuration.SetModuleSecret("MODULE_SECRET")
    client := decentroinkyc.NewAPIClient(configuration)

    t.Run("Test AadhaarXmlApiService GenerateOtp", func(t *testing.T) {

        
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

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test AadhaarXmlApiService InitiateSession", func(t *testing.T) {

        
        aadhaarxmlInitiateSessionRequest := *decentroinkyc.NewAadhaarxmlInitiateSessionRequest()
        aadhaarxmlInitiateSessionRequest.SetReferenceId("ABCDEF12345")
        aadhaarxmlInitiateSessionRequest.SetConsent(true)
        aadhaarxmlInitiateSessionRequest.SetPurpose("For Aadhaar verification")
        
        request := client.AadhaarXmlApi.InitiateSession(
            aadhaarxmlInitiateSessionRequest,
        )
        
        resp, httpRes, err := request.Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test AadhaarXmlApiService ReloadCaptcha", func(t *testing.T) {

        
        aadhaarxmlReloadCaptchaRequest := *decentroinkyc.NewAadhaarxmlReloadCaptchaRequest()
        aadhaarxmlReloadCaptchaRequest.SetReferenceId("ABCDEF12345")
        aadhaarxmlReloadCaptchaRequest.SetConsent(true)
        aadhaarxmlReloadCaptchaRequest.SetPurpose("For Aadhaar Verification")
        aadhaarxmlReloadCaptchaRequest.SetInitiationTransactionId("D009XXXXXXXXXXXXXXXXXX")
        
        request := client.AadhaarXmlApi.ReloadCaptcha(
            aadhaarxmlReloadCaptchaRequest,
        )
        
        resp, httpRes, err := request.Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test AadhaarXmlApiService ValidateOtp", func(t *testing.T) {

        
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

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

}