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

// ExtractText400ResponseError struct for ExtractText400ResponseError
type ExtractText400ResponseError struct {
	Message *string `json:"message,omitempty"`
	ResponseCode *string `json:"responseCode,omitempty"`
}

// NewExtractText400ResponseError instantiates a new ExtractText400ResponseError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExtractText400ResponseError() *ExtractText400ResponseError {
	this := ExtractText400ResponseError{}
	return &this
}

// NewExtractText400ResponseErrorWithDefaults instantiates a new ExtractText400ResponseError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExtractText400ResponseErrorWithDefaults() *ExtractText400ResponseError {
	this := ExtractText400ResponseError{}
	return &this
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *ExtractText400ResponseError) GetMessage() string {
	if o == nil || isNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExtractText400ResponseError) GetMessageOk() (*string, bool) {
	if o == nil || isNil(o.Message) {
    return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *ExtractText400ResponseError) HasMessage() bool {
	if o != nil && !isNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *ExtractText400ResponseError) SetMessage(v string) {
	o.Message = &v
}

// GetResponseCode returns the ResponseCode field value if set, zero value otherwise.
func (o *ExtractText400ResponseError) GetResponseCode() string {
	if o == nil || isNil(o.ResponseCode) {
		var ret string
		return ret
	}
	return *o.ResponseCode
}

// GetResponseCodeOk returns a tuple with the ResponseCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExtractText400ResponseError) GetResponseCodeOk() (*string, bool) {
	if o == nil || isNil(o.ResponseCode) {
    return nil, false
	}
	return o.ResponseCode, true
}

// HasResponseCode returns a boolean if a field has been set.
func (o *ExtractText400ResponseError) HasResponseCode() bool {
	if o != nil && !isNil(o.ResponseCode) {
		return true
	}

	return false
}

// SetResponseCode gets a reference to the given string and assigns it to the ResponseCode field.
func (o *ExtractText400ResponseError) SetResponseCode(v string) {
	o.ResponseCode = &v
}

func (o ExtractText400ResponseError) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	if !isNil(o.ResponseCode) {
		toSerialize["responseCode"] = o.ResponseCode
	}
	return json.Marshal(toSerialize)
}

type NullableExtractText400ResponseError struct {
	value *ExtractText400ResponseError
	isSet bool
}

func (v NullableExtractText400ResponseError) Get() *ExtractText400ResponseError {
	return v.value
}

func (v *NullableExtractText400ResponseError) Set(val *ExtractText400ResponseError) {
	v.value = val
	v.isSet = true
}

func (v NullableExtractText400ResponseError) IsSet() bool {
	return v.isSet
}

func (v *NullableExtractText400ResponseError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExtractText400ResponseError(val *ExtractText400ResponseError) *NullableExtractText400ResponseError {
	return &NullableExtractText400ResponseError{value: val, isSet: true}
}

func (v NullableExtractText400ResponseError) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExtractText400ResponseError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


