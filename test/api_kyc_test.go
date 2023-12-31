/*
decentro-in-kyc

Testing KYCApiService

*/


package decentroinkyc

import (
    "testing"
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/require"
    decentroinkyc "github.com/decentro-in/decentro-in-kyc-sdk-go"
    "os"
)


func Test_decentroinkyc_KYCApiService(t *testing.T) {

    configuration := decentroinkyc.NewConfiguration()
    configuration.SetHost("http://127.0.0.1:4010")
    configuration.SetClientId("CLIENT_ID")
    configuration.SetClientSecret("CLIENT_SECRET")
    configuration.SetModuleSecret("MODULE_SECRET")
    client := decentroinkyc.NewAPIClient(configuration)

    t.Run("Test KYCApiService CheckImageQuality", func(t *testing.T) {

        request := client.KYCApi.CheckImageQuality(
            "referenceId_example",
            true,
            "consentPurpose_example",
        )
        request.Image(os.NewFile(1234, "some_file"))
        request.QualityParameter("qualityParameter_example")
        request.ImageUrl("imageUrl_example")
        
        resp, httpRes, err := request.Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test KYCApiService CheckPhotocopy", func(t *testing.T) {

        request := client.KYCApi.CheckPhotocopy(
            "referenceId_example",
            true,
            "consentPurpose_example",
        )
        request.Image(os.NewFile(1234, "some_file"))
        request.ImageUrl("imageUrl_example")
        
        resp, httpRes, err := request.Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test KYCApiService CheckVideoLiveness", func(t *testing.T) {

        request := client.KYCApi.CheckVideoLiveness(
            "referenceId_example",
            "consent_example",
            "consentPurpose_example",
        )
        request.Video(os.NewFile(1234, "some_file"))
        request.VideoUrl("videoUrl_example")
        
        resp, httpRes, err := request.Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test KYCApiService ClassifyDocument", func(t *testing.T) {

        request := client.KYCApi.ClassifyDocument(
            "referenceId_example",
            "documentType_example",
            true,
            "consentPurpose_example",
        )
        request.Document(os.NewFile(1234, "some_file"))
        request.DocumentUrl("documentUrl_example")
        
        resp, httpRes, err := request.Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test KYCApiService ExtractText", func(t *testing.T) {

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

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test KYCApiService MaskAadhaarUid", func(t *testing.T) {

        request := client.KYCApi.MaskAadhaarUid(
            "referenceId_example",
            true,
            "consentPurpose_example",
        )
        request.Image(os.NewFile(1234, "some_file"))
        request.ImageUrl("imageUrl_example")
        
        resp, httpRes, err := request.Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test KYCApiService MatchFace", func(t *testing.T) {

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

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test KYCApiService Validate", func(t *testing.T) {

        
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

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

}