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

// LiabilitiesGetResponse LiabilitiesGetResponse defines the response schema for `/liabilities/get`
type LiabilitiesGetResponse struct {
	// An array of accounts associated with the Item
	Accounts []AccountBase `json:"accounts"`
	Item Item `json:"item"`
	Liabilities LiabilitiesObject `json:"liabilities"`
	// A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive.
	RequestId string `json:"request_id"`
	AdditionalProperties map[string]interface{}
}

type _LiabilitiesGetResponse LiabilitiesGetResponse

// NewLiabilitiesGetResponse instantiates a new LiabilitiesGetResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLiabilitiesGetResponse(accounts []AccountBase, item Item, liabilities LiabilitiesObject, requestId string) *LiabilitiesGetResponse {
	this := LiabilitiesGetResponse{}
	this.Accounts = accounts
	this.Item = item
	this.Liabilities = liabilities
	this.RequestId = requestId
	return &this
}

// NewLiabilitiesGetResponseWithDefaults instantiates a new LiabilitiesGetResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLiabilitiesGetResponseWithDefaults() *LiabilitiesGetResponse {
	this := LiabilitiesGetResponse{}
	return &this
}

// GetAccounts returns the Accounts field value
func (o *LiabilitiesGetResponse) GetAccounts() []AccountBase {
	if o == nil {
		var ret []AccountBase
		return ret
	}

	return o.Accounts
}

// GetAccountsOk returns a tuple with the Accounts field value
// and a boolean to check if the value has been set.
func (o *LiabilitiesGetResponse) GetAccountsOk() (*[]AccountBase, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Accounts, true
}

// SetAccounts sets field value
func (o *LiabilitiesGetResponse) SetAccounts(v []AccountBase) {
	o.Accounts = v
}

// GetItem returns the Item field value
func (o *LiabilitiesGetResponse) GetItem() Item {
	if o == nil {
		var ret Item
		return ret
	}

	return o.Item
}

// GetItemOk returns a tuple with the Item field value
// and a boolean to check if the value has been set.
func (o *LiabilitiesGetResponse) GetItemOk() (*Item, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Item, true
}

// SetItem sets field value
func (o *LiabilitiesGetResponse) SetItem(v Item) {
	o.Item = v
}

// GetLiabilities returns the Liabilities field value
func (o *LiabilitiesGetResponse) GetLiabilities() LiabilitiesObject {
	if o == nil {
		var ret LiabilitiesObject
		return ret
	}

	return o.Liabilities
}

// GetLiabilitiesOk returns a tuple with the Liabilities field value
// and a boolean to check if the value has been set.
func (o *LiabilitiesGetResponse) GetLiabilitiesOk() (*LiabilitiesObject, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Liabilities, true
}

// SetLiabilities sets field value
func (o *LiabilitiesGetResponse) SetLiabilities(v LiabilitiesObject) {
	o.Liabilities = v
}

// GetRequestId returns the RequestId field value
func (o *LiabilitiesGetResponse) GetRequestId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value
// and a boolean to check if the value has been set.
func (o *LiabilitiesGetResponse) GetRequestIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RequestId, true
}

// SetRequestId sets field value
func (o *LiabilitiesGetResponse) SetRequestId(v string) {
	o.RequestId = v
}

func (o LiabilitiesGetResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["accounts"] = o.Accounts
	}
	if true {
		toSerialize["item"] = o.Item
	}
	if true {
		toSerialize["liabilities"] = o.Liabilities
	}
	if true {
		toSerialize["request_id"] = o.RequestId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *LiabilitiesGetResponse) UnmarshalJSON(bytes []byte) (err error) {
	varLiabilitiesGetResponse := _LiabilitiesGetResponse{}

	if err = json.Unmarshal(bytes, &varLiabilitiesGetResponse); err == nil {
		*o = LiabilitiesGetResponse(varLiabilitiesGetResponse)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "accounts")
		delete(additionalProperties, "item")
		delete(additionalProperties, "liabilities")
		delete(additionalProperties, "request_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableLiabilitiesGetResponse struct {
	value *LiabilitiesGetResponse
	isSet bool
}

func (v NullableLiabilitiesGetResponse) Get() *LiabilitiesGetResponse {
	return v.value
}

func (v *NullableLiabilitiesGetResponse) Set(val *LiabilitiesGetResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableLiabilitiesGetResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableLiabilitiesGetResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLiabilitiesGetResponse(val *LiabilitiesGetResponse) *NullableLiabilitiesGetResponse {
	return &NullableLiabilitiesGetResponse{value: val, isSet: true}
}

func (v NullableLiabilitiesGetResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLiabilitiesGetResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


