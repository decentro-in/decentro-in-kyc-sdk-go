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

// ValidateResponseKycResultAddress struct for ValidateResponseKycResultAddress
type ValidateResponseKycResultAddress struct {
	District []string `json:"district,omitempty"`
	State *ValidateResponseKycResultAddressState `json:"state,omitempty"`
	City []string `json:"city,omitempty"`
	Pincode *string `json:"pincode,omitempty"`
	Country []string `json:"country,omitempty"`
	AddressLine *string `json:"addressLine,omitempty"`
	DistrictCode *string `json:"districtCode,omitempty"`
	DistrictName *string `json:"districtName,omitempty"`
	DistrictNameVernacular *string `json:"districtNameVernacular,omitempty"`
	StateCode *string `json:"stateCode,omitempty"`
}

// NewValidateResponseKycResultAddress instantiates a new ValidateResponseKycResultAddress object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewValidateResponseKycResultAddress() *ValidateResponseKycResultAddress {
	this := ValidateResponseKycResultAddress{}
	return &this
}

// NewValidateResponseKycResultAddressWithDefaults instantiates a new ValidateResponseKycResultAddress object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewValidateResponseKycResultAddressWithDefaults() *ValidateResponseKycResultAddress {
	this := ValidateResponseKycResultAddress{}
	return &this
}

// GetDistrict returns the District field value if set, zero value otherwise.
func (o *ValidateResponseKycResultAddress) GetDistrict() []string {
	if o == nil || isNil(o.District) {
		var ret []string
		return ret
	}
	return o.District
}

// GetDistrictOk returns a tuple with the District field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidateResponseKycResultAddress) GetDistrictOk() ([]string, bool) {
	if o == nil || isNil(o.District) {
    return nil, false
	}
	return o.District, true
}

// HasDistrict returns a boolean if a field has been set.
func (o *ValidateResponseKycResultAddress) HasDistrict() bool {
	if o != nil && !isNil(o.District) {
		return true
	}

	return false
}

// SetDistrict gets a reference to the given []string and assigns it to the District field.
func (o *ValidateResponseKycResultAddress) SetDistrict(v []string) {
	o.District = v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *ValidateResponseKycResultAddress) GetState() ValidateResponseKycResultAddressState {
	if o == nil || isNil(o.State) {
		var ret ValidateResponseKycResultAddressState
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidateResponseKycResultAddress) GetStateOk() (*ValidateResponseKycResultAddressState, bool) {
	if o == nil || isNil(o.State) {
    return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *ValidateResponseKycResultAddress) HasState() bool {
	if o != nil && !isNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given ValidateResponseKycResultAddressState and assigns it to the State field.
func (o *ValidateResponseKycResultAddress) SetState(v ValidateResponseKycResultAddressState) {
	o.State = &v
}

// GetCity returns the City field value if set, zero value otherwise.
func (o *ValidateResponseKycResultAddress) GetCity() []string {
	if o == nil || isNil(o.City) {
		var ret []string
		return ret
	}
	return o.City
}

// GetCityOk returns a tuple with the City field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidateResponseKycResultAddress) GetCityOk() ([]string, bool) {
	if o == nil || isNil(o.City) {
    return nil, false
	}
	return o.City, true
}

// HasCity returns a boolean if a field has been set.
func (o *ValidateResponseKycResultAddress) HasCity() bool {
	if o != nil && !isNil(o.City) {
		return true
	}

	return false
}

// SetCity gets a reference to the given []string and assigns it to the City field.
func (o *ValidateResponseKycResultAddress) SetCity(v []string) {
	o.City = v
}

// GetPincode returns the Pincode field value if set, zero value otherwise.
func (o *ValidateResponseKycResultAddress) GetPincode() string {
	if o == nil || isNil(o.Pincode) {
		var ret string
		return ret
	}
	return *o.Pincode
}

// GetPincodeOk returns a tuple with the Pincode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidateResponseKycResultAddress) GetPincodeOk() (*string, bool) {
	if o == nil || isNil(o.Pincode) {
    return nil, false
	}
	return o.Pincode, true
}

// HasPincode returns a boolean if a field has been set.
func (o *ValidateResponseKycResultAddress) HasPincode() bool {
	if o != nil && !isNil(o.Pincode) {
		return true
	}

	return false
}

// SetPincode gets a reference to the given string and assigns it to the Pincode field.
func (o *ValidateResponseKycResultAddress) SetPincode(v string) {
	o.Pincode = &v
}

// GetCountry returns the Country field value if set, zero value otherwise.
func (o *ValidateResponseKycResultAddress) GetCountry() []string {
	if o == nil || isNil(o.Country) {
		var ret []string
		return ret
	}
	return o.Country
}

// GetCountryOk returns a tuple with the Country field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidateResponseKycResultAddress) GetCountryOk() ([]string, bool) {
	if o == nil || isNil(o.Country) {
    return nil, false
	}
	return o.Country, true
}

