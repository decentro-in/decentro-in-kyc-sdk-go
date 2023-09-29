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

// CheckVideoLivenessRequest struct for CheckVideoLivenessRequest
type CheckVideoLivenessRequest struct {
	// 
	ReferenceId string `json:"reference_id"`
	// 
	Consent string `json:"consent"`
	// 
	ConsentPurpose string `json:"consent_purpose"`
	Video **os.File `json:"video,omitempty"`
	VideoUrl *string `json:"video_url,omitempty"`
}

// NewCheckVideoLivenessRequest instantiates a new CheckVideoLivenessRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCheckVideoLivenessRequest(referenceId string, consent string, consentPurpose string) *CheckVideoLivenessRequest {
	this := CheckVideoLivenessRequest{}
	this.ReferenceId = referenceId
	this.Consent = consent
	this.ConsentPurpose = consentPurpose
	return &this
}

// NewCheckVideoLivenessRequestWithDefaults instantiates a new CheckVideoLivenessRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCheckVideoLivenessRequestWithDefaults() *CheckVideoLivenessRequest {
	this := CheckVideoLivenessRequest{}
	return &this
}

// GetReferenceId returns the ReferenceId field value
func (o *CheckVideoLivenessRequest) GetReferenceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ReferenceId
}

// GetReferenceIdOk returns a tuple with the ReferenceId field value
// and a boolean to check if the value has been set.
func (o *CheckVideoLivenessRequest) GetReferenceIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.ReferenceId, true
}

// SetReferenceId sets field value
func (o *CheckVideoLivenessRequest) SetReferenceId(v string) {
	o.ReferenceId = v
}

// GetConsent returns the Consent field value
func (o *CheckVideoLivenessRequest) GetConsent() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Consent
}

// GetConsentOk returns a tuple with the Consent field value
// and a boolean to check if the value has been set.
func (o *CheckVideoLivenessRequest) GetConsentOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Consent, true
}

// SetConsent sets field value
func (o *CheckVideoLivenessRequest) SetConsent(v string) {
	o.Consent = v
}

// GetConsentPurpose returns the ConsentPurpose field value
func (o *CheckVideoLivenessRequest) GetConsentPurpose() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ConsentPurpose
}

// GetConsentPurposeOk returns a tuple with the ConsentPurpose field value
// and a boolean to check if the value has been set.
func (o *CheckVideoLivenessRequest) GetConsentPurposeOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.ConsentPurpose, true
}

// SetConsentPurpose sets field value
func (o *CheckVideoLivenessRequest) SetConsentPurpose(v string) {
	o.ConsentPurpose = v
}

// GetVideo returns the Video field value if set, zero value otherwise.
func (o *CheckVideoLivenessRequest) GetVideo() *os.File {
	if o == nil || isNil(o.Video) {
		var ret *os.File
		return ret
	}
	return *o.Video
}

// GetVideoOk returns a tuple with the Video field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckVideoLivenessRequest) GetVideoOk() (**os.File, bool) {
	if o == nil || isNil(o.Video) {
    return nil, false
	}
	return o.Video, true
}

// HasVideo returns a boolean if a field has been set.
func (o *CheckVideoLivenessRequest) HasVideo() bool {
	if o != nil && !isNil(o.Video) {
		return true
	}

	return false
}

// SetVideo gets a reference to the given *os.File and assigns it to the Video field.
func (o *CheckVideoLivenessRequest) SetVideo(v *os.File) {
	o.Video = &v
}

// GetVideoUrl returns the VideoUrl field value if set, zero value otherwise.
func (o *CheckVideoLivenessRequest) GetVideoUrl() string {
	if o == nil || isNil(o.VideoUrl) {
		var ret string
		return ret
	}
	return *o.VideoUrl
}

// GetVideoUrlOk returns a tuple with the VideoUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckVideoLivenessRequest) GetVideoUrlOk() (*string, bool) {
	if o == nil || isNil(o.VideoUrl) {
    return nil, false
	}
	return o.VideoUrl, true
}

// HasVideoUrl returns a boolean if a field has been set.
func (o *CheckVideoLivenessRequest) HasVideoUrl() bool {
	if o != nil && !isNil(o.VideoUrl) {
		return true
	}

	return false
}

// SetVideoUrl gets a reference to the given string and assigns it to the VideoUrl field.
func (o *CheckVideoLivenessRequest) SetVideoUrl(v string) {
	o.VideoUrl = &v
}

func (o CheckVideoLivenessRequest) MarshalJSON() ([]byte, error) {
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
	if !isNil(o.Video) {
		toSerialize["video"] = o.Video
	}
	if !isNil(o.VideoUrl) {
		toSerialize["video_url"] = o.VideoUrl
	}
	return json.Marshal(toSerialize)
}

type NullableCheckVideoLivenessRequest struct {
	value *CheckVideoLivenessRequest
	isSet bool
}

func (v NullableCheckVideoLivenessRequest) Get() *CheckVideoLivenessRequest {
	return v.value
}

func (v *NullableCheckVideoLivenessRequest) Set(val *CheckVideoLivenessRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCheckVideoLivenessRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCheckVideoLivenessRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCheckVideoLivenessRequest(val *CheckVideoLivenessRequest) *NullableCheckVideoLivenessRequest {
	return &NullableCheckVideoLivenessRequest{value: val, isSet: true}
}

func (v NullableCheckVideoLivenessRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCheckVideoLivenessRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


