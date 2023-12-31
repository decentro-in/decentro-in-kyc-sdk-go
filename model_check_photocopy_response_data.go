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

// CheckPhotocopyResponseData struct for CheckPhotocopyResponseData
type CheckPhotocopyResponseData struct {
	IsPhotocopy *string `json:"isPhotocopy,omitempty"`
	Score *string `json:"score,omitempty"`
}

// NewCheckPhotocopyResponseData instantiates a new CheckPhotocopyResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCheckPhotocopyResponseData() *CheckPhotocopyResponseData {
	this := CheckPhotocopyResponseData{}
	return &this
}

// NewCheckPhotocopyResponseDataWithDefaults instantiates a new CheckPhotocopyResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCheckPhotocopyResponseDataWithDefaults() *CheckPhotocopyResponseData {
	this := CheckPhotocopyResponseData{}
	return &this
}

// GetIsPhotocopy returns the IsPhotocopy field value if set, zero value otherwise.
func (o *CheckPhotocopyResponseData) GetIsPhotocopy() string {
	if o == nil || isNil(o.IsPhotocopy) {
		var ret string
		return ret
	}
	return *o.IsPhotocopy
}

// GetIsPhotocopyOk returns a tuple with the IsPhotocopy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckPhotocopyResponseData) GetIsPhotocopyOk() (*string, bool) {
	if o == nil || isNil(o.IsPhotocopy) {
    return nil, false
	}
	return o.IsPhotocopy, true
}

// HasIsPhotocopy returns a boolean if a field has been set.
func (o *CheckPhotocopyResponseData) HasIsPhotocopy() bool {
	if o != nil && !isNil(o.IsPhotocopy) {
		return true
	}

	return false
}

// SetIsPhotocopy gets a reference to the given string and assigns it to the IsPhotocopy field.
func (o *CheckPhotocopyResponseData) SetIsPhotocopy(v string) {
	o.IsPhotocopy = &v
}

// GetScore returns the Score field value if set, zero value otherwise.
func (o *CheckPhotocopyResponseData) GetScore() string {
	if o == nil || isNil(o.Score) {
		var ret string
		return ret
	}
	return *o.Score
}

// GetScoreOk returns a tuple with the Score field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckPhotocopyResponseData) GetScoreOk() (*string, bool) {
	if o == nil || isNil(o.Score) {
    return nil, false
	}
	return o.Score, true
}

// HasScore returns a boolean if a field has been set.
func (o *CheckPhotocopyResponseData) HasScore() bool {
	if o != nil && !isNil(o.Score) {
		return true
	}

	return false
}

// SetScore gets a reference to the given string and assigns it to the Score field.
func (o *CheckPhotocopyResponseData) SetScore(v string) {
	o.Score = &v
}

func (o CheckPhotocopyResponseData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.IsPhotocopy) {
		toSerialize["isPhotocopy"] = o.IsPhotocopy
	}
	if !isNil(o.Score) {
		toSerialize["score"] = o.Score
	}
	return json.Marshal(toSerialize)
}

type NullableCheckPhotocopyResponseData struct {
	value *CheckPhotocopyResponseData
	isSet bool
}

func (v NullableCheckPhotocopyResponseData) Get() *CheckPhotocopyResponseData {
	return v.value
}

func (v *NullableCheckPhotocopyResponseData) Set(val *CheckPhotocopyResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableCheckPhotocopyResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableCheckPhotocopyResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCheckPhotocopyResponseData(val *CheckPhotocopyResponseData) *NullableCheckPhotocopyResponseData {
	return &NullableCheckPhotocopyResponseData{value: val, isSet: true}
}

func (v NullableCheckPhotocopyResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCheckPhotocopyResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


