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

// PaymentInitiationPaymentGetResponse PaymentInitiationPaymentGetResponse defines the response schema for `/payment_initation/payment/get`
type PaymentInitiationPaymentGetResponse struct {
	// The ID of the payment. Like all Plaid identifiers, the `payment_id` is case sensitive.
	PaymentId string `json:"payment_id"`
	Amount PaymentAmount `json:"amount"`
	// The status of the payment.  `PAYMENT_STATUS_INPUT_NEEDED`: This is the initial state of all payments. It indicates that the payment is waiting on user input to continue processing. A payment may re-enter this state later on if further input is needed.  `PAYMENT_STATUS_PROCESSING`: The payment is currently being processed. The payment will automatically exit this state when processing is complete.  `PAYMENT_STATUS_INITIATED`: The payment has been successfully initiated and is considered complete.  `PAYMENT_STATUS_COMPLETED`: Indicates that the standing order has been successfully established. This state is only used for standing orders.  `PAYMENT_STATUS_INSUFFICIENT_FUNDS`: The payment has failed due to insufficient funds.  `PAYMENT_STATUS_FAILED`: The payment has failed to be initiated. This error is retryable once the root cause is resolved.  `PAYMENT_STATUS_BLOCKED`: The payment has been blocked. This is a retryable error.  `PAYMENT_STATUS_UNKNOWN`: The payment status is unknown.
	Status string `json:"status"`
	// The ID of the recipient
	RecipientId string `json:"recipient_id"`
	// A reference for the payment.
	Reference string `json:"reference"`
	// The value of the reference sent to the bank after adjustment to pass bank validation rules.
	AdjustedReference NullableString `json:"adjusted_reference,omitempty"`
	// The date and time of the last time the `status` was updated, in IS0 8601 format
	LastStatusUpdate time.Time `json:"last_status_update"`
	Schedule NullableExternalPaymentScheduleGet `json:"schedule,omitempty"`
	RefundDetails NullableExternalPaymentRefundDetails `json:"refund_details,omitempty"`
	Bacs NullableSenderBACSNullable `json:"bacs"`
	// The International Bank Account Number (IBAN) for the sender, if specified in the `/payment_initiation/payment/create` call.
	Iban NullableString `json:"iban"`
	// Initiated refunds associated with the payment.
	InitiatedRefunds *[]PaymentInitiationRefund `json:"initiated_refunds,omitempty"`
	// The EMI (E-Money Institution) account that this payment is associated with, if any. This EMI account is used as an intermediary account to enable Plaid to reconcile the settlement of funds for Payment Initiation requests.
	EmiAccountId NullableString `json:"emi_account_id,omitempty"`
	// A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive.
	RequestId string `json:"request_id"`
	AdditionalProperties map[string]interface{}
}

type _PaymentInitiationPaymentGetResponse PaymentInitiationPaymentGetResponse

// NewPaymentInitiationPaymentGetResponse instantiates a new PaymentInitiationPaymentGetResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaymentInitiationPaymentGetResponse(paymentId string, amount PaymentAmount, status string, recipientId string, reference string, lastStatusUpdate time.Time, bacs NullableSenderBACSNullable, iban NullableString, requestId string) *PaymentInitiationPaymentGetResponse {
	this := PaymentInitiationPaymentGetResponse{}
	this.PaymentId = paymentId
	this.Amount = amount
	this.Status = status
	this.RecipientId = recipientId
	this.Reference = reference
	this.LastStatusUpdate = lastStatusUpdate
	this.Bacs = bacs
	this.Iban = iban
	this.RequestId = requestId
	return &this
}

// NewPaymentInitiationPaymentGetResponseWithDefaults instantiates a new PaymentInitiationPaymentGetResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaymentInitiationPaymentGetResponseWithDefaults() *PaymentInitiationPaymentGetResponse {
	this := PaymentInitiationPaymentGetResponse{}
	return &this
}

// GetPaymentId returns the PaymentId field value
func (o *PaymentInitiationPaymentGetResponse) GetPaymentId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PaymentId
}

// GetPaymentIdOk returns a tuple with the PaymentId field value
// and a boolean to check if the value has been set.
func (o *PaymentInitiationPaymentGetResponse) GetPaymentIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.PaymentId, true
}

// SetPaymentId sets field value
func (o *PaymentInitiationPaymentGetResponse) SetPaymentId(v string) {
	o.PaymentId = v
}

