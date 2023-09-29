/*
decentro-in-kyc

KYC & Onboarding

API version: 1.0.0
Contact: admin@decentro.tech
*/


package decentroinkyc

import (
	"encoding/json"
)

// AadhaarxmlInitiateSession400ResponseData struct for AadhaarxmlInitiateSession400ResponseData
type AadhaarxmlInitiateSession400ResponseData struct {
	CaptchaImage *string `json:"captchaImage,omitempty"`
}

// NewAadhaarxmlInitiateSession400ResponseData instantiates a new AadhaarxmlInitiateSession400ResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAadhaarxmlInitiateSession400ResponseData() *AadhaarxmlInitiateSession400ResponseData {
	this := AadhaarxmlInitiateSession400ResponseData{}
	return &this
}

// NewAadhaarxmlInitiateSession400ResponseDataWithDefaults instantiates a new AadhaarxmlInitiateSession400ResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAadhaarxmlInitiateSession400ResponseDataWithDefaults() *AadhaarxmlInitiateSession400ResponseData {
	this := AadhaarxmlInitiateSession400ResponseData{}
	return &this
}

// GetCaptchaImage returns the CaptchaImage field value if set, zero value otherwise.
func (o *AadhaarxmlInitiateSession400ResponseData) GetCaptchaImage() string {
	if o == nil || isNil(o.CaptchaImage) {
		var ret string
		return ret
	}
	return *o.CaptchaImage
}

// GetCaptchaImageOk returns a tuple with the CaptchaImage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AadhaarxmlInitiateSession400ResponseData) GetCaptchaImageOk() (*string, bool) {
	if o == nil || isNil(o.CaptchaImage) {
    return nil, false
	}
	return o.CaptchaImage, true
}

// HasCaptchaImage returns a boolean if a field has been set.
func (o *AadhaarxmlInitiateSession400ResponseData) HasCaptchaImage() bool {
	if o != nil && !isNil(o.CaptchaImage) {
		return true
	}

	return false
}

// SetCaptchaImage gets a reference to the given string and assigns it to the CaptchaImage field.
func (o *AadhaarxmlInitiateSession400ResponseData) SetCaptchaImage(v string) {
	o.CaptchaImage = &v
}

func (o AadhaarxmlInitiateSession400ResponseData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.CaptchaImage) {
		toSerialize["captchaImage"] = o.CaptchaImage
	}
	return json.Marshal(toSerialize)
}

type NullableAadhaarxmlInitiateSession400ResponseData struct {
	value *AadhaarxmlInitiateSession400ResponseData
	isSet bool
}

func (v NullableAadhaarxmlInitiateSession400ResponseData) Get() *AadhaarxmlInitiateSession400ResponseData {
	return v.value
}

func (v *NullableAadhaarxmlInitiateSession400ResponseData) Set(val *AadhaarxmlInitiateSession400ResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableAadhaarxmlInitiateSession400ResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableAadhaarxmlInitiateSession400ResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAadhaarxmlInitiateSession400ResponseData(val *AadhaarxmlInitiateSession400ResponseData) *NullableAadhaarxmlInitiateSession400ResponseData {
	return &NullableAadhaarxmlInitiateSession400ResponseData{value: val, isSet: true}
}

func (v NullableAadhaarxmlInitiateSession400ResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAadhaarxmlInitiateSession400ResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


