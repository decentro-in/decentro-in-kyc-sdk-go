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

// CheckImageQualityResponseDataImageQualityQualityScoresSharpness struct for CheckImageQualityResponseDataImageQualityQualityScoresSharpness
type CheckImageQualityResponseDataImageQualityQualityScoresSharpness struct {
	Valid *string `json:"valid,omitempty"`
	Score *float32 `json:"score,omitempty"`
}

// NewCheckImageQualityResponseDataImageQualityQualityScoresSharpness instantiates a new CheckImageQualityResponseDataImageQualityQualityScoresSharpness object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCheckImageQualityResponseDataImageQualityQualityScoresSharpness() *CheckImageQualityResponseDataImageQualityQualityScoresSharpness {
	this := CheckImageQualityResponseDataImageQualityQualityScoresSharpness{}
	return &this
}

// NewCheckImageQualityResponseDataImageQualityQualityScoresSharpnessWithDefaults instantiates a new CheckImageQualityResponseDataImageQualityQualityScoresSharpness object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCheckImageQualityResponseDataImageQualityQualityScoresSharpnessWithDefaults() *CheckImageQualityResponseDataImageQualityQualityScoresSharpness {
	this := CheckImageQualityResponseDataImageQualityQualityScoresSharpness{}
	return &this
}

// GetValid returns the Valid field value if set, zero value otherwise.
func (o *CheckImageQualityResponseDataImageQualityQualityScoresSharpness) GetValid() string {
	if o == nil || isNil(o.Valid) {
		var ret string
		return ret
	}
	return *o.Valid
}

// GetValidOk returns a tuple with the Valid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckImageQualityResponseDataImageQualityQualityScoresSharpness) GetValidOk() (*string, bool) {
	if o == nil || isNil(o.Valid) {
    return nil, false
	}
	return o.Valid, true
}

// HasValid returns a boolean if a field has been set.
func (o *CheckImageQualityResponseDataImageQualityQualityScoresSharpness) HasValid() bool {
	if o != nil && !isNil(o.Valid) {
		return true
	}

	return false
}

// SetValid gets a reference to the given string and assigns it to the Valid field.
func (o *CheckImageQualityResponseDataImageQualityQualityScoresSharpness) SetValid(v string) {
	o.Valid = &v
}

// GetScore returns the Score field value if set, zero value otherwise.
func (o *CheckImageQualityResponseDataImageQualityQualityScoresSharpness) GetScore() float32 {
	if o == nil || isNil(o.Score) {
		var ret float32
		return ret
	}
	return *o.Score
}

// GetScoreOk returns a tuple with the Score field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckImageQualityResponseDataImageQualityQualityScoresSharpness) GetScoreOk() (*float32, bool) {
	if o == nil || isNil(o.Score) {
    return nil, false
	}
	return o.Score, true
}

// HasScore returns a boolean if a field has been set.
func (o *CheckImageQualityResponseDataImageQualityQualityScoresSharpness) HasScore() bool {
	if o != nil && !isNil(o.Score) {
		return true
	}

	return false
}

// SetScore gets a reference to the given float32 and assigns it to the Score field.
func (o *CheckImageQualityResponseDataImageQualityQualityScoresSharpness) SetScore(v float32) {
	o.Score = &v
}

func (o CheckImageQualityResponseDataImageQualityQualityScoresSharpness) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Valid) {
		toSerialize["valid"] = o.Valid
	}
	if !isNil(o.Score) {
		toSerialize["score"] = o.Score
	}
	return json.Marshal(toSerialize)
}

type NullableCheckImageQualityResponseDataImageQualityQualityScoresSharpness struct {
	value *CheckImageQualityResponseDataImageQualityQualityScoresSharpness
	isSet bool
}

func (v NullableCheckImageQualityResponseDataImageQualityQualityScoresSharpness) Get() *CheckImageQualityResponseDataImageQualityQualityScoresSharpness {
	return v.value
}

func (v *NullableCheckImageQualityResponseDataImageQualityQualityScoresSharpness) Set(val *CheckImageQualityResponseDataImageQualityQualityScoresSharpness) {
	v.value = val
	v.isSet = true
}

func (v NullableCheckImageQualityResponseDataImageQualityQualityScoresSharpness) IsSet() bool {
	return v.isSet
}

func (v *NullableCheckImageQualityResponseDataImageQualityQualityScoresSharpness) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCheckImageQualityResponseDataImageQualityQualityScoresSharpness(val *CheckImageQualityResponseDataImageQualityQualityScoresSharpness) *NullableCheckImageQualityResponseDataImageQualityQualityScoresSharpness {
	return &NullableCheckImageQualityResponseDataImageQualityQualityScoresSharpness{value: val, isSet: true}
}

func (v NullableCheckImageQualityResponseDataImageQualityQualityScoresSharpness) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCheckImageQualityResponseDataImageQualityQualityScoresSharpness) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


