/*
decentro-in-kyc

Testing DigilockerApiService

*/


package decentroinkyc

import (
    "testing"
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/require"
    decentroinkyc "github.com/decentro-in/decentro-in-kyc-sdk-go"
)

func Test_decentroinkyc_DigilockerApiService(t *testing.T) {

    configuration := decentroinkyc.NewConfiguration()
    configuration.SetHost("http://127.0.0.1:4010")
    configuration.SetClientId("CLIENT_ID")
    configuration.SetClientSecret("CLIENT_SECRET")
    configuration.SetModuleSecret("MODULE_SECRET")
    client := decentroinkyc.NewAPIClient(configuration)

    t.Run("Test DigilockerApiService DownloadFile", func(t *testing.T) {

        
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

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test DigilockerApiService EAadhaar", func(t *testing.T) {

        
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

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test DigilockerApiService FileData", func(t *testing.T) {

        
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

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test DigilockerApiService GenerateAccessToken", func(t *testing.T) {

        
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

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test DigilockerApiService InitiateSession", func(t *testing.T) {

        
        digilockerInitiateSessionRequest := *decentroinkyc.NewDigilockerInitiateSessionRequest()
        digilockerInitiateSessionRequest.SetConsent(true)
        digilockerInitiateSessionRequest.SetPurpose("for bank account verification")
        digilockerInitiateSessionRequest.SetReferenceId("ABCDEF123456")
        digilockerInitiateSessionRequest.SetRedirectUrl("<Redirect URL for DigiLocker>")
        
        request := client.DigilockerApi.InitiateSession(
            digilockerInitiateSessionRequest,
        )
        
        resp, httpRes, err := request.Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test DigilockerApiService IssuedFiles", func(t *testing.T) {

        
        digilockerIssuedFilesRequest := *decentroinkyc.NewDigilockerIssuedFilesRequest()
        digilockerIssuedFilesRequest.SetConsent(true)
        digilockerIssuedFilesRequest.SetPurpose("for bank account verification")
        digilockerIssuedFilesRequest.SetReferenceId("ABCDEF123456")
        
        request := client.DigilockerApi.IssuedFiles(
            "initialDecentroTxnId_example",
            digilockerIssuedFilesRequest,
        )
        
        resp, httpRes, err := request.Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test DigilockerApiService PullFile", func(t *testing.T) {

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

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test DigilockerApiService RevokeToken", func(t *testing.T) {

        request := client.DigilockerApi.RevokeToken(
            "initialDecentroTxnId_example",
        )
        
        resp, httpRes, err := request.Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

}