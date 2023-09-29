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

// ValidateResponseKycResultFilingStatusInner struct for ValidateResponseKycResultFilingStatusInner
type ValidateResponseKycResultFilingStatusInner struct {
	FilingYear *string `json:"filingYear,omitempty"`
	FilingForMonth *string `json:"filingForMonth,omitempty"`
	FilingMethod *string `json:"filingMethod,omitempty"`
	FilingDate *string `json:"filingDate,omitempty"`
	FilingGstType *string `json:"filingGstType,omitempty"`
	FilingAnnualReturn *string `json:"filingAnnualReturn,omitempty"`
	FilingStatus *string `json:"filingStatus,omitempty"`
}

// NewValidateResponseKycResultFilingStatusInner instantiates a new ValidateResponseKycResultFilingStatusInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewValidateResponseKycResultFilingStatusInner() *ValidateResponseKycResultFilingStatusInner {
	this := ValidateResponseKycResultFilingStatusInner{}
	return &this
}

// NewValidateResponseKycResultFilingStatusInnerWithDefaults instantiates a new ValidateResponseKycResultFilingStatusInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewValidateResponseKycResultFilingStatusInnerWithDefaults() *ValidateResponseKycResultFilingStatusInner {
	this := ValidateResponseKycResultFilingStatusInner{}
	return &this
}

// GetFilingYear returns the FilingYear field value if set, zero value otherwise.
func (o *ValidateResponseKycResultFilingStatusInner) GetFilingYear() string {
	if o == nil || isNil(o.FilingYear) {
		var ret string
		return ret
	}
	return *o.FilingYear
}

// GetFilingYearOk returns a tuple with the FilingYear field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidateResponseKycResultFilingStatusInner) GetFilingYearOk() (*string, bool) {
	if o == nil || isNil(o.FilingYear) {
    return nil, false
	}
	return o.FilingYear, true
}

// HasFilingYear returns a boolean if a field has been set.
func (o *ValidateResponseKycResultFilingStatusInner) HasFilingYear() bool {
	if o != nil && !isNil(o.FilingYear) {
		return true
	}

	return false
}

// SetFilingYear gets a reference to the given string and assigns it to the FilingYear field.
func (o *ValidateResponseKycResultFilingStatusInner) SetFilingYear(v string) {
	o.FilingYear = &v
}

// GetFilingForMonth returns the FilingForMonth field value if set, zero value otherwise.
func (o *ValidateResponseKycResultFilingStatusInner) GetFilingForMonth() string {
	if o == nil || isNil(o.FilingForMonth) {
		var ret string
		return ret
	}
	return *o.FilingForMonth
}

// GetFilingForMonthOk returns a tuple with the FilingForMonth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidateResponseKycResultFilingStatusInner) GetFilingForMonthOk() (*string, bool) {
	if o == nil || isNil(o.FilingForMonth) {
    return nil, false
	}
	return o.FilingForMonth, true
}

// HasFilingForMonth returns a boolean if a field has been set.
func (o *ValidateResponseKycResultFilingStatusInner) HasFilingForMonth() bool {
	if o != nil && !isNil(o.FilingForMonth) {
		return true
	}

	return false
}

// SetFilingForMonth gets a reference to the given string and assigns it to the FilingForMonth field.
func (o *ValidateResponseKycResultFilingStatusInner) SetFilingForMonth(v string) {
	o.FilingForMonth = &v
}

// GetFilingMethod returns the FilingMethod field value if set, zero value otherwise.
func (o *ValidateResponseKycResultFilingStatusInner) GetFilingMethod() string {
	if o == nil || isNil(o.FilingMethod) {
		var ret string
		return ret
	}
	return *o.FilingMethod
}

// GetFilingMethodOk returns a tuple with the FilingMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidateResponseKycResultFilingStatusInner) GetFilingMethodOk() (*string, bool) {
	if o == nil || isNil(o.FilingMethod) {
    return nil, false
	}
	return o.FilingMethod, true
}

// HasFilingMethod returns a boolean if a field has been set.
func (o *ValidateResponseKycResultFilingStatusInner) HasFilingMethod() bool {
	if o != nil && !isNil(o.FilingMethod) {
		return true
	}

	return false
}

// SetFilingMethod gets a reference to the given string and assigns it to the FilingMethod field.
func (o *ValidateResponseKycResultFilingStatusInner) SetFilingMethod(v string) {
	o.FilingMethod = &v
}

// GetFilingDate returns the FilingDate field value if set, zero value otherwise.
func (o *ValidateResponseKycResultFilingStatusInner) GetFilingDate() string {
	if o == nil || isNil(o.FilingDate) {
		var ret string
		return ret
	}
	return *o.FilingDate
}

// GetFilingDateOk returns a tuple with the FilingDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidateResponseKycResultFilingStatusInner) GetFilingDateOk() (*string, bool) {
	if o == nil || isNil(o.FilingDate) {
    return nil, false
	}
	return o.FilingDate, true
}

// HasFilingDate returns a boolean if a field has been set.
func (o *ValidateResponseKycResultFilingStatusInner) HasFilingDate() bool {
	if o != nil && !isNil(o.FilingDate) {
		return true
	}

	return false
}

// SetFilingDate gets a reference to the given string and assigns it to the FilingDate field.
func (o *ValidateResponseKycResultFilingStatusInner) SetFilingDate(v string) {
	o.FilingDate = &v
}

