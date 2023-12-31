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

// CheckImageQualityResponseDataImageQualityQualityScores struct for CheckImageQualityResponseDataImageQualityQualityScores
type CheckImageQualityResponseDataImageQualityQualityScores struct {
	TextQuality *CheckImageQualityResponseDataImageQualityQualityScoresTextQuality `json:"textQuality,omitempty"`
	Sharpness *CheckImageQualityResponseDataImageQualityQualityScoresSharpness `json:"sharpness,omitempty"`
	Brightness *CheckImageQualityResponseDataImageQualityQualityScoresBrightness `json:"brightness,omitempty"`
	CompressionQuality *CheckImageQualityResponseDataImageQualityQualityScoresCompressionQuality `json:"compressionQuality,omitempty"`
}

// NewCheckImageQualityResponseDataImageQualityQualityScores instantiates a new CheckImageQualityResponseDataImageQualityQualityScores object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCheckImageQualityResponseDataImageQualityQualityScores() *CheckImageQualityResponseDataImageQualityQualityScores {
	this := CheckImageQualityResponseDataImageQualityQualityScores{}
	return &this
}

// NewCheckImageQualityResponseDataImageQualityQualityScoresWithDefaults instantiates a new CheckImageQualityResponseDataImageQualityQualityScores object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCheckImageQualityResponseDataImageQualityQualityScoresWithDefaults() *CheckImageQualityResponseDataImageQualityQualityScores {
	this := CheckImageQualityResponseDataImageQualityQualityScores{}
	return &this
}

// GetTextQuality returns the TextQuality field value if set, zero value otherwise.
func (o *CheckImageQualityResponseDataImageQualityQualityScores) GetTextQuality() CheckImageQualityResponseDataImageQualityQualityScoresTextQuality {
	if o == nil || isNil(o.TextQuality) {
		var ret CheckImageQualityResponseDataImageQualityQualityScoresTextQuality
		return ret
	}
	return *o.TextQuality
}

// GetTextQualityOk returns a tuple with the TextQuality field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckImageQualityResponseDataImageQualityQualityScores) GetTextQualityOk() (*CheckImageQualityResponseDataImageQualityQualityScoresTextQuality, bool) {
	if o == nil || isNil(o.TextQuality) {
    return nil, false
	}
	return o.TextQuality, true
}

// HasTextQuality returns a boolean if a field has been set.
func (o *CheckImageQualityResponseDataImageQualityQualityScores) HasTextQuality() bool {
	if o != nil && !isNil(o.TextQuality) {
		return true
	}

	return false
}

// SetTextQuality gets a reference to the given CheckImageQualityResponseDataImageQualityQualityScoresTextQuality and assigns it to the TextQuality field.
func (o *CheckImageQualityResponseDataImageQualityQualityScores) SetTextQuality(v CheckImageQualityResponseDataImageQualityQualityScoresTextQuality) {
	o.TextQuality = &v
}

// GetSharpness returns the Sharpness field value if set, zero value otherwise.
func (o *CheckImageQualityResponseDataImageQualityQualityScores) GetSharpness() CheckImageQualityResponseDataImageQualityQualityScoresSharpness {
	if o == nil || isNil(o.Sharpness) {
		var ret CheckImageQualityResponseDataImageQualityQualityScoresSharpness
		return ret
	}
	return *o.Sharpness
}

// GetSharpnessOk returns a tuple with the Sharpness field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckImageQualityResponseDataImageQualityQualityScores) GetSharpnessOk() (*CheckImageQualityResponseDataImageQualityQualityScoresSharpness, bool) {
	if o == nil || isNil(o.Sharpness) {
    return nil, false
	}
	return o.Sharpness, true
}

// HasSharpness returns a boolean if a field has been set.
func (o *CheckImageQualityResponseDataImageQualityQualityScores) HasSharpness() bool {
	if o != nil && !isNil(o.Sharpness) {
		return true
	}

	return false
}

// SetSharpness gets a reference to the given CheckImageQualityResponseDataImageQualityQualityScoresSharpness and assigns it to the Sharpness field.
func (o *CheckImageQualityResponseDataImageQualityQualityScores) SetSharpness(v CheckImageQualityResponseDataImageQualityQualityScoresSharpness) {
	o.Sharpness = &v
}

