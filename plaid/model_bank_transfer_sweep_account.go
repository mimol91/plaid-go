/*
 * The Plaid API
 *
 * The Plaid REST API. Please see https://plaid.com/docs/api for more details.
 *
 * API version: 2020-09-14_1.54.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package plaid

import (
	"encoding/json"
)

// BankTransferSweepAccount The account where the funds are swept to.
type BankTransferSweepAccount struct {
	AccountNumber string `json:"account_number"`
	RoutingNumber string `json:"routing_number"`
	AdditionalProperties map[string]interface{}
}

type _BankTransferSweepAccount BankTransferSweepAccount

// NewBankTransferSweepAccount instantiates a new BankTransferSweepAccount object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBankTransferSweepAccount(accountNumber string, routingNumber string) *BankTransferSweepAccount {
	this := BankTransferSweepAccount{}
	this.AccountNumber = accountNumber
	this.RoutingNumber = routingNumber
	return &this
}

// NewBankTransferSweepAccountWithDefaults instantiates a new BankTransferSweepAccount object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBankTransferSweepAccountWithDefaults() *BankTransferSweepAccount {
	this := BankTransferSweepAccount{}
	return &this
}

// GetAccountNumber returns the AccountNumber field value
func (o *BankTransferSweepAccount) GetAccountNumber() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccountNumber
}

// GetAccountNumberOk returns a tuple with the AccountNumber field value
// and a boolean to check if the value has been set.
func (o *BankTransferSweepAccount) GetAccountNumberOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AccountNumber, true
}

// SetAccountNumber sets field value
func (o *BankTransferSweepAccount) SetAccountNumber(v string) {
	o.AccountNumber = v
}

// GetRoutingNumber returns the RoutingNumber field value
func (o *BankTransferSweepAccount) GetRoutingNumber() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RoutingNumber
}

// GetRoutingNumberOk returns a tuple with the RoutingNumber field value
// and a boolean to check if the value has been set.
func (o *BankTransferSweepAccount) GetRoutingNumberOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RoutingNumber, true
}

// SetRoutingNumber sets field value
func (o *BankTransferSweepAccount) SetRoutingNumber(v string) {
	o.RoutingNumber = v
}

func (o BankTransferSweepAccount) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["account_number"] = o.AccountNumber
	}
	if true {
		toSerialize["routing_number"] = o.RoutingNumber
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *BankTransferSweepAccount) UnmarshalJSON(bytes []byte) (err error) {
	varBankTransferSweepAccount := _BankTransferSweepAccount{}

	if err = json.Unmarshal(bytes, &varBankTransferSweepAccount); err == nil {
		*o = BankTransferSweepAccount(varBankTransferSweepAccount)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "account_number")
		delete(additionalProperties, "routing_number")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableBankTransferSweepAccount struct {
	value *BankTransferSweepAccount
	isSet bool
}

func (v NullableBankTransferSweepAccount) Get() *BankTransferSweepAccount {
	return v.value
}

func (v *NullableBankTransferSweepAccount) Set(val *BankTransferSweepAccount) {
	v.value = val
	v.isSet = true
}

func (v NullableBankTransferSweepAccount) IsSet() bool {
	return v.isSet
}

func (v *NullableBankTransferSweepAccount) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBankTransferSweepAccount(val *BankTransferSweepAccount) *NullableBankTransferSweepAccount {
	return &NullableBankTransferSweepAccount{value: val, isSet: true}
}

func (v NullableBankTransferSweepAccount) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBankTransferSweepAccount) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


