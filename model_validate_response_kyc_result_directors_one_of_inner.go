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

// ValidateResponseKycResultDirectorsOneOfInner struct for ValidateResponseKycResultDirectorsOneOfInner
type ValidateResponseKycResultDirectorsOneOfInner struct {
	EndDate *string `json:"endDate,omitempty"`
	SurrenderedDin *string `json:"surrenderedDin,omitempty"`
	DinOrPan *string `json:"dinOrPan,omitempty"`
	BeginDate *string `json:"beginDate,omitempty"`
	Name *string `json:"name,omitempty"`
}

// NewValidateResponseKycResultDirectorsOneOfInner instantiates a new ValidateResponseKycResultDirectorsOneOfInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewValidateResponseKycResultDirectorsOneOfInner() *ValidateResponseKycResultDirectorsOneOfInner {
	this := ValidateResponseKycResultDirectorsOneOfInner{}
	return &this
}

// NewValidateResponseKycResultDirectorsOneOfInnerWithDefaults instantiates a new ValidateResponseKycResultDirectorsOneOfInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewValidateResponseKycResultDirectorsOneOfInnerWithDefaults() *ValidateResponseKycResultDirectorsOneOfInner {
	this := ValidateResponseKycResultDirectorsOneOfInner{}
	return &this
}

// GetEndDate returns the EndDate field value if set, zero value otherwise.
func (o *ValidateResponseKycResultDirectorsOneOfInner) GetEndDate() string {
	if o == nil || isNil(o.EndDate) {
		var ret string
		return ret
	}
	return *o.EndDate
}

// GetEndDateOk returns a tuple with the EndDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidateResponseKycResultDirectorsOneOfInner) GetEndDateOk() (*string, bool) {
	if o == nil || isNil(o.EndDate) {
    return nil, false
	}
	return o.EndDate, true
}

// HasEndDate returns a boolean if a field has been set.
func (o *ValidateResponseKycResultDirectorsOneOfInner) HasEndDate() bool {
	if o != nil && !isNil(o.EndDate) {
		return true
	}

	return false
}

// SetEndDate gets a reference to the given string and assigns it to the EndDate field.
func (o *ValidateResponseKycResultDirectorsOneOfInner) SetEndDate(v string) {
	o.EndDate = &v
}

// GetSurrenderedDin returns the SurrenderedDin field value if set, zero value otherwise.
func (o *ValidateResponseKycResultDirectorsOneOfInner) GetSurrenderedDin() string {
	if o == nil || isNil(o.SurrenderedDin) {
		var ret string
		return ret
	}
	return *o.SurrenderedDin
}

// GetSurrenderedDinOk returns a tuple with the SurrenderedDin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidateResponseKycResultDirectorsOneOfInner) GetSurrenderedDinOk() (*string, bool) {
	if o == nil || isNil(o.SurrenderedDin) {
    return nil, false
	}
	return o.SurrenderedDin, true
}

// HasSurrenderedDin returns a boolean if a field has been set.
func (o *ValidateResponseKycResultDirectorsOneOfInner) HasSurrenderedDin() bool {
	if o != nil && !isNil(o.SurrenderedDin) {
		return true
	}

	return false
}

// SetSurrenderedDin gets a reference to the given string and assigns it to the SurrenderedDin field.
func (o *ValidateResponseKycResultDirectorsOneOfInner) SetSurrenderedDin(v string) {
	o.SurrenderedDin = &v
}

// GetDinOrPan returns the DinOrPan field value if set, zero value otherwise.
func (o *ValidateResponseKycResultDirectorsOneOfInner) GetDinOrPan() string {
	if o == nil || isNil(o.DinOrPan) {
		var ret string
		return ret
	}
	return *o.DinOrPan
}

// GetDinOrPanOk returns a tuple with the DinOrPan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidateResponseKycResultDirectorsOneOfInner) GetDinOrPanOk() (*string, bool) {
	if o == nil || isNil(o.DinOrPan) {
    return nil, false
	}
	return o.DinOrPan, true
}

// HasDinOrPan returns a boolean if a field has been set.
func (o *ValidateResponseKycResultDirectorsOneOfInner) HasDinOrPan() bool {
	if o != nil && !isNil(o.DinOrPan) {
		return true
	}

	return false
}

// SetDinOrPan gets a reference to the given string and assigns it to the DinOrPan field.
func (o *ValidateResponseKycResultDirectorsOneOfInner) SetDinOrPan(v string) {
	o.DinOrPan = &v
}

// GetBeginDate returns the BeginDate field value if set, zero value otherwise.
func (o *ValidateResponseKycResultDirectorsOneOfInner) GetBeginDate() string {
	if o == nil || isNil(o.BeginDate) {
		var ret string
		return ret
	}
	return *o.BeginDate
}

// GetBeginDateOk returns a tuple with the BeginDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidateResponseKycResultDirectorsOneOfInner) GetBeginDateOk() (*string, bool) {
	if o == nil || isNil(o.BeginDate) {
    return nil, false
	}
	return o.BeginDate, true
}

// HasBeginDate returns a boolean if a field has been set.
func (o *ValidateResponseKycResultDirectorsOneOfInner) HasBeginDate() bool {
	if o != nil && !isNil(o.BeginDate) {
		return true
	}

	return false
}

// SetBeginDate gets a reference to the given string and assigns it to the BeginDate field.
func (o *ValidateResponseKycResultDirectorsOneOfInner) SetBeginDate(v string) {
	o.BeginDate = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ValidateResponseKycResultDirectorsOneOfInner) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidateResponseKycResultDirectorsOneOfInner) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ValidateResponseKycResultDirectorsOneOfInner) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ValidateResponseKycResultDirectorsOneOfInner) SetName(v string) {
	o.Name = &v
}

func (o ValidateResponseKycResultDirectorsOneOfInner) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.EndDate) {
		toSerialize["endDate"] = o.EndDate
	}
	if !isNil(o.SurrenderedDin) {
		toSerialize["surrenderedDin"] = o.SurrenderedDin
	}
	if !isNil(o.DinOrPan) {
		toSerialize["dinOrPan"] = o.DinOrPan
	}
	if !isNil(o.BeginDate) {
		toSerialize["beginDate"] = o.BeginDate
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

type NullableValidateResponseKycResultDirectorsOneOfInner struct {
	value *ValidateResponseKycResultDirectorsOneOfInner
	isSet bool
}

func (v NullableValidateResponseKycResultDirectorsOneOfInner) Get() *ValidateResponseKycResultDirectorsOneOfInner {
	return v.value
}

func (v *NullableValidateResponseKycResultDirectorsOneOfInner) Set(val *ValidateResponseKycResultDirectorsOneOfInner) {
	v.value = val
	v.isSet = true
}

func (v NullableValidateResponseKycResultDirectorsOneOfInner) IsSet() bool {
	return v.isSet
}

func (v *NullableValidateResponseKycResultDirectorsOneOfInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableValidateResponseKycResultDirectorsOneOfInner(val *ValidateResponseKycResultDirectorsOneOfInner) *NullableValidateResponseKycResultDirectorsOneOfInner {
	return &NullableValidateResponseKycResultDirectorsOneOfInner{value: val, isSet: true}
}

func (v NullableValidateResponseKycResultDirectorsOneOfInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableValidateResponseKycResultDirectorsOneOfInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


