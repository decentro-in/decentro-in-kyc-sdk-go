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

// MaskAadhaarUidResponseData struct for MaskAadhaarUidResponseData
type MaskAadhaarUidResponseData struct {
	IsMasked *bool `json:"isMasked,omitempty"`
	MaskedImages []string `json:"maskedImages,omitempty"`
}

// NewMaskAadhaarUidResponseData instantiates a new MaskAadhaarUidResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMaskAadhaarUidResponseData() *MaskAadhaarUidResponseData {
	this := MaskAadhaarUidResponseData{}
	return &this
}

// NewMaskAadhaarUidResponseDataWithDefaults instantiates a new MaskAadhaarUidResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMaskAadhaarUidResponseDataWithDefaults() *MaskAadhaarUidResponseData {
	this := MaskAadhaarUidResponseData{}
	return &this
}

// GetIsMasked returns the IsMasked field value if set, zero value otherwise.
func (o *MaskAadhaarUidResponseData) GetIsMasked() bool {
	if o == nil || isNil(o.IsMasked) {
		var ret bool
		return ret
	}
	return *o.IsMasked
}

// GetIsMaskedOk returns a tuple with the IsMasked field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MaskAadhaarUidResponseData) GetIsMaskedOk() (*bool, bool) {
	if o == nil || isNil(o.IsMasked) {
    return nil, false
	}
	return o.IsMasked, true
}

// HasIsMasked returns a boolean if a field has been set.
func (o *MaskAadhaarUidResponseData) HasIsMasked() bool {
	if o != nil && !isNil(o.IsMasked) {
		return true
	}

	return false
}

// SetIsMasked gets a reference to the given bool and assigns it to the IsMasked field.
func (o *MaskAadhaarUidResponseData) SetIsMasked(v bool) {
	o.IsMasked = &v
}

// GetMaskedImages returns the MaskedImages field value if set, zero value otherwise.
func (o *MaskAadhaarUidResponseData) GetMaskedImages() []string {
	if o == nil || isNil(o.MaskedImages) {
		var ret []string
		return ret
	}
	return o.MaskedImages
}

// GetMaskedImagesOk returns a tuple with the MaskedImages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MaskAadhaarUidResponseData) GetMaskedImagesOk() ([]string, bool) {
	if o == nil || isNil(o.MaskedImages) {
    return nil, false
	}
	return o.MaskedImages, true
}

// HasMaskedImages returns a boolean if a field has been set.
func (o *MaskAadhaarUidResponseData) HasMaskedImages() bool {
	if o != nil && !isNil(o.MaskedImages) {
		return true
	}

	return false
}

// SetMaskedImages gets a reference to the given []string and assigns it to the MaskedImages field.
func (o *MaskAadhaarUidResponseData) SetMaskedImages(v []string) {
	o.MaskedImages = v
}

func (o MaskAadhaarUidResponseData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.IsMasked) {
		toSerialize["isMasked"] = o.IsMasked
	}
	if !isNil(o.MaskedImages) {
		toSerialize["maskedImages"] = o.MaskedImages
	}
	return json.Marshal(toSerialize)
}

type NullableMaskAadhaarUidResponseData struct {
	value *MaskAadhaarUidResponseData
	isSet bool
}

func (v NullableMaskAadhaarUidResponseData) Get() *MaskAadhaarUidResponseData {
	return v.value
}

func (v *NullableMaskAadhaarUidResponseData) Set(val *MaskAadhaarUidResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableMaskAadhaarUidResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableMaskAadhaarUidResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMaskAadhaarUidResponseData(val *MaskAadhaarUidResponseData) *NullableMaskAadhaarUidResponseData {
	return &NullableMaskAadhaarUidResponseData{value: val, isSet: true}
}

func (v NullableMaskAadhaarUidResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMaskAadhaarUidResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


