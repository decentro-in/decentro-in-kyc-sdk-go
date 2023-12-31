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

// AadhaarxmlValidateOtpRequest struct for AadhaarxmlValidateOtpRequest
type AadhaarxmlValidateOtpRequest struct {
	ReferenceId *string `json:"reference_id,omitempty"`
	Consent *bool `json:"consent,omitempty"`
	Purpose *string `json:"purpose,omitempty"`
	InitiationTransactionId *string `json:"initiation_transaction_id,omitempty"`
	GeneratePdf *bool `json:"generate_pdf,omitempty"`
	GenerateXml *bool `json:"generate_xml,omitempty"`
	ShareCode *bool `json:"share_code,omitempty"`
	Otp *string `json:"otp,omitempty"`
}

// NewAadhaarxmlValidateOtpRequest instantiates a new AadhaarxmlValidateOtpRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAadhaarxmlValidateOtpRequest() *AadhaarxmlValidateOtpRequest {
	this := AadhaarxmlValidateOtpRequest{}
	return &this
}

// NewAadhaarxmlValidateOtpRequestWithDefaults instantiates a new AadhaarxmlValidateOtpRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAadhaarxmlValidateOtpRequestWithDefaults() *AadhaarxmlValidateOtpRequest {
	this := AadhaarxmlValidateOtpRequest{}
	return &this
}

// GetReferenceId returns the ReferenceId field value if set, zero value otherwise.
func (o *AadhaarxmlValidateOtpRequest) GetReferenceId() string {
	if o == nil || isNil(o.ReferenceId) {
		var ret string
		return ret
	}
	return *o.ReferenceId
}

// GetReferenceIdOk returns a tuple with the ReferenceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AadhaarxmlValidateOtpRequest) GetReferenceIdOk() (*string, bool) {
	if o == nil || isNil(o.ReferenceId) {
    return nil, false
	}
	return o.ReferenceId, true
}

// HasReferenceId returns a boolean if a field has been set.
func (o *AadhaarxmlValidateOtpRequest) HasReferenceId() bool {
	if o != nil && !isNil(o.ReferenceId) {
		return true
	}

	return false
}

// SetReferenceId gets a reference to the given string and assigns it to the ReferenceId field.
func (o *AadhaarxmlValidateOtpRequest) SetReferenceId(v string) {
	o.ReferenceId = &v
}

// GetConsent returns the Consent field value if set, zero value otherwise.
func (o *AadhaarxmlValidateOtpRequest) GetConsent() bool {
	if o == nil || isNil(o.Consent) {
		var ret bool
		return ret
	}
	return *o.Consent
}

// GetConsentOk returns a tuple with the Consent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AadhaarxmlValidateOtpRequest) GetConsentOk() (*bool, bool) {
	if o == nil || isNil(o.Consent) {
    return nil, false
	}
	return o.Consent, true
}

// HasConsent returns a boolean if a field has been set.
func (o *AadhaarxmlValidateOtpRequest) HasConsent() bool {
	if o != nil && !isNil(o.Consent) {
		return true
	}

	return false
}

// SetConsent gets a reference to the given bool and assigns it to the Consent field.
func (o *AadhaarxmlValidateOtpRequest) SetConsent(v bool) {
	o.Consent = &v
}

// GetPurpose returns the Purpose field value if set, zero value otherwise.
func (o *AadhaarxmlValidateOtpRequest) GetPurpose() string {
	if o == nil || isNil(o.Purpose) {
		var ret string
		return ret
	}
	return *o.Purpose
}

// GetPurposeOk returns a tuple with the Purpose field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AadhaarxmlValidateOtpRequest) GetPurposeOk() (*string, bool) {
	if o == nil || isNil(o.Purpose) {
    return nil, false
	}
	return o.Purpose, true
}

// HasPurpose returns a boolean if a field has been set.
func (o *AadhaarxmlValidateOtpRequest) HasPurpose() bool {
	if o != nil && !isNil(o.Purpose) {
		return true
	}

	return false
}

// SetPurpose gets a reference to the given string and assigns it to the Purpose field.
func (o *AadhaarxmlValidateOtpRequest) SetPurpose(v string) {
	o.Purpose = &v
}

// GetInitiationTransactionId returns the InitiationTransactionId field value if set, zero value otherwise.
func (o *AadhaarxmlValidateOtpRequest) GetInitiationTransactionId() string {
	if o == nil || isNil(o.InitiationTransactionId) {
		var ret string
		return ret
	}
	return *o.InitiationTransactionId
}

// GetInitiationTransactionIdOk returns a tuple with the InitiationTransactionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AadhaarxmlValidateOtpRequest) GetInitiationTransactionIdOk() (*string, bool) {
	if o == nil || isNil(o.InitiationTransactionId) {
    return nil, false
	}
	return o.InitiationTransactionId, true
}

// HasInitiationTransactionId returns a boolean if a field has been set.
func (o *AadhaarxmlValidateOtpRequest) HasInitiationTransactionId() bool {
	if o != nil && !isNil(o.InitiationTransactionId) {
		return true
	}

	return false
}

// SetInitiationTransactionId gets a reference to the given string and assigns it to the InitiationTransactionId field.
func (o *AadhaarxmlValidateOtpRequest) SetInitiationTransactionId(v string) {
	o.InitiationTransactionId = &v
}

