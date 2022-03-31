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
	"fmt"
)

// RecurringTransactionFrequency Describes the frequency of the transaction stream.
type RecurringTransactionFrequency string

// List of RecurringTransactionFrequency
const (
	RECURRINGTRANSACTIONFREQUENCY_UNKNOWN RecurringTransactionFrequency = "UNKNOWN"
	RECURRINGTRANSACTIONFREQUENCY_WEEKLY RecurringTransactionFrequency = "WEEKLY"
	RECURRINGTRANSACTIONFREQUENCY_BIWEEKLY RecurringTransactionFrequency = "BIWEEKLY"
	RECURRINGTRANSACTIONFREQUENCY_SEMI_MONTHLY RecurringTransactionFrequency = "SEMI_MONTHLY"
	RECURRINGTRANSACTIONFREQUENCY_MONTHLY RecurringTransactionFrequency = "MONTHLY"
)

var allowedRecurringTransactionFrequencyEnumValues = []RecurringTransactionFrequency{
	"UNKNOWN",
	"WEEKLY",
	"BIWEEKLY",
	"SEMI_MONTHLY",
	"MONTHLY",
}

func (v *RecurringTransactionFrequency) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := RecurringTransactionFrequency(value)
	for _, existing := range allowedRecurringTransactionFrequencyEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid RecurringTransactionFrequency", value)
}

// NewRecurringTransactionFrequencyFromValue returns a pointer to a valid RecurringTransactionFrequency
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewRecurringTransactionFrequencyFromValue(v string) (*RecurringTransactionFrequency, error) {
	ev := RecurringTransactionFrequency(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for RecurringTransactionFrequency: valid values are %v", v, allowedRecurringTransactionFrequencyEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v RecurringTransactionFrequency) IsValid() bool {
	for _, existing := range allowedRecurringTransactionFrequencyEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to RecurringTransactionFrequency value
func (v RecurringTransactionFrequency) Ptr() *RecurringTransactionFrequency {
	return &v
}

type NullableRecurringTransactionFrequency struct {
	value *RecurringTransactionFrequency
	isSet bool
}

func (v NullableRecurringTransactionFrequency) Get() *RecurringTransactionFrequency {
	return v.value
}

func (v *NullableRecurringTransactionFrequency) Set(val *RecurringTransactionFrequency) {
	v.value = val
	v.isSet = true
}

func (v NullableRecurringTransactionFrequency) IsSet() bool {
	return v.isSet
}

func (v *NullableRecurringTransactionFrequency) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRecurringTransactionFrequency(val *RecurringTransactionFrequency) *NullableRecurringTransactionFrequency {
	return &NullableRecurringTransactionFrequency{value: val, isSet: true}
}

func (v NullableRecurringTransactionFrequency) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRecurringTransactionFrequency) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

