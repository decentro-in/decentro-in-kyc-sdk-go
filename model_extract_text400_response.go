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

// ExtractText400Response struct for ExtractText400Response
type ExtractText400Response struct {
	Status *string `json:"status,omitempty"`
	OcrStatus *string `json:"ocrStatus,omitempty"`
	KycStatus *string `json:"kycStatus,omitempty"`
	Error *ExtractText400ResponseError `json:"error,omitempty"`
	RequestTimestamp *string `json:"requestTimestamp,omitempty"`
	ResponseTimestamp *string `json:"responseTimestamp,omitempty"`
	DecentroTxnId *string `json:"decentroTxnId,omitempty"`
}

// NewExtractText400Response instantiates a new ExtractText400Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExtractText400Response() *ExtractText400Response {
	this := ExtractText400Response{}
	return &this
}

// NewExtractText400ResponseWithDefaults instantiates a new ExtractText400Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExtractText400ResponseWithDefaults() *ExtractText400Response {
	this := ExtractText400Response{}
	return &this
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ExtractText400Response) GetStatus() string {
	if o == nil || isNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExtractText400Response) GetStatusOk() (*string, bool) {
	if o == nil || isNil(o.Status) {
    return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ExtractText400Response) HasStatus() bool {
	if o != nil && !isNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *ExtractText400Response) SetStatus(v string) {
	o.Status = &v
}

// GetOcrStatus returns the OcrStatus field value if set, zero value otherwise.
func (o *ExtractText400Response) GetOcrStatus() string {
	if o == nil || isNil(o.OcrStatus) {
		var ret string
		return ret
	}
	return *o.OcrStatus
}

// GetOcrStatusOk returns a tuple with the OcrStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExtractText400Response) GetOcrStatusOk() (*string, bool) {
	if o == nil || isNil(o.OcrStatus) {
    return nil, false
	}
	return o.OcrStatus, true
}

// HasOcrStatus returns a boolean if a field has been set.
func (o *ExtractText400Response) HasOcrStatus() bool {
	if o != nil && !isNil(o.OcrStatus) {
		return true
	}

	return false
}

// SetOcrStatus gets a reference to the given string and assigns it to the OcrStatus field.
func (o *ExtractText400Response) SetOcrStatus(v string) {
	o.OcrStatus = &v
}

// GetKycStatus returns the KycStatus field value if set, zero value otherwise.
func (o *ExtractText400Response) GetKycStatus() string {
	if o == nil || isNil(o.KycStatus) {
		var ret string
		return ret
	}
	return *o.KycStatus
}

// GetKycStatusOk returns a tuple with the KycStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExtractText400Response) GetKycStatusOk() (*string, bool) {
	if o == nil || isNil(o.KycStatus) {
    return nil, false
	}
	return o.KycStatus, true
}

// HasKycStatus returns a boolean if a field has been set.
func (o *ExtractText400Response) HasKycStatus() bool {
	if o != nil && !isNil(o.KycStatus) {
		return true
	}

	return false
}

