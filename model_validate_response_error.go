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

// ValidateResponseError struct for ValidateResponseError
type ValidateResponseError struct {
	Message *string `json:"message,omitempty"`
	ResponseCode *string `json:"responseCode,omitempty"`
}

// NewValidateResponseError instantiates a new ValidateResponseError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewValidateResponseError() *ValidateResponseError {
	this := ValidateResponseError{}
	return &this
}

// NewValidateResponseErrorWithDefaults instantiates a new ValidateResponseError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewValidateResponseErrorWithDefaults() *ValidateResponseError {
	this := ValidateResponseError{}
	return &this
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *ValidateResponseError) GetMessage() string {
	if o == nil || isNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidateResponseError) GetMessageOk() (*string, bool) {
	if o == nil || isNil(o.Message) {
    return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *ValidateResponseError) HasMessage() bool {
	if o != nil && !isNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *ValidateResponseError) SetMessage(v string) {
	o.Message = &v
}

// GetResponseCode returns the ResponseCode field value if set, zero value otherwise.
func (o *ValidateResponseError) GetResponseCode() string {
	if o == nil || isNil(o.ResponseCode) {
		var ret string
		return ret
	}
	return *o.ResponseCode
}

// GetResponseCodeOk returns a tuple with the ResponseCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidateResponseError) GetResponseCodeOk() (*string, bool) {
	if o == nil || isNil(o.ResponseCode) {
    return nil, false
	}
	return o.ResponseCode, true
}

// HasResponseCode returns a boolean if a field has been set.
func (o *ValidateResponseError) HasResponseCode() bool {
	if o != nil && !isNil(o.ResponseCode) {
		return true
	}

	return false
}

// SetResponseCode gets a reference to the given string and assigns it to the ResponseCode field.
func (o *ValidateResponseError) SetResponseCode(v string) {
	o.ResponseCode = &v
}

func (o ValidateResponseError) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	if !isNil(o.ResponseCode) {
		toSerialize["responseCode"] = o.ResponseCode
	}
	return json.Marshal(toSerialize)
}

type NullableValidateResponseError struct {
	value *ValidateResponseError
	isSet bool
}

func (v NullableValidateResponseError) Get() *ValidateResponseError {
	return v.value
}

func (v *NullableValidateResponseError) Set(val *ValidateResponseError) {
	v.value = val
	v.isSet = true
}

func (v NullableValidateResponseError) IsSet() bool {
	return v.isSet
}

func (v *NullableValidateResponseError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableValidateResponseError(val *ValidateResponseError) *NullableValidateResponseError {
	return &NullableValidateResponseError{value: val, isSet: true}
}

func (v NullableValidateResponseError) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableValidateResponseError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


