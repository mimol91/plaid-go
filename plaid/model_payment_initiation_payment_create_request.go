/*
 * The Plaid API
 *
 * The Plaid REST API. Please see https://plaid.com/docs/api for more details.
 *
 * API version: 2020-09-14_1.94.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package plaid

import (
	"encoding/json"
)

// PaymentInitiationPaymentCreateRequest PaymentInitiationPaymentCreateRequest defines the request schema for `/payment_initiation/payment/create`
type PaymentInitiationPaymentCreateRequest struct {
	// Your Plaid API `client_id`. The `client_id` is required and may be provided either in the `PLAID-CLIENT-ID` header or as part of a request body.
	ClientId *string `json:"client_id,omitempty"`
	// Your Plaid API `secret`. The `secret` is required and may be provided either in the `PLAID-SECRET` header or as part of a request body.
	Secret *string `json:"secret,omitempty"`
	// The ID of the recipient the payment is for.
	RecipientId string `json:"recipient_id"`
	// A reference for the payment. This must be an alphanumeric string with at most 18 characters and must not contain any special characters (since not all institutions support them).
	Reference string `json:"reference"`
	Amount PaymentAmount `json:"amount"`
	Schedule *ExternalPaymentScheduleRequest `json:"schedule,omitempty"`
	Options NullableExternalPaymentOptions `json:"options,omitempty"`
}

// NewPaymentInitiationPaymentCreateRequest instantiates a new PaymentInitiationPaymentCreateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaymentInitiationPaymentCreateRequest(recipientId string, reference string, amount PaymentAmount) *PaymentInitiationPaymentCreateRequest {
	this := PaymentInitiationPaymentCreateRequest{}
	this.RecipientId = recipientId
	this.Reference = reference
	this.Amount = amount
	return &this
}

// NewPaymentInitiationPaymentCreateRequestWithDefaults instantiates a new PaymentInitiationPaymentCreateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaymentInitiationPaymentCreateRequestWithDefaults() *PaymentInitiationPaymentCreateRequest {
	this := PaymentInitiationPaymentCreateRequest{}
	return &this
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *PaymentInitiationPaymentCreateRequest) GetClientId() string {
	if o == nil || o.ClientId == nil {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentInitiationPaymentCreateRequest) GetClientIdOk() (*string, bool) {
	if o == nil || o.ClientId == nil {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *PaymentInitiationPaymentCreateRequest) HasClientId() bool {
	if o != nil && o.ClientId != nil {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *PaymentInitiationPaymentCreateRequest) SetClientId(v string) {
	o.ClientId = &v
}

// GetSecret returns the Secret field value if set, zero value otherwise.
func (o *PaymentInitiationPaymentCreateRequest) GetSecret() string {
	if o == nil || o.Secret == nil {
		var ret string
		return ret
	}
	return *o.Secret
}

// GetSecretOk returns a tuple with the Secret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentInitiationPaymentCreateRequest) GetSecretOk() (*string, bool) {
	if o == nil || o.Secret == nil {
		return nil, false
	}
	return o.Secret, true
}

// HasSecret returns a boolean if a field has been set.
func (o *PaymentInitiationPaymentCreateRequest) HasSecret() bool {
	if o != nil && o.Secret != nil {
		return true
	}

	return false
}

// SetSecret gets a reference to the given string and assigns it to the Secret field.
func (o *PaymentInitiationPaymentCreateRequest) SetSecret(v string) {
	o.Secret = &v
}

// GetRecipientId returns the RecipientId field value
func (o *PaymentInitiationPaymentCreateRequest) GetRecipientId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RecipientId
}

// GetRecipientIdOk returns a tuple with the RecipientId field value
// and a boolean to check if the value has been set.
func (o *PaymentInitiationPaymentCreateRequest) GetRecipientIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RecipientId, true
}

// SetRecipientId sets field value
func (o *PaymentInitiationPaymentCreateRequest) SetRecipientId(v string) {
	o.RecipientId = v
}

// GetReference returns the Reference field value
func (o *PaymentInitiationPaymentCreateRequest) GetReference() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value
// and a boolean to check if the value has been set.
func (o *PaymentInitiationPaymentCreateRequest) GetReferenceOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Reference, true
}

// SetReference sets field value
func (o *PaymentInitiationPaymentCreateRequest) SetReference(v string) {
	o.Reference = v
}

// GetAmount returns the Amount field value
func (o *PaymentInitiationPaymentCreateRequest) GetAmount() PaymentAmount {
	if o == nil {
		var ret PaymentAmount
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *PaymentInitiationPaymentCreateRequest) GetAmountOk() (*PaymentAmount, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *PaymentInitiationPaymentCreateRequest) SetAmount(v PaymentAmount) {
	o.Amount = v
}

// GetSchedule returns the Schedule field value if set, zero value otherwise.
func (o *PaymentInitiationPaymentCreateRequest) GetSchedule() ExternalPaymentScheduleRequest {
	if o == nil || o.Schedule == nil {
		var ret ExternalPaymentScheduleRequest
		return ret
	}
	return *o.Schedule
}

// GetScheduleOk returns a tuple with the Schedule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentInitiationPaymentCreateRequest) GetScheduleOk() (*ExternalPaymentScheduleRequest, bool) {
	if o == nil || o.Schedule == nil {
		return nil, false
	}
	return o.Schedule, true
}

// HasSchedule returns a boolean if a field has been set.
func (o *PaymentInitiationPaymentCreateRequest) HasSchedule() bool {
	if o != nil && o.Schedule != nil {
		return true
	}

	return false
}

// SetSchedule gets a reference to the given ExternalPaymentScheduleRequest and assigns it to the Schedule field.
func (o *PaymentInitiationPaymentCreateRequest) SetSchedule(v ExternalPaymentScheduleRequest) {
	o.Schedule = &v
}

// GetOptions returns the Options field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaymentInitiationPaymentCreateRequest) GetOptions() ExternalPaymentOptions {
	if o == nil || o.Options.Get() == nil {
		var ret ExternalPaymentOptions
		return ret
	}
	return *o.Options.Get()
}

// GetOptionsOk returns a tuple with the Options field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaymentInitiationPaymentCreateRequest) GetOptionsOk() (*ExternalPaymentOptions, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Options.Get(), o.Options.IsSet()
}

// HasOptions returns a boolean if a field has been set.
func (o *PaymentInitiationPaymentCreateRequest) HasOptions() bool {
	if o != nil && o.Options.IsSet() {
		return true
	}

	return false
}

// SetOptions gets a reference to the given NullableExternalPaymentOptions and assigns it to the Options field.
func (o *PaymentInitiationPaymentCreateRequest) SetOptions(v ExternalPaymentOptions) {
	o.Options.Set(&v)
}
// SetOptionsNil sets the value for Options to be an explicit nil
func (o *PaymentInitiationPaymentCreateRequest) SetOptionsNil() {
	o.Options.Set(nil)
}

// UnsetOptions ensures that no value is present for Options, not even an explicit nil
func (o *PaymentInitiationPaymentCreateRequest) UnsetOptions() {
	o.Options.Unset()
}

func (o PaymentInitiationPaymentCreateRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ClientId != nil {
		toSerialize["client_id"] = o.ClientId
	}
	if o.Secret != nil {
		toSerialize["secret"] = o.Secret
	}
	if true {
		toSerialize["recipient_id"] = o.RecipientId
	}
	if true {
		toSerialize["reference"] = o.Reference
	}
	if true {
		toSerialize["amount"] = o.Amount
	}
	if o.Schedule != nil {
		toSerialize["schedule"] = o.Schedule
	}
	if o.Options.IsSet() {
		toSerialize["options"] = o.Options.Get()
	}
	return json.Marshal(toSerialize)
}

type NullablePaymentInitiationPaymentCreateRequest struct {
	value *PaymentInitiationPaymentCreateRequest
	isSet bool
}

func (v NullablePaymentInitiationPaymentCreateRequest) Get() *PaymentInitiationPaymentCreateRequest {
	return v.value
}

func (v *NullablePaymentInitiationPaymentCreateRequest) Set(val *PaymentInitiationPaymentCreateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentInitiationPaymentCreateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentInitiationPaymentCreateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentInitiationPaymentCreateRequest(val *PaymentInitiationPaymentCreateRequest) *NullablePaymentInitiationPaymentCreateRequest {
	return &NullablePaymentInitiationPaymentCreateRequest{value: val, isSet: true}
}

func (v NullablePaymentInitiationPaymentCreateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentInitiationPaymentCreateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


