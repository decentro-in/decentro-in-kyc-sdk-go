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

// ValidateResponseKycResultLlpMasterData struct for ValidateResponseKycResultLlpMasterData
type ValidateResponseKycResultLlpMasterData struct {
	EmailId *string `json:"emailId,omitempty"`
	RegisteredAddress *string `json:"registeredAddress,omitempty"`
	DateOfLastFinancialYearEndDateForWhichAnnualReturnFiled *string `json:"dateOfLastFinancialYearEndDateForWhichAnnualReturnFiled,omitempty"`
	DateOfLastFinancialYearEndDateForWhichStatementOfAccountsAndSolvencyFiled *string `json:"dateOfLastFinancialYearEndDateForWhichStatementOfAccountsAndSolvencyFiled,omitempty"`
	MainDivisionOfBusinessActivityToBeCarriedOutInIndia *string `json:"mainDivisionOfBusinessActivityToBeCarriedOutInIndia,omitempty"`
	PreviousFircompanyDetailifApplicable *string `json:"previousFircompanyDetailifApplicable,omitempty"`
	RocCode *string `json:"rocCode,omitempty"`
	NumberOfDesignatedPartners *string `json:"numberOfDesignatedPartners,omitempty"`
	DateOfIncorporation *string `json:"dateOfIncorporation,omitempty"`
	LlpName *string `json:"llpName,omitempty"`
	TotalObligationOfContribution *string `json:"totalObligationOfContribution,omitempty"`
	Llpin *string `json:"llpin ,omitempty"`
	LlpStatus *string `json:"llpStatus,omitempty"`
	DescriptionOfMainDivision *string `json:"descriptionOfMainDivision,omitempty"`
	NumberOfPartners *string `json:"numberOfPartners,omitempty"`
}

// NewValidateResponseKycResultLlpMasterData instantiates a new ValidateResponseKycResultLlpMasterData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewValidateResponseKycResultLlpMasterData() *ValidateResponseKycResultLlpMasterData {
	this := ValidateResponseKycResultLlpMasterData{}
	return &this
}

// NewValidateResponseKycResultLlpMasterDataWithDefaults instantiates a new ValidateResponseKycResultLlpMasterData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewValidateResponseKycResultLlpMasterDataWithDefaults() *ValidateResponseKycResultLlpMasterData {
	this := ValidateResponseKycResultLlpMasterData{}
	return &this
}

// GetEmailId returns the EmailId field value if set, zero value otherwise.
func (o *ValidateResponseKycResultLlpMasterData) GetEmailId() string {
	if o == nil || isNil(o.EmailId) {
		var ret string
		return ret
	}
	return *o.EmailId
}

// GetEmailIdOk returns a tuple with the EmailId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidateResponseKycResultLlpMasterData) GetEmailIdOk() (*string, bool) {
	if o == nil || isNil(o.EmailId) {
    return nil, false
	}
	return o.EmailId, true
}

// HasEmailId returns a boolean if a field has been set.
func (o *ValidateResponseKycResultLlpMasterData) HasEmailId() bool {
	if o != nil && !isNil(o.EmailId) {
		return true
	}

	return false
}

// SetEmailId gets a reference to the given string and assigns it to the EmailId field.
func (o *ValidateResponseKycResultLlpMasterData) SetEmailId(v string) {
	o.EmailId = &v
}

// GetRegisteredAddress returns the RegisteredAddress field value if set, zero value otherwise.
func (o *ValidateResponseKycResultLlpMasterData) GetRegisteredAddress() string {
	if o == nil || isNil(o.RegisteredAddress) {
		var ret string
		return ret
	}
	return *o.RegisteredAddress
}

// GetRegisteredAddressOk returns a tuple with the RegisteredAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidateResponseKycResultLlpMasterData) GetRegisteredAddressOk() (*string, bool) {
	if o == nil || isNil(o.RegisteredAddress) {
    return nil, false
	}
	return o.RegisteredAddress, true
}

// HasRegisteredAddress returns a boolean if a field has been set.
func (o *ValidateResponseKycResultLlpMasterData) HasRegisteredAddress() bool {
	if o != nil && !isNil(o.RegisteredAddress) {
		return true
	}

	return false
}

// SetRegisteredAddress gets a reference to the given string and assigns it to the RegisteredAddress field.
func (o *ValidateResponseKycResultLlpMasterData) SetRegisteredAddress(v string) {
	o.RegisteredAddress = &v
}

