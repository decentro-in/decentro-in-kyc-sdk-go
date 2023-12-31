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

// DigilockerIssuedFilesResponseDataInner struct for DigilockerIssuedFilesResponseDataInner
type DigilockerIssuedFilesResponseDataInner struct {
	Description *string `json:"description,omitempty"`
	Name *string `json:"name,omitempty"`
	Type *string `json:"type,omitempty"`
	Size *string `json:"size,omitempty"`
	Date *string `json:"date,omitempty"`
	Parent *string `json:"parent,omitempty"`
	Mime []string `json:"mime,omitempty"`
	Uri *string `json:"uri,omitempty"`
	Doctype *string `json:"doctype,omitempty"`
	Issuerid *string `json:"issuerid,omitempty"`
	Issuer *string `json:"issuer,omitempty"`
}

// NewDigilockerIssuedFilesResponseDataInner instantiates a new DigilockerIssuedFilesResponseDataInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDigilockerIssuedFilesResponseDataInner() *DigilockerIssuedFilesResponseDataInner {
	this := DigilockerIssuedFilesResponseDataInner{}
	return &this
}

// NewDigilockerIssuedFilesResponseDataInnerWithDefaults instantiates a new DigilockerIssuedFilesResponseDataInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDigilockerIssuedFilesResponseDataInnerWithDefaults() *DigilockerIssuedFilesResponseDataInner {
	this := DigilockerIssuedFilesResponseDataInner{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *DigilockerIssuedFilesResponseDataInner) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DigilockerIssuedFilesResponseDataInner) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
    return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *DigilockerIssuedFilesResponseDataInner) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *DigilockerIssuedFilesResponseDataInner) SetDescription(v string) {
	o.Description = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *DigilockerIssuedFilesResponseDataInner) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DigilockerIssuedFilesResponseDataInner) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *DigilockerIssuedFilesResponseDataInner) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *DigilockerIssuedFilesResponseDataInner) SetName(v string) {
	o.Name = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *DigilockerIssuedFilesResponseDataInner) GetType() string {
	if o == nil || isNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DigilockerIssuedFilesResponseDataInner) GetTypeOk() (*string, bool) {
	if o == nil || isNil(o.Type) {
    return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *DigilockerIssuedFilesResponseDataInner) HasType() bool {
	if o != nil && !isNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *DigilockerIssuedFilesResponseDataInner) SetType(v string) {
	o.Type = &v
}

// GetSize returns the Size field value if set, zero value otherwise.
func (o *DigilockerIssuedFilesResponseDataInner) GetSize() string {
	if o == nil || isNil(o.Size) {
		var ret string
		return ret
	}
	return *o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DigilockerIssuedFilesResponseDataInner) GetSizeOk() (*string, bool) {
	if o == nil || isNil(o.Size) {
    return nil, false
	}
	return o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *DigilockerIssuedFilesResponseDataInner) HasSize() bool {
	if o != nil && !isNil(o.Size) {
		return true
	}

	return false
}

// SetSize gets a reference to the given string and assigns it to the Size field.
func (o *DigilockerIssuedFilesResponseDataInner) SetSize(v string) {
	o.Size = &v
}

// GetDate returns the Date field value if set, zero value otherwise.
func (o *DigilockerIssuedFilesResponseDataInner) GetDate() string {
	if o == nil || isNil(o.Date) {
		var ret string
		return ret
	}
	return *o.Date
}

// GetDateOk returns a tuple with the Date field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DigilockerIssuedFilesResponseDataInner) GetDateOk() (*string, bool) {
	if o == nil || isNil(o.Date) {
    return nil, false
	}
	return o.Date, true
}

// HasDate returns a boolean if a field has been set.
func (o *DigilockerIssuedFilesResponseDataInner) HasDate() bool {
	if o != nil && !isNil(o.Date) {
		return true
	}

	return false
}

// SetDate gets a reference to the given string and assigns it to the Date field.
func (o *DigilockerIssuedFilesResponseDataInner) SetDate(v string) {
	o.Date = &v
}

// GetParent returns the Parent field value if set, zero value otherwise.
func (o *DigilockerIssuedFilesResponseDataInner) GetParent() string {
	if o == nil || isNil(o.Parent) {
		var ret string
		return ret
	}
	return *o.Parent
}

// GetParentOk returns a tuple with the Parent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DigilockerIssuedFilesResponseDataInner) GetParentOk() (*string, bool) {
	if o == nil || isNil(o.Parent) {
    return nil, false
	}
	return o.Parent, true
}

// HasParent returns a boolean if a field has been set.
func (o *DigilockerIssuedFilesResponseDataInner) HasParent() bool {
	if o != nil && !isNil(o.Parent) {
		return true
	}

	return false
}

// SetParent gets a reference to the given string and assigns it to the Parent field.
func (o *DigilockerIssuedFilesResponseDataInner) SetParent(v string) {
	o.Parent = &v
}

// GetMime returns the Mime field value if set, zero value otherwise.
func (o *DigilockerIssuedFilesResponseDataInner) GetMime() []string {
	if o == nil || isNil(o.Mime) {
		var ret []string
		return ret
	}
	return o.Mime
}

// GetMimeOk returns a tuple with the Mime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DigilockerIssuedFilesResponseDataInner) GetMimeOk() ([]string, bool) {
	if o == nil || isNil(o.Mime) {
    return nil, false
	}
	return o.Mime, true
}

// HasMime returns a boolean if a field has been set.
func (o *DigilockerIssuedFilesResponseDataInner) HasMime() bool {
	if o != nil && !isNil(o.Mime) {
		return true
	}

	return false
}

// SetMime gets a reference to the given []string and assigns it to the Mime field.
func (o *DigilockerIssuedFilesResponseDataInner) SetMime(v []string) {
	o.Mime = v
}

