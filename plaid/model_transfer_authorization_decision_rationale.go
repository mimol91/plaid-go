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

// TransferAuthorizationDecisionRationale The rationale for Plaid's decision regarding a proposed transfer. Will be null for `approved` decisions.
type TransferAuthorizationDecisionRationale struct {
	// A code representing the rationale for permitting or declining the proposed transfer. Possible values are:  `NSF` – Transaction likely to result in a return due to insufficient funds.  `RISK` - Transaction is high-risk.  `MANUALLY_VERIFIED_ITEM` – Item created via same-day micro deposits, limited information available. Plaid can only offer `permitted` as a transaction decision.  `LOGIN_REQUIRED` – Unable to collect the account information required for an authorization decision due to Item staleness. Can be rectified using Link update mode.  `ERROR` – Unable to collect the account information required for an authorization decision due to an error.
	Code string `json:"code"`
	// A human-readable description of the code associated with a permitted transfer or transfer decline.
	Description string `json:"description"`
	AdditionalProperties map[string]interface{}
}

type _TransferAuthorizationDecisionRationale TransferAuthorizationDecisionRationale

// NewTransferAuthorizationDecisionRationale instantiates a new TransferAuthorizationDecisionRationale object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransferAuthorizationDecisionRationale(code string, description string) *TransferAuthorizationDecisionRationale {
	this := TransferAuthorizationDecisionRationale{}
	this.Code = code
	this.Description = description
	return &this
}

// NewTransferAuthorizationDecisionRationaleWithDefaults instantiates a new TransferAuthorizationDecisionRationale object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransferAuthorizationDecisionRationaleWithDefaults() *TransferAuthorizationDecisionRationale {
	this := TransferAuthorizationDecisionRationale{}
	return &this
}

// GetCode returns the Code field value
func (o *TransferAuthorizationDecisionRationale) GetCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *TransferAuthorizationDecisionRationale) GetCodeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *TransferAuthorizationDecisionRationale) SetCode(v string) {
	o.Code = v
}

// GetDescription returns the Description field value
func (o *TransferAuthorizationDecisionRationale) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *TransferAuthorizationDecisionRationale) GetDescriptionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *TransferAuthorizationDecisionRationale) SetDescription(v string) {
	o.Description = v
}

func (o TransferAuthorizationDecisionRationale) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["code"] = o.Code
	}
	if true {
		toSerialize["description"] = o.Description
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *TransferAuthorizationDecisionRationale) UnmarshalJSON(bytes []byte) (err error) {
	varTransferAuthorizationDecisionRationale := _TransferAuthorizationDecisionRationale{}

	if err = json.Unmarshal(bytes, &varTransferAuthorizationDecisionRationale); err == nil {
		*o = TransferAuthorizationDecisionRationale(varTransferAuthorizationDecisionRationale)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "code")
		delete(additionalProperties, "description")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTransferAuthorizationDecisionRationale struct {
	value *TransferAuthorizationDecisionRationale
	isSet bool
}

func (v NullableTransferAuthorizationDecisionRationale) Get() *TransferAuthorizationDecisionRationale {
	return v.value
}

func (v *NullableTransferAuthorizationDecisionRationale) Set(val *TransferAuthorizationDecisionRationale) {
	v.value = val
	v.isSet = true
}

func (v NullableTransferAuthorizationDecisionRationale) IsSet() bool {
	return v.isSet
}

func (v *NullableTransferAuthorizationDecisionRationale) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransferAuthorizationDecisionRationale(val *TransferAuthorizationDecisionRationale) *NullableTransferAuthorizationDecisionRationale {
	return &NullableTransferAuthorizationDecisionRationale{value: val, isSet: true}
}

func (v NullableTransferAuthorizationDecisionRationale) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransferAuthorizationDecisionRationale) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


