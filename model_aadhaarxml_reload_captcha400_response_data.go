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

// AadhaarxmlReloadCaptcha400ResponseData struct for AadhaarxmlReloadCaptcha400ResponseData
type AadhaarxmlReloadCaptcha400ResponseData struct {
	CaptchaImage *string `json:"captchaImage,omitempty"`
}

// NewAadhaarxmlReloadCaptcha400ResponseData instantiates a new AadhaarxmlReloadCaptcha400ResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAadhaarxmlReloadCaptcha400ResponseData() *AadhaarxmlReloadCaptcha400ResponseData {
	this := AadhaarxmlReloadCaptcha400ResponseData{}
	return &this
}

// NewAadhaarxmlReloadCaptcha400ResponseDataWithDefaults instantiates a new AadhaarxmlReloadCaptcha400ResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAadhaarxmlReloadCaptcha400ResponseDataWithDefaults() *AadhaarxmlReloadCaptcha400ResponseData {
	this := AadhaarxmlReloadCaptcha400ResponseData{}
	return &this
}

// GetCaptchaImage returns the CaptchaImage field value if set, zero value otherwise.
func (o *AadhaarxmlReloadCaptcha400ResponseData) GetCaptchaImage() string {
	if o == nil || isNil(o.CaptchaImage) {
		var ret string
		return ret
	}
	return *o.CaptchaImage
}

// GetCaptchaImageOk returns a tuple with the CaptchaImage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AadhaarxmlReloadCaptcha400ResponseData) GetCaptchaImageOk() (*string, bool) {
	if o == nil || isNil(o.CaptchaImage) {
    return nil, false
	}
	return o.CaptchaImage, true
}

// HasCaptchaImage returns a boolean if a field has been set.
func (o *AadhaarxmlReloadCaptcha400ResponseData) HasCaptchaImage() bool {
	if o != nil && !isNil(o.CaptchaImage) {
		return true
	}

	return false
}

// SetCaptchaImage gets a reference to the given string and assigns it to the CaptchaImage field.
func (o *AadhaarxmlReloadCaptcha400ResponseData) SetCaptchaImage(v string) {
	o.CaptchaImage = &v
}

func (o AadhaarxmlReloadCaptcha400ResponseData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.CaptchaImage) {
		toSerialize["captchaImage"] = o.CaptchaImage
	}
	return json.Marshal(toSerialize)
}

type NullableAadhaarxmlReloadCaptcha400ResponseData struct {
	value *AadhaarxmlReloadCaptcha400ResponseData
	isSet bool
}

func (v NullableAadhaarxmlReloadCaptcha400ResponseData) Get() *AadhaarxmlReloadCaptcha400ResponseData {
	return v.value
}

func (v *NullableAadhaarxmlReloadCaptcha400ResponseData) Set(val *AadhaarxmlReloadCaptcha400ResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableAadhaarxmlReloadCaptcha400ResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableAadhaarxmlReloadCaptcha400ResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAadhaarxmlReloadCaptcha400ResponseData(val *AadhaarxmlReloadCaptcha400ResponseData) *NullableAadhaarxmlReloadCaptcha400ResponseData {
	return &NullableAadhaarxmlReloadCaptcha400ResponseData{value: val, isSet: true}
}

func (v NullableAadhaarxmlReloadCaptcha400ResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAadhaarxmlReloadCaptcha400ResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


