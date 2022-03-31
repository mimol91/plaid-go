/*
 * The Plaid API
 *
 * The Plaid REST API. Please see https://plaid.com/docs/api for more details.
 *
 * API version: 2020-09-14_1.94.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package plaid

import (
	"encoding/json"
)

// CreditW2 W2 is an object that represents income data taken from a W2 tax document.
type CreditW2 struct {
	DocumentMetadata CreditDocumentMetadata `json:"document_metadata"`
	// An identifier of the document referenced by the document metadata.
	DocumentId string `json:"document_id"`
	Employer PayStubEmployer `json:"employer"`
	Employee PayStubEmployee `json:"employee"`
	// The tax year of the W2 document.
	TaxYear NullableString `json:"tax_year"`
	// An employee identification number or EIN.
	EmployerIdNumber NullableString `json:"employer_id_number"`
	// Wages from tips and other compensation.
	WagesTipsOtherComp NullableString `json:"wages_tips_other_comp"`
	// Federal income tax withheld for the tax year.
	FederalIncomeTaxWithheld NullableString `json:"federal_income_tax_withheld"`
	// Wages from social security.
	SocialSecurityWages NullableString `json:"social_security_wages"`
	// Social security tax withheld for the tax year.
	SocialSecurityTaxWithheld NullableString `json:"social_security_tax_withheld"`
	// Wages and tips from medicare.
	MedicareWagesAndTips NullableString `json:"medicare_wages_and_tips"`
	// Medicare tax withheld for the tax year.
	MedicareTaxWithheld NullableString `json:"medicare_tax_withheld"`
	// Tips from social security.
	SocialSecurityTips NullableString `json:"social_security_tips"`
	// Allocated tips.
	AllocatedTips NullableString `json:"allocated_tips"`
	// Contents from box 9 on the W2.
	Box9 NullableString `json:"box_9"`
	// Dependent care benefits.
	DependentCareBenefits NullableString `json:"dependent_care_benefits"`
	// Nonqualified plans.
	NonqualifiedPlans NullableString `json:"nonqualified_plans"`
	Box12 []W2Box12 `json:"box_12"`
	// Statutory employee.
	StatutoryEmployee NullableString `json:"statutory_employee"`
	// Retirement plan.
	RetirementPlan NullableString `json:"retirement_plan"`
	// Third party sick pay.
	ThirdPartySickPay NullableString `json:"third_party_sick_pay"`
	// Other.
	Other NullableString `json:"other"`
	StateAndLocalWages []W2StateAndLocalWages `json:"state_and_local_wages"`
	AdditionalProperties map[string]interface{}
}

type _CreditW2 CreditW2

// NewCreditW2 instantiates a new CreditW2 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreditW2(documentMetadata CreditDocumentMetadata, documentId string, employer PayStubEmployer, employee PayStubEmployee, taxYear NullableString, employerIdNumber NullableString, wagesTipsOtherComp NullableString, federalIncomeTaxWithheld NullableString, socialSecurityWages NullableString, socialSecurityTaxWithheld NullableString, medicareWagesAndTips NullableString, medicareTaxWithheld NullableString, socialSecurityTips NullableString, allocatedTips NullableString, box9 NullableString, dependentCareBenefits NullableString, nonqualifiedPlans NullableString, box12 []W2Box12, statutoryEmployee NullableString, retirementPlan NullableString, thirdPartySickPay NullableString, other NullableString, stateAndLocalWages []W2StateAndLocalWages) *CreditW2 {
	this := CreditW2{}
	this.DocumentMetadata = documentMetadata
	this.DocumentId = documentId
	this.Employer = employer
	this.Employee = employee
	this.TaxYear = taxYear
	this.EmployerIdNumber = employerIdNumber
	this.WagesTipsOtherComp = wagesTipsOtherComp
	this.FederalIncomeTaxWithheld = federalIncomeTaxWithheld
	this.SocialSecurityWages = socialSecurityWages
	this.SocialSecurityTaxWithheld = socialSecurityTaxWithheld
	this.MedicareWagesAndTips = medicareWagesAndTips
	this.MedicareTaxWithheld = medicareTaxWithheld
	this.SocialSecurityTips = socialSecurityTips
	this.AllocatedTips = allocatedTips
	this.Box9 = box9
	this.DependentCareBenefits = dependentCareBenefits
	this.NonqualifiedPlans = nonqualifiedPlans
	this.Box12 = box12
	this.StatutoryEmployee = statutoryEmployee
	this.RetirementPlan = retirementPlan
	this.ThirdPartySickPay = thirdPartySickPay
	this.Other = other
	this.StateAndLocalWages = stateAndLocalWages
	return &this
}

// NewCreditW2WithDefaults instantiates a new CreditW2 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreditW2WithDefaults() *CreditW2 {
	this := CreditW2{}
	return &this
}

// GetDocumentMetadata returns the DocumentMetadata field value
func (o *CreditW2) GetDocumentMetadata() CreditDocumentMetadata {
	if o == nil {
		var ret CreditDocumentMetadata
		return ret
	}

	return o.DocumentMetadata
}

// GetDocumentMetadataOk returns a tuple with the DocumentMetadata field value
// and a boolean to check if the value has been set.
func (o *CreditW2) GetDocumentMetadataOk() (*CreditDocumentMetadata, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.DocumentMetadata, true
}

// SetDocumentMetadata sets field value
func (o *CreditW2) SetDocumentMetadata(v CreditDocumentMetadata) {
	o.DocumentMetadata = v
}

// GetDocumentId returns the DocumentId field value
func (o *CreditW2) GetDocumentId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DocumentId
}

// GetDocumentIdOk returns a tuple with the DocumentId field value
// and a boolean to check if the value has been set.
func (o *CreditW2) GetDocumentIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.DocumentId, true
}

// SetDocumentId sets field value
func (o *CreditW2) SetDocumentId(v string) {
	o.DocumentId = v
}

// GetEmployer returns the Employer field value
func (o *CreditW2) GetEmployer() PayStubEmployer {
	if o == nil {
		var ret PayStubEmployer
		return ret
	}

	return o.Employer
}

// GetEmployerOk returns a tuple with the Employer field value
// and a boolean to check if the value has been set.
func (o *CreditW2) GetEmployerOk() (*PayStubEmployer, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Employer, true
}

// SetEmployer sets field value
func (o *CreditW2) SetEmployer(v PayStubEmployer) {
	o.Employer = v
}

// GetEmployee returns the Employee field value
func (o *CreditW2) GetEmployee() PayStubEmployee {
	if o == nil {
		var ret PayStubEmployee
		return ret
	}

	return o.Employee
}

// GetEmployeeOk returns a tuple with the Employee field value
// and a boolean to check if the value has been set.
func (o *CreditW2) GetEmployeeOk() (*PayStubEmployee, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Employee, true
}

// SetEmployee sets field value
func (o *CreditW2) SetEmployee(v PayStubEmployee) {
	o.Employee = v
}

// GetTaxYear returns the TaxYear field value
// If the value is explicit nil, the zero value for string will be returned
func (o *CreditW2) GetTaxYear() string {
	if o == nil || o.TaxYear.Get() == nil {
		var ret string
		return ret
	}

	return *o.TaxYear.Get()
}

// GetTaxYearOk returns a tuple with the TaxYear field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreditW2) GetTaxYearOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.TaxYear.Get(), o.TaxYear.IsSet()
}

// SetTaxYear sets field value
func (o *CreditW2) SetTaxYear(v string) {
	o.TaxYear.Set(&v)
}

// GetEmployerIdNumber returns the EmployerIdNumber field value
// If the value is explicit nil, the zero value for string will be returned
func (o *CreditW2) GetEmployerIdNumber() string {
	if o == nil || o.EmployerIdNumber.Get() == nil {
		var ret string
		return ret
	}

	return *o.EmployerIdNumber.Get()
}

// GetEmployerIdNumberOk returns a tuple with the EmployerIdNumber field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreditW2) GetEmployerIdNumberOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.EmployerIdNumber.Get(), o.EmployerIdNumber.IsSet()
}

// SetEmployerIdNumber sets field value
func (o *CreditW2) SetEmployerIdNumber(v string) {
	o.EmployerIdNumber.Set(&v)
}

// GetWagesTipsOtherComp returns the WagesTipsOtherComp field value
// If the value is explicit nil, the zero value for string will be returned
func (o *CreditW2) GetWagesTipsOtherComp() string {
	if o == nil || o.WagesTipsOtherComp.Get() == nil {
		var ret string
		return ret
	}

	return *o.WagesTipsOtherComp.Get()
}

// GetWagesTipsOtherCompOk returns a tuple with the WagesTipsOtherComp field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreditW2) GetWagesTipsOtherCompOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.WagesTipsOtherComp.Get(), o.WagesTipsOtherComp.IsSet()
}

// SetWagesTipsOtherComp sets field value
func (o *CreditW2) SetWagesTipsOtherComp(v string) {
	o.WagesTipsOtherComp.Set(&v)
}

// GetFederalIncomeTaxWithheld returns the FederalIncomeTaxWithheld field value
// If the value is explicit nil, the zero value for string will be returned
func (o *CreditW2) GetFederalIncomeTaxWithheld() string {
	if o == nil || o.FederalIncomeTaxWithheld.Get() == nil {
		var ret string
		return ret
	}

	return *o.FederalIncomeTaxWithheld.Get()
}

// GetFederalIncomeTaxWithheldOk returns a tuple with the FederalIncomeTaxWithheld field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreditW2) GetFederalIncomeTaxWithheldOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.FederalIncomeTaxWithheld.Get(), o.FederalIncomeTaxWithheld.IsSet()
}

// SetFederalIncomeTaxWithheld sets field value
func (o *CreditW2) SetFederalIncomeTaxWithheld(v string) {
	o.FederalIncomeTaxWithheld.Set(&v)
}

// GetSocialSecurityWages returns the SocialSecurityWages field value
// If the value is explicit nil, the zero value for string will be returned
func (o *CreditW2) GetSocialSecurityWages() string {
	if o == nil || o.SocialSecurityWages.Get() == nil {
		var ret string
		return ret
	}

	return *o.SocialSecurityWages.Get()
}

// GetSocialSecurityWagesOk returns a tuple with the SocialSecurityWages field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreditW2) GetSocialSecurityWagesOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.SocialSecurityWages.Get(), o.SocialSecurityWages.IsSet()
}

// SetSocialSecurityWages sets field value
func (o *CreditW2) SetSocialSecurityWages(v string) {
	o.SocialSecurityWages.Set(&v)
}

// GetSocialSecurityTaxWithheld returns the SocialSecurityTaxWithheld field value
// If the value is explicit nil, the zero value for string will be returned
func (o *CreditW2) GetSocialSecurityTaxWithheld() string {
	if o == nil || o.SocialSecurityTaxWithheld.Get() == nil {
		var ret string
		return ret
	}

	return *o.SocialSecurityTaxWithheld.Get()
}

// GetSocialSecurityTaxWithheldOk returns a tuple with the SocialSecurityTaxWithheld field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreditW2) GetSocialSecurityTaxWithheldOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.SocialSecurityTaxWithheld.Get(), o.SocialSecurityTaxWithheld.IsSet()
}

// SetSocialSecurityTaxWithheld sets field value
func (o *CreditW2) SetSocialSecurityTaxWithheld(v string) {
	o.SocialSecurityTaxWithheld.Set(&v)
}

// GetMedicareWagesAndTips returns the MedicareWagesAndTips field value
// If the value is explicit nil, the zero value for string will be returned
func (o *CreditW2) GetMedicareWagesAndTips() string {
	if o == nil || o.MedicareWagesAndTips.Get() == nil {
		var ret string
		return ret
	}

	return *o.MedicareWagesAndTips.Get()
}

// GetMedicareWagesAndTipsOk returns a tuple with the MedicareWagesAndTips field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreditW2) GetMedicareWagesAndTipsOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.MedicareWagesAndTips.Get(), o.MedicareWagesAndTips.IsSet()
}

// SetMedicareWagesAndTips sets field value
func (o *CreditW2) SetMedicareWagesAndTips(v string) {
	o.MedicareWagesAndTips.Set(&v)
}

// GetMedicareTaxWithheld returns the MedicareTaxWithheld field value
// If the value is explicit nil, the zero value for string will be returned
func (o *CreditW2) GetMedicareTaxWithheld() string {
	if o == nil || o.MedicareTaxWithheld.Get() == nil {
		var ret string
		return ret
	}

	return *o.MedicareTaxWithheld.Get()
}

// GetMedicareTaxWithheldOk returns a tuple with the MedicareTaxWithheld field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreditW2) GetMedicareTaxWithheldOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.MedicareTaxWithheld.Get(), o.MedicareTaxWithheld.IsSet()
}

// SetMedicareTaxWithheld sets field value
func (o *CreditW2) SetMedicareTaxWithheld(v string) {
	o.MedicareTaxWithheld.Set(&v)
}

// GetSocialSecurityTips returns the SocialSecurityTips field value
// If the value is explicit nil, the zero value for string will be returned
func (o *CreditW2) GetSocialSecurityTips() string {
	if o == nil || o.SocialSecurityTips.Get() == nil {
		var ret string
		return ret
	}

	return *o.SocialSecurityTips.Get()
}

// GetSocialSecurityTipsOk returns a tuple with the SocialSecurityTips field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreditW2) GetSocialSecurityTipsOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.SocialSecurityTips.Get(), o.SocialSecurityTips.IsSet()
}

// SetSocialSecurityTips sets field value
func (o *CreditW2) SetSocialSecurityTips(v string) {
	o.SocialSecurityTips.Set(&v)
}

// GetAllocatedTips returns the AllocatedTips field value
// If the value is explicit nil, the zero value for string will be returned
func (o *CreditW2) GetAllocatedTips() string {
	if o == nil || o.AllocatedTips.Get() == nil {
		var ret string
		return ret
	}

	return *o.AllocatedTips.Get()
}

// GetAllocatedTipsOk returns a tuple with the AllocatedTips field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreditW2) GetAllocatedTipsOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.AllocatedTips.Get(), o.AllocatedTips.IsSet()
}

// SetAllocatedTips sets field value
func (o *CreditW2) SetAllocatedTips(v string) {
	o.AllocatedTips.Set(&v)
}

// GetBox9 returns the Box9 field value
// If the value is explicit nil, the zero value for string will be returned
func (o *CreditW2) GetBox9() string {
	if o == nil || o.Box9.Get() == nil {
		var ret string
		return ret
	}

	return *o.Box9.Get()
}

// GetBox9Ok returns a tuple with the Box9 field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreditW2) GetBox9Ok() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Box9.Get(), o.Box9.IsSet()
}

// SetBox9 sets field value
func (o *CreditW2) SetBox9(v string) {
	o.Box9.Set(&v)
}

// GetDependentCareBenefits returns the DependentCareBenefits field value
// If the value is explicit nil, the zero value for string will be returned
func (o *CreditW2) GetDependentCareBenefits() string {
	if o == nil || o.DependentCareBenefits.Get() == nil {
		var ret string
		return ret
	}

	return *o.DependentCareBenefits.Get()
}

// GetDependentCareBenefitsOk returns a tuple with the DependentCareBenefits field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreditW2) GetDependentCareBenefitsOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.DependentCareBenefits.Get(), o.DependentCareBenefits.IsSet()
}

// SetDependentCareBenefits sets field value
func (o *CreditW2) SetDependentCareBenefits(v string) {
	o.DependentCareBenefits.Set(&v)
}

// GetNonqualifiedPlans returns the NonqualifiedPlans field value
// If the value is explicit nil, the zero value for string will be returned
func (o *CreditW2) GetNonqualifiedPlans() string {
	if o == nil || o.NonqualifiedPlans.Get() == nil {
		var ret string
		return ret
	}

	return *o.NonqualifiedPlans.Get()
}

// GetNonqualifiedPlansOk returns a tuple with the NonqualifiedPlans field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreditW2) GetNonqualifiedPlansOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.NonqualifiedPlans.Get(), o.NonqualifiedPlans.IsSet()
}

// SetNonqualifiedPlans sets field value
func (o *CreditW2) SetNonqualifiedPlans(v string) {
	o.NonqualifiedPlans.Set(&v)
}

// GetBox12 returns the Box12 field value
func (o *CreditW2) GetBox12() []W2Box12 {
	if o == nil {
		var ret []W2Box12
		return ret
	}

	return o.Box12
}

// GetBox12Ok returns a tuple with the Box12 field value
// and a boolean to check if the value has been set.
func (o *CreditW2) GetBox12Ok() (*[]W2Box12, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Box12, true
}

// SetBox12 sets field value
func (o *CreditW2) SetBox12(v []W2Box12) {
	o.Box12 = v
}

// GetStatutoryEmployee returns the StatutoryEmployee field value
// If the value is explicit nil, the zero value for string will be returned
func (o *CreditW2) GetStatutoryEmployee() string {
	if o == nil || o.StatutoryEmployee.Get() == nil {
		var ret string
		return ret
	}

	return *o.StatutoryEmployee.Get()
}

// GetStatutoryEmployeeOk returns a tuple with the StatutoryEmployee field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreditW2) GetStatutoryEmployeeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.StatutoryEmployee.Get(), o.StatutoryEmployee.IsSet()
}

// SetStatutoryEmployee sets field value
func (o *CreditW2) SetStatutoryEmployee(v string) {
	o.StatutoryEmployee.Set(&v)
}

// GetRetirementPlan returns the RetirementPlan field value
// If the value is explicit nil, the zero value for string will be returned
func (o *CreditW2) GetRetirementPlan() string {
	if o == nil || o.RetirementPlan.Get() == nil {
		var ret string
		return ret
	}

	return *o.RetirementPlan.Get()
}

// GetRetirementPlanOk returns a tuple with the RetirementPlan field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreditW2) GetRetirementPlanOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.RetirementPlan.Get(), o.RetirementPlan.IsSet()
}

// SetRetirementPlan sets field value
func (o *CreditW2) SetRetirementPlan(v string) {
	o.RetirementPlan.Set(&v)
}

// GetThirdPartySickPay returns the ThirdPartySickPay field value
// If the value is explicit nil, the zero value for string will be returned
func (o *CreditW2) GetThirdPartySickPay() string {
	if o == nil || o.ThirdPartySickPay.Get() == nil {
		var ret string
		return ret
	}

	return *o.ThirdPartySickPay.Get()
}

// GetThirdPartySickPayOk returns a tuple with the ThirdPartySickPay field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreditW2) GetThirdPartySickPayOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ThirdPartySickPay.Get(), o.ThirdPartySickPay.IsSet()
}

// SetThirdPartySickPay sets field value
func (o *CreditW2) SetThirdPartySickPay(v string) {
	o.ThirdPartySickPay.Set(&v)
}

// GetOther returns the Other field value
// If the value is explicit nil, the zero value for string will be returned
func (o *CreditW2) GetOther() string {
	if o == nil || o.Other.Get() == nil {
		var ret string
		return ret
	}

	return *o.Other.Get()
}

// GetOtherOk returns a tuple with the Other field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreditW2) GetOtherOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Other.Get(), o.Other.IsSet()
}

// SetOther sets field value
func (o *CreditW2) SetOther(v string) {
	o.Other.Set(&v)
}

// GetStateAndLocalWages returns the StateAndLocalWages field value
func (o *CreditW2) GetStateAndLocalWages() []W2StateAndLocalWages {
	if o == nil {
		var ret []W2StateAndLocalWages
		return ret
	}

	return o.StateAndLocalWages
}

// GetStateAndLocalWagesOk returns a tuple with the StateAndLocalWages field value
// and a boolean to check if the value has been set.
func (o *CreditW2) GetStateAndLocalWagesOk() (*[]W2StateAndLocalWages, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.StateAndLocalWages, true
}

// SetStateAndLocalWages sets field value
func (o *CreditW2) SetStateAndLocalWages(v []W2StateAndLocalWages) {
	o.StateAndLocalWages = v
}

func (o CreditW2) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["document_metadata"] = o.DocumentMetadata
	}
	if true {
		toSerialize["document_id"] = o.DocumentId
	}
	if true {
		toSerialize["employer"] = o.Employer
	}
	if true {
		toSerialize["employee"] = o.Employee
	}
	if true {
		toSerialize["tax_year"] = o.TaxYear.Get()
	}
	if true {
		toSerialize["employer_id_number"] = o.EmployerIdNumber.Get()
	}
	if true {
		toSerialize["wages_tips_other_comp"] = o.WagesTipsOtherComp.Get()
	}
	if true {
		toSerialize["federal_income_tax_withheld"] = o.FederalIncomeTaxWithheld.Get()
	}
	if true {
		toSerialize["social_security_wages"] = o.SocialSecurityWages.Get()
	}
	if true {
		toSerialize["social_security_tax_withheld"] = o.SocialSecurityTaxWithheld.Get()
	}
	if true {
		toSerialize["medicare_wages_and_tips"] = o.MedicareWagesAndTips.Get()
	}
	if true {
		toSerialize["medicare_tax_withheld"] = o.MedicareTaxWithheld.Get()
	}
	if true {
		toSerialize["social_security_tips"] = o.SocialSecurityTips.Get()
	}
	if true {
		toSerialize["allocated_tips"] = o.AllocatedTips.Get()
	}
	if true {
		toSerialize["box_9"] = o.Box9.Get()
	}
	if true {
		toSerialize["dependent_care_benefits"] = o.DependentCareBenefits.Get()
	}
	if true {
		toSerialize["nonqualified_plans"] = o.NonqualifiedPlans.Get()
	}
	if true {
		toSerialize["box_12"] = o.Box12
	}
	if true {
		toSerialize["statutory_employee"] = o.StatutoryEmployee.Get()
	}
	if true {
		toSerialize["retirement_plan"] = o.RetirementPlan.Get()
	}
	if true {
		toSerialize["third_party_sick_pay"] = o.ThirdPartySickPay.Get()
	}
	if true {
		toSerialize["other"] = o.Other.Get()
	}
	if true {
		toSerialize["state_and_local_wages"] = o.StateAndLocalWages
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *CreditW2) UnmarshalJSON(bytes []byte) (err error) {
	varCreditW2 := _CreditW2{}

	if err = json.Unmarshal(bytes, &varCreditW2); err == nil {
		*o = CreditW2(varCreditW2)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "document_metadata")
		delete(additionalProperties, "document_id")
		delete(additionalProperties, "employer")
		delete(additionalProperties, "employee")
		delete(additionalProperties, "tax_year")
		delete(additionalProperties, "employer_id_number")
		delete(additionalProperties, "wages_tips_other_comp")
		delete(additionalProperties, "federal_income_tax_withheld")
		delete(additionalProperties, "social_security_wages")
		delete(additionalProperties, "social_security_tax_withheld")
		delete(additionalProperties, "medicare_wages_and_tips")
		delete(additionalProperties, "medicare_tax_withheld")
		delete(additionalProperties, "social_security_tips")
		delete(additionalProperties, "allocated_tips")
		delete(additionalProperties, "box_9")
		delete(additionalProperties, "dependent_care_benefits")
		delete(additionalProperties, "nonqualified_plans")
		delete(additionalProperties, "box_12")
		delete(additionalProperties, "statutory_employee")
		delete(additionalProperties, "retirement_plan")
		delete(additionalProperties, "third_party_sick_pay")
		delete(additionalProperties, "other")
		delete(additionalProperties, "state_and_local_wages")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCreditW2 struct {
	value *CreditW2
	isSet bool
}

func (v NullableCreditW2) Get() *CreditW2 {
	return v.value
}

func (v *NullableCreditW2) Set(val *CreditW2) {
	v.value = val
	v.isSet = true
}

func (v NullableCreditW2) IsSet() bool {
	return v.isSet
}

func (v *NullableCreditW2) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreditW2(val *CreditW2) *NullableCreditW2 {
	return &NullableCreditW2{value: val, isSet: true}
}

func (v NullableCreditW2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreditW2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


