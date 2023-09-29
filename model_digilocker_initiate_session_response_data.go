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

// DigilockerInitiateSessionResponseData struct for DigilockerInitiateSessionResponseData
type DigilockerInitiateSessionResponseData struct {
	AuthorizationUrl *string `json:"authorizationUrl,omitempty"`
}

// NewDigilockerInitiateSessionResponseData instantiates a new DigilockerInitiateSessionResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDigilockerInitiateSessionResponseData() *DigilockerInitiateSessionResponseData {
	this := DigilockerInitiateSessionResponseData{}
	return &this
}

// NewDigilockerInitiateSessionResponseDataWithDefaults instantiates a new DigilockerInitiateSessionResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDigilockerInitiateSessionResponseDataWithDefaults() *DigilockerInitiateSessionResponseData {
	this := DigilockerInitiateSessionResponseData{}
	return &this
}

// GetAuthorizationUrl returns the AuthorizationUrl field value if set, zero value otherwise.
func (o *DigilockerInitiateSessionResponseData) GetAuthorizationUrl() string {
	if o == nil || isNil(o.AuthorizationUrl) {
		var ret string
		return ret
	}
	return *o.AuthorizationUrl
}

// GetAuthorizationUrlOk returns a tuple with the AuthorizationUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DigilockerInitiateSessionResponseData) GetAuthorizationUrlOk() (*string, bool) {
	if o == nil || isNil(o.AuthorizationUrl) {
    return nil, false
	}
	return o.AuthorizationUrl, true
}

// HasAuthorizationUrl returns a boolean if a field has been set.
func (o *DigilockerInitiateSessionResponseData) HasAuthorizationUrl() bool {
	if o != nil && !isNil(o.AuthorizationUrl) {
		return true
	}

	return false
}

// SetAuthorizationUrl gets a reference to the given string and assigns it to the AuthorizationUrl field.
func (o *DigilockerInitiateSessionResponseData) SetAuthorizationUrl(v string) {
	o.AuthorizationUrl = &v
}

func (o DigilockerInitiateSessionResponseData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.AuthorizationUrl) {
		toSerialize["authorizationUrl"] = o.AuthorizationUrl
	}
	return json.Marshal(toSerialize)
}

type NullableDigilockerInitiateSessionResponseData struct {
	value *DigilockerInitiateSessionResponseData
	isSet bool
}

func (v NullableDigilockerInitiateSessionResponseData) Get() *DigilockerInitiateSessionResponseData {
	return v.value
}

func (v *NullableDigilockerInitiateSessionResponseData) Set(val *DigilockerInitiateSessionResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableDigilockerInitiateSessionResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableDigilockerInitiateSessionResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDigilockerInitiateSessionResponseData(val *DigilockerInitiateSessionResponseData) *NullableDigilockerInitiateSessionResponseData {
	return &NullableDigilockerInitiateSessionResponseData{value: val, isSet: true}
}

func (v NullableDigilockerInitiateSessionResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDigilockerInitiateSessionResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


