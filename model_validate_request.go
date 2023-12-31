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

// ValidateRequest struct for ValidateRequest
type ValidateRequest struct {
	ReferenceId string `json:"reference_id"`
	DocumentType string `json:"document_type"`
	IdNumber string `json:"id_number"`
	Dob *string `json:"dob,omitempty"`
	Consent string `json:"consent"`
	ConsentPurpose string `json:"consent_purpose"`
	Name *string `json:"name,omitempty"`
}

// NewValidateRequest instantiates a new ValidateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewValidateRequest(referenceId string, documentType string, idNumber string, consent string, consentPurpose string) *ValidateRequest {
	this := ValidateRequest{}
	this.ReferenceId = referenceId
	this.DocumentType = documentType
	this.IdNumber = idNumber
	this.Consent = consent
	this.ConsentPurpose = consentPurpose
	return &this
}

// NewValidateRequestWithDefaults instantiates a new ValidateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewValidateRequestWithDefaults() *ValidateRequest {
	this := ValidateRequest{}
	return &this
}

// GetReferenceId returns the ReferenceId field value
func (o *ValidateRequest) GetReferenceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ReferenceId
}

// GetReferenceIdOk returns a tuple with the ReferenceId field value
// and a boolean to check if the value has been set.
func (o *ValidateRequest) GetReferenceIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.ReferenceId, true
}

// SetReferenceId sets field value
func (o *ValidateRequest) SetReferenceId(v string) {
	o.ReferenceId = v
}

// GetDocumentType returns the DocumentType field value
func (o *ValidateRequest) GetDocumentType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DocumentType
}

// GetDocumentTypeOk returns a tuple with the DocumentType field value
// and a boolean to check if the value has been set.
func (o *ValidateRequest) GetDocumentTypeOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.DocumentType, true
}

// SetDocumentType sets field value
func (o *ValidateRequest) SetDocumentType(v string) {
	o.DocumentType = v
}

// GetIdNumber returns the IdNumber field value
func (o *ValidateRequest) GetIdNumber() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.IdNumber
}

// GetIdNumberOk returns a tuple with the IdNumber field value
// and a boolean to check if the value has been set.
func (o *ValidateRequest) GetIdNumberOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.IdNumber, true
}

// SetIdNumber sets field value
func (o *ValidateRequest) SetIdNumber(v string) {
	o.IdNumber = v
}

// GetDob returns the Dob field value if set, zero value otherwise.
func (o *ValidateRequest) GetDob() string {
	if o == nil || isNil(o.Dob) {
		var ret string
		return ret
	}
	return *o.Dob
}

// GetDobOk returns a tuple with the Dob field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidateRequest) GetDobOk() (*string, bool) {
	if o == nil || isNil(o.Dob) {
    return nil, false
	}
	return o.Dob, true
}

// HasDob returns a boolean if a field has been set.
func (o *ValidateRequest) HasDob() bool {
	if o != nil && !isNil(o.Dob) {
		return true
	}

	return false
}

// SetDob gets a reference to the given string and assigns it to the Dob field.
func (o *ValidateRequest) SetDob(v string) {
	o.Dob = &v
}

// GetConsent returns the Consent field value
func (o *ValidateRequest) GetConsent() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Consent
}

// GetConsentOk returns a tuple with the Consent field value
// and a boolean to check if the value has been set.
func (o *ValidateRequest) GetConsentOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Consent, true
}

// SetConsent sets field value
func (o *ValidateRequest) SetConsent(v string) {
	o.Consent = v
}

// GetConsentPurpose returns the ConsentPurpose field value
func (o *ValidateRequest) GetConsentPurpose() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ConsentPurpose
}

// GetConsentPurposeOk returns a tuple with the ConsentPurpose field value
// and a boolean to check if the value has been set.
func (o *ValidateRequest) GetConsentPurposeOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.ConsentPurpose, true
}

// SetConsentPurpose sets field value
func (o *ValidateRequest) SetConsentPurpose(v string) {
	o.ConsentPurpose = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ValidateRequest) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidateRequest) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ValidateRequest) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ValidateRequest) SetName(v string) {
	o.Name = &v
}

func (o ValidateRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["reference_id"] = o.ReferenceId
	}
	if true {
		toSerialize["document_type"] = o.DocumentType
	}
	if true {
		toSerialize["id_number"] = o.IdNumber
	}
	if !isNil(o.Dob) {
		toSerialize["dob"] = o.Dob
	}
	if true {
		toSerialize["consent"] = o.Consent
	}
	if true {
		toSerialize["consent_purpose"] = o.ConsentPurpose
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

type NullableValidateRequest struct {
	value *ValidateRequest
	isSet bool
}

func (v NullableValidateRequest) Get() *ValidateRequest {
	return v.value
}

func (v *NullableValidateRequest) Set(val *ValidateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableValidateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableValidateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableValidateRequest(val *ValidateRequest) *NullableValidateRequest {
	return &NullableValidateRequest{value: val, isSet: true}
}

func (v NullableValidateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableValidateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


