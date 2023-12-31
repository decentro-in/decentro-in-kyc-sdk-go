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

// CheckVideoLivenessResponse struct for CheckVideoLivenessResponse
type CheckVideoLivenessResponse struct {
	Data *CheckVideoLivenessResponseData `json:"data,omitempty"`
	DecentroTxnId *string `json:"decentroTxnId,omitempty"`
	Message *string `json:"message,omitempty"`
	ResponseCode *string `json:"responseCode,omitempty"`
	Status *string `json:"status,omitempty"`
}

// NewCheckVideoLivenessResponse instantiates a new CheckVideoLivenessResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCheckVideoLivenessResponse() *CheckVideoLivenessResponse {
	this := CheckVideoLivenessResponse{}
	return &this
}

// NewCheckVideoLivenessResponseWithDefaults instantiates a new CheckVideoLivenessResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCheckVideoLivenessResponseWithDefaults() *CheckVideoLivenessResponse {
	this := CheckVideoLivenessResponse{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *CheckVideoLivenessResponse) GetData() CheckVideoLivenessResponseData {
	if o == nil || isNil(o.Data) {
		var ret CheckVideoLivenessResponseData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckVideoLivenessResponse) GetDataOk() (*CheckVideoLivenessResponseData, bool) {
	if o == nil || isNil(o.Data) {
    return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *CheckVideoLivenessResponse) HasData() bool {
	if o != nil && !isNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given CheckVideoLivenessResponseData and assigns it to the Data field.
func (o *CheckVideoLivenessResponse) SetData(v CheckVideoLivenessResponseData) {
	o.Data = &v
}

// GetDecentroTxnId returns the DecentroTxnId field value if set, zero value otherwise.
func (o *CheckVideoLivenessResponse) GetDecentroTxnId() string {
	if o == nil || isNil(o.DecentroTxnId) {
		var ret string
		return ret
	}
	return *o.DecentroTxnId
}

// GetDecentroTxnIdOk returns a tuple with the DecentroTxnId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckVideoLivenessResponse) GetDecentroTxnIdOk() (*string, bool) {
	if o == nil || isNil(o.DecentroTxnId) {
    return nil, false
	}
	return o.DecentroTxnId, true
}

// HasDecentroTxnId returns a boolean if a field has been set.
func (o *CheckVideoLivenessResponse) HasDecentroTxnId() bool {
	if o != nil && !isNil(o.DecentroTxnId) {
		return true
	}

	return false
}

// SetDecentroTxnId gets a reference to the given string and assigns it to the DecentroTxnId field.
func (o *CheckVideoLivenessResponse) SetDecentroTxnId(v string) {
	o.DecentroTxnId = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *CheckVideoLivenessResponse) GetMessage() string {
	if o == nil || isNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckVideoLivenessResponse) GetMessageOk() (*string, bool) {
	if o == nil || isNil(o.Message) {
    return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *CheckVideoLivenessResponse) HasMessage() bool {
	if o != nil && !isNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *CheckVideoLivenessResponse) SetMessage(v string) {
	o.Message = &v
}

// GetResponseCode returns the ResponseCode field value if set, zero value otherwise.
func (o *CheckVideoLivenessResponse) GetResponseCode() string {
	if o == nil || isNil(o.ResponseCode) {
		var ret string
		return ret
	}
	return *o.ResponseCode
}

// GetResponseCodeOk returns a tuple with the ResponseCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckVideoLivenessResponse) GetResponseCodeOk() (*string, bool) {
	if o == nil || isNil(o.ResponseCode) {
    return nil, false
	}
	return o.ResponseCode, true
}

// HasResponseCode returns a boolean if a field has been set.
func (o *CheckVideoLivenessResponse) HasResponseCode() bool {
	if o != nil && !isNil(o.ResponseCode) {
		return true
	}

	return false
}

// SetResponseCode gets a reference to the given string and assigns it to the ResponseCode field.
func (o *CheckVideoLivenessResponse) SetResponseCode(v string) {
	o.ResponseCode = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *CheckVideoLivenessResponse) GetStatus() string {
	if o == nil || isNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckVideoLivenessResponse) GetStatusOk() (*string, bool) {
	if o == nil || isNil(o.Status) {
    return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *CheckVideoLivenessResponse) HasStatus() bool {
	if o != nil && !isNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *CheckVideoLivenessResponse) SetStatus(v string) {
	o.Status = &v
}

func (o CheckVideoLivenessResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	if !isNil(o.DecentroTxnId) {
		toSerialize["decentroTxnId"] = o.DecentroTxnId
	}
	if !isNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	if !isNil(o.ResponseCode) {
		toSerialize["responseCode"] = o.ResponseCode
	}
	if !isNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	return json.Marshal(toSerialize)
}

type NullableCheckVideoLivenessResponse struct {
	value *CheckVideoLivenessResponse
	isSet bool
}

func (v NullableCheckVideoLivenessResponse) Get() *CheckVideoLivenessResponse {
	return v.value
}

func (v *NullableCheckVideoLivenessResponse) Set(val *CheckVideoLivenessResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCheckVideoLivenessResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCheckVideoLivenessResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCheckVideoLivenessResponse(val *CheckVideoLivenessResponse) *NullableCheckVideoLivenessResponse {
	return &NullableCheckVideoLivenessResponse{value: val, isSet: true}
}

func (v NullableCheckVideoLivenessResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCheckVideoLivenessResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


