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

// UpgradedaadhaarxmlGenerateOtpRequest struct for UpgradedaadhaarxmlGenerateOtpRequest
type UpgradedaadhaarxmlGenerateOtpRequest struct {
	ReferenceId *string `json:"reference_id,omitempty"`
	Consent *bool `json:"consent,omitempty"`
	Purpose *string `json:"purpose,omitempty"`
	AadhaarNumber *string `json:"aadhaar_number,omitempty"`
}

// NewUpgradedaadhaarxmlGenerateOtpRequest instantiates a new UpgradedaadhaarxmlGenerateOtpRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpgradedaadhaarxmlGenerateOtpRequest() *UpgradedaadhaarxmlGenerateOtpRequest {
	this := UpgradedaadhaarxmlGenerateOtpRequest{}
	return &this
}

// NewUpgradedaadhaarxmlGenerateOtpRequestWithDefaults instantiates a new UpgradedaadhaarxmlGenerateOtpRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpgradedaadhaarxmlGenerateOtpRequestWithDefaults() *UpgradedaadhaarxmlGenerateOtpRequest {
	this := UpgradedaadhaarxmlGenerateOtpRequest{}
	return &this
}

// GetReferenceId returns the ReferenceId field value if set, zero value otherwise.
func (o *UpgradedaadhaarxmlGenerateOtpRequest) GetReferenceId() string {
	if o == nil || isNil(o.ReferenceId) {
		var ret string
		return ret
	}
	return *o.ReferenceId
}

// GetReferenceIdOk returns a tuple with the ReferenceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpgradedaadhaarxmlGenerateOtpRequest) GetReferenceIdOk() (*string, bool) {
	if o == nil || isNil(o.ReferenceId) {
    return nil, false
	}
	return o.ReferenceId, true
}

// HasReferenceId returns a boolean if a field has been set.
func (o *UpgradedaadhaarxmlGenerateOtpRequest) HasReferenceId() bool {
	if o != nil && !isNil(o.ReferenceId) {
		return true
	}

	return false
}

// SetReferenceId gets a reference to the given string and assigns it to the ReferenceId field.
func (o *UpgradedaadhaarxmlGenerateOtpRequest) SetReferenceId(v string) {
	o.ReferenceId = &v
}

// GetConsent returns the Consent field value if set, zero value otherwise.
func (o *UpgradedaadhaarxmlGenerateOtpRequest) GetConsent() bool {
	if o == nil || isNil(o.Consent) {
		var ret bool
		return ret
	}
	return *o.Consent
}

// GetConsentOk returns a tuple with the Consent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpgradedaadhaarxmlGenerateOtpRequest) GetConsentOk() (*bool, bool) {
	if o == nil || isNil(o.Consent) {
    return nil, false
	}
	return o.Consent, true
}

// HasConsent returns a boolean if a field has been set.
func (o *UpgradedaadhaarxmlGenerateOtpRequest) HasConsent() bool {
	if o != nil && !isNil(o.Consent) {
		return true
	}

	return false
}

// SetConsent gets a reference to the given bool and assigns it to the Consent field.
func (o *UpgradedaadhaarxmlGenerateOtpRequest) SetConsent(v bool) {
	o.Consent = &v
}

// GetPurpose returns the Purpose field value if set, zero value otherwise.
func (o *UpgradedaadhaarxmlGenerateOtpRequest) GetPurpose() string {
	if o == nil || isNil(o.Purpose) {
		var ret string
		return ret
	}
	return *o.Purpose
}

// GetPurposeOk returns a tuple with the Purpose field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpgradedaadhaarxmlGenerateOtpRequest) GetPurposeOk() (*string, bool) {
	if o == nil || isNil(o.Purpose) {
    return nil, false
	}
	return o.Purpose, true
}

// HasPurpose returns a boolean if a field has been set.
func (o *UpgradedaadhaarxmlGenerateOtpRequest) HasPurpose() bool {
	if o != nil && !isNil(o.Purpose) {
		return true
	}

	return false
}

// SetPurpose gets a reference to the given string and assigns it to the Purpose field.
func (o *UpgradedaadhaarxmlGenerateOtpRequest) SetPurpose(v string) {
	o.Purpose = &v
}

// GetAadhaarNumber returns the AadhaarNumber field value if set, zero value otherwise.
func (o *UpgradedaadhaarxmlGenerateOtpRequest) GetAadhaarNumber() string {
	if o == nil || isNil(o.AadhaarNumber) {
		var ret string
		return ret
	}
	return *o.AadhaarNumber
}

// GetAadhaarNumberOk returns a tuple with the AadhaarNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpgradedaadhaarxmlGenerateOtpRequest) GetAadhaarNumberOk() (*string, bool) {
	if o == nil || isNil(o.AadhaarNumber) {
    return nil, false
	}
	return o.AadhaarNumber, true
}

// HasAadhaarNumber returns a boolean if a field has been set.
func (o *UpgradedaadhaarxmlGenerateOtpRequest) HasAadhaarNumber() bool {
	if o != nil && !isNil(o.AadhaarNumber) {
		return true
	}

	return false
}

// SetAadhaarNumber gets a reference to the given string and assigns it to the AadhaarNumber field.
func (o *UpgradedaadhaarxmlGenerateOtpRequest) SetAadhaarNumber(v string) {
	o.AadhaarNumber = &v
}

func (o UpgradedaadhaarxmlGenerateOtpRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.ReferenceId) {
		toSerialize["reference_id"] = o.ReferenceId
	}
	if !isNil(o.Consent) {
		toSerialize["consent"] = o.Consent
	}
	if !isNil(o.Purpose) {
		toSerialize["purpose"] = o.Purpose
	}
	if !isNil(o.AadhaarNumber) {
		toSerialize["aadhaar_number"] = o.AadhaarNumber
	}
	return json.Marshal(toSerialize)
}

type NullableUpgradedaadhaarxmlGenerateOtpRequest struct {
	value *UpgradedaadhaarxmlGenerateOtpRequest
	isSet bool
}

func (v NullableUpgradedaadhaarxmlGenerateOtpRequest) Get() *UpgradedaadhaarxmlGenerateOtpRequest {
	return v.value
}

func (v *NullableUpgradedaadhaarxmlGenerateOtpRequest) Set(val *UpgradedaadhaarxmlGenerateOtpRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpgradedaadhaarxmlGenerateOtpRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpgradedaadhaarxmlGenerateOtpRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpgradedaadhaarxmlGenerateOtpRequest(val *UpgradedaadhaarxmlGenerateOtpRequest) *NullableUpgradedaadhaarxmlGenerateOtpRequest {
	return &NullableUpgradedaadhaarxmlGenerateOtpRequest{value: val, isSet: true}
}

func (v NullableUpgradedaadhaarxmlGenerateOtpRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpgradedaadhaarxmlGenerateOtpRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


