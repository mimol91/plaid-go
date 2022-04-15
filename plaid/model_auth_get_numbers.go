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

// AuthGetNumbers An object containing identifying numbers used for making electronic transfers to and from the `accounts`. The identifying number type (ACH, EFT, IBAN, or BACS) used will depend on the country of the account. An account may have more than one number type. If a particular identifying number type is not used by any `accounts` for which data has been requested, the array for that type will be empty.
type AuthGetNumbers struct {
	// An array of ACH numbers identifying accounts.
	Ach []NumbersACH `json:"ach"`
	// An array of EFT numbers identifying accounts.
	Eft []NumbersEFT `json:"eft"`
	// An array of IBAN numbers identifying accounts.
	International []NumbersInternational `json:"international"`
	// An array of BACS numbers identifying accounts.
	Bacs []NumbersBACS `json:"bacs"`
	AdditionalProperties map[string]interface{}
}

type _AuthGetNumbers AuthGetNumbers

// NewAuthGetNumbers instantiates a new AuthGetNumbers object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthGetNumbers(ach []NumbersACH, eft []NumbersEFT, international []NumbersInternational, bacs []NumbersBACS) *AuthGetNumbers {
	this := AuthGetNumbers{}
	this.Ach = ach
	this.Eft = eft
	this.International = international
	this.Bacs = bacs
	return &this
}

// NewAuthGetNumbersWithDefaults instantiates a new AuthGetNumbers object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthGetNumbersWithDefaults() *AuthGetNumbers {
	this := AuthGetNumbers{}
	return &this
}

// GetAch returns the Ach field value
func (o *AuthGetNumbers) GetAch() []NumbersACH {
	if o == nil {
		var ret []NumbersACH
		return ret
	}

	return o.Ach
}

// GetAchOk returns a tuple with the Ach field value
// and a boolean to check if the value has been set.
func (o *AuthGetNumbers) GetAchOk() (*[]NumbersACH, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Ach, true
}

// SetAch sets field value
func (o *AuthGetNumbers) SetAch(v []NumbersACH) {
	o.Ach = v
}

// GetEft returns the Eft field value
func (o *AuthGetNumbers) GetEft() []NumbersEFT {
	if o == nil {
		var ret []NumbersEFT
		return ret
	}

	return o.Eft
}

// GetEftOk returns a tuple with the Eft field value
// and a boolean to check if the value has been set.
func (o *AuthGetNumbers) GetEftOk() (*[]NumbersEFT, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Eft, true
}

// SetEft sets field value
func (o *AuthGetNumbers) SetEft(v []NumbersEFT) {
	o.Eft = v
}

// GetInternational returns the International field value
func (o *AuthGetNumbers) GetInternational() []NumbersInternational {
	if o == nil {
		var ret []NumbersInternational
		return ret
	}

	return o.International
}

// GetInternationalOk returns a tuple with the International field value
// and a boolean to check if the value has been set.
func (o *AuthGetNumbers) GetInternationalOk() (*[]NumbersInternational, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.International, true
}

// SetInternational sets field value
func (o *AuthGetNumbers) SetInternational(v []NumbersInternational) {
	o.International = v
}

// GetBacs returns the Bacs field value
func (o *AuthGetNumbers) GetBacs() []NumbersBACS {
	if o == nil {
		var ret []NumbersBACS
		return ret
	}

	return o.Bacs
}

// GetBacsOk returns a tuple with the Bacs field value
// and a boolean to check if the value has been set.
func (o *AuthGetNumbers) GetBacsOk() (*[]NumbersBACS, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Bacs, true
}

// SetBacs sets field value
func (o *AuthGetNumbers) SetBacs(v []NumbersBACS) {
	o.Bacs = v
}

func (o AuthGetNumbers) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ach"] = o.Ach
	}
	if true {
		toSerialize["eft"] = o.Eft
	}
	if true {
		toSerialize["international"] = o.International
	}
	if true {
		toSerialize["bacs"] = o.Bacs
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *AuthGetNumbers) UnmarshalJSON(bytes []byte) (err error) {
	varAuthGetNumbers := _AuthGetNumbers{}

	if err = json.Unmarshal(bytes, &varAuthGetNumbers); err == nil {
		*o = AuthGetNumbers(varAuthGetNumbers)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ach")
		delete(additionalProperties, "eft")
		delete(additionalProperties, "international")
		delete(additionalProperties, "bacs")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAuthGetNumbers struct {
	value *AuthGetNumbers
	isSet bool
}

func (v NullableAuthGetNumbers) Get() *AuthGetNumbers {
	return v.value
}

func (v *NullableAuthGetNumbers) Set(val *AuthGetNumbers) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthGetNumbers) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthGetNumbers) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthGetNumbers(val *AuthGetNumbers) *NullableAuthGetNumbers {
	return &NullableAuthGetNumbers{value: val, isSet: true}
}

func (v NullableAuthGetNumbers) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthGetNumbers) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