// GetDateOfLastFinancialYearEndDateForWhichAnnualReturnFiled returns the DateOfLastFinancialYearEndDateForWhichAnnualReturnFiled field value if set, zero value otherwise.
func (o *ValidateResponseKycResultLlpMasterData) GetDateOfLastFinancialYearEndDateForWhichAnnualReturnFiled() string {
	if o == nil || isNil(o.DateOfLastFinancialYearEndDateForWhichAnnualReturnFiled) {
		var ret string
		return ret
	}
	return *o.DateOfLastFinancialYearEndDateForWhichAnnualReturnFiled
}

// GetDateOfLastFinancialYearEndDateForWhichAnnualReturnFiledOk returns a tuple with the DateOfLastFinancialYearEndDateForWhichAnnualReturnFiled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidateResponseKycResultLlpMasterData) GetDateOfLastFinancialYearEndDateForWhichAnnualReturnFiledOk() (*string, bool) {
	if o == nil || isNil(o.DateOfLastFinancialYearEndDateForWhichAnnualReturnFiled) {
    return nil, false
	}
	return o.DateOfLastFinancialYearEndDateForWhichAnnualReturnFiled, true
}

// HasDateOfLastFinancialYearEndDateForWhichAnnualReturnFiled returns a boolean if a field has been set.
func (o *ValidateResponseKycResultLlpMasterData) HasDateOfLastFinancialYearEndDateForWhichAnnualReturnFiled() bool {
	if o != nil && !isNil(o.DateOfLastFinancialYearEndDateForWhichAnnualReturnFiled) {
		return true
	}

	return false
}

// SetDateOfLastFinancialYearEndDateForWhichAnnualReturnFiled gets a reference to the given string and assigns it to the DateOfLastFinancialYearEndDateForWhichAnnualReturnFiled field.
func (o *ValidateResponseKycResultLlpMasterData) SetDateOfLastFinancialYearEndDateForWhichAnnualReturnFiled(v string) {
	o.DateOfLastFinancialYearEndDateForWhichAnnualReturnFiled = &v
}

// GetDateOfLastFinancialYearEndDateForWhichStatementOfAccountsAndSolvencyFiled returns the DateOfLastFinancialYearEndDateForWhichStatementOfAccountsAndSolvencyFiled field value if set, zero value otherwise.
func (o *ValidateResponseKycResultLlpMasterData) GetDateOfLastFinancialYearEndDateForWhichStatementOfAccountsAndSolvencyFiled() string {
	if o == nil || isNil(o.DateOfLastFinancialYearEndDateForWhichStatementOfAccountsAndSolvencyFiled) {
		var ret string
		return ret
	}
	return *o.DateOfLastFinancialYearEndDateForWhichStatementOfAccountsAndSolvencyFiled
}

// GetDateOfLastFinancialYearEndDateForWhichStatementOfAccountsAndSolvencyFiledOk returns a tuple with the DateOfLastFinancialYearEndDateForWhichStatementOfAccountsAndSolvencyFiled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidateResponseKycResultLlpMasterData) GetDateOfLastFinancialYearEndDateForWhichStatementOfAccountsAndSolvencyFiledOk() (*string, bool) {
	if o == nil || isNil(o.DateOfLastFinancialYearEndDateForWhichStatementOfAccountsAndSolvencyFiled) {
    return nil, false
	}
	return o.DateOfLastFinancialYearEndDateForWhichStatementOfAccountsAndSolvencyFiled, true
}

// HasDateOfLastFinancialYearEndDateForWhichStatementOfAccountsAndSolvencyFiled returns a boolean if a field has been set.
func (o *ValidateResponseKycResultLlpMasterData) HasDateOfLastFinancialYearEndDateForWhichStatementOfAccountsAndSolvencyFiled() bool {
	if o != nil && !isNil(o.DateOfLastFinancialYearEndDateForWhichStatementOfAccountsAndSolvencyFiled) {
		return true
	}

	return false
}

// SetDateOfLastFinancialYearEndDateForWhichStatementOfAccountsAndSolvencyFiled gets a reference to the given string and assigns it to the DateOfLastFinancialYearEndDateForWhichStatementOfAccountsAndSolvencyFiled field.
func (o *ValidateResponseKycResultLlpMasterData) SetDateOfLastFinancialYearEndDateForWhichStatementOfAccountsAndSolvencyFiled(v string) {
	o.DateOfLastFinancialYearEndDateForWhichStatementOfAccountsAndSolvencyFiled = &v
}

