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
	"time"
)

// TransactionAllOf struct for TransactionAllOf
type TransactionAllOf struct {
	// The channel used to make a payment. `online:` transactions that took place online.  `in store:` transactions that were made at a physical location.  `other:` transactions that relate to banks, e.g. fees or deposits.  This field replaces the `transaction_type` field. 
	PaymentChannel string `json:"payment_channel"`
	// The merchant name, as extracted by Plaid from the `name` field.
	MerchantName NullableString `json:"merchant_name"`
	// The date that the transaction was authorized. Dates are returned in an [ISO 8601](https://wikipedia.org/wiki/ISO_8601) format ( `YYYY-MM-DD` ).
	AuthorizedDate NullableString `json:"authorized_date"`
	// Date and time when a transaction was authorized in [ISO 8601](https://wikipedia.org/wiki/ISO_8601) format ( `YYYY-MM-DDTHH:mm:ssZ` ).  This field is only populated for UK institutions. For institutions in other countries, will be `null`.
	AuthorizedDatetime NullableTime `json:"authorized_datetime"`
	// Date and time when a transaction was posted in [ISO 8601](https://wikipedia.org/wiki/ISO_8601) format ( `YYYY-MM-DDTHH:mm:ssZ` ).  This field is only populated for UK institutions. For institutions in other countries, will be `null`.
	Datetime NullableTime `json:"datetime"`
	// The check number of the transaction. This field is only populated for check transactions.
	CheckNumber NullableString `json:"check_number"`
	TransactionCode NullableTransactionCode `json:"transaction_code"`
	PersonalFinanceCategory NullablePersonalFinanceCategory `json:"personal_finance_category,omitempty"`
}

// NewTransactionAllOf instantiates a new TransactionAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransactionAllOf(paymentChannel string, merchantName NullableString, authorizedDate NullableString, authorizedDatetime NullableTime, datetime NullableTime, checkNumber NullableString, transactionCode NullableTransactionCode) *TransactionAllOf {
	this := TransactionAllOf{}
	this.PaymentChannel = paymentChannel
	this.MerchantName = merchantName
	this.AuthorizedDate = authorizedDate
	this.AuthorizedDatetime = authorizedDatetime
	this.Datetime = datetime
	this.CheckNumber = checkNumber
	this.TransactionCode = transactionCode
	return &this
}

// NewTransactionAllOfWithDefaults instantiates a new TransactionAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransactionAllOfWithDefaults() *TransactionAllOf {
	this := TransactionAllOf{}
	return &this
}

// GetPaymentChannel returns the PaymentChannel field value
func (o *TransactionAllOf) GetPaymentChannel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PaymentChannel
}

// GetPaymentChannelOk returns a tuple with the PaymentChannel field value
// and a boolean to check if the value has been set.
func (o *TransactionAllOf) GetPaymentChannelOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.PaymentChannel, true
}

// SetPaymentChannel sets field value
func (o *TransactionAllOf) SetPaymentChannel(v string) {
	o.PaymentChannel = v
}

// GetMerchantName returns the MerchantName field value
// If the value is explicit nil, the zero value for string will be returned
func (o *TransactionAllOf) GetMerchantName() string {
	if o == nil || o.MerchantName.Get() == nil {
		var ret string
		return ret
	}

	return *o.MerchantName.Get()
}

// GetMerchantNameOk returns a tuple with the MerchantName field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TransactionAllOf) GetMerchantNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.MerchantName.Get(), o.MerchantName.IsSet()
}

// SetMerchantName sets field value
func (o *TransactionAllOf) SetMerchantName(v string) {
	o.MerchantName.Set(&v)
}

// GetAuthorizedDate returns the AuthorizedDate field value
// If the value is explicit nil, the zero value for string will be returned
func (o *TransactionAllOf) GetAuthorizedDate() string {
	if o == nil || o.AuthorizedDate.Get() == nil {
		var ret string
		return ret
	}

	return *o.AuthorizedDate.Get()
}

// GetAuthorizedDateOk returns a tuple with the AuthorizedDate field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TransactionAllOf) GetAuthorizedDateOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.AuthorizedDate.Get(), o.AuthorizedDate.IsSet()
}

// SetAuthorizedDate sets field value
func (o *TransactionAllOf) SetAuthorizedDate(v string) {
	o.AuthorizedDate.Set(&v)
}

// GetAuthorizedDatetime returns the AuthorizedDatetime field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *TransactionAllOf) GetAuthorizedDatetime() time.Time {
	if o == nil || o.AuthorizedDatetime.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.AuthorizedDatetime.Get()
}

// GetAuthorizedDatetimeOk returns a tuple with the AuthorizedDatetime field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TransactionAllOf) GetAuthorizedDatetimeOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.AuthorizedDatetime.Get(), o.AuthorizedDatetime.IsSet()
}

// SetAuthorizedDatetime sets field value
func (o *TransactionAllOf) SetAuthorizedDatetime(v time.Time) {
	o.AuthorizedDatetime.Set(&v)
}

