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

// PayStub An object representing an end user's pay stub.
type PayStub struct {
	Deductions PayStubDeductions `json:"deductions"`
	// An identifier of the document referenced by the document metadata.
	DocumentId NullableString `json:"document_id"`
	DocumentMetadata CreditDocumentMetadata `json:"document_metadata"`
	Earnings PayStubEarnings `json:"earnings"`
	Employee PayStubEmployee `json:"employee"`
	Employer PayStubEmployer `json:"employer"`
	NetPay PayStubNetPay `json:"net_pay"`
	PayPeriodDetails PayStubPayPeriodDetails `json:"pay_period_details"`
	Verification NullablePayStubVerification `json:"verification"`
	AdditionalProperties map[string]interface{}
}

type _PayStub PayStub

// NewPayStub instantiates a new PayStub object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPayStub(deductions PayStubDeductions, documentId NullableString, documentMetadata CreditDocumentMetadata, earnings PayStubEarnings, employee PayStubEmployee, employer PayStubEmployer, netPay PayStubNetPay, payPeriodDetails PayStubPayPeriodDetails, verification NullablePayStubVerification) *PayStub {
	this := PayStub{}
	this.Deductions = deductions
	this.DocumentId = documentId
	this.DocumentMetadata = documentMetadata
	this.Earnings = earnings
	this.Employee = employee
	this.Employer = employer
	this.NetPay = netPay
	this.PayPeriodDetails = payPeriodDetails
	this.Verification = verification
	return &this
}

// NewPayStubWithDefaults instantiates a new PayStub object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPayStubWithDefaults() *PayStub {
	this := PayStub{}
	return &this
}

// GetDeductions returns the Deductions field value
func (o *PayStub) GetDeductions() PayStubDeductions {
	if o == nil {
		var ret PayStubDeductions
		return ret
	}

	return o.Deductions
}

// GetDeductionsOk returns a tuple with the Deductions field value
// and a boolean to check if the value has been set.
func (o *PayStub) GetDeductionsOk() (*PayStubDeductions, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Deductions, true
}

// SetDeductions sets field value
func (o *PayStub) SetDeductions(v PayStubDeductions) {
	o.Deductions = v
}

// GetDocumentId returns the DocumentId field value
// If the value is explicit nil, the zero value for string will be returned
func (o *PayStub) GetDocumentId() string {
	if o == nil || o.DocumentId.Get() == nil {
		var ret string
		return ret
	}

	return *o.DocumentId.Get()
}

// GetDocumentIdOk returns a tuple with the DocumentId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PayStub) GetDocumentIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.DocumentId.Get(), o.DocumentId.IsSet()
}

// SetDocumentId sets field value
func (o *PayStub) SetDocumentId(v string) {
	o.DocumentId.Set(&v)
}

// GetDocumentMetadata returns the DocumentMetadata field value
func (o *PayStub) GetDocumentMetadata() CreditDocumentMetadata {
	if o == nil {
		var ret CreditDocumentMetadata
		return ret
	}

	return o.DocumentMetadata
}

// GetDocumentMetadataOk returns a tuple with the DocumentMetadata field value
// and a boolean to check if the value has been set.
func (o *PayStub) GetDocumentMetadataOk() (*CreditDocumentMetadata, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.DocumentMetadata, true
}

// SetDocumentMetadata sets field value
func (o *PayStub) SetDocumentMetadata(v CreditDocumentMetadata) {
	o.DocumentMetadata = v
}

// GetEarnings returns the Earnings field value
func (o *PayStub) GetEarnings() PayStubEarnings {
	if o == nil {
		var ret PayStubEarnings
		return ret
	}

	return o.Earnings
}

// GetEarningsOk returns a tuple with the Earnings field value
// and a boolean to check if the value has been set.
func (o *PayStub) GetEarningsOk() (*PayStubEarnings, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Earnings, true
}

// SetEarnings sets field value
func (o *PayStub) SetEarnings(v PayStubEarnings) {
	o.Earnings = v
}

// GetEmployee returns the Employee field value
func (o *PayStub) GetEmployee() PayStubEmployee {
	if o == nil {
		var ret PayStubEmployee
		return ret
	}

	return o.Employee
}

// GetEmployeeOk returns a tuple with the Employee field value
// and a boolean to check if the value has been set.
func (o *PayStub) GetEmployeeOk() (*PayStubEmployee, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Employee, true
}

