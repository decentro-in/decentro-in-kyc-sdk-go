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

// ValidateResponseKycResultAddressesInner struct for ValidateResponseKycResultAddressesInner
type ValidateResponseKycResultAddressesInner struct {
	AddressLine *string `json:"addressLine,omitempty"`
	CompleteAddress *string `json:"completeAddress,omitempty"`
	Country *string `json:"country,omitempty"`
	District *string `json:"district,omitempty"`
	Pin *string `json:"pin,omitempty"`
	State *string `json:"state,omitempty"`
	Type *string `json:"type,omitempty"`
}

// NewValidateResponseKycResultAddressesInner instantiates a new ValidateResponseKycResultAddressesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewValidateResponseKycResultAddressesInner() *ValidateResponseKycResultAddressesInner {
	this := ValidateResponseKycResultAddressesInner{}
	return &this
}

// NewValidateResponseKycResultAddressesInnerWithDefaults instantiates a new ValidateResponseKycResultAddressesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewValidateResponseKycResultAddressesInnerWithDefaults() *ValidateResponseKycResultAddressesInner {
	this := ValidateResponseKycResultAddressesInner{}
	return &this
}

// GetAddressLine returns the AddressLine field value if set, zero value otherwise.
func (o *ValidateResponseKycResultAddressesInner) GetAddressLine() string {
	if o == nil || isNil(o.AddressLine) {
		var ret string
		return ret
	}
	return *o.AddressLine
}

// GetAddressLineOk returns a tuple with the AddressLine field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidateResponseKycResultAddressesInner) GetAddressLineOk() (*string, bool) {
	if o == nil || isNil(o.AddressLine) {
    return nil, false
	}
	return o.AddressLine, true
}

// HasAddressLine returns a boolean if a field has been set.
func (o *ValidateResponseKycResultAddressesInner) HasAddressLine() bool {
	if o != nil && !isNil(o.AddressLine) {
		return true
	}

	return false
}

// SetAddressLine gets a reference to the given string and assigns it to the AddressLine field.
func (o *ValidateResponseKycResultAddressesInner) SetAddressLine(v string) {
	o.AddressLine = &v
}

// GetCompleteAddress returns the CompleteAddress field value if set, zero value otherwise.
func (o *ValidateResponseKycResultAddressesInner) GetCompleteAddress() string {
	if o == nil || isNil(o.CompleteAddress) {
		var ret string
		return ret
	}
	return *o.CompleteAddress
}

// GetCompleteAddressOk returns a tuple with the CompleteAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidateResponseKycResultAddressesInner) GetCompleteAddressOk() (*string, bool) {
	if o == nil || isNil(o.CompleteAddress) {
    return nil, false
	}
	return o.CompleteAddress, true
}

// HasCompleteAddress returns a boolean if a field has been set.
func (o *ValidateResponseKycResultAddressesInner) HasCompleteAddress() bool {
	if o != nil && !isNil(o.CompleteAddress) {
		return true
	}

	return false
}

// SetCompleteAddress gets a reference to the given string and assigns it to the CompleteAddress field.
func (o *ValidateResponseKycResultAddressesInner) SetCompleteAddress(v string) {
	o.CompleteAddress = &v
}

// GetCountry returns the Country field value if set, zero value otherwise.
func (o *ValidateResponseKycResultAddressesInner) GetCountry() string {
	if o == nil || isNil(o.Country) {
		var ret string
		return ret
	}
	return *o.Country
}

// GetCountryOk returns a tuple with the Country field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidateResponseKycResultAddressesInner) GetCountryOk() (*string, bool) {
	if o == nil || isNil(o.Country) {
    return nil, false
	}
	return o.Country, true
}

// HasCountry returns a boolean if a field has been set.
func (o *ValidateResponseKycResultAddressesInner) HasCountry() bool {
	if o != nil && !isNil(o.Country) {
		return true
	}

	return false
}

// SetCountry gets a reference to the given string and assigns it to the Country field.
func (o *ValidateResponseKycResultAddressesInner) SetCountry(v string) {
	o.Country = &v
}

