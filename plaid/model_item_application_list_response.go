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

// ItemApplicationListResponse Describes the connected application for a particular end user.
type ItemApplicationListResponse struct {
	// A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive.
	RequestId *string `json:"request_id,omitempty"`
	// A list of connected applications.
	Applications []ConnectedApplication `json:"applications"`
}

// NewItemApplicationListResponse instantiates a new ItemApplicationListResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewItemApplicationListResponse(applications []ConnectedApplication) *ItemApplicationListResponse {
	this := ItemApplicationListResponse{}
	this.Applications = applications
	return &this
}

// NewItemApplicationListResponseWithDefaults instantiates a new ItemApplicationListResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewItemApplicationListResponseWithDefaults() *ItemApplicationListResponse {
	this := ItemApplicationListResponse{}
	return &this
}

// GetRequestId returns the RequestId field value if set, zero value otherwise.
func (o *ItemApplicationListResponse) GetRequestId() string {
	if o == nil || o.RequestId == nil {
		var ret string
		return ret
	}
	return *o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ItemApplicationListResponse) GetRequestIdOk() (*string, bool) {
	if o == nil || o.RequestId == nil {
		return nil, false
	}
	return o.RequestId, true
}

// HasRequestId returns a boolean if a field has been set.
func (o *ItemApplicationListResponse) HasRequestId() bool {
	if o != nil && o.RequestId != nil {
		return true
	}

	return false
}

// SetRequestId gets a reference to the given string and assigns it to the RequestId field.
func (o *ItemApplicationListResponse) SetRequestId(v string) {
	o.RequestId = &v
}

// GetApplications returns the Applications field value
func (o *ItemApplicationListResponse) GetApplications() []ConnectedApplication {
	if o == nil {
		var ret []ConnectedApplication
		return ret
	}

	return o.Applications
}

// GetApplicationsOk returns a tuple with the Applications field value
// and a boolean to check if the value has been set.
func (o *ItemApplicationListResponse) GetApplicationsOk() (*[]ConnectedApplication, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Applications, true
}

// SetApplications sets field value
func (o *ItemApplicationListResponse) SetApplications(v []ConnectedApplication) {
	o.Applications = v
}

func (o ItemApplicationListResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.RequestId != nil {
		toSerialize["request_id"] = o.RequestId
	}
	if true {
		toSerialize["applications"] = o.Applications
	}
	return json.Marshal(toSerialize)
}

type NullableItemApplicationListResponse struct {
	value *ItemApplicationListResponse
	isSet bool
}

func (v NullableItemApplicationListResponse) Get() *ItemApplicationListResponse {
	return v.value
}

func (v *NullableItemApplicationListResponse) Set(val *ItemApplicationListResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableItemApplicationListResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableItemApplicationListResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableItemApplicationListResponse(val *ItemApplicationListResponse) *NullableItemApplicationListResponse {
	return &NullableItemApplicationListResponse{value: val, isSet: true}
}

func (v NullableItemApplicationListResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableItemApplicationListResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


