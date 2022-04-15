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
)

// LiabilitiesGetRequest LiabilitiesGetRequest defines the request schema for `/liabilities/get`
type LiabilitiesGetRequest struct {
	// Your Plaid API `client_id`. The `client_id` is required and may be provided either in the `PLAID-CLIENT-ID` header or as part of a request body.
	ClientId *string `json:"client_id,omitempty"`
	// Your Plaid API `secret`. The `secret` is required and may be provided either in the `PLAID-SECRET` header or as part of a request body.
	Secret *string `json:"secret,omitempty"`
	// The access token associated with the Item data is being requested for.
	AccessToken string `json:"access_token"`
	Options *LiabilitiesGetRequestOptions `json:"options,omitempty"`
}

// NewLiabilitiesGetRequest instantiates a new LiabilitiesGetRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLiabilitiesGetRequest(accessToken string) *LiabilitiesGetRequest {
	this := LiabilitiesGetRequest{}
	this.AccessToken = accessToken
	return &this
}

// NewLiabilitiesGetRequestWithDefaults instantiates a new LiabilitiesGetRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLiabilitiesGetRequestWithDefaults() *LiabilitiesGetRequest {
	this := LiabilitiesGetRequest{}
	return &this
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *LiabilitiesGetRequest) GetClientId() string {
	if o == nil || o.ClientId == nil {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LiabilitiesGetRequest) GetClientIdOk() (*string, bool) {
	if o == nil || o.ClientId == nil {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *LiabilitiesGetRequest) HasClientId() bool {
	if o != nil && o.ClientId != nil {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *LiabilitiesGetRequest) SetClientId(v string) {
	o.ClientId = &v
}

// GetSecret returns the Secret field value if set, zero value otherwise.
func (o *LiabilitiesGetRequest) GetSecret() string {
	if o == nil || o.Secret == nil {
		var ret string
		return ret
	}
	return *o.Secret
}

// GetSecretOk returns a tuple with the Secret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LiabilitiesGetRequest) GetSecretOk() (*string, bool) {
	if o == nil || o.Secret == nil {
		return nil, false
	}
	return o.Secret, true
}

// HasSecret returns a boolean if a field has been set.
func (o *LiabilitiesGetRequest) HasSecret() bool {
	if o != nil && o.Secret != nil {
		return true
	}

	return false
}

// SetSecret gets a reference to the given string and assigns it to the Secret field.
func (o *LiabilitiesGetRequest) SetSecret(v string) {
	o.Secret = &v
}

// GetAccessToken returns the AccessToken field value
func (o *LiabilitiesGetRequest) GetAccessToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccessToken
}

// GetAccessTokenOk returns a tuple with the AccessToken field value
// and a boolean to check if the value has been set.
func (o *LiabilitiesGetRequest) GetAccessTokenOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AccessToken, true
}

// SetAccessToken sets field value
func (o *LiabilitiesGetRequest) SetAccessToken(v string) {
	o.AccessToken = v
}

// GetOptions returns the Options field value if set, zero value otherwise.
func (o *LiabilitiesGetRequest) GetOptions() LiabilitiesGetRequestOptions {
	if o == nil || o.Options == nil {
		var ret LiabilitiesGetRequestOptions
		return ret
	}
	return *o.Options
}

// GetOptionsOk returns a tuple with the Options field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LiabilitiesGetRequest) GetOptionsOk() (*LiabilitiesGetRequestOptions, bool) {
	if o == nil || o.Options == nil {
		return nil, false
	}
	return o.Options, true
}

// HasOptions returns a boolean if a field has been set.
func (o *LiabilitiesGetRequest) HasOptions() bool {
	if o != nil && o.Options != nil {
		return true
	}

	return false
}

// SetOptions gets a reference to the given LiabilitiesGetRequestOptions and assigns it to the Options field.
func (o *LiabilitiesGetRequest) SetOptions(v LiabilitiesGetRequestOptions) {
	o.Options = &v
}

func (o LiabilitiesGetRequest) MarshalJSON() ([]byte, error) {
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
	if o.Options != nil {
		toSerialize["options"] = o.Options
	}
	return json.Marshal(toSerialize)
}

type NullableLiabilitiesGetRequest struct {
	value *LiabilitiesGetRequest
	isSet bool
}

func (v NullableLiabilitiesGetRequest) Get() *LiabilitiesGetRequest {
	return v.value
}

func (v *NullableLiabilitiesGetRequest) Set(val *LiabilitiesGetRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableLiabilitiesGetRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableLiabilitiesGetRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLiabilitiesGetRequest(val *LiabilitiesGetRequest) *NullableLiabilitiesGetRequest {
	return &NullableLiabilitiesGetRequest{value: val, isSet: true}
}

func (v NullableLiabilitiesGetRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLiabilitiesGetRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


