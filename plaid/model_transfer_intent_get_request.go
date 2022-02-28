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

// TransferIntentGetRequest Defines the request schema for `/transfer/intent/get`
type TransferIntentGetRequest struct {
	// Your Plaid API `client_id`. The `client_id` is required and may be provided either in the `PLAID-CLIENT-ID` header or as part of a request body.
	ClientId string `json:"client_id"`
	// Your Plaid API `secret`. The `secret` is required and may be provided either in the `PLAID-SECRET` header or as part of a request body.
	Secret string `json:"secret"`
	// Plaid's unique identifier for a transfer intent object.
	TransferIntentId string `json:"transfer_intent_id"`
	AdditionalProperties map[string]interface{}
}

type _TransferIntentGetRequest TransferIntentGetRequest

// NewTransferIntentGetRequest instantiates a new TransferIntentGetRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransferIntentGetRequest(clientId string, secret string, transferIntentId string) *TransferIntentGetRequest {
	this := TransferIntentGetRequest{}
	this.ClientId = clientId
	this.Secret = secret
	this.TransferIntentId = transferIntentId
	return &this
}

// NewTransferIntentGetRequestWithDefaults instantiates a new TransferIntentGetRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransferIntentGetRequestWithDefaults() *TransferIntentGetRequest {
	this := TransferIntentGetRequest{}
	return &this
}

// GetClientId returns the ClientId field value
func (o *TransferIntentGetRequest) GetClientId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value
// and a boolean to check if the value has been set.
func (o *TransferIntentGetRequest) GetClientIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ClientId, true
}

// SetClientId sets field value
func (o *TransferIntentGetRequest) SetClientId(v string) {
	o.ClientId = v
}

// GetSecret returns the Secret field value
func (o *TransferIntentGetRequest) GetSecret() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Secret
}

// GetSecretOk returns a tuple with the Secret field value
// and a boolean to check if the value has been set.
func (o *TransferIntentGetRequest) GetSecretOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Secret, true
}

// SetSecret sets field value
func (o *TransferIntentGetRequest) SetSecret(v string) {
	o.Secret = v
}

// GetTransferIntentId returns the TransferIntentId field value
func (o *TransferIntentGetRequest) GetTransferIntentId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TransferIntentId
}

// GetTransferIntentIdOk returns a tuple with the TransferIntentId field value
// and a boolean to check if the value has been set.
func (o *TransferIntentGetRequest) GetTransferIntentIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.TransferIntentId, true
}

// SetTransferIntentId sets field value
func (o *TransferIntentGetRequest) SetTransferIntentId(v string) {
	o.TransferIntentId = v
}

func (o TransferIntentGetRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["client_id"] = o.ClientId
	}
	if true {
		toSerialize["secret"] = o.Secret
	}
	if true {
		toSerialize["transfer_intent_id"] = o.TransferIntentId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *TransferIntentGetRequest) UnmarshalJSON(bytes []byte) (err error) {
	varTransferIntentGetRequest := _TransferIntentGetRequest{}

	if err = json.Unmarshal(bytes, &varTransferIntentGetRequest); err == nil {
		*o = TransferIntentGetRequest(varTransferIntentGetRequest)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "client_id")
		delete(additionalProperties, "secret")
		delete(additionalProperties, "transfer_intent_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTransferIntentGetRequest struct {
	value *TransferIntentGetRequest
	isSet bool
}

func (v NullableTransferIntentGetRequest) Get() *TransferIntentGetRequest {
	return v.value
}

func (v *NullableTransferIntentGetRequest) Set(val *TransferIntentGetRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableTransferIntentGetRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableTransferIntentGetRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransferIntentGetRequest(val *TransferIntentGetRequest) *NullableTransferIntentGetRequest {
	return &NullableTransferIntentGetRequest{value: val, isSet: true}
}

func (v NullableTransferIntentGetRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransferIntentGetRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


