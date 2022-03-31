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

// VerificationStatus The verification status. One of the following:  `\"VERIFIED\"`: The information was successfully verified.  `\"UNVERIFIED\"`: The verification has not yet been performed.  `\"NEEDS_INFO\"`: The verification was attempted but could not be completed due to missing information.  \"`UNABLE_TO_VERIFY`\": The verification was performed and the information could not be verified.  `\"UNKNOWN\"`: The verification status is unknown.
type VerificationStatus string

// List of VerificationStatus
const (
	VERIFICATIONSTATUS_VERIFIED VerificationStatus = "VERIFIED"
	VERIFICATIONSTATUS_UNVERIFIED VerificationStatus = "UNVERIFIED"
	VERIFICATIONSTATUS_NEEDS_INFO VerificationStatus = "NEEDS_INFO"
	VERIFICATIONSTATUS_UNABLE_TO_VERIFY VerificationStatus = "UNABLE_TO_VERIFY"
	VERIFICATIONSTATUS_UNKNOWN VerificationStatus = "UNKNOWN"
)

var allowedVerificationStatusEnumValues = []VerificationStatus{
	"VERIFIED",
	"UNVERIFIED",
	"NEEDS_INFO",
	"UNABLE_TO_VERIFY",
	"UNKNOWN",
}

func (v *VerificationStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := VerificationStatus(value)
	for _, existing := range allowedVerificationStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid VerificationStatus", value)
}

// NewVerificationStatusFromValue returns a pointer to a valid VerificationStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewVerificationStatusFromValue(v string) (*VerificationStatus, error) {
	ev := VerificationStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for VerificationStatus: valid values are %v", v, allowedVerificationStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v VerificationStatus) IsValid() bool {
	for _, existing := range allowedVerificationStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to VerificationStatus value
func (v VerificationStatus) Ptr() *VerificationStatus {
	return &v
}

type NullableVerificationStatus struct {
	value *VerificationStatus
	isSet bool
}

func (v NullableVerificationStatus) Get() *VerificationStatus {
	return v.value
}

func (v *NullableVerificationStatus) Set(val *VerificationStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableVerificationStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableVerificationStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVerificationStatus(val *VerificationStatus) *NullableVerificationStatus {
	return &NullableVerificationStatus{value: val, isSet: true}
}

func (v NullableVerificationStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVerificationStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

