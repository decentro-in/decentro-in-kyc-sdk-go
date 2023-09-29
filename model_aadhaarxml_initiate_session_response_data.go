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

// AadhaarxmlInitiateSessionResponseData struct for AadhaarxmlInitiateSessionResponseData
type AadhaarxmlInitiateSessionResponseData struct {
	CaptchaImage *string `json:"captchaImage,omitempty"`
}

// NewAadhaarxmlInitiateSessionResponseData instantiates a new AadhaarxmlInitiateSessionResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAadhaarxmlInitiateSessionResponseData() *AadhaarxmlInitiateSessionResponseData {
	this := AadhaarxmlInitiateSessionResponseData{}
	return &this
}

// NewAadhaarxmlInitiateSessionResponseDataWithDefaults instantiates a new AadhaarxmlInitiateSessionResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAadhaarxmlInitiateSessionResponseDataWithDefaults() *AadhaarxmlInitiateSessionResponseData {
	this := AadhaarxmlInitiateSessionResponseData{}
	return &this
}

// GetCaptchaImage returns the CaptchaImage field value if set, zero value otherwise.
func (o *AadhaarxmlInitiateSessionResponseData) GetCaptchaImage() string {
	if o == nil || isNil(o.CaptchaImage) {
		var ret string
		return ret
	}
	return *o.CaptchaImage
}

// GetCaptchaImageOk returns a tuple with the CaptchaImage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AadhaarxmlInitiateSessionResponseData) GetCaptchaImageOk() (*string, bool) {
	if o == nil || isNil(o.CaptchaImage) {
    return nil, false
	}
	return o.CaptchaImage, true
}

// HasCaptchaImage returns a boolean if a field has been set.
func (o *AadhaarxmlInitiateSessionResponseData) HasCaptchaImage() bool {
	if o != nil && !isNil(o.CaptchaImage) {
		return true
	}

	return false
}

// SetCaptchaImage gets a reference to the given string and assigns it to the CaptchaImage field.
func (o *AadhaarxmlInitiateSessionResponseData) SetCaptchaImage(v string) {
	o.CaptchaImage = &v
}

func (o AadhaarxmlInitiateSessionResponseData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.CaptchaImage) {
		toSerialize["captchaImage"] = o.CaptchaImage
	}
	return json.Marshal(toSerialize)
}

type NullableAadhaarxmlInitiateSessionResponseData struct {
	value *AadhaarxmlInitiateSessionResponseData
	isSet bool
}

func (v NullableAadhaarxmlInitiateSessionResponseData) Get() *AadhaarxmlInitiateSessionResponseData {
	return v.value
}

func (v *NullableAadhaarxmlInitiateSessionResponseData) Set(val *AadhaarxmlInitiateSessionResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableAadhaarxmlInitiateSessionResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableAadhaarxmlInitiateSessionResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAadhaarxmlInitiateSessionResponseData(val *AadhaarxmlInitiateSessionResponseData) *NullableAadhaarxmlInitiateSessionResponseData {
	return &NullableAadhaarxmlInitiateSessionResponseData{value: val, isSet: true}
}

func (v NullableAadhaarxmlInitiateSessionResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAadhaarxmlInitiateSessionResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


