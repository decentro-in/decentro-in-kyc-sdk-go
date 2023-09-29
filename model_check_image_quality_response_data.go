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

// CheckImageQualityResponseData struct for CheckImageQualityResponseData
type CheckImageQualityResponseData struct {
	ImageQuality *CheckImageQualityResponseDataImageQuality `json:"imageQuality,omitempty"`
}

// NewCheckImageQualityResponseData instantiates a new CheckImageQualityResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCheckImageQualityResponseData() *CheckImageQualityResponseData {
	this := CheckImageQualityResponseData{}
	return &this
}

// NewCheckImageQualityResponseDataWithDefaults instantiates a new CheckImageQualityResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCheckImageQualityResponseDataWithDefaults() *CheckImageQualityResponseData {
	this := CheckImageQualityResponseData{}
	return &this
}

// GetImageQuality returns the ImageQuality field value if set, zero value otherwise.
func (o *CheckImageQualityResponseData) GetImageQuality() CheckImageQualityResponseDataImageQuality {
	if o == nil || isNil(o.ImageQuality) {
		var ret CheckImageQualityResponseDataImageQuality
		return ret
	}
	return *o.ImageQuality
}

// GetImageQualityOk returns a tuple with the ImageQuality field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckImageQualityResponseData) GetImageQualityOk() (*CheckImageQualityResponseDataImageQuality, bool) {
	if o == nil || isNil(o.ImageQuality) {
    return nil, false
	}
	return o.ImageQuality, true
}

// HasImageQuality returns a boolean if a field has been set.
func (o *CheckImageQualityResponseData) HasImageQuality() bool {
	if o != nil && !isNil(o.ImageQuality) {
		return true
	}

	return false
}

// SetImageQuality gets a reference to the given CheckImageQualityResponseDataImageQuality and assigns it to the ImageQuality field.
func (o *CheckImageQualityResponseData) SetImageQuality(v CheckImageQualityResponseDataImageQuality) {
	o.ImageQuality = &v
}

func (o CheckImageQualityResponseData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.ImageQuality) {
		toSerialize["imageQuality"] = o.ImageQuality
	}
	return json.Marshal(toSerialize)
}

type NullableCheckImageQualityResponseData struct {
	value *CheckImageQualityResponseData
	isSet bool
}

func (v NullableCheckImageQualityResponseData) Get() *CheckImageQualityResponseData {
	return v.value
}

func (v *NullableCheckImageQualityResponseData) Set(val *CheckImageQualityResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableCheckImageQualityResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableCheckImageQualityResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCheckImageQualityResponseData(val *CheckImageQualityResponseData) *NullableCheckImageQualityResponseData {
	return &NullableCheckImageQualityResponseData{value: val, isSet: true}
}

func (v NullableCheckImageQualityResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCheckImageQualityResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