// HasCountry returns a boolean if a field has been set.
func (o *ValidateResponseKycResultAddress) HasCountry() bool {
	if o != nil && !isNil(o.Country) {
		return true
	}

	return false
}

// SetCountry gets a reference to the given []string and assigns it to the Country field.
func (o *ValidateResponseKycResultAddress) SetCountry(v []string) {
	o.Country = v
}

// GetAddressLine returns the AddressLine field value if set, zero value otherwise.
func (o *ValidateResponseKycResultAddress) GetAddressLine() string {
	if o == nil || isNil(o.AddressLine) {
		var ret string
		return ret
	}
	return *o.AddressLine
}

// GetAddressLineOk returns a tuple with the AddressLine field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidateResponseKycResultAddress) GetAddressLineOk() (*string, bool) {
	if o == nil || isNil(o.AddressLine) {
    return nil, false
	}
	return o.AddressLine, true
}

// HasAddressLine returns a boolean if a field has been set.
func (o *ValidateResponseKycResultAddress) HasAddressLine() bool {
	if o != nil && !isNil(o.AddressLine) {
		return true
	}

	return false
}

// SetAddressLine gets a reference to the given string and assigns it to the AddressLine field.
func (o *ValidateResponseKycResultAddress) SetAddressLine(v string) {
	o.AddressLine = &v
}

// GetDistrictCode returns the DistrictCode field value if set, zero value otherwise.
func (o *ValidateResponseKycResultAddress) GetDistrictCode() string {
	if o == nil || isNil(o.DistrictCode) {
		var ret string
		return ret
	}
	return *o.DistrictCode
}

// GetDistrictCodeOk returns a tuple with the DistrictCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidateResponseKycResultAddress) GetDistrictCodeOk() (*string, bool) {
	if o == nil || isNil(o.DistrictCode) {
    return nil, false
	}
	return o.DistrictCode, true
}

// HasDistrictCode returns a boolean if a field has been set.
func (o *ValidateResponseKycResultAddress) HasDistrictCode() bool {
	if o != nil && !isNil(o.DistrictCode) {
		return true
	}

	return false
}

// SetDistrictCode gets a reference to the given string and assigns it to the DistrictCode field.
func (o *ValidateResponseKycResultAddress) SetDistrictCode(v string) {
	o.DistrictCode = &v
}

// GetDistrictName returns the DistrictName field value if set, zero value otherwise.
func (o *ValidateResponseKycResultAddress) GetDistrictName() string {
	if o == nil || isNil(o.DistrictName) {
		var ret string
		return ret
	}
	return *o.DistrictName
}

// GetDistrictNameOk returns a tuple with the DistrictName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidateResponseKycResultAddress) GetDistrictNameOk() (*string, bool) {
	if o == nil || isNil(o.DistrictName) {
    return nil, false
	}
	return o.DistrictName, true
}

// HasDistrictName returns a boolean if a field has been set.
func (o *ValidateResponseKycResultAddress) HasDistrictName() bool {
	if o != nil && !isNil(o.DistrictName) {
		return true
	}

	return false
}

// SetDistrictName gets a reference to the given string and assigns it to the DistrictName field.
func (o *ValidateResponseKycResultAddress) SetDistrictName(v string) {
	o.DistrictName = &v
}

