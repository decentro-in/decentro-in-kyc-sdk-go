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

// ValidateResponseKycResultPremissesAddress struct for ValidateResponseKycResultPremissesAddress
type ValidateResponseKycResultPremissesAddress struct {
	District *string `json:"district,omitempty"`
	Address *string `json:"address,omitempty"`
	State *string `json:"state,omitempty"`
	Pincode *string `json:"pincode,omitempty"`
}

// NewValidateResponseKycResultPremissesAddress instantiates a new ValidateResponseKycResultPremissesAddress object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewValidateResponseKycResultPremissesAddress() *ValidateResponseKycResultPremissesAddress {
	this := ValidateResponseKycResultPremissesAddress{}
	return &this
}

// NewValidateResponseKycResultPremissesAddressWithDefaults instantiates a new ValidateResponseKycResultPremissesAddress object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewValidateResponseKycResultPremissesAddressWithDefaults() *ValidateResponseKycResultPremissesAddress {
	this := ValidateResponseKycResultPremissesAddress{}
	return &this
}

// GetDistrict returns the District field value if set, zero value otherwise.
func (o *ValidateResponseKycResultPremissesAddress) GetDistrict() string {
	if o == nil || isNil(o.District) {
		var ret string
		return ret
	}
	return *o.District
}

// GetDistrictOk returns a tuple with the District field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidateResponseKycResultPremissesAddress) GetDistrictOk() (*string, bool) {
	if o == nil || isNil(o.District) {
    return nil, false
	}
	return o.District, true
}

// HasDistrict returns a boolean if a field has been set.
func (o *ValidateResponseKycResultPremissesAddress) HasDistrict() bool {
	if o != nil && !isNil(o.District) {
		return true
	}

	return false
}

// SetDistrict gets a reference to the given string and assigns it to the District field.
func (o *ValidateResponseKycResultPremissesAddress) SetDistrict(v string) {
	o.District = &v
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *ValidateResponseKycResultPremissesAddress) GetAddress() string {
	if o == nil || isNil(o.Address) {
		var ret string
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidateResponseKycResultPremissesAddress) GetAddressOk() (*string, bool) {
	if o == nil || isNil(o.Address) {
    return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *ValidateResponseKycResultPremissesAddress) HasAddress() bool {
	if o != nil && !isNil(o.Address) {
		return true
	}

	return false
}

// SetAddress gets a reference to the given string and assigns it to the Address field.
func (o *ValidateResponseKycResultPremissesAddress) SetAddress(v string) {
	o.Address = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *ValidateResponseKycResultPremissesAddress) GetState() string {
	if o == nil || isNil(o.State) {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidateResponseKycResultPremissesAddress) GetStateOk() (*string, bool) {
	if o == nil || isNil(o.State) {
    return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *ValidateResponseKycResultPremissesAddress) HasState() bool {
	if o != nil && !isNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *ValidateResponseKycResultPremissesAddress) SetState(v string) {
	o.State = &v
}

// GetPincode returns the Pincode field value if set, zero value otherwise.
func (o *ValidateResponseKycResultPremissesAddress) GetPincode() string {
	if o == nil || isNil(o.Pincode) {
		var ret string
		return ret
	}
	return *o.Pincode
}

// GetPincodeOk returns a tuple with the Pincode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidateResponseKycResultPremissesAddress) GetPincodeOk() (*string, bool) {
	if o == nil || isNil(o.Pincode) {
    return nil, false
	}
	return o.Pincode, true
}

// HasPincode returns a boolean if a field has been set.
func (o *ValidateResponseKycResultPremissesAddress) HasPincode() bool {
	if o != nil && !isNil(o.Pincode) {
		return true
	}

	return false
}

// SetPincode gets a reference to the given string and assigns it to the Pincode field.
func (o *ValidateResponseKycResultPremissesAddress) SetPincode(v string) {
	o.Pincode = &v
}

func (o ValidateResponseKycResultPremissesAddress) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.District) {
		toSerialize["district"] = o.District
	}
	if !isNil(o.Address) {
		toSerialize["address"] = o.Address
	}
	if !isNil(o.State) {
		toSerialize["state"] = o.State
	}
	if !isNil(o.Pincode) {
		toSerialize["pincode"] = o.Pincode
	}
	return json.Marshal(toSerialize)
}

type NullableValidateResponseKycResultPremissesAddress struct {
	value *ValidateResponseKycResultPremissesAddress
	isSet bool
}

func (v NullableValidateResponseKycResultPremissesAddress) Get() *ValidateResponseKycResultPremissesAddress {
	return v.value
}

func (v *NullableValidateResponseKycResultPremissesAddress) Set(val *ValidateResponseKycResultPremissesAddress) {
	v.value = val
	v.isSet = true
}

func (v NullableValidateResponseKycResultPremissesAddress) IsSet() bool {
	return v.isSet
}

func (v *NullableValidateResponseKycResultPremissesAddress) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableValidateResponseKycResultPremissesAddress(val *ValidateResponseKycResultPremissesAddress) *NullableValidateResponseKycResultPremissesAddress {
	return &NullableValidateResponseKycResultPremissesAddress{value: val, isSet: true}
}

func (v NullableValidateResponseKycResultPremissesAddress) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableValidateResponseKycResultPremissesAddress) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


