/*
 * The Plaid API
 *
 * The Plaid REST API. Please see https://plaid.com/docs/api for more details.
 *
 * API version: 2020-09-14_1.54.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package plaid

import (
	"encoding/json"
)

// TransferAuthorizationCreateResponse Defines the response schema for `/transfer/authorization/create`
type TransferAuthorizationCreateResponse struct {
	Authorization TransferAuthorization `json:"authorization"`
	// A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive.
	RequestId string `json:"request_id"`
	AdditionalProperties map[string]interface{}
}

type _TransferAuthorizationCreateResponse TransferAuthorizationCreateResponse

// NewTransferAuthorizationCreateResponse instantiates a new TransferAuthorizationCreateResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransferAuthorizationCreateResponse(authorization TransferAuthorization, requestId string) *TransferAuthorizationCreateResponse {
	this := TransferAuthorizationCreateResponse{}
	this.Authorization = authorization
	this.RequestId = requestId
	return &this
}

// NewTransferAuthorizationCreateResponseWithDefaults instantiates a new TransferAuthorizationCreateResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransferAuthorizationCreateResponseWithDefaults() *TransferAuthorizationCreateResponse {
	this := TransferAuthorizationCreateResponse{}
	return &this
}

// GetAuthorization returns the Authorization field value
func (o *TransferAuthorizationCreateResponse) GetAuthorization() TransferAuthorization {
	if o == nil {
		var ret TransferAuthorization
		return ret
	}

	return o.Authorization
}

// GetAuthorizationOk returns a tuple with the Authorization field value
// and a boolean to check if the value has been set.
func (o *TransferAuthorizationCreateResponse) GetAuthorizationOk() (*TransferAuthorization, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Authorization, true
}

// SetAuthorization sets field value
func (o *TransferAuthorizationCreateResponse) SetAuthorization(v TransferAuthorization) {
	o.Authorization = v
}

// GetRequestId returns the RequestId field value
func (o *TransferAuthorizationCreateResponse) GetRequestId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value
// and a boolean to check if the value has been set.
func (o *TransferAuthorizationCreateResponse) GetRequestIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RequestId, true
}

// SetRequestId sets field value
func (o *TransferAuthorizationCreateResponse) SetRequestId(v string) {
	o.RequestId = v
}

func (o TransferAuthorizationCreateResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["authorization"] = o.Authorization
	}
	if true {
		toSerialize["request_id"] = o.RequestId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *TransferAuthorizationCreateResponse) UnmarshalJSON(bytes []byte) (err error) {
	varTransferAuthorizationCreateResponse := _TransferAuthorizationCreateResponse{}

	if err = json.Unmarshal(bytes, &varTransferAuthorizationCreateResponse); err == nil {
		*o = TransferAuthorizationCreateResponse(varTransferAuthorizationCreateResponse)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "authorization")
		delete(additionalProperties, "request_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTransferAuthorizationCreateResponse struct {
	value *TransferAuthorizationCreateResponse
	isSet bool
}

func (v NullableTransferAuthorizationCreateResponse) Get() *TransferAuthorizationCreateResponse {
	return v.value
}

func (v *NullableTransferAuthorizationCreateResponse) Set(val *TransferAuthorizationCreateResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableTransferAuthorizationCreateResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableTransferAuthorizationCreateResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransferAuthorizationCreateResponse(val *TransferAuthorizationCreateResponse) *NullableTransferAuthorizationCreateResponse {
	return &NullableTransferAuthorizationCreateResponse{value: val, isSet: true}
}

func (v NullableTransferAuthorizationCreateResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransferAuthorizationCreateResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


