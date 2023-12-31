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

// DigilockerPullFileRequest struct for DigilockerPullFileRequest
type DigilockerPullFileRequest struct {
	Consent *bool `json:"consent,omitempty"`
	Purpose *string `json:"purpose,omitempty"`
	ReferenceId *string `json:"reference_id,omitempty"`
	DocumentType *string `json:"document_type,omitempty"`
	DocumentTypeRelatedParameters *DigilockerPullFileRequestDocumentTypeRelatedParameters `json:"document_type_related_parameters,omitempty"`
}

// NewDigilockerPullFileRequest instantiates a new DigilockerPullFileRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDigilockerPullFileRequest() *DigilockerPullFileRequest {
	this := DigilockerPullFileRequest{}
	return &this
}

// NewDigilockerPullFileRequestWithDefaults instantiates a new DigilockerPullFileRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDigilockerPullFileRequestWithDefaults() *DigilockerPullFileRequest {
	this := DigilockerPullFileRequest{}
	return &this
}

// GetConsent returns the Consent field value if set, zero value otherwise.
func (o *DigilockerPullFileRequest) GetConsent() bool {
	if o == nil || isNil(o.Consent) {
		var ret bool
		return ret
	}
	return *o.Consent
}

// GetConsentOk returns a tuple with the Consent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DigilockerPullFileRequest) GetConsentOk() (*bool, bool) {
	if o == nil || isNil(o.Consent) {
    return nil, false
	}
	return o.Consent, true
}

// HasConsent returns a boolean if a field has been set.
func (o *DigilockerPullFileRequest) HasConsent() bool {
	if o != nil && !isNil(o.Consent) {
		return true
	}

	return false
}

// SetConsent gets a reference to the given bool and assigns it to the Consent field.
func (o *DigilockerPullFileRequest) SetConsent(v bool) {
	o.Consent = &v
}

// GetPurpose returns the Purpose field value if set, zero value otherwise.
func (o *DigilockerPullFileRequest) GetPurpose() string {
	if o == nil || isNil(o.Purpose) {
		var ret string
		return ret
	}
	return *o.Purpose
}

// GetPurposeOk returns a tuple with the Purpose field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DigilockerPullFileRequest) GetPurposeOk() (*string, bool) {
	if o == nil || isNil(o.Purpose) {
    return nil, false
	}
	return o.Purpose, true
}

// HasPurpose returns a boolean if a field has been set.
func (o *DigilockerPullFileRequest) HasPurpose() bool {
	if o != nil && !isNil(o.Purpose) {
		return true
	}

	return false
}

// SetPurpose gets a reference to the given string and assigns it to the Purpose field.
func (o *DigilockerPullFileRequest) SetPurpose(v string) {
	o.Purpose = &v
}

// GetReferenceId returns the ReferenceId field value if set, zero value otherwise.
func (o *DigilockerPullFileRequest) GetReferenceId() string {
	if o == nil || isNil(o.ReferenceId) {
		var ret string
		return ret
	}
	return *o.ReferenceId
}

// GetReferenceIdOk returns a tuple with the ReferenceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DigilockerPullFileRequest) GetReferenceIdOk() (*string, bool) {
	if o == nil || isNil(o.ReferenceId) {
    return nil, false
	}
	return o.ReferenceId, true
}

// HasReferenceId returns a boolean if a field has been set.
func (o *DigilockerPullFileRequest) HasReferenceId() bool {
	if o != nil && !isNil(o.ReferenceId) {
		return true
	}

	return false
}

// SetReferenceId gets a reference to the given string and assigns it to the ReferenceId field.
func (o *DigilockerPullFileRequest) SetReferenceId(v string) {
	o.ReferenceId = &v
}

// GetDocumentType returns the DocumentType field value if set, zero value otherwise.
func (o *DigilockerPullFileRequest) GetDocumentType() string {
	if o == nil || isNil(o.DocumentType) {
		var ret string
		return ret
	}
	return *o.DocumentType
}

// GetDocumentTypeOk returns a tuple with the DocumentType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DigilockerPullFileRequest) GetDocumentTypeOk() (*string, bool) {
	if o == nil || isNil(o.DocumentType) {
    return nil, false
	}
	return o.DocumentType, true
}

// HasDocumentType returns a boolean if a field has been set.
func (o *DigilockerPullFileRequest) HasDocumentType() bool {
	if o != nil && !isNil(o.DocumentType) {
		return true
	}

	return false
}

// SetDocumentType gets a reference to the given string and assigns it to the DocumentType field.
func (o *DigilockerPullFileRequest) SetDocumentType(v string) {
	o.DocumentType = &v
}

// GetDocumentTypeRelatedParameters returns the DocumentTypeRelatedParameters field value if set, zero value otherwise.
func (o *DigilockerPullFileRequest) GetDocumentTypeRelatedParameters() DigilockerPullFileRequestDocumentTypeRelatedParameters {
	if o == nil || isNil(o.DocumentTypeRelatedParameters) {
		var ret DigilockerPullFileRequestDocumentTypeRelatedParameters
		return ret
	}
	return *o.DocumentTypeRelatedParameters
}

// GetDocumentTypeRelatedParametersOk returns a tuple with the DocumentTypeRelatedParameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DigilockerPullFileRequest) GetDocumentTypeRelatedParametersOk() (*DigilockerPullFileRequestDocumentTypeRelatedParameters, bool) {
	if o == nil || isNil(o.DocumentTypeRelatedParameters) {
    return nil, false
	}
	return o.DocumentTypeRelatedParameters, true
}

// HasDocumentTypeRelatedParameters returns a boolean if a field has been set.
func (o *DigilockerPullFileRequest) HasDocumentTypeRelatedParameters() bool {
	if o != nil && !isNil(o.DocumentTypeRelatedParameters) {
		return true
	}

	return false
}

// SetDocumentTypeRelatedParameters gets a reference to the given DigilockerPullFileRequestDocumentTypeRelatedParameters and assigns it to the DocumentTypeRelatedParameters field.
func (o *DigilockerPullFileRequest) SetDocumentTypeRelatedParameters(v DigilockerPullFileRequestDocumentTypeRelatedParameters) {
	o.DocumentTypeRelatedParameters = &v
}

func (o DigilockerPullFileRequest) MarshalJSON() ([]byte, error) {
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
	if !isNil(o.DocumentType) {
		toSerialize["document_type"] = o.DocumentType
	}
	if !isNil(o.DocumentTypeRelatedParameters) {
		toSerialize["document_type_related_parameters"] = o.DocumentTypeRelatedParameters
	}
	return json.Marshal(toSerialize)
}

type NullableDigilockerPullFileRequest struct {
	value *DigilockerPullFileRequest
	isSet bool
}

func (v NullableDigilockerPullFileRequest) Get() *DigilockerPullFileRequest {
	return v.value
}

func (v *NullableDigilockerPullFileRequest) Set(val *DigilockerPullFileRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableDigilockerPullFileRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableDigilockerPullFileRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDigilockerPullFileRequest(val *DigilockerPullFileRequest) *NullableDigilockerPullFileRequest {
	return &NullableDigilockerPullFileRequest{value: val, isSet: true}
}

func (v NullableDigilockerPullFileRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDigilockerPullFileRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


