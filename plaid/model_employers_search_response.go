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

// EmployersSearchResponse EmployersSearchResponse defines the response schema for `/employers/search`.
type EmployersSearchResponse struct {
	// A list of employers matching the search criteria.
	Employers []Employer `json:"employers"`
	// A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive.
	RequestId string `json:"request_id"`
	AdditionalProperties map[string]interface{}
}

type _EmployersSearchResponse EmployersSearchResponse

// NewEmployersSearchResponse instantiates a new EmployersSearchResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEmployersSearchResponse(employers []Employer, requestId string) *EmployersSearchResponse {
	this := EmployersSearchResponse{}
	this.Employers = employers
	this.RequestId = requestId
	return &this
}

// NewEmployersSearchResponseWithDefaults instantiates a new EmployersSearchResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEmployersSearchResponseWithDefaults() *EmployersSearchResponse {
	this := EmployersSearchResponse{}
	return &this
}

// GetEmployers returns the Employers field value
func (o *EmployersSearchResponse) GetEmployers() []Employer {
	if o == nil {
		var ret []Employer
		return ret
	}

	return o.Employers
}

// GetEmployersOk returns a tuple with the Employers field value
// and a boolean to check if the value has been set.
func (o *EmployersSearchResponse) GetEmployersOk() (*[]Employer, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Employers, true
}

// SetEmployers sets field value
func (o *EmployersSearchResponse) SetEmployers(v []Employer) {
	o.Employers = v
}

// GetRequestId returns the RequestId field value
func (o *EmployersSearchResponse) GetRequestId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value
// and a boolean to check if the value has been set.
func (o *EmployersSearchResponse) GetRequestIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RequestId, true
}

// SetRequestId sets field value
func (o *EmployersSearchResponse) SetRequestId(v string) {
	o.RequestId = v
}

func (o EmployersSearchResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["employers"] = o.Employers
	}
	if true {
		toSerialize["request_id"] = o.RequestId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *EmployersSearchResponse) UnmarshalJSON(bytes []byte) (err error) {
	varEmployersSearchResponse := _EmployersSearchResponse{}

	if err = json.Unmarshal(bytes, &varEmployersSearchResponse); err == nil {
		*o = EmployersSearchResponse(varEmployersSearchResponse)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "employers")
		delete(additionalProperties, "request_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableEmployersSearchResponse struct {
	value *EmployersSearchResponse
	isSet bool
}

func (v NullableEmployersSearchResponse) Get() *EmployersSearchResponse {
	return v.value
}

func (v *NullableEmployersSearchResponse) Set(val *EmployersSearchResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableEmployersSearchResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableEmployersSearchResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEmployersSearchResponse(val *EmployersSearchResponse) *NullableEmployersSearchResponse {
	return &NullableEmployersSearchResponse{value: val, isSet: true}
}

func (v NullableEmployersSearchResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEmployersSearchResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


