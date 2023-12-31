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

// ValidateResponseKycResultPrimaryBusinessContact struct for ValidateResponseKycResultPrimaryBusinessContact
type ValidateResponseKycResultPrimaryBusinessContact struct {
	Email *string `json:"email,omitempty"`
	Address *string `json:"address,omitempty"`
	MobileNumber *ValidateResponseKycResultPrimaryBusinessContactMobileNumber `json:"mobileNumber,omitempty"`
	NatureOfBusinessAtAddress *string `json:"natureOfBusinessAtAddress,omitempty"`
	LastUpdatedDate *string `json:"lastUpdatedDate,omitempty"`
}

// NewValidateResponseKycResultPrimaryBusinessContact instantiates a new ValidateResponseKycResultPrimaryBusinessContact object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewValidateResponseKycResultPrimaryBusinessContact() *ValidateResponseKycResultPrimaryBusinessContact {
	this := ValidateResponseKycResultPrimaryBusinessContact{}
	return &this
}

// NewValidateResponseKycResultPrimaryBusinessContactWithDefaults instantiates a new ValidateResponseKycResultPrimaryBusinessContact object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewValidateResponseKycResultPrimaryBusinessContactWithDefaults() *ValidateResponseKycResultPrimaryBusinessContact {
	this := ValidateResponseKycResultPrimaryBusinessContact{}
	return &this
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *ValidateResponseKycResultPrimaryBusinessContact) GetEmail() string {
	if o == nil || isNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidateResponseKycResultPrimaryBusinessContact) GetEmailOk() (*string, bool) {
	if o == nil || isNil(o.Email) {
    return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *ValidateResponseKycResultPrimaryBusinessContact) HasEmail() bool {
	if o != nil && !isNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *ValidateResponseKycResultPrimaryBusinessContact) SetEmail(v string) {
	o.Email = &v
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *ValidateResponseKycResultPrimaryBusinessContact) GetAddress() string {
	if o == nil || isNil(o.Address) {
		var ret string
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidateResponseKycResultPrimaryBusinessContact) GetAddressOk() (*string, bool) {
	if o == nil || isNil(o.Address) {
    return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *ValidateResponseKycResultPrimaryBusinessContact) HasAddress() bool {
	if o != nil && !isNil(o.Address) {
		return true
	}

	return false
}

// SetAddress gets a reference to the given string and assigns it to the Address field.
func (o *ValidateResponseKycResultPrimaryBusinessContact) SetAddress(v string) {
	o.Address = &v
}

// GetMobileNumber returns the MobileNumber field value if set, zero value otherwise.
func (o *ValidateResponseKycResultPrimaryBusinessContact) GetMobileNumber() ValidateResponseKycResultPrimaryBusinessContactMobileNumber {
	if o == nil || isNil(o.MobileNumber) {
		var ret ValidateResponseKycResultPrimaryBusinessContactMobileNumber
		return ret
	}
	return *o.MobileNumber
}

// GetMobileNumberOk returns a tuple with the MobileNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidateResponseKycResultPrimaryBusinessContact) GetMobileNumberOk() (*ValidateResponseKycResultPrimaryBusinessContactMobileNumber, bool) {
	if o == nil || isNil(o.MobileNumber) {
    return nil, false
	}
	return o.MobileNumber, true
}

// HasMobileNumber returns a boolean if a field has been set.
func (o *ValidateResponseKycResultPrimaryBusinessContact) HasMobileNumber() bool {
	if o != nil && !isNil(o.MobileNumber) {
		return true
	}

	return false
}

// SetMobileNumber gets a reference to the given ValidateResponseKycResultPrimaryBusinessContactMobileNumber and assigns it to the MobileNumber field.
func (o *ValidateResponseKycResultPrimaryBusinessContact) SetMobileNumber(v ValidateResponseKycResultPrimaryBusinessContactMobileNumber) {
	o.MobileNumber = &v
}

// GetNatureOfBusinessAtAddress returns the NatureOfBusinessAtAddress field value if set, zero value otherwise.
func (o *ValidateResponseKycResultPrimaryBusinessContact) GetNatureOfBusinessAtAddress() string {
	if o == nil || isNil(o.NatureOfBusinessAtAddress) {
		var ret string
		return ret
	}
	return *o.NatureOfBusinessAtAddress
}

// GetNatureOfBusinessAtAddressOk returns a tuple with the NatureOfBusinessAtAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidateResponseKycResultPrimaryBusinessContact) GetNatureOfBusinessAtAddressOk() (*string, bool) {
	if o == nil || isNil(o.NatureOfBusinessAtAddress) {
    return nil, false
	}
	return o.NatureOfBusinessAtAddress, true
}

// HasNatureOfBusinessAtAddress returns a boolean if a field has been set.
func (o *ValidateResponseKycResultPrimaryBusinessContact) HasNatureOfBusinessAtAddress() bool {
	if o != nil && !isNil(o.NatureOfBusinessAtAddress) {
		return true
	}

	return false
}

// SetNatureOfBusinessAtAddress gets a reference to the given string and assigns it to the NatureOfBusinessAtAddress field.
func (o *ValidateResponseKycResultPrimaryBusinessContact) SetNatureOfBusinessAtAddress(v string) {
	o.NatureOfBusinessAtAddress = &v
}

// GetLastUpdatedDate returns the LastUpdatedDate field value if set, zero value otherwise.
func (o *ValidateResponseKycResultPrimaryBusinessContact) GetLastUpdatedDate() string {
	if o == nil || isNil(o.LastUpdatedDate) {
		var ret string
		return ret
	}
	return *o.LastUpdatedDate
}

// GetLastUpdatedDateOk returns a tuple with the LastUpdatedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidateResponseKycResultPrimaryBusinessContact) GetLastUpdatedDateOk() (*string, bool) {
	if o == nil || isNil(o.LastUpdatedDate) {
    return nil, false
	}
	return o.LastUpdatedDate, true
}

