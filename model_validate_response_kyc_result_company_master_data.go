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

// ValidateResponseKycResultCompanyMasterData struct for ValidateResponseKycResultCompanyMasterData
type ValidateResponseKycResultCompanyMasterData struct {
	CompanyCategory *string `json:"companyCategory,omitempty"`
	EmailId *string `json:"emailId,omitempty"`
	ClassOfCompany *string `json:"classOfCompany,omitempty"`
	NumberOfMembersApplicableInCaseOfCompanyWithoutShareCapital *string `json:"numberOfMembersApplicableInCaseOfCompanyWithoutShareCapital,omitempty"`
	AddressOtherThanRegisteredOfficeWhereAllOrAnyBooksOfAccountAndPapersAreMaintained *string `json:"addressOtherThanRegisteredOfficeWhereAllOrAnyBooksOfAccountAndPapersAreMaintained,omitempty"`
	DateOfLastAgm *string `json:"dateOfLastAgm,omitempty"`
	RegisteredAddress *string `json:"registeredAddress,omitempty"`
	ActiveCompliance *string `json:"activeCompliance,omitempty"`
	RegistrationNumber *string `json:"registrationNumber,omitempty"`
	PaidUpCapitalInInr *string `json:"paidUpCapitalInInr,omitempty"`
	WhetherListedOrNot *string `json:"whetherListedOrNot,omitempty"`
	SuspendedAtStockExchange *string `json:"suspendedAtStockExchange,omitempty"`
	CompanySubcategory *string `json:"companySubcategory,omitempty"`
	AuthorisedCapitalInInr *string `json:"authorisedCapitalInInr,omitempty"`
	CompanyStatusForEFiling *string `json:"companyStatusForEFiling,omitempty"`
	RocCode *string `json:"rocCode,omitempty"`
	DateOfBalanceSheet *string `json:"dateOfBalanceSheet,omitempty"`
	DateOfIncorporation *string `json:"dateOfIncorporation,omitempty"`
	Cin *string `json:"cin ,omitempty"`
	CompanyName *string `json:"companyName,omitempty"`
}

// NewValidateResponseKycResultCompanyMasterData instantiates a new ValidateResponseKycResultCompanyMasterData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewValidateResponseKycResultCompanyMasterData() *ValidateResponseKycResultCompanyMasterData {
	this := ValidateResponseKycResultCompanyMasterData{}
	return &this
}

// NewValidateResponseKycResultCompanyMasterDataWithDefaults instantiates a new ValidateResponseKycResultCompanyMasterData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewValidateResponseKycResultCompanyMasterDataWithDefaults() *ValidateResponseKycResultCompanyMasterData {
	this := ValidateResponseKycResultCompanyMasterData{}
	return &this
}

// GetCompanyCategory returns the CompanyCategory field value if set, zero value otherwise.
func (o *ValidateResponseKycResultCompanyMasterData) GetCompanyCategory() string {
	if o == nil || isNil(o.CompanyCategory) {
		var ret string
		return ret
	}
	return *o.CompanyCategory
}

// GetCompanyCategoryOk returns a tuple with the CompanyCategory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidateResponseKycResultCompanyMasterData) GetCompanyCategoryOk() (*string, bool) {
	if o == nil || isNil(o.CompanyCategory) {
    return nil, false
	}
	return o.CompanyCategory, true
}

// HasCompanyCategory returns a boolean if a field has been set.
func (o *ValidateResponseKycResultCompanyMasterData) HasCompanyCategory() bool {
	if o != nil && !isNil(o.CompanyCategory) {
		return true
	}

	return false
}

// SetCompanyCategory gets a reference to the given string and assigns it to the CompanyCategory field.
func (o *ValidateResponseKycResultCompanyMasterData) SetCompanyCategory(v string) {
	o.CompanyCategory = &v
}

// GetEmailId returns the EmailId field value if set, zero value otherwise.
func (o *ValidateResponseKycResultCompanyMasterData) GetEmailId() string {
	if o == nil || isNil(o.EmailId) {
		var ret string
		return ret
	}
	return *o.EmailId
}