// GetMainDivisionOfBusinessActivityToBeCarriedOutInIndia returns the MainDivisionOfBusinessActivityToBeCarriedOutInIndia field value if set, zero value otherwise.
func (o *ValidateResponseKycResultLlpMasterData) GetMainDivisionOfBusinessActivityToBeCarriedOutInIndia() string {
	if o == nil || isNil(o.MainDivisionOfBusinessActivityToBeCarriedOutInIndia) {
		var ret string
		return ret
	}
	return *o.MainDivisionOfBusinessActivityToBeCarriedOutInIndia
}

// GetMainDivisionOfBusinessActivityToBeCarriedOutInIndiaOk returns a tuple with the MainDivisionOfBusinessActivityToBeCarriedOutInIndia field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidateResponseKycResultLlpMasterData) GetMainDivisionOfBusinessActivityToBeCarriedOutInIndiaOk() (*string, bool) {
	if o == nil || isNil(o.MainDivisionOfBusinessActivityToBeCarriedOutInIndia) {
    return nil, false
	}
	return o.MainDivisionOfBusinessActivityToBeCarriedOutInIndia, true
}

// HasMainDivisionOfBusinessActivityToBeCarriedOutInIndia returns a boolean if a field has been set.
func (o *ValidateResponseKycResultLlpMasterData) HasMainDivisionOfBusinessActivityToBeCarriedOutInIndia() bool {
	if o != nil && !isNil(o.MainDivisionOfBusinessActivityToBeCarriedOutInIndia) {
		return true
	}

	return false
}

// SetMainDivisionOfBusinessActivityToBeCarriedOutInIndia gets a reference to the given string and assigns it to the MainDivisionOfBusinessActivityToBeCarriedOutInIndia field.
func (o *ValidateResponseKycResultLlpMasterData) SetMainDivisionOfBusinessActivityToBeCarriedOutInIndia(v string) {
	o.MainDivisionOfBusinessActivityToBeCarriedOutInIndia = &v
}

// GetPreviousFircompanyDetailifApplicable returns the PreviousFircompanyDetailifApplicable field value if set, zero value otherwise.
func (o *ValidateResponseKycResultLlpMasterData) GetPreviousFircompanyDetailifApplicable() string {
	if o == nil || isNil(o.PreviousFircompanyDetailifApplicable) {
		var ret string
		return ret
	}
	return *o.PreviousFircompanyDetailifApplicable
}

// GetPreviousFircompanyDetailifApplicableOk returns a tuple with the PreviousFircompanyDetailifApplicable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidateResponseKycResultLlpMasterData) GetPreviousFircompanyDetailifApplicableOk() (*string, bool) {
	if o == nil || isNil(o.PreviousFircompanyDetailifApplicable) {
    return nil, false
	}
	return o.PreviousFircompanyDetailifApplicable, true
}

// HasPreviousFircompanyDetailifApplicable returns a boolean if a field has been set.
func (o *ValidateResponseKycResultLlpMasterData) HasPreviousFircompanyDetailifApplicable() bool {
	if o != nil && !isNil(o.PreviousFircompanyDetailifApplicable) {
		return true
	}

	return false
}

// SetPreviousFircompanyDetailifApplicable gets a reference to the given string and assigns it to the PreviousFircompanyDetailifApplicable field.
func (o *ValidateResponseKycResultLlpMasterData) SetPreviousFircompanyDetailifApplicable(v string) {
	o.PreviousFircompanyDetailifApplicable = &v
}

// GetRocCode returns the RocCode field value if set, zero value otherwise.
func (o *ValidateResponseKycResultLlpMasterData) GetRocCode() string {
	if o == nil || isNil(o.RocCode) {
		var ret string
		return ret
	}
	return *o.RocCode
}

// GetRocCodeOk returns a tuple with the RocCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidateResponseKycResultLlpMasterData) GetRocCodeOk() (*string, bool) {
	if o == nil || isNil(o.RocCode) {
    return nil, false
	}
	return o.RocCode, true
}