// HasLastUpdatedDate returns a boolean if a field has been set.
func (o *ValidateResponseKycResultPrimaryBusinessContact) HasLastUpdatedDate() bool {
	if o != nil && !isNil(o.LastUpdatedDate) {
		return true
	}

	return false
}

// SetLastUpdatedDate gets a reference to the given string and assigns it to the LastUpdatedDate field.
func (o *ValidateResponseKycResultPrimaryBusinessContact) SetLastUpdatedDate(v string) {
	o.LastUpdatedDate = &v
}

func (o ValidateResponseKycResultPrimaryBusinessContact) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Email) {
		toSerialize["email"] = o.Email
	}
	if !isNil(o.Address) {
		toSerialize["address"] = o.Address
	}
	if !isNil(o.MobileNumber) {
		toSerialize["mobileNumber"] = o.MobileNumber
	}
	if !isNil(o.NatureOfBusinessAtAddress) {
		toSerialize["natureOfBusinessAtAddress"] = o.NatureOfBusinessAtAddress
	}
	if !isNil(o.LastUpdatedDate) {
		toSerialize["lastUpdatedDate"] = o.LastUpdatedDate
	}
	return json.Marshal(toSerialize)
}

type NullableValidateResponseKycResultPrimaryBusinessContact struct {
	value *ValidateResponseKycResultPrimaryBusinessContact
	isSet bool
}

func (v NullableValidateResponseKycResultPrimaryBusinessContact) Get() *ValidateResponseKycResultPrimaryBusinessContact {
	return v.value
}

func (v *NullableValidateResponseKycResultPrimaryBusinessContact) Set(val *ValidateResponseKycResultPrimaryBusinessContact) {
	v.value = val
	v.isSet = true
}

func (v NullableValidateResponseKycResultPrimaryBusinessContact) IsSet() bool {
	return v.isSet
}

func (v *NullableValidateResponseKycResultPrimaryBusinessContact) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableValidateResponseKycResultPrimaryBusinessContact(val *ValidateResponseKycResultPrimaryBusinessContact) *NullableValidateResponseKycResultPrimaryBusinessContact {
	return &NullableValidateResponseKycResultPrimaryBusinessContact{value: val, isSet: true}
}

func (v NullableValidateResponseKycResultPrimaryBusinessContact) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableValidateResponseKycResultPrimaryBusinessContact) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


