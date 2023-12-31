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

// DigilockerIssuedFiles400ResponseDataInner struct for DigilockerIssuedFiles400ResponseDataInner
type DigilockerIssuedFiles400ResponseDataInner struct {
	Description *string `json:"description,omitempty"`
	Date *string `json:"date,omitempty"`
	Doctype *string `json:"doctype,omitempty"`
	Issuer *string `json:"issuer,omitempty"`
	Issuerid *string `json:"issuerid,omitempty"`
	Mime []string `json:"mime,omitempty"`
	Name *string `json:"name,omitempty"`
	Parent *string `json:"parent,omitempty"`
	Size *string `json:"size,omitempty"`
	Type *string `json:"type,omitempty"`
	Uri *string `json:"uri,omitempty"`
}

// NewDigilockerIssuedFiles400ResponseDataInner instantiates a new DigilockerIssuedFiles400ResponseDataInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDigilockerIssuedFiles400ResponseDataInner() *DigilockerIssuedFiles400ResponseDataInner {
	this := DigilockerIssuedFiles400ResponseDataInner{}
	return &this
}

// NewDigilockerIssuedFiles400ResponseDataInnerWithDefaults instantiates a new DigilockerIssuedFiles400ResponseDataInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDigilockerIssuedFiles400ResponseDataInnerWithDefaults() *DigilockerIssuedFiles400ResponseDataInner {
	this := DigilockerIssuedFiles400ResponseDataInner{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *DigilockerIssuedFiles400ResponseDataInner) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DigilockerIssuedFiles400ResponseDataInner) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
    return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *DigilockerIssuedFiles400ResponseDataInner) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *DigilockerIssuedFiles400ResponseDataInner) SetDescription(v string) {
	o.Description = &v
}

// GetDate returns the Date field value if set, zero value otherwise.
func (o *DigilockerIssuedFiles400ResponseDataInner) GetDate() string {
	if o == nil || isNil(o.Date) {
		var ret string
		return ret
	}
	return *o.Date
}

// GetDateOk returns a tuple with the Date field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DigilockerIssuedFiles400ResponseDataInner) GetDateOk() (*string, bool) {
	if o == nil || isNil(o.Date) {
    return nil, false
	}
	return o.Date, true
}

// HasDate returns a boolean if a field has been set.
func (o *DigilockerIssuedFiles400ResponseDataInner) HasDate() bool {
	if o != nil && !isNil(o.Date) {
		return true
	}

	return false
}

// SetDate gets a reference to the given string and assigns it to the Date field.
func (o *DigilockerIssuedFiles400ResponseDataInner) SetDate(v string) {
	o.Date = &v
}

// GetDoctype returns the Doctype field value if set, zero value otherwise.
func (o *DigilockerIssuedFiles400ResponseDataInner) GetDoctype() string {
	if o == nil || isNil(o.Doctype) {
		var ret string
		return ret
	}
	return *o.Doctype
}

// GetDoctypeOk returns a tuple with the Doctype field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DigilockerIssuedFiles400ResponseDataInner) GetDoctypeOk() (*string, bool) {
	if o == nil || isNil(o.Doctype) {
    return nil, false
	}
	return o.Doctype, true
}

// HasDoctype returns a boolean if a field has been set.
func (o *DigilockerIssuedFiles400ResponseDataInner) HasDoctype() bool {
	if o != nil && !isNil(o.Doctype) {
		return true
	}

	return false
}

// SetDoctype gets a reference to the given string and assigns it to the Doctype field.
func (o *DigilockerIssuedFiles400ResponseDataInner) SetDoctype(v string) {
	o.Doctype = &v
}

// GetIssuer returns the Issuer field value if set, zero value otherwise.
func (o *DigilockerIssuedFiles400ResponseDataInner) GetIssuer() string {
	if o == nil || isNil(o.Issuer) {
		var ret string
		return ret
	}
	return *o.Issuer
}

