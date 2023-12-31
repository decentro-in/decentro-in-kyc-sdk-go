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

// DigilockerPullFile400Response struct for DigilockerPullFile400Response
type DigilockerPullFile400Response struct {
	DecentroTxnId *string `json:"decentroTxnId,omitempty"`
	Message *string `json:"message,omitempty"`
	ResponseCode *string `json:"responseCode,omitempty"`
	ResponseKey *string `json:"responseKey,omitempty"`
	Status *string `json:"status,omitempty"`
}

// NewDigilockerPullFile400Response instantiates a new DigilockerPullFile400Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDigilockerPullFile400Response() *DigilockerPullFile400Response {
	this := DigilockerPullFile400Response{}
	return &this
}

// NewDigilockerPullFile400ResponseWithDefaults instantiates a new DigilockerPullFile400Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDigilockerPullFile400ResponseWithDefaults() *DigilockerPullFile400Response {
	this := DigilockerPullFile400Response{}
	return &this
}

// GetDecentroTxnId returns the DecentroTxnId field value if set, zero value otherwise.
func (o *DigilockerPullFile400Response) GetDecentroTxnId() string {
	if o == nil || isNil(o.DecentroTxnId) {
		var ret string
		return ret
	}
	return *o.DecentroTxnId
}

// GetDecentroTxnIdOk returns a tuple with the DecentroTxnId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DigilockerPullFile400Response) GetDecentroTxnIdOk() (*string, bool) {
	if o == nil || isNil(o.DecentroTxnId) {
    return nil, false
	}
	return o.DecentroTxnId, true
}

// HasDecentroTxnId returns a boolean if a field has been set.
func (o *DigilockerPullFile400Response) HasDecentroTxnId() bool {
	if o != nil && !isNil(o.DecentroTxnId) {
		return true
	}

	return false
}

// SetDecentroTxnId gets a reference to the given string and assigns it to the DecentroTxnId field.
func (o *DigilockerPullFile400Response) SetDecentroTxnId(v string) {
	o.DecentroTxnId = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *DigilockerPullFile400Response) GetMessage() string {
	if o == nil || isNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DigilockerPullFile400Response) GetMessageOk() (*string, bool) {
	if o == nil || isNil(o.Message) {
    return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *DigilockerPullFile400Response) HasMessage() bool {
	if o != nil && !isNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *DigilockerPullFile400Response) SetMessage(v string) {
	o.Message = &v
}

// GetResponseCode returns the ResponseCode field value if set, zero value otherwise.
func (o *DigilockerPullFile400Response) GetResponseCode() string {
	if o == nil || isNil(o.ResponseCode) {
		var ret string
		return ret
	}
	return *o.ResponseCode
}

// GetResponseCodeOk returns a tuple with the ResponseCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DigilockerPullFile400Response) GetResponseCodeOk() (*string, bool) {
	if o == nil || isNil(o.ResponseCode) {
    return nil, false
	}
	return o.ResponseCode, true
}

// HasResponseCode returns a boolean if a field has been set.
func (o *DigilockerPullFile400Response) HasResponseCode() bool {
	if o != nil && !isNil(o.ResponseCode) {
		return true
	}

	return false
}

// SetResponseCode gets a reference to the given string and assigns it to the ResponseCode field.
func (o *DigilockerPullFile400Response) SetResponseCode(v string) {
	o.ResponseCode = &v
}

// GetResponseKey returns the ResponseKey field value if set, zero value otherwise.
func (o *DigilockerPullFile400Response) GetResponseKey() string {
	if o == nil || isNil(o.ResponseKey) {
		var ret string
		return ret
	}
	return *o.ResponseKey
}

// GetResponseKeyOk returns a tuple with the ResponseKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DigilockerPullFile400Response) GetResponseKeyOk() (*string, bool) {
	if o == nil || isNil(o.ResponseKey) {
    return nil, false
	}
	return o.ResponseKey, true
}

// HasResponseKey returns a boolean if a field has been set.
func (o *DigilockerPullFile400Response) HasResponseKey() bool {
	if o != nil && !isNil(o.ResponseKey) {
		return true
	}

	return false
}

// SetResponseKey gets a reference to the given string and assigns it to the ResponseKey field.
func (o *DigilockerPullFile400Response) SetResponseKey(v string) {
	o.ResponseKey = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *DigilockerPullFile400Response) GetStatus() string {
	if o == nil || isNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DigilockerPullFile400Response) GetStatusOk() (*string, bool) {
	if o == nil || isNil(o.Status) {
    return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *DigilockerPullFile400Response) HasStatus() bool {
	if o != nil && !isNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *DigilockerPullFile400Response) SetStatus(v string) {
	o.Status = &v
}

func (o DigilockerPullFile400Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.DecentroTxnId) {
		toSerialize["decentroTxnId"] = o.DecentroTxnId
	}
	if !isNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	if !isNil(o.ResponseCode) {
		toSerialize["responseCode"] = o.ResponseCode
	}
	if !isNil(o.ResponseKey) {
		toSerialize["responseKey"] = o.ResponseKey
	}
	if !isNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	return json.Marshal(toSerialize)
}

type NullableDigilockerPullFile400Response struct {
	value *DigilockerPullFile400Response
	isSet bool
}

func (v NullableDigilockerPullFile400Response) Get() *DigilockerPullFile400Response {
	return v.value
}

func (v *NullableDigilockerPullFile400Response) Set(val *DigilockerPullFile400Response) {
	v.value = val
	v.isSet = true
}

func (v NullableDigilockerPullFile400Response) IsSet() bool {
	return v.isSet
}

func (v *NullableDigilockerPullFile400Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDigilockerPullFile400Response(val *DigilockerPullFile400Response) *NullableDigilockerPullFile400Response {
	return &NullableDigilockerPullFile400Response{value: val, isSet: true}
}

func (v NullableDigilockerPullFile400Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDigilockerPullFile400Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


