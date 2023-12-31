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

// AadhaarxmlGenerateOtpRequest struct for AadhaarxmlGenerateOtpRequest
type AadhaarxmlGenerateOtpRequest struct {
	ReferenceId *string `json:"reference_id,omitempty"`
	Consent *bool `json:"consent,omitempty"`
	Purpose *string `json:"purpose,omitempty"`
	InitiationTransactionId *string `json:"initiation_transaction_id,omitempty"`
	AadhaarNumber *string `json:"aadhaar_number,omitempty"`
	Captcha *string `json:"captcha,omitempty"`
}

// NewAadhaarxmlGenerateOtpRequest instantiates a new AadhaarxmlGenerateOtpRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAadhaarxmlGenerateOtpRequest() *AadhaarxmlGenerateOtpRequest {
	this := AadhaarxmlGenerateOtpRequest{}
	return &this
}

// NewAadhaarxmlGenerateOtpRequestWithDefaults instantiates a new AadhaarxmlGenerateOtpRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAadhaarxmlGenerateOtpRequestWithDefaults() *AadhaarxmlGenerateOtpRequest {
	this := AadhaarxmlGenerateOtpRequest{}
	return &this
}

// GetReferenceId returns the ReferenceId field value if set, zero value otherwise.
func (o *AadhaarxmlGenerateOtpRequest) GetReferenceId() string {
	if o == nil || isNil(o.ReferenceId) {
		var ret string
		return ret
	}
	return *o.ReferenceId
}

// GetReferenceIdOk returns a tuple with the ReferenceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AadhaarxmlGenerateOtpRequest) GetReferenceIdOk() (*string, bool) {
	if o == nil || isNil(o.ReferenceId) {
    return nil, false
	}
	return o.ReferenceId, true
}

// HasReferenceId returns a boolean if a field has been set.
func (o *AadhaarxmlGenerateOtpRequest) HasReferenceId() bool {
	if o != nil && !isNil(o.ReferenceId) {
		return true
	}

	return false
}

// SetReferenceId gets a reference to the given string and assigns it to the ReferenceId field.
func (o *AadhaarxmlGenerateOtpRequest) SetReferenceId(v string) {
	o.ReferenceId = &v
}

// GetConsent returns the Consent field value if set, zero value otherwise.
func (o *AadhaarxmlGenerateOtpRequest) GetConsent() bool {
	if o == nil || isNil(o.Consent) {
		var ret bool
		return ret
	}
	return *o.Consent
}

// GetConsentOk returns a tuple with the Consent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AadhaarxmlGenerateOtpRequest) GetConsentOk() (*bool, bool) {
	if o == nil || isNil(o.Consent) {
    return nil, false
	}
	return o.Consent, true
}

// HasConsent returns a boolean if a field has been set.
func (o *AadhaarxmlGenerateOtpRequest) HasConsent() bool {
	if o != nil && !isNil(o.Consent) {
		return true
	}

	return false
}

// SetConsent gets a reference to the given bool and assigns it to the Consent field.
func (o *AadhaarxmlGenerateOtpRequest) SetConsent(v bool) {
	o.Consent = &v
}

// GetPurpose returns the Purpose field value if set, zero value otherwise.
func (o *AadhaarxmlGenerateOtpRequest) GetPurpose() string {
	if o == nil || isNil(o.Purpose) {
		var ret string
		return ret
	}
	return *o.Purpose
}

// GetPurposeOk returns a tuple with the Purpose field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AadhaarxmlGenerateOtpRequest) GetPurposeOk() (*string, bool) {
	if o == nil || isNil(o.Purpose) {
    return nil, false
	}
	return o.Purpose, true
}

// HasPurpose returns a boolean if a field has been set.
func (o *AadhaarxmlGenerateOtpRequest) HasPurpose() bool {
	if o != nil && !isNil(o.Purpose) {
		return true
	}

	return false
}

// SetPurpose gets a reference to the given string and assigns it to the Purpose field.
func (o *AadhaarxmlGenerateOtpRequest) SetPurpose(v string) {
	o.Purpose = &v
}

// GetInitiationTransactionId returns the InitiationTransactionId field value if set, zero value otherwise.
func (o *AadhaarxmlGenerateOtpRequest) GetInitiationTransactionId() string {
	if o == nil || isNil(o.InitiationTransactionId) {
		var ret string
		return ret
	}
	return *o.InitiationTransactionId
}

