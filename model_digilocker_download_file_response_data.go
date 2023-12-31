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

// DigilockerDownloadFileResponseData struct for DigilockerDownloadFileResponseData
type DigilockerDownloadFileResponseData struct {
	File *string `json:"file,omitempty"`
}

// NewDigilockerDownloadFileResponseData instantiates a new DigilockerDownloadFileResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDigilockerDownloadFileResponseData() *DigilockerDownloadFileResponseData {
	this := DigilockerDownloadFileResponseData{}
	return &this
}

// NewDigilockerDownloadFileResponseDataWithDefaults instantiates a new DigilockerDownloadFileResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDigilockerDownloadFileResponseDataWithDefaults() *DigilockerDownloadFileResponseData {
	this := DigilockerDownloadFileResponseData{}
	return &this
}

// GetFile returns the File field value if set, zero value otherwise.
func (o *DigilockerDownloadFileResponseData) GetFile() string {
	if o == nil || isNil(o.File) {
		var ret string
		return ret
	}
	return *o.File
}

// GetFileOk returns a tuple with the File field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DigilockerDownloadFileResponseData) GetFileOk() (*string, bool) {
	if o == nil || isNil(o.File) {
    return nil, false
	}
	return o.File, true
}

// HasFile returns a boolean if a field has been set.
func (o *DigilockerDownloadFileResponseData) HasFile() bool {
	if o != nil && !isNil(o.File) {
		return true
	}

	return false
}

// SetFile gets a reference to the given string and assigns it to the File field.
func (o *DigilockerDownloadFileResponseData) SetFile(v string) {
	o.File = &v
}

func (o DigilockerDownloadFileResponseData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.File) {
		toSerialize["file"] = o.File
	}
	return json.Marshal(toSerialize)
}

type NullableDigilockerDownloadFileResponseData struct {
	value *DigilockerDownloadFileResponseData
	isSet bool
}

func (v NullableDigilockerDownloadFileResponseData) Get() *DigilockerDownloadFileResponseData {
	return v.value
}

func (v *NullableDigilockerDownloadFileResponseData) Set(val *DigilockerDownloadFileResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableDigilockerDownloadFileResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableDigilockerDownloadFileResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDigilockerDownloadFileResponseData(val *DigilockerDownloadFileResponseData) *NullableDigilockerDownloadFileResponseData {
	return &NullableDigilockerDownloadFileResponseData{value: val, isSet: true}
}

func (v NullableDigilockerDownloadFileResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDigilockerDownloadFileResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


