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

// ExtractTextResponseOcrResult struct for ExtractTextResponseOcrResult
type ExtractTextResponseOcrResult struct {
	CardNo *string `json:"cardNo,omitempty"`
	DateInfo *string `json:"dateInfo,omitempty"`
	DateType *string `json:"dateType,omitempty"`
	FatherName *string `json:"fatherName,omitempty"`
	Name *string `json:"name,omitempty"`
	Gender *string `json:"gender,omitempty"`
	Vid *string `json:"vid,omitempty"`
	Address *string `json:"address,omitempty"`
	SonOf *string `json:"sonOf,omitempty"`
	HusbandOf *string `json:"husbandOf,omitempty"`
}

// NewExtractTextResponseOcrResult instantiates a new ExtractTextResponseOcrResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExtractTextResponseOcrResult() *ExtractTextResponseOcrResult {
	this := ExtractTextResponseOcrResult{}
	return &this
}

// NewExtractTextResponseOcrResultWithDefaults instantiates a new ExtractTextResponseOcrResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExtractTextResponseOcrResultWithDefaults() *ExtractTextResponseOcrResult {
	this := ExtractTextResponseOcrResult{}
	return &this
}

// GetCardNo returns the CardNo field value if set, zero value otherwise.
func (o *ExtractTextResponseOcrResult) GetCardNo() string {
	if o == nil || isNil(o.CardNo) {
		var ret string
		return ret
	}
	return *o.CardNo
}

// GetCardNoOk returns a tuple with the CardNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExtractTextResponseOcrResult) GetCardNoOk() (*string, bool) {
	if o == nil || isNil(o.CardNo) {
    return nil, false
	}
	return o.CardNo, true
}

// HasCardNo returns a boolean if a field has been set.
func (o *ExtractTextResponseOcrResult) HasCardNo() bool {
	if o != nil && !isNil(o.CardNo) {
		return true
	}

	return false
}

// SetCardNo gets a reference to the given string and assigns it to the CardNo field.
func (o *ExtractTextResponseOcrResult) SetCardNo(v string) {
	o.CardNo = &v
}

// GetDateInfo returns the DateInfo field value if set, zero value otherwise.
func (o *ExtractTextResponseOcrResult) GetDateInfo() string {
	if o == nil || isNil(o.DateInfo) {
		var ret string
		return ret
	}
	return *o.DateInfo
}

// GetDateInfoOk returns a tuple with the DateInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExtractTextResponseOcrResult) GetDateInfoOk() (*string, bool) {
	if o == nil || isNil(o.DateInfo) {
    return nil, false
	}
	return o.DateInfo, true
}

// HasDateInfo returns a boolean if a field has been set.
func (o *ExtractTextResponseOcrResult) HasDateInfo() bool {
	if o != nil && !isNil(o.DateInfo) {
		return true
	}

	return false
}

// SetDateInfo gets a reference to the given string and assigns it to the DateInfo field.
func (o *ExtractTextResponseOcrResult) SetDateInfo(v string) {
	o.DateInfo = &v
}

// GetDateType returns the DateType field value if set, zero value otherwise.
func (o *ExtractTextResponseOcrResult) GetDateType() string {
	if o == nil || isNil(o.DateType) {
		var ret string
		return ret
	}
	return *o.DateType
}

// GetDateTypeOk returns a tuple with the DateType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExtractTextResponseOcrResult) GetDateTypeOk() (*string, bool) {
	if o == nil || isNil(o.DateType) {
    return nil, false
	}
	return o.DateType, true
}

// HasDateType returns a boolean if a field has been set.
func (o *ExtractTextResponseOcrResult) HasDateType() bool {
	if o != nil && !isNil(o.DateType) {
		return true
	}

	return false
}

// SetDateType gets a reference to the given string and assigns it to the DateType field.
func (o *ExtractTextResponseOcrResult) SetDateType(v string) {
	o.DateType = &v
}

// GetFatherName returns the FatherName field value if set, zero value otherwise.
func (o *ExtractTextResponseOcrResult) GetFatherName() string {
	if o == nil || isNil(o.FatherName) {
		var ret string
		return ret
	}
	return *o.FatherName
}

// GetFatherNameOk returns a tuple with the FatherName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExtractTextResponseOcrResult) GetFatherNameOk() (*string, bool) {
	if o == nil || isNil(o.FatherName) {
    return nil, false
	}
	return o.FatherName, true
}

// HasFatherName returns a boolean if a field has been set.
func (o *ExtractTextResponseOcrResult) HasFatherName() bool {
	if o != nil && !isNil(o.FatherName) {
		return true
	}

	return false
}

// SetFatherName gets a reference to the given string and assigns it to the FatherName field.
func (o *ExtractTextResponseOcrResult) SetFatherName(v string) {
	o.FatherName = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ExtractTextResponseOcrResult) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExtractTextResponseOcrResult) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ExtractTextResponseOcrResult) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ExtractTextResponseOcrResult) SetName(v string) {
	o.Name = &v
}

// GetGender returns the Gender field value if set, zero value otherwise.
func (o *ExtractTextResponseOcrResult) GetGender() string {
	if o == nil || isNil(o.Gender) {
		var ret string
		return ret
	}
	return *o.Gender
}

// GetGenderOk returns a tuple with the Gender field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExtractTextResponseOcrResult) GetGenderOk() (*string, bool) {
	if o == nil || isNil(o.Gender) {
    return nil, false
	}
	return o.Gender, true
}

