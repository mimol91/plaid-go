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

// PaymentInitiationPaymentTokenCreateResponse PaymentInitiationPaymentTokenCreateResponse defines the response schema for `/payment_initiation/payment/token/create`
type PaymentInitiationPaymentTokenCreateResponse struct {
	// A `payment_token` that can be provided to Link initialization to enter the payment initiation flow
	PaymentToken string `json:"payment_token"`
	// The date and time at which the token will expire, in [ISO 8601](https://wikipedia.org/wiki/ISO_8601) format. A `payment_token` expires after 15 minutes.
	PaymentTokenExpirationTime time.Time `json:"payment_token_expiration_time"`
	// A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive.
	RequestId string `json:"request_id"`
	AdditionalProperties map[string]interface{}
}

type _PaymentInitiationPaymentTokenCreateResponse PaymentInitiationPaymentTokenCreateResponse

// NewPaymentInitiationPaymentTokenCreateResponse instantiates a new PaymentInitiationPaymentTokenCreateResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaymentInitiationPaymentTokenCreateResponse(paymentToken string, paymentTokenExpirationTime time.Time, requestId string) *PaymentInitiationPaymentTokenCreateResponse {
	this := PaymentInitiationPaymentTokenCreateResponse{}
	this.PaymentToken = paymentToken
	this.PaymentTokenExpirationTime = paymentTokenExpirationTime
	this.RequestId = requestId
	return &this
}

// NewPaymentInitiationPaymentTokenCreateResponseWithDefaults instantiates a new PaymentInitiationPaymentTokenCreateResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaymentInitiationPaymentTokenCreateResponseWithDefaults() *PaymentInitiationPaymentTokenCreateResponse {
	this := PaymentInitiationPaymentTokenCreateResponse{}
	return &this
}

// GetPaymentToken returns the PaymentToken field value
func (o *PaymentInitiationPaymentTokenCreateResponse) GetPaymentToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PaymentToken
}

// GetPaymentTokenOk returns a tuple with the PaymentToken field value
// and a boolean to check if the value has been set.
func (o *PaymentInitiationPaymentTokenCreateResponse) GetPaymentTokenOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.PaymentToken, true
}

// SetPaymentToken sets field value
func (o *PaymentInitiationPaymentTokenCreateResponse) SetPaymentToken(v string) {
	o.PaymentToken = v
}

// GetPaymentTokenExpirationTime returns the PaymentTokenExpirationTime field value
func (o *PaymentInitiationPaymentTokenCreateResponse) GetPaymentTokenExpirationTime() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.PaymentTokenExpirationTime
}

// GetPaymentTokenExpirationTimeOk returns a tuple with the PaymentTokenExpirationTime field value
// and a boolean to check if the value has been set.
func (o *PaymentInitiationPaymentTokenCreateResponse) GetPaymentTokenExpirationTimeOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.PaymentTokenExpirationTime, true
}

// SetPaymentTokenExpirationTime sets field value
func (o *PaymentInitiationPaymentTokenCreateResponse) SetPaymentTokenExpirationTime(v time.Time) {
	o.PaymentTokenExpirationTime = v
}

// GetRequestId returns the RequestId field value
func (o *PaymentInitiationPaymentTokenCreateResponse) GetRequestId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value
// and a boolean to check if the value has been set.
func (o *PaymentInitiationPaymentTokenCreateResponse) GetRequestIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RequestId, true
}

// SetRequestId sets field value
func (o *PaymentInitiationPaymentTokenCreateResponse) SetRequestId(v string) {
	o.RequestId = v
}

func (o PaymentInitiationPaymentTokenCreateResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["payment_token"] = o.PaymentToken
	}
	if true {
		toSerialize["payment_token_expiration_time"] = o.PaymentTokenExpirationTime
	}
	if true {
		toSerialize["request_id"] = o.RequestId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *PaymentInitiationPaymentTokenCreateResponse) UnmarshalJSON(bytes []byte) (err error) {
	varPaymentInitiationPaymentTokenCreateResponse := _PaymentInitiationPaymentTokenCreateResponse{}

	if err = json.Unmarshal(bytes, &varPaymentInitiationPaymentTokenCreateResponse); err == nil {
		*o = PaymentInitiationPaymentTokenCreateResponse(varPaymentInitiationPaymentTokenCreateResponse)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "payment_token")
		delete(additionalProperties, "payment_token_expiration_time")
		delete(additionalProperties, "request_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePaymentInitiationPaymentTokenCreateResponse struct {
	value *PaymentInitiationPaymentTokenCreateResponse
	isSet bool
}

func (v NullablePaymentInitiationPaymentTokenCreateResponse) Get() *PaymentInitiationPaymentTokenCreateResponse {
	return v.value
}

func (v *NullablePaymentInitiationPaymentTokenCreateResponse) Set(val *PaymentInitiationPaymentTokenCreateResponse) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentInitiationPaymentTokenCreateResponse) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentInitiationPaymentTokenCreateResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentInitiationPaymentTokenCreateResponse(val *PaymentInitiationPaymentTokenCreateResponse) *NullablePaymentInitiationPaymentTokenCreateResponse {
	return &NullablePaymentInitiationPaymentTokenCreateResponse{value: val, isSet: true}
}

func (v NullablePaymentInitiationPaymentTokenCreateResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentInitiationPaymentTokenCreateResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


