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

// AadhaarxmlValidateOtpResponseDataAadhaarFile struct for AadhaarxmlValidateOtpResponseDataAadhaarFile
type AadhaarxmlValidateOtpResponseDataAadhaarFile struct {
	AadhaarZip *string `json:"aadhaarZip,omitempty"`
	ShareCode *string `json:"shareCode,omitempty"`
}

// NewAadhaarxmlValidateOtpResponseDataAadhaarFile instantiates a new AadhaarxmlValidateOtpResponseDataAadhaarFile object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAadhaarxmlValidateOtpResponseDataAadhaarFile() *AadhaarxmlValidateOtpResponseDataAadhaarFile {
	this := AadhaarxmlValidateOtpResponseDataAadhaarFile{}
	return &this
}

// NewAadhaarxmlValidateOtpResponseDataAadhaarFileWithDefaults instantiates a new AadhaarxmlValidateOtpResponseDataAadhaarFile object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAadhaarxmlValidateOtpResponseDataAadhaarFileWithDefaults() *AadhaarxmlValidateOtpResponseDataAadhaarFile {
	this := AadhaarxmlValidateOtpResponseDataAadhaarFile{}
	return &this
}

// GetAadhaarZip returns the AadhaarZip field value if set, zero value otherwise.
func (o *AadhaarxmlValidateOtpResponseDataAadhaarFile) GetAadhaarZip() string {
	if o == nil || isNil(o.AadhaarZip) {
		var ret string
		return ret
	}
	return *o.AadhaarZip
}

// GetAadhaarZipOk returns a tuple with the AadhaarZip field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AadhaarxmlValidateOtpResponseDataAadhaarFile) GetAadhaarZipOk() (*string, bool) {
	if o == nil || isNil(o.AadhaarZip) {
    return nil, false
	}
	return o.AadhaarZip, true
}

// HasAadhaarZip returns a boolean if a field has been set.
func (o *AadhaarxmlValidateOtpResponseDataAadhaarFile) HasAadhaarZip() bool {
	if o != nil && !isNil(o.AadhaarZip) {
		return true
	}

	return false
}

// SetAadhaarZip gets a reference to the given string and assigns it to the AadhaarZip field.
func (o *AadhaarxmlValidateOtpResponseDataAadhaarFile) SetAadhaarZip(v string) {
	o.AadhaarZip = &v
}

// GetShareCode returns the ShareCode field value if set, zero value otherwise.
func (o *AadhaarxmlValidateOtpResponseDataAadhaarFile) GetShareCode() string {
	if o == nil || isNil(o.ShareCode) {
		var ret string
		return ret
	}
	return *o.ShareCode
}

// GetShareCodeOk returns a tuple with the ShareCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AadhaarxmlValidateOtpResponseDataAadhaarFile) GetShareCodeOk() (*string, bool) {
	if o == nil || isNil(o.ShareCode) {
    return nil, false
	}
	return o.ShareCode, true
}

// HasShareCode returns a boolean if a field has been set.
func (o *AadhaarxmlValidateOtpResponseDataAadhaarFile) HasShareCode() bool {
	if o != nil && !isNil(o.ShareCode) {
		return true
	}

	return false
}

// SetShareCode gets a reference to the given string and assigns it to the ShareCode field.
func (o *AadhaarxmlValidateOtpResponseDataAadhaarFile) SetShareCode(v string) {
	o.ShareCode = &v
}

func (o AadhaarxmlValidateOtpResponseDataAadhaarFile) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.AadhaarZip) {
		toSerialize["aadhaarZip"] = o.AadhaarZip
	}
	if !isNil(o.ShareCode) {
		toSerialize["shareCode"] = o.ShareCode
	}
	return json.Marshal(toSerialize)
}

type NullableAadhaarxmlValidateOtpResponseDataAadhaarFile struct {
	value *AadhaarxmlValidateOtpResponseDataAadhaarFile
	isSet bool
}

func (v NullableAadhaarxmlValidateOtpResponseDataAadhaarFile) Get() *AadhaarxmlValidateOtpResponseDataAadhaarFile {
	return v.value
}

func (v *NullableAadhaarxmlValidateOtpResponseDataAadhaarFile) Set(val *AadhaarxmlValidateOtpResponseDataAadhaarFile) {
	v.value = val
	v.isSet = true
}

func (v NullableAadhaarxmlValidateOtpResponseDataAadhaarFile) IsSet() bool {
	return v.isSet
}

func (v *NullableAadhaarxmlValidateOtpResponseDataAadhaarFile) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAadhaarxmlValidateOtpResponseDataAadhaarFile(val *AadhaarxmlValidateOtpResponseDataAadhaarFile) *NullableAadhaarxmlValidateOtpResponseDataAadhaarFile {
	return &NullableAadhaarxmlValidateOtpResponseDataAadhaarFile{value: val, isSet: true}
}

func (v NullableAadhaarxmlValidateOtpResponseDataAadhaarFile) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAadhaarxmlValidateOtpResponseDataAadhaarFile) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