// HasRocCode returns a boolean if a field has been set.
func (o *ValidateResponseKycResultLlpMasterData) HasRocCode() bool {
	if o != nil && !isNil(o.RocCode) {
		return true
	}

	return false
}

// SetRocCode gets a reference to the given string and assigns it to the RocCode field.
func (o *ValidateResponseKycResultLlpMasterData) SetRocCode(v string) {
	o.RocCode = &v
}

// GetNumberOfDesignatedPartners returns the NumberOfDesignatedPartners field value if set, zero value otherwise.
func (o *ValidateResponseKycResultLlpMasterData) GetNumberOfDesignatedPartners() string {
	if o == nil || isNil(o.NumberOfDesignatedPartners) {
		var ret string
		return ret
	}
	return *o.NumberOfDesignatedPartners
}

// GetNumberOfDesignatedPartnersOk returns a tuple with the NumberOfDesignatedPartners field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidateResponseKycResultLlpMasterData) GetNumberOfDesignatedPartnersOk() (*string, bool) {
	if o == nil || isNil(o.NumberOfDesignatedPartners) {
    return nil, false
	}
	return o.NumberOfDesignatedPartners, true
}

// HasNumberOfDesignatedPartners returns a boolean if a field has been set.
func (o *ValidateResponseKycResultLlpMasterData) HasNumberOfDesignatedPartners() bool {
	if o != nil && !isNil(o.NumberOfDesignatedPartners) {
		return true
	}

	return false
}

// SetNumberOfDesignatedPartners gets a reference to the given string and assigns it to the NumberOfDesignatedPartners field.
func (o *ValidateResponseKycResultLlpMasterData) SetNumberOfDesignatedPartners(v string) {
	o.NumberOfDesignatedPartners = &v
}

// GetDateOfIncorporation returns the DateOfIncorporation field value if set, zero value otherwise.
func (o *ValidateResponseKycResultLlpMasterData) GetDateOfIncorporation() string {
	if o == nil || isNil(o.DateOfIncorporation) {
		var ret string
		return ret
	}
	return *o.DateOfIncorporation
}

// GetDateOfIncorporationOk returns a tuple with the DateOfIncorporation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidateResponseKycResultLlpMasterData) GetDateOfIncorporationOk() (*string, bool) {
	if o == nil || isNil(o.DateOfIncorporation) {
    return nil, false
	}
	return o.DateOfIncorporation, true
}

// HasDateOfIncorporation returns a boolean if a field has been set.
func (o *ValidateResponseKycResultLlpMasterData) HasDateOfIncorporation() bool {
	if o != nil && !isNil(o.DateOfIncorporation) {
		return true
	}

	return false
}

// SetDateOfIncorporation gets a reference to the given string and assigns it to the DateOfIncorporation field.
func (o *ValidateResponseKycResultLlpMasterData) SetDateOfIncorporation(v string) {
	o.DateOfIncorporation = &v
}

// GetLlpName returns the LlpName field value if set, zero value otherwise.
func (o *ValidateResponseKycResultLlpMasterData) GetLlpName() string {
	if o == nil || isNil(o.LlpName) {
		var ret string
		return ret
	}
	return *o.LlpName
}

// GetLlpNameOk returns a tuple with the LlpName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidateResponseKycResultLlpMasterData) GetLlpNameOk() (*string, bool) {
	if o == nil || isNil(o.LlpName) {
    return nil, false
	}
	return o.LlpName, true
}

// HasLlpName returns a boolean if a field has been set.
func (o *ValidateResponseKycResultLlpMasterData) HasLlpName() bool {
	if o != nil && !isNil(o.LlpName) {
		return true
	}

	return false
}

// SetLlpName gets a reference to the given string and assigns it to the LlpName field.
func (o *ValidateResponseKycResultLlpMasterData) SetLlpName(v string) {
	o.LlpName = &v
}

// GetTotalObligationOfContribution returns the TotalObligationOfContribution field value if set, zero value otherwise.
func (o *ValidateResponseKycResultLlpMasterData) GetTotalObligationOfContribution() string {
	if o == nil || isNil(o.TotalObligationOfContribution) {
		var ret string
		return ret
	}
	return *o.TotalObligationOfContribution
}

// GetTotalObligationOfContributionOk returns a tuple with the TotalObligationOfContribution field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidateResponseKycResultLlpMasterData) GetTotalObligationOfContributionOk() (*string, bool) {
	if o == nil || isNil(o.TotalObligationOfContribution) {
    return nil, false
	}
	return o.TotalObligationOfContribution, true
}