// GetEmailIdOk returns a tuple with the EmailId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidateResponseKycResultCompanyMasterData) GetEmailIdOk() (*string, bool) {
	if o == nil || isNil(o.EmailId) {
    return nil, false
	}
	return o.EmailId, true
}

// HasEmailId returns a boolean if a field has been set.
func (o *ValidateResponseKycResultCompanyMasterData) HasEmailId() bool {
	if o != nil && !isNil(o.EmailId) {
		return true
	}

	return false
}

// SetEmailId gets a reference to the given string and assigns it to the EmailId field.
func (o *ValidateResponseKycResultCompanyMasterData) SetEmailId(v string) {
	o.EmailId = &v
}

// GetClassOfCompany returns the ClassOfCompany field value if set, zero value otherwise.
func (o *ValidateResponseKycResultCompanyMasterData) GetClassOfCompany() string {
	if o == nil || isNil(o.ClassOfCompany) {
		var ret string
		return ret
	}
	return *o.ClassOfCompany
}

// GetClassOfCompanyOk returns a tuple with the ClassOfCompany field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidateResponseKycResultCompanyMasterData) GetClassOfCompanyOk() (*string, bool) {
	if o == nil || isNil(o.ClassOfCompany) {
    return nil, false
	}
	return o.ClassOfCompany, true
}

// HasClassOfCompany returns a boolean if a field has been set.
func (o *ValidateResponseKycResultCompanyMasterData) HasClassOfCompany() bool {
	if o != nil && !isNil(o.ClassOfCompany) {
		return true
	}

	return false
}

// SetClassOfCompany gets a reference to the given string and assigns it to the ClassOfCompany field.
func (o *ValidateResponseKycResultCompanyMasterData) SetClassOfCompany(v string) {
	o.ClassOfCompany = &v
}

// GetNumberOfMembersApplicableInCaseOfCompanyWithoutShareCapital returns the NumberOfMembersApplicableInCaseOfCompanyWithoutShareCapital field value if set, zero value otherwise.
func (o *ValidateResponseKycResultCompanyMasterData) GetNumberOfMembersApplicableInCaseOfCompanyWithoutShareCapital() string {
	if o == nil || isNil(o.NumberOfMembersApplicableInCaseOfCompanyWithoutShareCapital) {
		var ret string
		return ret
	}
	return *o.NumberOfMembersApplicableInCaseOfCompanyWithoutShareCapital
}

// GetNumberOfMembersApplicableInCaseOfCompanyWithoutShareCapitalOk returns a tuple with the NumberOfMembersApplicableInCaseOfCompanyWithoutShareCapital field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidateResponseKycResultCompanyMasterData) GetNumberOfMembersApplicableInCaseOfCompanyWithoutShareCapitalOk() (*string, bool) {
	if o == nil || isNil(o.NumberOfMembersApplicableInCaseOfCompanyWithoutShareCapital) {
    return nil, false
	}
	return o.NumberOfMembersApplicableInCaseOfCompanyWithoutShareCapital, true
}

// HasNumberOfMembersApplicableInCaseOfCompanyWithoutShareCapital returns a boolean if a field has been set.
func (o *ValidateResponseKycResultCompanyMasterData) HasNumberOfMembersApplicableInCaseOfCompanyWithoutShareCapital() bool {
	if o != nil && !isNil(o.NumberOfMembersApplicableInCaseOfCompanyWithoutShareCapital) {
		return true
	}

	return false
}

// SetNumberOfMembersApplicableInCaseOfCompanyWithoutShareCapital gets a reference to the given string and assigns it to the NumberOfMembersApplicableInCaseOfCompanyWithoutShareCapital field.
func (o *ValidateResponseKycResultCompanyMasterData) SetNumberOfMembersApplicableInCaseOfCompanyWithoutShareCapital(v string) {
	o.NumberOfMembersApplicableInCaseOfCompanyWithoutShareCapital = &v
}

