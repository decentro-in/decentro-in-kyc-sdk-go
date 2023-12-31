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

// CheckImageQualityResponseDataImageQuality struct for CheckImageQualityResponseDataImageQuality
type CheckImageQualityResponseDataImageQuality struct {
	Summary *string `json:"summary,omitempty"`
	QualityScores *CheckImageQualityResponseDataImageQualityQualityScores `json:"qualityScores,omitempty"`
	ExtractionQuality *string `json:"extractionQuality,omitempty"`
	Score *float32 `json:"score,omitempty"`
	Msg *string `json:"msg,omitempty"`
}

// NewCheckImageQualityResponseDataImageQuality instantiates a new CheckImageQualityResponseDataImageQuality object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCheckImageQualityResponseDataImageQuality() *CheckImageQualityResponseDataImageQuality {
	this := CheckImageQualityResponseDataImageQuality{}
	return &this
}

// NewCheckImageQualityResponseDataImageQualityWithDefaults instantiates a new CheckImageQualityResponseDataImageQuality object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCheckImageQualityResponseDataImageQualityWithDefaults() *CheckImageQualityResponseDataImageQuality {
	this := CheckImageQualityResponseDataImageQuality{}
	return &this
}

// GetSummary returns the Summary field value if set, zero value otherwise.
func (o *CheckImageQualityResponseDataImageQuality) GetSummary() string {
	if o == nil || isNil(o.Summary) {
		var ret string
		return ret
	}
	return *o.Summary
}

// GetSummaryOk returns a tuple with the Summary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckImageQualityResponseDataImageQuality) GetSummaryOk() (*string, bool) {
	if o == nil || isNil(o.Summary) {
    return nil, false
	}
	return o.Summary, true
}

// HasSummary returns a boolean if a field has been set.
func (o *CheckImageQualityResponseDataImageQuality) HasSummary() bool {
	if o != nil && !isNil(o.Summary) {
		return true
	}

	return false
}

// SetSummary gets a reference to the given string and assigns it to the Summary field.
func (o *CheckImageQualityResponseDataImageQuality) SetSummary(v string) {
	o.Summary = &v
}

// GetQualityScores returns the QualityScores field value if set, zero value otherwise.
func (o *CheckImageQualityResponseDataImageQuality) GetQualityScores() CheckImageQualityResponseDataImageQualityQualityScores {
	if o == nil || isNil(o.QualityScores) {
		var ret CheckImageQualityResponseDataImageQualityQualityScores
		return ret
	}
	return *o.QualityScores
}

// GetQualityScoresOk returns a tuple with the QualityScores field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckImageQualityResponseDataImageQuality) GetQualityScoresOk() (*CheckImageQualityResponseDataImageQualityQualityScores, bool) {
	if o == nil || isNil(o.QualityScores) {
    return nil, false
	}
	return o.QualityScores, true
}

// HasQualityScores returns a boolean if a field has been set.
func (o *CheckImageQualityResponseDataImageQuality) HasQualityScores() bool {
	if o != nil && !isNil(o.QualityScores) {
		return true
	}

	return false
}

// SetQualityScores gets a reference to the given CheckImageQualityResponseDataImageQualityQualityScores and assigns it to the QualityScores field.
func (o *CheckImageQualityResponseDataImageQuality) SetQualityScores(v CheckImageQualityResponseDataImageQualityQualityScores) {
	o.QualityScores = &v
}

// GetExtractionQuality returns the ExtractionQuality field value if set, zero value otherwise.
func (o *CheckImageQualityResponseDataImageQuality) GetExtractionQuality() string {
	if o == nil || isNil(o.ExtractionQuality) {
		var ret string
		return ret
	}
	return *o.ExtractionQuality
}

// GetExtractionQualityOk returns a tuple with the ExtractionQuality field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckImageQualityResponseDataImageQuality) GetExtractionQualityOk() (*string, bool) {
	if o == nil || isNil(o.ExtractionQuality) {
    return nil, false
	}
	return o.ExtractionQuality, true
}

