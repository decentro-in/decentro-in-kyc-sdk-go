/*
decentro-in-kyc

Testing UpgradedAadhaarXmlApiService

*/


package decentroinkyc

import (
    "testing"
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/require"
    decentroinkyc "github.com/decentro-in/decentro-in-kyc-sdk-go"
)

func Test_decentroinkyc_UpgradedAadhaarXmlApiService(t *testing.T) {

    configuration := decentroinkyc.NewConfiguration()
    configuration.SetHost("http://127.0.0.1:4010")
    configuration.SetClientId("CLIENT_ID")
    configuration.SetClientSecret("CLIENT_SECRET")
    configuration.SetModuleSecret("MODULE_SECRET")
    client := decentroinkyc.NewAPIClient(configuration)

    t.Run("Test UpgradedAadhaarXmlApiService GenerateOtp", func(t *testing.T) {

        
        upgradedaadhaarxmlGenerateOtpRequest := *decentroinkyc.NewUpgradedaadhaarxmlGenerateOtpRequest()
        upgradedaadhaarxmlGenerateOtpRequest.SetReferenceId("0f2ef482-5802-42aa-bbf2-fbfb770188d4")
        upgradedaadhaarxmlGenerateOtpRequest.SetConsent(true)
        upgradedaadhaarxmlGenerateOtpRequest.SetPurpose("for bank account verification")
        upgradedaadhaarxmlGenerateOtpRequest.SetAadhaarNumber("898989898989")
        
        request := client.UpgradedAadhaarXmlApi.GenerateOtp(
            upgradedaadhaarxmlGenerateOtpRequest,
        )
        
        resp, httpRes, err := request.Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test UpgradedAadhaarXmlApiService ValidateOtp", func(t *testing.T) {

        
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

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

}