// GetAddressOtherThanRegisteredOfficeWhereAllOrAnyBooksOfAccountAndPapersAreMaintained returns the AddressOtherThanRegisteredOfficeWhereAllOrAnyBooksOfAccountAndPapersAreMaintained field value if set, zero value otherwise.
func (o *ValidateResponseKycResultCompanyMasterData) GetAddressOtherThanRegisteredOfficeWhereAllOrAnyBooksOfAccountAndPapersAreMaintained() string {
	if o == nil || isNil(o.AddressOtherThanRegisteredOfficeWhereAllOrAnyBooksOfAccountAndPapersAreMaintained) {
		var ret string
		return ret
	}
	return *o.AddressOtherThanRegisteredOfficeWhereAllOrAnyBooksOfAccountAndPapersAreMaintained
}

// GetAddressOtherThanRegisteredOfficeWhereAllOrAnyBooksOfAccountAndPapersAreMaintainedOk returns a tuple with the AddressOtherThanRegisteredOfficeWhereAllOrAnyBooksOfAccountAndPapersAreMaintained field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidateResponseKycResultCompanyMasterData) GetAddressOtherThanRegisteredOfficeWhereAllOrAnyBooksOfAccountAndPapersAreMaintainedOk() (*string, bool) {
	if o == nil || isNil(o.AddressOtherThanRegisteredOfficeWhereAllOrAnyBooksOfAccountAndPapersAreMaintained) {
    return nil, false
	}
	return o.AddressOtherThanRegisteredOfficeWhereAllOrAnyBooksOfAccountAndPapersAreMaintained, true
}

// HasAddressOtherThanRegisteredOfficeWhereAllOrAnyBooksOfAccountAndPapersAreMaintained returns a boolean if a field has been set.
func (o *ValidateResponseKycResultCompanyMasterData) HasAddressOtherThanRegisteredOfficeWhereAllOrAnyBooksOfAccountAndPapersAreMaintained() bool {
	if o != nil && !isNil(o.AddressOtherThanRegisteredOfficeWhereAllOrAnyBooksOfAccountAndPapersAreMaintained) {
		return true
	}

	return false
}

// SetAddressOtherThanRegisteredOfficeWhereAllOrAnyBooksOfAccountAndPapersAreMaintained gets a reference to the given string and assigns it to the AddressOtherThanRegisteredOfficeWhereAllOrAnyBooksOfAccountAndPapersAreMaintained field.
func (o *ValidateResponseKycResultCompanyMasterData) SetAddressOtherThanRegisteredOfficeWhereAllOrAnyBooksOfAccountAndPapersAreMaintained(v string) {
	o.AddressOtherThanRegisteredOfficeWhereAllOrAnyBooksOfAccountAndPapersAreMaintained = &v
}

// GetDateOfLastAgm returns the DateOfLastAgm field value if set, zero value otherwise.
func (o *ValidateResponseKycResultCompanyMasterData) GetDateOfLastAgm() string {
	if o == nil || isNil(o.DateOfLastAgm) {
		var ret string
		return ret
	}
	return *o.DateOfLastAgm
}

// GetDateOfLastAgmOk returns a tuple with the DateOfLastAgm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidateResponseKycResultCompanyMasterData) GetDateOfLastAgmOk() (*string, bool) {
	if o == nil || isNil(o.DateOfLastAgm) {
    return nil, false
	}
	return o.DateOfLastAgm, true
}

// HasDateOfLastAgm returns a boolean if a field has been set.
func (o *ValidateResponseKycResultCompanyMasterData) HasDateOfLastAgm() bool {
	if o != nil && !isNil(o.DateOfLastAgm) {
		return true
	}

	return false
}

// SetDateOfLastAgm gets a reference to the given string and assigns it to the DateOfLastAgm field.
func (o *ValidateResponseKycResultCompanyMasterData) SetDateOfLastAgm(v string) {
	o.DateOfLastAgm = &v
}

// GetRegisteredAddress returns the RegisteredAddress field value if set, zero value otherwise.
func (o *ValidateResponseKycResultCompanyMasterData) GetRegisteredAddress() string {
	if o == nil || isNil(o.RegisteredAddress) {
		var ret string
		return ret
	}
	return *o.RegisteredAddress
}

