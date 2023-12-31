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

// AadhaarxmlValidateOtp400ResponseDataAadhaarFile struct for AadhaarxmlValidateOtp400ResponseDataAadhaarFile
type AadhaarxmlValidateOtp400ResponseDataAadhaarFile struct {
	AadhaarZip *string `json:"aadhaarZip,omitempty"`
	ShareCode *string `json:"shareCode,omitempty"`
}

// NewAadhaarxmlValidateOtp400ResponseDataAadhaarFile instantiates a new AadhaarxmlValidateOtp400ResponseDataAadhaarFile object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAadhaarxmlValidateOtp400ResponseDataAadhaarFile() *AadhaarxmlValidateOtp400ResponseDataAadhaarFile {
	this := AadhaarxmlValidateOtp400ResponseDataAadhaarFile{}
	return &this
}

// NewAadhaarxmlValidateOtp400ResponseDataAadhaarFileWithDefaults instantiates a new AadhaarxmlValidateOtp400ResponseDataAadhaarFile object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAadhaarxmlValidateOtp400ResponseDataAadhaarFileWithDefaults() *AadhaarxmlValidateOtp400ResponseDataAadhaarFile {
	this := AadhaarxmlValidateOtp400ResponseDataAadhaarFile{}
	return &this
}

// GetAadhaarZip returns the AadhaarZip field value if set, zero value otherwise.
func (o *AadhaarxmlValidateOtp400ResponseDataAadhaarFile) GetAadhaarZip() string {
	if o == nil || isNil(o.AadhaarZip) {
		var ret string
		return ret
	}
	return *o.AadhaarZip
}

// GetAadhaarZipOk returns a tuple with the AadhaarZip field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AadhaarxmlValidateOtp400ResponseDataAadhaarFile) GetAadhaarZipOk() (*string, bool) {
	if o == nil || isNil(o.AadhaarZip) {
    return nil, false
	}
	return o.AadhaarZip, true
}

// HasAadhaarZip returns a boolean if a field has been set.
func (o *AadhaarxmlValidateOtp400ResponseDataAadhaarFile) HasAadhaarZip() bool {
	if o != nil && !isNil(o.AadhaarZip) {
		return true
	}

	return false
}

// SetAadhaarZip gets a reference to the given string and assigns it to the AadhaarZip field.
func (o *AadhaarxmlValidateOtp400ResponseDataAadhaarFile) SetAadhaarZip(v string) {
	o.AadhaarZip = &v
}

// GetShareCode returns the ShareCode field value if set, zero value otherwise.
func (o *AadhaarxmlValidateOtp400ResponseDataAadhaarFile) GetShareCode() string {
	if o == nil || isNil(o.ShareCode) {
		var ret string
		return ret
	}
	return *o.ShareCode
}

// GetShareCodeOk returns a tuple with the ShareCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AadhaarxmlValidateOtp400ResponseDataAadhaarFile) GetShareCodeOk() (*string, bool) {
	if o == nil || isNil(o.ShareCode) {
    return nil, false
	}
	return o.ShareCode, true
}

// HasShareCode returns a boolean if a field has been set.
func (o *AadhaarxmlValidateOtp400ResponseDataAadhaarFile) HasShareCode() bool {
	if o != nil && !isNil(o.ShareCode) {
		return true
	}

	return false
}

// SetShareCode gets a reference to the given string and assigns it to the ShareCode field.
func (o *AadhaarxmlValidateOtp400ResponseDataAadhaarFile) SetShareCode(v string) {
	o.ShareCode = &v
}

func (o AadhaarxmlValidateOtp400ResponseDataAadhaarFile) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.AadhaarZip) {
		toSerialize["aadhaarZip"] = o.AadhaarZip
	}
	if !isNil(o.ShareCode) {
		toSerialize["shareCode"] = o.ShareCode
	}
	return json.Marshal(toSerialize)
}

type NullableAadhaarxmlValidateOtp400ResponseDataAadhaarFile struct {
	value *AadhaarxmlValidateOtp400ResponseDataAadhaarFile
	isSet bool
}

func (v NullableAadhaarxmlValidateOtp400ResponseDataAadhaarFile) Get() *AadhaarxmlValidateOtp400ResponseDataAadhaarFile {
	return v.value
}

func (v *NullableAadhaarxmlValidateOtp400ResponseDataAadhaarFile) Set(val *AadhaarxmlValidateOtp400ResponseDataAadhaarFile) {
	v.value = val
	v.isSet = true
}

func (v NullableAadhaarxmlValidateOtp400ResponseDataAadhaarFile) IsSet() bool {
	return v.isSet
}

func (v *NullableAadhaarxmlValidateOtp400ResponseDataAadhaarFile) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAadhaarxmlValidateOtp400ResponseDataAadhaarFile(val *AadhaarxmlValidateOtp400ResponseDataAadhaarFile) *NullableAadhaarxmlValidateOtp400ResponseDataAadhaarFile {
	return &NullableAadhaarxmlValidateOtp400ResponseDataAadhaarFile{value: val, isSet: true}
}

func (v NullableAadhaarxmlValidateOtp400ResponseDataAadhaarFile) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAadhaarxmlValidateOtp400ResponseDataAadhaarFile) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


