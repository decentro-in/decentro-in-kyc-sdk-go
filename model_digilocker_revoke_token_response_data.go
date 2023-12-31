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

// DigilockerRevokeTokenResponseData struct for DigilockerRevokeTokenResponseData
type DigilockerRevokeTokenResponseData struct {
	Revoked *bool `json:"revoked,omitempty"`
}

// NewDigilockerRevokeTokenResponseData instantiates a new DigilockerRevokeTokenResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDigilockerRevokeTokenResponseData() *DigilockerRevokeTokenResponseData {
	this := DigilockerRevokeTokenResponseData{}
	return &this
}

// NewDigilockerRevokeTokenResponseDataWithDefaults instantiates a new DigilockerRevokeTokenResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDigilockerRevokeTokenResponseDataWithDefaults() *DigilockerRevokeTokenResponseData {
	this := DigilockerRevokeTokenResponseData{}
	return &this
}

// GetRevoked returns the Revoked field value if set, zero value otherwise.
func (o *DigilockerRevokeTokenResponseData) GetRevoked() bool {
	if o == nil || isNil(o.Revoked) {
		var ret bool
		return ret
	}
	return *o.Revoked
}

// GetRevokedOk returns a tuple with the Revoked field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DigilockerRevokeTokenResponseData) GetRevokedOk() (*bool, bool) {
	if o == nil || isNil(o.Revoked) {
    return nil, false
	}
	return o.Revoked, true
}

// HasRevoked returns a boolean if a field has been set.
func (o *DigilockerRevokeTokenResponseData) HasRevoked() bool {
	if o != nil && !isNil(o.Revoked) {
		return true
	}

	return false
}

// SetRevoked gets a reference to the given bool and assigns it to the Revoked field.
func (o *DigilockerRevokeTokenResponseData) SetRevoked(v bool) {
	o.Revoked = &v
}

func (o DigilockerRevokeTokenResponseData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Revoked) {
		toSerialize["revoked"] = o.Revoked
	}
	return json.Marshal(toSerialize)
}

type NullableDigilockerRevokeTokenResponseData struct {
	value *DigilockerRevokeTokenResponseData
	isSet bool
}

func (v NullableDigilockerRevokeTokenResponseData) Get() *DigilockerRevokeTokenResponseData {
	return v.value
}

func (v *NullableDigilockerRevokeTokenResponseData) Set(val *DigilockerRevokeTokenResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableDigilockerRevokeTokenResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableDigilockerRevokeTokenResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDigilockerRevokeTokenResponseData(val *DigilockerRevokeTokenResponseData) *NullableDigilockerRevokeTokenResponseData {
	return &NullableDigilockerRevokeTokenResponseData{value: val, isSet: true}
}

func (v NullableDigilockerRevokeTokenResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDigilockerRevokeTokenResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


