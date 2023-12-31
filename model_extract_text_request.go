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

// ExtractTextRequest struct for ExtractTextRequest
type ExtractTextRequest struct {
	// 
	ReferenceId string `json:"reference_id"`
	// 
	DocumentType string `json:"document_type"`
	// 
	Consent string `json:"consent"`
	// 
	ConsentPurpose string `json:"consent_purpose"`
	// 
	KycValidate int32 `json:"kyc_validate"`
	Document **os.File `json:"document,omitempty"`
	DocumentUrl *string `json:"document_url,omitempty"`
	DocumentBack **os.File `json:"document_back,omitempty"`
	DocumentBackUrl *string `json:"document_back_url,omitempty"`
}

// NewExtractTextRequest instantiates a new ExtractTextRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExtractTextRequest(referenceId string, documentType string, consent string, consentPurpose string, kycValidate int32) *ExtractTextRequest {
	this := ExtractTextRequest{}
	this.ReferenceId = referenceId
	this.DocumentType = documentType
	this.Consent = consent
	this.ConsentPurpose = consentPurpose
	this.KycValidate = kycValidate
	return &this
}

// NewExtractTextRequestWithDefaults instantiates a new ExtractTextRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExtractTextRequestWithDefaults() *ExtractTextRequest {
	this := ExtractTextRequest{}
	return &this
}

// GetReferenceId returns the ReferenceId field value
func (o *ExtractTextRequest) GetReferenceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ReferenceId
}

// GetReferenceIdOk returns a tuple with the ReferenceId field value
// and a boolean to check if the value has been set.
func (o *ExtractTextRequest) GetReferenceIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.ReferenceId, true
}

// SetReferenceId sets field value
func (o *ExtractTextRequest) SetReferenceId(v string) {
	o.ReferenceId = v
}

// GetDocumentType returns the DocumentType field value
func (o *ExtractTextRequest) GetDocumentType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DocumentType
}

// GetDocumentTypeOk returns a tuple with the DocumentType field value
// and a boolean to check if the value has been set.
func (o *ExtractTextRequest) GetDocumentTypeOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.DocumentType, true
}

// SetDocumentType sets field value
func (o *ExtractTextRequest) SetDocumentType(v string) {
	o.DocumentType = v
}

// GetConsent returns the Consent field value
func (o *ExtractTextRequest) GetConsent() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Consent
}

// GetConsentOk returns a tuple with the Consent field value
// and a boolean to check if the value has been set.
func (o *ExtractTextRequest) GetConsentOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Consent, true
}

// SetConsent sets field value
func (o *ExtractTextRequest) SetConsent(v string) {
	o.Consent = v
}

// GetConsentPurpose returns the ConsentPurpose field value
func (o *ExtractTextRequest) GetConsentPurpose() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ConsentPurpose
}

// GetConsentPurposeOk returns a tuple with the ConsentPurpose field value
// and a boolean to check if the value has been set.
func (o *ExtractTextRequest) GetConsentPurposeOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.ConsentPurpose, true
}

// SetConsentPurpose sets field value
func (o *ExtractTextRequest) SetConsentPurpose(v string) {
	o.ConsentPurpose = v
}

// GetKycValidate returns the KycValidate field value
func (o *ExtractTextRequest) GetKycValidate() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.KycValidate
}

// GetKycValidateOk returns a tuple with the KycValidate field value
// and a boolean to check if the value has been set.
func (o *ExtractTextRequest) GetKycValidateOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return &o.KycValidate, true
}

// SetKycValidate sets field value
func (o *ExtractTextRequest) SetKycValidate(v int32) {
	o.KycValidate = v
}

// GetDocument returns the Document field value if set, zero value otherwise.
func (o *ExtractTextRequest) GetDocument() *os.File {
	if o == nil || isNil(o.Document) {
		var ret *os.File
		return ret
	}
	return *o.Document
}

// GetDocumentOk returns a tuple with the Document field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExtractTextRequest) GetDocumentOk() (**os.File, bool) {
	if o == nil || isNil(o.Document) {
    return nil, false
	}
	return o.Document, true
}

// HasDocument returns a boolean if a field has been set.
func (o *ExtractTextRequest) HasDocument() bool {
	if o != nil && !isNil(o.Document) {
		return true
	}

	return false
}

