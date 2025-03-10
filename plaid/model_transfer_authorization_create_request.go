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

// TransferAuthorizationCreateRequest Defines the request schema for `/transfer/authorization/create`
type TransferAuthorizationCreateRequest struct {
	// Your Plaid API `client_id`. The `client_id` is required and may be provided either in the `PLAID-CLIENT-ID` header or as part of a request body.
	ClientId *string `json:"client_id,omitempty"`
	// Your Plaid API `secret`. The `secret` is required and may be provided either in the `PLAID-SECRET` header or as part of a request body.
	Secret *string `json:"secret,omitempty"`
	// The Plaid `access_token` for the account that will be debited or credited.
	AccessToken string `json:"access_token"`
	// The Plaid `account_id` for the account that will be debited or credited.
	AccountId string `json:"account_id"`
	Type TransferType `json:"type"`
	Network TransferNetwork `json:"network"`
	// The amount of the transfer (decimal string with two digits of precision e.g. \"10.00\").
	Amount string `json:"amount"`
	AchClass ACHClass `json:"ach_class"`
	User TransferUserInRequest `json:"user"`
	Device *TransferAuthorizationDevice `json:"device,omitempty"`
	// Plaid's unique identifier for the origination account for this authorization. If not specified, the default account will be used.
	OriginationAccountId *string `json:"origination_account_id,omitempty"`
	// The currency of the transfer amount. The default value is \"USD\".
	IsoCurrencyCode *string `json:"iso_currency_code,omitempty"`
}

// NewTransferAuthorizationCreateRequest instantiates a new TransferAuthorizationCreateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransferAuthorizationCreateRequest(accessToken string, accountId string, type_ TransferType, network TransferNetwork, amount string, achClass ACHClass, user TransferUserInRequest) *TransferAuthorizationCreateRequest {
	this := TransferAuthorizationCreateRequest{}
	this.AccessToken = accessToken
	this.AccountId = accountId
	this.Type = type_
	this.Network = network
	this.Amount = amount
	this.AchClass = achClass
	this.User = user
	return &this
}

// NewTransferAuthorizationCreateRequestWithDefaults instantiates a new TransferAuthorizationCreateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransferAuthorizationCreateRequestWithDefaults() *TransferAuthorizationCreateRequest {
	this := TransferAuthorizationCreateRequest{}
	return &this
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *TransferAuthorizationCreateRequest) GetClientId() string {
	if o == nil || o.ClientId == nil {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferAuthorizationCreateRequest) GetClientIdOk() (*string, bool) {
	if o == nil || o.ClientId == nil {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *TransferAuthorizationCreateRequest) HasClientId() bool {
	if o != nil && o.ClientId != nil {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *TransferAuthorizationCreateRequest) SetClientId(v string) {
	o.ClientId = &v
}

// GetSecret returns the Secret field value if set, zero value otherwise.
func (o *TransferAuthorizationCreateRequest) GetSecret() string {
	if o == nil || o.Secret == nil {
		var ret string
		return ret
	}
	return *o.Secret
}

// GetSecretOk returns a tuple with the Secret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferAuthorizationCreateRequest) GetSecretOk() (*string, bool) {
	if o == nil || o.Secret == nil {
		return nil, false
	}
	return o.Secret, true
}

// HasSecret returns a boolean if a field has been set.
func (o *TransferAuthorizationCreateRequest) HasSecret() bool {
	if o != nil && o.Secret != nil {
		return true
	}

	return false
}

// SetSecret gets a reference to the given string and assigns it to the Secret field.
func (o *TransferAuthorizationCreateRequest) SetSecret(v string) {
	o.Secret = &v
}

// GetAccessToken returns the AccessToken field value
func (o *TransferAuthorizationCreateRequest) GetAccessToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccessToken
}

// GetAccessTokenOk returns a tuple with the AccessToken field value
// and a boolean to check if the value has been set.
func (o *TransferAuthorizationCreateRequest) GetAccessTokenOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AccessToken, true
}

// SetAccessToken sets field value
func (o *TransferAuthorizationCreateRequest) SetAccessToken(v string) {
	o.AccessToken = v
}

// GetAccountId returns the AccountId field value
func (o *TransferAuthorizationCreateRequest) GetAccountId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value
// and a boolean to check if the value has been set.
func (o *TransferAuthorizationCreateRequest) GetAccountIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AccountId, true
}

// SetAccountId sets field value
func (o *TransferAuthorizationCreateRequest) SetAccountId(v string) {
	o.AccountId = v
}

// GetType returns the Type field value
func (o *TransferAuthorizationCreateRequest) GetType() TransferType {
	if o == nil {
		var ret TransferType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *TransferAuthorizationCreateRequest) GetTypeOk() (*TransferType, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *TransferAuthorizationCreateRequest) SetType(v TransferType) {
	o.Type = v
}

// GetNetwork returns the Network field value
func (o *TransferAuthorizationCreateRequest) GetNetwork() TransferNetwork {
	if o == nil {
		var ret TransferNetwork
		return ret
	}

	return o.Network
}

// GetNetworkOk returns a tuple with the Network field value
// and a boolean to check if the value has been set.
func (o *TransferAuthorizationCreateRequest) GetNetworkOk() (*TransferNetwork, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Network, true
}

// SetNetwork sets field value
func (o *TransferAuthorizationCreateRequest) SetNetwork(v TransferNetwork) {
	o.Network = v
}

// GetAmount returns the Amount field value
func (o *TransferAuthorizationCreateRequest) GetAmount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *TransferAuthorizationCreateRequest) GetAmountOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *TransferAuthorizationCreateRequest) SetAmount(v string) {
	o.Amount = v
}

