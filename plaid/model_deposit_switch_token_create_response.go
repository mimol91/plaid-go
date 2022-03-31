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

// DepositSwitchTokenCreateResponse DepositSwitchTokenCreateResponse defines the response schema for `/deposit_switch/token/create`
type DepositSwitchTokenCreateResponse struct {
	// Deposit switch token, used to initialize Link for the Deposit Switch product
	DepositSwitchToken string `json:"deposit_switch_token"`
	// Expiration time of the token, in [ISO 8601](https://wikipedia.org/wiki/ISO_8601) format
	DepositSwitchTokenExpirationTime string `json:"deposit_switch_token_expiration_time"`
	// A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive.
	RequestId string `json:"request_id"`
	AdditionalProperties map[string]interface{}
}

type _DepositSwitchTokenCreateResponse DepositSwitchTokenCreateResponse

// NewDepositSwitchTokenCreateResponse instantiates a new DepositSwitchTokenCreateResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDepositSwitchTokenCreateResponse(depositSwitchToken string, depositSwitchTokenExpirationTime string, requestId string) *DepositSwitchTokenCreateResponse {
	this := DepositSwitchTokenCreateResponse{}
	this.DepositSwitchToken = depositSwitchToken
	this.DepositSwitchTokenExpirationTime = depositSwitchTokenExpirationTime
	this.RequestId = requestId
	return &this
}

// NewDepositSwitchTokenCreateResponseWithDefaults instantiates a new DepositSwitchTokenCreateResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDepositSwitchTokenCreateResponseWithDefaults() *DepositSwitchTokenCreateResponse {
	this := DepositSwitchTokenCreateResponse{}
	return &this
}

// GetDepositSwitchToken returns the DepositSwitchToken field value
func (o *DepositSwitchTokenCreateResponse) GetDepositSwitchToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DepositSwitchToken
}

// GetDepositSwitchTokenOk returns a tuple with the DepositSwitchToken field value
// and a boolean to check if the value has been set.
func (o *DepositSwitchTokenCreateResponse) GetDepositSwitchTokenOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.DepositSwitchToken, true
}

// SetDepositSwitchToken sets field value
func (o *DepositSwitchTokenCreateResponse) SetDepositSwitchToken(v string) {
	o.DepositSwitchToken = v
}

// GetDepositSwitchTokenExpirationTime returns the DepositSwitchTokenExpirationTime field value
func (o *DepositSwitchTokenCreateResponse) GetDepositSwitchTokenExpirationTime() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DepositSwitchTokenExpirationTime
}

// GetDepositSwitchTokenExpirationTimeOk returns a tuple with the DepositSwitchTokenExpirationTime field value
// and a boolean to check if the value has been set.
func (o *DepositSwitchTokenCreateResponse) GetDepositSwitchTokenExpirationTimeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.DepositSwitchTokenExpirationTime, true
}

// SetDepositSwitchTokenExpirationTime sets field value
func (o *DepositSwitchTokenCreateResponse) SetDepositSwitchTokenExpirationTime(v string) {
	o.DepositSwitchTokenExpirationTime = v
}

// GetRequestId returns the RequestId field value
func (o *DepositSwitchTokenCreateResponse) GetRequestId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value
// and a boolean to check if the value has been set.
func (o *DepositSwitchTokenCreateResponse) GetRequestIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RequestId, true
}

// SetRequestId sets field value
func (o *DepositSwitchTokenCreateResponse) SetRequestId(v string) {
	o.RequestId = v
}

func (o DepositSwitchTokenCreateResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["deposit_switch_token"] = o.DepositSwitchToken
	}
	if true {
		toSerialize["deposit_switch_token_expiration_time"] = o.DepositSwitchTokenExpirationTime
	}
	if true {
		toSerialize["request_id"] = o.RequestId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *DepositSwitchTokenCreateResponse) UnmarshalJSON(bytes []byte) (err error) {
	varDepositSwitchTokenCreateResponse := _DepositSwitchTokenCreateResponse{}

	if err = json.Unmarshal(bytes, &varDepositSwitchTokenCreateResponse); err == nil {
		*o = DepositSwitchTokenCreateResponse(varDepositSwitchTokenCreateResponse)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "deposit_switch_token")
		delete(additionalProperties, "deposit_switch_token_expiration_time")
		delete(additionalProperties, "request_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDepositSwitchTokenCreateResponse struct {
	value *DepositSwitchTokenCreateResponse
	isSet bool
}

func (v NullableDepositSwitchTokenCreateResponse) Get() *DepositSwitchTokenCreateResponse {
	return v.value
}

func (v *NullableDepositSwitchTokenCreateResponse) Set(val *DepositSwitchTokenCreateResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableDepositSwitchTokenCreateResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableDepositSwitchTokenCreateResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDepositSwitchTokenCreateResponse(val *DepositSwitchTokenCreateResponse) *NullableDepositSwitchTokenCreateResponse {
	return &NullableDepositSwitchTokenCreateResponse{value: val, isSet: true}
}

func (v NullableDepositSwitchTokenCreateResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDepositSwitchTokenCreateResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


