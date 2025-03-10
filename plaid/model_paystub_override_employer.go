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

// PaystubOverrideEmployer The employer on the paystub.
type PaystubOverrideEmployer struct {
	// The name of the employer.
	Name *string `json:"name,omitempty"`
}

// NewPaystubOverrideEmployer instantiates a new PaystubOverrideEmployer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaystubOverrideEmployer() *PaystubOverrideEmployer {
	this := PaystubOverrideEmployer{}
	return &this
}

// NewPaystubOverrideEmployerWithDefaults instantiates a new PaystubOverrideEmployer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaystubOverrideEmployerWithDefaults() *PaystubOverrideEmployer {
	this := PaystubOverrideEmployer{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PaystubOverrideEmployer) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaystubOverrideEmployer) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PaystubOverrideEmployer) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PaystubOverrideEmployer) SetName(v string) {
	o.Name = &v
}

func (o PaystubOverrideEmployer) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

type NullablePaystubOverrideEmployer struct {
	value *PaystubOverrideEmployer
	isSet bool
}

func (v NullablePaystubOverrideEmployer) Get() *PaystubOverrideEmployer {
	return v.value
}

func (v *NullablePaystubOverrideEmployer) Set(val *PaystubOverrideEmployer) {
	v.value = val
	v.isSet = true
}

func (v NullablePaystubOverrideEmployer) IsSet() bool {
	return v.isSet
}

func (v *NullablePaystubOverrideEmployer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaystubOverrideEmployer(val *PaystubOverrideEmployer) *NullablePaystubOverrideEmployer {
	return &NullablePaystubOverrideEmployer{value: val, isSet: true}
}

func (v NullablePaystubOverrideEmployer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaystubOverrideEmployer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


