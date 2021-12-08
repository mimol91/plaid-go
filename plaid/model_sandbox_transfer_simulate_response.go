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

// SandboxTransferSimulateResponse Defines the response schema for `/sandbox/transfer/simulate`
type SandboxTransferSimulateResponse struct {
	// A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive.
	RequestId string `json:"request_id"`
	AdditionalProperties map[string]interface{}
}

type _SandboxTransferSimulateResponse SandboxTransferSimulateResponse

// NewSandboxTransferSimulateResponse instantiates a new SandboxTransferSimulateResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSandboxTransferSimulateResponse(requestId string) *SandboxTransferSimulateResponse {
	this := SandboxTransferSimulateResponse{}
	this.RequestId = requestId
	return &this
}

// NewSandboxTransferSimulateResponseWithDefaults instantiates a new SandboxTransferSimulateResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSandboxTransferSimulateResponseWithDefaults() *SandboxTransferSimulateResponse {
	this := SandboxTransferSimulateResponse{}
	return &this
}

// GetRequestId returns the RequestId field value
func (o *SandboxTransferSimulateResponse) GetRequestId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value
// and a boolean to check if the value has been set.
func (o *SandboxTransferSimulateResponse) GetRequestIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RequestId, true
}

// SetRequestId sets field value
func (o *SandboxTransferSimulateResponse) SetRequestId(v string) {
	o.RequestId = v
}

func (o SandboxTransferSimulateResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["request_id"] = o.RequestId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *SandboxTransferSimulateResponse) UnmarshalJSON(bytes []byte) (err error) {
	varSandboxTransferSimulateResponse := _SandboxTransferSimulateResponse{}

	if err = json.Unmarshal(bytes, &varSandboxTransferSimulateResponse); err == nil {
		*o = SandboxTransferSimulateResponse(varSandboxTransferSimulateResponse)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "request_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSandboxTransferSimulateResponse struct {
	value *SandboxTransferSimulateResponse
	isSet bool
}

func (v NullableSandboxTransferSimulateResponse) Get() *SandboxTransferSimulateResponse {
	return v.value
}

func (v *NullableSandboxTransferSimulateResponse) Set(val *SandboxTransferSimulateResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSandboxTransferSimulateResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSandboxTransferSimulateResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSandboxTransferSimulateResponse(val *SandboxTransferSimulateResponse) *NullableSandboxTransferSimulateResponse {
	return &NullableSandboxTransferSimulateResponse{value: val, isSet: true}
}

func (v NullableSandboxTransferSimulateResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSandboxTransferSimulateResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


