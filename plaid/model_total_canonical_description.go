/*
 * The Plaid API
 *
 * The Plaid REST API. Please see https://plaid.com/docs/api for more details.
 *
 * API version: 2020-09-14_1.78.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package plaid

import (
	"encoding/json"
	"fmt"
)

// TotalCanonicalDescription Commonly used term to describe the line item.
type TotalCanonicalDescription string

// List of TotalCanonicalDescription
const (
	TOTALCANONICALDESCRIPTION_BONUS TotalCanonicalDescription = "BONUS"
	TOTALCANONICALDESCRIPTION_COMMISSION TotalCanonicalDescription = "COMMISSION"
	TOTALCANONICALDESCRIPTION_OVERTIME TotalCanonicalDescription = "OVERTIME"
	TOTALCANONICALDESCRIPTION_PAID_TIME_OFF TotalCanonicalDescription = "PAID TIME OFF"
	TOTALCANONICALDESCRIPTION_REGULAR_PAY TotalCanonicalDescription = "REGULAR PAY"
	TOTALCANONICALDESCRIPTION_VACATION TotalCanonicalDescription = "VACATION"
	TOTALCANONICALDESCRIPTION_EMPLOYEE_MEDICARE TotalCanonicalDescription = "EMPLOYEE MEDICARE"
	TOTALCANONICALDESCRIPTION_FICA TotalCanonicalDescription = "FICA"
	TOTALCANONICALDESCRIPTION_SOCIAL_SECURITY_EMPLOYEE_TAX TotalCanonicalDescription = "SOCIAL SECURITY EMPLOYEE TAX"
	TOTALCANONICALDESCRIPTION_MEDICAL TotalCanonicalDescription = "MEDICAL"
	TOTALCANONICALDESCRIPTION_VISION TotalCanonicalDescription = "VISION"
	TOTALCANONICALDESCRIPTION_DENTAL TotalCanonicalDescription = "DENTAL"
	TOTALCANONICALDESCRIPTION_NET_PAY TotalCanonicalDescription = "NET PAY"
	TOTALCANONICALDESCRIPTION_TAXES TotalCanonicalDescription = "TAXES"
	TOTALCANONICALDESCRIPTION_NOT_FOUND TotalCanonicalDescription = "NOT_FOUND"
	TOTALCANONICALDESCRIPTION_OTHER TotalCanonicalDescription = "OTHER"
	TOTALCANONICALDESCRIPTION_NULL TotalCanonicalDescription = "null"
)

var allowedTotalCanonicalDescriptionEnumValues = []TotalCanonicalDescription{
	"BONUS",
	"COMMISSION",
	"OVERTIME",
	"PAID TIME OFF",
	"REGULAR PAY",
	"VACATION",
	"EMPLOYEE MEDICARE",
	"FICA",
	"SOCIAL SECURITY EMPLOYEE TAX",
	"MEDICAL",
	"VISION",
	"DENTAL",
	"NET PAY",
	"TAXES",
	"NOT_FOUND",
	"OTHER",
	"null",
}

func (v *TotalCanonicalDescription) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := TotalCanonicalDescription(value)
	for _, existing := range allowedTotalCanonicalDescriptionEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid TotalCanonicalDescription", value)
}

// NewTotalCanonicalDescriptionFromValue returns a pointer to a valid TotalCanonicalDescription
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewTotalCanonicalDescriptionFromValue(v string) (*TotalCanonicalDescription, error) {
	ev := TotalCanonicalDescription(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for TotalCanonicalDescription: valid values are %v", v, allowedTotalCanonicalDescriptionEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v TotalCanonicalDescription) IsValid() bool {
	for _, existing := range allowedTotalCanonicalDescriptionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TotalCanonicalDescription value
func (v TotalCanonicalDescription) Ptr() *TotalCanonicalDescription {
	return &v
}

type NullableTotalCanonicalDescription struct {
	value *TotalCanonicalDescription
	isSet bool
}

func (v NullableTotalCanonicalDescription) Get() *TotalCanonicalDescription {
	return v.value
}

func (v *NullableTotalCanonicalDescription) Set(val *TotalCanonicalDescription) {
	v.value = val
	v.isSet = true
}

func (v NullableTotalCanonicalDescription) IsSet() bool {
	return v.isSet
}

func (v *NullableTotalCanonicalDescription) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTotalCanonicalDescription(val *TotalCanonicalDescription) *NullableTotalCanonicalDescription {
	return &NullableTotalCanonicalDescription{value: val, isSet: true}
}

func (v NullableTotalCanonicalDescription) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTotalCanonicalDescription) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

