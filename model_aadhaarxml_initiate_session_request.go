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

// AadhaarxmlInitiateSessionRequest struct for AadhaarxmlInitiateSessionRequest
type AadhaarxmlInitiateSessionRequest struct {
	ReferenceId *string `json:"reference_id,omitempty"`
	Consent *bool `json:"consent,omitempty"`
	Purpose *string `json:"purpose,omitempty"`
}

// NewAadhaarxmlInitiateSessionRequest instantiates a new AadhaarxmlInitiateSessionRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAadhaarxmlInitiateSessionRequest() *AadhaarxmlInitiateSessionRequest {
	this := AadhaarxmlInitiateSessionRequest{}
	return &this
}

// NewAadhaarxmlInitiateSessionRequestWithDefaults instantiates a new AadhaarxmlInitiateSessionRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAadhaarxmlInitiateSessionRequestWithDefaults() *AadhaarxmlInitiateSessionRequest {
	this := AadhaarxmlInitiateSessionRequest{}
	return &this
}

// GetReferenceId returns the ReferenceId field value if set, zero value otherwise.
func (o *AadhaarxmlInitiateSessionRequest) GetReferenceId() string {
	if o == nil || isNil(o.ReferenceId) {
		var ret string
		return ret
	}
	return *o.ReferenceId
}

// GetReferenceIdOk returns a tuple with the ReferenceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AadhaarxmlInitiateSessionRequest) GetReferenceIdOk() (*string, bool) {
	if o == nil || isNil(o.ReferenceId) {
    return nil, false
	}
	return o.ReferenceId, true
}

// HasReferenceId returns a boolean if a field has been set.
func (o *AadhaarxmlInitiateSessionRequest) HasReferenceId() bool {
	if o != nil && !isNil(o.ReferenceId) {
		return true
	}

	return false
}

// SetReferenceId gets a reference to the given string and assigns it to the ReferenceId field.
func (o *AadhaarxmlInitiateSessionRequest) SetReferenceId(v string) {
	o.ReferenceId = &v
}

// GetConsent returns the Consent field value if set, zero value otherwise.
func (o *AadhaarxmlInitiateSessionRequest) GetConsent() bool {
	if o == nil || isNil(o.Consent) {
		var ret bool
		return ret
	}
	return *o.Consent
}

// GetConsentOk returns a tuple with the Consent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AadhaarxmlInitiateSessionRequest) GetConsentOk() (*bool, bool) {
	if o == nil || isNil(o.Consent) {
    return nil, false
	}
	return o.Consent, true
}

// HasConsent returns a boolean if a field has been set.
func (o *AadhaarxmlInitiateSessionRequest) HasConsent() bool {
	if o != nil && !isNil(o.Consent) {
		return true
	}

	return false
}

// SetConsent gets a reference to the given bool and assigns it to the Consent field.
func (o *AadhaarxmlInitiateSessionRequest) SetConsent(v bool) {
	o.Consent = &v
}

// GetPurpose returns the Purpose field value if set, zero value otherwise.
func (o *AadhaarxmlInitiateSessionRequest) GetPurpose() string {
	if o == nil || isNil(o.Purpose) {
		var ret string
		return ret
	}
	return *o.Purpose
}

// GetPurposeOk returns a tuple with the Purpose field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AadhaarxmlInitiateSessionRequest) GetPurposeOk() (*string, bool) {
	if o == nil || isNil(o.Purpose) {
    return nil, false
	}
	return o.Purpose, true
}

// HasPurpose returns a boolean if a field has been set.
func (o *AadhaarxmlInitiateSessionRequest) HasPurpose() bool {
	if o != nil && !isNil(o.Purpose) {
		return true
	}

	return false
}

// SetPurpose gets a reference to the given string and assigns it to the Purpose field.
func (o *AadhaarxmlInitiateSessionRequest) SetPurpose(v string) {
	o.Purpose = &v
}

func (o AadhaarxmlInitiateSessionRequest) MarshalJSON() ([]byte, error) {
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
	return json.Marshal(toSerialize)
}

type NullableAadhaarxmlInitiateSessionRequest struct {
	value *AadhaarxmlInitiateSessionRequest
	isSet bool
}

func (v NullableAadhaarxmlInitiateSessionRequest) Get() *AadhaarxmlInitiateSessionRequest {
	return v.value
}

func (v *NullableAadhaarxmlInitiateSessionRequest) Set(val *AadhaarxmlInitiateSessionRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAadhaarxmlInitiateSessionRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAadhaarxmlInitiateSessionRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAadhaarxmlInitiateSessionRequest(val *AadhaarxmlInitiateSessionRequest) *NullableAadhaarxmlInitiateSessionRequest {
	return &NullableAadhaarxmlInitiateSessionRequest{value: val, isSet: true}
}

func (v NullableAadhaarxmlInitiateSessionRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAadhaarxmlInitiateSessionRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


