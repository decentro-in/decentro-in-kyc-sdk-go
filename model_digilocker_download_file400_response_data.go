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

// DigilockerDownloadFile400ResponseData struct for DigilockerDownloadFile400ResponseData
type DigilockerDownloadFile400ResponseData struct {
	File *string `json:"file,omitempty"`
}

// NewDigilockerDownloadFile400ResponseData instantiates a new DigilockerDownloadFile400ResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDigilockerDownloadFile400ResponseData() *DigilockerDownloadFile400ResponseData {
	this := DigilockerDownloadFile400ResponseData{}
	return &this
}

// NewDigilockerDownloadFile400ResponseDataWithDefaults instantiates a new DigilockerDownloadFile400ResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDigilockerDownloadFile400ResponseDataWithDefaults() *DigilockerDownloadFile400ResponseData {
	this := DigilockerDownloadFile400ResponseData{}
	return &this
}

// GetFile returns the File field value if set, zero value otherwise.
func (o *DigilockerDownloadFile400ResponseData) GetFile() string {
	if o == nil || isNil(o.File) {
		var ret string
		return ret
	}
	return *o.File
}

// GetFileOk returns a tuple with the File field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DigilockerDownloadFile400ResponseData) GetFileOk() (*string, bool) {
	if o == nil || isNil(o.File) {
    return nil, false
	}
	return o.File, true
}

// HasFile returns a boolean if a field has been set.
func (o *DigilockerDownloadFile400ResponseData) HasFile() bool {
	if o != nil && !isNil(o.File) {
		return true
	}

	return false
}

// SetFile gets a reference to the given string and assigns it to the File field.
func (o *DigilockerDownloadFile400ResponseData) SetFile(v string) {
	o.File = &v
}

func (o DigilockerDownloadFile400ResponseData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.File) {
		toSerialize["file"] = o.File
	}
	return json.Marshal(toSerialize)
}

type NullableDigilockerDownloadFile400ResponseData struct {
	value *DigilockerDownloadFile400ResponseData
	isSet bool
}

func (v NullableDigilockerDownloadFile400ResponseData) Get() *DigilockerDownloadFile400ResponseData {
	return v.value
}

func (v *NullableDigilockerDownloadFile400ResponseData) Set(val *DigilockerDownloadFile400ResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableDigilockerDownloadFile400ResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableDigilockerDownloadFile400ResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDigilockerDownloadFile400ResponseData(val *DigilockerDownloadFile400ResponseData) *NullableDigilockerDownloadFile400ResponseData {
	return &NullableDigilockerDownloadFile400ResponseData{value: val, isSet: true}
}

func (v NullableDigilockerDownloadFile400ResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDigilockerDownloadFile400ResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


