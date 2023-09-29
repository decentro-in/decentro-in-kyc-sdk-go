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

// DigilockerGenerateAccessTokenRequest struct for DigilockerGenerateAccessTokenRequest
type DigilockerGenerateAccessTokenRequest struct {
	Code *string `json:"code,omitempty"`
	Consent *bool `json:"consent,omitempty"`
	Purpose *string `json:"purpose,omitempty"`
	ReferenceId *string `json:"reference_id,omitempty"`
}

// NewDigilockerGenerateAccessTokenRequest instantiates a new DigilockerGenerateAccessTokenRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDigilockerGenerateAccessTokenRequest() *DigilockerGenerateAccessTokenRequest {
	this := DigilockerGenerateAccessTokenRequest{}
	return &this
}

// NewDigilockerGenerateAccessTokenRequestWithDefaults instantiates a new DigilockerGenerateAccessTokenRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDigilockerGenerateAccessTokenRequestWithDefaults() *DigilockerGenerateAccessTokenRequest {
	this := DigilockerGenerateAccessTokenRequest{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *DigilockerGenerateAccessTokenRequest) GetCode() string {
	if o == nil || isNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DigilockerGenerateAccessTokenRequest) GetCodeOk() (*string, bool) {
	if o == nil || isNil(o.Code) {
    return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *DigilockerGenerateAccessTokenRequest) HasCode() bool {
	if o != nil && !isNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *DigilockerGenerateAccessTokenRequest) SetCode(v string) {
	o.Code = &v
}

// GetConsent returns the Consent field value if set, zero value otherwise.
func (o *DigilockerGenerateAccessTokenRequest) GetConsent() bool {
	if o == nil || isNil(o.Consent) {
		var ret bool
		return ret
	}
	return *o.Consent
}

// GetConsentOk returns a tuple with the Consent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DigilockerGenerateAccessTokenRequest) GetConsentOk() (*bool, bool) {
	if o == nil || isNil(o.Consent) {
    return nil, false
	}
	return o.Consent, true
}

// HasConsent returns a boolean if a field has been set.
func (o *DigilockerGenerateAccessTokenRequest) HasConsent() bool {
	if o != nil && !isNil(o.Consent) {
		return true
	}

	return false
}

// SetConsent gets a reference to the given bool and assigns it to the Consent field.
func (o *DigilockerGenerateAccessTokenRequest) SetConsent(v bool) {
	o.Consent = &v
}

// GetPurpose returns the Purpose field value if set, zero value otherwise.
func (o *DigilockerGenerateAccessTokenRequest) GetPurpose() string {
	if o == nil || isNil(o.Purpose) {
		var ret string
		return ret
	}
	return *o.Purpose
}

// GetPurposeOk returns a tuple with the Purpose field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DigilockerGenerateAccessTokenRequest) GetPurposeOk() (*string, bool) {
	if o == nil || isNil(o.Purpose) {
    return nil, false
	}
	return o.Purpose, true
}

// HasPurpose returns a boolean if a field has been set.
func (o *DigilockerGenerateAccessTokenRequest) HasPurpose() bool {
	if o != nil && !isNil(o.Purpose) {
		return true
	}

	return false
}

// SetPurpose gets a reference to the given string and assigns it to the Purpose field.
func (o *DigilockerGenerateAccessTokenRequest) SetPurpose(v string) {
	o.Purpose = &v
}

// GetReferenceId returns the ReferenceId field value if set, zero value otherwise.
func (o *DigilockerGenerateAccessTokenRequest) GetReferenceId() string {
	if o == nil || isNil(o.ReferenceId) {
		var ret string
		return ret
	}
	return *o.ReferenceId
}

// GetReferenceIdOk returns a tuple with the ReferenceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DigilockerGenerateAccessTokenRequest) GetReferenceIdOk() (*string, bool) {
	if o == nil || isNil(o.ReferenceId) {
    return nil, false
	}
	return o.ReferenceId, true
}

// HasReferenceId returns a boolean if a field has been set.
func (o *DigilockerGenerateAccessTokenRequest) HasReferenceId() bool {
	if o != nil && !isNil(o.ReferenceId) {
		return true
	}

	return false
}

// SetReferenceId gets a reference to the given string and assigns it to the ReferenceId field.
func (o *DigilockerGenerateAccessTokenRequest) SetReferenceId(v string) {
	o.ReferenceId = &v
}

func (o DigilockerGenerateAccessTokenRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	if !isNil(o.Consent) {
		toSerialize["consent"] = o.Consent
	}
	if !isNil(o.Purpose) {
		toSerialize["purpose"] = o.Purpose
	}
	if !isNil(o.ReferenceId) {
		toSerialize["reference_id"] = o.ReferenceId
	}
	return json.Marshal(toSerialize)
}

type NullableDigilockerGenerateAccessTokenRequest struct {
	value *DigilockerGenerateAccessTokenRequest
	isSet bool
}

func (v NullableDigilockerGenerateAccessTokenRequest) Get() *DigilockerGenerateAccessTokenRequest {
	return v.value
}

func (v *NullableDigilockerGenerateAccessTokenRequest) Set(val *DigilockerGenerateAccessTokenRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableDigilockerGenerateAccessTokenRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableDigilockerGenerateAccessTokenRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDigilockerGenerateAccessTokenRequest(val *DigilockerGenerateAccessTokenRequest) *NullableDigilockerGenerateAccessTokenRequest {
	return &NullableDigilockerGenerateAccessTokenRequest{value: val, isSet: true}
}

func (v NullableDigilockerGenerateAccessTokenRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDigilockerGenerateAccessTokenRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


