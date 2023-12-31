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

// DigilockerRevokeTokenResponse struct for DigilockerRevokeTokenResponse
type DigilockerRevokeTokenResponse struct {
	DecentroTxnId *string `json:"decentroTxnId,omitempty"`
	Status *string `json:"status,omitempty"`
	ResponseCode *string `json:"responseCode,omitempty"`
	Message *string `json:"message,omitempty"`
	Data *DigilockerRevokeTokenResponseData `json:"data,omitempty"`
	ResponseKey *string `json:"responseKey,omitempty"`
}

// NewDigilockerRevokeTokenResponse instantiates a new DigilockerRevokeTokenResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDigilockerRevokeTokenResponse() *DigilockerRevokeTokenResponse {
	this := DigilockerRevokeTokenResponse{}
	return &this
}

// NewDigilockerRevokeTokenResponseWithDefaults instantiates a new DigilockerRevokeTokenResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDigilockerRevokeTokenResponseWithDefaults() *DigilockerRevokeTokenResponse {
	this := DigilockerRevokeTokenResponse{}
	return &this
}

// GetDecentroTxnId returns the DecentroTxnId field value if set, zero value otherwise.
func (o *DigilockerRevokeTokenResponse) GetDecentroTxnId() string {
	if o == nil || isNil(o.DecentroTxnId) {
		var ret string
		return ret
	}
	return *o.DecentroTxnId
}

// GetDecentroTxnIdOk returns a tuple with the DecentroTxnId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DigilockerRevokeTokenResponse) GetDecentroTxnIdOk() (*string, bool) {
	if o == nil || isNil(o.DecentroTxnId) {
    return nil, false
	}
	return o.DecentroTxnId, true
}

// HasDecentroTxnId returns a boolean if a field has been set.
func (o *DigilockerRevokeTokenResponse) HasDecentroTxnId() bool {
	if o != nil && !isNil(o.DecentroTxnId) {
		return true
	}

	return false
}

// SetDecentroTxnId gets a reference to the given string and assigns it to the DecentroTxnId field.
func (o *DigilockerRevokeTokenResponse) SetDecentroTxnId(v string) {
	o.DecentroTxnId = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *DigilockerRevokeTokenResponse) GetStatus() string {
	if o == nil || isNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DigilockerRevokeTokenResponse) GetStatusOk() (*string, bool) {
	if o == nil || isNil(o.Status) {
    return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *DigilockerRevokeTokenResponse) HasStatus() bool {
	if o != nil && !isNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *DigilockerRevokeTokenResponse) SetStatus(v string) {
	o.Status = &v
}

// GetResponseCode returns the ResponseCode field value if set, zero value otherwise.
func (o *DigilockerRevokeTokenResponse) GetResponseCode() string {
	if o == nil || isNil(o.ResponseCode) {
		var ret string
		return ret
	}
	return *o.ResponseCode
}

// GetResponseCodeOk returns a tuple with the ResponseCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DigilockerRevokeTokenResponse) GetResponseCodeOk() (*string, bool) {
	if o == nil || isNil(o.ResponseCode) {
    return nil, false
	}
	return o.ResponseCode, true
}

// HasResponseCode returns a boolean if a field has been set.
func (o *DigilockerRevokeTokenResponse) HasResponseCode() bool {
	if o != nil && !isNil(o.ResponseCode) {
		return true
	}

	return false
}

// SetResponseCode gets a reference to the given string and assigns it to the ResponseCode field.
func (o *DigilockerRevokeTokenResponse) SetResponseCode(v string) {
	o.ResponseCode = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *DigilockerRevokeTokenResponse) GetMessage() string {
	if o == nil || isNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DigilockerRevokeTokenResponse) GetMessageOk() (*string, bool) {
	if o == nil || isNil(o.Message) {
    return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *DigilockerRevokeTokenResponse) HasMessage() bool {
	if o != nil && !isNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *DigilockerRevokeTokenResponse) SetMessage(v string) {
	o.Message = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *DigilockerRevokeTokenResponse) GetData() DigilockerRevokeTokenResponseData {
	if o == nil || isNil(o.Data) {
		var ret DigilockerRevokeTokenResponseData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DigilockerRevokeTokenResponse) GetDataOk() (*DigilockerRevokeTokenResponseData, bool) {
	if o == nil || isNil(o.Data) {
    return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *DigilockerRevokeTokenResponse) HasData() bool {
	if o != nil && !isNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given DigilockerRevokeTokenResponseData and assigns it to the Data field.
func (o *DigilockerRevokeTokenResponse) SetData(v DigilockerRevokeTokenResponseData) {
	o.Data = &v
}

// GetResponseKey returns the ResponseKey field value if set, zero value otherwise.
func (o *DigilockerRevokeTokenResponse) GetResponseKey() string {
	if o == nil || isNil(o.ResponseKey) {
		var ret string
		return ret
	}
	return *o.ResponseKey
}

// GetResponseKeyOk returns a tuple with the ResponseKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DigilockerRevokeTokenResponse) GetResponseKeyOk() (*string, bool) {
	if o == nil || isNil(o.ResponseKey) {
    return nil, false
	}
	return o.ResponseKey, true
}

// HasResponseKey returns a boolean if a field has been set.
func (o *DigilockerRevokeTokenResponse) HasResponseKey() bool {
	if o != nil && !isNil(o.ResponseKey) {
		return true
	}

	return false
}

// SetResponseKey gets a reference to the given string and assigns it to the ResponseKey field.
func (o *DigilockerRevokeTokenResponse) SetResponseKey(v string) {
	o.ResponseKey = &v
}

func (o DigilockerRevokeTokenResponse) MarshalJSON() ([]byte, error) {
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
	if !isNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	if !isNil(o.ResponseKey) {
		toSerialize["responseKey"] = o.ResponseKey
	}
	return json.Marshal(toSerialize)
}

type NullableDigilockerRevokeTokenResponse struct {
	value *DigilockerRevokeTokenResponse
	isSet bool
}

func (v NullableDigilockerRevokeTokenResponse) Get() *DigilockerRevokeTokenResponse {
	return v.value
}

func (v *NullableDigilockerRevokeTokenResponse) Set(val *DigilockerRevokeTokenResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableDigilockerRevokeTokenResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableDigilockerRevokeTokenResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDigilockerRevokeTokenResponse(val *DigilockerRevokeTokenResponse) *NullableDigilockerRevokeTokenResponse {
	return &NullableDigilockerRevokeTokenResponse{value: val, isSet: true}
}

func (v NullableDigilockerRevokeTokenResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDigilockerRevokeTokenResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