// GetRegisteredAddressOk returns a tuple with the RegisteredAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidateResponseKycResultCompanyMasterData) GetRegisteredAddressOk() (*string, bool) {
	if o == nil || isNil(o.RegisteredAddress) {
    return nil, false
	}
	return o.RegisteredAddress, true
}

// HasRegisteredAddress returns a boolean if a field has been set.
func (o *ValidateResponseKycResultCompanyMasterData) HasRegisteredAddress() bool {
	if o != nil && !isNil(o.RegisteredAddress) {
		return true
	}

	return false
}

// SetRegisteredAddress gets a reference to the given string and assigns it to the RegisteredAddress field.
func (o *ValidateResponseKycResultCompanyMasterData) SetRegisteredAddress(v string) {
	o.RegisteredAddress = &v
}

// GetActiveCompliance returns the ActiveCompliance field value if set, zero value otherwise.
func (o *ValidateResponseKycResultCompanyMasterData) GetActiveCompliance() string {
	if o == nil || isNil(o.ActiveCompliance) {
		var ret string
		return ret
	}
	return *o.ActiveCompliance
}

// GetActiveComplianceOk returns a tuple with the ActiveCompliance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidateResponseKycResultCompanyMasterData) GetActiveComplianceOk() (*string, bool) {
	if o == nil || isNil(o.ActiveCompliance) {
    return nil, false
	}
	return o.ActiveCompliance, true
}

// HasActiveCompliance returns a boolean if a field has been set.
func (o *ValidateResponseKycResultCompanyMasterData) HasActiveCompliance() bool {
	if o != nil && !isNil(o.ActiveCompliance) {
		return true
	}

	return false
}

// SetActiveCompliance gets a reference to the given string and assigns it to the ActiveCompliance field.
func (o *ValidateResponseKycResultCompanyMasterData) SetActiveCompliance(v string) {
	o.ActiveCompliance = &v
}

// GetRegistrationNumber returns the RegistrationNumber field value if set, zero value otherwise.
func (o *ValidateResponseKycResultCompanyMasterData) GetRegistrationNumber() string {
	if o == nil || isNil(o.RegistrationNumber) {
		var ret string
		return ret
	}
	return *o.RegistrationNumber
}

// GetRegistrationNumberOk returns a tuple with the RegistrationNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidateResponseKycResultCompanyMasterData) GetRegistrationNumberOk() (*string, bool) {
	if o == nil || isNil(o.RegistrationNumber) {
    return nil, false
	}
	return o.RegistrationNumber, true
}

// HasRegistrationNumber returns a boolean if a field has been set.
func (o *ValidateResponseKycResultCompanyMasterData) HasRegistrationNumber() bool {
	if o != nil && !isNil(o.RegistrationNumber) {
		return true
	}

	return false
}

// SetRegistrationNumber gets a reference to the given string and assigns it to the RegistrationNumber field.
func (o *ValidateResponseKycResultCompanyMasterData) SetRegistrationNumber(v string) {
	o.RegistrationNumber = &v
}

// GetPaidUpCapitalInInr returns the PaidUpCapitalInInr field value if set, zero value otherwise.
func (o *ValidateResponseKycResultCompanyMasterData) GetPaidUpCapitalInInr() string {
	if o == nil || isNil(o.PaidUpCapitalInInr) {
		var ret string
		return ret
	}
	return *o.PaidUpCapitalInInr
}

// GetPaidUpCapitalInInrOk returns a tuple with the PaidUpCapitalInInr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidateResponseKycResultCompanyMasterData) GetPaidUpCapitalInInrOk() (*string, bool) {
	if o == nil || isNil(o.PaidUpCapitalInInr) {
    return nil, false
	}
	return o.PaidUpCapitalInInr, true
}

// HasPaidUpCapitalInInr returns a boolean if a field has been set.
func (o *ValidateResponseKycResultCompanyMasterData) HasPaidUpCapitalInInr() bool {
	if o != nil && !isNil(o.PaidUpCapitalInInr) {
		return true
	}

	return false
}

