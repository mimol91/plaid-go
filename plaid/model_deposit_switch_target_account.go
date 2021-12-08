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

// DepositSwitchTargetAccount struct for DepositSwitchTargetAccount
type DepositSwitchTargetAccount struct {
	// Account number for deposit switch destination
	AccountNumber string `json:"account_number"`
	// Routing number for deposit switch destination
	RoutingNumber string `json:"routing_number"`
	// The name of the deposit switch destination account, as it will be displayed to the end user in the Deposit Switch interface. It is not required to match the name used in online banking.
	AccountName string `json:"account_name"`
	// The account subtype of the account, either `checking` or `savings`.
	AccountSubtype string `json:"account_subtype"`
	AdditionalProperties map[string]interface{}
}

type _DepositSwitchTargetAccount DepositSwitchTargetAccount

// NewDepositSwitchTargetAccount instantiates a new DepositSwitchTargetAccount object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDepositSwitchTargetAccount(accountNumber string, routingNumber string, accountName string, accountSubtype string) *DepositSwitchTargetAccount {
	this := DepositSwitchTargetAccount{}
	this.AccountNumber = accountNumber
	this.RoutingNumber = routingNumber
	this.AccountName = accountName
	this.AccountSubtype = accountSubtype
	return &this
}

// NewDepositSwitchTargetAccountWithDefaults instantiates a new DepositSwitchTargetAccount object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDepositSwitchTargetAccountWithDefaults() *DepositSwitchTargetAccount {
	this := DepositSwitchTargetAccount{}
	return &this
}

// GetAccountNumber returns the AccountNumber field value
func (o *DepositSwitchTargetAccount) GetAccountNumber() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccountNumber
}

// GetAccountNumberOk returns a tuple with the AccountNumber field value
// and a boolean to check if the value has been set.
func (o *DepositSwitchTargetAccount) GetAccountNumberOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AccountNumber, true
}

// SetAccountNumber sets field value
func (o *DepositSwitchTargetAccount) SetAccountNumber(v string) {
	o.AccountNumber = v
}

// GetRoutingNumber returns the RoutingNumber field value
func (o *DepositSwitchTargetAccount) GetRoutingNumber() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RoutingNumber
}

// GetRoutingNumberOk returns a tuple with the RoutingNumber field value
// and a boolean to check if the value has been set.
func (o *DepositSwitchTargetAccount) GetRoutingNumberOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RoutingNumber, true
}

// SetRoutingNumber sets field value
func (o *DepositSwitchTargetAccount) SetRoutingNumber(v string) {
	o.RoutingNumber = v
}

// GetAccountName returns the AccountName field value
func (o *DepositSwitchTargetAccount) GetAccountName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccountName
}

// GetAccountNameOk returns a tuple with the AccountName field value
// and a boolean to check if the value has been set.
func (o *DepositSwitchTargetAccount) GetAccountNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AccountName, true
}

// SetAccountName sets field value
func (o *DepositSwitchTargetAccount) SetAccountName(v string) {
	o.AccountName = v
}

// GetAccountSubtype returns the AccountSubtype field value
func (o *DepositSwitchTargetAccount) GetAccountSubtype() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccountSubtype
}

// GetAccountSubtypeOk returns a tuple with the AccountSubtype field value
// and a boolean to check if the value has been set.
func (o *DepositSwitchTargetAccount) GetAccountSubtypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AccountSubtype, true
}

// SetAccountSubtype sets field value
func (o *DepositSwitchTargetAccount) SetAccountSubtype(v string) {
	o.AccountSubtype = v
}

func (o DepositSwitchTargetAccount) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["account_number"] = o.AccountNumber
	}
	if true {
		toSerialize["routing_number"] = o.RoutingNumber
	}
	if true {
		toSerialize["account_name"] = o.AccountName
	}
	if true {
		toSerialize["account_subtype"] = o.AccountSubtype
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *DepositSwitchTargetAccount) UnmarshalJSON(bytes []byte) (err error) {
	varDepositSwitchTargetAccount := _DepositSwitchTargetAccount{}

	if err = json.Unmarshal(bytes, &varDepositSwitchTargetAccount); err == nil {
		*o = DepositSwitchTargetAccount(varDepositSwitchTargetAccount)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "account_number")
		delete(additionalProperties, "routing_number")
		delete(additionalProperties, "account_name")
		delete(additionalProperties, "account_subtype")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDepositSwitchTargetAccount struct {
	value *DepositSwitchTargetAccount
	isSet bool
}

func (v NullableDepositSwitchTargetAccount) Get() *DepositSwitchTargetAccount {
	return v.value
}

func (v *NullableDepositSwitchTargetAccount) Set(val *DepositSwitchTargetAccount) {
	v.value = val
	v.isSet = true
}

func (v NullableDepositSwitchTargetAccount) IsSet() bool {
	return v.isSet
}

func (v *NullableDepositSwitchTargetAccount) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDepositSwitchTargetAccount(val *DepositSwitchTargetAccount) *NullableDepositSwitchTargetAccount {
	return &NullableDepositSwitchTargetAccount{value: val, isSet: true}
}

func (v NullableDepositSwitchTargetAccount) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDepositSwitchTargetAccount) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


