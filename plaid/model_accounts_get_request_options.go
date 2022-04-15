/*
 * The Plaid API
 *
 * The Plaid REST API. Please see https://plaid.com/docs/api for more details.
 *
 * API version: 2020-09-14_1.97.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package plaid

import (
	"encoding/json"
)

// AccountsGetRequestOptions An optional object to filter `/accounts/get` results.
type AccountsGetRequestOptions struct {
	// An array of `account_ids` to retrieve for the Account.
	AccountIds *[]string `json:"account_ids,omitempty"`
}

// NewAccountsGetRequestOptions instantiates a new AccountsGetRequestOptions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountsGetRequestOptions() *AccountsGetRequestOptions {
	this := AccountsGetRequestOptions{}
	return &this
}

// NewAccountsGetRequestOptionsWithDefaults instantiates a new AccountsGetRequestOptions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountsGetRequestOptionsWithDefaults() *AccountsGetRequestOptions {
	this := AccountsGetRequestOptions{}
	return &this
}

// GetAccountIds returns the AccountIds field value if set, zero value otherwise.
func (o *AccountsGetRequestOptions) GetAccountIds() []string {
	if o == nil || o.AccountIds == nil {
		var ret []string
		return ret
	}
	return *o.AccountIds
}

// GetAccountIdsOk returns a tuple with the AccountIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountsGetRequestOptions) GetAccountIdsOk() (*[]string, bool) {
	if o == nil || o.AccountIds == nil {
		return nil, false
	}
	return o.AccountIds, true
}

// HasAccountIds returns a boolean if a field has been set.
func (o *AccountsGetRequestOptions) HasAccountIds() bool {
	if o != nil && o.AccountIds != nil {
		return true
	}

	return false
}

// SetAccountIds gets a reference to the given []string and assigns it to the AccountIds field.
func (o *AccountsGetRequestOptions) SetAccountIds(v []string) {
	o.AccountIds = &v
}

func (o AccountsGetRequestOptions) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AccountIds != nil {
		toSerialize["account_ids"] = o.AccountIds
	}
	return json.Marshal(toSerialize)
}

type NullableAccountsGetRequestOptions struct {
	value *AccountsGetRequestOptions
	isSet bool
}

func (v NullableAccountsGetRequestOptions) Get() *AccountsGetRequestOptions {
	return v.value
}

func (v *NullableAccountsGetRequestOptions) Set(val *AccountsGetRequestOptions) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountsGetRequestOptions) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountsGetRequestOptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountsGetRequestOptions(val *AccountsGetRequestOptions) *NullableAccountsGetRequestOptions {
	return &NullableAccountsGetRequestOptions{value: val, isSet: true}
}

func (v NullableAccountsGetRequestOptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountsGetRequestOptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


