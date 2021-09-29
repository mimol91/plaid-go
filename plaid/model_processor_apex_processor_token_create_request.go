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

// ProcessorApexProcessorTokenCreateRequest ProcessorApexProcessorTokenCreateRequest defines the request schema for `/processor/apex/processor_token/create`
type ProcessorApexProcessorTokenCreateRequest struct {
	// Your Plaid API `client_id`. The `client_id` is required and may be provided either in the `PLAID-CLIENT-ID` header or as part of a request body.
	ClientId *string `json:"client_id,omitempty"`
	// Your Plaid API `secret`. The `secret` is required and may be provided either in the `PLAID-SECRET` header or as part of a request body.
	Secret *string `json:"secret,omitempty"`
	// The access token associated with the Item data is being requested for.
	AccessToken string `json:"access_token"`
	// The `account_id` value obtained from the `onSuccess` callback in Link
	AccountId string `json:"account_id"`
}

// NewProcessorApexProcessorTokenCreateRequest instantiates a new ProcessorApexProcessorTokenCreateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProcessorApexProcessorTokenCreateRequest(accessToken string, accountId string) *ProcessorApexProcessorTokenCreateRequest {
	this := ProcessorApexProcessorTokenCreateRequest{}
	this.AccessToken = accessToken
	this.AccountId = accountId
	return &this
}

// NewProcessorApexProcessorTokenCreateRequestWithDefaults instantiates a new ProcessorApexProcessorTokenCreateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProcessorApexProcessorTokenCreateRequestWithDefaults() *ProcessorApexProcessorTokenCreateRequest {
	this := ProcessorApexProcessorTokenCreateRequest{}
	return &this
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *ProcessorApexProcessorTokenCreateRequest) GetClientId() string {
	if o == nil || o.ClientId == nil {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProcessorApexProcessorTokenCreateRequest) GetClientIdOk() (*string, bool) {
	if o == nil || o.ClientId == nil {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *ProcessorApexProcessorTokenCreateRequest) HasClientId() bool {
	if o != nil && o.ClientId != nil {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *ProcessorApexProcessorTokenCreateRequest) SetClientId(v string) {
	o.ClientId = &v
}

// GetSecret returns the Secret field value if set, zero value otherwise.
func (o *ProcessorApexProcessorTokenCreateRequest) GetSecret() string {
	if o == nil || o.Secret == nil {
		var ret string
		return ret
	}
	return *o.Secret
}

// GetSecretOk returns a tuple with the Secret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProcessorApexProcessorTokenCreateRequest) GetSecretOk() (*string, bool) {
	if o == nil || o.Secret == nil {
		return nil, false
	}
	return o.Secret, true
}

// HasSecret returns a boolean if a field has been set.
func (o *ProcessorApexProcessorTokenCreateRequest) HasSecret() bool {
	if o != nil && o.Secret != nil {
		return true
	}

	return false
}

// SetSecret gets a reference to the given string and assigns it to the Secret field.
func (o *ProcessorApexProcessorTokenCreateRequest) SetSecret(v string) {
	o.Secret = &v
}

// GetAccessToken returns the AccessToken field value
func (o *ProcessorApexProcessorTokenCreateRequest) GetAccessToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccessToken
}

// GetAccessTokenOk returns a tuple with the AccessToken field value
// and a boolean to check if the value has been set.
func (o *ProcessorApexProcessorTokenCreateRequest) GetAccessTokenOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AccessToken, true
}

// SetAccessToken sets field value
func (o *ProcessorApexProcessorTokenCreateRequest) SetAccessToken(v string) {
	o.AccessToken = v
}

// GetAccountId returns the AccountId field value
func (o *ProcessorApexProcessorTokenCreateRequest) GetAccountId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value
// and a boolean to check if the value has been set.
func (o *ProcessorApexProcessorTokenCreateRequest) GetAccountIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AccountId, true
}

// SetAccountId sets field value
func (o *ProcessorApexProcessorTokenCreateRequest) SetAccountId(v string) {
	o.AccountId = v
}

func (o ProcessorApexProcessorTokenCreateRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ClientId != nil {
		toSerialize["client_id"] = o.ClientId
	}
	if o.Secret != nil {
		toSerialize["secret"] = o.Secret
	}
	if true {
		toSerialize["access_token"] = o.AccessToken
	}
	if true {
		toSerialize["account_id"] = o.AccountId
	}
	return json.Marshal(toSerialize)
}

type NullableProcessorApexProcessorTokenCreateRequest struct {
	value *ProcessorApexProcessorTokenCreateRequest
	isSet bool
}

func (v NullableProcessorApexProcessorTokenCreateRequest) Get() *ProcessorApexProcessorTokenCreateRequest {
	return v.value
}

func (v *NullableProcessorApexProcessorTokenCreateRequest) Set(val *ProcessorApexProcessorTokenCreateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableProcessorApexProcessorTokenCreateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableProcessorApexProcessorTokenCreateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProcessorApexProcessorTokenCreateRequest(val *ProcessorApexProcessorTokenCreateRequest) *NullableProcessorApexProcessorTokenCreateRequest {
	return &NullableProcessorApexProcessorTokenCreateRequest{value: val, isSet: true}
}

func (v NullableProcessorApexProcessorTokenCreateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProcessorApexProcessorTokenCreateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