// GetDatetime returns the Datetime field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *TransactionAllOf) GetDatetime() time.Time {
	if o == nil || o.Datetime.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.Datetime.Get()
}

// GetDatetimeOk returns a tuple with the Datetime field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TransactionAllOf) GetDatetimeOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Datetime.Get(), o.Datetime.IsSet()
}

// SetDatetime sets field value
func (o *TransactionAllOf) SetDatetime(v time.Time) {
	o.Datetime.Set(&v)
}

// GetCheckNumber returns the CheckNumber field value
// If the value is explicit nil, the zero value for string will be returned
func (o *TransactionAllOf) GetCheckNumber() string {
	if o == nil || o.CheckNumber.Get() == nil {
		var ret string
		return ret
	}

	return *o.CheckNumber.Get()
}

// GetCheckNumberOk returns a tuple with the CheckNumber field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TransactionAllOf) GetCheckNumberOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.CheckNumber.Get(), o.CheckNumber.IsSet()
}

// SetCheckNumber sets field value
func (o *TransactionAllOf) SetCheckNumber(v string) {
	o.CheckNumber.Set(&v)
}

// GetTransactionCode returns the TransactionCode field value
// If the value is explicit nil, the zero value for TransactionCode will be returned
func (o *TransactionAllOf) GetTransactionCode() TransactionCode {
	if o == nil || o.TransactionCode.Get() == nil {
		var ret TransactionCode
		return ret
	}

	return *o.TransactionCode.Get()
}

// GetTransactionCodeOk returns a tuple with the TransactionCode field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TransactionAllOf) GetTransactionCodeOk() (*TransactionCode, bool) {
	if o == nil  {
		return nil, false
	}
	return o.TransactionCode.Get(), o.TransactionCode.IsSet()
}

// SetTransactionCode sets field value
func (o *TransactionAllOf) SetTransactionCode(v TransactionCode) {
	o.TransactionCode.Set(&v)
}

// GetPersonalFinanceCategory returns the PersonalFinanceCategory field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TransactionAllOf) GetPersonalFinanceCategory() PersonalFinanceCategory {
	if o == nil || o.PersonalFinanceCategory.Get() == nil {
		var ret PersonalFinanceCategory
		return ret
	}
	return *o.PersonalFinanceCategory.Get()
}

// GetPersonalFinanceCategoryOk returns a tuple with the PersonalFinanceCategory field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TransactionAllOf) GetPersonalFinanceCategoryOk() (*PersonalFinanceCategory, bool) {
	if o == nil  {
		return nil, false
	}
	return o.PersonalFinanceCategory.Get(), o.PersonalFinanceCategory.IsSet()
}

// HasPersonalFinanceCategory returns a boolean if a field has been set.
func (o *TransactionAllOf) HasPersonalFinanceCategory() bool {
	if o != nil && o.PersonalFinanceCategory.IsSet() {
		return true
	}

	return false
}

// SetPersonalFinanceCategory gets a reference to the given NullablePersonalFinanceCategory and assigns it to the PersonalFinanceCategory field.
func (o *TransactionAllOf) SetPersonalFinanceCategory(v PersonalFinanceCategory) {
	o.PersonalFinanceCategory.Set(&v)
}
// SetPersonalFinanceCategoryNil sets the value for PersonalFinanceCategory to be an explicit nil
func (o *TransactionAllOf) SetPersonalFinanceCategoryNil() {
	o.PersonalFinanceCategory.Set(nil)
}

// UnsetPersonalFinanceCategory ensures that no value is present for PersonalFinanceCategory, not even an explicit nil
func (o *TransactionAllOf) UnsetPersonalFinanceCategory() {
	o.PersonalFinanceCategory.Unset()
}

func (o TransactionAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["payment_channel"] = o.PaymentChannel
	}
	if true {
		toSerialize["merchant_name"] = o.MerchantName.Get()
	}
	if true {
		toSerialize["authorized_date"] = o.AuthorizedDate.Get()
	}
	if true {
		toSerialize["authorized_datetime"] = o.AuthorizedDatetime.Get()
	}
	if true {
		toSerialize["datetime"] = o.Datetime.Get()
	}
	if true {
		toSerialize["check_number"] = o.CheckNumber.Get()
	}
	if true {
		toSerialize["transaction_code"] = o.TransactionCode.Get()
	}
	if o.PersonalFinanceCategory.IsSet() {
		toSerialize["personal_finance_category"] = o.PersonalFinanceCategory.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableTransactionAllOf struct {
	value *TransactionAllOf
	isSet bool
}

func (v NullableTransactionAllOf) Get() *TransactionAllOf {
	return v.value
}

func (v *NullableTransactionAllOf) Set(val *TransactionAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableTransactionAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableTransactionAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransactionAllOf(val *TransactionAllOf) *NullableTransactionAllOf {
	return &NullableTransactionAllOf{value: val, isSet: true}
}

func (v NullableTransactionAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransactionAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