// GetIssuerOk returns a tuple with the Issuer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DigilockerIssuedFiles400ResponseDataInner) GetIssuerOk() (*string, bool) {
	if o == nil || isNil(o.Issuer) {
    return nil, false
	}
	return o.Issuer, true
}

// HasIssuer returns a boolean if a field has been set.
func (o *DigilockerIssuedFiles400ResponseDataInner) HasIssuer() bool {
	if o != nil && !isNil(o.Issuer) {
		return true
	}

	return false
}

// SetIssuer gets a reference to the given string and assigns it to the Issuer field.
func (o *DigilockerIssuedFiles400ResponseDataInner) SetIssuer(v string) {
	o.Issuer = &v
}

// GetIssuerid returns the Issuerid field value if set, zero value otherwise.
func (o *DigilockerIssuedFiles400ResponseDataInner) GetIssuerid() string {
	if o == nil || isNil(o.Issuerid) {
		var ret string
		return ret
	}
	return *o.Issuerid
}

// GetIssueridOk returns a tuple with the Issuerid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DigilockerIssuedFiles400ResponseDataInner) GetIssueridOk() (*string, bool) {
	if o == nil || isNil(o.Issuerid) {
    return nil, false
	}
	return o.Issuerid, true
}

// HasIssuerid returns a boolean if a field has been set.
func (o *DigilockerIssuedFiles400ResponseDataInner) HasIssuerid() bool {
	if o != nil && !isNil(o.Issuerid) {
		return true
	}

	return false
}

// SetIssuerid gets a reference to the given string and assigns it to the Issuerid field.
func (o *DigilockerIssuedFiles400ResponseDataInner) SetIssuerid(v string) {
	o.Issuerid = &v
}

// GetMime returns the Mime field value if set, zero value otherwise.
func (o *DigilockerIssuedFiles400ResponseDataInner) GetMime() []string {
	if o == nil || isNil(o.Mime) {
		var ret []string
		return ret
	}
	return o.Mime
}

// GetMimeOk returns a tuple with the Mime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DigilockerIssuedFiles400ResponseDataInner) GetMimeOk() ([]string, bool) {
	if o == nil || isNil(o.Mime) {
    return nil, false
	}
	return o.Mime, true
}

// HasMime returns a boolean if a field has been set.
func (o *DigilockerIssuedFiles400ResponseDataInner) HasMime() bool {
	if o != nil && !isNil(o.Mime) {
		return true
	}

	return false
}

// SetMime gets a reference to the given []string and assigns it to the Mime field.
func (o *DigilockerIssuedFiles400ResponseDataInner) SetMime(v []string) {
	o.Mime = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *DigilockerIssuedFiles400ResponseDataInner) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DigilockerIssuedFiles400ResponseDataInner) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *DigilockerIssuedFiles400ResponseDataInner) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *DigilockerIssuedFiles400ResponseDataInner) SetName(v string) {
	o.Name = &v
}

// GetParent returns the Parent field value if set, zero value otherwise.
func (o *DigilockerIssuedFiles400ResponseDataInner) GetParent() string {
	if o == nil || isNil(o.Parent) {
		var ret string
		return ret
	}
	return *o.Parent
}

// GetParentOk returns a tuple with the Parent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DigilockerIssuedFiles400ResponseDataInner) GetParentOk() (*string, bool) {
	if o == nil || isNil(o.Parent) {
    return nil, false
	}
	return o.Parent, true
}

// HasParent returns a boolean if a field has been set.
func (o *DigilockerIssuedFiles400ResponseDataInner) HasParent() bool {
	if o != nil && !isNil(o.Parent) {
		return true
	}

	return false
}

// SetParent gets a reference to the given string and assigns it to the Parent field.
func (o *DigilockerIssuedFiles400ResponseDataInner) SetParent(v string) {
	o.Parent = &v
}

