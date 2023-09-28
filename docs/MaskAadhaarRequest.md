# MaskAadhaarRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReferenceId** | **string** |  | 
**Consent** | **bool** |  | 
**ConsentPurpose** | **string** |  | 
**Image** | Pointer to ***os.File** |  | [optional] 
**ImageUrl** | Pointer to **string** |  | [optional] 

## Methods

### NewMaskAadhaarRequest

`func NewMaskAadhaarRequest(referenceId string, consent bool, consentPurpose string, ) *MaskAadhaarRequest`

NewMaskAadhaarRequest instantiates a new MaskAadhaarRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMaskAadhaarRequestWithDefaults

`func NewMaskAadhaarRequestWithDefaults() *MaskAadhaarRequest`

NewMaskAadhaarRequestWithDefaults instantiates a new MaskAadhaarRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReferenceId

`func (o *MaskAadhaarRequest) GetReferenceId() string`

GetReferenceId returns the ReferenceId field if non-nil, zero value otherwise.

### GetReferenceIdOk

`func (o *MaskAadhaarRequest) GetReferenceIdOk() (*string, bool)`

GetReferenceIdOk returns a tuple with the ReferenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceId

`func (o *MaskAadhaarRequest) SetReferenceId(v string)`

SetReferenceId sets ReferenceId field to given value.


### GetConsent

`func (o *MaskAadhaarRequest) GetConsent() bool`

GetConsent returns the Consent field if non-nil, zero value otherwise.

### GetConsentOk

`func (o *MaskAadhaarRequest) GetConsentOk() (*bool, bool)`

GetConsentOk returns a tuple with the Consent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsent

`func (o *MaskAadhaarRequest) SetConsent(v bool)`

SetConsent sets Consent field to given value.


### GetConsentPurpose

`func (o *MaskAadhaarRequest) GetConsentPurpose() string`

GetConsentPurpose returns the ConsentPurpose field if non-nil, zero value otherwise.

### GetConsentPurposeOk

`func (o *MaskAadhaarRequest) GetConsentPurposeOk() (*string, bool)`

GetConsentPurposeOk returns a tuple with the ConsentPurpose field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsentPurpose

`func (o *MaskAadhaarRequest) SetConsentPurpose(v string)`

SetConsentPurpose sets ConsentPurpose field to given value.


### GetImage

`func (o *MaskAadhaarRequest) GetImage() *os.File`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *MaskAadhaarRequest) GetImageOk() (**os.File, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *MaskAadhaarRequest) SetImage(v *os.File)`

SetImage sets Image field to given value.

### HasImage

`func (o *MaskAadhaarRequest) HasImage() bool`

HasImage returns a boolean if a field has been set.

### GetImageUrl

`func (o *MaskAadhaarRequest) GetImageUrl() string`

GetImageUrl returns the ImageUrl field if non-nil, zero value otherwise.

### GetImageUrlOk

`func (o *MaskAadhaarRequest) GetImageUrlOk() (*string, bool)`

GetImageUrlOk returns a tuple with the ImageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageUrl

`func (o *MaskAadhaarRequest) SetImageUrl(v string)`

SetImageUrl sets ImageUrl field to given value.

### HasImageUrl

`func (o *MaskAadhaarRequest) HasImageUrl() bool`

HasImageUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


