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

// ClassifyDocumentResponseData struct for ClassifyDocumentResponseData
type ClassifyDocumentResponseData struct {
	DocumentClassificationStatus *string `json:"documentClassificationStatus,omitempty"`
	IdTypeClassified *string `json:"idTypeClassified,omitempty"`
}

// NewClassifyDocumentResponseData instantiates a new ClassifyDocumentResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClassifyDocumentResponseData() *ClassifyDocumentResponseData {
	this := ClassifyDocumentResponseData{}
	return &this
}

// NewClassifyDocumentResponseDataWithDefaults instantiates a new ClassifyDocumentResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClassifyDocumentResponseDataWithDefaults() *ClassifyDocumentResponseData {
	this := ClassifyDocumentResponseData{}
	return &this
}

// GetDocumentClassificationStatus returns the DocumentClassificationStatus field value if set, zero value otherwise.
func (o *ClassifyDocumentResponseData) GetDocumentClassificationStatus() string {
	if o == nil || isNil(o.DocumentClassificationStatus) {
		var ret string
		return ret
	}
	return *o.DocumentClassificationStatus
}

// GetDocumentClassificationStatusOk returns a tuple with the DocumentClassificationStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClassifyDocumentResponseData) GetDocumentClassificationStatusOk() (*string, bool) {
	if o == nil || isNil(o.DocumentClassificationStatus) {
    return nil, false
	}
	return o.DocumentClassificationStatus, true
}

// HasDocumentClassificationStatus returns a boolean if a field has been set.
func (o *ClassifyDocumentResponseData) HasDocumentClassificationStatus() bool {
	if o != nil && !isNil(o.DocumentClassificationStatus) {
		return true
	}

	return false
}

// SetDocumentClassificationStatus gets a reference to the given string and assigns it to the DocumentClassificationStatus field.
func (o *ClassifyDocumentResponseData) SetDocumentClassificationStatus(v string) {
	o.DocumentClassificationStatus = &v
}

// GetIdTypeClassified returns the IdTypeClassified field value if set, zero value otherwise.
func (o *ClassifyDocumentResponseData) GetIdTypeClassified() string {
	if o == nil || isNil(o.IdTypeClassified) {
		var ret string
		return ret
	}
	return *o.IdTypeClassified
}

// GetIdTypeClassifiedOk returns a tuple with the IdTypeClassified field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClassifyDocumentResponseData) GetIdTypeClassifiedOk() (*string, bool) {
	if o == nil || isNil(o.IdTypeClassified) {
    return nil, false
	}
	return o.IdTypeClassified, true
}

// HasIdTypeClassified returns a boolean if a field has been set.
func (o *ClassifyDocumentResponseData) HasIdTypeClassified() bool {
	if o != nil && !isNil(o.IdTypeClassified) {
		return true
	}

	return false
}

// SetIdTypeClassified gets a reference to the given string and assigns it to the IdTypeClassified field.
func (o *ClassifyDocumentResponseData) SetIdTypeClassified(v string) {
	o.IdTypeClassified = &v
}

func (o ClassifyDocumentResponseData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.DocumentClassificationStatus) {
		toSerialize["documentClassificationStatus"] = o.DocumentClassificationStatus
	}
	if !isNil(o.IdTypeClassified) {
		toSerialize["idTypeClassified"] = o.IdTypeClassified
	}
	return json.Marshal(toSerialize)
}

type NullableClassifyDocumentResponseData struct {
	value *ClassifyDocumentResponseData
	isSet bool
}

func (v NullableClassifyDocumentResponseData) Get() *ClassifyDocumentResponseData {
	return v.value
}

func (v *NullableClassifyDocumentResponseData) Set(val *ClassifyDocumentResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableClassifyDocumentResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableClassifyDocumentResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClassifyDocumentResponseData(val *ClassifyDocumentResponseData) *NullableClassifyDocumentResponseData {
	return &NullableClassifyDocumentResponseData{value: val, isSet: true}
}

func (v NullableClassifyDocumentResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClassifyDocumentResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