// GetBrightness returns the Brightness field value if set, zero value otherwise.
func (o *CheckImageQualityResponseDataImageQualityQualityScores) GetBrightness() CheckImageQualityResponseDataImageQualityQualityScoresBrightness {
	if o == nil || isNil(o.Brightness) {
		var ret CheckImageQualityResponseDataImageQualityQualityScoresBrightness
		return ret
	}
	return *o.Brightness
}

// GetBrightnessOk returns a tuple with the Brightness field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckImageQualityResponseDataImageQualityQualityScores) GetBrightnessOk() (*CheckImageQualityResponseDataImageQualityQualityScoresBrightness, bool) {
	if o == nil || isNil(o.Brightness) {
    return nil, false
	}
	return o.Brightness, true
}

// HasBrightness returns a boolean if a field has been set.
func (o *CheckImageQualityResponseDataImageQualityQualityScores) HasBrightness() bool {
	if o != nil && !isNil(o.Brightness) {
		return true
	}

	return false
}

// SetBrightness gets a reference to the given CheckImageQualityResponseDataImageQualityQualityScoresBrightness and assigns it to the Brightness field.
func (o *CheckImageQualityResponseDataImageQualityQualityScores) SetBrightness(v CheckImageQualityResponseDataImageQualityQualityScoresBrightness) {
	o.Brightness = &v
}

// GetCompressionQuality returns the CompressionQuality field value if set, zero value otherwise.
func (o *CheckImageQualityResponseDataImageQualityQualityScores) GetCompressionQuality() CheckImageQualityResponseDataImageQualityQualityScoresCompressionQuality {
	if o == nil || isNil(o.CompressionQuality) {
		var ret CheckImageQualityResponseDataImageQualityQualityScoresCompressionQuality
		return ret
	}
	return *o.CompressionQuality
}

// GetCompressionQualityOk returns a tuple with the CompressionQuality field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckImageQualityResponseDataImageQualityQualityScores) GetCompressionQualityOk() (*CheckImageQualityResponseDataImageQualityQualityScoresCompressionQuality, bool) {
	if o == nil || isNil(o.CompressionQuality) {
    return nil, false
	}
	return o.CompressionQuality, true
}

// HasCompressionQuality returns a boolean if a field has been set.
func (o *CheckImageQualityResponseDataImageQualityQualityScores) HasCompressionQuality() bool {
	if o != nil && !isNil(o.CompressionQuality) {
		return true
	}

	return false
}

// SetCompressionQuality gets a reference to the given CheckImageQualityResponseDataImageQualityQualityScoresCompressionQuality and assigns it to the CompressionQuality field.
func (o *CheckImageQualityResponseDataImageQualityQualityScores) SetCompressionQuality(v CheckImageQualityResponseDataImageQualityQualityScoresCompressionQuality) {
	o.CompressionQuality = &v
}

func (o CheckImageQualityResponseDataImageQualityQualityScores) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.TextQuality) {
		toSerialize["textQuality"] = o.TextQuality
	}
	if !isNil(o.Sharpness) {
		toSerialize["sharpness"] = o.Sharpness
	}
	if !isNil(o.Brightness) {
		toSerialize["brightness"] = o.Brightness
	}
	if !isNil(o.CompressionQuality) {
		toSerialize["compressionQuality"] = o.CompressionQuality
	}
	return json.Marshal(toSerialize)
}

type NullableCheckImageQualityResponseDataImageQualityQualityScores struct {
	value *CheckImageQualityResponseDataImageQualityQualityScores
	isSet bool
}

func (v NullableCheckImageQualityResponseDataImageQualityQualityScores) Get() *CheckImageQualityResponseDataImageQualityQualityScores {
	return v.value
}

func (v *NullableCheckImageQualityResponseDataImageQualityQualityScores) Set(val *CheckImageQualityResponseDataImageQualityQualityScores) {
	v.value = val
	v.isSet = true
}

func (v NullableCheckImageQualityResponseDataImageQualityQualityScores) IsSet() bool {
	return v.isSet
}

func (v *NullableCheckImageQualityResponseDataImageQualityQualityScores) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCheckImageQualityResponseDataImageQualityQualityScores(val *CheckImageQualityResponseDataImageQualityQualityScores) *NullableCheckImageQualityResponseDataImageQualityQualityScores {
	return &NullableCheckImageQualityResponseDataImageQualityQualityScores{value: val, isSet: true}
}

func (v NullableCheckImageQualityResponseDataImageQualityQualityScores) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCheckImageQualityResponseDataImageQualityQualityScores) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