// SetEmployee sets field value
func (o *PayStub) SetEmployee(v PayStubEmployee) {
	o.Employee = v
}

// GetEmployer returns the Employer field value
func (o *PayStub) GetEmployer() PayStubEmployer {
	if o == nil {
		var ret PayStubEmployer
		return ret
	}

	return o.Employer
}

// GetEmployerOk returns a tuple with the Employer field value
// and a boolean to check if the value has been set.
func (o *PayStub) GetEmployerOk() (*PayStubEmployer, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Employer, true
}

// SetEmployer sets field value
func (o *PayStub) SetEmployer(v PayStubEmployer) {
	o.Employer = v
}

// GetNetPay returns the NetPay field value
func (o *PayStub) GetNetPay() PayStubNetPay {
	if o == nil {
		var ret PayStubNetPay
		return ret
	}

	return o.NetPay
}

// GetNetPayOk returns a tuple with the NetPay field value
// and a boolean to check if the value has been set.
func (o *PayStub) GetNetPayOk() (*PayStubNetPay, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.NetPay, true
}

// SetNetPay sets field value
func (o *PayStub) SetNetPay(v PayStubNetPay) {
	o.NetPay = v
}

// GetPayPeriodDetails returns the PayPeriodDetails field value
func (o *PayStub) GetPayPeriodDetails() PayStubPayPeriodDetails {
	if o == nil {
		var ret PayStubPayPeriodDetails
		return ret
	}

	return o.PayPeriodDetails
}

// GetPayPeriodDetailsOk returns a tuple with the PayPeriodDetails field value
// and a boolean to check if the value has been set.
func (o *PayStub) GetPayPeriodDetailsOk() (*PayStubPayPeriodDetails, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.PayPeriodDetails, true
}

// SetPayPeriodDetails sets field value
func (o *PayStub) SetPayPeriodDetails(v PayStubPayPeriodDetails) {
	o.PayPeriodDetails = v
}

// GetVerification returns the Verification field value
// If the value is explicit nil, the zero value for PayStubVerification will be returned
func (o *PayStub) GetVerification() PayStubVerification {
	if o == nil || o.Verification.Get() == nil {
		var ret PayStubVerification
		return ret
	}

	return *o.Verification.Get()
}

// GetVerificationOk returns a tuple with the Verification field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PayStub) GetVerificationOk() (*PayStubVerification, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Verification.Get(), o.Verification.IsSet()
}

// SetVerification sets field value
func (o *PayStub) SetVerification(v PayStubVerification) {
	o.Verification.Set(&v)
}

func (o PayStub) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["deductions"] = o.Deductions
	}
	if true {
		toSerialize["document_id"] = o.DocumentId.Get()
	}
	if true {
		toSerialize["document_metadata"] = o.DocumentMetadata
	}
	if true {
		toSerialize["earnings"] = o.Earnings
	}
	if true {
		toSerialize["employee"] = o.Employee
	}
	if true {
		toSerialize["employer"] = o.Employer
	}
	if true {
		toSerialize["net_pay"] = o.NetPay
	}
	if true {
		toSerialize["pay_period_details"] = o.PayPeriodDetails
	}
	if true {
		toSerialize["verification"] = o.Verification.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *PayStub) UnmarshalJSON(bytes []byte) (err error) {
	varPayStub := _PayStub{}

	if err = json.Unmarshal(bytes, &varPayStub); err == nil {
		*o = PayStub(varPayStub)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "deductions")
		delete(additionalProperties, "document_id")
		delete(additionalProperties, "document_metadata")
		delete(additionalProperties, "earnings")
		delete(additionalProperties, "employee")
		delete(additionalProperties, "employer")
		delete(additionalProperties, "net_pay")
		delete(additionalProperties, "pay_period_details")
		delete(additionalProperties, "verification")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePayStub struct {
	value *PayStub
	isSet bool
}

func (v NullablePayStub) Get() *PayStub {
	return v.value
}

func (v *NullablePayStub) Set(val *PayStub) {
	v.value = val
	v.isSet = true
}

func (v NullablePayStub) IsSet() bool {
	return v.isSet
}

func (v *NullablePayStub) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePayStub(val *PayStub) *NullablePayStub {
	return &NullablePayStub{value: val, isSet: true}
}

func (v NullablePayStub) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePayStub) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


