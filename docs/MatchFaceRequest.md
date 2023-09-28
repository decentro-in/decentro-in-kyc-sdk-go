# MatchFaceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReferenceId** | **string** |  | 
**Consent** | **string** |  | 
**ConsentPurpose** | **string** |  | 
**Image1** | Pointer to ***os.File** |  | [optional] 
**Image2** | Pointer to ***os.File** |  | [optional] 
**Threshold** | Pointer to **int32** |  | [optional] 
**Image1Url** | Pointer to **string** |  | [optional] 
**Image2Url** | Pointer to **string** |  | [optional] 

## Methods

### NewMatchFaceRequest

`func NewMatchFaceRequest(referenceId string, consent string, consentPurpose string, ) *MatchFaceRequest`

NewMatchFaceRequest instantiates a new MatchFaceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMatchFaceRequestWithDefaults

`func NewMatchFaceRequestWithDefaults() *MatchFaceRequest`

NewMatchFaceRequestWithDefaults instantiates a new MatchFaceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReferenceId

`func (o *MatchFaceRequest) GetReferenceId() string`

GetReferenceId returns the ReferenceId field if non-nil, zero value otherwise.

### GetReferenceIdOk

`func (o *MatchFaceRequest) GetReferenceIdOk() (*string, bool)`

GetReferenceIdOk returns a tuple with the ReferenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceId

`func (o *MatchFaceRequest) SetReferenceId(v string)`

SetReferenceId sets ReferenceId field to given value.


### GetConsent

`func (o *MatchFaceRequest) GetConsent() string`

GetConsent returns the Consent field if non-nil, zero value otherwise.

### GetConsentOk

`func (o *MatchFaceRequest) GetConsentOk() (*string, bool)`

GetConsentOk returns a tuple with the Consent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsent

`func (o *MatchFaceRequest) SetConsent(v string)`

SetConsent sets Consent field to given value.


### GetConsentPurpose

`func (o *MatchFaceRequest) GetConsentPurpose() string`

GetConsentPurpose returns the ConsentPurpose field if non-nil, zero value otherwise.

### GetConsentPurposeOk

`func (o *MatchFaceRequest) GetConsentPurposeOk() (*string, bool)`

GetConsentPurposeOk returns a tuple with the ConsentPurpose field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsentPurpose

`func (o *MatchFaceRequest) SetConsentPurpose(v string)`

SetConsentPurpose sets ConsentPurpose field to given value.


### GetImage1

`func (o *MatchFaceRequest) GetImage1() *os.File`

GetImage1 returns the Image1 field if non-nil, zero value otherwise.

### GetImage1Ok

`func (o *MatchFaceRequest) GetImage1Ok() (**os.File, bool)`

GetImage1Ok returns a tuple with the Image1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage1

`func (o *MatchFaceRequest) SetImage1(v *os.File)`

SetImage1 sets Image1 field to given value.

### HasImage1

`func (o *MatchFaceRequest) HasImage1() bool`

HasImage1 returns a boolean if a field has been set.

### GetImage2

`func (o *MatchFaceRequest) GetImage2() *os.File`

GetImage2 returns the Image2 field if non-nil, zero value otherwise.

### GetImage2Ok

`func (o *MatchFaceRequest) GetImage2Ok() (**os.File, bool)`

GetImage2Ok returns a tuple with the Image2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage2

`func (o *MatchFaceRequest) SetImage2(v *os.File)`

SetImage2 sets Image2 field to given value.

### HasImage2

`func (o *MatchFaceRequest) HasImage2() bool`

HasImage2 returns a boolean if a field has been set.

### GetThreshold

`func (o *MatchFaceRequest) GetThreshold() int32`

GetThreshold returns the Threshold field if non-nil, zero value otherwise.

### GetThresholdOk

`func (o *MatchFaceRequest) GetThresholdOk() (*int32, bool)`

GetThresholdOk returns a tuple with the Threshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreshold

`func (o *MatchFaceRequest) SetThreshold(v int32)`

SetThreshold sets Threshold field to given value.

### HasThreshold

`func (o *MatchFaceRequest) HasThreshold() bool`

HasThreshold returns a boolean if a field has been set.

### GetImage1Url

`func (o *MatchFaceRequest) GetImage1Url() string`

GetImage1Url returns the Image1Url field if non-nil, zero value otherwise.

### GetImage1UrlOk

`func (o *MatchFaceRequest) GetImage1UrlOk() (*string, bool)`

GetImage1UrlOk returns a tuple with the Image1Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage1Url

`func (o *MatchFaceRequest) SetImage1Url(v string)`

SetImage1Url sets Image1Url field to given value.

### HasImage1Url

`func (o *MatchFaceRequest) HasImage1Url() bool`

HasImage1Url returns a boolean if a field has been set.

### GetImage2Url

`func (o *MatchFaceRequest) GetImage2Url() string`

GetImage2Url returns the Image2Url field if non-nil, zero value otherwise.

### GetImage2UrlOk

`func (o *MatchFaceRequest) GetImage2UrlOk() (*string, bool)`

GetImage2UrlOk returns a tuple with the Image2Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage2Url

`func (o *MatchFaceRequest) SetImage2Url(v string)`

SetImage2Url sets Image2Url field to given value.

### HasImage2Url

`func (o *MatchFaceRequest) HasImage2Url() bool`

HasImage2Url returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


