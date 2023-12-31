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

// CheckImageQualityRequest struct for CheckImageQualityRequest
type CheckImageQualityRequest struct {
	ReferenceId string `json:"reference_id"`
	Consent bool `json:"consent"`
	ConsentPurpose string `json:"consent_purpose"`
	Image **os.File `json:"image,omitempty"`
	QualityParameter *string `json:"quality_parameter,omitempty"`
	ImageUrl *string `json:"image_url,omitempty"`
}

// NewCheckImageQualityRequest instantiates a new CheckImageQualityRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCheckImageQualityRequest(referenceId string, consent bool, consentPurpose string) *CheckImageQualityRequest {
	this := CheckImageQualityRequest{}
	this.ReferenceId = referenceId
	this.Consent = consent
	this.ConsentPurpose = consentPurpose
	return &this
}

// NewCheckImageQualityRequestWithDefaults instantiates a new CheckImageQualityRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCheckImageQualityRequestWithDefaults() *CheckImageQualityRequest {
	this := CheckImageQualityRequest{}
	return &this
}

// GetReferenceId returns the ReferenceId field value
func (o *CheckImageQualityRequest) GetReferenceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ReferenceId
}

// GetReferenceIdOk returns a tuple with the ReferenceId field value
// and a boolean to check if the value has been set.
func (o *CheckImageQualityRequest) GetReferenceIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.ReferenceId, true
}

// SetReferenceId sets field value
func (o *CheckImageQualityRequest) SetReferenceId(v string) {
	o.ReferenceId = v
}

// GetConsent returns the Consent field value
func (o *CheckImageQualityRequest) GetConsent() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Consent
}

// GetConsentOk returns a tuple with the Consent field value
// and a boolean to check if the value has been set.
func (o *CheckImageQualityRequest) GetConsentOk() (*bool, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Consent, true
}

// SetConsent sets field value
func (o *CheckImageQualityRequest) SetConsent(v bool) {
	o.Consent = v
}

// GetConsentPurpose returns the ConsentPurpose field value
func (o *CheckImageQualityRequest) GetConsentPurpose() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ConsentPurpose
}

// GetConsentPurposeOk returns a tuple with the ConsentPurpose field value
// and a boolean to check if the value has been set.
func (o *CheckImageQualityRequest) GetConsentPurposeOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.ConsentPurpose, true
}

// SetConsentPurpose sets field value
func (o *CheckImageQualityRequest) SetConsentPurpose(v string) {
	o.ConsentPurpose = v
}

// GetImage returns the Image field value if set, zero value otherwise.
func (o *CheckImageQualityRequest) GetImage() *os.File {
	if o == nil || isNil(o.Image) {
		var ret *os.File
		return ret
	}
	return *o.Image
}

// GetImageOk returns a tuple with the Image field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckImageQualityRequest) GetImageOk() (**os.File, bool) {
	if o == nil || isNil(o.Image) {
    return nil, false
	}
	return o.Image, true
}

// HasImage returns a boolean if a field has been set.
func (o *CheckImageQualityRequest) HasImage() bool {
	if o != nil && !isNil(o.Image) {
		return true
	}

	return false
}

// SetImage gets a reference to the given *os.File and assigns it to the Image field.
func (o *CheckImageQualityRequest) SetImage(v *os.File) {
	o.Image = &v
}

// GetQualityParameter returns the QualityParameter field value if set, zero value otherwise.
func (o *CheckImageQualityRequest) GetQualityParameter() string {
	if o == nil || isNil(o.QualityParameter) {
		var ret string
		return ret
	}
	return *o.QualityParameter
}

// GetQualityParameterOk returns a tuple with the QualityParameter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckImageQualityRequest) GetQualityParameterOk() (*string, bool) {
	if o == nil || isNil(o.QualityParameter) {
    return nil, false
	}
	return o.QualityParameter, true
}

// HasQualityParameter returns a boolean if a field has been set.
func (o *CheckImageQualityRequest) HasQualityParameter() bool {
	if o != nil && !isNil(o.QualityParameter) {
		return true
	}

	return false
}

// SetQualityParameter gets a reference to the given string and assigns it to the QualityParameter field.
func (o *CheckImageQualityRequest) SetQualityParameter(v string) {
	o.QualityParameter = &v
}

// GetImageUrl returns the ImageUrl field value if set, zero value otherwise.
func (o *CheckImageQualityRequest) GetImageUrl() string {
	if o == nil || isNil(o.ImageUrl) {
		var ret string
		return ret
	}
	return *o.ImageUrl
}

// GetImageUrlOk returns a tuple with the ImageUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckImageQualityRequest) GetImageUrlOk() (*string, bool) {
	if o == nil || isNil(o.ImageUrl) {
    return nil, false
	}
	return o.ImageUrl, true
}

// HasImageUrl returns a boolean if a field has been set.
func (o *CheckImageQualityRequest) HasImageUrl() bool {
	if o != nil && !isNil(o.ImageUrl) {
		return true
	}

	return false
}

// SetImageUrl gets a reference to the given string and assigns it to the ImageUrl field.
func (o *CheckImageQualityRequest) SetImageUrl(v string) {
	o.ImageUrl = &v
}

func (o CheckImageQualityRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["reference_id"] = o.ReferenceId
	}
	if true {
		toSerialize["consent"] = o.Consent
	}
	if true {
		toSerialize["consent_purpose"] = o.ConsentPurpose
	}
	if !isNil(o.Image) {
		toSerialize["image"] = o.Image
	}
	if !isNil(o.QualityParameter) {
		toSerialize["quality_parameter"] = o.QualityParameter
	}
	if !isNil(o.ImageUrl) {
		toSerialize["image_url"] = o.ImageUrl
	}
	return json.Marshal(toSerialize)
}

type NullableCheckImageQualityRequest struct {
	value *CheckImageQualityRequest
	isSet bool
}

func (v NullableCheckImageQualityRequest) Get() *CheckImageQualityRequest {
	return v.value
}

func (v *NullableCheckImageQualityRequest) Set(val *CheckImageQualityRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCheckImageQualityRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCheckImageQualityRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCheckImageQualityRequest(val *CheckImageQualityRequest) *NullableCheckImageQualityRequest {
	return &NullableCheckImageQualityRequest{value: val, isSet: true}
}

func (v NullableCheckImageQualityRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCheckImageQualityRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


