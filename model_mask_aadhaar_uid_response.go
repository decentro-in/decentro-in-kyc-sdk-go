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

// MaskAadhaarUidResponse struct for MaskAadhaarUidResponse
type MaskAadhaarUidResponse struct {
	Data *MaskAadhaarUidResponseData `json:"data,omitempty"`
	DecentroTxnId *string `json:"decentroTxnId,omitempty"`
	Message *string `json:"message,omitempty"`
	ResponseCode *string `json:"responseCode,omitempty"`
	Status *string `json:"status,omitempty"`
}

// NewMaskAadhaarUidResponse instantiates a new MaskAadhaarUidResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMaskAadhaarUidResponse() *MaskAadhaarUidResponse {
	this := MaskAadhaarUidResponse{}
	return &this
}

// NewMaskAadhaarUidResponseWithDefaults instantiates a new MaskAadhaarUidResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMaskAadhaarUidResponseWithDefaults() *MaskAadhaarUidResponse {
	this := MaskAadhaarUidResponse{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *MaskAadhaarUidResponse) GetData() MaskAadhaarUidResponseData {
	if o == nil || isNil(o.Data) {
		var ret MaskAadhaarUidResponseData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MaskAadhaarUidResponse) GetDataOk() (*MaskAadhaarUidResponseData, bool) {
	if o == nil || isNil(o.Data) {
    return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *MaskAadhaarUidResponse) HasData() bool {
	if o != nil && !isNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given MaskAadhaarUidResponseData and assigns it to the Data field.
func (o *MaskAadhaarUidResponse) SetData(v MaskAadhaarUidResponseData) {
	o.Data = &v
}

// GetDecentroTxnId returns the DecentroTxnId field value if set, zero value otherwise.
func (o *MaskAadhaarUidResponse) GetDecentroTxnId() string {
	if o == nil || isNil(o.DecentroTxnId) {
		var ret string
		return ret
	}
	return *o.DecentroTxnId
}

// GetDecentroTxnIdOk returns a tuple with the DecentroTxnId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MaskAadhaarUidResponse) GetDecentroTxnIdOk() (*string, bool) {
	if o == nil || isNil(o.DecentroTxnId) {
    return nil, false
	}
	return o.DecentroTxnId, true
}

// HasDecentroTxnId returns a boolean if a field has been set.
func (o *MaskAadhaarUidResponse) HasDecentroTxnId() bool {
	if o != nil && !isNil(o.DecentroTxnId) {
		return true
	}

	return false
}

// SetDecentroTxnId gets a reference to the given string and assigns it to the DecentroTxnId field.
func (o *MaskAadhaarUidResponse) SetDecentroTxnId(v string) {
	o.DecentroTxnId = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *MaskAadhaarUidResponse) GetMessage() string {
	if o == nil || isNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MaskAadhaarUidResponse) GetMessageOk() (*string, bool) {
	if o == nil || isNil(o.Message) {
    return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *MaskAadhaarUidResponse) HasMessage() bool {
	if o != nil && !isNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *MaskAadhaarUidResponse) SetMessage(v string) {
	o.Message = &v
}

// GetResponseCode returns the ResponseCode field value if set, zero value otherwise.
func (o *MaskAadhaarUidResponse) GetResponseCode() string {
	if o == nil || isNil(o.ResponseCode) {
		var ret string
		return ret
	}
	return *o.ResponseCode
}

// GetResponseCodeOk returns a tuple with the ResponseCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MaskAadhaarUidResponse) GetResponseCodeOk() (*string, bool) {
	if o == nil || isNil(o.ResponseCode) {
    return nil, false
	}
	return o.ResponseCode, true
}

// HasResponseCode returns a boolean if a field has been set.
func (o *MaskAadhaarUidResponse) HasResponseCode() bool {
	if o != nil && !isNil(o.ResponseCode) {
		return true
	}

	return false
}

// SetResponseCode gets a reference to the given string and assigns it to the ResponseCode field.
func (o *MaskAadhaarUidResponse) SetResponseCode(v string) {
	o.ResponseCode = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *MaskAadhaarUidResponse) GetStatus() string {
	if o == nil || isNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MaskAadhaarUidResponse) GetStatusOk() (*string, bool) {
	if o == nil || isNil(o.Status) {
    return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *MaskAadhaarUidResponse) HasStatus() bool {
	if o != nil && !isNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *MaskAadhaarUidResponse) SetStatus(v string) {
	o.Status = &v
}

func (o MaskAadhaarUidResponse) MarshalJSON() ([]byte, error) {
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

type NullableMaskAadhaarUidResponse struct {
	value *MaskAadhaarUidResponse
	isSet bool
}

func (v NullableMaskAadhaarUidResponse) Get() *MaskAadhaarUidResponse {
	return v.value
}

func (v *NullableMaskAadhaarUidResponse) Set(val *MaskAadhaarUidResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableMaskAadhaarUidResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableMaskAadhaarUidResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMaskAadhaarUidResponse(val *MaskAadhaarUidResponse) *NullableMaskAadhaarUidResponse {
	return &NullableMaskAadhaarUidResponse{value: val, isSet: true}
}

func (v NullableMaskAadhaarUidResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMaskAadhaarUidResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


