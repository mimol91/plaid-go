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

// UserCreateResponse UserCreateResponse defines the response schema for `/user/create`
type UserCreateResponse struct {
	// The user token associated with the User data is being requested for.
	UserToken string `json:"user_token"`
	UserId string `json:"user_id"`
	// A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive.
	RequestId string `json:"request_id"`
	AdditionalProperties map[string]interface{}
}

type _UserCreateResponse UserCreateResponse

// NewUserCreateResponse instantiates a new UserCreateResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserCreateResponse(userToken string, userId string, requestId string) *UserCreateResponse {
	this := UserCreateResponse{}
	this.UserToken = userToken
	this.UserId = userId
	this.RequestId = requestId
	return &this
}

// NewUserCreateResponseWithDefaults instantiates a new UserCreateResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserCreateResponseWithDefaults() *UserCreateResponse {
	this := UserCreateResponse{}
	return &this
}

// GetUserToken returns the UserToken field value
func (o *UserCreateResponse) GetUserToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UserToken
}

// GetUserTokenOk returns a tuple with the UserToken field value
// and a boolean to check if the value has been set.
func (o *UserCreateResponse) GetUserTokenOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.UserToken, true
}

// SetUserToken sets field value
func (o *UserCreateResponse) SetUserToken(v string) {
	o.UserToken = v
}

// GetUserId returns the UserId field value
func (o *UserCreateResponse) GetUserId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value
// and a boolean to check if the value has been set.
func (o *UserCreateResponse) GetUserIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.UserId, true
}

// SetUserId sets field value
func (o *UserCreateResponse) SetUserId(v string) {
	o.UserId = v
}

// GetRequestId returns the RequestId field value
func (o *UserCreateResponse) GetRequestId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value
// and a boolean to check if the value has been set.
func (o *UserCreateResponse) GetRequestIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RequestId, true
}

// SetRequestId sets field value
func (o *UserCreateResponse) SetRequestId(v string) {
	o.RequestId = v
}

func (o UserCreateResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["user_token"] = o.UserToken
	}
	if true {
		toSerialize["user_id"] = o.UserId
	}
	if true {
		toSerialize["request_id"] = o.RequestId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *UserCreateResponse) UnmarshalJSON(bytes []byte) (err error) {
	varUserCreateResponse := _UserCreateResponse{}

	if err = json.Unmarshal(bytes, &varUserCreateResponse); err == nil {
		*o = UserCreateResponse(varUserCreateResponse)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "user_token")
		delete(additionalProperties, "user_id")
		delete(additionalProperties, "request_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUserCreateResponse struct {
	value *UserCreateResponse
	isSet bool
}

func (v NullableUserCreateResponse) Get() *UserCreateResponse {
	return v.value
}

func (v *NullableUserCreateResponse) Set(val *UserCreateResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableUserCreateResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableUserCreateResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserCreateResponse(val *UserCreateResponse) *NullableUserCreateResponse {
	return &NullableUserCreateResponse{value: val, isSet: true}
}

func (v NullableUserCreateResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserCreateResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


