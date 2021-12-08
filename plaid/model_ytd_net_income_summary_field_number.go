/*
 * The Plaid API
 *
 * The Plaid REST API. Please see https://plaid.com/docs/api for more details.
 *
 * API version: 2020-09-14_1.56.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package plaid

import (
	"encoding/json"
)

// YTDNetIncomeSummaryFieldNumber struct for YTDNetIncomeSummaryFieldNumber
type YTDNetIncomeSummaryFieldNumber struct {
	// The value of the field.
	Value float32 `json:"value"`
	VerificationStatus VerificationStatus `json:"verification_status"`
}

// NewYTDNetIncomeSummaryFieldNumber instantiates a new YTDNetIncomeSummaryFieldNumber object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewYTDNetIncomeSummaryFieldNumber(value float32, verificationStatus VerificationStatus) *YTDNetIncomeSummaryFieldNumber {
	this := YTDNetIncomeSummaryFieldNumber{}
	this.Value = value
	this.VerificationStatus = verificationStatus
	return &this
}

// NewYTDNetIncomeSummaryFieldNumberWithDefaults instantiates a new YTDNetIncomeSummaryFieldNumber object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewYTDNetIncomeSummaryFieldNumberWithDefaults() *YTDNetIncomeSummaryFieldNumber {
	this := YTDNetIncomeSummaryFieldNumber{}
	return &this
}

// GetValue returns the Value field value
func (o *YTDNetIncomeSummaryFieldNumber) GetValue() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *YTDNetIncomeSummaryFieldNumber) GetValueOk() (*float32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *YTDNetIncomeSummaryFieldNumber) SetValue(v float32) {
	o.Value = v
}

// GetVerificationStatus returns the VerificationStatus field value
func (o *YTDNetIncomeSummaryFieldNumber) GetVerificationStatus() VerificationStatus {
	if o == nil {
		var ret VerificationStatus
		return ret
	}

	return o.VerificationStatus
}

// GetVerificationStatusOk returns a tuple with the VerificationStatus field value
// and a boolean to check if the value has been set.
func (o *YTDNetIncomeSummaryFieldNumber) GetVerificationStatusOk() (*VerificationStatus, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.VerificationStatus, true
}

// SetVerificationStatus sets field value
func (o *YTDNetIncomeSummaryFieldNumber) SetVerificationStatus(v VerificationStatus) {
	o.VerificationStatus = v
}

func (o YTDNetIncomeSummaryFieldNumber) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["value"] = o.Value
	}
	if true {
		toSerialize["verification_status"] = o.VerificationStatus
	}
	return json.Marshal(toSerialize)
}

type NullableYTDNetIncomeSummaryFieldNumber struct {
	value *YTDNetIncomeSummaryFieldNumber
	isSet bool
}

func (v NullableYTDNetIncomeSummaryFieldNumber) Get() *YTDNetIncomeSummaryFieldNumber {
	return v.value
}

func (v *NullableYTDNetIncomeSummaryFieldNumber) Set(val *YTDNetIncomeSummaryFieldNumber) {
	v.value = val
	v.isSet = true
}

func (v NullableYTDNetIncomeSummaryFieldNumber) IsSet() bool {
	return v.isSet
}

func (v *NullableYTDNetIncomeSummaryFieldNumber) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableYTDNetIncomeSummaryFieldNumber(val *YTDNetIncomeSummaryFieldNumber) *NullableYTDNetIncomeSummaryFieldNumber {
	return &NullableYTDNetIncomeSummaryFieldNumber{value: val, isSet: true}
}

func (v NullableYTDNetIncomeSummaryFieldNumber) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableYTDNetIncomeSummaryFieldNumber) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


