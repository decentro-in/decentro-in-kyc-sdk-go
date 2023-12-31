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

// MatchFaceResponseData struct for MatchFaceResponseData
type MatchFaceResponseData struct {
	Status *string `json:"status,omitempty"`
	Match *string `json:"match,omitempty"`
}

// NewMatchFaceResponseData instantiates a new MatchFaceResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMatchFaceResponseData() *MatchFaceResponseData {
	this := MatchFaceResponseData{}
	return &this
}

// NewMatchFaceResponseDataWithDefaults instantiates a new MatchFaceResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMatchFaceResponseDataWithDefaults() *MatchFaceResponseData {
	this := MatchFaceResponseData{}
	return &this
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *MatchFaceResponseData) GetStatus() string {
	if o == nil || isNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MatchFaceResponseData) GetStatusOk() (*string, bool) {
	if o == nil || isNil(o.Status) {
    return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *MatchFaceResponseData) HasStatus() bool {
	if o != nil && !isNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *MatchFaceResponseData) SetStatus(v string) {
	o.Status = &v
}

// GetMatch returns the Match field value if set, zero value otherwise.
func (o *MatchFaceResponseData) GetMatch() string {
	if o == nil || isNil(o.Match) {
		var ret string
		return ret
	}
	return *o.Match
}

// GetMatchOk returns a tuple with the Match field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MatchFaceResponseData) GetMatchOk() (*string, bool) {
	if o == nil || isNil(o.Match) {
    return nil, false
	}
	return o.Match, true
}

// HasMatch returns a boolean if a field has been set.
func (o *MatchFaceResponseData) HasMatch() bool {
	if o != nil && !isNil(o.Match) {
		return true
	}

	return false
}

// SetMatch gets a reference to the given string and assigns it to the Match field.
func (o *MatchFaceResponseData) SetMatch(v string) {
	o.Match = &v
}

func (o MatchFaceResponseData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !isNil(o.Match) {
		toSerialize["match"] = o.Match
	}
	return json.Marshal(toSerialize)
}

type NullableMatchFaceResponseData struct {
	value *MatchFaceResponseData
	isSet bool
}

func (v NullableMatchFaceResponseData) Get() *MatchFaceResponseData {
	return v.value
}

func (v *NullableMatchFaceResponseData) Set(val *MatchFaceResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableMatchFaceResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableMatchFaceResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMatchFaceResponseData(val *MatchFaceResponseData) *NullableMatchFaceResponseData {
	return &NullableMatchFaceResponseData{value: val, isSet: true}
}

func (v NullableMatchFaceResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMatchFaceResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


