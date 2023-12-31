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

// DigilockerGenerateAccessTokenResponseData struct for DigilockerGenerateAccessTokenResponseData
type DigilockerGenerateAccessTokenResponseData struct {
	ExpiresIn *float32 `json:"expiresIn,omitempty"`
	Name *string `json:"name,omitempty"`
	Dob *string `json:"dob,omitempty"`
	Gender *string `json:"gender,omitempty"`
	Eaadhaar *string `json:"eaadhaar,omitempty"`
}

// NewDigilockerGenerateAccessTokenResponseData instantiates a new DigilockerGenerateAccessTokenResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDigilockerGenerateAccessTokenResponseData() *DigilockerGenerateAccessTokenResponseData {
	this := DigilockerGenerateAccessTokenResponseData{}
	return &this
}

// NewDigilockerGenerateAccessTokenResponseDataWithDefaults instantiates a new DigilockerGenerateAccessTokenResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDigilockerGenerateAccessTokenResponseDataWithDefaults() *DigilockerGenerateAccessTokenResponseData {
	this := DigilockerGenerateAccessTokenResponseData{}
	return &this
}

// GetExpiresIn returns the ExpiresIn field value if set, zero value otherwise.
func (o *DigilockerGenerateAccessTokenResponseData) GetExpiresIn() float32 {
	if o == nil || isNil(o.ExpiresIn) {
		var ret float32
		return ret
	}
	return *o.ExpiresIn
}

// GetExpiresInOk returns a tuple with the ExpiresIn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DigilockerGenerateAccessTokenResponseData) GetExpiresInOk() (*float32, bool) {
	if o == nil || isNil(o.ExpiresIn) {
    return nil, false
	}
	return o.ExpiresIn, true
}

// HasExpiresIn returns a boolean if a field has been set.
func (o *DigilockerGenerateAccessTokenResponseData) HasExpiresIn() bool {
	if o != nil && !isNil(o.ExpiresIn) {
		return true
	}

	return false
}

// SetExpiresIn gets a reference to the given float32 and assigns it to the ExpiresIn field.
func (o *DigilockerGenerateAccessTokenResponseData) SetExpiresIn(v float32) {
	o.ExpiresIn = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *DigilockerGenerateAccessTokenResponseData) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DigilockerGenerateAccessTokenResponseData) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *DigilockerGenerateAccessTokenResponseData) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *DigilockerGenerateAccessTokenResponseData) SetName(v string) {
	o.Name = &v
}

// GetDob returns the Dob field value if set, zero value otherwise.
func (o *DigilockerGenerateAccessTokenResponseData) GetDob() string {
	if o == nil || isNil(o.Dob) {
		var ret string
		return ret
	}
	return *o.Dob
}

// GetDobOk returns a tuple with the Dob field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DigilockerGenerateAccessTokenResponseData) GetDobOk() (*string, bool) {
	if o == nil || isNil(o.Dob) {
    return nil, false
	}
	return o.Dob, true
}

// HasDob returns a boolean if a field has been set.
func (o *DigilockerGenerateAccessTokenResponseData) HasDob() bool {
	if o != nil && !isNil(o.Dob) {
		return true
	}

	return false
}

// SetDob gets a reference to the given string and assigns it to the Dob field.
func (o *DigilockerGenerateAccessTokenResponseData) SetDob(v string) {
	o.Dob = &v
}

// GetGender returns the Gender field value if set, zero value otherwise.
func (o *DigilockerGenerateAccessTokenResponseData) GetGender() string {
	if o == nil || isNil(o.Gender) {
		var ret string
		return ret
	}
	return *o.Gender
}

// GetGenderOk returns a tuple with the Gender field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DigilockerGenerateAccessTokenResponseData) GetGenderOk() (*string, bool) {
	if o == nil || isNil(o.Gender) {
    return nil, false
	}
	return o.Gender, true
}

// HasGender returns a boolean if a field has been set.
func (o *DigilockerGenerateAccessTokenResponseData) HasGender() bool {
	if o != nil && !isNil(o.Gender) {
		return true
	}

	return false
}

// SetGender gets a reference to the given string and assigns it to the Gender field.
func (o *DigilockerGenerateAccessTokenResponseData) SetGender(v string) {
	o.Gender = &v
}

// GetEaadhaar returns the Eaadhaar field value if set, zero value otherwise.
func (o *DigilockerGenerateAccessTokenResponseData) GetEaadhaar() string {
	if o == nil || isNil(o.Eaadhaar) {
		var ret string
		return ret
	}
	return *o.Eaadhaar
}

// GetEaadhaarOk returns a tuple with the Eaadhaar field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DigilockerGenerateAccessTokenResponseData) GetEaadhaarOk() (*string, bool) {
	if o == nil || isNil(o.Eaadhaar) {
    return nil, false
	}
	return o.Eaadhaar, true
}

// HasEaadhaar returns a boolean if a field has been set.
func (o *DigilockerGenerateAccessTokenResponseData) HasEaadhaar() bool {
	if o != nil && !isNil(o.Eaadhaar) {
		return true
	}

	return false
}

// SetEaadhaar gets a reference to the given string and assigns it to the Eaadhaar field.
func (o *DigilockerGenerateAccessTokenResponseData) SetEaadhaar(v string) {
	o.Eaadhaar = &v
}

func (o DigilockerGenerateAccessTokenResponseData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.ExpiresIn) {
		toSerialize["expiresIn"] = o.ExpiresIn
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Dob) {
		toSerialize["dob"] = o.Dob
	}
	if !isNil(o.Gender) {
		toSerialize["gender"] = o.Gender
	}
	if !isNil(o.Eaadhaar) {
		toSerialize["eaadhaar"] = o.Eaadhaar
	}
	return json.Marshal(toSerialize)
}

type NullableDigilockerGenerateAccessTokenResponseData struct {
	value *DigilockerGenerateAccessTokenResponseData
	isSet bool
}

func (v NullableDigilockerGenerateAccessTokenResponseData) Get() *DigilockerGenerateAccessTokenResponseData {
	return v.value
}

func (v *NullableDigilockerGenerateAccessTokenResponseData) Set(val *DigilockerGenerateAccessTokenResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableDigilockerGenerateAccessTokenResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableDigilockerGenerateAccessTokenResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDigilockerGenerateAccessTokenResponseData(val *DigilockerGenerateAccessTokenResponseData) *NullableDigilockerGenerateAccessTokenResponseData {
	return &NullableDigilockerGenerateAccessTokenResponseData{value: val, isSet: true}
}

func (v NullableDigilockerGenerateAccessTokenResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDigilockerGenerateAccessTokenResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