// SetKycStatus gets a reference to the given string and assigns it to the KycStatus field.
func (o *ExtractText400Response) SetKycStatus(v string) {
	o.KycStatus = &v
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *ExtractText400Response) GetError() ExtractText400ResponseError {
	if o == nil || isNil(o.Error) {
		var ret ExtractText400ResponseError
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExtractText400Response) GetErrorOk() (*ExtractText400ResponseError, bool) {
	if o == nil || isNil(o.Error) {
    return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *ExtractText400Response) HasError() bool {
	if o != nil && !isNil(o.Error) {
		return true
	}

	return false
}

// SetError gets a reference to the given ExtractText400ResponseError and assigns it to the Error field.
func (o *ExtractText400Response) SetError(v ExtractText400ResponseError) {
	o.Error = &v
}

// GetRequestTimestamp returns the RequestTimestamp field value if set, zero value otherwise.
func (o *ExtractText400Response) GetRequestTimestamp() string {
	if o == nil || isNil(o.RequestTimestamp) {
		var ret string
		return ret
	}
	return *o.RequestTimestamp
}

// GetRequestTimestampOk returns a tuple with the RequestTimestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExtractText400Response) GetRequestTimestampOk() (*string, bool) {
	if o == nil || isNil(o.RequestTimestamp) {
    return nil, false
	}
	return o.RequestTimestamp, true
}

// HasRequestTimestamp returns a boolean if a field has been set.
func (o *ExtractText400Response) HasRequestTimestamp() bool {
	if o != nil && !isNil(o.RequestTimestamp) {
		return true
	}

	return false
}

// SetRequestTimestamp gets a reference to the given string and assigns it to the RequestTimestamp field.
func (o *ExtractText400Response) SetRequestTimestamp(v string) {
	o.RequestTimestamp = &v
}

// GetResponseTimestamp returns the ResponseTimestamp field value if set, zero value otherwise.
func (o *ExtractText400Response) GetResponseTimestamp() string {
	if o == nil || isNil(o.ResponseTimestamp) {
		var ret string
		return ret
	}
	return *o.ResponseTimestamp
}

// GetResponseTimestampOk returns a tuple with the ResponseTimestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExtractText400Response) GetResponseTimestampOk() (*string, bool) {
	if o == nil || isNil(o.ResponseTimestamp) {
    return nil, false
	}
	return o.ResponseTimestamp, true
}

// HasResponseTimestamp returns a boolean if a field has been set.
func (o *ExtractText400Response) HasResponseTimestamp() bool {
	if o != nil && !isNil(o.ResponseTimestamp) {
		return true
	}

	return false
}

// SetResponseTimestamp gets a reference to the given string and assigns it to the ResponseTimestamp field.
func (o *ExtractText400Response) SetResponseTimestamp(v string) {
	o.ResponseTimestamp = &v
}

// GetDecentroTxnId returns the DecentroTxnId field value if set, zero value otherwise.
func (o *ExtractText400Response) GetDecentroTxnId() string {
	if o == nil || isNil(o.DecentroTxnId) {
		var ret string
		return ret
	}
	return *o.DecentroTxnId
}

// GetDecentroTxnIdOk returns a tuple with the DecentroTxnId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExtractText400Response) GetDecentroTxnIdOk() (*string, bool) {
	if o == nil || isNil(o.DecentroTxnId) {
    return nil, false
	}
	return o.DecentroTxnId, true
}

// HasDecentroTxnId returns a boolean if a field has been set.
func (o *ExtractText400Response) HasDecentroTxnId() bool {
	if o != nil && !isNil(o.DecentroTxnId) {
		return true
	}

	return false
}

// SetDecentroTxnId gets a reference to the given string and assigns it to the DecentroTxnId field.
func (o *ExtractText400Response) SetDecentroTxnId(v string) {
	o.DecentroTxnId = &v
}

func (o ExtractText400Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !isNil(o.OcrStatus) {
		toSerialize["ocrStatus"] = o.OcrStatus
	}
	if !isNil(o.KycStatus) {
		toSerialize["kycStatus"] = o.KycStatus
	}
	if !isNil(o.Error) {
		toSerialize["error"] = o.Error
	}
	if !isNil(o.RequestTimestamp) {
		toSerialize["requestTimestamp"] = o.RequestTimestamp
	}
	if !isNil(o.ResponseTimestamp) {
		toSerialize["responseTimestamp"] = o.ResponseTimestamp
	}
	if !isNil(o.DecentroTxnId) {
		toSerialize["decentroTxnId"] = o.DecentroTxnId
	}
	return json.Marshal(toSerialize)
}

type NullableExtractText400Response struct {
	value *ExtractText400Response
	isSet bool
}

func (v NullableExtractText400Response) Get() *ExtractText400Response {
	return v.value
}

func (v *NullableExtractText400Response) Set(val *ExtractText400Response) {
	v.value = val
	v.isSet = true
}

func (v NullableExtractText400Response) IsSet() bool {
	return v.isSet
}

func (v *NullableExtractText400Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExtractText400Response(val *ExtractText400Response) *NullableExtractText400Response {
	return &NullableExtractText400Response{value: val, isSet: true}
}

func (v NullableExtractText400Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExtractText400Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