// GetDistrict returns the District field value if set, zero value otherwise.
func (o *ValidateResponseKycResultAddressesInner) GetDistrict() string {
	if o == nil || isNil(o.District) {
		var ret string
		return ret
	}
	return *o.District
}

// GetDistrictOk returns a tuple with the District field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidateResponseKycResultAddressesInner) GetDistrictOk() (*string, bool) {
	if o == nil || isNil(o.District) {
    return nil, false
	}
	return o.District, true
}

// HasDistrict returns a boolean if a field has been set.
func (o *ValidateResponseKycResultAddressesInner) HasDistrict() bool {
	if o != nil && !isNil(o.District) {
		return true
	}

	return false
}

// SetDistrict gets a reference to the given string and assigns it to the District field.
func (o *ValidateResponseKycResultAddressesInner) SetDistrict(v string) {
	o.District = &v
}

// GetPin returns the Pin field value if set, zero value otherwise.
func (o *ValidateResponseKycResultAddressesInner) GetPin() string {
	if o == nil || isNil(o.Pin) {
		var ret string
		return ret
	}
	return *o.Pin
}

// GetPinOk returns a tuple with the Pin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidateResponseKycResultAddressesInner) GetPinOk() (*string, bool) {
	if o == nil || isNil(o.Pin) {
    return nil, false
	}
	return o.Pin, true
}

// HasPin returns a boolean if a field has been set.
func (o *ValidateResponseKycResultAddressesInner) HasPin() bool {
	if o != nil && !isNil(o.Pin) {
		return true
	}

	return false
}

// SetPin gets a reference to the given string and assigns it to the Pin field.
func (o *ValidateResponseKycResultAddressesInner) SetPin(v string) {
	o.Pin = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *ValidateResponseKycResultAddressesInner) GetState() string {
	if o == nil || isNil(o.State) {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidateResponseKycResultAddressesInner) GetStateOk() (*string, bool) {
	if o == nil || isNil(o.State) {
    return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *ValidateResponseKycResultAddressesInner) HasState() bool {
	if o != nil && !isNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *ValidateResponseKycResultAddressesInner) SetState(v string) {
	o.State = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ValidateResponseKycResultAddressesInner) GetType() string {
	if o == nil || isNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidateResponseKycResultAddressesInner) GetTypeOk() (*string, bool) {
	if o == nil || isNil(o.Type) {
    return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ValidateResponseKycResultAddressesInner) HasType() bool {
	if o != nil && !isNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *ValidateResponseKycResultAddressesInner) SetType(v string) {
	o.Type = &v
}

func (o ValidateResponseKycResultAddressesInner) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.AddressLine) {
		toSerialize["addressLine"] = o.AddressLine
	}
	if !isNil(o.CompleteAddress) {
		toSerialize["completeAddress"] = o.CompleteAddress
	}
	if !isNil(o.Country) {
		toSerialize["country"] = o.Country
	}
	if !isNil(o.District) {
		toSerialize["district"] = o.District
	}
	if !isNil(o.Pin) {
		toSerialize["pin"] = o.Pin
	}
	if !isNil(o.State) {
		toSerialize["state"] = o.State
	}
	if !isNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableValidateResponseKycResultAddressesInner struct {
	value *ValidateResponseKycResultAddressesInner
	isSet bool
}

func (v NullableValidateResponseKycResultAddressesInner) Get() *ValidateResponseKycResultAddressesInner {
	return v.value
}

func (v *NullableValidateResponseKycResultAddressesInner) Set(val *ValidateResponseKycResultAddressesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableValidateResponseKycResultAddressesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableValidateResponseKycResultAddressesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableValidateResponseKycResultAddressesInner(val *ValidateResponseKycResultAddressesInner) *NullableValidateResponseKycResultAddressesInner {
	return &NullableValidateResponseKycResultAddressesInner{value: val, isSet: true}
}

func (v NullableValidateResponseKycResultAddressesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableValidateResponseKycResultAddressesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


