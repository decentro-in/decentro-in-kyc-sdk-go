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

// DigilockerIssuedFilesResponse struct for DigilockerIssuedFilesResponse
type DigilockerIssuedFilesResponse struct {
	DecentroTxnId *string `json:"decentroTxnId,omitempty"`
	Status *string `json:"status,omitempty"`
	ResponseCode *string `json:"responseCode,omitempty"`
	Message *string `json:"message,omitempty"`
	Data []DigilockerIssuedFilesResponseDataInner `json:"data,omitempty"`
	ResponseKey *string `json:"responseKey,omitempty"`
}

// NewDigilockerIssuedFilesResponse instantiates a new DigilockerIssuedFilesResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDigilockerIssuedFilesResponse() *DigilockerIssuedFilesResponse {
	this := DigilockerIssuedFilesResponse{}
	return &this
}

// NewDigilockerIssuedFilesResponseWithDefaults instantiates a new DigilockerIssuedFilesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDigilockerIssuedFilesResponseWithDefaults() *DigilockerIssuedFilesResponse {
	this := DigilockerIssuedFilesResponse{}
	return &this
}

// GetDecentroTxnId returns the DecentroTxnId field value if set, zero value otherwise.
func (o *DigilockerIssuedFilesResponse) GetDecentroTxnId() string {
	if o == nil || isNil(o.DecentroTxnId) {
		var ret string
		return ret
	}
	return *o.DecentroTxnId
}

// GetDecentroTxnIdOk returns a tuple with the DecentroTxnId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DigilockerIssuedFilesResponse) GetDecentroTxnIdOk() (*string, bool) {
	if o == nil || isNil(o.DecentroTxnId) {
    return nil, false
	}
	return o.DecentroTxnId, true
}

// HasDecentroTxnId returns a boolean if a field has been set.
func (o *DigilockerIssuedFilesResponse) HasDecentroTxnId() bool {
	if o != nil && !isNil(o.DecentroTxnId) {
		return true
	}

	return false
}

// SetDecentroTxnId gets a reference to the given string and assigns it to the DecentroTxnId field.
func (o *DigilockerIssuedFilesResponse) SetDecentroTxnId(v string) {
	o.DecentroTxnId = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *DigilockerIssuedFilesResponse) GetStatus() string {
	if o == nil || isNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DigilockerIssuedFilesResponse) GetStatusOk() (*string, bool) {
	if o == nil || isNil(o.Status) {
    return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *DigilockerIssuedFilesResponse) HasStatus() bool {
	if o != nil && !isNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *DigilockerIssuedFilesResponse) SetStatus(v string) {
	o.Status = &v
}

// GetResponseCode returns the ResponseCode field value if set, zero value otherwise.
func (o *DigilockerIssuedFilesResponse) GetResponseCode() string {
	if o == nil || isNil(o.ResponseCode) {
		var ret string
		return ret
	}
	return *o.ResponseCode
}

// GetResponseCodeOk returns a tuple with the ResponseCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DigilockerIssuedFilesResponse) GetResponseCodeOk() (*string, bool) {
	if o == nil || isNil(o.ResponseCode) {
    return nil, false
	}
	return o.ResponseCode, true
}

// HasResponseCode returns a boolean if a field has been set.
func (o *DigilockerIssuedFilesResponse) HasResponseCode() bool {
	if o != nil && !isNil(o.ResponseCode) {
		return true
	}

	return false
}

// SetResponseCode gets a reference to the given string and assigns it to the ResponseCode field.
func (o *DigilockerIssuedFilesResponse) SetResponseCode(v string) {
	o.ResponseCode = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *DigilockerIssuedFilesResponse) GetMessage() string {
	if o == nil || isNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DigilockerIssuedFilesResponse) GetMessageOk() (*string, bool) {
	if o == nil || isNil(o.Message) {
    return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *DigilockerIssuedFilesResponse) HasMessage() bool {
	if o != nil && !isNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *DigilockerIssuedFilesResponse) SetMessage(v string) {
	o.Message = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *DigilockerIssuedFilesResponse) GetData() []DigilockerIssuedFilesResponseDataInner {
	if o == nil || isNil(o.Data) {
		var ret []DigilockerIssuedFilesResponseDataInner
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DigilockerIssuedFilesResponse) GetDataOk() ([]DigilockerIssuedFilesResponseDataInner, bool) {
	if o == nil || isNil(o.Data) {
    return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *DigilockerIssuedFilesResponse) HasData() bool {
	if o != nil && !isNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []DigilockerIssuedFilesResponseDataInner and assigns it to the Data field.
func (o *DigilockerIssuedFilesResponse) SetData(v []DigilockerIssuedFilesResponseDataInner) {
	o.Data = v
}

// GetResponseKey returns the ResponseKey field value if set, zero value otherwise.
func (o *DigilockerIssuedFilesResponse) GetResponseKey() string {
	if o == nil || isNil(o.ResponseKey) {
		var ret string
		return ret
	}
	return *o.ResponseKey
}

// GetResponseKeyOk returns a tuple with the ResponseKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DigilockerIssuedFilesResponse) GetResponseKeyOk() (*string, bool) {
	if o == nil || isNil(o.ResponseKey) {
    return nil, false
	}
	return o.ResponseKey, true
}

// HasResponseKey returns a boolean if a field has been set.
func (o *DigilockerIssuedFilesResponse) HasResponseKey() bool {
	if o != nil && !isNil(o.ResponseKey) {
		return true
	}

	return false
}

// SetResponseKey gets a reference to the given string and assigns it to the ResponseKey field.
func (o *DigilockerIssuedFilesResponse) SetResponseKey(v string) {
	o.ResponseKey = &v
}

func (o DigilockerIssuedFilesResponse) MarshalJSON() ([]byte, error) {
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

type NullableDigilockerIssuedFilesResponse struct {
	value *DigilockerIssuedFilesResponse
	isSet bool
}

func (v NullableDigilockerIssuedFilesResponse) Get() *DigilockerIssuedFilesResponse {
	return v.value
}

func (v *NullableDigilockerIssuedFilesResponse) Set(val *DigilockerIssuedFilesResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableDigilockerIssuedFilesResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableDigilockerIssuedFilesResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDigilockerIssuedFilesResponse(val *DigilockerIssuedFilesResponse) *NullableDigilockerIssuedFilesResponse {
	return &NullableDigilockerIssuedFilesResponse{value: val, isSet: true}
}

func (v NullableDigilockerIssuedFilesResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDigilockerIssuedFilesResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


