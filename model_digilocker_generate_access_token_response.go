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

// DigilockerGenerateAccessTokenResponse struct for DigilockerGenerateAccessTokenResponse
type DigilockerGenerateAccessTokenResponse struct {
	DecentroTxnId *string `json:"decentroTxnId,omitempty"`
	Status *string `json:"status,omitempty"`
	ResponseCode *string `json:"responseCode,omitempty"`
	Message *string `json:"message,omitempty"`
	Data *DigilockerGenerateAccessTokenResponseData `json:"data,omitempty"`
	ResponseKey *string `json:"responseKey,omitempty"`
}

// NewDigilockerGenerateAccessTokenResponse instantiates a new DigilockerGenerateAccessTokenResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDigilockerGenerateAccessTokenResponse() *DigilockerGenerateAccessTokenResponse {
	this := DigilockerGenerateAccessTokenResponse{}
	return &this
}

// NewDigilockerGenerateAccessTokenResponseWithDefaults instantiates a new DigilockerGenerateAccessTokenResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDigilockerGenerateAccessTokenResponseWithDefaults() *DigilockerGenerateAccessTokenResponse {
	this := DigilockerGenerateAccessTokenResponse{}
	return &this
}

// GetDecentroTxnId returns the DecentroTxnId field value if set, zero value otherwise.
func (o *DigilockerGenerateAccessTokenResponse) GetDecentroTxnId() string {
	if o == nil || isNil(o.DecentroTxnId) {
		var ret string
		return ret
	}
	return *o.DecentroTxnId
}

// GetDecentroTxnIdOk returns a tuple with the DecentroTxnId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DigilockerGenerateAccessTokenResponse) GetDecentroTxnIdOk() (*string, bool) {
	if o == nil || isNil(o.DecentroTxnId) {
    return nil, false
	}
	return o.DecentroTxnId, true
}

// HasDecentroTxnId returns a boolean if a field has been set.
func (o *DigilockerGenerateAccessTokenResponse) HasDecentroTxnId() bool {
	if o != nil && !isNil(o.DecentroTxnId) {
		return true
	}

	return false
}

// SetDecentroTxnId gets a reference to the given string and assigns it to the DecentroTxnId field.
func (o *DigilockerGenerateAccessTokenResponse) SetDecentroTxnId(v string) {
	o.DecentroTxnId = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *DigilockerGenerateAccessTokenResponse) GetStatus() string {
	if o == nil || isNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DigilockerGenerateAccessTokenResponse) GetStatusOk() (*string, bool) {
	if o == nil || isNil(o.Status) {
    return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *DigilockerGenerateAccessTokenResponse) HasStatus() bool {
	if o != nil && !isNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *DigilockerGenerateAccessTokenResponse) SetStatus(v string) {
	o.Status = &v
}

// GetResponseCode returns the ResponseCode field value if set, zero value otherwise.
func (o *DigilockerGenerateAccessTokenResponse) GetResponseCode() string {
	if o == nil || isNil(o.ResponseCode) {
		var ret string
		return ret
	}
	return *o.ResponseCode
}

// GetResponseCodeOk returns a tuple with the ResponseCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DigilockerGenerateAccessTokenResponse) GetResponseCodeOk() (*string, bool) {
	if o == nil || isNil(o.ResponseCode) {
    return nil, false
	}
	return o.ResponseCode, true
}

// HasResponseCode returns a boolean if a field has been set.
func (o *DigilockerGenerateAccessTokenResponse) HasResponseCode() bool {
	if o != nil && !isNil(o.ResponseCode) {
		return true
	}

	return false
}

// SetResponseCode gets a reference to the given string and assigns it to the ResponseCode field.
func (o *DigilockerGenerateAccessTokenResponse) SetResponseCode(v string) {
	o.ResponseCode = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *DigilockerGenerateAccessTokenResponse) GetMessage() string {
	if o == nil || isNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DigilockerGenerateAccessTokenResponse) GetMessageOk() (*string, bool) {
	if o == nil || isNil(o.Message) {
    return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *DigilockerGenerateAccessTokenResponse) HasMessage() bool {
	if o != nil && !isNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *DigilockerGenerateAccessTokenResponse) SetMessage(v string) {
	o.Message = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *DigilockerGenerateAccessTokenResponse) GetData() DigilockerGenerateAccessTokenResponseData {
	if o == nil || isNil(o.Data) {
		var ret DigilockerGenerateAccessTokenResponseData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DigilockerGenerateAccessTokenResponse) GetDataOk() (*DigilockerGenerateAccessTokenResponseData, bool) {
	if o == nil || isNil(o.Data) {
    return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *DigilockerGenerateAccessTokenResponse) HasData() bool {
	if o != nil && !isNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given DigilockerGenerateAccessTokenResponseData and assigns it to the Data field.
func (o *DigilockerGenerateAccessTokenResponse) SetData(v DigilockerGenerateAccessTokenResponseData) {
	o.Data = &v
}

// GetResponseKey returns the ResponseKey field value if set, zero value otherwise.
func (o *DigilockerGenerateAccessTokenResponse) GetResponseKey() string {
	if o == nil || isNil(o.ResponseKey) {
		var ret string
		return ret
	}
	return *o.ResponseKey
}

// GetResponseKeyOk returns a tuple with the ResponseKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DigilockerGenerateAccessTokenResponse) GetResponseKeyOk() (*string, bool) {
	if o == nil || isNil(o.ResponseKey) {
    return nil, false
	}
	return o.ResponseKey, true
}

// HasResponseKey returns a boolean if a field has been set.
func (o *DigilockerGenerateAccessTokenResponse) HasResponseKey() bool {
	if o != nil && !isNil(o.ResponseKey) {
		return true
	}

	return false
}

// SetResponseKey gets a reference to the given string and assigns it to the ResponseKey field.
func (o *DigilockerGenerateAccessTokenResponse) SetResponseKey(v string) {
	o.ResponseKey = &v
}

func (o DigilockerGenerateAccessTokenResponse) MarshalJSON() ([]byte, error) {
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

type NullableDigilockerGenerateAccessTokenResponse struct {
	value *DigilockerGenerateAccessTokenResponse
	isSet bool
}

func (v NullableDigilockerGenerateAccessTokenResponse) Get() *DigilockerGenerateAccessTokenResponse {
	return v.value
}

func (v *NullableDigilockerGenerateAccessTokenResponse) Set(val *DigilockerGenerateAccessTokenResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableDigilockerGenerateAccessTokenResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableDigilockerGenerateAccessTokenResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDigilockerGenerateAccessTokenResponse(val *DigilockerGenerateAccessTokenResponse) *NullableDigilockerGenerateAccessTokenResponse {
	return &NullableDigilockerGenerateAccessTokenResponse{value: val, isSet: true}
}

func (v NullableDigilockerGenerateAccessTokenResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDigilockerGenerateAccessTokenResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


