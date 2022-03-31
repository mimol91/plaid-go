/*
 * The Plaid API
 *
 * The Plaid REST API. Please see https://plaid.com/docs/api for more details.
 *
 * API version: 2020-09-14_1.94.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package plaid

import (
	"encoding/json"
)

// SandboxTransferSweepSimulateResponse Defines the response schema for `/sandbox/transfer/sweep/simulate`
type SandboxTransferSweepSimulateResponse struct {
	Sweep *SimulatedTransferSweep `json:"sweep,omitempty"`
	// A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive.
	RequestId string `json:"request_id"`
	AdditionalProperties map[string]interface{}
}

type _SandboxTransferSweepSimulateResponse SandboxTransferSweepSimulateResponse

// NewSandboxTransferSweepSimulateResponse instantiates a new SandboxTransferSweepSimulateResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSandboxTransferSweepSimulateResponse(requestId string) *SandboxTransferSweepSimulateResponse {
	this := SandboxTransferSweepSimulateResponse{}
	this.RequestId = requestId
	return &this
}

// NewSandboxTransferSweepSimulateResponseWithDefaults instantiates a new SandboxTransferSweepSimulateResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSandboxTransferSweepSimulateResponseWithDefaults() *SandboxTransferSweepSimulateResponse {
	this := SandboxTransferSweepSimulateResponse{}
	return &this
}

// GetSweep returns the Sweep field value if set, zero value otherwise.
func (o *SandboxTransferSweepSimulateResponse) GetSweep() SimulatedTransferSweep {
	if o == nil || o.Sweep == nil {
		var ret SimulatedTransferSweep
		return ret
	}
	return *o.Sweep
}

// GetSweepOk returns a tuple with the Sweep field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SandboxTransferSweepSimulateResponse) GetSweepOk() (*SimulatedTransferSweep, bool) {
	if o == nil || o.Sweep == nil {
		return nil, false
	}
	return o.Sweep, true
}

// HasSweep returns a boolean if a field has been set.
func (o *SandboxTransferSweepSimulateResponse) HasSweep() bool {
	if o != nil && o.Sweep != nil {
		return true
	}

	return false
}

// SetSweep gets a reference to the given SimulatedTransferSweep and assigns it to the Sweep field.
func (o *SandboxTransferSweepSimulateResponse) SetSweep(v SimulatedTransferSweep) {
	o.Sweep = &v
}

// GetRequestId returns the RequestId field value
func (o *SandboxTransferSweepSimulateResponse) GetRequestId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value
// and a boolean to check if the value has been set.
func (o *SandboxTransferSweepSimulateResponse) GetRequestIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RequestId, true
}

// SetRequestId sets field value
func (o *SandboxTransferSweepSimulateResponse) SetRequestId(v string) {
	o.RequestId = v
}

func (o SandboxTransferSweepSimulateResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Sweep != nil {
		toSerialize["sweep"] = o.Sweep
	}
	if true {
		toSerialize["request_id"] = o.RequestId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *SandboxTransferSweepSimulateResponse) UnmarshalJSON(bytes []byte) (err error) {
	varSandboxTransferSweepSimulateResponse := _SandboxTransferSweepSimulateResponse{}

	if err = json.Unmarshal(bytes, &varSandboxTransferSweepSimulateResponse); err == nil {
		*o = SandboxTransferSweepSimulateResponse(varSandboxTransferSweepSimulateResponse)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "sweep")
		delete(additionalProperties, "request_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSandboxTransferSweepSimulateResponse struct {
	value *SandboxTransferSweepSimulateResponse
	isSet bool
}

func (v NullableSandboxTransferSweepSimulateResponse) Get() *SandboxTransferSweepSimulateResponse {
	return v.value
}

func (v *NullableSandboxTransferSweepSimulateResponse) Set(val *SandboxTransferSweepSimulateResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSandboxTransferSweepSimulateResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSandboxTransferSweepSimulateResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSandboxTransferSweepSimulateResponse(val *SandboxTransferSweepSimulateResponse) *NullableSandboxTransferSweepSimulateResponse {
	return &NullableSandboxTransferSweepSimulateResponse{value: val, isSet: true}
}

func (v NullableSandboxTransferSweepSimulateResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSandboxTransferSweepSimulateResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


