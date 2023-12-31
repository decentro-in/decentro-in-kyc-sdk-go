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

// ValidateResponseKycResultAllClassOfVehicleInner struct for ValidateResponseKycResultAllClassOfVehicleInner
type ValidateResponseKycResultAllClassOfVehicleInner struct {
	Cov *string `json:"cov,omitempty"`
	ExpiryDate *string `json:"expiryDate,omitempty"`
	IssueDate *string `json:"issueDate,omitempty"`
	CovCategory *string `json:"covCategory,omitempty"`
}

// NewValidateResponseKycResultAllClassOfVehicleInner instantiates a new ValidateResponseKycResultAllClassOfVehicleInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewValidateResponseKycResultAllClassOfVehicleInner() *ValidateResponseKycResultAllClassOfVehicleInner {
	this := ValidateResponseKycResultAllClassOfVehicleInner{}
	return &this
}

// NewValidateResponseKycResultAllClassOfVehicleInnerWithDefaults instantiates a new ValidateResponseKycResultAllClassOfVehicleInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewValidateResponseKycResultAllClassOfVehicleInnerWithDefaults() *ValidateResponseKycResultAllClassOfVehicleInner {
	this := ValidateResponseKycResultAllClassOfVehicleInner{}
	return &this
}

// GetCov returns the Cov field value if set, zero value otherwise.
func (o *ValidateResponseKycResultAllClassOfVehicleInner) GetCov() string {
	if o == nil || isNil(o.Cov) {
		var ret string
		return ret
	}
	return *o.Cov
}

// GetCovOk returns a tuple with the Cov field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidateResponseKycResultAllClassOfVehicleInner) GetCovOk() (*string, bool) {
	if o == nil || isNil(o.Cov) {
    return nil, false
	}
	return o.Cov, true
}

// HasCov returns a boolean if a field has been set.
func (o *ValidateResponseKycResultAllClassOfVehicleInner) HasCov() bool {
	if o != nil && !isNil(o.Cov) {
		return true
	}

	return false
}

// SetCov gets a reference to the given string and assigns it to the Cov field.
func (o *ValidateResponseKycResultAllClassOfVehicleInner) SetCov(v string) {
	o.Cov = &v
}

// GetExpiryDate returns the ExpiryDate field value if set, zero value otherwise.
func (o *ValidateResponseKycResultAllClassOfVehicleInner) GetExpiryDate() string {
	if o == nil || isNil(o.ExpiryDate) {
		var ret string
		return ret
	}
	return *o.ExpiryDate
}

// GetExpiryDateOk returns a tuple with the ExpiryDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidateResponseKycResultAllClassOfVehicleInner) GetExpiryDateOk() (*string, bool) {
	if o == nil || isNil(o.ExpiryDate) {
    return nil, false
	}
	return o.ExpiryDate, true
}

// HasExpiryDate returns a boolean if a field has been set.
func (o *ValidateResponseKycResultAllClassOfVehicleInner) HasExpiryDate() bool {
	if o != nil && !isNil(o.ExpiryDate) {
		return true
	}

	return false
}

// SetExpiryDate gets a reference to the given string and assigns it to the ExpiryDate field.
func (o *ValidateResponseKycResultAllClassOfVehicleInner) SetExpiryDate(v string) {
	o.ExpiryDate = &v
}

// GetIssueDate returns the IssueDate field value if set, zero value otherwise.
func (o *ValidateResponseKycResultAllClassOfVehicleInner) GetIssueDate() string {
	if o == nil || isNil(o.IssueDate) {
		var ret string
		return ret
	}
	return *o.IssueDate
}

// GetIssueDateOk returns a tuple with the IssueDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidateResponseKycResultAllClassOfVehicleInner) GetIssueDateOk() (*string, bool) {
	if o == nil || isNil(o.IssueDate) {
    return nil, false
	}
	return o.IssueDate, true
}

// HasIssueDate returns a boolean if a field has been set.
func (o *ValidateResponseKycResultAllClassOfVehicleInner) HasIssueDate() bool {
	if o != nil && !isNil(o.IssueDate) {
		return true
	}

	return false
}

// SetIssueDate gets a reference to the given string and assigns it to the IssueDate field.
func (o *ValidateResponseKycResultAllClassOfVehicleInner) SetIssueDate(v string) {
	o.IssueDate = &v
}

// GetCovCategory returns the CovCategory field value if set, zero value otherwise.
func (o *ValidateResponseKycResultAllClassOfVehicleInner) GetCovCategory() string {
	if o == nil || isNil(o.CovCategory) {
		var ret string
		return ret
	}
	return *o.CovCategory
}

// GetCovCategoryOk returns a tuple with the CovCategory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidateResponseKycResultAllClassOfVehicleInner) GetCovCategoryOk() (*string, bool) {
	if o == nil || isNil(o.CovCategory) {
    return nil, false
	}
	return o.CovCategory, true
}

// HasCovCategory returns a boolean if a field has been set.
func (o *ValidateResponseKycResultAllClassOfVehicleInner) HasCovCategory() bool {
	if o != nil && !isNil(o.CovCategory) {
		return true
	}

	return false
}

// SetCovCategory gets a reference to the given string and assigns it to the CovCategory field.
func (o *ValidateResponseKycResultAllClassOfVehicleInner) SetCovCategory(v string) {
	o.CovCategory = &v
}

func (o ValidateResponseKycResultAllClassOfVehicleInner) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Cov) {
		toSerialize["cov"] = o.Cov
	}
	if !isNil(o.ExpiryDate) {
		toSerialize["expiryDate"] = o.ExpiryDate
	}
	if !isNil(o.IssueDate) {
		toSerialize["issueDate"] = o.IssueDate
	}
	if !isNil(o.CovCategory) {
		toSerialize["covCategory"] = o.CovCategory
	}
	return json.Marshal(toSerialize)
}

type NullableValidateResponseKycResultAllClassOfVehicleInner struct {
	value *ValidateResponseKycResultAllClassOfVehicleInner
	isSet bool
}

func (v NullableValidateResponseKycResultAllClassOfVehicleInner) Get() *ValidateResponseKycResultAllClassOfVehicleInner {
	return v.value
}

func (v *NullableValidateResponseKycResultAllClassOfVehicleInner) Set(val *ValidateResponseKycResultAllClassOfVehicleInner) {
	v.value = val
	v.isSet = true
}

func (v NullableValidateResponseKycResultAllClassOfVehicleInner) IsSet() bool {
	return v.isSet
}

func (v *NullableValidateResponseKycResultAllClassOfVehicleInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableValidateResponseKycResultAllClassOfVehicleInner(val *ValidateResponseKycResultAllClassOfVehicleInner) *NullableValidateResponseKycResultAllClassOfVehicleInner {
	return &NullableValidateResponseKycResultAllClassOfVehicleInner{value: val, isSet: true}
}

func (v NullableValidateResponseKycResultAllClassOfVehicleInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableValidateResponseKycResultAllClassOfVehicleInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


