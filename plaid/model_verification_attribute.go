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

// VerificationAttribute Details about a certain reason as to why a document could potentially be fraudulent
type VerificationAttribute struct {
	// Message indicating the reason as to why the verification failed
	Type NullableString `json:"type"`
	AdditionalProperties map[string]interface{}
}

type _VerificationAttribute VerificationAttribute

// NewVerificationAttribute instantiates a new VerificationAttribute object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVerificationAttribute(type_ NullableString) *VerificationAttribute {
	this := VerificationAttribute{}
	this.Type = type_
	return &this
}

// NewVerificationAttributeWithDefaults instantiates a new VerificationAttribute object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVerificationAttributeWithDefaults() *VerificationAttribute {
	this := VerificationAttribute{}
	return &this
}

// GetType returns the Type field value
// If the value is explicit nil, the zero value for string will be returned
func (o *VerificationAttribute) GetType() string {
	if o == nil || o.Type.Get() == nil {
		var ret string
		return ret
	}

	return *o.Type.Get()
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VerificationAttribute) GetTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Type.Get(), o.Type.IsSet()
}

// SetType sets field value
func (o *VerificationAttribute) SetType(v string) {
	o.Type.Set(&v)
}

func (o VerificationAttribute) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *VerificationAttribute) UnmarshalJSON(bytes []byte) (err error) {
	varVerificationAttribute := _VerificationAttribute{}

	if err = json.Unmarshal(bytes, &varVerificationAttribute); err == nil {
		*o = VerificationAttribute(varVerificationAttribute)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "type")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableVerificationAttribute struct {
	value *VerificationAttribute
	isSet bool
}

func (v NullableVerificationAttribute) Get() *VerificationAttribute {
	return v.value
}

func (v *NullableVerificationAttribute) Set(val *VerificationAttribute) {
	v.value = val
	v.isSet = true
}

func (v NullableVerificationAttribute) IsSet() bool {
	return v.isSet
}

func (v *NullableVerificationAttribute) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVerificationAttribute(val *VerificationAttribute) *NullableVerificationAttribute {
	return &NullableVerificationAttribute{value: val, isSet: true}
}

func (v NullableVerificationAttribute) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVerificationAttribute) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


