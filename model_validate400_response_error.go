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

// Validate400ResponseError struct for Validate400ResponseError
type Validate400ResponseError struct {
	Message *string `json:"message,omitempty"`
	ResponseCode *string `json:"responseCode,omitempty"`
}

// NewValidate400ResponseError instantiates a new Validate400ResponseError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewValidate400ResponseError() *Validate400ResponseError {
	this := Validate400ResponseError{}
	return &this
}

// NewValidate400ResponseErrorWithDefaults instantiates a new Validate400ResponseError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewValidate400ResponseErrorWithDefaults() *Validate400ResponseError {
	this := Validate400ResponseError{}
	return &this
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *Validate400ResponseError) GetMessage() string {
	if o == nil || isNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Validate400ResponseError) GetMessageOk() (*string, bool) {
	if o == nil || isNil(o.Message) {
    return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *Validate400ResponseError) HasMessage() bool {
	if o != nil && !isNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *Validate400ResponseError) SetMessage(v string) {
	o.Message = &v
}

// GetResponseCode returns the ResponseCode field value if set, zero value otherwise.
func (o *Validate400ResponseError) GetResponseCode() string {
	if o == nil || isNil(o.ResponseCode) {
		var ret string
		return ret
	}
	return *o.ResponseCode
}

// GetResponseCodeOk returns a tuple with the ResponseCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Validate400ResponseError) GetResponseCodeOk() (*string, bool) {
	if o == nil || isNil(o.ResponseCode) {
    return nil, false
	}
	return o.ResponseCode, true
}

// HasResponseCode returns a boolean if a field has been set.
func (o *Validate400ResponseError) HasResponseCode() bool {
	if o != nil && !isNil(o.ResponseCode) {
		return true
	}

	return false
}

// SetResponseCode gets a reference to the given string and assigns it to the ResponseCode field.
func (o *Validate400ResponseError) SetResponseCode(v string) {
	o.ResponseCode = &v
}

func (o Validate400ResponseError) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	if !isNil(o.ResponseCode) {
		toSerialize["responseCode"] = o.ResponseCode
	}
	return json.Marshal(toSerialize)
}

type NullableValidate400ResponseError struct {
	value *Validate400ResponseError
	isSet bool
}

func (v NullableValidate400ResponseError) Get() *Validate400ResponseError {
	return v.value
}

func (v *NullableValidate400ResponseError) Set(val *Validate400ResponseError) {
	v.value = val
	v.isSet = true
}

func (v NullableValidate400ResponseError) IsSet() bool {
	return v.isSet
}

func (v *NullableValidate400ResponseError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableValidate400ResponseError(val *Validate400ResponseError) *NullableValidate400ResponseError {
	return &NullableValidate400ResponseError{value: val, isSet: true}
}

func (v NullableValidate400ResponseError) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableValidate400ResponseError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


