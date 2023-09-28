# ExtractTextRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReferenceId** | **string** |  | 
**DocumentType** | **string** |  | 
**Consent** | **string** |  | 
**ConsentPurpose** | **string** |  | 
**KycValidate** | **int32** |  | 
**Document** | Pointer to ***os.File** |  | [optional] 
**DocumentUrl** | Pointer to **string** |  | [optional] 
**DocumentBack** | Pointer to ***os.File** |  | [optional] 
**DocumentBackUrl** | Pointer to **string** |  | [optional] 

## Methods

### NewExtractTextRequest

`func NewExtractTextRequest(referenceId string, documentType string, consent string, consentPurpose string, kycValidate int32, ) *ExtractTextRequest`

NewExtractTextRequest instantiates a new ExtractTextRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExtractTextRequestWithDefaults

`func NewExtractTextRequestWithDefaults() *ExtractTextRequest`

NewExtractTextRequestWithDefaults instantiates a new ExtractTextRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReferenceId

`func (o *ExtractTextRequest) GetReferenceId() string`

GetReferenceId returns the ReferenceId field if non-nil, zero value otherwise.

### GetReferenceIdOk

`func (o *ExtractTextRequest) GetReferenceIdOk() (*string, bool)`

GetReferenceIdOk returns a tuple with the ReferenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceId

`func (o *ExtractTextRequest) SetReferenceId(v string)`

SetReferenceId sets ReferenceId field to given value.


### GetDocumentType

`func (o *ExtractTextRequest) GetDocumentType() string`

GetDocumentType returns the DocumentType field if non-nil, zero value otherwise.

### GetDocumentTypeOk

`func (o *ExtractTextRequest) GetDocumentTypeOk() (*string, bool)`

GetDocumentTypeOk returns a tuple with the DocumentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentType

`func (o *ExtractTextRequest) SetDocumentType(v string)`

SetDocumentType sets DocumentType field to given value.


### GetConsent

`func (o *ExtractTextRequest) GetConsent() string`

GetConsent returns the Consent field if non-nil, zero value otherwise.

### GetConsentOk

`func (o *ExtractTextRequest) GetConsentOk() (*string, bool)`

GetConsentOk returns a tuple with the Consent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsent

`func (o *ExtractTextRequest) SetConsent(v string)`

SetConsent sets Consent field to given value.


### GetConsentPurpose

`func (o *ExtractTextRequest) GetConsentPurpose() string`

GetConsentPurpose returns the ConsentPurpose field if non-nil, zero value otherwise.

### GetConsentPurposeOk

`func (o *ExtractTextRequest) GetConsentPurposeOk() (*string, bool)`

GetConsentPurposeOk returns a tuple with the ConsentPurpose field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsentPurpose

`func (o *ExtractTextRequest) SetConsentPurpose(v string)`

SetConsentPurpose sets ConsentPurpose field to given value.


### GetKycValidate

`func (o *ExtractTextRequest) GetKycValidate() int32`

GetKycValidate returns the KycValidate field if non-nil, zero value otherwise.

### GetKycValidateOk

`func (o *ExtractTextRequest) GetKycValidateOk() (*int32, bool)`

GetKycValidateOk returns a tuple with the KycValidate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKycValidate

`func (o *ExtractTextRequest) SetKycValidate(v int32)`

SetKycValidate sets KycValidate field to given value.


### GetDocument

`func (o *ExtractTextRequest) GetDocument() *os.File`

GetDocument returns the Document field if non-nil, zero value otherwise.

### GetDocumentOk

`func (o *ExtractTextRequest) GetDocumentOk() (**os.File, bool)`

GetDocumentOk returns a tuple with the Document field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocument

`func (o *ExtractTextRequest) SetDocument(v *os.File)`

SetDocument sets Document field to given value.

### HasDocument

`func (o *ExtractTextRequest) HasDocument() bool`

HasDocument returns a boolean if a field has been set.

### GetDocumentUrl

`func (o *ExtractTextRequest) GetDocumentUrl() string`

GetDocumentUrl returns the DocumentUrl field if non-nil, zero value otherwise.

### GetDocumentUrlOk

`func (o *ExtractTextRequest) GetDocumentUrlOk() (*string, bool)`

GetDocumentUrlOk returns a tuple with the DocumentUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentUrl

`func (o *ExtractTextRequest) SetDocumentUrl(v string)`

SetDocumentUrl sets DocumentUrl field to given value.

### HasDocumentUrl

`func (o *ExtractTextRequest) HasDocumentUrl() bool`

HasDocumentUrl returns a boolean if a field has been set.

### GetDocumentBack

`func (o *ExtractTextRequest) GetDocumentBack() *os.File`

GetDocumentBack returns the DocumentBack field if non-nil, zero value otherwise.

### GetDocumentBackOk

`func (o *ExtractTextRequest) GetDocumentBackOk() (**os.File, bool)`

GetDocumentBackOk returns a tuple with the DocumentBack field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentBack

`func (o *ExtractTextRequest) SetDocumentBack(v *os.File)`

SetDocumentBack sets DocumentBack field to given value.

### HasDocumentBack

`func (o *ExtractTextRequest) HasDocumentBack() bool`

HasDocumentBack returns a boolean if a field has been set.

### GetDocumentBackUrl

`func (o *ExtractTextRequest) GetDocumentBackUrl() string`

GetDocumentBackUrl returns the DocumentBackUrl field if non-nil, zero value otherwise.

### GetDocumentBackUrlOk

`func (o *ExtractTextRequest) GetDocumentBackUrlOk() (*string, bool)`

GetDocumentBackUrlOk returns a tuple with the DocumentBackUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentBackUrl

`func (o *ExtractTextRequest) SetDocumentBackUrl(v string)`

SetDocumentBackUrl sets DocumentBackUrl field to given value.

### HasDocumentBackUrl

`func (o *ExtractTextRequest) HasDocumentBackUrl() bool`

HasDocumentBackUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


