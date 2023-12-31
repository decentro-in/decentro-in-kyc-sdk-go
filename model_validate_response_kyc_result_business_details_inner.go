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

// ValidateResponseKycResultBusinessDetailsInner struct for ValidateResponseKycResultBusinessDetailsInner
type ValidateResponseKycResultBusinessDetailsInner struct {
	Hsn *string `json:"hsn,omitempty"`
	ServiceDetail *string `json:"serviceDetail,omitempty"`
}

// NewValidateResponseKycResultBusinessDetailsInner instantiates a new ValidateResponseKycResultBusinessDetailsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewValidateResponseKycResultBusinessDetailsInner() *ValidateResponseKycResultBusinessDetailsInner {
	this := ValidateResponseKycResultBusinessDetailsInner{}
	return &this
}

// NewValidateResponseKycResultBusinessDetailsInnerWithDefaults instantiates a new ValidateResponseKycResultBusinessDetailsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewValidateResponseKycResultBusinessDetailsInnerWithDefaults() *ValidateResponseKycResultBusinessDetailsInner {
	this := ValidateResponseKycResultBusinessDetailsInner{}
	return &this
}

// GetHsn returns the Hsn field value if set, zero value otherwise.
func (o *ValidateResponseKycResultBusinessDetailsInner) GetHsn() string {
	if o == nil || isNil(o.Hsn) {
		var ret string
		return ret
	}
	return *o.Hsn
}

// GetHsnOk returns a tuple with the Hsn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidateResponseKycResultBusinessDetailsInner) GetHsnOk() (*string, bool) {
	if o == nil || isNil(o.Hsn) {
    return nil, false
	}
	return o.Hsn, true
}

// HasHsn returns a boolean if a field has been set.
func (o *ValidateResponseKycResultBusinessDetailsInner) HasHsn() bool {
	if o != nil && !isNil(o.Hsn) {
		return true
	}

	return false
}

// SetHsn gets a reference to the given string and assigns it to the Hsn field.
func (o *ValidateResponseKycResultBusinessDetailsInner) SetHsn(v string) {
	o.Hsn = &v
}

// GetServiceDetail returns the ServiceDetail field value if set, zero value otherwise.
func (o *ValidateResponseKycResultBusinessDetailsInner) GetServiceDetail() string {
	if o == nil || isNil(o.ServiceDetail) {
		var ret string
		return ret
	}
	return *o.ServiceDetail
}

// GetServiceDetailOk returns a tuple with the ServiceDetail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidateResponseKycResultBusinessDetailsInner) GetServiceDetailOk() (*string, bool) {
	if o == nil || isNil(o.ServiceDetail) {
    return nil, false
	}
	return o.ServiceDetail, true
}

// HasServiceDetail returns a boolean if a field has been set.
func (o *ValidateResponseKycResultBusinessDetailsInner) HasServiceDetail() bool {
	if o != nil && !isNil(o.ServiceDetail) {
		return true
	}

	return false
}

// SetServiceDetail gets a reference to the given string and assigns it to the ServiceDetail field.
func (o *ValidateResponseKycResultBusinessDetailsInner) SetServiceDetail(v string) {
	o.ServiceDetail = &v
}

func (o ValidateResponseKycResultBusinessDetailsInner) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Hsn) {
		toSerialize["hsn"] = o.Hsn
	}
	if !isNil(o.ServiceDetail) {
		toSerialize["serviceDetail"] = o.ServiceDetail
	}
	return json.Marshal(toSerialize)
}

type NullableValidateResponseKycResultBusinessDetailsInner struct {
	value *ValidateResponseKycResultBusinessDetailsInner
	isSet bool
}

func (v NullableValidateResponseKycResultBusinessDetailsInner) Get() *ValidateResponseKycResultBusinessDetailsInner {
	return v.value
}

func (v *NullableValidateResponseKycResultBusinessDetailsInner) Set(val *ValidateResponseKycResultBusinessDetailsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableValidateResponseKycResultBusinessDetailsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableValidateResponseKycResultBusinessDetailsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableValidateResponseKycResultBusinessDetailsInner(val *ValidateResponseKycResultBusinessDetailsInner) *NullableValidateResponseKycResultBusinessDetailsInner {
	return &NullableValidateResponseKycResultBusinessDetailsInner{value: val, isSet: true}
}

func (v NullableValidateResponseKycResultBusinessDetailsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableValidateResponseKycResultBusinessDetailsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