// GetAchClass returns the AchClass field value
func (o *TransferAuthorizationCreateRequest) GetAchClass() ACHClass {
	if o == nil {
		var ret ACHClass
		return ret
	}

	return o.AchClass
}

// GetAchClassOk returns a tuple with the AchClass field value
// and a boolean to check if the value has been set.
func (o *TransferAuthorizationCreateRequest) GetAchClassOk() (*ACHClass, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AchClass, true
}

// SetAchClass sets field value
func (o *TransferAuthorizationCreateRequest) SetAchClass(v ACHClass) {
	o.AchClass = v
}

// GetUser returns the User field value
func (o *TransferAuthorizationCreateRequest) GetUser() TransferUserInRequest {
	if o == nil {
		var ret TransferUserInRequest
		return ret
	}

	return o.User
}

// GetUserOk returns a tuple with the User field value
// and a boolean to check if the value has been set.
func (o *TransferAuthorizationCreateRequest) GetUserOk() (*TransferUserInRequest, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.User, true
}

// SetUser sets field value
func (o *TransferAuthorizationCreateRequest) SetUser(v TransferUserInRequest) {
	o.User = v
}

// GetDevice returns the Device field value if set, zero value otherwise.
func (o *TransferAuthorizationCreateRequest) GetDevice() TransferAuthorizationDevice {
	if o == nil || o.Device == nil {
		var ret TransferAuthorizationDevice
		return ret
	}
	return *o.Device
}

// GetDeviceOk returns a tuple with the Device field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferAuthorizationCreateRequest) GetDeviceOk() (*TransferAuthorizationDevice, bool) {
	if o == nil || o.Device == nil {
		return nil, false
	}
	return o.Device, true
}

// HasDevice returns a boolean if a field has been set.
func (o *TransferAuthorizationCreateRequest) HasDevice() bool {
	if o != nil && o.Device != nil {
		return true
	}

	return false
}

// SetDevice gets a reference to the given TransferAuthorizationDevice and assigns it to the Device field.
func (o *TransferAuthorizationCreateRequest) SetDevice(v TransferAuthorizationDevice) {
	o.Device = &v
}

// GetOriginationAccountId returns the OriginationAccountId field value if set, zero value otherwise.
func (o *TransferAuthorizationCreateRequest) GetOriginationAccountId() string {
	if o == nil || o.OriginationAccountId == nil {
		var ret string
		return ret
	}
	return *o.OriginationAccountId
}

// GetOriginationAccountIdOk returns a tuple with the OriginationAccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferAuthorizationCreateRequest) GetOriginationAccountIdOk() (*string, bool) {
	if o == nil || o.OriginationAccountId == nil {
		return nil, false
	}
	return o.OriginationAccountId, true
}

// HasOriginationAccountId returns a boolean if a field has been set.
func (o *TransferAuthorizationCreateRequest) HasOriginationAccountId() bool {
	if o != nil && o.OriginationAccountId != nil {
		return true
	}

	return false
}

// SetOriginationAccountId gets a reference to the given string and assigns it to the OriginationAccountId field.
func (o *TransferAuthorizationCreateRequest) SetOriginationAccountId(v string) {
	o.OriginationAccountId = &v
}

// GetIsoCurrencyCode returns the IsoCurrencyCode field value if set, zero value otherwise.
func (o *TransferAuthorizationCreateRequest) GetIsoCurrencyCode() string {
	if o == nil || o.IsoCurrencyCode == nil {
		var ret string
		return ret
	}
	return *o.IsoCurrencyCode
}

// GetIsoCurrencyCodeOk returns a tuple with the IsoCurrencyCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferAuthorizationCreateRequest) GetIsoCurrencyCodeOk() (*string, bool) {
	if o == nil || o.IsoCurrencyCode == nil {
		return nil, false
	}
	return o.IsoCurrencyCode, true
}

// HasIsoCurrencyCode returns a boolean if a field has been set.
func (o *TransferAuthorizationCreateRequest) HasIsoCurrencyCode() bool {
	if o != nil && o.IsoCurrencyCode != nil {
		return true
	}

	return false
}

// SetIsoCurrencyCode gets a reference to the given string and assigns it to the IsoCurrencyCode field.
func (o *TransferAuthorizationCreateRequest) SetIsoCurrencyCode(v string) {
	o.IsoCurrencyCode = &v
}

func (o TransferAuthorizationCreateRequest) MarshalJSON() ([]byte, error) {
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
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["network"] = o.Network
	}
	if true {
		toSerialize["amount"] = o.Amount
	}
	if true {
		toSerialize["ach_class"] = o.AchClass
	}
	if true {
		toSerialize["user"] = o.User
	}
	if o.Device != nil {
		toSerialize["device"] = o.Device
	}
	if o.OriginationAccountId != nil {
		toSerialize["origination_account_id"] = o.OriginationAccountId
	}
	if o.IsoCurrencyCode != nil {
		toSerialize["iso_currency_code"] = o.IsoCurrencyCode
	}
	return json.Marshal(toSerialize)
}

type NullableTransferAuthorizationCreateRequest struct {
	value *TransferAuthorizationCreateRequest
	isSet bool
}

func (v NullableTransferAuthorizationCreateRequest) Get() *TransferAuthorizationCreateRequest {
	return v.value
}

func (v *NullableTransferAuthorizationCreateRequest) Set(val *TransferAuthorizationCreateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableTransferAuthorizationCreateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableTransferAuthorizationCreateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransferAuthorizationCreateRequest(val *TransferAuthorizationCreateRequest) *NullableTransferAuthorizationCreateRequest {
	return &NullableTransferAuthorizationCreateRequest{value: val, isSet: true}
}

func (v NullableTransferAuthorizationCreateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransferAuthorizationCreateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


