# CheckVideoLivenessRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReferenceId** | **string** |  | 
**Consent** | **string** |  | 
**ConsentPurpose** | **string** |  | 
**Video** | Pointer to ***os.File** |  | [optional] 
**VideoUrl** | Pointer to **string** |  | [optional] 

## Methods

### NewCheckVideoLivenessRequest

`func NewCheckVideoLivenessRequest(referenceId string, consent string, consentPurpose string, ) *CheckVideoLivenessRequest`

NewCheckVideoLivenessRequest instantiates a new CheckVideoLivenessRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCheckVideoLivenessRequestWithDefaults

`func NewCheckVideoLivenessRequestWithDefaults() *CheckVideoLivenessRequest`

NewCheckVideoLivenessRequestWithDefaults instantiates a new CheckVideoLivenessRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReferenceId

`func (o *CheckVideoLivenessRequest) GetReferenceId() string`

GetReferenceId returns the ReferenceId field if non-nil, zero value otherwise.

### GetReferenceIdOk

`func (o *CheckVideoLivenessRequest) GetReferenceIdOk() (*string, bool)`

GetReferenceIdOk returns a tuple with the ReferenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceId

`func (o *CheckVideoLivenessRequest) SetReferenceId(v string)`

SetReferenceId sets ReferenceId field to given value.


### GetConsent

`func (o *CheckVideoLivenessRequest) GetConsent() string`

GetConsent returns the Consent field if non-nil, zero value otherwise.

### GetConsentOk

`func (o *CheckVideoLivenessRequest) GetConsentOk() (*string, bool)`

GetConsentOk returns a tuple with the Consent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsent

`func (o *CheckVideoLivenessRequest) SetConsent(v string)`

SetConsent sets Consent field to given value.


### GetConsentPurpose

`func (o *CheckVideoLivenessRequest) GetConsentPurpose() string`

GetConsentPurpose returns the ConsentPurpose field if non-nil, zero value otherwise.

### GetConsentPurposeOk

`func (o *CheckVideoLivenessRequest) GetConsentPurposeOk() (*string, bool)`

GetConsentPurposeOk returns a tuple with the ConsentPurpose field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsentPurpose

`func (o *CheckVideoLivenessRequest) SetConsentPurpose(v string)`

SetConsentPurpose sets ConsentPurpose field to given value.


### GetVideo

`func (o *CheckVideoLivenessRequest) GetVideo() *os.File`

GetVideo returns the Video field if non-nil, zero value otherwise.

### GetVideoOk

`func (o *CheckVideoLivenessRequest) GetVideoOk() (**os.File, bool)`

GetVideoOk returns a tuple with the Video field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideo

`func (o *CheckVideoLivenessRequest) SetVideo(v *os.File)`

SetVideo sets Video field to given value.

### HasVideo

`func (o *CheckVideoLivenessRequest) HasVideo() bool`

HasVideo returns a boolean if a field has been set.

### GetVideoUrl

`func (o *CheckVideoLivenessRequest) GetVideoUrl() string`

GetVideoUrl returns the VideoUrl field if non-nil, zero value otherwise.

### GetVideoUrlOk

`func (o *CheckVideoLivenessRequest) GetVideoUrlOk() (*string, bool)`

GetVideoUrlOk returns a tuple with the VideoUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideoUrl

`func (o *CheckVideoLivenessRequest) SetVideoUrl(v string)`

SetVideoUrl sets VideoUrl field to given value.

### HasVideoUrl

`func (o *CheckVideoLivenessRequest) HasVideoUrl() bool`

HasVideoUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


