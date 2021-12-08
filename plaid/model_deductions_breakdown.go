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

// DeductionsBreakdown An object representing the deduction line items for the pay period
type DeductionsBreakdown struct {
	// Raw amount of the deduction
	CurrentAmount NullableFloat32 `json:"current_amount,omitempty"`
	// Description of the deduction line item
	Description NullableString `json:"description,omitempty"`
	// The ISO-4217 currency code of the line item. Always `null` if `unofficial_currency_code` is non-null.
	IsoCurrencyCode NullableString `json:"iso_currency_code,omitempty"`
	// The unofficial currency code associated with the line item. Always `null` if `iso_currency_code` is non-`null`. Unofficial currency codes are used for currencies that do not have official ISO currency codes, such as cryptocurrencies and the currencies of certain countries.  See the [currency code schema](https://plaid.com/docs/api/accounts#currency-code-schema) for a full listing of supported `iso_currency_code`s.
	UnofficialCurrencyCode NullableString `json:"unofficial_currency_code,omitempty"`
	// The year-to-date amount of the deduction
	YtdAmount NullableFloat32 `json:"ytd_amount,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _DeductionsBreakdown DeductionsBreakdown

// NewDeductionsBreakdown instantiates a new DeductionsBreakdown object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeductionsBreakdown() *DeductionsBreakdown {
	this := DeductionsBreakdown{}
	return &this
}

// NewDeductionsBreakdownWithDefaults instantiates a new DeductionsBreakdown object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeductionsBreakdownWithDefaults() *DeductionsBreakdown {
	this := DeductionsBreakdown{}
	return &this
}

// GetCurrentAmount returns the CurrentAmount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DeductionsBreakdown) GetCurrentAmount() float32 {
	if o == nil || o.CurrentAmount.Get() == nil {
		var ret float32
		return ret
	}
	return *o.CurrentAmount.Get()
}

// GetCurrentAmountOk returns a tuple with the CurrentAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DeductionsBreakdown) GetCurrentAmountOk() (*float32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.CurrentAmount.Get(), o.CurrentAmount.IsSet()
}

// HasCurrentAmount returns a boolean if a field has been set.
func (o *DeductionsBreakdown) HasCurrentAmount() bool {
	if o != nil && o.CurrentAmount.IsSet() {
		return true
	}

	return false
}

// SetCurrentAmount gets a reference to the given NullableFloat32 and assigns it to the CurrentAmount field.
func (o *DeductionsBreakdown) SetCurrentAmount(v float32) {
	o.CurrentAmount.Set(&v)
}
// SetCurrentAmountNil sets the value for CurrentAmount to be an explicit nil
func (o *DeductionsBreakdown) SetCurrentAmountNil() {
	o.CurrentAmount.Set(nil)
}

// UnsetCurrentAmount ensures that no value is present for CurrentAmount, not even an explicit nil
func (o *DeductionsBreakdown) UnsetCurrentAmount() {
	o.CurrentAmount.Unset()
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DeductionsBreakdown) GetDescription() string {
	if o == nil || o.Description.Get() == nil {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DeductionsBreakdown) GetDescriptionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *DeductionsBreakdown) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *DeductionsBreakdown) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *DeductionsBreakdown) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *DeductionsBreakdown) UnsetDescription() {
	o.Description.Unset()
}

// GetIsoCurrencyCode returns the IsoCurrencyCode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DeductionsBreakdown) GetIsoCurrencyCode() string {
	if o == nil || o.IsoCurrencyCode.Get() == nil {
		var ret string
		return ret
	}
	return *o.IsoCurrencyCode.Get()
}

// GetIsoCurrencyCodeOk returns a tuple with the IsoCurrencyCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DeductionsBreakdown) GetIsoCurrencyCodeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.IsoCurrencyCode.Get(), o.IsoCurrencyCode.IsSet()
}

// HasIsoCurrencyCode returns a boolean if a field has been set.
func (o *DeductionsBreakdown) HasIsoCurrencyCode() bool {
	if o != nil && o.IsoCurrencyCode.IsSet() {
		return true
	}

	return false
}

// SetIsoCurrencyCode gets a reference to the given NullableString and assigns it to the IsoCurrencyCode field.
func (o *DeductionsBreakdown) SetIsoCurrencyCode(v string) {
	o.IsoCurrencyCode.Set(&v)
}
// SetIsoCurrencyCodeNil sets the value for IsoCurrencyCode to be an explicit nil
func (o *DeductionsBreakdown) SetIsoCurrencyCodeNil() {
	o.IsoCurrencyCode.Set(nil)
}

// UnsetIsoCurrencyCode ensures that no value is present for IsoCurrencyCode, not even an explicit nil
func (o *DeductionsBreakdown) UnsetIsoCurrencyCode() {
	o.IsoCurrencyCode.Unset()
}

// GetUnofficialCurrencyCode returns the UnofficialCurrencyCode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DeductionsBreakdown) GetUnofficialCurrencyCode() string {
	if o == nil || o.UnofficialCurrencyCode.Get() == nil {
		var ret string
		return ret
	}
	return *o.UnofficialCurrencyCode.Get()
}

// GetUnofficialCurrencyCodeOk returns a tuple with the UnofficialCurrencyCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DeductionsBreakdown) GetUnofficialCurrencyCodeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.UnofficialCurrencyCode.Get(), o.UnofficialCurrencyCode.IsSet()
}

// HasUnofficialCurrencyCode returns a boolean if a field has been set.
func (o *DeductionsBreakdown) HasUnofficialCurrencyCode() bool {
	if o != nil && o.UnofficialCurrencyCode.IsSet() {
		return true
	}

	return false
}

// SetUnofficialCurrencyCode gets a reference to the given NullableString and assigns it to the UnofficialCurrencyCode field.
func (o *DeductionsBreakdown) SetUnofficialCurrencyCode(v string) {
	o.UnofficialCurrencyCode.Set(&v)
}
// SetUnofficialCurrencyCodeNil sets the value for UnofficialCurrencyCode to be an explicit nil
func (o *DeductionsBreakdown) SetUnofficialCurrencyCodeNil() {
	o.UnofficialCurrencyCode.Set(nil)
}

// UnsetUnofficialCurrencyCode ensures that no value is present for UnofficialCurrencyCode, not even an explicit nil
func (o *DeductionsBreakdown) UnsetUnofficialCurrencyCode() {
	o.UnofficialCurrencyCode.Unset()
}

// GetYtdAmount returns the YtdAmount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DeductionsBreakdown) GetYtdAmount() float32 {
	if o == nil || o.YtdAmount.Get() == nil {
		var ret float32
		return ret
	}
	return *o.YtdAmount.Get()
}

// GetYtdAmountOk returns a tuple with the YtdAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DeductionsBreakdown) GetYtdAmountOk() (*float32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.YtdAmount.Get(), o.YtdAmount.IsSet()
}

// HasYtdAmount returns a boolean if a field has been set.
func (o *DeductionsBreakdown) HasYtdAmount() bool {
	if o != nil && o.YtdAmount.IsSet() {
		return true
	}

	return false
}

// SetYtdAmount gets a reference to the given NullableFloat32 and assigns it to the YtdAmount field.
func (o *DeductionsBreakdown) SetYtdAmount(v float32) {
	o.YtdAmount.Set(&v)
}
// SetYtdAmountNil sets the value for YtdAmount to be an explicit nil
func (o *DeductionsBreakdown) SetYtdAmountNil() {
	o.YtdAmount.Set(nil)
}

// UnsetYtdAmount ensures that no value is present for YtdAmount, not even an explicit nil
func (o *DeductionsBreakdown) UnsetYtdAmount() {
	o.YtdAmount.Unset()
}

func (o DeductionsBreakdown) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CurrentAmount.IsSet() {
		toSerialize["current_amount"] = o.CurrentAmount.Get()
	}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if o.IsoCurrencyCode.IsSet() {
		toSerialize["iso_currency_code"] = o.IsoCurrencyCode.Get()
	}
	if o.UnofficialCurrencyCode.IsSet() {
		toSerialize["unofficial_currency_code"] = o.UnofficialCurrencyCode.Get()
	}
	if o.YtdAmount.IsSet() {
		toSerialize["ytd_amount"] = o.YtdAmount.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *DeductionsBreakdown) UnmarshalJSON(bytes []byte) (err error) {
	varDeductionsBreakdown := _DeductionsBreakdown{}

	if err = json.Unmarshal(bytes, &varDeductionsBreakdown); err == nil {
		*o = DeductionsBreakdown(varDeductionsBreakdown)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "current_amount")
		delete(additionalProperties, "description")
		delete(additionalProperties, "iso_currency_code")
		delete(additionalProperties, "unofficial_currency_code")
		delete(additionalProperties, "ytd_amount")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDeductionsBreakdown struct {
	value *DeductionsBreakdown
	isSet bool
}

func (v NullableDeductionsBreakdown) Get() *DeductionsBreakdown {
	return v.value
}

func (v *NullableDeductionsBreakdown) Set(val *DeductionsBreakdown) {
	v.value = val
	v.isSet = true
}

func (v NullableDeductionsBreakdown) IsSet() bool {
	return v.isSet
}

func (v *NullableDeductionsBreakdown) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeductionsBreakdown(val *DeductionsBreakdown) *NullableDeductionsBreakdown {
	return &NullableDeductionsBreakdown{value: val, isSet: true}
}

func (v NullableDeductionsBreakdown) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeductionsBreakdown) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


