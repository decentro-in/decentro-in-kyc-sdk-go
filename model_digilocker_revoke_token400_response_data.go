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

// DigilockerRevokeToken400ResponseData struct for DigilockerRevokeToken400ResponseData
type DigilockerRevokeToken400ResponseData struct {
	Revoked *bool `json:"revoked,omitempty"`
}

// NewDigilockerRevokeToken400ResponseData instantiates a new DigilockerRevokeToken400ResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDigilockerRevokeToken400ResponseData() *DigilockerRevokeToken400ResponseData {
	this := DigilockerRevokeToken400ResponseData{}
	return &this
}

// NewDigilockerRevokeToken400ResponseDataWithDefaults instantiates a new DigilockerRevokeToken400ResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDigilockerRevokeToken400ResponseDataWithDefaults() *DigilockerRevokeToken400ResponseData {
	this := DigilockerRevokeToken400ResponseData{}
	return &this
}

// GetRevoked returns the Revoked field value if set, zero value otherwise.
func (o *DigilockerRevokeToken400ResponseData) GetRevoked() bool {
	if o == nil || isNil(o.Revoked) {
		var ret bool
		return ret
	}
	return *o.Revoked
}

// GetRevokedOk returns a tuple with the Revoked field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DigilockerRevokeToken400ResponseData) GetRevokedOk() (*bool, bool) {
	if o == nil || isNil(o.Revoked) {
    return nil, false
	}
	return o.Revoked, true
}

// HasRevoked returns a boolean if a field has been set.
func (o *DigilockerRevokeToken400ResponseData) HasRevoked() bool {
	if o != nil && !isNil(o.Revoked) {
		return true
	}

	return false
}

// SetRevoked gets a reference to the given bool and assigns it to the Revoked field.
func (o *DigilockerRevokeToken400ResponseData) SetRevoked(v bool) {
	o.Revoked = &v
}

func (o DigilockerRevokeToken400ResponseData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Revoked) {
		toSerialize["revoked"] = o.Revoked
	}
	return json.Marshal(toSerialize)
}

type NullableDigilockerRevokeToken400ResponseData struct {
	value *DigilockerRevokeToken400ResponseData
	isSet bool
}

func (v NullableDigilockerRevokeToken400ResponseData) Get() *DigilockerRevokeToken400ResponseData {
	return v.value
}

func (v *NullableDigilockerRevokeToken400ResponseData) Set(val *DigilockerRevokeToken400ResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableDigilockerRevokeToken400ResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableDigilockerRevokeToken400ResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDigilockerRevokeToken400ResponseData(val *DigilockerRevokeToken400ResponseData) *NullableDigilockerRevokeToken400ResponseData {
	return &NullableDigilockerRevokeToken400ResponseData{value: val, isSet: true}
}

func (v NullableDigilockerRevokeToken400ResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDigilockerRevokeToken400ResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