// SetPaidUpCapitalInInr gets a reference to the given string and assigns it to the PaidUpCapitalInInr field.
func (o *ValidateResponseKycResultCompanyMasterData) SetPaidUpCapitalInInr(v string) {
	o.PaidUpCapitalInInr = &v
}

// GetWhetherListedOrNot returns the WhetherListedOrNot field value if set, zero value otherwise.
func (o *ValidateResponseKycResultCompanyMasterData) GetWhetherListedOrNot() string {
	if o == nil || isNil(o.WhetherListedOrNot) {
		var ret string
		return ret
	}
	return *o.WhetherListedOrNot
}

// GetWhetherListedOrNotOk returns a tuple with the WhetherListedOrNot field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidateResponseKycResultCompanyMasterData) GetWhetherListedOrNotOk() (*string, bool) {
	if o == nil || isNil(o.WhetherListedOrNot) {
    return nil, false
	}
	return o.WhetherListedOrNot, true
}

// HasWhetherListedOrNot returns a boolean if a field has been set.
func (o *ValidateResponseKycResultCompanyMasterData) HasWhetherListedOrNot() bool {
	if o != nil && !isNil(o.WhetherListedOrNot) {
		return true
	}

	return false
}

// SetWhetherListedOrNot gets a reference to the given string and assigns it to the WhetherListedOrNot field.
func (o *ValidateResponseKycResultCompanyMasterData) SetWhetherListedOrNot(v string) {
	o.WhetherListedOrNot = &v
}

// GetSuspendedAtStockExchange returns the SuspendedAtStockExchange field value if set, zero value otherwise.
func (o *ValidateResponseKycResultCompanyMasterData) GetSuspendedAtStockExchange() string {
	if o == nil || isNil(o.SuspendedAtStockExchange) {
		var ret string
		return ret
	}
	return *o.SuspendedAtStockExchange
}

// GetSuspendedAtStockExchangeOk returns a tuple with the SuspendedAtStockExchange field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidateResponseKycResultCompanyMasterData) GetSuspendedAtStockExchangeOk() (*string, bool) {
	if o == nil || isNil(o.SuspendedAtStockExchange) {
    return nil, false
	}
	return o.SuspendedAtStockExchange, true
}

// HasSuspendedAtStockExchange returns a boolean if a field has been set.
func (o *ValidateResponseKycResultCompanyMasterData) HasSuspendedAtStockExchange() bool {
	if o != nil && !isNil(o.SuspendedAtStockExchange) {
		return true
	}

	return false
}

// SetSuspendedAtStockExchange gets a reference to the given string and assigns it to the SuspendedAtStockExchange field.
func (o *ValidateResponseKycResultCompanyMasterData) SetSuspendedAtStockExchange(v string) {
	o.SuspendedAtStockExchange = &v
}

// GetCompanySubcategory returns the CompanySubcategory field value if set, zero value otherwise.
func (o *ValidateResponseKycResultCompanyMasterData) GetCompanySubcategory() string {
	if o == nil || isNil(o.CompanySubcategory) {
		var ret string
		return ret
	}
	return *o.CompanySubcategory
}

// GetCompanySubcategoryOk returns a tuple with the CompanySubcategory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidateResponseKycResultCompanyMasterData) GetCompanySubcategoryOk() (*string, bool) {
	if o == nil || isNil(o.CompanySubcategory) {
    return nil, false
	}
	return o.CompanySubcategory, true
}

// HasCompanySubcategory returns a boolean if a field has been set.
func (o *ValidateResponseKycResultCompanyMasterData) HasCompanySubcategory() bool {
	if o != nil && !isNil(o.CompanySubcategory) {
		return true
	}

	return false
}

// SetCompanySubcategory gets a reference to the given string and assigns it to the CompanySubcategory field.
func (o *ValidateResponseKycResultCompanyMasterData) SetCompanySubcategory(v string) {
	o.CompanySubcategory = &v
}

