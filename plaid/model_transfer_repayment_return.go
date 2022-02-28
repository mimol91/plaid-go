/*
 * The Plaid API
 *
 * The Plaid REST API. Please see https://plaid.com/docs/api for more details.
 *
 * API version: 2020-09-14_1.77.1
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package plaid

import (
	"encoding/json"
)

// TransferRepaymentReturn Represents a return on a Guaranteed ACH transfer that is included in the specified repayment.
type TransferRepaymentReturn struct {
	// The unique identifier of the guaranteed transfer that was returned.
	TransferId string `json:"transfer_id"`
	// The unique identifier of the corresponding `reversed` transfer event.
	EventId int32 `json:"event_id"`
	// The value of the returned transfer.
	Amount string `json:"amount"`
	// The currency of the repayment, e.g. \"USD\".
	IsoCurrencyCode string `json:"iso_currency_code"`
	AdditionalProperties map[string]interface{}
}

type _TransferRepaymentReturn TransferRepaymentReturn

// NewTransferRepaymentReturn instantiates a new TransferRepaymentReturn object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransferRepaymentReturn(transferId string, eventId int32, amount string, isoCurrencyCode string) *TransferRepaymentReturn {
	this := TransferRepaymentReturn{}
	this.TransferId = transferId
	this.EventId = eventId
	this.Amount = amount
	this.IsoCurrencyCode = isoCurrencyCode
	return &this
}

// NewTransferRepaymentReturnWithDefaults instantiates a new TransferRepaymentReturn object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransferRepaymentReturnWithDefaults() *TransferRepaymentReturn {
	this := TransferRepaymentReturn{}
	return &this
}

// GetTransferId returns the TransferId field value
func (o *TransferRepaymentReturn) GetTransferId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TransferId
}

// GetTransferIdOk returns a tuple with the TransferId field value
// and a boolean to check if the value has been set.
func (o *TransferRepaymentReturn) GetTransferIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.TransferId, true
}

// SetTransferId sets field value
func (o *TransferRepaymentReturn) SetTransferId(v string) {
	o.TransferId = v
}

// GetEventId returns the EventId field value
func (o *TransferRepaymentReturn) GetEventId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.EventId
}

// GetEventIdOk returns a tuple with the EventId field value
// and a boolean to check if the value has been set.
func (o *TransferRepaymentReturn) GetEventIdOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.EventId, true
}

// SetEventId sets field value
func (o *TransferRepaymentReturn) SetEventId(v int32) {
	o.EventId = v
}

// GetAmount returns the Amount field value
func (o *TransferRepaymentReturn) GetAmount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *TransferRepaymentReturn) GetAmountOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *TransferRepaymentReturn) SetAmount(v string) {
	o.Amount = v
}

// GetIsoCurrencyCode returns the IsoCurrencyCode field value
func (o *TransferRepaymentReturn) GetIsoCurrencyCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.IsoCurrencyCode
}

// GetIsoCurrencyCodeOk returns a tuple with the IsoCurrencyCode field value
// and a boolean to check if the value has been set.
func (o *TransferRepaymentReturn) GetIsoCurrencyCodeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.IsoCurrencyCode, true
}

// SetIsoCurrencyCode sets field value
func (o *TransferRepaymentReturn) SetIsoCurrencyCode(v string) {
	o.IsoCurrencyCode = v
}

func (o TransferRepaymentReturn) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["transfer_id"] = o.TransferId
	}
	if true {
		toSerialize["event_id"] = o.EventId
	}
	if true {
		toSerialize["amount"] = o.Amount
	}
	if true {
		toSerialize["iso_currency_code"] = o.IsoCurrencyCode
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *TransferRepaymentReturn) UnmarshalJSON(bytes []byte) (err error) {
	varTransferRepaymentReturn := _TransferRepaymentReturn{}

	if err = json.Unmarshal(bytes, &varTransferRepaymentReturn); err == nil {
		*o = TransferRepaymentReturn(varTransferRepaymentReturn)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "transfer_id")
		delete(additionalProperties, "event_id")
		delete(additionalProperties, "amount")
		delete(additionalProperties, "iso_currency_code")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTransferRepaymentReturn struct {
	value *TransferRepaymentReturn
	isSet bool
}

func (v NullableTransferRepaymentReturn) Get() *TransferRepaymentReturn {
	return v.value
}

func (v *NullableTransferRepaymentReturn) Set(val *TransferRepaymentReturn) {
	v.value = val
	v.isSet = true
}

func (v NullableTransferRepaymentReturn) IsSet() bool {
	return v.isSet
}

func (v *NullableTransferRepaymentReturn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransferRepaymentReturn(val *TransferRepaymentReturn) *NullableTransferRepaymentReturn {
	return &NullableTransferRepaymentReturn{value: val, isSet: true}
}

func (v NullableTransferRepaymentReturn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransferRepaymentReturn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


