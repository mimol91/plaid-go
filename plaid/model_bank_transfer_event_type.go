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

// BankTransferEventType The type of event that this bank transfer represents.  `pending`: A new transfer was created; it is in the pending state.  `cancelled`: The transfer was cancelled by the client.  `failed`: The transfer failed, no funds were moved.  `posted`: The transfer has been successfully submitted to the payment network.  `reversed`: A posted transfer was reversed.
type BankTransferEventType string

// List of BankTransferEventType
const (
	BANKTRANSFEREVENTTYPE_PENDING BankTransferEventType = "pending"
	BANKTRANSFEREVENTTYPE_CANCELLED BankTransferEventType = "cancelled"
	BANKTRANSFEREVENTTYPE_FAILED BankTransferEventType = "failed"
	BANKTRANSFEREVENTTYPE_POSTED BankTransferEventType = "posted"
	BANKTRANSFEREVENTTYPE_REVERSED BankTransferEventType = "reversed"
)

var allowedBankTransferEventTypeEnumValues = []BankTransferEventType{
	"pending",
	"cancelled",
	"failed",
	"posted",
	"reversed",
}

func (v *BankTransferEventType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := BankTransferEventType(value)
	for _, existing := range allowedBankTransferEventTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid BankTransferEventType", value)
}

// NewBankTransferEventTypeFromValue returns a pointer to a valid BankTransferEventType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewBankTransferEventTypeFromValue(v string) (*BankTransferEventType, error) {
	ev := BankTransferEventType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for BankTransferEventType: valid values are %v", v, allowedBankTransferEventTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v BankTransferEventType) IsValid() bool {
	for _, existing := range allowedBankTransferEventTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to BankTransferEventType value
func (v BankTransferEventType) Ptr() *BankTransferEventType {
	return &v
}

type NullableBankTransferEventType struct {
	value *BankTransferEventType
	isSet bool
}

func (v NullableBankTransferEventType) Get() *BankTransferEventType {
	return v.value
}

func (v *NullableBankTransferEventType) Set(val *BankTransferEventType) {
	v.value = val
	v.isSet = true
}

func (v NullableBankTransferEventType) IsSet() bool {
	return v.isSet
}

func (v *NullableBankTransferEventType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBankTransferEventType(val *BankTransferEventType) *NullableBankTransferEventType {
	return &NullableBankTransferEventType{value: val, isSet: true}
}

func (v NullableBankTransferEventType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBankTransferEventType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

