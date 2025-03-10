/*
 * The Plaid API
 *
 * The Plaid REST API. Please see https://plaid.com/docs/api for more details.
 *
 * API version: 2020-09-14_1.97.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package plaid

import (
	"encoding/json"
)

// PaystubOverride An object representing data from a paystub.
type PaystubOverride struct {
	Employer *PaystubOverrideEmployer `json:"employer,omitempty"`
	Employee *PaystubOverrideEmployee `json:"employee,omitempty"`
	IncomeBreakdown *[]IncomeBreakdown `json:"income_breakdown,omitempty"`
	PayPeriodDetails *PayPeriodDetails `json:"pay_period_details,omitempty"`
}

// NewPaystubOverride instantiates a new PaystubOverride object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaystubOverride() *PaystubOverride {
	this := PaystubOverride{}
	return &this
}

// NewPaystubOverrideWithDefaults instantiates a new PaystubOverride object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaystubOverrideWithDefaults() *PaystubOverride {
	this := PaystubOverride{}
	return &this
}

// GetEmployer returns the Employer field value if set, zero value otherwise.
func (o *PaystubOverride) GetEmployer() PaystubOverrideEmployer {
	if o == nil || o.Employer == nil {
		var ret PaystubOverrideEmployer
		return ret
	}
	return *o.Employer
}

// GetEmployerOk returns a tuple with the Employer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaystubOverride) GetEmployerOk() (*PaystubOverrideEmployer, bool) {
	if o == nil || o.Employer == nil {
		return nil, false
	}
	return o.Employer, true
}

// HasEmployer returns a boolean if a field has been set.
func (o *PaystubOverride) HasEmployer() bool {
	if o != nil && o.Employer != nil {
		return true
	}

	return false
}

// SetEmployer gets a reference to the given PaystubOverrideEmployer and assigns it to the Employer field.
func (o *PaystubOverride) SetEmployer(v PaystubOverrideEmployer) {
	o.Employer = &v
}

// GetEmployee returns the Employee field value if set, zero value otherwise.
func (o *PaystubOverride) GetEmployee() PaystubOverrideEmployee {
	if o == nil || o.Employee == nil {
		var ret PaystubOverrideEmployee
		return ret
	}
	return *o.Employee
}

// GetEmployeeOk returns a tuple with the Employee field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaystubOverride) GetEmployeeOk() (*PaystubOverrideEmployee, bool) {
	if o == nil || o.Employee == nil {
		return nil, false
	}
	return o.Employee, true
}

// HasEmployee returns a boolean if a field has been set.
func (o *PaystubOverride) HasEmployee() bool {
	if o != nil && o.Employee != nil {
		return true
	}

	return false
}

// SetEmployee gets a reference to the given PaystubOverrideEmployee and assigns it to the Employee field.
func (o *PaystubOverride) SetEmployee(v PaystubOverrideEmployee) {
	o.Employee = &v
}

// GetIncomeBreakdown returns the IncomeBreakdown field value if set, zero value otherwise.
func (o *PaystubOverride) GetIncomeBreakdown() []IncomeBreakdown {
	if o == nil || o.IncomeBreakdown == nil {
		var ret []IncomeBreakdown
		return ret
	}
	return *o.IncomeBreakdown
}

// GetIncomeBreakdownOk returns a tuple with the IncomeBreakdown field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaystubOverride) GetIncomeBreakdownOk() (*[]IncomeBreakdown, bool) {
	if o == nil || o.IncomeBreakdown == nil {
		return nil, false
	}
	return o.IncomeBreakdown, true
}

// HasIncomeBreakdown returns a boolean if a field has been set.
func (o *PaystubOverride) HasIncomeBreakdown() bool {
	if o != nil && o.IncomeBreakdown != nil {
		return true
	}

	return false
}

// SetIncomeBreakdown gets a reference to the given []IncomeBreakdown and assigns it to the IncomeBreakdown field.
func (o *PaystubOverride) SetIncomeBreakdown(v []IncomeBreakdown) {
	o.IncomeBreakdown = &v
}

// GetPayPeriodDetails returns the PayPeriodDetails field value if set, zero value otherwise.
func (o *PaystubOverride) GetPayPeriodDetails() PayPeriodDetails {
	if o == nil || o.PayPeriodDetails == nil {
		var ret PayPeriodDetails
		return ret
	}
	return *o.PayPeriodDetails
}

// GetPayPeriodDetailsOk returns a tuple with the PayPeriodDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaystubOverride) GetPayPeriodDetailsOk() (*PayPeriodDetails, bool) {
	if o == nil || o.PayPeriodDetails == nil {
		return nil, false
	}
	return o.PayPeriodDetails, true
}

// HasPayPeriodDetails returns a boolean if a field has been set.
func (o *PaystubOverride) HasPayPeriodDetails() bool {
	if o != nil && o.PayPeriodDetails != nil {
		return true
	}

	return false
}

// SetPayPeriodDetails gets a reference to the given PayPeriodDetails and assigns it to the PayPeriodDetails field.
func (o *PaystubOverride) SetPayPeriodDetails(v PayPeriodDetails) {
	o.PayPeriodDetails = &v
}

func (o PaystubOverride) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Employer != nil {
		toSerialize["employer"] = o.Employer
	}
	if o.Employee != nil {
		toSerialize["employee"] = o.Employee
	}
	if o.IncomeBreakdown != nil {
		toSerialize["income_breakdown"] = o.IncomeBreakdown
	}
	if o.PayPeriodDetails != nil {
		toSerialize["pay_period_details"] = o.PayPeriodDetails
	}
	return json.Marshal(toSerialize)
}

type NullablePaystubOverride struct {
	value *PaystubOverride
	isSet bool
}

func (v NullablePaystubOverride) Get() *PaystubOverride {
	return v.value
}

func (v *NullablePaystubOverride) Set(val *PaystubOverride) {
	v.value = val
	v.isSet = true
}

func (v NullablePaystubOverride) IsSet() bool {
	return v.isSet
}

func (v *NullablePaystubOverride) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaystubOverride(val *PaystubOverride) *NullablePaystubOverride {
	return &NullablePaystubOverride{value: val, isSet: true}
}

func (v NullablePaystubOverride) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaystubOverride) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


