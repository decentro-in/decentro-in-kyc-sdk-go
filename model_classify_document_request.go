/*
decentro-in-kyc

KYC & Onboarding

API version: 1.0.0
Contact: admin@decentro.tech
*/


package decentroinkyc

import (
	"encoding/json"
	"os"
)

// ClassifyDocumentRequest struct for ClassifyDocumentRequest
type ClassifyDocumentRequest struct {
	// 
	ReferenceId string `json:"reference_id"`
	// 
	DocumentType string `json:"document_type"`
	// 
	Consent bool `json:"consent"`
	// 
	ConsentPurpose string `json:"consent_purpose"`
	Document **os.File `json:"document,omitempty"`
	DocumentUrl *string `json:"document_url,omitempty"`
}

// NewClassifyDocumentRequest instantiates a new ClassifyDocumentRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClassifyDocumentRequest(referenceId string, documentType string, consent bool, consentPurpose string) *ClassifyDocumentRequest {
	this := ClassifyDocumentRequest{}
	this.ReferenceId = referenceId
	this.DocumentType = documentType
	this.Consent = consent
	this.ConsentPurpose = consentPurpose
	return &this
}

// NewClassifyDocumentRequestWithDefaults instantiates a new ClassifyDocumentRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClassifyDocumentRequestWithDefaults() *ClassifyDocumentRequest {
	this := ClassifyDocumentRequest{}
	return &this
}

// GetReferenceId returns the ReferenceId field value
func (o *ClassifyDocumentRequest) GetReferenceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ReferenceId
}

// GetReferenceIdOk returns a tuple with the ReferenceId field value
// and a boolean to check if the value has been set.
func (o *ClassifyDocumentRequest) GetReferenceIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.ReferenceId, true
}

// SetReferenceId sets field value
func (o *ClassifyDocumentRequest) SetReferenceId(v string) {
	o.ReferenceId = v
}

// GetDocumentType returns the DocumentType field value
func (o *ClassifyDocumentRequest) GetDocumentType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DocumentType
}

// GetDocumentTypeOk returns a tuple with the DocumentType field value
// and a boolean to check if the value has been set.
func (o *ClassifyDocumentRequest) GetDocumentTypeOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.DocumentType, true
}

// SetDocumentType sets field value
func (o *ClassifyDocumentRequest) SetDocumentType(v string) {
	o.DocumentType = v
}

// GetConsent returns the Consent field value
func (o *ClassifyDocumentRequest) GetConsent() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Consent
}

// GetConsentOk returns a tuple with the Consent field value
// and a boolean to check if the value has been set.
func (o *ClassifyDocumentRequest) GetConsentOk() (*bool, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Consent, true
}

// SetConsent sets field value
func (o *ClassifyDocumentRequest) SetConsent(v bool) {
	o.Consent = v
}

// GetConsentPurpose returns the ConsentPurpose field value
func (o *ClassifyDocumentRequest) GetConsentPurpose() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ConsentPurpose
}

// GetConsentPurposeOk returns a tuple with the ConsentPurpose field value
// and a boolean to check if the value has been set.
func (o *ClassifyDocumentRequest) GetConsentPurposeOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.ConsentPurpose, true
}

// SetConsentPurpose sets field value
func (o *ClassifyDocumentRequest) SetConsentPurpose(v string) {
	o.ConsentPurpose = v
}

// GetDocument returns the Document field value if set, zero value otherwise.
func (o *ClassifyDocumentRequest) GetDocument() *os.File {
	if o == nil || isNil(o.Document) {
		var ret *os.File
		return ret
	}
	return *o.Document
}

// GetDocumentOk returns a tuple with the Document field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClassifyDocumentRequest) GetDocumentOk() (**os.File, bool) {
	if o == nil || isNil(o.Document) {
    return nil, false
	}
	return o.Document, true
}

// HasDocument returns a boolean if a field has been set.
func (o *ClassifyDocumentRequest) HasDocument() bool {
	if o != nil && !isNil(o.Document) {
		return true
	}

	return false
}

// SetDocument gets a reference to the given *os.File and assigns it to the Document field.
func (o *ClassifyDocumentRequest) SetDocument(v *os.File) {
	o.Document = &v
}

// GetDocumentUrl returns the DocumentUrl field value if set, zero value otherwise.
func (o *ClassifyDocumentRequest) GetDocumentUrl() string {
	if o == nil || isNil(o.DocumentUrl) {
		var ret string
		return ret
	}
	return *o.DocumentUrl
}

// GetDocumentUrlOk returns a tuple with the DocumentUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClassifyDocumentRequest) GetDocumentUrlOk() (*string, bool) {
	if o == nil || isNil(o.DocumentUrl) {
    return nil, false
	}
	return o.DocumentUrl, true
}

// HasDocumentUrl returns a boolean if a field has been set.
func (o *ClassifyDocumentRequest) HasDocumentUrl() bool {
	if o != nil && !isNil(o.DocumentUrl) {
		return true
	}

	return false
}

// SetDocumentUrl gets a reference to the given string and assigns it to the DocumentUrl field.
func (o *ClassifyDocumentRequest) SetDocumentUrl(v string) {
	o.DocumentUrl = &v
}

func (o ClassifyDocumentRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["reference_id"] = o.ReferenceId
	}
	if true {
		toSerialize["document_type"] = o.DocumentType
	}
	if true {
		toSerialize["consent"] = o.Consent
	}
	if true {
		toSerialize["consent_purpose"] = o.ConsentPurpose
	}
	if !isNil(o.Document) {
		toSerialize["document"] = o.Document
	}
	if !isNil(o.DocumentUrl) {
		toSerialize["document_url"] = o.DocumentUrl
	}
	return json.Marshal(toSerialize)
}

type NullableClassifyDocumentRequest struct {
	value *ClassifyDocumentRequest
	isSet bool
}

func (v NullableClassifyDocumentRequest) Get() *ClassifyDocumentRequest {
	return v.value
}

func (v *NullableClassifyDocumentRequest) Set(val *ClassifyDocumentRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableClassifyDocumentRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableClassifyDocumentRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClassifyDocumentRequest(val *ClassifyDocumentRequest) *NullableClassifyDocumentRequest {
	return &NullableClassifyDocumentRequest{value: val, isSet: true}
}

func (v NullableClassifyDocumentRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClassifyDocumentRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


