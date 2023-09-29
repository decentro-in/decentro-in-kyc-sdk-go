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

// AadhaarxmlReloadCaptchaRequest struct for AadhaarxmlReloadCaptchaRequest
type AadhaarxmlReloadCaptchaRequest struct {
	ReferenceId *string `json:"reference_id,omitempty"`
	Consent *bool `json:"consent,omitempty"`
	Purpose *string `json:"purpose,omitempty"`
	InitiationTransactionId *string `json:"initiation_transaction_id,omitempty"`
}

// NewAadhaarxmlReloadCaptchaRequest instantiates a new AadhaarxmlReloadCaptchaRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAadhaarxmlReloadCaptchaRequest() *AadhaarxmlReloadCaptchaRequest {
	this := AadhaarxmlReloadCaptchaRequest{}
	return &this
}

// NewAadhaarxmlReloadCaptchaRequestWithDefaults instantiates a new AadhaarxmlReloadCaptchaRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAadhaarxmlReloadCaptchaRequestWithDefaults() *AadhaarxmlReloadCaptchaRequest {
	this := AadhaarxmlReloadCaptchaRequest{}
	return &this
}

// GetReferenceId returns the ReferenceId field value if set, zero value otherwise.
func (o *AadhaarxmlReloadCaptchaRequest) GetReferenceId() string {
	if o == nil || isNil(o.ReferenceId) {
		var ret string
		return ret
	}
	return *o.ReferenceId
}

// GetReferenceIdOk returns a tuple with the ReferenceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AadhaarxmlReloadCaptchaRequest) GetReferenceIdOk() (*string, bool) {
	if o == nil || isNil(o.ReferenceId) {
    return nil, false
	}
	return o.ReferenceId, true
}

// HasReferenceId returns a boolean if a field has been set.
func (o *AadhaarxmlReloadCaptchaRequest) HasReferenceId() bool {
	if o != nil && !isNil(o.ReferenceId) {
		return true
	}

	return false
}

// SetReferenceId gets a reference to the given string and assigns it to the ReferenceId field.
func (o *AadhaarxmlReloadCaptchaRequest) SetReferenceId(v string) {
	o.ReferenceId = &v
}

// GetConsent returns the Consent field value if set, zero value otherwise.
func (o *AadhaarxmlReloadCaptchaRequest) GetConsent() bool {
	if o == nil || isNil(o.Consent) {
		var ret bool
		return ret
	}
	return *o.Consent
}

// GetConsentOk returns a tuple with the Consent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AadhaarxmlReloadCaptchaRequest) GetConsentOk() (*bool, bool) {
	if o == nil || isNil(o.Consent) {
    return nil, false
	}
	return o.Consent, true
}

// HasConsent returns a boolean if a field has been set.
func (o *AadhaarxmlReloadCaptchaRequest) HasConsent() bool {
	if o != nil && !isNil(o.Consent) {
		return true
	}

	return false
}

// SetConsent gets a reference to the given bool and assigns it to the Consent field.
func (o *AadhaarxmlReloadCaptchaRequest) SetConsent(v bool) {
	o.Consent = &v
}

// GetPurpose returns the Purpose field value if set, zero value otherwise.
func (o *AadhaarxmlReloadCaptchaRequest) GetPurpose() string {
	if o == nil || isNil(o.Purpose) {
		var ret string
		return ret
	}
	return *o.Purpose
}

// GetPurposeOk returns a tuple with the Purpose field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AadhaarxmlReloadCaptchaRequest) GetPurposeOk() (*string, bool) {
	if o == nil || isNil(o.Purpose) {
    return nil, false
	}
	return o.Purpose, true
}

// HasPurpose returns a boolean if a field has been set.
func (o *AadhaarxmlReloadCaptchaRequest) HasPurpose() bool {
	if o != nil && !isNil(o.Purpose) {
		return true
	}

	return false
}

// SetPurpose gets a reference to the given string and assigns it to the Purpose field.
func (o *AadhaarxmlReloadCaptchaRequest) SetPurpose(v string) {
	o.Purpose = &v
}

// GetInitiationTransactionId returns the InitiationTransactionId field value if set, zero value otherwise.
func (o *AadhaarxmlReloadCaptchaRequest) GetInitiationTransactionId() string {
	if o == nil || isNil(o.InitiationTransactionId) {
		var ret string
		return ret
	}
	return *o.InitiationTransactionId
}

// GetInitiationTransactionIdOk returns a tuple with the InitiationTransactionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AadhaarxmlReloadCaptchaRequest) GetInitiationTransactionIdOk() (*string, bool) {
	if o == nil || isNil(o.InitiationTransactionId) {
    return nil, false
	}
	return o.InitiationTransactionId, true
}

// HasInitiationTransactionId returns a boolean if a field has been set.
func (o *AadhaarxmlReloadCaptchaRequest) HasInitiationTransactionId() bool {
	if o != nil && !isNil(o.InitiationTransactionId) {
		return true
	}

	return false
}

// SetInitiationTransactionId gets a reference to the given string and assigns it to the InitiationTransactionId field.
func (o *AadhaarxmlReloadCaptchaRequest) SetInitiationTransactionId(v string) {
	o.InitiationTransactionId = &v
}

func (o AadhaarxmlReloadCaptchaRequest) MarshalJSON() ([]byte, error) {
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
	if !isNil(o.InitiationTransactionId) {
		toSerialize["initiation_transaction_id"] = o.InitiationTransactionId
	}
	return json.Marshal(toSerialize)
}

type NullableAadhaarxmlReloadCaptchaRequest struct {
	value *AadhaarxmlReloadCaptchaRequest
	isSet bool
}

func (v NullableAadhaarxmlReloadCaptchaRequest) Get() *AadhaarxmlReloadCaptchaRequest {
	return v.value
}

func (v *NullableAadhaarxmlReloadCaptchaRequest) Set(val *AadhaarxmlReloadCaptchaRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAadhaarxmlReloadCaptchaRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAadhaarxmlReloadCaptchaRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAadhaarxmlReloadCaptchaRequest(val *AadhaarxmlReloadCaptchaRequest) *NullableAadhaarxmlReloadCaptchaRequest {
	return &NullableAadhaarxmlReloadCaptchaRequest{value: val, isSet: true}
}

func (v NullableAadhaarxmlReloadCaptchaRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAadhaarxmlReloadCaptchaRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


