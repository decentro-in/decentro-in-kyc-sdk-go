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

// CheckImageQualityResponseDataImageQualityQualityScoresTextQuality struct for CheckImageQualityResponseDataImageQualityQualityScoresTextQuality
type CheckImageQualityResponseDataImageQualityQualityScoresTextQuality struct {
	Valid *string `json:"valid,omitempty"`
	Score *float32 `json:"score,omitempty"`
}

// NewCheckImageQualityResponseDataImageQualityQualityScoresTextQuality instantiates a new CheckImageQualityResponseDataImageQualityQualityScoresTextQuality object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCheckImageQualityResponseDataImageQualityQualityScoresTextQuality() *CheckImageQualityResponseDataImageQualityQualityScoresTextQuality {
	this := CheckImageQualityResponseDataImageQualityQualityScoresTextQuality{}
	return &this
}

// NewCheckImageQualityResponseDataImageQualityQualityScoresTextQualityWithDefaults instantiates a new CheckImageQualityResponseDataImageQualityQualityScoresTextQuality object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCheckImageQualityResponseDataImageQualityQualityScoresTextQualityWithDefaults() *CheckImageQualityResponseDataImageQualityQualityScoresTextQuality {
	this := CheckImageQualityResponseDataImageQualityQualityScoresTextQuality{}
	return &this
}

// GetValid returns the Valid field value if set, zero value otherwise.
func (o *CheckImageQualityResponseDataImageQualityQualityScoresTextQuality) GetValid() string {
	if o == nil || isNil(o.Valid) {
		var ret string
		return ret
	}
	return *o.Valid
}

// GetValidOk returns a tuple with the Valid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckImageQualityResponseDataImageQualityQualityScoresTextQuality) GetValidOk() (*string, bool) {
	if o == nil || isNil(o.Valid) {
    return nil, false
	}
	return o.Valid, true
}

// HasValid returns a boolean if a field has been set.
func (o *CheckImageQualityResponseDataImageQualityQualityScoresTextQuality) HasValid() bool {
	if o != nil && !isNil(o.Valid) {
		return true
	}

	return false
}

// SetValid gets a reference to the given string and assigns it to the Valid field.
func (o *CheckImageQualityResponseDataImageQualityQualityScoresTextQuality) SetValid(v string) {
	o.Valid = &v
}

// GetScore returns the Score field value if set, zero value otherwise.
func (o *CheckImageQualityResponseDataImageQualityQualityScoresTextQuality) GetScore() float32 {
	if o == nil || isNil(o.Score) {
		var ret float32
		return ret
	}
	return *o.Score
}

// GetScoreOk returns a tuple with the Score field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckImageQualityResponseDataImageQualityQualityScoresTextQuality) GetScoreOk() (*float32, bool) {
	if o == nil || isNil(o.Score) {
    return nil, false
	}
	return o.Score, true
}

// HasScore returns a boolean if a field has been set.
func (o *CheckImageQualityResponseDataImageQualityQualityScoresTextQuality) HasScore() bool {
	if o != nil && !isNil(o.Score) {
		return true
	}

	return false
}

// SetScore gets a reference to the given float32 and assigns it to the Score field.
func (o *CheckImageQualityResponseDataImageQualityQualityScoresTextQuality) SetScore(v float32) {
	o.Score = &v
}

func (o CheckImageQualityResponseDataImageQualityQualityScoresTextQuality) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Valid) {
		toSerialize["valid"] = o.Valid
	}
	if !isNil(o.Score) {
		toSerialize["score"] = o.Score
	}
	return json.Marshal(toSerialize)
}

type NullableCheckImageQualityResponseDataImageQualityQualityScoresTextQuality struct {
	value *CheckImageQualityResponseDataImageQualityQualityScoresTextQuality
	isSet bool
}

func (v NullableCheckImageQualityResponseDataImageQualityQualityScoresTextQuality) Get() *CheckImageQualityResponseDataImageQualityQualityScoresTextQuality {
	return v.value
}

func (v *NullableCheckImageQualityResponseDataImageQualityQualityScoresTextQuality) Set(val *CheckImageQualityResponseDataImageQualityQualityScoresTextQuality) {
	v.value = val
	v.isSet = true
}

func (v NullableCheckImageQualityResponseDataImageQualityQualityScoresTextQuality) IsSet() bool {
	return v.isSet
}

func (v *NullableCheckImageQualityResponseDataImageQualityQualityScoresTextQuality) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCheckImageQualityResponseDataImageQualityQualityScoresTextQuality(val *CheckImageQualityResponseDataImageQualityQualityScoresTextQuality) *NullableCheckImageQualityResponseDataImageQualityQualityScoresTextQuality {
	return &NullableCheckImageQualityResponseDataImageQualityQualityScoresTextQuality{value: val, isSet: true}
}

func (v NullableCheckImageQualityResponseDataImageQualityQualityScoresTextQuality) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCheckImageQualityResponseDataImageQualityQualityScoresTextQuality) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