// GetAmount returns the Amount field value
func (o *PaymentInitiationPaymentGetResponse) GetAmount() PaymentAmount {
	if o == nil {
		var ret PaymentAmount
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *PaymentInitiationPaymentGetResponse) GetAmountOk() (*PaymentAmount, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *PaymentInitiationPaymentGetResponse) SetAmount(v PaymentAmount) {
	o.Amount = v
}

// GetStatus returns the Status field value
func (o *PaymentInitiationPaymentGetResponse) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *PaymentInitiationPaymentGetResponse) GetStatusOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *PaymentInitiationPaymentGetResponse) SetStatus(v string) {
	o.Status = v
}

// GetRecipientId returns the RecipientId field value
func (o *PaymentInitiationPaymentGetResponse) GetRecipientId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RecipientId
}

// GetRecipientIdOk returns a tuple with the RecipientId field value
// and a boolean to check if the value has been set.
func (o *PaymentInitiationPaymentGetResponse) GetRecipientIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RecipientId, true
}

// SetRecipientId sets field value
func (o *PaymentInitiationPaymentGetResponse) SetRecipientId(v string) {
	o.RecipientId = v
}

// GetReference returns the Reference field value
func (o *PaymentInitiationPaymentGetResponse) GetReference() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value
// and a boolean to check if the value has been set.
func (o *PaymentInitiationPaymentGetResponse) GetReferenceOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Reference, true
}

// SetReference sets field value
func (o *PaymentInitiationPaymentGetResponse) SetReference(v string) {
	o.Reference = v
}

// GetAdjustedReference returns the AdjustedReference field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaymentInitiationPaymentGetResponse) GetAdjustedReference() string {
	if o == nil || o.AdjustedReference.Get() == nil {
		var ret string
		return ret
	}
	return *o.AdjustedReference.Get()
}

// GetAdjustedReferenceOk returns a tuple with the AdjustedReference field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaymentInitiationPaymentGetResponse) GetAdjustedReferenceOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.AdjustedReference.Get(), o.AdjustedReference.IsSet()
}

// HasAdjustedReference returns a boolean if a field has been set.
func (o *PaymentInitiationPaymentGetResponse) HasAdjustedReference() bool {
	if o != nil && o.AdjustedReference.IsSet() {
		return true
	}

	return false
}

// SetAdjustedReference gets a reference to the given NullableString and assigns it to the AdjustedReference field.
func (o *PaymentInitiationPaymentGetResponse) SetAdjustedReference(v string) {
	o.AdjustedReference.Set(&v)
}
// SetAdjustedReferenceNil sets the value for AdjustedReference to be an explicit nil
func (o *PaymentInitiationPaymentGetResponse) SetAdjustedReferenceNil() {
	o.AdjustedReference.Set(nil)
}

// UnsetAdjustedReference ensures that no value is present for AdjustedReference, not even an explicit nil
func (o *PaymentInitiationPaymentGetResponse) UnsetAdjustedReference() {
	o.AdjustedReference.Unset()
}

// GetLastStatusUpdate returns the LastStatusUpdate field value
func (o *PaymentInitiationPaymentGetResponse) GetLastStatusUpdate() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.LastStatusUpdate
}

// GetLastStatusUpdateOk returns a tuple with the LastStatusUpdate field value
// and a boolean to check if the value has been set.
func (o *PaymentInitiationPaymentGetResponse) GetLastStatusUpdateOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.LastStatusUpdate, true
}

// SetLastStatusUpdate sets field value
func (o *PaymentInitiationPaymentGetResponse) SetLastStatusUpdate(v time.Time) {
	o.LastStatusUpdate = v
}

// GetSchedule returns the Schedule field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaymentInitiationPaymentGetResponse) GetSchedule() ExternalPaymentScheduleGet {
	if o == nil || o.Schedule.Get() == nil {
		var ret ExternalPaymentScheduleGet
		return ret
	}
	return *o.Schedule.Get()
}

// GetScheduleOk returns a tuple with the Schedule field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaymentInitiationPaymentGetResponse) GetScheduleOk() (*ExternalPaymentScheduleGet, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Schedule.Get(), o.Schedule.IsSet()
}

// HasSchedule returns a boolean if a field has been set.
func (o *PaymentInitiationPaymentGetResponse) HasSchedule() bool {
	if o != nil && o.Schedule.IsSet() {
		return true
	}

	return false
}

// SetSchedule gets a reference to the given NullableExternalPaymentScheduleGet and assigns it to the Schedule field.
func (o *PaymentInitiationPaymentGetResponse) SetSchedule(v ExternalPaymentScheduleGet) {
	o.Schedule.Set(&v)
}
// SetScheduleNil sets the value for Schedule to be an explicit nil
func (o *PaymentInitiationPaymentGetResponse) SetScheduleNil() {
	o.Schedule.Set(nil)
}