// HasTotalObligationOfContribution returns a boolean if a field has been set.
func (o *ValidateResponseKycResultLlpMasterData) HasTotalObligationOfContribution() bool {
	if o != nil && !isNil(o.TotalObligationOfContribution) {
		return true
	}

	return false
}

// SetTotalObligationOfContribution gets a reference to the given string and assigns it to the TotalObligationOfContribution field.
func (o *ValidateResponseKycResultLlpMasterData) SetTotalObligationOfContribution(v string) {
	o.TotalObligationOfContribution = &v
}

// GetLlpin returns the Llpin field value if set, zero value otherwise.
func (o *ValidateResponseKycResultLlpMasterData) GetLlpin() string {
	if o == nil || isNil(o.Llpin) {
		var ret string
		return ret
	}
	return *o.Llpin
}

// GetLlpinOk returns a tuple with the Llpin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidateResponseKycResultLlpMasterData) GetLlpinOk() (*string, bool) {
	if o == nil || isNil(o.Llpin) {
    return nil, false
	}
	return o.Llpin, true
}

// HasLlpin returns a boolean if a field has been set.
func (o *ValidateResponseKycResultLlpMasterData) HasLlpin() bool {
	if o != nil && !isNil(o.Llpin) {
		return true
	}

	return false
}

// SetLlpin gets a reference to the given string and assigns it to the Llpin field.
func (o *ValidateResponseKycResultLlpMasterData) SetLlpin(v string) {
	o.Llpin = &v
}

// GetLlpStatus returns the LlpStatus field value if set, zero value otherwise.
func (o *ValidateResponseKycResultLlpMasterData) GetLlpStatus() string {
	if o == nil || isNil(o.LlpStatus) {
		var ret string
		return ret
	}
	return *o.LlpStatus
}

// GetLlpStatusOk returns a tuple with the LlpStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidateResponseKycResultLlpMasterData) GetLlpStatusOk() (*string, bool) {
	if o == nil || isNil(o.LlpStatus) {
    return nil, false
	}
	return o.LlpStatus, true
}

// HasLlpStatus returns a boolean if a field has been set.
func (o *ValidateResponseKycResultLlpMasterData) HasLlpStatus() bool {
	if o != nil && !isNil(o.LlpStatus) {
		return true
	}

	return false
}

// SetLlpStatus gets a reference to the given string and assigns it to the LlpStatus field.
func (o *ValidateResponseKycResultLlpMasterData) SetLlpStatus(v string) {
	o.LlpStatus = &v
}

// GetDescriptionOfMainDivision returns the DescriptionOfMainDivision field value if set, zero value otherwise.
func (o *ValidateResponseKycResultLlpMasterData) GetDescriptionOfMainDivision() string {
	if o == nil || isNil(o.DescriptionOfMainDivision) {
		var ret string
		return ret
	}
	return *o.DescriptionOfMainDivision
}

// GetDescriptionOfMainDivisionOk returns a tuple with the DescriptionOfMainDivision field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidateResponseKycResultLlpMasterData) GetDescriptionOfMainDivisionOk() (*string, bool) {
	if o == nil || isNil(o.DescriptionOfMainDivision) {
    return nil, false
	}
	return o.DescriptionOfMainDivision, true
}

// HasDescriptionOfMainDivision returns a boolean if a field has been set.
func (o *ValidateResponseKycResultLlpMasterData) HasDescriptionOfMainDivision() bool {
	if o != nil && !isNil(o.DescriptionOfMainDivision) {
		return true
	}

	return false
}

// SetDescriptionOfMainDivision gets a reference to the given string and assigns it to the DescriptionOfMainDivision field.
func (o *ValidateResponseKycResultLlpMasterData) SetDescriptionOfMainDivision(v string) {
	o.DescriptionOfMainDivision = &v
}

// GetNumberOfPartners returns the NumberOfPartners field value if set, zero value otherwise.
func (o *ValidateResponseKycResultLlpMasterData) GetNumberOfPartners() string {
	if o == nil || isNil(o.NumberOfPartners) {
		var ret string
		return ret
	}
	return *o.NumberOfPartners
}

