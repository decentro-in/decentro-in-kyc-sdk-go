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

// UpgradedaadhaarxmlGenerateOtpResponse struct for UpgradedaadhaarxmlGenerateOtpResponse
type UpgradedaadhaarxmlGenerateOtpResponse struct {
	DecentroTxnId *string `json:"decentroTxnId,omitempty"`
	Status *string `json:"status,omitempty"`
	ResponseCode *string `json:"responseCode,omitempty"`
	Message *string `json:"message,omitempty"`
	ResponseKey *string `json:"responseKey,omitempty"`
}

// NewUpgradedaadhaarxmlGenerateOtpResponse instantiates a new UpgradedaadhaarxmlGenerateOtpResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpgradedaadhaarxmlGenerateOtpResponse() *UpgradedaadhaarxmlGenerateOtpResponse {
	this := UpgradedaadhaarxmlGenerateOtpResponse{}
	return &this
}

// NewUpgradedaadhaarxmlGenerateOtpResponseWithDefaults instantiates a new UpgradedaadhaarxmlGenerateOtpResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpgradedaadhaarxmlGenerateOtpResponseWithDefaults() *UpgradedaadhaarxmlGenerateOtpResponse {
	this := UpgradedaadhaarxmlGenerateOtpResponse{}
	return &this
}

// GetDecentroTxnId returns the DecentroTxnId field value if set, zero value otherwise.
func (o *UpgradedaadhaarxmlGenerateOtpResponse) GetDecentroTxnId() string {
	if o == nil || isNil(o.DecentroTxnId) {
		var ret string
		return ret
	}
	return *o.DecentroTxnId
}

// GetDecentroTxnIdOk returns a tuple with the DecentroTxnId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpgradedaadhaarxmlGenerateOtpResponse) GetDecentroTxnIdOk() (*string, bool) {
	if o == nil || isNil(o.DecentroTxnId) {
    return nil, false
	}
	return o.DecentroTxnId, true
}

// HasDecentroTxnId returns a boolean if a field has been set.
func (o *UpgradedaadhaarxmlGenerateOtpResponse) HasDecentroTxnId() bool {
	if o != nil && !isNil(o.DecentroTxnId) {
		return true
	}

	return false
}

// SetDecentroTxnId gets a reference to the given string and assigns it to the DecentroTxnId field.
func (o *UpgradedaadhaarxmlGenerateOtpResponse) SetDecentroTxnId(v string) {
	o.DecentroTxnId = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *UpgradedaadhaarxmlGenerateOtpResponse) GetStatus() string {
	if o == nil || isNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpgradedaadhaarxmlGenerateOtpResponse) GetStatusOk() (*string, bool) {
	if o == nil || isNil(o.Status) {
    return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *UpgradedaadhaarxmlGenerateOtpResponse) HasStatus() bool {
	if o != nil && !isNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *UpgradedaadhaarxmlGenerateOtpResponse) SetStatus(v string) {
	o.Status = &v
}

// GetResponseCode returns the ResponseCode field value if set, zero value otherwise.
func (o *UpgradedaadhaarxmlGenerateOtpResponse) GetResponseCode() string {
	if o == nil || isNil(o.ResponseCode) {
		var ret string
		return ret
	}
	return *o.ResponseCode
}

// GetResponseCodeOk returns a tuple with the ResponseCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpgradedaadhaarxmlGenerateOtpResponse) GetResponseCodeOk() (*string, bool) {
	if o == nil || isNil(o.ResponseCode) {
    return nil, false
	}
	return o.ResponseCode, true
}

// HasResponseCode returns a boolean if a field has been set.
func (o *UpgradedaadhaarxmlGenerateOtpResponse) HasResponseCode() bool {
	if o != nil && !isNil(o.ResponseCode) {
		return true
	}

	return false
}

// SetResponseCode gets a reference to the given string and assigns it to the ResponseCode field.
func (o *UpgradedaadhaarxmlGenerateOtpResponse) SetResponseCode(v string) {
	o.ResponseCode = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *UpgradedaadhaarxmlGenerateOtpResponse) GetMessage() string {
	if o == nil || isNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpgradedaadhaarxmlGenerateOtpResponse) GetMessageOk() (*string, bool) {
	if o == nil || isNil(o.Message) {
    return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *UpgradedaadhaarxmlGenerateOtpResponse) HasMessage() bool {
	if o != nil && !isNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *UpgradedaadhaarxmlGenerateOtpResponse) SetMessage(v string) {
	o.Message = &v
}

// GetResponseKey returns the ResponseKey field value if set, zero value otherwise.
func (o *UpgradedaadhaarxmlGenerateOtpResponse) GetResponseKey() string {
	if o == nil || isNil(o.ResponseKey) {
		var ret string
		return ret
	}
	return *o.ResponseKey
}

// GetResponseKeyOk returns a tuple with the ResponseKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpgradedaadhaarxmlGenerateOtpResponse) GetResponseKeyOk() (*string, bool) {
	if o == nil || isNil(o.ResponseKey) {
    return nil, false
	}
	return o.ResponseKey, true
}

// HasResponseKey returns a boolean if a field has been set.
func (o *UpgradedaadhaarxmlGenerateOtpResponse) HasResponseKey() bool {
	if o != nil && !isNil(o.ResponseKey) {
		return true
	}

	return false
}

// SetResponseKey gets a reference to the given string and assigns it to the ResponseKey field.
func (o *UpgradedaadhaarxmlGenerateOtpResponse) SetResponseKey(v string) {
	o.ResponseKey = &v
}

func (o UpgradedaadhaarxmlGenerateOtpResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.DecentroTxnId) {
		toSerialize["decentroTxnId"] = o.DecentroTxnId
	}
	if !isNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !isNil(o.ResponseCode) {
		toSerialize["responseCode"] = o.ResponseCode
	}
	if !isNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	if !isNil(o.ResponseKey) {
		toSerialize["responseKey"] = o.ResponseKey
	}
	return json.Marshal(toSerialize)
}

type NullableUpgradedaadhaarxmlGenerateOtpResponse struct {
	value *UpgradedaadhaarxmlGenerateOtpResponse
	isSet bool
}

func (v NullableUpgradedaadhaarxmlGenerateOtpResponse) Get() *UpgradedaadhaarxmlGenerateOtpResponse {
	return v.value
}

func (v *NullableUpgradedaadhaarxmlGenerateOtpResponse) Set(val *UpgradedaadhaarxmlGenerateOtpResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableUpgradedaadhaarxmlGenerateOtpResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableUpgradedaadhaarxmlGenerateOtpResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpgradedaadhaarxmlGenerateOtpResponse(val *UpgradedaadhaarxmlGenerateOtpResponse) *NullableUpgradedaadhaarxmlGenerateOtpResponse {
	return &NullableUpgradedaadhaarxmlGenerateOtpResponse{value: val, isSet: true}
}

func (v NullableUpgradedaadhaarxmlGenerateOtpResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpgradedaadhaarxmlGenerateOtpResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


