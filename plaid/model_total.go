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

// Total An object representing both the current pay period and year to date amount for a category.
type Total struct {
	// Commonly used term to describe the line item.
	CanonicalDescription NullableString `json:"canonical_description,omitempty"`
	// Text of the line item as printed on the paystub.
	Description NullableString `json:"description,omitempty"`
	CurrentPay *Pay `json:"current_pay,omitempty"`
	YtdPay *Pay `json:"ytd_pay,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Total Total

// NewTotal instantiates a new Total object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTotal() *Total {
	this := Total{}
	return &this
}

// NewTotalWithDefaults instantiates a new Total object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTotalWithDefaults() *Total {
	this := Total{}
	return &this
}

// GetCanonicalDescription returns the CanonicalDescription field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Total) GetCanonicalDescription() string {
	if o == nil || o.CanonicalDescription.Get() == nil {
		var ret string
		return ret
	}
	return *o.CanonicalDescription.Get()
}

// GetCanonicalDescriptionOk returns a tuple with the CanonicalDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Total) GetCanonicalDescriptionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.CanonicalDescription.Get(), o.CanonicalDescription.IsSet()
}

// HasCanonicalDescription returns a boolean if a field has been set.
func (o *Total) HasCanonicalDescription() bool {
	if o != nil && o.CanonicalDescription.IsSet() {
		return true
	}

	return false
}

// SetCanonicalDescription gets a reference to the given NullableString and assigns it to the CanonicalDescription field.
func (o *Total) SetCanonicalDescription(v string) {
	o.CanonicalDescription.Set(&v)
}
// SetCanonicalDescriptionNil sets the value for CanonicalDescription to be an explicit nil
func (o *Total) SetCanonicalDescriptionNil() {
	o.CanonicalDescription.Set(nil)
}

// UnsetCanonicalDescription ensures that no value is present for CanonicalDescription, not even an explicit nil
func (o *Total) UnsetCanonicalDescription() {
	o.CanonicalDescription.Unset()
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Total) GetDescription() string {
	if o == nil || o.Description.Get() == nil {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Total) GetDescriptionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *Total) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *Total) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *Total) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *Total) UnsetDescription() {
	o.Description.Unset()
}

// GetCurrentPay returns the CurrentPay field value if set, zero value otherwise.
func (o *Total) GetCurrentPay() Pay {
	if o == nil || o.CurrentPay == nil {
		var ret Pay
		return ret
	}
	return *o.CurrentPay
}

// GetCurrentPayOk returns a tuple with the CurrentPay field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Total) GetCurrentPayOk() (*Pay, bool) {
	if o == nil || o.CurrentPay == nil {
		return nil, false
	}
	return o.CurrentPay, true
}

// HasCurrentPay returns a boolean if a field has been set.
func (o *Total) HasCurrentPay() bool {
	if o != nil && o.CurrentPay != nil {
		return true
	}

	return false
}

// SetCurrentPay gets a reference to the given Pay and assigns it to the CurrentPay field.
func (o *Total) SetCurrentPay(v Pay) {
	o.CurrentPay = &v
}

// GetYtdPay returns the YtdPay field value if set, zero value otherwise.
func (o *Total) GetYtdPay() Pay {
	if o == nil || o.YtdPay == nil {
		var ret Pay
		return ret
	}
	return *o.YtdPay
}

// GetYtdPayOk returns a tuple with the YtdPay field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Total) GetYtdPayOk() (*Pay, bool) {
	if o == nil || o.YtdPay == nil {
		return nil, false
	}
	return o.YtdPay, true
}

// HasYtdPay returns a boolean if a field has been set.
func (o *Total) HasYtdPay() bool {
	if o != nil && o.YtdPay != nil {
		return true
	}

	return false
}

// SetYtdPay gets a reference to the given Pay and assigns it to the YtdPay field.
func (o *Total) SetYtdPay(v Pay) {
	o.YtdPay = &v
}

func (o Total) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CanonicalDescription.IsSet() {
		toSerialize["canonical_description"] = o.CanonicalDescription.Get()
	}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if o.CurrentPay != nil {
		toSerialize["current_pay"] = o.CurrentPay
	}
	if o.YtdPay != nil {
		toSerialize["ytd_pay"] = o.YtdPay
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *Total) UnmarshalJSON(bytes []byte) (err error) {
	varTotal := _Total{}

	if err = json.Unmarshal(bytes, &varTotal); err == nil {
		*o = Total(varTotal)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "canonical_description")
		delete(additionalProperties, "description")
		delete(additionalProperties, "current_pay")
		delete(additionalProperties, "ytd_pay")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTotal struct {
	value *Total
	isSet bool
}

func (v NullableTotal) Get() *Total {
	return v.value
}

func (v *NullableTotal) Set(val *Total) {
	v.value = val
	v.isSet = true
}

func (v NullableTotal) IsSet() bool {
	return v.isSet
}

func (v *NullableTotal) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTotal(val *Total) *NullableTotal {
	return &NullableTotal{value: val, isSet: true}
}

func (v NullableTotal) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTotal) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