// UnsetSchedule ensures that no value is present for Schedule, not even an explicit nil
func (o *PaymentInitiationPaymentGetResponse) UnsetSchedule() {
	o.Schedule.Unset()
}

// GetRefundDetails returns the RefundDetails field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaymentInitiationPaymentGetResponse) GetRefundDetails() ExternalPaymentRefundDetails {
	if o == nil || o.RefundDetails.Get() == nil {
		var ret ExternalPaymentRefundDetails
		return ret
	}
	return *o.RefundDetails.Get()
}

// GetRefundDetailsOk returns a tuple with the RefundDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaymentInitiationPaymentGetResponse) GetRefundDetailsOk() (*ExternalPaymentRefundDetails, bool) {
	if o == nil  {
		return nil, false
	}
	return o.RefundDetails.Get(), o.RefundDetails.IsSet()
}

// HasRefundDetails returns a boolean if a field has been set.
func (o *PaymentInitiationPaymentGetResponse) HasRefundDetails() bool {
	if o != nil && o.RefundDetails.IsSet() {
		return true
	}

	return false
}

// SetRefundDetails gets a reference to the given NullableExternalPaymentRefundDetails and assigns it to the RefundDetails field.
func (o *PaymentInitiationPaymentGetResponse) SetRefundDetails(v ExternalPaymentRefundDetails) {
	o.RefundDetails.Set(&v)
}
// SetRefundDetailsNil sets the value for RefundDetails to be an explicit nil
func (o *PaymentInitiationPaymentGetResponse) SetRefundDetailsNil() {
	o.RefundDetails.Set(nil)
}

// UnsetRefundDetails ensures that no value is present for RefundDetails, not even an explicit nil
func (o *PaymentInitiationPaymentGetResponse) UnsetRefundDetails() {
	o.RefundDetails.Unset()
}

// GetBacs returns the Bacs field value
// If the value is explicit nil, the zero value for SenderBACSNullable will be returned
func (o *PaymentInitiationPaymentGetResponse) GetBacs() SenderBACSNullable {
	if o == nil || o.Bacs.Get() == nil {
		var ret SenderBACSNullable
		return ret
	}

	return *o.Bacs.Get()
}

// GetBacsOk returns a tuple with the Bacs field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaymentInitiationPaymentGetResponse) GetBacsOk() (*SenderBACSNullable, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Bacs.Get(), o.Bacs.IsSet()
}

// SetBacs sets field value
func (o *PaymentInitiationPaymentGetResponse) SetBacs(v SenderBACSNullable) {
	o.Bacs.Set(&v)
}

// GetIban returns the Iban field value
// If the value is explicit nil, the zero value for string will be returned
func (o *PaymentInitiationPaymentGetResponse) GetIban() string {
	if o == nil || o.Iban.Get() == nil {
		var ret string
		return ret
	}

	return *o.Iban.Get()
}

// GetIbanOk returns a tuple with the Iban field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaymentInitiationPaymentGetResponse) GetIbanOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Iban.Get(), o.Iban.IsSet()
}

// SetIban sets field value
func (o *PaymentInitiationPaymentGetResponse) SetIban(v string) {
	o.Iban.Set(&v)
}

// GetInitiatedRefunds returns the InitiatedRefunds field value if set, zero value otherwise.
func (o *PaymentInitiationPaymentGetResponse) GetInitiatedRefunds() []PaymentInitiationRefund {
	if o == nil || o.InitiatedRefunds == nil {
		var ret []PaymentInitiationRefund
		return ret
	}
	return *o.InitiatedRefunds
}

// GetInitiatedRefundsOk returns a tuple with the InitiatedRefunds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentInitiationPaymentGetResponse) GetInitiatedRefundsOk() (*[]PaymentInitiationRefund, bool) {
	if o == nil || o.InitiatedRefunds == nil {
		return nil, false
	}
	return o.InitiatedRefunds, true
}

// HasInitiatedRefunds returns a boolean if a field has been set.
func (o *PaymentInitiationPaymentGetResponse) HasInitiatedRefunds() bool {
	if o != nil && o.InitiatedRefunds != nil {
		return true
	}

	return false
}

// SetInitiatedRefunds gets a reference to the given []PaymentInitiationRefund and assigns it to the InitiatedRefunds field.
func (o *PaymentInitiationPaymentGetResponse) SetInitiatedRefunds(v []PaymentInitiationRefund) {
	o.InitiatedRefunds = &v
}

// GetEmiAccountId returns the EmiAccountId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaymentInitiationPaymentGetResponse) GetEmiAccountId() string {
	if o == nil || o.EmiAccountId.Get() == nil {
		var ret string
		return ret
	}
	return *o.EmiAccountId.Get()
}