// GetNumberOfPartnersOk returns a tuple with the NumberOfPartners field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidateResponseKycResultLlpMasterData) GetNumberOfPartnersOk() (*string, bool) {
	if o == nil || isNil(o.NumberOfPartners) {
    return nil, false
	}
	return o.NumberOfPartners, true
}

// HasNumberOfPartners returns a boolean if a field has been set.
func (o *ValidateResponseKycResultLlpMasterData) HasNumberOfPartners() bool {
	if o != nil && !isNil(o.NumberOfPartners) {
		return true
	}

	return false
}

// SetNumberOfPartners gets a reference to the given string and assigns it to the NumberOfPartners field.
func (o *ValidateResponseKycResultLlpMasterData) SetNumberOfPartners(v string) {
	o.NumberOfPartners = &v
}

func (o ValidateResponseKycResultLlpMasterData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.EmailId) {
		toSerialize["emailId"] = o.EmailId
	}
	if !isNil(o.RegisteredAddress) {
		toSerialize["registeredAddress"] = o.RegisteredAddress
	}
	if !isNil(o.DateOfLastFinancialYearEndDateForWhichAnnualReturnFiled) {
		toSerialize["dateOfLastFinancialYearEndDateForWhichAnnualReturnFiled"] = o.DateOfLastFinancialYearEndDateForWhichAnnualReturnFiled
	}
	if !isNil(o.DateOfLastFinancialYearEndDateForWhichStatementOfAccountsAndSolvencyFiled) {
		toSerialize["dateOfLastFinancialYearEndDateForWhichStatementOfAccountsAndSolvencyFiled"] = o.DateOfLastFinancialYearEndDateForWhichStatementOfAccountsAndSolvencyFiled
	}
	if !isNil(o.MainDivisionOfBusinessActivityToBeCarriedOutInIndia) {
		toSerialize["mainDivisionOfBusinessActivityToBeCarriedOutInIndia"] = o.MainDivisionOfBusinessActivityToBeCarriedOutInIndia
	}
	if !isNil(o.PreviousFircompanyDetailifApplicable) {
		toSerialize["previousFircompanyDetailifApplicable"] = o.PreviousFircompanyDetailifApplicable
	}
	if !isNil(o.RocCode) {
		toSerialize["rocCode"] = o.RocCode
	}
	if !isNil(o.NumberOfDesignatedPartners) {
		toSerialize["numberOfDesignatedPartners"] = o.NumberOfDesignatedPartners
	}
	if !isNil(o.DateOfIncorporation) {
		toSerialize["dateOfIncorporation"] = o.DateOfIncorporation
	}
	if !isNil(o.LlpName) {
		toSerialize["llpName"] = o.LlpName
	}
	if !isNil(o.TotalObligationOfContribution) {
		toSerialize["totalObligationOfContribution"] = o.TotalObligationOfContribution
	}
	if !isNil(o.Llpin) {
		toSerialize["llpin "] = o.Llpin
	}
	if !isNil(o.LlpStatus) {
		toSerialize["llpStatus"] = o.LlpStatus
	}
	if !isNil(o.DescriptionOfMainDivision) {
		toSerialize["descriptionOfMainDivision"] = o.DescriptionOfMainDivision
	}
	if !isNil(o.NumberOfPartners) {
		toSerialize["numberOfPartners"] = o.NumberOfPartners
	}
	return json.Marshal(toSerialize)
}

type NullableValidateResponseKycResultLlpMasterData struct {
	value *ValidateResponseKycResultLlpMasterData
	isSet bool
}

func (v NullableValidateResponseKycResultLlpMasterData) Get() *ValidateResponseKycResultLlpMasterData {
	return v.value
}

func (v *NullableValidateResponseKycResultLlpMasterData) Set(val *ValidateResponseKycResultLlpMasterData) {
	v.value = val
	v.isSet = true
}

func (v NullableValidateResponseKycResultLlpMasterData) IsSet() bool {
	return v.isSet
}

func (v *NullableValidateResponseKycResultLlpMasterData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableValidateResponseKycResultLlpMasterData(val *ValidateResponseKycResultLlpMasterData) *NullableValidateResponseKycResultLlpMasterData {
	return &NullableValidateResponseKycResultLlpMasterData{value: val, isSet: true}
}

func (v NullableValidateResponseKycResultLlpMasterData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableValidateResponseKycResultLlpMasterData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


