/*
 * The Plaid API
 *
 * The Plaid REST API. Please see https://plaid.com/docs/api for more details.
 *
 * API version: 2020-09-14_1.36.1
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package plaid

import (
	"encoding/json"
)

// LiabilitiesGetRequestOptions An optional object to filter `/liabilities/get` results. If provided, `options` cannot be null.
type LiabilitiesGetRequestOptions struct {
	// A list of accounts to retrieve for the Item.  An error will be returned if a provided `account_id` is not associated with the Item
	AccountIds *[]string `json:"account_ids,omitempty"`
}

// NewLiabilitiesGetRequestOptions instantiates a new LiabilitiesGetRequestOptions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLiabilitiesGetRequestOptions() *LiabilitiesGetRequestOptions {
	this := LiabilitiesGetRequestOptions{}
	return &this
}

// NewLiabilitiesGetRequestOptionsWithDefaults instantiates a new LiabilitiesGetRequestOptions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLiabilitiesGetRequestOptionsWithDefaults() *LiabilitiesGetRequestOptions {
	this := LiabilitiesGetRequestOptions{}
	return &this
}

// GetAccountIds returns the AccountIds field value if set, zero value otherwise.
func (o *LiabilitiesGetRequestOptions) GetAccountIds() []string {
	if o == nil || o.AccountIds == nil {
		var ret []string
		return ret
	}
	return *o.AccountIds
}

// GetAccountIdsOk returns a tuple with the AccountIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LiabilitiesGetRequestOptions) GetAccountIdsOk() (*[]string, bool) {
	if o == nil || o.AccountIds == nil {
		return nil, false
	}
	return o.AccountIds, true
}

// HasAccountIds returns a boolean if a field has been set.
func (o *LiabilitiesGetRequestOptions) HasAccountIds() bool {
	if o != nil && o.AccountIds != nil {
		return true
	}

	return false
}

// SetAccountIds gets a reference to the given []string and assigns it to the AccountIds field.
func (o *LiabilitiesGetRequestOptions) SetAccountIds(v []string) {
	o.AccountIds = &v
}

func (o LiabilitiesGetRequestOptions) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AccountIds != nil {
		toSerialize["account_ids"] = o.AccountIds
	}
	return json.Marshal(toSerialize)
}

type NullableLiabilitiesGetRequestOptions struct {
	value *LiabilitiesGetRequestOptions
	isSet bool
}

func (v NullableLiabilitiesGetRequestOptions) Get() *LiabilitiesGetRequestOptions {
	return v.value
}

func (v *NullableLiabilitiesGetRequestOptions) Set(val *LiabilitiesGetRequestOptions) {
	v.value = val
	v.isSet = true
}

func (v NullableLiabilitiesGetRequestOptions) IsSet() bool {
	return v.isSet
}

func (v *NullableLiabilitiesGetRequestOptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLiabilitiesGetRequestOptions(val *LiabilitiesGetRequestOptions) *NullableLiabilitiesGetRequestOptions {
	return &NullableLiabilitiesGetRequestOptions{value: val, isSet: true}
}

func (v NullableLiabilitiesGetRequestOptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLiabilitiesGetRequestOptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


