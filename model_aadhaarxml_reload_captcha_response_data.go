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

// AadhaarxmlReloadCaptchaResponseData struct for AadhaarxmlReloadCaptchaResponseData
type AadhaarxmlReloadCaptchaResponseData struct {
	CaptchaImage *string `json:"captchaImage,omitempty"`
}

// NewAadhaarxmlReloadCaptchaResponseData instantiates a new AadhaarxmlReloadCaptchaResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAadhaarxmlReloadCaptchaResponseData() *AadhaarxmlReloadCaptchaResponseData {
	this := AadhaarxmlReloadCaptchaResponseData{}
	return &this
}

// NewAadhaarxmlReloadCaptchaResponseDataWithDefaults instantiates a new AadhaarxmlReloadCaptchaResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAadhaarxmlReloadCaptchaResponseDataWithDefaults() *AadhaarxmlReloadCaptchaResponseData {
	this := AadhaarxmlReloadCaptchaResponseData{}
	return &this
}

// GetCaptchaImage returns the CaptchaImage field value if set, zero value otherwise.
func (o *AadhaarxmlReloadCaptchaResponseData) GetCaptchaImage() string {
	if o == nil || isNil(o.CaptchaImage) {
		var ret string
		return ret
	}
	return *o.CaptchaImage
}

// GetCaptchaImageOk returns a tuple with the CaptchaImage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AadhaarxmlReloadCaptchaResponseData) GetCaptchaImageOk() (*string, bool) {
	if o == nil || isNil(o.CaptchaImage) {
    return nil, false
	}
	return o.CaptchaImage, true
}

// HasCaptchaImage returns a boolean if a field has been set.
func (o *AadhaarxmlReloadCaptchaResponseData) HasCaptchaImage() bool {
	if o != nil && !isNil(o.CaptchaImage) {
		return true
	}

	return false
}

// SetCaptchaImage gets a reference to the given string and assigns it to the CaptchaImage field.
func (o *AadhaarxmlReloadCaptchaResponseData) SetCaptchaImage(v string) {
	o.CaptchaImage = &v
}

func (o AadhaarxmlReloadCaptchaResponseData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.CaptchaImage) {
		toSerialize["captchaImage"] = o.CaptchaImage
	}
	return json.Marshal(toSerialize)
}

type NullableAadhaarxmlReloadCaptchaResponseData struct {
	value *AadhaarxmlReloadCaptchaResponseData
	isSet bool
}

func (v NullableAadhaarxmlReloadCaptchaResponseData) Get() *AadhaarxmlReloadCaptchaResponseData {
	return v.value
}

func (v *NullableAadhaarxmlReloadCaptchaResponseData) Set(val *AadhaarxmlReloadCaptchaResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableAadhaarxmlReloadCaptchaResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableAadhaarxmlReloadCaptchaResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAadhaarxmlReloadCaptchaResponseData(val *AadhaarxmlReloadCaptchaResponseData) *NullableAadhaarxmlReloadCaptchaResponseData {
	return &NullableAadhaarxmlReloadCaptchaResponseData{value: val, isSet: true}
}

func (v NullableAadhaarxmlReloadCaptchaResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAadhaarxmlReloadCaptchaResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


