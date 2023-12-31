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

// DigilockerPullFileRequestDocumentTypeRelatedParameters struct for DigilockerPullFileRequestDocumentTypeRelatedParameters
type DigilockerPullFileRequestDocumentTypeRelatedParameters struct {
	IdNumber *string `json:"id_number,omitempty"`
	FullName *string `json:"full_name,omitempty"`
}

// NewDigilockerPullFileRequestDocumentTypeRelatedParameters instantiates a new DigilockerPullFileRequestDocumentTypeRelatedParameters object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDigilockerPullFileRequestDocumentTypeRelatedParameters() *DigilockerPullFileRequestDocumentTypeRelatedParameters {
	this := DigilockerPullFileRequestDocumentTypeRelatedParameters{}
	return &this
}

// NewDigilockerPullFileRequestDocumentTypeRelatedParametersWithDefaults instantiates a new DigilockerPullFileRequestDocumentTypeRelatedParameters object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDigilockerPullFileRequestDocumentTypeRelatedParametersWithDefaults() *DigilockerPullFileRequestDocumentTypeRelatedParameters {
	this := DigilockerPullFileRequestDocumentTypeRelatedParameters{}
	return &this
}

// GetIdNumber returns the IdNumber field value if set, zero value otherwise.
func (o *DigilockerPullFileRequestDocumentTypeRelatedParameters) GetIdNumber() string {
	if o == nil || isNil(o.IdNumber) {
		var ret string
		return ret
	}
	return *o.IdNumber
}

// GetIdNumberOk returns a tuple with the IdNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DigilockerPullFileRequestDocumentTypeRelatedParameters) GetIdNumberOk() (*string, bool) {
	if o == nil || isNil(o.IdNumber) {
    return nil, false
	}
	return o.IdNumber, true
}

// HasIdNumber returns a boolean if a field has been set.
func (o *DigilockerPullFileRequestDocumentTypeRelatedParameters) HasIdNumber() bool {
	if o != nil && !isNil(o.IdNumber) {
		return true
	}

	return false
}

// SetIdNumber gets a reference to the given string and assigns it to the IdNumber field.
func (o *DigilockerPullFileRequestDocumentTypeRelatedParameters) SetIdNumber(v string) {
	o.IdNumber = &v
}

// GetFullName returns the FullName field value if set, zero value otherwise.
func (o *DigilockerPullFileRequestDocumentTypeRelatedParameters) GetFullName() string {
	if o == nil || isNil(o.FullName) {
		var ret string
		return ret
	}
	return *o.FullName
}

// GetFullNameOk returns a tuple with the FullName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DigilockerPullFileRequestDocumentTypeRelatedParameters) GetFullNameOk() (*string, bool) {
	if o == nil || isNil(o.FullName) {
    return nil, false
	}
	return o.FullName, true
}

// HasFullName returns a boolean if a field has been set.
func (o *DigilockerPullFileRequestDocumentTypeRelatedParameters) HasFullName() bool {
	if o != nil && !isNil(o.FullName) {
		return true
	}

	return false
}

// SetFullName gets a reference to the given string and assigns it to the FullName field.
func (o *DigilockerPullFileRequestDocumentTypeRelatedParameters) SetFullName(v string) {
	o.FullName = &v
}

func (o DigilockerPullFileRequestDocumentTypeRelatedParameters) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.IdNumber) {
		toSerialize["id_number"] = o.IdNumber
	}
	if !isNil(o.FullName) {
		toSerialize["full_name"] = o.FullName
	}
	return json.Marshal(toSerialize)
}

type NullableDigilockerPullFileRequestDocumentTypeRelatedParameters struct {
	value *DigilockerPullFileRequestDocumentTypeRelatedParameters
	isSet bool
}

func (v NullableDigilockerPullFileRequestDocumentTypeRelatedParameters) Get() *DigilockerPullFileRequestDocumentTypeRelatedParameters {
	return v.value
}

func (v *NullableDigilockerPullFileRequestDocumentTypeRelatedParameters) Set(val *DigilockerPullFileRequestDocumentTypeRelatedParameters) {
	v.value = val
	v.isSet = true
}

func (v NullableDigilockerPullFileRequestDocumentTypeRelatedParameters) IsSet() bool {
	return v.isSet
}

func (v *NullableDigilockerPullFileRequestDocumentTypeRelatedParameters) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDigilockerPullFileRequestDocumentTypeRelatedParameters(val *DigilockerPullFileRequestDocumentTypeRelatedParameters) *NullableDigilockerPullFileRequestDocumentTypeRelatedParameters {
	return &NullableDigilockerPullFileRequestDocumentTypeRelatedParameters{value: val, isSet: true}
}

func (v NullableDigilockerPullFileRequestDocumentTypeRelatedParameters) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDigilockerPullFileRequestDocumentTypeRelatedParameters) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


