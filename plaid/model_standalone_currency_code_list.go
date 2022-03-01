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
)

// StandaloneCurrencyCodeList The following currency codes are supported by Plaid.
type StandaloneCurrencyCodeList struct {
	// Plaid supports all ISO 4217 currency codes.
	IsoCurrencyCode string `json:"iso_currency_code"`
	// List of unofficial currency codes
	UnofficialCurrencyCode string `json:"unofficial_currency_code"`
	AdditionalProperties map[string]interface{}
}

type _StandaloneCurrencyCodeList StandaloneCurrencyCodeList

// NewStandaloneCurrencyCodeList instantiates a new StandaloneCurrencyCodeList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStandaloneCurrencyCodeList(isoCurrencyCode string, unofficialCurrencyCode string) *StandaloneCurrencyCodeList {
	this := StandaloneCurrencyCodeList{}
	this.IsoCurrencyCode = isoCurrencyCode
	this.UnofficialCurrencyCode = unofficialCurrencyCode
	return &this
}

// NewStandaloneCurrencyCodeListWithDefaults instantiates a new StandaloneCurrencyCodeList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStandaloneCurrencyCodeListWithDefaults() *StandaloneCurrencyCodeList {
	this := StandaloneCurrencyCodeList{}
	return &this
}

// GetIsoCurrencyCode returns the IsoCurrencyCode field value
func (o *StandaloneCurrencyCodeList) GetIsoCurrencyCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.IsoCurrencyCode
}

// GetIsoCurrencyCodeOk returns a tuple with the IsoCurrencyCode field value
// and a boolean to check if the value has been set.
func (o *StandaloneCurrencyCodeList) GetIsoCurrencyCodeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.IsoCurrencyCode, true
}

// SetIsoCurrencyCode sets field value
func (o *StandaloneCurrencyCodeList) SetIsoCurrencyCode(v string) {
	o.IsoCurrencyCode = v
}

// GetUnofficialCurrencyCode returns the UnofficialCurrencyCode field value
func (o *StandaloneCurrencyCodeList) GetUnofficialCurrencyCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UnofficialCurrencyCode
}

// GetUnofficialCurrencyCodeOk returns a tuple with the UnofficialCurrencyCode field value
// and a boolean to check if the value has been set.
func (o *StandaloneCurrencyCodeList) GetUnofficialCurrencyCodeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.UnofficialCurrencyCode, true
}

// SetUnofficialCurrencyCode sets field value
func (o *StandaloneCurrencyCodeList) SetUnofficialCurrencyCode(v string) {
	o.UnofficialCurrencyCode = v
}

func (o StandaloneCurrencyCodeList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["iso_currency_code"] = o.IsoCurrencyCode
	}
	if true {
		toSerialize["unofficial_currency_code"] = o.UnofficialCurrencyCode
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *StandaloneCurrencyCodeList) UnmarshalJSON(bytes []byte) (err error) {
	varStandaloneCurrencyCodeList := _StandaloneCurrencyCodeList{}

	if err = json.Unmarshal(bytes, &varStandaloneCurrencyCodeList); err == nil {
		*o = StandaloneCurrencyCodeList(varStandaloneCurrencyCodeList)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "iso_currency_code")
		delete(additionalProperties, "unofficial_currency_code")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableStandaloneCurrencyCodeList struct {
	value *StandaloneCurrencyCodeList
	isSet bool
}

func (v NullableStandaloneCurrencyCodeList) Get() *StandaloneCurrencyCodeList {
	return v.value
}

func (v *NullableStandaloneCurrencyCodeList) Set(val *StandaloneCurrencyCodeList) {
	v.value = val
	v.isSet = true
}

func (v NullableStandaloneCurrencyCodeList) IsSet() bool {
	return v.isSet
}

func (v *NullableStandaloneCurrencyCodeList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStandaloneCurrencyCodeList(val *StandaloneCurrencyCodeList) *NullableStandaloneCurrencyCodeList {
	return &NullableStandaloneCurrencyCodeList{value: val, isSet: true}
}

func (v NullableStandaloneCurrencyCodeList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStandaloneCurrencyCodeList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


