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

// DigilockerDownloadFile400Response struct for DigilockerDownloadFile400Response
type DigilockerDownloadFile400Response struct {
	Data *DigilockerDownloadFile400ResponseData `json:"data,omitempty"`
	DecentroTxnId *string `json:"decentroTxnId,omitempty"`
	Message *string `json:"message,omitempty"`
	ResponseCode *string `json:"responseCode,omitempty"`
	ResponseKey *string `json:"responseKey,omitempty"`
	Status *string `json:"status,omitempty"`
}

// NewDigilockerDownloadFile400Response instantiates a new DigilockerDownloadFile400Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDigilockerDownloadFile400Response() *DigilockerDownloadFile400Response {
	this := DigilockerDownloadFile400Response{}
	return &this
}

// NewDigilockerDownloadFile400ResponseWithDefaults instantiates a new DigilockerDownloadFile400Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDigilockerDownloadFile400ResponseWithDefaults() *DigilockerDownloadFile400Response {
	this := DigilockerDownloadFile400Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *DigilockerDownloadFile400Response) GetData() DigilockerDownloadFile400ResponseData {
	if o == nil || isNil(o.Data) {
		var ret DigilockerDownloadFile400ResponseData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DigilockerDownloadFile400Response) GetDataOk() (*DigilockerDownloadFile400ResponseData, bool) {
	if o == nil || isNil(o.Data) {
    return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *DigilockerDownloadFile400Response) HasData() bool {
	if o != nil && !isNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given DigilockerDownloadFile400ResponseData and assigns it to the Data field.
func (o *DigilockerDownloadFile400Response) SetData(v DigilockerDownloadFile400ResponseData) {
	o.Data = &v
}

// GetDecentroTxnId returns the DecentroTxnId field value if set, zero value otherwise.
func (o *DigilockerDownloadFile400Response) GetDecentroTxnId() string {
	if o == nil || isNil(o.DecentroTxnId) {
		var ret string
		return ret
	}
	return *o.DecentroTxnId
}

// GetDecentroTxnIdOk returns a tuple with the DecentroTxnId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DigilockerDownloadFile400Response) GetDecentroTxnIdOk() (*string, bool) {
	if o == nil || isNil(o.DecentroTxnId) {
    return nil, false
	}
	return o.DecentroTxnId, true
}

// HasDecentroTxnId returns a boolean if a field has been set.
func (o *DigilockerDownloadFile400Response) HasDecentroTxnId() bool {
	if o != nil && !isNil(o.DecentroTxnId) {
		return true
	}

	return false
}

// SetDecentroTxnId gets a reference to the given string and assigns it to the DecentroTxnId field.
func (o *DigilockerDownloadFile400Response) SetDecentroTxnId(v string) {
	o.DecentroTxnId = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *DigilockerDownloadFile400Response) GetMessage() string {
	if o == nil || isNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DigilockerDownloadFile400Response) GetMessageOk() (*string, bool) {
	if o == nil || isNil(o.Message) {
    return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *DigilockerDownloadFile400Response) HasMessage() bool {
	if o != nil && !isNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *DigilockerDownloadFile400Response) SetMessage(v string) {
	o.Message = &v
}

// GetResponseCode returns the ResponseCode field value if set, zero value otherwise.
func (o *DigilockerDownloadFile400Response) GetResponseCode() string {
	if o == nil || isNil(o.ResponseCode) {
		var ret string
		return ret
	}
	return *o.ResponseCode
}

// GetResponseCodeOk returns a tuple with the ResponseCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DigilockerDownloadFile400Response) GetResponseCodeOk() (*string, bool) {
	if o == nil || isNil(o.ResponseCode) {
    return nil, false
	}
	return o.ResponseCode, true
}

// HasResponseCode returns a boolean if a field has been set.
func (o *DigilockerDownloadFile400Response) HasResponseCode() bool {
	if o != nil && !isNil(o.ResponseCode) {
		return true
	}

	return false
}

// SetResponseCode gets a reference to the given string and assigns it to the ResponseCode field.
func (o *DigilockerDownloadFile400Response) SetResponseCode(v string) {
	o.ResponseCode = &v
}

// GetResponseKey returns the ResponseKey field value if set, zero value otherwise.
func (o *DigilockerDownloadFile400Response) GetResponseKey() string {
	if o == nil || isNil(o.ResponseKey) {
		var ret string
		return ret
	}
	return *o.ResponseKey
}

// GetResponseKeyOk returns a tuple with the ResponseKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DigilockerDownloadFile400Response) GetResponseKeyOk() (*string, bool) {
	if o == nil || isNil(o.ResponseKey) {
    return nil, false
	}
	return o.ResponseKey, true
}

// HasResponseKey returns a boolean if a field has been set.
func (o *DigilockerDownloadFile400Response) HasResponseKey() bool {
	if o != nil && !isNil(o.ResponseKey) {
		return true
	}

	return false
}

// SetResponseKey gets a reference to the given string and assigns it to the ResponseKey field.
func (o *DigilockerDownloadFile400Response) SetResponseKey(v string) {
	o.ResponseKey = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *DigilockerDownloadFile400Response) GetStatus() string {
	if o == nil || isNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DigilockerDownloadFile400Response) GetStatusOk() (*string, bool) {
	if o == nil || isNil(o.Status) {
    return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *DigilockerDownloadFile400Response) HasStatus() bool {
	if o != nil && !isNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *DigilockerDownloadFile400Response) SetStatus(v string) {
	o.Status = &v
}

func (o DigilockerDownloadFile400Response) MarshalJSON() ([]byte, error) {
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
	if !isNil(o.ResponseKey) {
		toSerialize["responseKey"] = o.ResponseKey
	}
	if !isNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	return json.Marshal(toSerialize)
}

type NullableDigilockerDownloadFile400Response struct {
	value *DigilockerDownloadFile400Response
	isSet bool
}

func (v NullableDigilockerDownloadFile400Response) Get() *DigilockerDownloadFile400Response {
	return v.value
}

func (v *NullableDigilockerDownloadFile400Response) Set(val *DigilockerDownloadFile400Response) {
	v.value = val
	v.isSet = true
}

func (v NullableDigilockerDownloadFile400Response) IsSet() bool {
	return v.isSet
}

func (v *NullableDigilockerDownloadFile400Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDigilockerDownloadFile400Response(val *DigilockerDownloadFile400Response) *NullableDigilockerDownloadFile400Response {
	return &NullableDigilockerDownloadFile400Response{value: val, isSet: true}
}

func (v NullableDigilockerDownloadFile400Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDigilockerDownloadFile400Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


