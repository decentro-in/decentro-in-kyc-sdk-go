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

// ValidateResponseKycResultDirectorData struct for ValidateResponseKycResultDirectorData
type ValidateResponseKycResultDirectorData struct {
	Din *string `json:"din,omitempty"`
	Name *string `json:"name,omitempty"`
}

// NewValidateResponseKycResultDirectorData instantiates a new ValidateResponseKycResultDirectorData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewValidateResponseKycResultDirectorData() *ValidateResponseKycResultDirectorData {
	this := ValidateResponseKycResultDirectorData{}
	return &this
}

// NewValidateResponseKycResultDirectorDataWithDefaults instantiates a new ValidateResponseKycResultDirectorData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewValidateResponseKycResultDirectorDataWithDefaults() *ValidateResponseKycResultDirectorData {
	this := ValidateResponseKycResultDirectorData{}
	return &this
}

// GetDin returns the Din field value if set, zero value otherwise.
func (o *ValidateResponseKycResultDirectorData) GetDin() string {
	if o == nil || isNil(o.Din) {
		var ret string
		return ret
	}
	return *o.Din
}

// GetDinOk returns a tuple with the Din field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidateResponseKycResultDirectorData) GetDinOk() (*string, bool) {
	if o == nil || isNil(o.Din) {
    return nil, false
	}
	return o.Din, true
}

// HasDin returns a boolean if a field has been set.
func (o *ValidateResponseKycResultDirectorData) HasDin() bool {
	if o != nil && !isNil(o.Din) {
		return true
	}

	return false
}

// SetDin gets a reference to the given string and assigns it to the Din field.
func (o *ValidateResponseKycResultDirectorData) SetDin(v string) {
	o.Din = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ValidateResponseKycResultDirectorData) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidateResponseKycResultDirectorData) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ValidateResponseKycResultDirectorData) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ValidateResponseKycResultDirectorData) SetName(v string) {
	o.Name = &v
}

func (o ValidateResponseKycResultDirectorData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Din) {
		toSerialize["din"] = o.Din
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

type NullableValidateResponseKycResultDirectorData struct {
	value *ValidateResponseKycResultDirectorData
	isSet bool
}

func (v NullableValidateResponseKycResultDirectorData) Get() *ValidateResponseKycResultDirectorData {
	return v.value
}

func (v *NullableValidateResponseKycResultDirectorData) Set(val *ValidateResponseKycResultDirectorData) {
	v.value = val
	v.isSet = true
}

func (v NullableValidateResponseKycResultDirectorData) IsSet() bool {
	return v.isSet
}

func (v *NullableValidateResponseKycResultDirectorData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableValidateResponseKycResultDirectorData(val *ValidateResponseKycResultDirectorData) *NullableValidateResponseKycResultDirectorData {
	return &NullableValidateResponseKycResultDirectorData{value: val, isSet: true}
}

func (v NullableValidateResponseKycResultDirectorData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableValidateResponseKycResultDirectorData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