// GetUri returns the Uri field value if set, zero value otherwise.
func (o *DigilockerIssuedFilesResponseDataInner) GetUri() string {
	if o == nil || isNil(o.Uri) {
		var ret string
		return ret
	}
	return *o.Uri
}

// GetUriOk returns a tuple with the Uri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DigilockerIssuedFilesResponseDataInner) GetUriOk() (*string, bool) {
	if o == nil || isNil(o.Uri) {
    return nil, false
	}
	return o.Uri, true
}

// HasUri returns a boolean if a field has been set.
func (o *DigilockerIssuedFilesResponseDataInner) HasUri() bool {
	if o != nil && !isNil(o.Uri) {
		return true
	}

	return false
}

// SetUri gets a reference to the given string and assigns it to the Uri field.
func (o *DigilockerIssuedFilesResponseDataInner) SetUri(v string) {
	o.Uri = &v
}

// GetDoctype returns the Doctype field value if set, zero value otherwise.
func (o *DigilockerIssuedFilesResponseDataInner) GetDoctype() string {
	if o == nil || isNil(o.Doctype) {
		var ret string
		return ret
	}
	return *o.Doctype
}

// GetDoctypeOk returns a tuple with the Doctype field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DigilockerIssuedFilesResponseDataInner) GetDoctypeOk() (*string, bool) {
	if o == nil || isNil(o.Doctype) {
    return nil, false
	}
	return o.Doctype, true
}

// HasDoctype returns a boolean if a field has been set.
func (o *DigilockerIssuedFilesResponseDataInner) HasDoctype() bool {
	if o != nil && !isNil(o.Doctype) {
		return true
	}

	return false
}

// SetDoctype gets a reference to the given string and assigns it to the Doctype field.
func (o *DigilockerIssuedFilesResponseDataInner) SetDoctype(v string) {
	o.Doctype = &v
}

// GetIssuerid returns the Issuerid field value if set, zero value otherwise.
func (o *DigilockerIssuedFilesResponseDataInner) GetIssuerid() string {
	if o == nil || isNil(o.Issuerid) {
		var ret string
		return ret
	}
	return *o.Issuerid
}

// GetIssueridOk returns a tuple with the Issuerid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DigilockerIssuedFilesResponseDataInner) GetIssueridOk() (*string, bool) {
	if o == nil || isNil(o.Issuerid) {
    return nil, false
	}
	return o.Issuerid, true
}

// HasIssuerid returns a boolean if a field has been set.
func (o *DigilockerIssuedFilesResponseDataInner) HasIssuerid() bool {
	if o != nil && !isNil(o.Issuerid) {
		return true
	}

	return false
}

// SetIssuerid gets a reference to the given string and assigns it to the Issuerid field.
func (o *DigilockerIssuedFilesResponseDataInner) SetIssuerid(v string) {
	o.Issuerid = &v
}

// GetIssuer returns the Issuer field value if set, zero value otherwise.
func (o *DigilockerIssuedFilesResponseDataInner) GetIssuer() string {
	if o == nil || isNil(o.Issuer) {
		var ret string
		return ret
	}
	return *o.Issuer
}

// GetIssuerOk returns a tuple with the Issuer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DigilockerIssuedFilesResponseDataInner) GetIssuerOk() (*string, bool) {
	if o == nil || isNil(o.Issuer) {
    return nil, false
	}
	return o.Issuer, true
}

// HasIssuer returns a boolean if a field has been set.
func (o *DigilockerIssuedFilesResponseDataInner) HasIssuer() bool {
	if o != nil && !isNil(o.Issuer) {
		return true
	}

	return false
}

// SetIssuer gets a reference to the given string and assigns it to the Issuer field.
func (o *DigilockerIssuedFilesResponseDataInner) SetIssuer(v string) {
	o.Issuer = &v
}

func (o DigilockerIssuedFilesResponseDataInner) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !isNil(o.Size) {
		toSerialize["size"] = o.Size
	}
	if !isNil(o.Date) {
		toSerialize["date"] = o.Date
	}
	if !isNil(o.Parent) {
		toSerialize["parent"] = o.Parent
	}
	if !isNil(o.Mime) {
		toSerialize["mime"] = o.Mime
	}
	if !isNil(o.Uri) {
		toSerialize["uri"] = o.Uri
	}
	if !isNil(o.Doctype) {
		toSerialize["doctype"] = o.Doctype
	}
	if !isNil(o.Issuerid) {
		toSerialize["issuerid"] = o.Issuerid
	}
	if !isNil(o.Issuer) {
		toSerialize["issuer"] = o.Issuer
	}
	return json.Marshal(toSerialize)
}

type NullableDigilockerIssuedFilesResponseDataInner struct {
	value *DigilockerIssuedFilesResponseDataInner
	isSet bool
}

func (v NullableDigilockerIssuedFilesResponseDataInner) Get() *DigilockerIssuedFilesResponseDataInner {
	return v.value
}

func (v *NullableDigilockerIssuedFilesResponseDataInner) Set(val *DigilockerIssuedFilesResponseDataInner) {
	v.value = val
	v.isSet = true
}

func (v NullableDigilockerIssuedFilesResponseDataInner) IsSet() bool {
	return v.isSet
}

func (v *NullableDigilockerIssuedFilesResponseDataInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDigilockerIssuedFilesResponseDataInner(val *DigilockerIssuedFilesResponseDataInner) *NullableDigilockerIssuedFilesResponseDataInner {
	return &NullableDigilockerIssuedFilesResponseDataInner{value: val, isSet: true}
}

func (v NullableDigilockerIssuedFilesResponseDataInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDigilockerIssuedFilesResponseDataInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