// GetAuthorisedCapitalInInr returns the AuthorisedCapitalInInr field value if set, zero value otherwise.
func (o *ValidateResponseKycResultCompanyMasterData) GetAuthorisedCapitalInInr() string {
	if o == nil || isNil(o.AuthorisedCapitalInInr) {
		var ret string
		return ret
	}
	return *o.AuthorisedCapitalInInr
}

// GetAuthorisedCapitalInInrOk returns a tuple with the AuthorisedCapitalInInr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidateResponseKycResultCompanyMasterData) GetAuthorisedCapitalInInrOk() (*string, bool) {
	if o == nil || isNil(o.AuthorisedCapitalInInr) {
    return nil, false
	}
	return o.AuthorisedCapitalInInr, true
}

// HasAuthorisedCapitalInInr returns a boolean if a field has been set.
func (o *ValidateResponseKycResultCompanyMasterData) HasAuthorisedCapitalInInr() bool {
	if o != nil && !isNil(o.AuthorisedCapitalInInr) {
		return true
	}

	return false
}

// SetAuthorisedCapitalInInr gets a reference to the given string and assigns it to the AuthorisedCapitalInInr field.
func (o *ValidateResponseKycResultCompanyMasterData) SetAuthorisedCapitalInInr(v string) {
	o.AuthorisedCapitalInInr = &v
}

// GetCompanyStatusForEFiling returns the CompanyStatusForEFiling field value if set, zero value otherwise.
func (o *ValidateResponseKycResultCompanyMasterData) GetCompanyStatusForEFiling() string {
	if o == nil || isNil(o.CompanyStatusForEFiling) {
		var ret string
		return ret
	}
	return *o.CompanyStatusForEFiling
}

// GetCompanyStatusForEFilingOk returns a tuple with the CompanyStatusForEFiling field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidateResponseKycResultCompanyMasterData) GetCompanyStatusForEFilingOk() (*string, bool) {
	if o == nil || isNil(o.CompanyStatusForEFiling) {
    return nil, false
	}
	return o.CompanyStatusForEFiling, true
}

// HasCompanyStatusForEFiling returns a boolean if a field has been set.
func (o *ValidateResponseKycResultCompanyMasterData) HasCompanyStatusForEFiling() bool {
	if o != nil && !isNil(o.CompanyStatusForEFiling) {
		return true
	}

	return false
}

// SetCompanyStatusForEFiling gets a reference to the given string and assigns it to the CompanyStatusForEFiling field.
func (o *ValidateResponseKycResultCompanyMasterData) SetCompanyStatusForEFiling(v string) {
	o.CompanyStatusForEFiling = &v
}

// GetRocCode returns the RocCode field value if set, zero value otherwise.
func (o *ValidateResponseKycResultCompanyMasterData) GetRocCode() string {
	if o == nil || isNil(o.RocCode) {
		var ret string
		return ret
	}
	return *o.RocCode
}

// GetRocCodeOk returns a tuple with the RocCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidateResponseKycResultCompanyMasterData) GetRocCodeOk() (*string, bool) {
	if o == nil || isNil(o.RocCode) {
    return nil, false
	}
	return o.RocCode, true
}

// HasRocCode returns a boolean if a field has been set.
func (o *ValidateResponseKycResultCompanyMasterData) HasRocCode() bool {
	if o != nil && !isNil(o.RocCode) {
		return true
	}

	return false
}

// SetRocCode gets a reference to the given string and assigns it to the RocCode field.
func (o *ValidateResponseKycResultCompanyMasterData) SetRocCode(v string) {
	o.RocCode = &v
}

// GetDateOfBalanceSheet returns the DateOfBalanceSheet field value if set, zero value otherwise.
func (o *ValidateResponseKycResultCompanyMasterData) GetDateOfBalanceSheet() string {
	if o == nil || isNil(o.DateOfBalanceSheet) {
		var ret string
		return ret
	}
	return *o.DateOfBalanceSheet
}

