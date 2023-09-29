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

// CheckVideoLivenessResponseData struct for CheckVideoLivenessResponseData
type CheckVideoLivenessResponseData struct {
	Confidence *string `json:"confidence,omitempty"`
	Status *string `json:"status,omitempty"`
}

// NewCheckVideoLivenessResponseData instantiates a new CheckVideoLivenessResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCheckVideoLivenessResponseData() *CheckVideoLivenessResponseData {
	this := CheckVideoLivenessResponseData{}
	return &this
}

// NewCheckVideoLivenessResponseDataWithDefaults instantiates a new CheckVideoLivenessResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCheckVideoLivenessResponseDataWithDefaults() *CheckVideoLivenessResponseData {
	this := CheckVideoLivenessResponseData{}
	return &this
}

// GetConfidence returns the Confidence field value if set, zero value otherwise.
func (o *CheckVideoLivenessResponseData) GetConfidence() string {
	if o == nil || isNil(o.Confidence) {
		var ret string
		return ret
	}
	return *o.Confidence
}

// GetConfidenceOk returns a tuple with the Confidence field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckVideoLivenessResponseData) GetConfidenceOk() (*string, bool) {
	if o == nil || isNil(o.Confidence) {
    return nil, false
	}
	return o.Confidence, true
}

// HasConfidence returns a boolean if a field has been set.
func (o *CheckVideoLivenessResponseData) HasConfidence() bool {
	if o != nil && !isNil(o.Confidence) {
		return true
	}

	return false
}

// SetConfidence gets a reference to the given string and assigns it to the Confidence field.
func (o *CheckVideoLivenessResponseData) SetConfidence(v string) {
	o.Confidence = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *CheckVideoLivenessResponseData) GetStatus() string {
	if o == nil || isNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckVideoLivenessResponseData) GetStatusOk() (*string, bool) {
	if o == nil || isNil(o.Status) {
    return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *CheckVideoLivenessResponseData) HasStatus() bool {
	if o != nil && !isNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *CheckVideoLivenessResponseData) SetStatus(v string) {
	o.Status = &v
}

func (o CheckVideoLivenessResponseData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Confidence) {
		toSerialize["confidence"] = o.Confidence
	}
	if !isNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	return json.Marshal(toSerialize)
}

type NullableCheckVideoLivenessResponseData struct {
	value *CheckVideoLivenessResponseData
	isSet bool
}

func (v NullableCheckVideoLivenessResponseData) Get() *CheckVideoLivenessResponseData {
	return v.value
}

func (v *NullableCheckVideoLivenessResponseData) Set(val *CheckVideoLivenessResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableCheckVideoLivenessResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableCheckVideoLivenessResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCheckVideoLivenessResponseData(val *CheckVideoLivenessResponseData) *NullableCheckVideoLivenessResponseData {
	return &NullableCheckVideoLivenessResponseData{value: val, isSet: true}
}

func (v NullableCheckVideoLivenessResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCheckVideoLivenessResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


