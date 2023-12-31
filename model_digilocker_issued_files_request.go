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

// DigilockerIssuedFilesRequest struct for DigilockerIssuedFilesRequest
type DigilockerIssuedFilesRequest struct {
	Consent *bool `json:"consent,omitempty"`
	Purpose *string `json:"purpose,omitempty"`
	ReferenceId *string `json:"reference_id,omitempty"`
}

// NewDigilockerIssuedFilesRequest instantiates a new DigilockerIssuedFilesRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDigilockerIssuedFilesRequest() *DigilockerIssuedFilesRequest {
	this := DigilockerIssuedFilesRequest{}
	return &this
}

// NewDigilockerIssuedFilesRequestWithDefaults instantiates a new DigilockerIssuedFilesRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDigilockerIssuedFilesRequestWithDefaults() *DigilockerIssuedFilesRequest {
	this := DigilockerIssuedFilesRequest{}
	return &this
}

// GetConsent returns the Consent field value if set, zero value otherwise.
func (o *DigilockerIssuedFilesRequest) GetConsent() bool {
	if o == nil || isNil(o.Consent) {
		var ret bool
		return ret
	}
	return *o.Consent
}

// GetConsentOk returns a tuple with the Consent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DigilockerIssuedFilesRequest) GetConsentOk() (*bool, bool) {
	if o == nil || isNil(o.Consent) {
    return nil, false
	}
	return o.Consent, true
}

// HasConsent returns a boolean if a field has been set.
func (o *DigilockerIssuedFilesRequest) HasConsent() bool {
	if o != nil && !isNil(o.Consent) {
		return true
	}

	return false
}

// SetConsent gets a reference to the given bool and assigns it to the Consent field.
func (o *DigilockerIssuedFilesRequest) SetConsent(v bool) {
	o.Consent = &v
}

// GetPurpose returns the Purpose field value if set, zero value otherwise.
func (o *DigilockerIssuedFilesRequest) GetPurpose() string {
	if o == nil || isNil(o.Purpose) {
		var ret string
		return ret
	}
	return *o.Purpose
}

// GetPurposeOk returns a tuple with the Purpose field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DigilockerIssuedFilesRequest) GetPurposeOk() (*string, bool) {
	if o == nil || isNil(o.Purpose) {
    return nil, false
	}
	return o.Purpose, true
}

// HasPurpose returns a boolean if a field has been set.
func (o *DigilockerIssuedFilesRequest) HasPurpose() bool {
	if o != nil && !isNil(o.Purpose) {
		return true
	}

	return false
}

// SetPurpose gets a reference to the given string and assigns it to the Purpose field.
func (o *DigilockerIssuedFilesRequest) SetPurpose(v string) {
	o.Purpose = &v
}

// GetReferenceId returns the ReferenceId field value if set, zero value otherwise.
func (o *DigilockerIssuedFilesRequest) GetReferenceId() string {
	if o == nil || isNil(o.ReferenceId) {
		var ret string
		return ret
	}
	return *o.ReferenceId
}

// GetReferenceIdOk returns a tuple with the ReferenceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DigilockerIssuedFilesRequest) GetReferenceIdOk() (*string, bool) {
	if o == nil || isNil(o.ReferenceId) {
    return nil, false
	}
	return o.ReferenceId, true
}

// HasReferenceId returns a boolean if a field has been set.
func (o *DigilockerIssuedFilesRequest) HasReferenceId() bool {
	if o != nil && !isNil(o.ReferenceId) {
		return true
	}

	return false
}

// SetReferenceId gets a reference to the given string and assigns it to the ReferenceId field.
func (o *DigilockerIssuedFilesRequest) SetReferenceId(v string) {
	o.ReferenceId = &v
}

func (o DigilockerIssuedFilesRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
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

type NullableDigilockerIssuedFilesRequest struct {
	value *DigilockerIssuedFilesRequest
	isSet bool
}

func (v NullableDigilockerIssuedFilesRequest) Get() *DigilockerIssuedFilesRequest {
	return v.value
}

func (v *NullableDigilockerIssuedFilesRequest) Set(val *DigilockerIssuedFilesRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableDigilockerIssuedFilesRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableDigilockerIssuedFilesRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDigilockerIssuedFilesRequest(val *DigilockerIssuedFilesRequest) *NullableDigilockerIssuedFilesRequest {
	return &NullableDigilockerIssuedFilesRequest{value: val, isSet: true}
}

func (v NullableDigilockerIssuedFilesRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDigilockerIssuedFilesRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