// SetDocument gets a reference to the given *os.File and assigns it to the Document field.
func (o *ExtractTextRequest) SetDocument(v *os.File) {
	o.Document = &v
}

// GetDocumentUrl returns the DocumentUrl field value if set, zero value otherwise.
func (o *ExtractTextRequest) GetDocumentUrl() string {
	if o == nil || isNil(o.DocumentUrl) {
		var ret string
		return ret
	}
	return *o.DocumentUrl
}

// GetDocumentUrlOk returns a tuple with the DocumentUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExtractTextRequest) GetDocumentUrlOk() (*string, bool) {
	if o == nil || isNil(o.DocumentUrl) {
    return nil, false
	}
	return o.DocumentUrl, true
}

// HasDocumentUrl returns a boolean if a field has been set.
func (o *ExtractTextRequest) HasDocumentUrl() bool {
	if o != nil && !isNil(o.DocumentUrl) {
		return true
	}

	return false
}

// SetDocumentUrl gets a reference to the given string and assigns it to the DocumentUrl field.
func (o *ExtractTextRequest) SetDocumentUrl(v string) {
	o.DocumentUrl = &v
}

// GetDocumentBack returns the DocumentBack field value if set, zero value otherwise.
func (o *ExtractTextRequest) GetDocumentBack() *os.File {
	if o == nil || isNil(o.DocumentBack) {
		var ret *os.File
		return ret
	}
	return *o.DocumentBack
}

// GetDocumentBackOk returns a tuple with the DocumentBack field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExtractTextRequest) GetDocumentBackOk() (**os.File, bool) {
	if o == nil || isNil(o.DocumentBack) {
    return nil, false
	}
	return o.DocumentBack, true
}

// HasDocumentBack returns a boolean if a field has been set.
func (o *ExtractTextRequest) HasDocumentBack() bool {
	if o != nil && !isNil(o.DocumentBack) {
		return true
	}

	return false
}

// SetDocumentBack gets a reference to the given *os.File and assigns it to the DocumentBack field.
func (o *ExtractTextRequest) SetDocumentBack(v *os.File) {
	o.DocumentBack = &v
}

// GetDocumentBackUrl returns the DocumentBackUrl field value if set, zero value otherwise.
func (o *ExtractTextRequest) GetDocumentBackUrl() string {
	if o == nil || isNil(o.DocumentBackUrl) {
		var ret string
		return ret
	}
	return *o.DocumentBackUrl
}

// GetDocumentBackUrlOk returns a tuple with the DocumentBackUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExtractTextRequest) GetDocumentBackUrlOk() (*string, bool) {
	if o == nil || isNil(o.DocumentBackUrl) {
    return nil, false
	}
	return o.DocumentBackUrl, true
}

// HasDocumentBackUrl returns a boolean if a field has been set.
func (o *ExtractTextRequest) HasDocumentBackUrl() bool {
	if o != nil && !isNil(o.DocumentBackUrl) {
		return true
	}

	return false
}

// SetDocumentBackUrl gets a reference to the given string and assigns it to the DocumentBackUrl field.
func (o *ExtractTextRequest) SetDocumentBackUrl(v string) {
	o.DocumentBackUrl = &v
}

func (o ExtractTextRequest) MarshalJSON() ([]byte, error) {
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
	if true {
		toSerialize["kyc_validate"] = o.KycValidate
	}
	if !isNil(o.Document) {
		toSerialize["document"] = o.Document
	}
	if !isNil(o.DocumentUrl) {
		toSerialize["document_url"] = o.DocumentUrl
	}
	if !isNil(o.DocumentBack) {
		toSerialize["document_back"] = o.DocumentBack
	}
	if !isNil(o.DocumentBackUrl) {
		toSerialize["document_back_url"] = o.DocumentBackUrl
	}
	return json.Marshal(toSerialize)
}

type NullableExtractTextRequest struct {
	value *ExtractTextRequest
	isSet bool
}

func (v NullableExtractTextRequest) Get() *ExtractTextRequest {
	return v.value
}

func (v *NullableExtractTextRequest) Set(val *ExtractTextRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableExtractTextRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableExtractTextRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExtractTextRequest(val *ExtractTextRequest) *NullableExtractTextRequest {
	return &NullableExtractTextRequest{value: val, isSet: true}
}

func (v NullableExtractTextRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExtractTextRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