// GetEmiAccountIdOk returns a tuple with the EmiAccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaymentInitiationPaymentGetResponse) GetEmiAccountIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.EmiAccountId.Get(), o.EmiAccountId.IsSet()
}

// HasEmiAccountId returns a boolean if a field has been set.
func (o *PaymentInitiationPaymentGetResponse) HasEmiAccountId() bool {
	if o != nil && o.EmiAccountId.IsSet() {
		return true
	}

	return false
}

// SetEmiAccountId gets a reference to the given NullableString and assigns it to the EmiAccountId field.
func (o *PaymentInitiationPaymentGetResponse) SetEmiAccountId(v string) {
	o.EmiAccountId.Set(&v)
}
// SetEmiAccountIdNil sets the value for EmiAccountId to be an explicit nil
func (o *PaymentInitiationPaymentGetResponse) SetEmiAccountIdNil() {
	o.EmiAccountId.Set(nil)
}

// UnsetEmiAccountId ensures that no value is present for EmiAccountId, not even an explicit nil
func (o *PaymentInitiationPaymentGetResponse) UnsetEmiAccountId() {
	o.EmiAccountId.Unset()
}

// GetRequestId returns the RequestId field value
func (o *PaymentInitiationPaymentGetResponse) GetRequestId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value
// and a boolean to check if the value has been set.
func (o *PaymentInitiationPaymentGetResponse) GetRequestIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RequestId, true
}

// SetRequestId sets field value
func (o *PaymentInitiationPaymentGetResponse) SetRequestId(v string) {
	o.RequestId = v
}

func (o PaymentInitiationPaymentGetResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["payment_id"] = o.PaymentId
	}
	if true {
		toSerialize["amount"] = o.Amount
	}
	if true {
		toSerialize["status"] = o.Status
	}
	if true {
		toSerialize["recipient_id"] = o.RecipientId
	}
	if true {
		toSerialize["reference"] = o.Reference
	}
	if o.AdjustedReference.IsSet() {
		toSerialize["adjusted_reference"] = o.AdjustedReference.Get()
	}
	if true {
		toSerialize["last_status_update"] = o.LastStatusUpdate
	}
	if o.Schedule.IsSet() {
		toSerialize["schedule"] = o.Schedule.Get()
	}
	if o.RefundDetails.IsSet() {
		toSerialize["refund_details"] = o.RefundDetails.Get()
	}
	if true {
		toSerialize["bacs"] = o.Bacs.Get()
	}
	if true {
		toSerialize["iban"] = o.Iban.Get()
	}
	if o.InitiatedRefunds != nil {
		toSerialize["initiated_refunds"] = o.InitiatedRefunds
	}
	if o.EmiAccountId.IsSet() {
		toSerialize["emi_account_id"] = o.EmiAccountId.Get()
	}
	if true {
		toSerialize["request_id"] = o.RequestId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *PaymentInitiationPaymentGetResponse) UnmarshalJSON(bytes []byte) (err error) {
	varPaymentInitiationPaymentGetResponse := _PaymentInitiationPaymentGetResponse{}

	if err = json.Unmarshal(bytes, &varPaymentInitiationPaymentGetResponse); err == nil {
		*o = PaymentInitiationPaymentGetResponse(varPaymentInitiationPaymentGetResponse)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "payment_id")
		delete(additionalProperties, "amount")
		delete(additionalProperties, "status")
		delete(additionalProperties, "recipient_id")
		delete(additionalProperties, "reference")
		delete(additionalProperties, "adjusted_reference")
		delete(additionalProperties, "last_status_update")
		delete(additionalProperties, "schedule")
		delete(additionalProperties, "refund_details")
		delete(additionalProperties, "bacs")
		delete(additionalProperties, "iban")
		delete(additionalProperties, "initiated_refunds")
		delete(additionalProperties, "emi_account_id")
		delete(additionalProperties, "request_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePaymentInitiationPaymentGetResponse struct {
	value *PaymentInitiationPaymentGetResponse
	isSet bool
}

func (v NullablePaymentInitiationPaymentGetResponse) Get() *PaymentInitiationPaymentGetResponse {
	return v.value
}

func (v *NullablePaymentInitiationPaymentGetResponse) Set(val *PaymentInitiationPaymentGetResponse) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentInitiationPaymentGetResponse) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentInitiationPaymentGetResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentInitiationPaymentGetResponse(val *PaymentInitiationPaymentGetResponse) *NullablePaymentInitiationPaymentGetResponse {
	return &NullablePaymentInitiationPaymentGetResponse{value: val, isSet: true}
}

func (v NullablePaymentInitiationPaymentGetResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentInitiationPaymentGetResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