// GetGeneratePdf returns the GeneratePdf field value if set, zero value otherwise.
func (o *AadhaarxmlValidateOtpRequest) GetGeneratePdf() bool {
	if o == nil || isNil(o.GeneratePdf) {
		var ret bool
		return ret
	}
	return *o.GeneratePdf
}

// GetGeneratePdfOk returns a tuple with the GeneratePdf field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AadhaarxmlValidateOtpRequest) GetGeneratePdfOk() (*bool, bool) {
	if o == nil || isNil(o.GeneratePdf) {
    return nil, false
	}
	return o.GeneratePdf, true
}

// HasGeneratePdf returns a boolean if a field has been set.
func (o *AadhaarxmlValidateOtpRequest) HasGeneratePdf() bool {
	if o != nil && !isNil(o.GeneratePdf) {
		return true
	}

	return false
}

// SetGeneratePdf gets a reference to the given bool and assigns it to the GeneratePdf field.
func (o *AadhaarxmlValidateOtpRequest) SetGeneratePdf(v bool) {
	o.GeneratePdf = &v
}

// GetGenerateXml returns the GenerateXml field value if set, zero value otherwise.
func (o *AadhaarxmlValidateOtpRequest) GetGenerateXml() bool {
	if o == nil || isNil(o.GenerateXml) {
		var ret bool
		return ret
	}
	return *o.GenerateXml
}

// GetGenerateXmlOk returns a tuple with the GenerateXml field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AadhaarxmlValidateOtpRequest) GetGenerateXmlOk() (*bool, bool) {
	if o == nil || isNil(o.GenerateXml) {
    return nil, false
	}
	return o.GenerateXml, true
}

// HasGenerateXml returns a boolean if a field has been set.
func (o *AadhaarxmlValidateOtpRequest) HasGenerateXml() bool {
	if o != nil && !isNil(o.GenerateXml) {
		return true
	}

	return false
}

// SetGenerateXml gets a reference to the given bool and assigns it to the GenerateXml field.
func (o *AadhaarxmlValidateOtpRequest) SetGenerateXml(v bool) {
	o.GenerateXml = &v
}

// GetShareCode returns the ShareCode field value if set, zero value otherwise.
func (o *AadhaarxmlValidateOtpRequest) GetShareCode() bool {
	if o == nil || isNil(o.ShareCode) {
		var ret bool
		return ret
	}
	return *o.ShareCode
}

// GetShareCodeOk returns a tuple with the ShareCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AadhaarxmlValidateOtpRequest) GetShareCodeOk() (*bool, bool) {
	if o == nil || isNil(o.ShareCode) {
    return nil, false
	}
	return o.ShareCode, true
}

// HasShareCode returns a boolean if a field has been set.
func (o *AadhaarxmlValidateOtpRequest) HasShareCode() bool {
	if o != nil && !isNil(o.ShareCode) {
		return true
	}

	return false
}

// SetShareCode gets a reference to the given bool and assigns it to the ShareCode field.
func (o *AadhaarxmlValidateOtpRequest) SetShareCode(v bool) {
	o.ShareCode = &v
}

// GetOtp returns the Otp field value if set, zero value otherwise.
func (o *AadhaarxmlValidateOtpRequest) GetOtp() string {
	if o == nil || isNil(o.Otp) {
		var ret string
		return ret
	}
	return *o.Otp
}

// GetOtpOk returns a tuple with the Otp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AadhaarxmlValidateOtpRequest) GetOtpOk() (*string, bool) {
	if o == nil || isNil(o.Otp) {
    return nil, false
	}
	return o.Otp, true
}

// HasOtp returns a boolean if a field has been set.
func (o *AadhaarxmlValidateOtpRequest) HasOtp() bool {
	if o != nil && !isNil(o.Otp) {
		return true
	}

	return false
}

// SetOtp gets a reference to the given string and assigns it to the Otp field.
func (o *AadhaarxmlValidateOtpRequest) SetOtp(v string) {
	o.Otp = &v
}

func (o AadhaarxmlValidateOtpRequest) MarshalJSON() ([]byte, error) {
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
	if !isNil(o.GeneratePdf) {
		toSerialize["generate_pdf"] = o.GeneratePdf
	}
	if !isNil(o.GenerateXml) {
		toSerialize["generate_xml"] = o.GenerateXml
	}
	if !isNil(o.ShareCode) {
		toSerialize["share_code"] = o.ShareCode
	}
	if !isNil(o.Otp) {
		toSerialize["otp"] = o.Otp
	}
	return json.Marshal(toSerialize)
}

type NullableAadhaarxmlValidateOtpRequest struct {
	value *AadhaarxmlValidateOtpRequest
	isSet bool
}

func (v NullableAadhaarxmlValidateOtpRequest) Get() *AadhaarxmlValidateOtpRequest {
	return v.value
}

func (v *NullableAadhaarxmlValidateOtpRequest) Set(val *AadhaarxmlValidateOtpRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAadhaarxmlValidateOtpRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAadhaarxmlValidateOtpRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAadhaarxmlValidateOtpRequest(val *AadhaarxmlValidateOtpRequest) *NullableAadhaarxmlValidateOtpRequest {
	return &NullableAadhaarxmlValidateOtpRequest{value: val, isSet: true}
}

func (v NullableAadhaarxmlValidateOtpRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAadhaarxmlValidateOtpRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


