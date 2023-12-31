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

// DigilockerInitiateSessionRequest struct for DigilockerInitiateSessionRequest
type DigilockerInitiateSessionRequest struct {
	Consent *bool `json:"consent,omitempty"`
	Purpose *string `json:"purpose,omitempty"`
	ReferenceId *string `json:"reference_id,omitempty"`
	RedirectUrl *string `json:"redirect_url,omitempty"`
}

// NewDigilockerInitiateSessionRequest instantiates a new DigilockerInitiateSessionRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDigilockerInitiateSessionRequest() *DigilockerInitiateSessionRequest {
	this := DigilockerInitiateSessionRequest{}
	return &this
}

// NewDigilockerInitiateSessionRequestWithDefaults instantiates a new DigilockerInitiateSessionRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDigilockerInitiateSessionRequestWithDefaults() *DigilockerInitiateSessionRequest {
	this := DigilockerInitiateSessionRequest{}
	return &this
}

// GetConsent returns the Consent field value if set, zero value otherwise.
func (o *DigilockerInitiateSessionRequest) GetConsent() bool {
	if o == nil || isNil(o.Consent) {
		var ret bool
		return ret
	}
	return *o.Consent
}

// GetConsentOk returns a tuple with the Consent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DigilockerInitiateSessionRequest) GetConsentOk() (*bool, bool) {
	if o == nil || isNil(o.Consent) {
    return nil, false
	}
	return o.Consent, true
}

// HasConsent returns a boolean if a field has been set.
func (o *DigilockerInitiateSessionRequest) HasConsent() bool {
	if o != nil && !isNil(o.Consent) {
		return true
	}

	return false
}

// SetConsent gets a reference to the given bool and assigns it to the Consent field.
func (o *DigilockerInitiateSessionRequest) SetConsent(v bool) {
	o.Consent = &v
}

// GetPurpose returns the Purpose field value if set, zero value otherwise.
func (o *DigilockerInitiateSessionRequest) GetPurpose() string {
	if o == nil || isNil(o.Purpose) {
		var ret string
		return ret
	}
	return *o.Purpose
}

// GetPurposeOk returns a tuple with the Purpose field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DigilockerInitiateSessionRequest) GetPurposeOk() (*string, bool) {
	if o == nil || isNil(o.Purpose) {
    return nil, false
	}
	return o.Purpose, true
}

// HasPurpose returns a boolean if a field has been set.
func (o *DigilockerInitiateSessionRequest) HasPurpose() bool {
	if o != nil && !isNil(o.Purpose) {
		return true
	}

	return false
}

// SetPurpose gets a reference to the given string and assigns it to the Purpose field.
func (o *DigilockerInitiateSessionRequest) SetPurpose(v string) {
	o.Purpose = &v
}

// GetReferenceId returns the ReferenceId field value if set, zero value otherwise.
func (o *DigilockerInitiateSessionRequest) GetReferenceId() string {
	if o == nil || isNil(o.ReferenceId) {
		var ret string
		return ret
	}
	return *o.ReferenceId
}

// GetReferenceIdOk returns a tuple with the ReferenceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DigilockerInitiateSessionRequest) GetReferenceIdOk() (*string, bool) {
	if o == nil || isNil(o.ReferenceId) {
    return nil, false
	}
	return o.ReferenceId, true
}

// HasReferenceId returns a boolean if a field has been set.
func (o *DigilockerInitiateSessionRequest) HasReferenceId() bool {
	if o != nil && !isNil(o.ReferenceId) {
		return true
	}

	return false
}

// SetReferenceId gets a reference to the given string and assigns it to the ReferenceId field.
func (o *DigilockerInitiateSessionRequest) SetReferenceId(v string) {
	o.ReferenceId = &v
}

// GetRedirectUrl returns the RedirectUrl field value if set, zero value otherwise.
func (o *DigilockerInitiateSessionRequest) GetRedirectUrl() string {
	if o == nil || isNil(o.RedirectUrl) {
		var ret string
		return ret
	}
	return *o.RedirectUrl
}

// GetRedirectUrlOk returns a tuple with the RedirectUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DigilockerInitiateSessionRequest) GetRedirectUrlOk() (*string, bool) {
	if o == nil || isNil(o.RedirectUrl) {
    return nil, false
	}
	return o.RedirectUrl, true
}

// HasRedirectUrl returns a boolean if a field has been set.
func (o *DigilockerInitiateSessionRequest) HasRedirectUrl() bool {
	if o != nil && !isNil(o.RedirectUrl) {
		return true
	}

	return false
}

// SetRedirectUrl gets a reference to the given string and assigns it to the RedirectUrl field.
func (o *DigilockerInitiateSessionRequest) SetRedirectUrl(v string) {
	o.RedirectUrl = &v
}

func (o DigilockerInitiateSessionRequest) MarshalJSON() ([]byte, error) {
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
	if !isNil(o.RedirectUrl) {
		toSerialize["redirect_url"] = o.RedirectUrl
	}
	return json.Marshal(toSerialize)
}

type NullableDigilockerInitiateSessionRequest struct {
	value *DigilockerInitiateSessionRequest
	isSet bool
}

func (v NullableDigilockerInitiateSessionRequest) Get() *DigilockerInitiateSessionRequest {
	return v.value
}

func (v *NullableDigilockerInitiateSessionRequest) Set(val *DigilockerInitiateSessionRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableDigilockerInitiateSessionRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableDigilockerInitiateSessionRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDigilockerInitiateSessionRequest(val *DigilockerInitiateSessionRequest) *NullableDigilockerInitiateSessionRequest {
	return &NullableDigilockerInitiateSessionRequest{value: val, isSet: true}
}

func (v NullableDigilockerInitiateSessionRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDigilockerInitiateSessionRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