// HasExtractionQuality returns a boolean if a field has been set.
func (o *CheckImageQualityResponseDataImageQuality) HasExtractionQuality() bool {
	if o != nil && !isNil(o.ExtractionQuality) {
		return true
	}

	return false
}

// SetExtractionQuality gets a reference to the given string and assigns it to the ExtractionQuality field.
func (o *CheckImageQualityResponseDataImageQuality) SetExtractionQuality(v string) {
	o.ExtractionQuality = &v
}

// GetScore returns the Score field value if set, zero value otherwise.
func (o *CheckImageQualityResponseDataImageQuality) GetScore() float32 {
	if o == nil || isNil(o.Score) {
		var ret float32
		return ret
	}
	return *o.Score
}

// GetScoreOk returns a tuple with the Score field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckImageQualityResponseDataImageQuality) GetScoreOk() (*float32, bool) {
	if o == nil || isNil(o.Score) {
    return nil, false
	}
	return o.Score, true
}

// HasScore returns a boolean if a field has been set.
func (o *CheckImageQualityResponseDataImageQuality) HasScore() bool {
	if o != nil && !isNil(o.Score) {
		return true
	}

	return false
}

// SetScore gets a reference to the given float32 and assigns it to the Score field.
func (o *CheckImageQualityResponseDataImageQuality) SetScore(v float32) {
	o.Score = &v
}

// GetMsg returns the Msg field value if set, zero value otherwise.
func (o *CheckImageQualityResponseDataImageQuality) GetMsg() string {
	if o == nil || isNil(o.Msg) {
		var ret string
		return ret
	}
	return *o.Msg
}

// GetMsgOk returns a tuple with the Msg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckImageQualityResponseDataImageQuality) GetMsgOk() (*string, bool) {
	if o == nil || isNil(o.Msg) {
    return nil, false
	}
	return o.Msg, true
}

// HasMsg returns a boolean if a field has been set.
func (o *CheckImageQualityResponseDataImageQuality) HasMsg() bool {
	if o != nil && !isNil(o.Msg) {
		return true
	}

	return false
}

// SetMsg gets a reference to the given string and assigns it to the Msg field.
func (o *CheckImageQualityResponseDataImageQuality) SetMsg(v string) {
	o.Msg = &v
}

func (o CheckImageQualityResponseDataImageQuality) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Summary) {
		toSerialize["summary"] = o.Summary
	}
	if !isNil(o.QualityScores) {
		toSerialize["qualityScores"] = o.QualityScores
	}
	if !isNil(o.ExtractionQuality) {
		toSerialize["extractionQuality"] = o.ExtractionQuality
	}
	if !isNil(o.Score) {
		toSerialize["score"] = o.Score
	}
	if !isNil(o.Msg) {
		toSerialize["msg"] = o.Msg
	}
	return json.Marshal(toSerialize)
}

type NullableCheckImageQualityResponseDataImageQuality struct {
	value *CheckImageQualityResponseDataImageQuality
	isSet bool
}

func (v NullableCheckImageQualityResponseDataImageQuality) Get() *CheckImageQualityResponseDataImageQuality {
	return v.value
}

func (v *NullableCheckImageQualityResponseDataImageQuality) Set(val *CheckImageQualityResponseDataImageQuality) {
	v.value = val
	v.isSet = true
}

func (v NullableCheckImageQualityResponseDataImageQuality) IsSet() bool {
	return v.isSet
}

func (v *NullableCheckImageQualityResponseDataImageQuality) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCheckImageQualityResponseDataImageQuality(val *CheckImageQualityResponseDataImageQuality) *NullableCheckImageQualityResponseDataImageQuality {
	return &NullableCheckImageQualityResponseDataImageQuality{value: val, isSet: true}
}

func (v NullableCheckImageQualityResponseDataImageQuality) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCheckImageQualityResponseDataImageQuality) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