// GetFilingGstType returns the FilingGstType field value if set, zero value otherwise.
func (o *ValidateResponseKycResultFilingStatusInner) GetFilingGstType() string {
	if o == nil || isNil(o.FilingGstType) {
		var ret string
		return ret
	}
	return *o.FilingGstType
}

// GetFilingGstTypeOk returns a tuple with the FilingGstType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidateResponseKycResultFilingStatusInner) GetFilingGstTypeOk() (*string, bool) {
	if o == nil || isNil(o.FilingGstType) {
    return nil, false
	}
	return o.FilingGstType, true
}

// HasFilingGstType returns a boolean if a field has been set.
func (o *ValidateResponseKycResultFilingStatusInner) HasFilingGstType() bool {
	if o != nil && !isNil(o.FilingGstType) {
		return true
	}

	return false
}

// SetFilingGstType gets a reference to the given string and assigns it to the FilingGstType field.
func (o *ValidateResponseKycResultFilingStatusInner) SetFilingGstType(v string) {
	o.FilingGstType = &v
}

// GetFilingAnnualReturn returns the FilingAnnualReturn field value if set, zero value otherwise.
func (o *ValidateResponseKycResultFilingStatusInner) GetFilingAnnualReturn() string {
	if o == nil || isNil(o.FilingAnnualReturn) {
		var ret string
		return ret
	}
	return *o.FilingAnnualReturn
}

// GetFilingAnnualReturnOk returns a tuple with the FilingAnnualReturn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidateResponseKycResultFilingStatusInner) GetFilingAnnualReturnOk() (*string, bool) {
	if o == nil || isNil(o.FilingAnnualReturn) {
    return nil, false
	}
	return o.FilingAnnualReturn, true
}

// HasFilingAnnualReturn returns a boolean if a field has been set.
func (o *ValidateResponseKycResultFilingStatusInner) HasFilingAnnualReturn() bool {
	if o != nil && !isNil(o.FilingAnnualReturn) {
		return true
	}

	return false
}

// SetFilingAnnualReturn gets a reference to the given string and assigns it to the FilingAnnualReturn field.
func (o *ValidateResponseKycResultFilingStatusInner) SetFilingAnnualReturn(v string) {
	o.FilingAnnualReturn = &v
}

// GetFilingStatus returns the FilingStatus field value if set, zero value otherwise.
func (o *ValidateResponseKycResultFilingStatusInner) GetFilingStatus() string {
	if o == nil || isNil(o.FilingStatus) {
		var ret string
		return ret
	}
	return *o.FilingStatus
}

// GetFilingStatusOk returns a tuple with the FilingStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidateResponseKycResultFilingStatusInner) GetFilingStatusOk() (*string, bool) {
	if o == nil || isNil(o.FilingStatus) {
    return nil, false
	}
	return o.FilingStatus, true
}

// HasFilingStatus returns a boolean if a field has been set.
func (o *ValidateResponseKycResultFilingStatusInner) HasFilingStatus() bool {
	if o != nil && !isNil(o.FilingStatus) {
		return true
	}

	return false
}

// SetFilingStatus gets a reference to the given string and assigns it to the FilingStatus field.
func (o *ValidateResponseKycResultFilingStatusInner) SetFilingStatus(v string) {
	o.FilingStatus = &v
}

func (o ValidateResponseKycResultFilingStatusInner) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.FilingYear) {
		toSerialize["filingYear"] = o.FilingYear
	}
	if !isNil(o.FilingForMonth) {
		toSerialize["filingForMonth"] = o.FilingForMonth
	}
	if !isNil(o.FilingMethod) {
		toSerialize["filingMethod"] = o.FilingMethod
	}
	if !isNil(o.FilingDate) {
		toSerialize["filingDate"] = o.FilingDate
	}
	if !isNil(o.FilingGstType) {
		toSerialize["filingGstType"] = o.FilingGstType
	}
	if !isNil(o.FilingAnnualReturn) {
		toSerialize["filingAnnualReturn"] = o.FilingAnnualReturn
	}
	if !isNil(o.FilingStatus) {
		toSerialize["filingStatus"] = o.FilingStatus
	}
	return json.Marshal(toSerialize)
}

type NullableValidateResponseKycResultFilingStatusInner struct {
	value *ValidateResponseKycResultFilingStatusInner
	isSet bool
}

func (v NullableValidateResponseKycResultFilingStatusInner) Get() *ValidateResponseKycResultFilingStatusInner {
	return v.value
}

func (v *NullableValidateResponseKycResultFilingStatusInner) Set(val *ValidateResponseKycResultFilingStatusInner) {
	v.value = val
	v.isSet = true
}

func (v NullableValidateResponseKycResultFilingStatusInner) IsSet() bool {
	return v.isSet
}

func (v *NullableValidateResponseKycResultFilingStatusInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableValidateResponseKycResultFilingStatusInner(val *ValidateResponseKycResultFilingStatusInner) *NullableValidateResponseKycResultFilingStatusInner {
	return &NullableValidateResponseKycResultFilingStatusInner{value: val, isSet: true}
}

func (v NullableValidateResponseKycResultFilingStatusInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableValidateResponseKycResultFilingStatusInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


