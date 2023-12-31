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

// MatchFaceRequest struct for MatchFaceRequest
type MatchFaceRequest struct {
	// 
	ReferenceId string `json:"reference_id"`
	// 
	Consent string `json:"consent"`
	// 
	ConsentPurpose string `json:"consent_purpose"`
	Image1 **os.File `json:"image1,omitempty"`
	Image2 **os.File `json:"image2,omitempty"`
	Threshold *int32 `json:"threshold,omitempty"`
	Image1Url *string `json:"image1_url,omitempty"`
	Image2Url *string `json:"image2_url,omitempty"`
}

// NewMatchFaceRequest instantiates a new MatchFaceRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMatchFaceRequest(referenceId string, consent string, consentPurpose string) *MatchFaceRequest {
	this := MatchFaceRequest{}
	this.ReferenceId = referenceId
	this.Consent = consent
	this.ConsentPurpose = consentPurpose
	return &this
}

// NewMatchFaceRequestWithDefaults instantiates a new MatchFaceRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMatchFaceRequestWithDefaults() *MatchFaceRequest {
	this := MatchFaceRequest{}
	return &this
}

// GetReferenceId returns the ReferenceId field value
func (o *MatchFaceRequest) GetReferenceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ReferenceId
}

// GetReferenceIdOk returns a tuple with the ReferenceId field value
// and a boolean to check if the value has been set.
func (o *MatchFaceRequest) GetReferenceIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.ReferenceId, true
}

// SetReferenceId sets field value
func (o *MatchFaceRequest) SetReferenceId(v string) {
	o.ReferenceId = v
}

// GetConsent returns the Consent field value
func (o *MatchFaceRequest) GetConsent() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Consent
}

// GetConsentOk returns a tuple with the Consent field value
// and a boolean to check if the value has been set.
func (o *MatchFaceRequest) GetConsentOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Consent, true
}

// SetConsent sets field value
func (o *MatchFaceRequest) SetConsent(v string) {
	o.Consent = v
}

// GetConsentPurpose returns the ConsentPurpose field value
func (o *MatchFaceRequest) GetConsentPurpose() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ConsentPurpose
}

// GetConsentPurposeOk returns a tuple with the ConsentPurpose field value
// and a boolean to check if the value has been set.
func (o *MatchFaceRequest) GetConsentPurposeOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.ConsentPurpose, true
}

// SetConsentPurpose sets field value
func (o *MatchFaceRequest) SetConsentPurpose(v string) {
	o.ConsentPurpose = v
}

// GetImage1 returns the Image1 field value if set, zero value otherwise.
func (o *MatchFaceRequest) GetImage1() *os.File {
	if o == nil || isNil(o.Image1) {
		var ret *os.File
		return ret
	}
	return *o.Image1
}

// GetImage1Ok returns a tuple with the Image1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MatchFaceRequest) GetImage1Ok() (**os.File, bool) {
	if o == nil || isNil(o.Image1) {
    return nil, false
	}
	return o.Image1, true
}

// HasImage1 returns a boolean if a field has been set.
func (o *MatchFaceRequest) HasImage1() bool {
	if o != nil && !isNil(o.Image1) {
		return true
	}

	return false
}

// SetImage1 gets a reference to the given *os.File and assigns it to the Image1 field.
func (o *MatchFaceRequest) SetImage1(v *os.File) {
	o.Image1 = &v
}

// GetImage2 returns the Image2 field value if set, zero value otherwise.
func (o *MatchFaceRequest) GetImage2() *os.File {
	if o == nil || isNil(o.Image2) {
		var ret *os.File
		return ret
	}
	return *o.Image2
}

// GetImage2Ok returns a tuple with the Image2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MatchFaceRequest) GetImage2Ok() (**os.File, bool) {
	if o == nil || isNil(o.Image2) {
    return nil, false
	}
	return o.Image2, true
}

// HasImage2 returns a boolean if a field has been set.
func (o *MatchFaceRequest) HasImage2() bool {
	if o != nil && !isNil(o.Image2) {
		return true
	}

	return false
}

