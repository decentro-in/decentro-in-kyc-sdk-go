# AadhaarxmlValidateOtpRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReferenceId** | Pointer to **string** |  | [optional] 
**Consent** | Pointer to **bool** |  | [optional] 
**Purpose** | Pointer to **string** |  | [optional] 
**InitiationTransactionId** | Pointer to **string** |  | [optional] 
**GeneratePdf** | Pointer to **bool** |  | [optional] 
**GenerateXml** | Pointer to **bool** |  | [optional] 
**ShareCode** | Pointer to **bool** |  | [optional] 
**Otp** | Pointer to **string** |  | [optional] 

## Methods

### NewAadhaarxmlValidateOtpRequest

`func NewAadhaarxmlValidateOtpRequest() *AadhaarxmlValidateOtpRequest`

NewAadhaarxmlValidateOtpRequest instantiates a new AadhaarxmlValidateOtpRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAadhaarxmlValidateOtpRequestWithDefaults

`func NewAadhaarxmlValidateOtpRequestWithDefaults() *AadhaarxmlValidateOtpRequest`

NewAadhaarxmlValidateOtpRequestWithDefaults instantiates a new AadhaarxmlValidateOtpRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReferenceId

`func (o *AadhaarxmlValidateOtpRequest) GetReferenceId() string`

GetReferenceId returns the ReferenceId field if non-nil, zero value otherwise.

### GetReferenceIdOk

`func (o *AadhaarxmlValidateOtpRequest) GetReferenceIdOk() (*string, bool)`

GetReferenceIdOk returns a tuple with the ReferenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceId

`func (o *AadhaarxmlValidateOtpRequest) SetReferenceId(v string)`

SetReferenceId sets ReferenceId field to given value.

### HasReferenceId

`func (o *AadhaarxmlValidateOtpRequest) HasReferenceId() bool`

HasReferenceId returns a boolean if a field has been set.

### GetConsent

`func (o *AadhaarxmlValidateOtpRequest) GetConsent() bool`

GetConsent returns the Consent field if non-nil, zero value otherwise.

### GetConsentOk

`func (o *AadhaarxmlValidateOtpRequest) GetConsentOk() (*bool, bool)`

GetConsentOk returns a tuple with the Consent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsent

`func (o *AadhaarxmlValidateOtpRequest) SetConsent(v bool)`

SetConsent sets Consent field to given value.

### HasConsent

`func (o *AadhaarxmlValidateOtpRequest) HasConsent() bool`

HasConsent returns a boolean if a field has been set.

### GetPurpose

`func (o *AadhaarxmlValidateOtpRequest) GetPurpose() string`

GetPurpose returns the Purpose field if non-nil, zero value otherwise.

### GetPurposeOk

`func (o *AadhaarxmlValidateOtpRequest) GetPurposeOk() (*string, bool)`

GetPurposeOk returns a tuple with the Purpose field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurpose

`func (o *AadhaarxmlValidateOtpRequest) SetPurpose(v string)`

SetPurpose sets Purpose field to given value.

### HasPurpose

`func (o *AadhaarxmlValidateOtpRequest) HasPurpose() bool`

HasPurpose returns a boolean if a field has been set.

### GetInitiationTransactionId

`func (o *AadhaarxmlValidateOtpRequest) GetInitiationTransactionId() string`

GetInitiationTransactionId returns the InitiationTransactionId field if non-nil, zero value otherwise.

### GetInitiationTransactionIdOk

`func (o *AadhaarxmlValidateOtpRequest) GetInitiationTransactionIdOk() (*string, bool)`

GetInitiationTransactionIdOk returns a tuple with the InitiationTransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitiationTransactionId

`func (o *AadhaarxmlValidateOtpRequest) SetInitiationTransactionId(v string)`

SetInitiationTransactionId sets InitiationTransactionId field to given value.

### HasInitiationTransactionId

`func (o *AadhaarxmlValidateOtpRequest) HasInitiationTransactionId() bool`

HasInitiationTransactionId returns a boolean if a field has been set.

### GetGeneratePdf

`func (o *AadhaarxmlValidateOtpRequest) GetGeneratePdf() bool`

GetGeneratePdf returns the GeneratePdf field if non-nil, zero value otherwise.

### GetGeneratePdfOk

`func (o *AadhaarxmlValidateOtpRequest) GetGeneratePdfOk() (*bool, bool)`

GetGeneratePdfOk returns a tuple with the GeneratePdf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeneratePdf

`func (o *AadhaarxmlValidateOtpRequest) SetGeneratePdf(v bool)`

SetGeneratePdf sets GeneratePdf field to given value.

### HasGeneratePdf

`func (o *AadhaarxmlValidateOtpRequest) HasGeneratePdf() bool`

HasGeneratePdf returns a boolean if a field has been set.

### GetGenerateXml

`func (o *AadhaarxmlValidateOtpRequest) GetGenerateXml() bool`

GetGenerateXml returns the GenerateXml field if non-nil, zero value otherwise.

### GetGenerateXmlOk

`func (o *AadhaarxmlValidateOtpRequest) GetGenerateXmlOk() (*bool, bool)`

GetGenerateXmlOk returns a tuple with the GenerateXml field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenerateXml

`func (o *AadhaarxmlValidateOtpRequest) SetGenerateXml(v bool)`

SetGenerateXml sets GenerateXml field to given value.

### HasGenerateXml

`func (o *AadhaarxmlValidateOtpRequest) HasGenerateXml() bool`

HasGenerateXml returns a boolean if a field has been set.

### GetShareCode

`func (o *AadhaarxmlValidateOtpRequest) GetShareCode() bool`

GetShareCode returns the ShareCode field if non-nil, zero value otherwise.

### GetShareCodeOk

`func (o *AadhaarxmlValidateOtpRequest) GetShareCodeOk() (*bool, bool)`

GetShareCodeOk returns a tuple with the ShareCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShareCode

`func (o *AadhaarxmlValidateOtpRequest) SetShareCode(v bool)`

SetShareCode sets ShareCode field to given value.

### HasShareCode

`func (o *AadhaarxmlValidateOtpRequest) HasShareCode() bool`

HasShareCode returns a boolean if a field has been set.

### GetOtp

`func (o *AadhaarxmlValidateOtpRequest) GetOtp() string`

GetOtp returns the Otp field if non-nil, zero value otherwise.

### GetOtpOk

`func (o *AadhaarxmlValidateOtpRequest) GetOtpOk() (*string, bool)`

GetOtpOk returns a tuple with the Otp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtp

`func (o *AadhaarxmlValidateOtpRequest) SetOtp(v string)`

SetOtp sets Otp field to given value.

### HasOtp

`func (o *AadhaarxmlValidateOtpRequest) HasOtp() bool`

HasOtp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


