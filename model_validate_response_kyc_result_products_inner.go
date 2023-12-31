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

// ValidateResponseKycResultProductsInner struct for ValidateResponseKycResultProductsInner
type ValidateResponseKycResultProductsInner struct {
	SlNo *string `json:"slNo,omitempty"`
	FoodProductCategory *string `json:"foodProductCategory,omitempty"`
}

// NewValidateResponseKycResultProductsInner instantiates a new ValidateResponseKycResultProductsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewValidateResponseKycResultProductsInner() *ValidateResponseKycResultProductsInner {
	this := ValidateResponseKycResultProductsInner{}
	return &this
}

// NewValidateResponseKycResultProductsInnerWithDefaults instantiates a new ValidateResponseKycResultProductsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewValidateResponseKycResultProductsInnerWithDefaults() *ValidateResponseKycResultProductsInner {
	this := ValidateResponseKycResultProductsInner{}
	return &this
}

// GetSlNo returns the SlNo field value if set, zero value otherwise.
func (o *ValidateResponseKycResultProductsInner) GetSlNo() string {
	if o == nil || isNil(o.SlNo) {
		var ret string
		return ret
	}
	return *o.SlNo
}

// GetSlNoOk returns a tuple with the SlNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidateResponseKycResultProductsInner) GetSlNoOk() (*string, bool) {
	if o == nil || isNil(o.SlNo) {
    return nil, false
	}
	return o.SlNo, true
}

// HasSlNo returns a boolean if a field has been set.
func (o *ValidateResponseKycResultProductsInner) HasSlNo() bool {
	if o != nil && !isNil(o.SlNo) {
		return true
	}

	return false
}

// SetSlNo gets a reference to the given string and assigns it to the SlNo field.
func (o *ValidateResponseKycResultProductsInner) SetSlNo(v string) {
	o.SlNo = &v
}

// GetFoodProductCategory returns the FoodProductCategory field value if set, zero value otherwise.
func (o *ValidateResponseKycResultProductsInner) GetFoodProductCategory() string {
	if o == nil || isNil(o.FoodProductCategory) {
		var ret string
		return ret
	}
	return *o.FoodProductCategory
}

// GetFoodProductCategoryOk returns a tuple with the FoodProductCategory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidateResponseKycResultProductsInner) GetFoodProductCategoryOk() (*string, bool) {
	if o == nil || isNil(o.FoodProductCategory) {
    return nil, false
	}
	return o.FoodProductCategory, true
}

// HasFoodProductCategory returns a boolean if a field has been set.
func (o *ValidateResponseKycResultProductsInner) HasFoodProductCategory() bool {
	if o != nil && !isNil(o.FoodProductCategory) {
		return true
	}

	return false
}

// SetFoodProductCategory gets a reference to the given string and assigns it to the FoodProductCategory field.
func (o *ValidateResponseKycResultProductsInner) SetFoodProductCategory(v string) {
	o.FoodProductCategory = &v
}

func (o ValidateResponseKycResultProductsInner) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.SlNo) {
		toSerialize["slNo"] = o.SlNo
	}
	if !isNil(o.FoodProductCategory) {
		toSerialize["foodProductCategory"] = o.FoodProductCategory
	}
	return json.Marshal(toSerialize)
}

type NullableValidateResponseKycResultProductsInner struct {
	value *ValidateResponseKycResultProductsInner
	isSet bool
}

func (v NullableValidateResponseKycResultProductsInner) Get() *ValidateResponseKycResultProductsInner {
	return v.value
}

func (v *NullableValidateResponseKycResultProductsInner) Set(val *ValidateResponseKycResultProductsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableValidateResponseKycResultProductsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableValidateResponseKycResultProductsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableValidateResponseKycResultProductsInner(val *ValidateResponseKycResultProductsInner) *NullableValidateResponseKycResultProductsInner {
	return &NullableValidateResponseKycResultProductsInner{value: val, isSet: true}
}

func (v NullableValidateResponseKycResultProductsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableValidateResponseKycResultProductsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


