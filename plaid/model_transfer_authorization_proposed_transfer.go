/*
 * The Plaid API
 *
 * The Plaid REST API. Please see https://plaid.com/docs/api for more details.
 *
 * API version: 2020-09-14_1.77.1
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package plaid

import (
	"encoding/json"
)

// TransferAuthorizationProposedTransfer Details regarding the proposed transfer.
type TransferAuthorizationProposedTransfer struct {
	AchClass ACHClass `json:"ach_class"`
	// The Plaid `account_id` for the account that will be debited or credited.
	AccountId string `json:"account_id"`
	Type TransferType `json:"type"`
	User TransferUserInResponse `json:"user"`
	// The amount of the transfer (decimal string with two digits of precision e.g. \"10.00\").
	Amount string `json:"amount"`
	// The network or rails used for the transfer.
	Network string `json:"network"`
	// Plaid's unique identifier for the origination account that was used for this transfer.
	OriginationAccountId string `json:"origination_account_id"`
	// The currency of the transfer amount. The default value is \"USD\".
	IsoCurrencyCode string `json:"iso_currency_code"`
	AdditionalProperties map[string]interface{}
}

type _TransferAuthorizationProposedTransfer TransferAuthorizationProposedTransfer

// NewTransferAuthorizationProposedTransfer instantiates a new TransferAuthorizationProposedTransfer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransferAuthorizationProposedTransfer(achClass ACHClass, accountId string, type_ TransferType, user TransferUserInResponse, amount string, network string, originationAccountId string, isoCurrencyCode string) *TransferAuthorizationProposedTransfer {
	this := TransferAuthorizationProposedTransfer{}
	this.AchClass = achClass
	this.AccountId = accountId
	this.Type = type_
	this.User = user
	this.Amount = amount
	this.Network = network
	this.OriginationAccountId = originationAccountId
	this.IsoCurrencyCode = isoCurrencyCode
	return &this
}

// NewTransferAuthorizationProposedTransferWithDefaults instantiates a new TransferAuthorizationProposedTransfer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransferAuthorizationProposedTransferWithDefaults() *TransferAuthorizationProposedTransfer {
	this := TransferAuthorizationProposedTransfer{}
	return &this
}

// GetAchClass returns the AchClass field value
func (o *TransferAuthorizationProposedTransfer) GetAchClass() ACHClass {
	if o == nil {
		var ret ACHClass
		return ret
	}

	return o.AchClass
}

// GetAchClassOk returns a tuple with the AchClass field value
// and a boolean to check if the value has been set.
func (o *TransferAuthorizationProposedTransfer) GetAchClassOk() (*ACHClass, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AchClass, true
}

// SetAchClass sets field value
func (o *TransferAuthorizationProposedTransfer) SetAchClass(v ACHClass) {
	o.AchClass = v
}

// GetAccountId returns the AccountId field value
func (o *TransferAuthorizationProposedTransfer) GetAccountId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value
// and a boolean to check if the value has been set.
func (o *TransferAuthorizationProposedTransfer) GetAccountIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AccountId, true
}

// SetAccountId sets field value
func (o *TransferAuthorizationProposedTransfer) SetAccountId(v string) {
	o.AccountId = v
}

// GetType returns the Type field value
func (o *TransferAuthorizationProposedTransfer) GetType() TransferType {
	if o == nil {
		var ret TransferType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *TransferAuthorizationProposedTransfer) GetTypeOk() (*TransferType, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *TransferAuthorizationProposedTransfer) SetType(v TransferType) {
	o.Type = v
}

// GetUser returns the User field value
func (o *TransferAuthorizationProposedTransfer) GetUser() TransferUserInResponse {
	if o == nil {
		var ret TransferUserInResponse
		return ret
	}

	return o.User
}

// GetUserOk returns a tuple with the User field value
// and a boolean to check if the value has been set.
func (o *TransferAuthorizationProposedTransfer) GetUserOk() (*TransferUserInResponse, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.User, true
}

// SetUser sets field value
func (o *TransferAuthorizationProposedTransfer) SetUser(v TransferUserInResponse) {
	o.User = v
}

// GetAmount returns the Amount field value
func (o *TransferAuthorizationProposedTransfer) GetAmount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *TransferAuthorizationProposedTransfer) GetAmountOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *TransferAuthorizationProposedTransfer) SetAmount(v string) {
	o.Amount = v
}

// GetNetwork returns the Network field value
func (o *TransferAuthorizationProposedTransfer) GetNetwork() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Network
}

// GetNetworkOk returns a tuple with the Network field value
// and a boolean to check if the value has been set.
func (o *TransferAuthorizationProposedTransfer) GetNetworkOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Network, true
}

// SetNetwork sets field value
func (o *TransferAuthorizationProposedTransfer) SetNetwork(v string) {
	o.Network = v
}

// GetOriginationAccountId returns the OriginationAccountId field value
func (o *TransferAuthorizationProposedTransfer) GetOriginationAccountId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OriginationAccountId
}

// GetOriginationAccountIdOk returns a tuple with the OriginationAccountId field value
// and a boolean to check if the value has been set.
func (o *TransferAuthorizationProposedTransfer) GetOriginationAccountIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.OriginationAccountId, true
}

// SetOriginationAccountId sets field value
func (o *TransferAuthorizationProposedTransfer) SetOriginationAccountId(v string) {
	o.OriginationAccountId = v
}

// GetIsoCurrencyCode returns the IsoCurrencyCode field value
func (o *TransferAuthorizationProposedTransfer) GetIsoCurrencyCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.IsoCurrencyCode
}

// GetIsoCurrencyCodeOk returns a tuple with the IsoCurrencyCode field value
// and a boolean to check if the value has been set.
func (o *TransferAuthorizationProposedTransfer) GetIsoCurrencyCodeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.IsoCurrencyCode, true
}

// SetIsoCurrencyCode sets field value
func (o *TransferAuthorizationProposedTransfer) SetIsoCurrencyCode(v string) {
	o.IsoCurrencyCode = v
}

func (o TransferAuthorizationProposedTransfer) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ach_class"] = o.AchClass
	}
	if true {
		toSerialize["account_id"] = o.AccountId
	}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["user"] = o.User
	}
	if true {
		toSerialize["amount"] = o.Amount
	}
	if true {
		toSerialize["network"] = o.Network
	}
	if true {
		toSerialize["origination_account_id"] = o.OriginationAccountId
	}
	if true {
		toSerialize["iso_currency_code"] = o.IsoCurrencyCode
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *TransferAuthorizationProposedTransfer) UnmarshalJSON(bytes []byte) (err error) {
	varTransferAuthorizationProposedTransfer := _TransferAuthorizationProposedTransfer{}

	if err = json.Unmarshal(bytes, &varTransferAuthorizationProposedTransfer); err == nil {
		*o = TransferAuthorizationProposedTransfer(varTransferAuthorizationProposedTransfer)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ach_class")
		delete(additionalProperties, "account_id")
		delete(additionalProperties, "type")
		delete(additionalProperties, "user")
		delete(additionalProperties, "amount")
		delete(additionalProperties, "network")
		delete(additionalProperties, "origination_account_id")
		delete(additionalProperties, "iso_currency_code")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTransferAuthorizationProposedTransfer struct {
	value *TransferAuthorizationProposedTransfer
	isSet bool
}

func (v NullableTransferAuthorizationProposedTransfer) Get() *TransferAuthorizationProposedTransfer {
	return v.value
}

func (v *NullableTransferAuthorizationProposedTransfer) Set(val *TransferAuthorizationProposedTransfer) {
	v.value = val
	v.isSet = true
}

func (v NullableTransferAuthorizationProposedTransfer) IsSet() bool {
	return v.isSet
}

func (v *NullableTransferAuthorizationProposedTransfer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransferAuthorizationProposedTransfer(val *TransferAuthorizationProposedTransfer) *NullableTransferAuthorizationProposedTransfer {
	return &NullableTransferAuthorizationProposedTransfer{value: val, isSet: true}
}

func (v NullableTransferAuthorizationProposedTransfer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransferAuthorizationProposedTransfer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