// GetSize returns the Size field value if set, zero value otherwise.
func (o *DigilockerIssuedFiles400ResponseDataInner) GetSize() string {
	if o == nil || isNil(o.Size) {
		var ret string
		return ret
	}
	return *o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DigilockerIssuedFiles400ResponseDataInner) GetSizeOk() (*string, bool) {
	if o == nil || isNil(o.Size) {
    return nil, false
	}
	return o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *DigilockerIssuedFiles400ResponseDataInner) HasSize() bool {
	if o != nil && !isNil(o.Size) {
		return true
	}

	return false
}

// SetSize gets a reference to the given string and assigns it to the Size field.
func (o *DigilockerIssuedFiles400ResponseDataInner) SetSize(v string) {
	o.Size = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *DigilockerIssuedFiles400ResponseDataInner) GetType() string {
	if o == nil || isNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DigilockerIssuedFiles400ResponseDataInner) GetTypeOk() (*string, bool) {
	if o == nil || isNil(o.Type) {
    return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *DigilockerIssuedFiles400ResponseDataInner) HasType() bool {
	if o != nil && !isNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *DigilockerIssuedFiles400ResponseDataInner) SetType(v string) {
	o.Type = &v
}

// GetUri returns the Uri field value if set, zero value otherwise.
func (o *DigilockerIssuedFiles400ResponseDataInner) GetUri() string {
	if o == nil || isNil(o.Uri) {
		var ret string
		return ret
	}
	return *o.Uri
}

// GetUriOk returns a tuple with the Uri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DigilockerIssuedFiles400ResponseDataInner) GetUriOk() (*string, bool) {
	if o == nil || isNil(o.Uri) {
    return nil, false
	}
	return o.Uri, true
}

// HasUri returns a boolean if a field has been set.
func (o *DigilockerIssuedFiles400ResponseDataInner) HasUri() bool {
	if o != nil && !isNil(o.Uri) {
		return true
	}

	return false
}

// SetUri gets a reference to the given string and assigns it to the Uri field.
func (o *DigilockerIssuedFiles400ResponseDataInner) SetUri(v string) {
	o.Uri = &v
}

func (o DigilockerIssuedFiles400ResponseDataInner) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !isNil(o.Date) {
		toSerialize["date"] = o.Date
	}
	if !isNil(o.Doctype) {
		toSerialize["doctype"] = o.Doctype
	}
	if !isNil(o.Issuer) {
		toSerialize["issuer"] = o.Issuer
	}
	if !isNil(o.Issuerid) {
		toSerialize["issuerid"] = o.Issuerid
	}
	if !isNil(o.Mime) {
		toSerialize["mime"] = o.Mime
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Parent) {
		toSerialize["parent"] = o.Parent
	}
	if !isNil(o.Size) {
		toSerialize["size"] = o.Size
	}
	if !isNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !isNil(o.Uri) {
		toSerialize["uri"] = o.Uri
	}
	return json.Marshal(toSerialize)
}

type NullableDigilockerIssuedFiles400ResponseDataInner struct {
	value *DigilockerIssuedFiles400ResponseDataInner
	isSet bool
}

func (v NullableDigilockerIssuedFiles400ResponseDataInner) Get() *DigilockerIssuedFiles400ResponseDataInner {
	return v.value
}

func (v *NullableDigilockerIssuedFiles400ResponseDataInner) Set(val *DigilockerIssuedFiles400ResponseDataInner) {
	v.value = val
	v.isSet = true
}

func (v NullableDigilockerIssuedFiles400ResponseDataInner) IsSet() bool {
	return v.isSet
}

func (v *NullableDigilockerIssuedFiles400ResponseDataInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDigilockerIssuedFiles400ResponseDataInner(val *DigilockerIssuedFiles400ResponseDataInner) *NullableDigilockerIssuedFiles400ResponseDataInner {
	return &NullableDigilockerIssuedFiles400ResponseDataInner{value: val, isSet: true}
}

func (v NullableDigilockerIssuedFiles400ResponseDataInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDigilockerIssuedFiles400ResponseDataInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


