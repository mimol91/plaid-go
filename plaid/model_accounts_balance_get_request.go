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

// AccountsBalanceGetRequest AccountsBalanceGetRequest defines the request schema for `/accounts/balance/get`
type AccountsBalanceGetRequest struct {
	// The access token associated with the Item data is being requested for.
	AccessToken string `json:"access_token"`
	// Your Plaid API `secret`. The `secret` is required and may be provided either in the `PLAID-SECRET` header or as part of a request body.
	Secret *string `json:"secret,omitempty"`
	// Your Plaid API `client_id`. The `client_id` is required and may be provided either in the `PLAID-CLIENT-ID` header or as part of a request body.
	ClientId *string `json:"client_id,omitempty"`
	Options *AccountsBalanceGetRequestOptions `json:"options,omitempty"`
}

// NewAccountsBalanceGetRequest instantiates a new AccountsBalanceGetRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountsBalanceGetRequest(accessToken string) *AccountsBalanceGetRequest {
	this := AccountsBalanceGetRequest{}
	this.AccessToken = accessToken
	return &this
}

// NewAccountsBalanceGetRequestWithDefaults instantiates a new AccountsBalanceGetRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountsBalanceGetRequestWithDefaults() *AccountsBalanceGetRequest {
	this := AccountsBalanceGetRequest{}
	return &this
}

// GetAccessToken returns the AccessToken field value
func (o *AccountsBalanceGetRequest) GetAccessToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccessToken
}

// GetAccessTokenOk returns a tuple with the AccessToken field value
// and a boolean to check if the value has been set.
func (o *AccountsBalanceGetRequest) GetAccessTokenOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AccessToken, true
}

// SetAccessToken sets field value
func (o *AccountsBalanceGetRequest) SetAccessToken(v string) {
	o.AccessToken = v
}

// GetSecret returns the Secret field value if set, zero value otherwise.
func (o *AccountsBalanceGetRequest) GetSecret() string {
	if o == nil || o.Secret == nil {
		var ret string
		return ret
	}
	return *o.Secret
}

// GetSecretOk returns a tuple with the Secret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountsBalanceGetRequest) GetSecretOk() (*string, bool) {
	if o == nil || o.Secret == nil {
		return nil, false
	}
	return o.Secret, true
}

// HasSecret returns a boolean if a field has been set.
func (o *AccountsBalanceGetRequest) HasSecret() bool {
	if o != nil && o.Secret != nil {
		return true
	}

	return false
}

// SetSecret gets a reference to the given string and assigns it to the Secret field.
func (o *AccountsBalanceGetRequest) SetSecret(v string) {
	o.Secret = &v
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *AccountsBalanceGetRequest) GetClientId() string {
	if o == nil || o.ClientId == nil {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountsBalanceGetRequest) GetClientIdOk() (*string, bool) {
	if o == nil || o.ClientId == nil {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *AccountsBalanceGetRequest) HasClientId() bool {
	if o != nil && o.ClientId != nil {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *AccountsBalanceGetRequest) SetClientId(v string) {
	o.ClientId = &v
}

// GetOptions returns the Options field value if set, zero value otherwise.
func (o *AccountsBalanceGetRequest) GetOptions() AccountsBalanceGetRequestOptions {
	if o == nil || o.Options == nil {
		var ret AccountsBalanceGetRequestOptions
		return ret
	}
	return *o.Options
}

// GetOptionsOk returns a tuple with the Options field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountsBalanceGetRequest) GetOptionsOk() (*AccountsBalanceGetRequestOptions, bool) {
	if o == nil || o.Options == nil {
		return nil, false
	}
	return o.Options, true
}

// HasOptions returns a boolean if a field has been set.
func (o *AccountsBalanceGetRequest) HasOptions() bool {
	if o != nil && o.Options != nil {
		return true
	}

	return false
}

// SetOptions gets a reference to the given AccountsBalanceGetRequestOptions and assigns it to the Options field.
func (o *AccountsBalanceGetRequest) SetOptions(v AccountsBalanceGetRequestOptions) {
	o.Options = &v
}

func (o AccountsBalanceGetRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["access_token"] = o.AccessToken
	}
	if o.Secret != nil {
		toSerialize["secret"] = o.Secret
	}
	if o.ClientId != nil {
		toSerialize["client_id"] = o.ClientId
	}
	if o.Options != nil {
		toSerialize["options"] = o.Options
	}
	return json.Marshal(toSerialize)
}

type NullableAccountsBalanceGetRequest struct {
	value *AccountsBalanceGetRequest
	isSet bool
}

func (v NullableAccountsBalanceGetRequest) Get() *AccountsBalanceGetRequest {
	return v.value
}

func (v *NullableAccountsBalanceGetRequest) Set(val *AccountsBalanceGetRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountsBalanceGetRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountsBalanceGetRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountsBalanceGetRequest(val *AccountsBalanceGetRequest) *NullableAccountsBalanceGetRequest {
	return &NullableAccountsBalanceGetRequest{value: val, isSet: true}
}

func (v NullableAccountsBalanceGetRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountsBalanceGetRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


