/*
decentro-in-kyc

KYC & Onboarding

API version: 1.0.0
Contact: admin@decentro.tech
*/


package decentroinkyc

import (
	"encoding/json"
	"os"
)

// CheckPhotocopyRequest struct for CheckPhotocopyRequest
type CheckPhotocopyRequest struct {
	// 
	ReferenceId string `json:"reference_id"`
	// 
	Consent bool `json:"consent"`
	// 
	ConsentPurpose string `json:"consent_purpose"`
	Image **os.File `json:"image,omitempty"`
	ImageUrl *string `json:"image_url,omitempty"`
}

// NewCheckPhotocopyRequest instantiates a new CheckPhotocopyRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCheckPhotocopyRequest(referenceId string, consent bool, consentPurpose string) *CheckPhotocopyRequest {
	this := CheckPhotocopyRequest{}
	this.ReferenceId = referenceId
	this.Consent = consent
	this.ConsentPurpose = consentPurpose
	return &this
}

// NewCheckPhotocopyRequestWithDefaults instantiates a new CheckPhotocopyRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCheckPhotocopyRequestWithDefaults() *CheckPhotocopyRequest {
	this := CheckPhotocopyRequest{}
	return &this
}

// GetReferenceId returns the ReferenceId field value
func (o *CheckPhotocopyRequest) GetReferenceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ReferenceId
}

// GetReferenceIdOk returns a tuple with the ReferenceId field value
// and a boolean to check if the value has been set.
func (o *CheckPhotocopyRequest) GetReferenceIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.ReferenceId, true
}

// SetReferenceId sets field value
func (o *CheckPhotocopyRequest) SetReferenceId(v string) {
	o.ReferenceId = v
}

// GetConsent returns the Consent field value
func (o *CheckPhotocopyRequest) GetConsent() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Consent
}

// GetConsentOk returns a tuple with the Consent field value
// and a boolean to check if the value has been set.
func (o *CheckPhotocopyRequest) GetConsentOk() (*bool, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Consent, true
}

// SetConsent sets field value
func (o *CheckPhotocopyRequest) SetConsent(v bool) {
	o.Consent = v
}

// GetConsentPurpose returns the ConsentPurpose field value
func (o *CheckPhotocopyRequest) GetConsentPurpose() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ConsentPurpose
}

// GetConsentPurposeOk returns a tuple with the ConsentPurpose field value
// and a boolean to check if the value has been set.
func (o *CheckPhotocopyRequest) GetConsentPurposeOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.ConsentPurpose, true
}

// SetConsentPurpose sets field value
func (o *CheckPhotocopyRequest) SetConsentPurpose(v string) {
	o.ConsentPurpose = v
}

// GetImage returns the Image field value if set, zero value otherwise.
func (o *CheckPhotocopyRequest) GetImage() *os.File {
	if o == nil || isNil(o.Image) {
		var ret *os.File
		return ret
	}
	return *o.Image
}

// GetImageOk returns a tuple with the Image field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckPhotocopyRequest) GetImageOk() (**os.File, bool) {
	if o == nil || isNil(o.Image) {
    return nil, false
	}
	return o.Image, true
}

// HasImage returns a boolean if a field has been set.
func (o *CheckPhotocopyRequest) HasImage() bool {
	if o != nil && !isNil(o.Image) {
		return true
	}

	return false
}

// SetImage gets a reference to the given *os.File and assigns it to the Image field.
func (o *CheckPhotocopyRequest) SetImage(v *os.File) {
	o.Image = &v
}

// GetImageUrl returns the ImageUrl field value if set, zero value otherwise.
func (o *CheckPhotocopyRequest) GetImageUrl() string {
	if o == nil || isNil(o.ImageUrl) {
		var ret string
		return ret
	}
	return *o.ImageUrl
}

// GetImageUrlOk returns a tuple with the ImageUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckPhotocopyRequest) GetImageUrlOk() (*string, bool) {
	if o == nil || isNil(o.ImageUrl) {
    return nil, false
	}
	return o.ImageUrl, true
}

// HasImageUrl returns a boolean if a field has been set.
func (o *CheckPhotocopyRequest) HasImageUrl() bool {
	if o != nil && !isNil(o.ImageUrl) {
		return true
	}

	return false
}

// SetImageUrl gets a reference to the given string and assigns it to the ImageUrl field.
func (o *CheckPhotocopyRequest) SetImageUrl(v string) {
	o.ImageUrl = &v
}

func (o CheckPhotocopyRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["reference_id"] = o.ReferenceId
	}
	if true {
		toSerialize["consent"] = o.Consent
	}
	if true {
		toSerialize["consent_purpose"] = o.ConsentPurpose
	}
	if !isNil(o.Image) {
		toSerialize["image"] = o.Image
	}
	if !isNil(o.ImageUrl) {
		toSerialize["image_url"] = o.ImageUrl
	}
	return json.Marshal(toSerialize)
}

type NullableCheckPhotocopyRequest struct {
	value *CheckPhotocopyRequest
	isSet bool
}

func (v NullableCheckPhotocopyRequest) Get() *CheckPhotocopyRequest {
	return v.value
}

func (v *NullableCheckPhotocopyRequest) Set(val *CheckPhotocopyRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCheckPhotocopyRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCheckPhotocopyRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCheckPhotocopyRequest(val *CheckPhotocopyRequest) *NullableCheckPhotocopyRequest {
	return &NullableCheckPhotocopyRequest{value: val, isSet: true}
}

func (v NullableCheckPhotocopyRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCheckPhotocopyRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