// GetDistrictNameVernacular returns the DistrictNameVernacular field value if set, zero value otherwise.
func (o *ValidateResponseKycResultAddress) GetDistrictNameVernacular() string {
	if o == nil || isNil(o.DistrictNameVernacular) {
		var ret string
		return ret
	}
	return *o.DistrictNameVernacular
}

// GetDistrictNameVernacularOk returns a tuple with the DistrictNameVernacular field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidateResponseKycResultAddress) GetDistrictNameVernacularOk() (*string, bool) {
	if o == nil || isNil(o.DistrictNameVernacular) {
    return nil, false
	}
	return o.DistrictNameVernacular, true
}

// HasDistrictNameVernacular returns a boolean if a field has been set.
func (o *ValidateResponseKycResultAddress) HasDistrictNameVernacular() bool {
	if o != nil && !isNil(o.DistrictNameVernacular) {
		return true
	}

	return false
}

// SetDistrictNameVernacular gets a reference to the given string and assigns it to the DistrictNameVernacular field.
func (o *ValidateResponseKycResultAddress) SetDistrictNameVernacular(v string) {
	o.DistrictNameVernacular = &v
}

// GetStateCode returns the StateCode field value if set, zero value otherwise.
func (o *ValidateResponseKycResultAddress) GetStateCode() string {
	if o == nil || isNil(o.StateCode) {
		var ret string
		return ret
	}
	return *o.StateCode
}

// GetStateCodeOk returns a tuple with the StateCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidateResponseKycResultAddress) GetStateCodeOk() (*string, bool) {
	if o == nil || isNil(o.StateCode) {
    return nil, false
	}
	return o.StateCode, true
}

// HasStateCode returns a boolean if a field has been set.
func (o *ValidateResponseKycResultAddress) HasStateCode() bool {
	if o != nil && !isNil(o.StateCode) {
		return true
	}

	return false
}

// SetStateCode gets a reference to the given string and assigns it to the StateCode field.
func (o *ValidateResponseKycResultAddress) SetStateCode(v string) {
	o.StateCode = &v
}

func (o ValidateResponseKycResultAddress) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.District) {
		toSerialize["district"] = o.District
	}
	if !isNil(o.State) {
		toSerialize["state"] = o.State
	}
	if !isNil(o.City) {
		toSerialize["city"] = o.City
	}
	if !isNil(o.Pincode) {
		toSerialize["pincode"] = o.Pincode
	}
	if !isNil(o.Country) {
		toSerialize["country"] = o.Country
	}
	if !isNil(o.AddressLine) {
		toSerialize["addressLine"] = o.AddressLine
	}
	if !isNil(o.DistrictCode) {
		toSerialize["districtCode"] = o.DistrictCode
	}
	if !isNil(o.DistrictName) {
		toSerialize["districtName"] = o.DistrictName
	}
	if !isNil(o.DistrictNameVernacular) {
		toSerialize["districtNameVernacular"] = o.DistrictNameVernacular
	}
	if !isNil(o.StateCode) {
		toSerialize["stateCode"] = o.StateCode
	}
	return json.Marshal(toSerialize)
}

type NullableValidateResponseKycResultAddress struct {
	value *ValidateResponseKycResultAddress
	isSet bool
}

func (v NullableValidateResponseKycResultAddress) Get() *ValidateResponseKycResultAddress {
	return v.value
}

func (v *NullableValidateResponseKycResultAddress) Set(val *ValidateResponseKycResultAddress) {
	v.value = val
	v.isSet = true
}

func (v NullableValidateResponseKycResultAddress) IsSet() bool {
	return v.isSet
}

func (v *NullableValidateResponseKycResultAddress) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableValidateResponseKycResultAddress(val *ValidateResponseKycResultAddress) *NullableValidateResponseKycResultAddress {
	return &NullableValidateResponseKycResultAddress{value: val, isSet: true}
}

func (v NullableValidateResponseKycResultAddress) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableValidateResponseKycResultAddress) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


