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

// SignalDecisionReportRequest SignalDecisionReportRequest defines the request schema for `/signal/decision/report`
type SignalDecisionReportRequest struct {
	// Your Plaid API `client_id`. The `client_id` is required and may be provided either in the `PLAID-CLIENT-ID` header or as part of a request body.
	ClientId *string `json:"client_id,omitempty"`
	// Your Plaid API `secret`. The `secret` is required and may be provided either in the `PLAID-SECRET` header or as part of a request body.
	Secret *string `json:"secret,omitempty"`
	// Must be the same as the `client_transaction_id` supplied when calling `/signal/evaluate`
	ClientTransactionId string `json:"client_transaction_id"`
	// `true` if the ACH transaction was initiated, `false` otherwise.
	Initiated bool `json:"initiated"`
	// The actual number of days (hold time) since the ACH debit transaction that you wait before making funds available to your customers. The holding time could affect the ACH return rate. For example, use 0 if you make funds available to your customers instantly or the same day following the debit transaction, or 1 if you make funds available the next day following the debit initialization.
	DaysFundsOnHold NullableInt32 `json:"days_funds_on_hold,omitempty"`
}

// NewSignalDecisionReportRequest instantiates a new SignalDecisionReportRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSignalDecisionReportRequest(clientTransactionId string, initiated bool) *SignalDecisionReportRequest {
	this := SignalDecisionReportRequest{}
	this.ClientTransactionId = clientTransactionId
	this.Initiated = initiated
	return &this
}

// NewSignalDecisionReportRequestWithDefaults instantiates a new SignalDecisionReportRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSignalDecisionReportRequestWithDefaults() *SignalDecisionReportRequest {
	this := SignalDecisionReportRequest{}
	return &this
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *SignalDecisionReportRequest) GetClientId() string {
	if o == nil || o.ClientId == nil {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignalDecisionReportRequest) GetClientIdOk() (*string, bool) {
	if o == nil || o.ClientId == nil {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *SignalDecisionReportRequest) HasClientId() bool {
	if o != nil && o.ClientId != nil {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *SignalDecisionReportRequest) SetClientId(v string) {
	o.ClientId = &v
}

// GetSecret returns the Secret field value if set, zero value otherwise.
func (o *SignalDecisionReportRequest) GetSecret() string {
	if o == nil || o.Secret == nil {
		var ret string
		return ret
	}
	return *o.Secret
}

// GetSecretOk returns a tuple with the Secret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignalDecisionReportRequest) GetSecretOk() (*string, bool) {
	if o == nil || o.Secret == nil {
		return nil, false
	}
	return o.Secret, true
}

// HasSecret returns a boolean if a field has been set.
func (o *SignalDecisionReportRequest) HasSecret() bool {
	if o != nil && o.Secret != nil {
		return true
	}

	return false
}

// SetSecret gets a reference to the given string and assigns it to the Secret field.
func (o *SignalDecisionReportRequest) SetSecret(v string) {
	o.Secret = &v
}

// GetClientTransactionId returns the ClientTransactionId field value
func (o *SignalDecisionReportRequest) GetClientTransactionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClientTransactionId
}

// GetClientTransactionIdOk returns a tuple with the ClientTransactionId field value
// and a boolean to check if the value has been set.
func (o *SignalDecisionReportRequest) GetClientTransactionIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ClientTransactionId, true
}

// SetClientTransactionId sets field value
func (o *SignalDecisionReportRequest) SetClientTransactionId(v string) {
	o.ClientTransactionId = v
}

// GetInitiated returns the Initiated field value
func (o *SignalDecisionReportRequest) GetInitiated() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Initiated
}

// GetInitiatedOk returns a tuple with the Initiated field value
// and a boolean to check if the value has been set.
func (o *SignalDecisionReportRequest) GetInitiatedOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Initiated, true
}

// SetInitiated sets field value
func (o *SignalDecisionReportRequest) SetInitiated(v bool) {
	o.Initiated = v
}

// GetDaysFundsOnHold returns the DaysFundsOnHold field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SignalDecisionReportRequest) GetDaysFundsOnHold() int32 {
	if o == nil || o.DaysFundsOnHold.Get() == nil {
		var ret int32
		return ret
	}
	return *o.DaysFundsOnHold.Get()
}

// GetDaysFundsOnHoldOk returns a tuple with the DaysFundsOnHold field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SignalDecisionReportRequest) GetDaysFundsOnHoldOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.DaysFundsOnHold.Get(), o.DaysFundsOnHold.IsSet()
}

// HasDaysFundsOnHold returns a boolean if a field has been set.
func (o *SignalDecisionReportRequest) HasDaysFundsOnHold() bool {
	if o != nil && o.DaysFundsOnHold.IsSet() {
		return true
	}

	return false
}

// SetDaysFundsOnHold gets a reference to the given NullableInt32 and assigns it to the DaysFundsOnHold field.
func (o *SignalDecisionReportRequest) SetDaysFundsOnHold(v int32) {
	o.DaysFundsOnHold.Set(&v)
}
// SetDaysFundsOnHoldNil sets the value for DaysFundsOnHold to be an explicit nil
func (o *SignalDecisionReportRequest) SetDaysFundsOnHoldNil() {
	o.DaysFundsOnHold.Set(nil)
}

// UnsetDaysFundsOnHold ensures that no value is present for DaysFundsOnHold, not even an explicit nil
func (o *SignalDecisionReportRequest) UnsetDaysFundsOnHold() {
	o.DaysFundsOnHold.Unset()
}

func (o SignalDecisionReportRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ClientId != nil {
		toSerialize["client_id"] = o.ClientId
	}
	if o.Secret != nil {
		toSerialize["secret"] = o.Secret
	}
	if true {
		toSerialize["client_transaction_id"] = o.ClientTransactionId
	}
	if true {
		toSerialize["initiated"] = o.Initiated
	}
	if o.DaysFundsOnHold.IsSet() {
		toSerialize["days_funds_on_hold"] = o.DaysFundsOnHold.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableSignalDecisionReportRequest struct {
	value *SignalDecisionReportRequest
	isSet bool
}

func (v NullableSignalDecisionReportRequest) Get() *SignalDecisionReportRequest {
	return v.value
}

func (v *NullableSignalDecisionReportRequest) Set(val *SignalDecisionReportRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSignalDecisionReportRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSignalDecisionReportRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSignalDecisionReportRequest(val *SignalDecisionReportRequest) *NullableSignalDecisionReportRequest {
	return &NullableSignalDecisionReportRequest{value: val, isSet: true}
}

func (v NullableSignalDecisionReportRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSignalDecisionReportRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