// GetDateOfBalanceSheetOk returns a tuple with the DateOfBalanceSheet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidateResponseKycResultCompanyMasterData) GetDateOfBalanceSheetOk() (*string, bool) {
	if o == nil || isNil(o.DateOfBalanceSheet) {
    return nil, false
	}
	return o.DateOfBalanceSheet, true
}

// HasDateOfBalanceSheet returns a boolean if a field has been set.
func (o *ValidateResponseKycResultCompanyMasterData) HasDateOfBalanceSheet() bool {
	if o != nil && !isNil(o.DateOfBalanceSheet) {
		return true
	}

	return false
}

// SetDateOfBalanceSheet gets a reference to the given string and assigns it to the DateOfBalanceSheet field.
func (o *ValidateResponseKycResultCompanyMasterData) SetDateOfBalanceSheet(v string) {
	o.DateOfBalanceSheet = &v
}

// GetDateOfIncorporation returns the DateOfIncorporation field value if set, zero value otherwise.
func (o *ValidateResponseKycResultCompanyMasterData) GetDateOfIncorporation() string {
	if o == nil || isNil(o.DateOfIncorporation) {
		var ret string
		return ret
	}
	return *o.DateOfIncorporation
}

// GetDateOfIncorporationOk returns a tuple with the DateOfIncorporation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidateResponseKycResultCompanyMasterData) GetDateOfIncorporationOk() (*string, bool) {
	if o == nil || isNil(o.DateOfIncorporation) {
    return nil, false
	}
	return o.DateOfIncorporation, true
}

// HasDateOfIncorporation returns a boolean if a field has been set.
func (o *ValidateResponseKycResultCompanyMasterData) HasDateOfIncorporation() bool {
	if o != nil && !isNil(o.DateOfIncorporation) {
		return true
	}

	return false
}

// SetDateOfIncorporation gets a reference to the given string and assigns it to the DateOfIncorporation field.
func (o *ValidateResponseKycResultCompanyMasterData) SetDateOfIncorporation(v string) {
	o.DateOfIncorporation = &v
}

// GetCin returns the Cin field value if set, zero value otherwise.
func (o *ValidateResponseKycResultCompanyMasterData) GetCin() string {
	if o == nil || isNil(o.Cin) {
		var ret string
		return ret
	}
	return *o.Cin
}

// GetCinOk returns a tuple with the Cin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidateResponseKycResultCompanyMasterData) GetCinOk() (*string, bool) {
	if o == nil || isNil(o.Cin) {
    return nil, false
	}
	return o.Cin, true
}

// HasCin returns a boolean if a field has been set.
func (o *ValidateResponseKycResultCompanyMasterData) HasCin() bool {
	if o != nil && !isNil(o.Cin) {
		return true
	}

	return false
}

// SetCin gets a reference to the given string and assigns it to the Cin field.
func (o *ValidateResponseKycResultCompanyMasterData) SetCin(v string) {
	o.Cin = &v
}

// GetCompanyName returns the CompanyName field value if set, zero value otherwise.
func (o *ValidateResponseKycResultCompanyMasterData) GetCompanyName() string {
	if o == nil || isNil(o.CompanyName) {
		var ret string
		return ret
	}
	return *o.CompanyName
}

// GetCompanyNameOk returns a tuple with the CompanyName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidateResponseKycResultCompanyMasterData) GetCompanyNameOk() (*string, bool) {
	if o == nil || isNil(o.CompanyName) {
    return nil, false
	}
	return o.CompanyName, true
}

// HasCompanyName returns a boolean if a field has been set.
func (o *ValidateResponseKycResultCompanyMasterData) HasCompanyName() bool {
	if o != nil && !isNil(o.CompanyName) {
		return true
	}

	return false
}

// SetCompanyName gets a reference to the given string and assigns it to the CompanyName field.
func (o *ValidateResponseKycResultCompanyMasterData) SetCompanyName(v string) {
	o.CompanyName = &v
}

