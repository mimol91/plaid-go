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
	"time"
)

// ProcessorBalanceGetRequestOptions An optional object to filter `/processor/balance/get` results.
type ProcessorBalanceGetRequestOptions struct {
	// Timestamp in [ISO 8601](https://wikipedia.org/wiki/ISO_8601) format (`YYYY-MM-DDTHH:mm:ssZ`) indicating the oldest acceptable balance when making a request to `/accounts/balance/get`.  If the balance that is pulled for `ins_128026` (Capital One) is older than the given timestamp, an `INVALID_REQUEST` error with the code of `LAST_UPDATED_DATETIME_OUT_OF_RANGE` will be returned with the most recent timestamp for the requested account contained in the response.  This field is only used when the institution is `ins_128026` (Capital One), in which case a value must be provided or an `INVALID_REQUEST` error with the code of `INVALID_FIELD` will be returned. For all other institutions, this field is ignored.
	MinLastUpdatedDatetime *time.Time `json:"min_last_updated_datetime,omitempty"`
}

// NewProcessorBalanceGetRequestOptions instantiates a new ProcessorBalanceGetRequestOptions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProcessorBalanceGetRequestOptions() *ProcessorBalanceGetRequestOptions {
	this := ProcessorBalanceGetRequestOptions{}
	return &this
}

// NewProcessorBalanceGetRequestOptionsWithDefaults instantiates a new ProcessorBalanceGetRequestOptions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProcessorBalanceGetRequestOptionsWithDefaults() *ProcessorBalanceGetRequestOptions {
	this := ProcessorBalanceGetRequestOptions{}
	return &this
}

// GetMinLastUpdatedDatetime returns the MinLastUpdatedDatetime field value if set, zero value otherwise.
func (o *ProcessorBalanceGetRequestOptions) GetMinLastUpdatedDatetime() time.Time {
	if o == nil || o.MinLastUpdatedDatetime == nil {
		var ret time.Time
		return ret
	}
	return *o.MinLastUpdatedDatetime
}

// GetMinLastUpdatedDatetimeOk returns a tuple with the MinLastUpdatedDatetime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProcessorBalanceGetRequestOptions) GetMinLastUpdatedDatetimeOk() (*time.Time, bool) {
	if o == nil || o.MinLastUpdatedDatetime == nil {
		return nil, false
	}
	return o.MinLastUpdatedDatetime, true
}

// HasMinLastUpdatedDatetime returns a boolean if a field has been set.
func (o *ProcessorBalanceGetRequestOptions) HasMinLastUpdatedDatetime() bool {
	if o != nil && o.MinLastUpdatedDatetime != nil {
		return true
	}

	return false
}

// SetMinLastUpdatedDatetime gets a reference to the given time.Time and assigns it to the MinLastUpdatedDatetime field.
func (o *ProcessorBalanceGetRequestOptions) SetMinLastUpdatedDatetime(v time.Time) {
	o.MinLastUpdatedDatetime = &v
}

func (o ProcessorBalanceGetRequestOptions) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.MinLastUpdatedDatetime != nil {
		toSerialize["min_last_updated_datetime"] = o.MinLastUpdatedDatetime
	}
	return json.Marshal(toSerialize)
}

type NullableProcessorBalanceGetRequestOptions struct {
	value *ProcessorBalanceGetRequestOptions
	isSet bool
}

func (v NullableProcessorBalanceGetRequestOptions) Get() *ProcessorBalanceGetRequestOptions {
	return v.value
}

func (v *NullableProcessorBalanceGetRequestOptions) Set(val *ProcessorBalanceGetRequestOptions) {
	v.value = val
	v.isSet = true
}

func (v NullableProcessorBalanceGetRequestOptions) IsSet() bool {
	return v.isSet
}

func (v *NullableProcessorBalanceGetRequestOptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProcessorBalanceGetRequestOptions(val *ProcessorBalanceGetRequestOptions) *NullableProcessorBalanceGetRequestOptions {
	return &NullableProcessorBalanceGetRequestOptions{value: val, isSet: true}
}

func (v NullableProcessorBalanceGetRequestOptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProcessorBalanceGetRequestOptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


