# Validate400Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **string** |  | [optional] 
**KycStatus** | Pointer to **string** |  | [optional] 
**Error** | Pointer to [**Validate400ResponseError**](Validate400ResponseError.md) |  | [optional] 
**RequestTimestamp** | Pointer to **string** |  | [optional] 
**ResponseTimestamp** | Pointer to **string** |  | [optional] 
**DecentroTxnId** | Pointer to **string** |  | [optional] 

## Methods

### NewValidate400Response

`func NewValidate400Response() *Validate400Response`

NewValidate400Response instantiates a new Validate400Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewValidate400ResponseWithDefaults

`func NewValidate400ResponseWithDefaults() *Validate400Response`

NewValidate400ResponseWithDefaults instantiates a new Validate400Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *Validate400Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Validate400Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Validate400Response) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Validate400Response) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetKycStatus

`func (o *Validate400Response) GetKycStatus() string`

GetKycStatus returns the KycStatus field if non-nil, zero value otherwise.

### GetKycStatusOk

`func (o *Validate400Response) GetKycStatusOk() (*string, bool)`

GetKycStatusOk returns a tuple with the KycStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKycStatus

`func (o *Validate400Response) SetKycStatus(v string)`

SetKycStatus sets KycStatus field to given value.

### HasKycStatus

`func (o *Validate400Response) HasKycStatus() bool`

HasKycStatus returns a boolean if a field has been set.

### GetError

`func (o *Validate400Response) GetError() Validate400ResponseError`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *Validate400Response) GetErrorOk() (*Validate400ResponseError, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *Validate400Response) SetError(v Validate400ResponseError)`

SetError sets Error field to given value.

### HasError

`func (o *Validate400Response) HasError() bool`

HasError returns a boolean if a field has been set.

### GetRequestTimestamp

`func (o *Validate400Response) GetRequestTimestamp() string`

GetRequestTimestamp returns the RequestTimestamp field if non-nil, zero value otherwise.

### GetRequestTimestampOk

`func (o *Validate400Response) GetRequestTimestampOk() (*string, bool)`

GetRequestTimestampOk returns a tuple with the RequestTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestTimestamp

`func (o *Validate400Response) SetRequestTimestamp(v string)`

SetRequestTimestamp sets RequestTimestamp field to given value.

### HasRequestTimestamp

`func (o *Validate400Response) HasRequestTimestamp() bool`

HasRequestTimestamp returns a boolean if a field has been set.

### GetResponseTimestamp

`func (o *Validate400Response) GetResponseTimestamp() string`

GetResponseTimestamp returns the ResponseTimestamp field if non-nil, zero value otherwise.

### GetResponseTimestampOk

`func (o *Validate400Response) GetResponseTimestampOk() (*string, bool)`

GetResponseTimestampOk returns a tuple with the ResponseTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseTimestamp

`func (o *Validate400Response) SetResponseTimestamp(v string)`

SetResponseTimestamp sets ResponseTimestamp field to given value.

### HasResponseTimestamp

`func (o *Validate400Response) HasResponseTimestamp() bool`

HasResponseTimestamp returns a boolean if a field has been set.

### GetDecentroTxnId

`func (o *Validate400Response) GetDecentroTxnId() string`

GetDecentroTxnId returns the DecentroTxnId field if non-nil, zero value otherwise.

### GetDecentroTxnIdOk

`func (o *Validate400Response) GetDecentroTxnIdOk() (*string, bool)`

GetDecentroTxnIdOk returns a tuple with the DecentroTxnId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecentroTxnId

`func (o *Validate400Response) SetDecentroTxnId(v string)`

SetDecentroTxnId sets DecentroTxnId field to given value.

### HasDecentroTxnId

`func (o *Validate400Response) HasDecentroTxnId() bool`

HasDecentroTxnId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


