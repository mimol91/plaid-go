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

// ItemPublicTokenExchangeRequest ItemPublicTokenExchangeRequest defines the request schema for `/item/public_token/exchange`
type ItemPublicTokenExchangeRequest struct {
	// Your Plaid API `client_id`. The `client_id` is required and may be provided either in the `PLAID-CLIENT-ID` header or as part of a request body.
	ClientId *string `json:"client_id,omitempty"`
	// Your Plaid API `secret`. The `secret` is required and may be provided either in the `PLAID-SECRET` header or as part of a request body.
	Secret *string `json:"secret,omitempty"`
	// Your `public_token`, obtained from the Link `onSuccess` callback or `/sandbox/item/public_token/create`.
	PublicToken string `json:"public_token"`
}

// NewItemPublicTokenExchangeRequest instantiates a new ItemPublicTokenExchangeRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewItemPublicTokenExchangeRequest(publicToken string) *ItemPublicTokenExchangeRequest {
	this := ItemPublicTokenExchangeRequest{}
	this.PublicToken = publicToken
	return &this
}

// NewItemPublicTokenExchangeRequestWithDefaults instantiates a new ItemPublicTokenExchangeRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewItemPublicTokenExchangeRequestWithDefaults() *ItemPublicTokenExchangeRequest {
	this := ItemPublicTokenExchangeRequest{}
	return &this
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *ItemPublicTokenExchangeRequest) GetClientId() string {
	if o == nil || o.ClientId == nil {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ItemPublicTokenExchangeRequest) GetClientIdOk() (*string, bool) {
	if o == nil || o.ClientId == nil {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *ItemPublicTokenExchangeRequest) HasClientId() bool {
	if o != nil && o.ClientId != nil {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *ItemPublicTokenExchangeRequest) SetClientId(v string) {
	o.ClientId = &v
}

// GetSecret returns the Secret field value if set, zero value otherwise.
func (o *ItemPublicTokenExchangeRequest) GetSecret() string {
	if o == nil || o.Secret == nil {
		var ret string
		return ret
	}
	return *o.Secret
}

// GetSecretOk returns a tuple with the Secret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ItemPublicTokenExchangeRequest) GetSecretOk() (*string, bool) {
	if o == nil || o.Secret == nil {
		return nil, false
	}
	return o.Secret, true
}

// HasSecret returns a boolean if a field has been set.
func (o *ItemPublicTokenExchangeRequest) HasSecret() bool {
	if o != nil && o.Secret != nil {
		return true
	}

	return false
}

// SetSecret gets a reference to the given string and assigns it to the Secret field.
func (o *ItemPublicTokenExchangeRequest) SetSecret(v string) {
	o.Secret = &v
}

// GetPublicToken returns the PublicToken field value
func (o *ItemPublicTokenExchangeRequest) GetPublicToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PublicToken
}

// GetPublicTokenOk returns a tuple with the PublicToken field value
// and a boolean to check if the value has been set.
func (o *ItemPublicTokenExchangeRequest) GetPublicTokenOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.PublicToken, true
}

// SetPublicToken sets field value
func (o *ItemPublicTokenExchangeRequest) SetPublicToken(v string) {
	o.PublicToken = v
}

func (o ItemPublicTokenExchangeRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ClientId != nil {
		toSerialize["client_id"] = o.ClientId
	}
	if o.Secret != nil {
		toSerialize["secret"] = o.Secret
	}
	if true {
		toSerialize["public_token"] = o.PublicToken
	}
	return json.Marshal(toSerialize)
}

type NullableItemPublicTokenExchangeRequest struct {
	value *ItemPublicTokenExchangeRequest
	isSet bool
}

func (v NullableItemPublicTokenExchangeRequest) Get() *ItemPublicTokenExchangeRequest {
	return v.value
}

func (v *NullableItemPublicTokenExchangeRequest) Set(val *ItemPublicTokenExchangeRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableItemPublicTokenExchangeRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableItemPublicTokenExchangeRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableItemPublicTokenExchangeRequest(val *ItemPublicTokenExchangeRequest) *NullableItemPublicTokenExchangeRequest {
	return &NullableItemPublicTokenExchangeRequest{value: val, isSet: true}
}

func (v NullableItemPublicTokenExchangeRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableItemPublicTokenExchangeRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