// HasGender returns a boolean if a field has been set.
func (o *ExtractTextResponseOcrResult) HasGender() bool {
	if o != nil && !isNil(o.Gender) {
		return true
	}

	return false
}

// SetGender gets a reference to the given string and assigns it to the Gender field.
func (o *ExtractTextResponseOcrResult) SetGender(v string) {
	o.Gender = &v
}

// GetVid returns the Vid field value if set, zero value otherwise.
func (o *ExtractTextResponseOcrResult) GetVid() string {
	if o == nil || isNil(o.Vid) {
		var ret string
		return ret
	}
	return *o.Vid
}

// GetVidOk returns a tuple with the Vid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExtractTextResponseOcrResult) GetVidOk() (*string, bool) {
	if o == nil || isNil(o.Vid) {
    return nil, false
	}
	return o.Vid, true
}

// HasVid returns a boolean if a field has been set.
func (o *ExtractTextResponseOcrResult) HasVid() bool {
	if o != nil && !isNil(o.Vid) {
		return true
	}

	return false
}

// SetVid gets a reference to the given string and assigns it to the Vid field.
func (o *ExtractTextResponseOcrResult) SetVid(v string) {
	o.Vid = &v
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *ExtractTextResponseOcrResult) GetAddress() string {
	if o == nil || isNil(o.Address) {
		var ret string
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExtractTextResponseOcrResult) GetAddressOk() (*string, bool) {
	if o == nil || isNil(o.Address) {
    return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *ExtractTextResponseOcrResult) HasAddress() bool {
	if o != nil && !isNil(o.Address) {
		return true
	}

	return false
}

// SetAddress gets a reference to the given string and assigns it to the Address field.
func (o *ExtractTextResponseOcrResult) SetAddress(v string) {
	o.Address = &v
}

// GetSonOf returns the SonOf field value if set, zero value otherwise.
func (o *ExtractTextResponseOcrResult) GetSonOf() string {
	if o == nil || isNil(o.SonOf) {
		var ret string
		return ret
	}
	return *o.SonOf
}

// GetSonOfOk returns a tuple with the SonOf field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExtractTextResponseOcrResult) GetSonOfOk() (*string, bool) {
	if o == nil || isNil(o.SonOf) {
    return nil, false
	}
	return o.SonOf, true
}

// HasSonOf returns a boolean if a field has been set.
func (o *ExtractTextResponseOcrResult) HasSonOf() bool {
	if o != nil && !isNil(o.SonOf) {
		return true
	}

	return false
}

// SetSonOf gets a reference to the given string and assigns it to the SonOf field.
func (o *ExtractTextResponseOcrResult) SetSonOf(v string) {
	o.SonOf = &v
}

// GetHusbandOf returns the HusbandOf field value if set, zero value otherwise.
func (o *ExtractTextResponseOcrResult) GetHusbandOf() string {
	if o == nil || isNil(o.HusbandOf) {
		var ret string
		return ret
	}
	return *o.HusbandOf
}

// GetHusbandOfOk returns a tuple with the HusbandOf field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExtractTextResponseOcrResult) GetHusbandOfOk() (*string, bool) {
	if o == nil || isNil(o.HusbandOf) {
    return nil, false
	}
	return o.HusbandOf, true
}

// HasHusbandOf returns a boolean if a field has been set.
func (o *ExtractTextResponseOcrResult) HasHusbandOf() bool {
	if o != nil && !isNil(o.HusbandOf) {
		return true
	}

	return false
}

// SetHusbandOf gets a reference to the given string and assigns it to the HusbandOf field.
func (o *ExtractTextResponseOcrResult) SetHusbandOf(v string) {
	o.HusbandOf = &v
}

func (o ExtractTextResponseOcrResult) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.CardNo) {
		toSerialize["cardNo"] = o.CardNo
	}
	if !isNil(o.DateInfo) {
		toSerialize["dateInfo"] = o.DateInfo
	}
	if !isNil(o.DateType) {
		toSerialize["dateType"] = o.DateType
	}
	if !isNil(o.FatherName) {
		toSerialize["fatherName"] = o.FatherName
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Gender) {
		toSerialize["gender"] = o.Gender
	}
	if !isNil(o.Vid) {
		toSerialize["vid"] = o.Vid
	}
	if !isNil(o.Address) {
		toSerialize["address"] = o.Address
	}
	if !isNil(o.SonOf) {
		toSerialize["sonOf"] = o.SonOf
	}
	if !isNil(o.HusbandOf) {
		toSerialize["husbandOf"] = o.HusbandOf
	}
	return json.Marshal(toSerialize)
}

type NullableExtractTextResponseOcrResult struct {
	value *ExtractTextResponseOcrResult
	isSet bool
}

func (v NullableExtractTextResponseOcrResult) Get() *ExtractTextResponseOcrResult {
	return v.value
}

func (v *NullableExtractTextResponseOcrResult) Set(val *ExtractTextResponseOcrResult) {
	v.value = val
	v.isSet = true
}

func (v NullableExtractTextResponseOcrResult) IsSet() bool {
	return v.isSet
}

func (v *NullableExtractTextResponseOcrResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExtractTextResponseOcrResult(val *ExtractTextResponseOcrResult) *NullableExtractTextResponseOcrResult {
	return &NullableExtractTextResponseOcrResult{value: val, isSet: true}
}

func (v NullableExtractTextResponseOcrResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExtractTextResponseOcrResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


