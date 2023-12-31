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

// ValidateResponseKycResultPollingBoothDetails struct for ValidateResponseKycResultPollingBoothDetails
type ValidateResponseKycResultPollingBoothDetails struct {
	LatLong *string `json:"latLong,omitempty"`
	Name *string `json:"name,omitempty"`
	NameVernacular *string `json:"nameVernacular,omitempty"`
	Number *string `json:"number,omitempty"`
}

// NewValidateResponseKycResultPollingBoothDetails instantiates a new ValidateResponseKycResultPollingBoothDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewValidateResponseKycResultPollingBoothDetails() *ValidateResponseKycResultPollingBoothDetails {
	this := ValidateResponseKycResultPollingBoothDetails{}
	return &this
}

// NewValidateResponseKycResultPollingBoothDetailsWithDefaults instantiates a new ValidateResponseKycResultPollingBoothDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewValidateResponseKycResultPollingBoothDetailsWithDefaults() *ValidateResponseKycResultPollingBoothDetails {
	this := ValidateResponseKycResultPollingBoothDetails{}
	return &this
}

// GetLatLong returns the LatLong field value if set, zero value otherwise.
func (o *ValidateResponseKycResultPollingBoothDetails) GetLatLong() string {
	if o == nil || isNil(o.LatLong) {
		var ret string
		return ret
	}
	return *o.LatLong
}

// GetLatLongOk returns a tuple with the LatLong field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidateResponseKycResultPollingBoothDetails) GetLatLongOk() (*string, bool) {
	if o == nil || isNil(o.LatLong) {
    return nil, false
	}
	return o.LatLong, true
}

// HasLatLong returns a boolean if a field has been set.
func (o *ValidateResponseKycResultPollingBoothDetails) HasLatLong() bool {
	if o != nil && !isNil(o.LatLong) {
		return true
	}

	return false
}

// SetLatLong gets a reference to the given string and assigns it to the LatLong field.
func (o *ValidateResponseKycResultPollingBoothDetails) SetLatLong(v string) {
	o.LatLong = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ValidateResponseKycResultPollingBoothDetails) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidateResponseKycResultPollingBoothDetails) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ValidateResponseKycResultPollingBoothDetails) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ValidateResponseKycResultPollingBoothDetails) SetName(v string) {
	o.Name = &v
}

// GetNameVernacular returns the NameVernacular field value if set, zero value otherwise.
func (o *ValidateResponseKycResultPollingBoothDetails) GetNameVernacular() string {
	if o == nil || isNil(o.NameVernacular) {
		var ret string
		return ret
	}
	return *o.NameVernacular
}

// GetNameVernacularOk returns a tuple with the NameVernacular field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidateResponseKycResultPollingBoothDetails) GetNameVernacularOk() (*string, bool) {
	if o == nil || isNil(o.NameVernacular) {
    return nil, false
	}
	return o.NameVernacular, true
}

// HasNameVernacular returns a boolean if a field has been set.
func (o *ValidateResponseKycResultPollingBoothDetails) HasNameVernacular() bool {
	if o != nil && !isNil(o.NameVernacular) {
		return true
	}

	return false
}

// SetNameVernacular gets a reference to the given string and assigns it to the NameVernacular field.
func (o *ValidateResponseKycResultPollingBoothDetails) SetNameVernacular(v string) {
	o.NameVernacular = &v
}

// GetNumber returns the Number field value if set, zero value otherwise.
func (o *ValidateResponseKycResultPollingBoothDetails) GetNumber() string {
	if o == nil || isNil(o.Number) {
		var ret string
		return ret
	}
	return *o.Number
}

// GetNumberOk returns a tuple with the Number field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidateResponseKycResultPollingBoothDetails) GetNumberOk() (*string, bool) {
	if o == nil || isNil(o.Number) {
    return nil, false
	}
	return o.Number, true
}

// HasNumber returns a boolean if a field has been set.
func (o *ValidateResponseKycResultPollingBoothDetails) HasNumber() bool {
	if o != nil && !isNil(o.Number) {
		return true
	}

	return false
}

// SetNumber gets a reference to the given string and assigns it to the Number field.
func (o *ValidateResponseKycResultPollingBoothDetails) SetNumber(v string) {
	o.Number = &v
}

func (o ValidateResponseKycResultPollingBoothDetails) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.LatLong) {
		toSerialize["latLong"] = o.LatLong
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.NameVernacular) {
		toSerialize["nameVernacular"] = o.NameVernacular
	}
	if !isNil(o.Number) {
		toSerialize["number"] = o.Number
	}
	return json.Marshal(toSerialize)
}

type NullableValidateResponseKycResultPollingBoothDetails struct {
	value *ValidateResponseKycResultPollingBoothDetails
	isSet bool
}

func (v NullableValidateResponseKycResultPollingBoothDetails) Get() *ValidateResponseKycResultPollingBoothDetails {
	return v.value
}

func (v *NullableValidateResponseKycResultPollingBoothDetails) Set(val *ValidateResponseKycResultPollingBoothDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableValidateResponseKycResultPollingBoothDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableValidateResponseKycResultPollingBoothDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableValidateResponseKycResultPollingBoothDetails(val *ValidateResponseKycResultPollingBoothDetails) *NullableValidateResponseKycResultPollingBoothDetails {
	return &NullableValidateResponseKycResultPollingBoothDetails{value: val, isSet: true}
}

func (v NullableValidateResponseKycResultPollingBoothDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableValidateResponseKycResultPollingBoothDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


