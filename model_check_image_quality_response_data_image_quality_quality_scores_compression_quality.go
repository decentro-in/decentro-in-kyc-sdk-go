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

// CheckImageQualityResponseDataImageQualityQualityScoresCompressionQuality struct for CheckImageQualityResponseDataImageQualityQualityScoresCompressionQuality
type CheckImageQualityResponseDataImageQualityQualityScoresCompressionQuality struct {
	Valid *string `json:"valid,omitempty"`
	Score *float32 `json:"score,omitempty"`
}

// NewCheckImageQualityResponseDataImageQualityQualityScoresCompressionQuality instantiates a new CheckImageQualityResponseDataImageQualityQualityScoresCompressionQuality object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCheckImageQualityResponseDataImageQualityQualityScoresCompressionQuality() *CheckImageQualityResponseDataImageQualityQualityScoresCompressionQuality {
	this := CheckImageQualityResponseDataImageQualityQualityScoresCompressionQuality{}
	return &this
}

// NewCheckImageQualityResponseDataImageQualityQualityScoresCompressionQualityWithDefaults instantiates a new CheckImageQualityResponseDataImageQualityQualityScoresCompressionQuality object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCheckImageQualityResponseDataImageQualityQualityScoresCompressionQualityWithDefaults() *CheckImageQualityResponseDataImageQualityQualityScoresCompressionQuality {
	this := CheckImageQualityResponseDataImageQualityQualityScoresCompressionQuality{}
	return &this
}

// GetValid returns the Valid field value if set, zero value otherwise.
func (o *CheckImageQualityResponseDataImageQualityQualityScoresCompressionQuality) GetValid() string {
	if o == nil || isNil(o.Valid) {
		var ret string
		return ret
	}
	return *o.Valid
}

// GetValidOk returns a tuple with the Valid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckImageQualityResponseDataImageQualityQualityScoresCompressionQuality) GetValidOk() (*string, bool) {
	if o == nil || isNil(o.Valid) {
    return nil, false
	}
	return o.Valid, true
}

// HasValid returns a boolean if a field has been set.
func (o *CheckImageQualityResponseDataImageQualityQualityScoresCompressionQuality) HasValid() bool {
	if o != nil && !isNil(o.Valid) {
		return true
	}

	return false
}

// SetValid gets a reference to the given string and assigns it to the Valid field.
func (o *CheckImageQualityResponseDataImageQualityQualityScoresCompressionQuality) SetValid(v string) {
	o.Valid = &v
}

// GetScore returns the Score field value if set, zero value otherwise.
func (o *CheckImageQualityResponseDataImageQualityQualityScoresCompressionQuality) GetScore() float32 {
	if o == nil || isNil(o.Score) {
		var ret float32
		return ret
	}
	return *o.Score
}

// GetScoreOk returns a tuple with the Score field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckImageQualityResponseDataImageQualityQualityScoresCompressionQuality) GetScoreOk() (*float32, bool) {
	if o == nil || isNil(o.Score) {
    return nil, false
	}
	return o.Score, true
}

// HasScore returns a boolean if a field has been set.
func (o *CheckImageQualityResponseDataImageQualityQualityScoresCompressionQuality) HasScore() bool {
	if o != nil && !isNil(o.Score) {
		return true
	}

	return false
}

// SetScore gets a reference to the given float32 and assigns it to the Score field.
func (o *CheckImageQualityResponseDataImageQualityQualityScoresCompressionQuality) SetScore(v float32) {
	o.Score = &v
}

func (o CheckImageQualityResponseDataImageQualityQualityScoresCompressionQuality) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Valid) {
		toSerialize["valid"] = o.Valid
	}
	if !isNil(o.Score) {
		toSerialize["score"] = o.Score
	}
	return json.Marshal(toSerialize)
}

type NullableCheckImageQualityResponseDataImageQualityQualityScoresCompressionQuality struct {
	value *CheckImageQualityResponseDataImageQualityQualityScoresCompressionQuality
	isSet bool
}

func (v NullableCheckImageQualityResponseDataImageQualityQualityScoresCompressionQuality) Get() *CheckImageQualityResponseDataImageQualityQualityScoresCompressionQuality {
	return v.value
}

func (v *NullableCheckImageQualityResponseDataImageQualityQualityScoresCompressionQuality) Set(val *CheckImageQualityResponseDataImageQualityQualityScoresCompressionQuality) {
	v.value = val
	v.isSet = true
}

func (v NullableCheckImageQualityResponseDataImageQualityQualityScoresCompressionQuality) IsSet() bool {
	return v.isSet
}

func (v *NullableCheckImageQualityResponseDataImageQualityQualityScoresCompressionQuality) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCheckImageQualityResponseDataImageQualityQualityScoresCompressionQuality(val *CheckImageQualityResponseDataImageQualityQualityScoresCompressionQuality) *NullableCheckImageQualityResponseDataImageQualityQualityScoresCompressionQuality {
	return &NullableCheckImageQualityResponseDataImageQualityQualityScoresCompressionQuality{value: val, isSet: true}
}

func (v NullableCheckImageQualityResponseDataImageQualityQualityScoresCompressionQuality) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCheckImageQualityResponseDataImageQualityQualityScoresCompressionQuality) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


