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

// CheckPhotocopy400Response struct for CheckPhotocopy400Response
type CheckPhotocopy400Response struct {
	DecentroTxnId *string `json:"decentroTxnId,omitempty"`
	Status *string `json:"status,omitempty"`
	ResponseCode *string `json:"responseCode,omitempty"`
	Message *string `json:"message,omitempty"`
}

// NewCheckPhotocopy400Response instantiates a new CheckPhotocopy400Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCheckPhotocopy400Response() *CheckPhotocopy400Response {
	this := CheckPhotocopy400Response{}
	return &this
}

// NewCheckPhotocopy400ResponseWithDefaults instantiates a new CheckPhotocopy400Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCheckPhotocopy400ResponseWithDefaults() *CheckPhotocopy400Response {
	this := CheckPhotocopy400Response{}
	return &this
}

// GetDecentroTxnId returns the DecentroTxnId field value if set, zero value otherwise.
func (o *CheckPhotocopy400Response) GetDecentroTxnId() string {
	if o == nil || isNil(o.DecentroTxnId) {
		var ret string
		return ret
	}
	return *o.DecentroTxnId
}

// GetDecentroTxnIdOk returns a tuple with the DecentroTxnId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckPhotocopy400Response) GetDecentroTxnIdOk() (*string, bool) {
	if o == nil || isNil(o.DecentroTxnId) {
    return nil, false
	}
	return o.DecentroTxnId, true
}

// HasDecentroTxnId returns a boolean if a field has been set.
func (o *CheckPhotocopy400Response) HasDecentroTxnId() bool {
	if o != nil && !isNil(o.DecentroTxnId) {
		return true
	}

	return false
}

// SetDecentroTxnId gets a reference to the given string and assigns it to the DecentroTxnId field.
func (o *CheckPhotocopy400Response) SetDecentroTxnId(v string) {
	o.DecentroTxnId = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *CheckPhotocopy400Response) GetStatus() string {
	if o == nil || isNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckPhotocopy400Response) GetStatusOk() (*string, bool) {
	if o == nil || isNil(o.Status) {
    return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *CheckPhotocopy400Response) HasStatus() bool {
	if o != nil && !isNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *CheckPhotocopy400Response) SetStatus(v string) {
	o.Status = &v
}

// GetResponseCode returns the ResponseCode field value if set, zero value otherwise.
func (o *CheckPhotocopy400Response) GetResponseCode() string {
	if o == nil || isNil(o.ResponseCode) {
		var ret string
		return ret
	}
	return *o.ResponseCode
}

// GetResponseCodeOk returns a tuple with the ResponseCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckPhotocopy400Response) GetResponseCodeOk() (*string, bool) {
	if o == nil || isNil(o.ResponseCode) {
    return nil, false
	}
	return o.ResponseCode, true
}

// HasResponseCode returns a boolean if a field has been set.
func (o *CheckPhotocopy400Response) HasResponseCode() bool {
	if o != nil && !isNil(o.ResponseCode) {
		return true
	}

	return false
}

// SetResponseCode gets a reference to the given string and assigns it to the ResponseCode field.
func (o *CheckPhotocopy400Response) SetResponseCode(v string) {
	o.ResponseCode = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *CheckPhotocopy400Response) GetMessage() string {
	if o == nil || isNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckPhotocopy400Response) GetMessageOk() (*string, bool) {
	if o == nil || isNil(o.Message) {
    return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *CheckPhotocopy400Response) HasMessage() bool {
	if o != nil && !isNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *CheckPhotocopy400Response) SetMessage(v string) {
	o.Message = &v
}

func (o CheckPhotocopy400Response) MarshalJSON() ([]byte, error) {
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
	return json.Marshal(toSerialize)
}

type NullableCheckPhotocopy400Response struct {
	value *CheckPhotocopy400Response
	isSet bool
}

func (v NullableCheckPhotocopy400Response) Get() *CheckPhotocopy400Response {
	return v.value
}

func (v *NullableCheckPhotocopy400Response) Set(val *CheckPhotocopy400Response) {
	v.value = val
	v.isSet = true
}

func (v NullableCheckPhotocopy400Response) IsSet() bool {
	return v.isSet
}

func (v *NullableCheckPhotocopy400Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCheckPhotocopy400Response(val *CheckPhotocopy400Response) *NullableCheckPhotocopy400Response {
	return &NullableCheckPhotocopy400Response{value: val, isSet: true}
}

func (v NullableCheckPhotocopy400Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCheckPhotocopy400Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


