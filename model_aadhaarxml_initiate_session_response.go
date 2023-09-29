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

// AadhaarxmlInitiateSessionResponse struct for AadhaarxmlInitiateSessionResponse
type AadhaarxmlInitiateSessionResponse struct {
	DecentroTxnId *string `json:"decentroTxnId,omitempty"`
	Status *string `json:"status,omitempty"`
	ResponseCode *string `json:"responseCode,omitempty"`
	Message *string `json:"message,omitempty"`
	Data *AadhaarxmlInitiateSessionResponseData `json:"data,omitempty"`
	ResponseKey *string `json:"responseKey,omitempty"`
}

// NewAadhaarxmlInitiateSessionResponse instantiates a new AadhaarxmlInitiateSessionResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAadhaarxmlInitiateSessionResponse() *AadhaarxmlInitiateSessionResponse {
	this := AadhaarxmlInitiateSessionResponse{}
	return &this
}

// NewAadhaarxmlInitiateSessionResponseWithDefaults instantiates a new AadhaarxmlInitiateSessionResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAadhaarxmlInitiateSessionResponseWithDefaults() *AadhaarxmlInitiateSessionResponse {
	this := AadhaarxmlInitiateSessionResponse{}
	return &this
}

// GetDecentroTxnId returns the DecentroTxnId field value if set, zero value otherwise.
func (o *AadhaarxmlInitiateSessionResponse) GetDecentroTxnId() string {
	if o == nil || isNil(o.DecentroTxnId) {
		var ret string
		return ret
	}
	return *o.DecentroTxnId
}

// GetDecentroTxnIdOk returns a tuple with the DecentroTxnId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AadhaarxmlInitiateSessionResponse) GetDecentroTxnIdOk() (*string, bool) {
	if o == nil || isNil(o.DecentroTxnId) {
    return nil, false
	}
	return o.DecentroTxnId, true
}

// HasDecentroTxnId returns a boolean if a field has been set.
func (o *AadhaarxmlInitiateSessionResponse) HasDecentroTxnId() bool {
	if o != nil && !isNil(o.DecentroTxnId) {
		return true
	}

	return false
}

// SetDecentroTxnId gets a reference to the given string and assigns it to the DecentroTxnId field.
func (o *AadhaarxmlInitiateSessionResponse) SetDecentroTxnId(v string) {
	o.DecentroTxnId = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *AadhaarxmlInitiateSessionResponse) GetStatus() string {
	if o == nil || isNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AadhaarxmlInitiateSessionResponse) GetStatusOk() (*string, bool) {
	if o == nil || isNil(o.Status) {
    return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *AadhaarxmlInitiateSessionResponse) HasStatus() bool {
	if o != nil && !isNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *AadhaarxmlInitiateSessionResponse) SetStatus(v string) {
	o.Status = &v
}

// GetResponseCode returns the ResponseCode field value if set, zero value otherwise.
func (o *AadhaarxmlInitiateSessionResponse) GetResponseCode() string {
	if o == nil || isNil(o.ResponseCode) {
		var ret string
		return ret
	}
	return *o.ResponseCode
}

// GetResponseCodeOk returns a tuple with the ResponseCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AadhaarxmlInitiateSessionResponse) GetResponseCodeOk() (*string, bool) {
	if o == nil || isNil(o.ResponseCode) {
    return nil, false
	}
	return o.ResponseCode, true
}

// HasResponseCode returns a boolean if a field has been set.
func (o *AadhaarxmlInitiateSessionResponse) HasResponseCode() bool {
	if o != nil && !isNil(o.ResponseCode) {
		return true
	}

	return false
}

// SetResponseCode gets a reference to the given string and assigns it to the ResponseCode field.
func (o *AadhaarxmlInitiateSessionResponse) SetResponseCode(v string) {
	o.ResponseCode = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *AadhaarxmlInitiateSessionResponse) GetMessage() string {
	if o == nil || isNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AadhaarxmlInitiateSessionResponse) GetMessageOk() (*string, bool) {
	if o == nil || isNil(o.Message) {
    return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *AadhaarxmlInitiateSessionResponse) HasMessage() bool {
	if o != nil && !isNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *AadhaarxmlInitiateSessionResponse) SetMessage(v string) {
	o.Message = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *AadhaarxmlInitiateSessionResponse) GetData() AadhaarxmlInitiateSessionResponseData {
	if o == nil || isNil(o.Data) {
		var ret AadhaarxmlInitiateSessionResponseData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AadhaarxmlInitiateSessionResponse) GetDataOk() (*AadhaarxmlInitiateSessionResponseData, bool) {
	if o == nil || isNil(o.Data) {
    return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *AadhaarxmlInitiateSessionResponse) HasData() bool {
	if o != nil && !isNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given AadhaarxmlInitiateSessionResponseData and assigns it to the Data field.
func (o *AadhaarxmlInitiateSessionResponse) SetData(v AadhaarxmlInitiateSessionResponseData) {
	o.Data = &v
}

// GetResponseKey returns the ResponseKey field value if set, zero value otherwise.
func (o *AadhaarxmlInitiateSessionResponse) GetResponseKey() string {
	if o == nil || isNil(o.ResponseKey) {
		var ret string
		return ret
	}
	return *o.ResponseKey
}

// GetResponseKeyOk returns a tuple with the ResponseKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AadhaarxmlInitiateSessionResponse) GetResponseKeyOk() (*string, bool) {
	if o == nil || isNil(o.ResponseKey) {
    return nil, false
	}
	return o.ResponseKey, true
}

// HasResponseKey returns a boolean if a field has been set.
func (o *AadhaarxmlInitiateSessionResponse) HasResponseKey() bool {
	if o != nil && !isNil(o.ResponseKey) {
		return true
	}

	return false
}

// SetResponseKey gets a reference to the given string and assigns it to the ResponseKey field.
func (o *AadhaarxmlInitiateSessionResponse) SetResponseKey(v string) {
	o.ResponseKey = &v
}

func (o AadhaarxmlInitiateSessionResponse) MarshalJSON() ([]byte, error) {
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

type NullableAadhaarxmlInitiateSessionResponse struct {
	value *AadhaarxmlInitiateSessionResponse
	isSet bool
}

func (v NullableAadhaarxmlInitiateSessionResponse) Get() *AadhaarxmlInitiateSessionResponse {
	return v.value
}

func (v *NullableAadhaarxmlInitiateSessionResponse) Set(val *AadhaarxmlInitiateSessionResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAadhaarxmlInitiateSessionResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAadhaarxmlInitiateSessionResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAadhaarxmlInitiateSessionResponse(val *AadhaarxmlInitiateSessionResponse) *NullableAadhaarxmlInitiateSessionResponse {
	return &NullableAadhaarxmlInitiateSessionResponse{value: val, isSet: true}
}

func (v NullableAadhaarxmlInitiateSessionResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAadhaarxmlInitiateSessionResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


