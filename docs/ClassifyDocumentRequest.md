# ClassifyDocumentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReferenceId** | **string** |  | 
**DocumentType** | **string** |  | 
**Consent** | **bool** |  | 
**ConsentPurpose** | **string** |  | 
**Document** | Pointer to ***os.File** |  | [optional] 
**DocumentUrl** | Pointer to **string** |  | [optional] 

## Methods

### NewClassifyDocumentRequest

`func NewClassifyDocumentRequest(referenceId string, documentType string, consent bool, consentPurpose string, ) *ClassifyDocumentRequest`

NewClassifyDocumentRequest instantiates a new ClassifyDocumentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClassifyDocumentRequestWithDefaults

`func NewClassifyDocumentRequestWithDefaults() *ClassifyDocumentRequest`

NewClassifyDocumentRequestWithDefaults instantiates a new ClassifyDocumentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReferenceId

`func (o *ClassifyDocumentRequest) GetReferenceId() string`

GetReferenceId returns the ReferenceId field if non-nil, zero value otherwise.

### GetReferenceIdOk

`func (o *ClassifyDocumentRequest) GetReferenceIdOk() (*string, bool)`

GetReferenceIdOk returns a tuple with the ReferenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceId

`func (o *ClassifyDocumentRequest) SetReferenceId(v string)`

SetReferenceId sets ReferenceId field to given value.


### GetDocumentType

`func (o *ClassifyDocumentRequest) GetDocumentType() string`

GetDocumentType returns the DocumentType field if non-nil, zero value otherwise.

### GetDocumentTypeOk

`func (o *ClassifyDocumentRequest) GetDocumentTypeOk() (*string, bool)`

GetDocumentTypeOk returns a tuple with the DocumentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentType

`func (o *ClassifyDocumentRequest) SetDocumentType(v string)`

SetDocumentType sets DocumentType field to given value.


### GetConsent

`func (o *ClassifyDocumentRequest) GetConsent() bool`

GetConsent returns the Consent field if non-nil, zero value otherwise.

### GetConsentOk

`func (o *ClassifyDocumentRequest) GetConsentOk() (*bool, bool)`

GetConsentOk returns a tuple with the Consent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsent

`func (o *ClassifyDocumentRequest) SetConsent(v bool)`

SetConsent sets Consent field to given value.


### GetConsentPurpose

`func (o *ClassifyDocumentRequest) GetConsentPurpose() string`

GetConsentPurpose returns the ConsentPurpose field if non-nil, zero value otherwise.

### GetConsentPurposeOk

`func (o *ClassifyDocumentRequest) GetConsentPurposeOk() (*string, bool)`

GetConsentPurposeOk returns a tuple with the ConsentPurpose field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsentPurpose

`func (o *ClassifyDocumentRequest) SetConsentPurpose(v string)`

SetConsentPurpose sets ConsentPurpose field to given value.


### GetDocument

`func (o *ClassifyDocumentRequest) GetDocument() *os.File`

GetDocument returns the Document field if non-nil, zero value otherwise.

### GetDocumentOk

`func (o *ClassifyDocumentRequest) GetDocumentOk() (**os.File, bool)`

GetDocumentOk returns a tuple with the Document field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocument

`func (o *ClassifyDocumentRequest) SetDocument(v *os.File)`

SetDocument sets Document field to given value.

### HasDocument

`func (o *ClassifyDocumentRequest) HasDocument() bool`

HasDocument returns a boolean if a field has been set.

### GetDocumentUrl

`func (o *ClassifyDocumentRequest) GetDocumentUrl() string`

GetDocumentUrl returns the DocumentUrl field if non-nil, zero value otherwise.

### GetDocumentUrlOk

`func (o *ClassifyDocumentRequest) GetDocumentUrlOk() (*string, bool)`

GetDocumentUrlOk returns a tuple with the DocumentUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentUrl

`func (o *ClassifyDocumentRequest) SetDocumentUrl(v string)`

SetDocumentUrl sets DocumentUrl field to given value.

### HasDocumentUrl

`func (o *ClassifyDocumentRequest) HasDocumentUrl() bool`

HasDocumentUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