// GetInitiationTransactionIdOk returns a tuple with the InitiationTransactionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AadhaarxmlGenerateOtpRequest) GetInitiationTransactionIdOk() (*string, bool) {
	if o == nil || isNil(o.InitiationTransactionId) {
    return nil, false
	}
	return o.InitiationTransactionId, true
}

// HasInitiationTransactionId returns a boolean if a field has been set.
func (o *AadhaarxmlGenerateOtpRequest) HasInitiationTransactionId() bool {
	if o != nil && !isNil(o.InitiationTransactionId) {
		return true
	}

	return false
}

// SetInitiationTransactionId gets a reference to the given string and assigns it to the InitiationTransactionId field.
func (o *AadhaarxmlGenerateOtpRequest) SetInitiationTransactionId(v string) {
	o.InitiationTransactionId = &v
}

// GetAadhaarNumber returns the AadhaarNumber field value if set, zero value otherwise.
func (o *AadhaarxmlGenerateOtpRequest) GetAadhaarNumber() string {
	if o == nil || isNil(o.AadhaarNumber) {
		var ret string
		return ret
	}
	return *o.AadhaarNumber
}

// GetAadhaarNumberOk returns a tuple with the AadhaarNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AadhaarxmlGenerateOtpRequest) GetAadhaarNumberOk() (*string, bool) {
	if o == nil || isNil(o.AadhaarNumber) {
    return nil, false
	}
	return o.AadhaarNumber, true
}

// HasAadhaarNumber returns a boolean if a field has been set.
func (o *AadhaarxmlGenerateOtpRequest) HasAadhaarNumber() bool {
	if o != nil && !isNil(o.AadhaarNumber) {
		return true
	}

	return false
}

// SetAadhaarNumber gets a reference to the given string and assigns it to the AadhaarNumber field.
func (o *AadhaarxmlGenerateOtpRequest) SetAadhaarNumber(v string) {
	o.AadhaarNumber = &v
}

// GetCaptcha returns the Captcha field value if set, zero value otherwise.
func (o *AadhaarxmlGenerateOtpRequest) GetCaptcha() string {
	if o == nil || isNil(o.Captcha) {
		var ret string
		return ret
	}
	return *o.Captcha
}

// GetCaptchaOk returns a tuple with the Captcha field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AadhaarxmlGenerateOtpRequest) GetCaptchaOk() (*string, bool) {
	if o == nil || isNil(o.Captcha) {
    return nil, false
	}
	return o.Captcha, true
}

// HasCaptcha returns a boolean if a field has been set.
func (o *AadhaarxmlGenerateOtpRequest) HasCaptcha() bool {
	if o != nil && !isNil(o.Captcha) {
		return true
	}

	return false
}

// SetCaptcha gets a reference to the given string and assigns it to the Captcha field.
func (o *AadhaarxmlGenerateOtpRequest) SetCaptcha(v string) {
	o.Captcha = &v
}

func (o AadhaarxmlGenerateOtpRequest) MarshalJSON() ([]byte, error) {
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
	if !isNil(o.AadhaarNumber) {
		toSerialize["aadhaar_number"] = o.AadhaarNumber
	}
	if !isNil(o.Captcha) {
		toSerialize["captcha"] = o.Captcha
	}
	return json.Marshal(toSerialize)
}

type NullableAadhaarxmlGenerateOtpRequest struct {
	value *AadhaarxmlGenerateOtpRequest
	isSet bool
}

func (v NullableAadhaarxmlGenerateOtpRequest) Get() *AadhaarxmlGenerateOtpRequest {
	return v.value
}

func (v *NullableAadhaarxmlGenerateOtpRequest) Set(val *AadhaarxmlGenerateOtpRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAadhaarxmlGenerateOtpRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAadhaarxmlGenerateOtpRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAadhaarxmlGenerateOtpRequest(val *AadhaarxmlGenerateOtpRequest) *NullableAadhaarxmlGenerateOtpRequest {
	return &NullableAadhaarxmlGenerateOtpRequest{value: val, isSet: true}
}

func (v NullableAadhaarxmlGenerateOtpRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAadhaarxmlGenerateOtpRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


