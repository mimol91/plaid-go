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

// LinkTokenCreateRequestUpdate Specifies options for initializing Link for [update mode](https://plaid.com/docs/link/update-mode).
type LinkTokenCreateRequestUpdate struct {
	// If `true`, enables [update mode with Account Select](https://plaid.com/docs/link/update-mode/#using-update-mode-to-request-new-accounts).
	AccountSelectionEnabled *bool `json:"account_selection_enabled,omitempty"`
}

// NewLinkTokenCreateRequestUpdate instantiates a new LinkTokenCreateRequestUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLinkTokenCreateRequestUpdate() *LinkTokenCreateRequestUpdate {
	this := LinkTokenCreateRequestUpdate{}
	var accountSelectionEnabled bool = false
	this.AccountSelectionEnabled = &accountSelectionEnabled
	return &this
}

// NewLinkTokenCreateRequestUpdateWithDefaults instantiates a new LinkTokenCreateRequestUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLinkTokenCreateRequestUpdateWithDefaults() *LinkTokenCreateRequestUpdate {
	this := LinkTokenCreateRequestUpdate{}
	var accountSelectionEnabled bool = false
	this.AccountSelectionEnabled = &accountSelectionEnabled
	return &this
}

// GetAccountSelectionEnabled returns the AccountSelectionEnabled field value if set, zero value otherwise.
func (o *LinkTokenCreateRequestUpdate) GetAccountSelectionEnabled() bool {
	if o == nil || o.AccountSelectionEnabled == nil {
		var ret bool
		return ret
	}
	return *o.AccountSelectionEnabled
}

// GetAccountSelectionEnabledOk returns a tuple with the AccountSelectionEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinkTokenCreateRequestUpdate) GetAccountSelectionEnabledOk() (*bool, bool) {
	if o == nil || o.AccountSelectionEnabled == nil {
		return nil, false
	}
	return o.AccountSelectionEnabled, true
}

// HasAccountSelectionEnabled returns a boolean if a field has been set.
func (o *LinkTokenCreateRequestUpdate) HasAccountSelectionEnabled() bool {
	if o != nil && o.AccountSelectionEnabled != nil {
		return true
	}

	return false
}

// SetAccountSelectionEnabled gets a reference to the given bool and assigns it to the AccountSelectionEnabled field.
func (o *LinkTokenCreateRequestUpdate) SetAccountSelectionEnabled(v bool) {
	o.AccountSelectionEnabled = &v
}

func (o LinkTokenCreateRequestUpdate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AccountSelectionEnabled != nil {
		toSerialize["account_selection_enabled"] = o.AccountSelectionEnabled
	}
	return json.Marshal(toSerialize)
}

type NullableLinkTokenCreateRequestUpdate struct {
	value *LinkTokenCreateRequestUpdate
	isSet bool
}

func (v NullableLinkTokenCreateRequestUpdate) Get() *LinkTokenCreateRequestUpdate {
	return v.value
}

func (v *NullableLinkTokenCreateRequestUpdate) Set(val *LinkTokenCreateRequestUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullableLinkTokenCreateRequestUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullableLinkTokenCreateRequestUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLinkTokenCreateRequestUpdate(val *LinkTokenCreateRequestUpdate) *NullableLinkTokenCreateRequestUpdate {
	return &NullableLinkTokenCreateRequestUpdate{value: val, isSet: true}
}

func (v NullableLinkTokenCreateRequestUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLinkTokenCreateRequestUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