func (o ValidateResponseKycResultCompanyMasterData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.CompanyCategory) {
		toSerialize["companyCategory"] = o.CompanyCategory
	}
	if !isNil(o.EmailId) {
		toSerialize["emailId"] = o.EmailId
	}
	if !isNil(o.ClassOfCompany) {
		toSerialize["classOfCompany"] = o.ClassOfCompany
	}
	if !isNil(o.NumberOfMembersApplicableInCaseOfCompanyWithoutShareCapital) {
		toSerialize["numberOfMembersApplicableInCaseOfCompanyWithoutShareCapital"] = o.NumberOfMembersApplicableInCaseOfCompanyWithoutShareCapital
	}
	if !isNil(o.AddressOtherThanRegisteredOfficeWhereAllOrAnyBooksOfAccountAndPapersAreMaintained) {
		toSerialize["addressOtherThanRegisteredOfficeWhereAllOrAnyBooksOfAccountAndPapersAreMaintained"] = o.AddressOtherThanRegisteredOfficeWhereAllOrAnyBooksOfAccountAndPapersAreMaintained
	}
	if !isNil(o.DateOfLastAgm) {
		toSerialize["dateOfLastAgm"] = o.DateOfLastAgm
	}
	if !isNil(o.RegisteredAddress) {
		toSerialize["registeredAddress"] = o.RegisteredAddress
	}
	if !isNil(o.ActiveCompliance) {
		toSerialize["activeCompliance"] = o.ActiveCompliance
	}
	if !isNil(o.RegistrationNumber) {
		toSerialize["registrationNumber"] = o.RegistrationNumber
	}
	if !isNil(o.PaidUpCapitalInInr) {
		toSerialize["paidUpCapitalInInr"] = o.PaidUpCapitalInInr
	}
	if !isNil(o.WhetherListedOrNot) {
		toSerialize["whetherListedOrNot"] = o.WhetherListedOrNot
	}
	if !isNil(o.SuspendedAtStockExchange) {
		toSerialize["suspendedAtStockExchange"] = o.SuspendedAtStockExchange
	}
	if !isNil(o.CompanySubcategory) {
		toSerialize["companySubcategory"] = o.CompanySubcategory
	}
	if !isNil(o.AuthorisedCapitalInInr) {
		toSerialize["authorisedCapitalInInr"] = o.AuthorisedCapitalInInr
	}
	if !isNil(o.CompanyStatusForEFiling) {
		toSerialize["companyStatusForEFiling"] = o.CompanyStatusForEFiling
	}
	if !isNil(o.RocCode) {
		toSerialize["rocCode"] = o.RocCode
	}
	if !isNil(o.DateOfBalanceSheet) {
		toSerialize["dateOfBalanceSheet"] = o.DateOfBalanceSheet
	}
	if !isNil(o.DateOfIncorporation) {
		toSerialize["dateOfIncorporation"] = o.DateOfIncorporation
	}
	if !isNil(o.Cin) {
		toSerialize["cin "] = o.Cin
	}
	if !isNil(o.CompanyName) {
		toSerialize["companyName"] = o.CompanyName
	}
	return json.Marshal(toSerialize)
}

type NullableValidateResponseKycResultCompanyMasterData struct {
	value *ValidateResponseKycResultCompanyMasterData
	isSet bool
}

func (v NullableValidateResponseKycResultCompanyMasterData) Get() *ValidateResponseKycResultCompanyMasterData {
	return v.value
}

func (v *NullableValidateResponseKycResultCompanyMasterData) Set(val *ValidateResponseKycResultCompanyMasterData) {
	v.value = val
	v.isSet = true
}

func (v NullableValidateResponseKycResultCompanyMasterData) IsSet() bool {
	return v.isSet
}

func (v *NullableValidateResponseKycResultCompanyMasterData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableValidateResponseKycResultCompanyMasterData(val *ValidateResponseKycResultCompanyMasterData) *NullableValidateResponseKycResultCompanyMasterData {
	return &NullableValidateResponseKycResultCompanyMasterData{value: val, isSet: true}
}

func (v NullableValidateResponseKycResultCompanyMasterData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableValidateResponseKycResultCompanyMasterData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


