# CheckImageQualityRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReferenceId** | **string** |  | 
**Consent** | **bool** |  | 
**ConsentPurpose** | **string** |  | 
**Image** | Pointer to ***os.File** |  | [optional] 
**QualityParameter** | Pointer to **string** |  | [optional] 
**ImageUrl** | Pointer to **string** |  | [optional] 

## Methods

### NewCheckImageQualityRequest

`func NewCheckImageQualityRequest(referenceId string, consent bool, consentPurpose string, ) *CheckImageQualityRequest`

NewCheckImageQualityRequest instantiates a new CheckImageQualityRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCheckImageQualityRequestWithDefaults

`func NewCheckImageQualityRequestWithDefaults() *CheckImageQualityRequest`

NewCheckImageQualityRequestWithDefaults instantiates a new CheckImageQualityRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReferenceId

`func (o *CheckImageQualityRequest) GetReferenceId() string`

GetReferenceId returns the ReferenceId field if non-nil, zero value otherwise.

### GetReferenceIdOk

`func (o *CheckImageQualityRequest) GetReferenceIdOk() (*string, bool)`

GetReferenceIdOk returns a tuple with the ReferenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceId

`func (o *CheckImageQualityRequest) SetReferenceId(v string)`

SetReferenceId sets ReferenceId field to given value.


### GetConsent

`func (o *CheckImageQualityRequest) GetConsent() bool`

GetConsent returns the Consent field if non-nil, zero value otherwise.

### GetConsentOk

`func (o *CheckImageQualityRequest) GetConsentOk() (*bool, bool)`

GetConsentOk returns a tuple with the Consent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsent

`func (o *CheckImageQualityRequest) SetConsent(v bool)`

SetConsent sets Consent field to given value.


### GetConsentPurpose

`func (o *CheckImageQualityRequest) GetConsentPurpose() string`

GetConsentPurpose returns the ConsentPurpose field if non-nil, zero value otherwise.

### GetConsentPurposeOk

`func (o *CheckImageQualityRequest) GetConsentPurposeOk() (*string, bool)`

GetConsentPurposeOk returns a tuple with the ConsentPurpose field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsentPurpose

`func (o *CheckImageQualityRequest) SetConsentPurpose(v string)`

SetConsentPurpose sets ConsentPurpose field to given value.


### GetImage

`func (o *CheckImageQualityRequest) GetImage() *os.File`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *CheckImageQualityRequest) GetImageOk() (**os.File, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *CheckImageQualityRequest) SetImage(v *os.File)`

SetImage sets Image field to given value.

### HasImage

`func (o *CheckImageQualityRequest) HasImage() bool`

HasImage returns a boolean if a field has been set.

### GetQualityParameter

`func (o *CheckImageQualityRequest) GetQualityParameter() string`

GetQualityParameter returns the QualityParameter field if non-nil, zero value otherwise.

### GetQualityParameterOk

`func (o *CheckImageQualityRequest) GetQualityParameterOk() (*string, bool)`

GetQualityParameterOk returns a tuple with the QualityParameter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQualityParameter

`func (o *CheckImageQualityRequest) SetQualityParameter(v string)`

SetQualityParameter sets QualityParameter field to given value.

### HasQualityParameter

`func (o *CheckImageQualityRequest) HasQualityParameter() bool`

HasQualityParameter returns a boolean if a field has been set.

### GetImageUrl

`func (o *CheckImageQualityRequest) GetImageUrl() string`

GetImageUrl returns the ImageUrl field if non-nil, zero value otherwise.

### GetImageUrlOk

`func (o *CheckImageQualityRequest) GetImageUrlOk() (*string, bool)`

GetImageUrlOk returns a tuple with the ImageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageUrl

`func (o *CheckImageQualityRequest) SetImageUrl(v string)`

SetImageUrl sets ImageUrl field to given value.

### HasImageUrl

`func (o *CheckImageQualityRequest) HasImageUrl() bool`

HasImageUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