// SetImage2 gets a reference to the given *os.File and assigns it to the Image2 field.
func (o *MatchFaceRequest) SetImage2(v *os.File) {
	o.Image2 = &v
}

// GetThreshold returns the Threshold field value if set, zero value otherwise.
func (o *MatchFaceRequest) GetThreshold() int32 {
	if o == nil || isNil(o.Threshold) {
		var ret int32
		return ret
	}
	return *o.Threshold
}

// GetThresholdOk returns a tuple with the Threshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MatchFaceRequest) GetThresholdOk() (*int32, bool) {
	if o == nil || isNil(o.Threshold) {
    return nil, false
	}
	return o.Threshold, true
}

// HasThreshold returns a boolean if a field has been set.
func (o *MatchFaceRequest) HasThreshold() bool {
	if o != nil && !isNil(o.Threshold) {
		return true
	}

	return false
}

// SetThreshold gets a reference to the given int32 and assigns it to the Threshold field.
func (o *MatchFaceRequest) SetThreshold(v int32) {
	o.Threshold = &v
}

// GetImage1Url returns the Image1Url field value if set, zero value otherwise.
func (o *MatchFaceRequest) GetImage1Url() string {
	if o == nil || isNil(o.Image1Url) {
		var ret string
		return ret
	}
	return *o.Image1Url
}

// GetImage1UrlOk returns a tuple with the Image1Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MatchFaceRequest) GetImage1UrlOk() (*string, bool) {
	if o == nil || isNil(o.Image1Url) {
    return nil, false
	}
	return o.Image1Url, true
}

// HasImage1Url returns a boolean if a field has been set.
func (o *MatchFaceRequest) HasImage1Url() bool {
	if o != nil && !isNil(o.Image1Url) {
		return true
	}

	return false
}

// SetImage1Url gets a reference to the given string and assigns it to the Image1Url field.
func (o *MatchFaceRequest) SetImage1Url(v string) {
	o.Image1Url = &v
}

// GetImage2Url returns the Image2Url field value if set, zero value otherwise.
func (o *MatchFaceRequest) GetImage2Url() string {
	if o == nil || isNil(o.Image2Url) {
		var ret string
		return ret
	}
	return *o.Image2Url
}

// GetImage2UrlOk returns a tuple with the Image2Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MatchFaceRequest) GetImage2UrlOk() (*string, bool) {
	if o == nil || isNil(o.Image2Url) {
    return nil, false
	}
	return o.Image2Url, true
}

// HasImage2Url returns a boolean if a field has been set.
func (o *MatchFaceRequest) HasImage2Url() bool {
	if o != nil && !isNil(o.Image2Url) {
		return true
	}

	return false
}

// SetImage2Url gets a reference to the given string and assigns it to the Image2Url field.
func (o *MatchFaceRequest) SetImage2Url(v string) {
	o.Image2Url = &v
}

func (o MatchFaceRequest) MarshalJSON() ([]byte, error) {
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
	if !isNil(o.Image1) {
		toSerialize["image1"] = o.Image1
	}
	if !isNil(o.Image2) {
		toSerialize["image2"] = o.Image2
	}
	if !isNil(o.Threshold) {
		toSerialize["threshold"] = o.Threshold
	}
	if !isNil(o.Image1Url) {
		toSerialize["image1_url"] = o.Image1Url
	}
	if !isNil(o.Image2Url) {
		toSerialize["image2_url"] = o.Image2Url
	}
	return json.Marshal(toSerialize)
}

type NullableMatchFaceRequest struct {
	value *MatchFaceRequest
	isSet bool
}

func (v NullableMatchFaceRequest) Get() *MatchFaceRequest {
	return v.value
}

func (v *NullableMatchFaceRequest) Set(val *MatchFaceRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableMatchFaceRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableMatchFaceRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMatchFaceRequest(val *MatchFaceRequest) *NullableMatchFaceRequest {
	return &NullableMatchFaceRequest{value: val, isSet: true}
}

func (v NullableMatchFaceRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMatchFaceRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


