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

// IncomeVerificationPrecheckResponse IncomeVerificationPrecheckResponse defines the response schema for `/income/verification/precheck`.
type IncomeVerificationPrecheckResponse struct {
	// ID of the precheck.
	PrecheckId string `json:"precheck_id"`
	// A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive.
	RequestId string `json:"request_id"`
	Confidence IncomeVerificationPrecheckConfidence `json:"confidence"`
}

// NewIncomeVerificationPrecheckResponse instantiates a new IncomeVerificationPrecheckResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIncomeVerificationPrecheckResponse(precheckId string, requestId string, confidence IncomeVerificationPrecheckConfidence) *IncomeVerificationPrecheckResponse {
	this := IncomeVerificationPrecheckResponse{}
	this.PrecheckId = precheckId
	this.RequestId = requestId
	this.Confidence = confidence
	return &this
}

// NewIncomeVerificationPrecheckResponseWithDefaults instantiates a new IncomeVerificationPrecheckResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIncomeVerificationPrecheckResponseWithDefaults() *IncomeVerificationPrecheckResponse {
	this := IncomeVerificationPrecheckResponse{}
	return &this
}

// GetPrecheckId returns the PrecheckId field value
func (o *IncomeVerificationPrecheckResponse) GetPrecheckId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PrecheckId
}

// GetPrecheckIdOk returns a tuple with the PrecheckId field value
// and a boolean to check if the value has been set.
func (o *IncomeVerificationPrecheckResponse) GetPrecheckIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.PrecheckId, true
}

// SetPrecheckId sets field value
func (o *IncomeVerificationPrecheckResponse) SetPrecheckId(v string) {
	o.PrecheckId = v
}

// GetRequestId returns the RequestId field value
func (o *IncomeVerificationPrecheckResponse) GetRequestId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value
// and a boolean to check if the value has been set.
func (o *IncomeVerificationPrecheckResponse) GetRequestIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RequestId, true
}

// SetRequestId sets field value
func (o *IncomeVerificationPrecheckResponse) SetRequestId(v string) {
	o.RequestId = v
}

// GetConfidence returns the Confidence field value
func (o *IncomeVerificationPrecheckResponse) GetConfidence() IncomeVerificationPrecheckConfidence {
	if o == nil {
		var ret IncomeVerificationPrecheckConfidence
		return ret
	}

	return o.Confidence
}

// GetConfidenceOk returns a tuple with the Confidence field value
// and a boolean to check if the value has been set.
func (o *IncomeVerificationPrecheckResponse) GetConfidenceOk() (*IncomeVerificationPrecheckConfidence, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Confidence, true
}

// SetConfidence sets field value
func (o *IncomeVerificationPrecheckResponse) SetConfidence(v IncomeVerificationPrecheckConfidence) {
	o.Confidence = v
}

func (o IncomeVerificationPrecheckResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["precheck_id"] = o.PrecheckId
	}
	if true {
		toSerialize["request_id"] = o.RequestId
	}
	if true {
		toSerialize["confidence"] = o.Confidence
	}
	return json.Marshal(toSerialize)
}

type NullableIncomeVerificationPrecheckResponse struct {
	value *IncomeVerificationPrecheckResponse
	isSet bool
}

func (v NullableIncomeVerificationPrecheckResponse) Get() *IncomeVerificationPrecheckResponse {
	return v.value
}

func (v *NullableIncomeVerificationPrecheckResponse) Set(val *IncomeVerificationPrecheckResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableIncomeVerificationPrecheckResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableIncomeVerificationPrecheckResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIncomeVerificationPrecheckResponse(val *IncomeVerificationPrecheckResponse) *NullableIncomeVerificationPrecheckResponse {
	return &NullableIncomeVerificationPrecheckResponse{value: val, isSet: true}
}

func (v NullableIncomeVerificationPrecheckResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIncomeVerificationPrecheckResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


