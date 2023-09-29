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

// DigilockerDownloadFileRequest struct for DigilockerDownloadFileRequest
type DigilockerDownloadFileRequest struct {
	FileUrn *string `json:"file_urn,omitempty"`
	Consent *bool `json:"consent,omitempty"`
	Purpose *string `json:"purpose,omitempty"`
	ReferenceId *string `json:"reference_id,omitempty"`
}

// NewDigilockerDownloadFileRequest instantiates a new DigilockerDownloadFileRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDigilockerDownloadFileRequest() *DigilockerDownloadFileRequest {
	this := DigilockerDownloadFileRequest{}
	return &this
}

// NewDigilockerDownloadFileRequestWithDefaults instantiates a new DigilockerDownloadFileRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDigilockerDownloadFileRequestWithDefaults() *DigilockerDownloadFileRequest {
	this := DigilockerDownloadFileRequest{}
	return &this
}

// GetFileUrn returns the FileUrn field value if set, zero value otherwise.
func (o *DigilockerDownloadFileRequest) GetFileUrn() string {
	if o == nil || isNil(o.FileUrn) {
		var ret string
		return ret
	}
	return *o.FileUrn
}

// GetFileUrnOk returns a tuple with the FileUrn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DigilockerDownloadFileRequest) GetFileUrnOk() (*string, bool) {
	if o == nil || isNil(o.FileUrn) {
    return nil, false
	}
	return o.FileUrn, true
}

// HasFileUrn returns a boolean if a field has been set.
func (o *DigilockerDownloadFileRequest) HasFileUrn() bool {
	if o != nil && !isNil(o.FileUrn) {
		return true
	}

	return false
}

// SetFileUrn gets a reference to the given string and assigns it to the FileUrn field.
func (o *DigilockerDownloadFileRequest) SetFileUrn(v string) {
	o.FileUrn = &v
}

// GetConsent returns the Consent field value if set, zero value otherwise.
func (o *DigilockerDownloadFileRequest) GetConsent() bool {
	if o == nil || isNil(o.Consent) {
		var ret bool
		return ret
	}
	return *o.Consent
}

// GetConsentOk returns a tuple with the Consent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DigilockerDownloadFileRequest) GetConsentOk() (*bool, bool) {
	if o == nil || isNil(o.Consent) {
    return nil, false
	}
	return o.Consent, true
}

// HasConsent returns a boolean if a field has been set.
func (o *DigilockerDownloadFileRequest) HasConsent() bool {
	if o != nil && !isNil(o.Consent) {
		return true
	}

	return false
}

// SetConsent gets a reference to the given bool and assigns it to the Consent field.
func (o *DigilockerDownloadFileRequest) SetConsent(v bool) {
	o.Consent = &v
}

// GetPurpose returns the Purpose field value if set, zero value otherwise.
func (o *DigilockerDownloadFileRequest) GetPurpose() string {
	if o == nil || isNil(o.Purpose) {
		var ret string
		return ret
	}
	return *o.Purpose
}

// GetPurposeOk returns a tuple with the Purpose field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DigilockerDownloadFileRequest) GetPurposeOk() (*string, bool) {
	if o == nil || isNil(o.Purpose) {
    return nil, false
	}
	return o.Purpose, true
}

// HasPurpose returns a boolean if a field has been set.
func (o *DigilockerDownloadFileRequest) HasPurpose() bool {
	if o != nil && !isNil(o.Purpose) {
		return true
	}

	return false
}

// SetPurpose gets a reference to the given string and assigns it to the Purpose field.
func (o *DigilockerDownloadFileRequest) SetPurpose(v string) {
	o.Purpose = &v
}

// GetReferenceId returns the ReferenceId field value if set, zero value otherwise.
func (o *DigilockerDownloadFileRequest) GetReferenceId() string {
	if o == nil || isNil(o.ReferenceId) {
		var ret string
		return ret
	}
	return *o.ReferenceId
}

// GetReferenceIdOk returns a tuple with the ReferenceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DigilockerDownloadFileRequest) GetReferenceIdOk() (*string, bool) {
	if o == nil || isNil(o.ReferenceId) {
    return nil, false
	}
	return o.ReferenceId, true
}

// HasReferenceId returns a boolean if a field has been set.
func (o *DigilockerDownloadFileRequest) HasReferenceId() bool {
	if o != nil && !isNil(o.ReferenceId) {
		return true
	}

	return false
}

// SetReferenceId gets a reference to the given string and assigns it to the ReferenceId field.
func (o *DigilockerDownloadFileRequest) SetReferenceId(v string) {
	o.ReferenceId = &v
}

func (o DigilockerDownloadFileRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.FileUrn) {
		toSerialize["file_urn"] = o.FileUrn
	}
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

type NullableDigilockerDownloadFileRequest struct {
	value *DigilockerDownloadFileRequest
	isSet bool
}

func (v NullableDigilockerDownloadFileRequest) Get() *DigilockerDownloadFileRequest {
	return v.value
}

func (v *NullableDigilockerDownloadFileRequest) Set(val *DigilockerDownloadFileRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableDigilockerDownloadFileRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableDigilockerDownloadFileRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDigilockerDownloadFileRequest(val *DigilockerDownloadFileRequest) *NullableDigilockerDownloadFileRequest {
	return &NullableDigilockerDownloadFileRequest{value: val, isSet: true}
}

func (v NullableDigilockerDownloadFileRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDigilockerDownloadFileRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


