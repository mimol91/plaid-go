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

// PaymentChannel The channel used to make a payment. `online:` transactions that took place online.  `in store:` transactions that were made at a physical location.  `other:` transactions that relate to banks, e.g. fees or deposits.
type PaymentChannel string

// List of PaymentChannel
const (
	PAYMENTCHANNEL_ONLINE PaymentChannel = "online"
	PAYMENTCHANNEL_IN_STORE PaymentChannel = "in store"
	PAYMENTCHANNEL_OTHER PaymentChannel = "other"
)

var allowedPaymentChannelEnumValues = []PaymentChannel{
	"online",
	"in store",
	"other",
}

func (v *PaymentChannel) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PaymentChannel(value)
	for _, existing := range allowedPaymentChannelEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PaymentChannel", value)
}

// NewPaymentChannelFromValue returns a pointer to a valid PaymentChannel
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPaymentChannelFromValue(v string) (*PaymentChannel, error) {
	ev := PaymentChannel(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PaymentChannel: valid values are %v", v, allowedPaymentChannelEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PaymentChannel) IsValid() bool {
	for _, existing := range allowedPaymentChannelEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PaymentChannel value
func (v PaymentChannel) Ptr() *PaymentChannel {
	return &v
}

type NullablePaymentChannel struct {
	value *PaymentChannel
	isSet bool
}

func (v NullablePaymentChannel) Get() *PaymentChannel {
	return v.value
}

func (v *NullablePaymentChannel) Set(val *PaymentChannel) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentChannel) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentChannel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentChannel(val *PaymentChannel) *NullablePaymentChannel {
	return &NullablePaymentChannel{value: val, isSet: true}
}

func (v NullablePaymentChannel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentChannel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

