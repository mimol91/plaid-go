/*
 * The Plaid API
 *
 * The Plaid REST API. Please see https://plaid.com/docs/api for more details.
 *
 * API version: 2020-09-14_1.56.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package plaid

import (
	"encoding/json"
)

// TransferIntentGetFailureReason The reason for a failed transfer intent. Returned only if the transfer intent status is `failed`. Null otherwise.
type TransferIntentGetFailureReason struct {
	// A broad categorization of the error.
	ErrorType *string `json:"error_type,omitempty"`
	// A code representing the reason for a failed transfer intent (i.e., an API error or the authorization being declined).  For a full listing of bank transfer errors, see [Bank Transfers errors](https://plaid.com/docs/errors/bank-transfers/).
	ErrorCode *string `json:"error_code,omitempty"`
	// A human-readable description of the code associated with a failed transfer intent.
	ErrorMessage *string `json:"error_message,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _TransferIntentGetFailureReason TransferIntentGetFailureReason

// NewTransferIntentGetFailureReason instantiates a new TransferIntentGetFailureReason object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransferIntentGetFailureReason() *TransferIntentGetFailureReason {
	this := TransferIntentGetFailureReason{}
	return &this
}

// NewTransferIntentGetFailureReasonWithDefaults instantiates a new TransferIntentGetFailureReason object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransferIntentGetFailureReasonWithDefaults() *TransferIntentGetFailureReason {
	this := TransferIntentGetFailureReason{}
	return &this
}

// GetErrorType returns the ErrorType field value if set, zero value otherwise.
func (o *TransferIntentGetFailureReason) GetErrorType() string {
	if o == nil || o.ErrorType == nil {
		var ret string
		return ret
	}
	return *o.ErrorType
}

// GetErrorTypeOk returns a tuple with the ErrorType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferIntentGetFailureReason) GetErrorTypeOk() (*string, bool) {
	if o == nil || o.ErrorType == nil {
		return nil, false
	}
	return o.ErrorType, true
}

// HasErrorType returns a boolean if a field has been set.
func (o *TransferIntentGetFailureReason) HasErrorType() bool {
	if o != nil && o.ErrorType != nil {
		return true
	}

	return false
}

// SetErrorType gets a reference to the given string and assigns it to the ErrorType field.
func (o *TransferIntentGetFailureReason) SetErrorType(v string) {
	o.ErrorType = &v
}

// GetErrorCode returns the ErrorCode field value if set, zero value otherwise.
func (o *TransferIntentGetFailureReason) GetErrorCode() string {
	if o == nil || o.ErrorCode == nil {
		var ret string
		return ret
	}
	return *o.ErrorCode
}

// GetErrorCodeOk returns a tuple with the ErrorCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferIntentGetFailureReason) GetErrorCodeOk() (*string, bool) {
	if o == nil || o.ErrorCode == nil {
		return nil, false
	}
	return o.ErrorCode, true
}

// HasErrorCode returns a boolean if a field has been set.
func (o *TransferIntentGetFailureReason) HasErrorCode() bool {
	if o != nil && o.ErrorCode != nil {
		return true
	}

	return false
}

// SetErrorCode gets a reference to the given string and assigns it to the ErrorCode field.
func (o *TransferIntentGetFailureReason) SetErrorCode(v string) {
	o.ErrorCode = &v
}

// GetErrorMessage returns the ErrorMessage field value if set, zero value otherwise.
func (o *TransferIntentGetFailureReason) GetErrorMessage() string {
	if o == nil || o.ErrorMessage == nil {
		var ret string
		return ret
	}
	return *o.ErrorMessage
}

// GetErrorMessageOk returns a tuple with the ErrorMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferIntentGetFailureReason) GetErrorMessageOk() (*string, bool) {
	if o == nil || o.ErrorMessage == nil {
		return nil, false
	}
	return o.ErrorMessage, true
}

// HasErrorMessage returns a boolean if a field has been set.
func (o *TransferIntentGetFailureReason) HasErrorMessage() bool {
	if o != nil && o.ErrorMessage != nil {
		return true
	}

	return false
}

// SetErrorMessage gets a reference to the given string and assigns it to the ErrorMessage field.
func (o *TransferIntentGetFailureReason) SetErrorMessage(v string) {
	o.ErrorMessage = &v
}

func (o TransferIntentGetFailureReason) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ErrorType != nil {
		toSerialize["error_type"] = o.ErrorType
	}
	if o.ErrorCode != nil {
		toSerialize["error_code"] = o.ErrorCode
	}
	if o.ErrorMessage != nil {
		toSerialize["error_message"] = o.ErrorMessage
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *TransferIntentGetFailureReason) UnmarshalJSON(bytes []byte) (err error) {
	varTransferIntentGetFailureReason := _TransferIntentGetFailureReason{}

	if err = json.Unmarshal(bytes, &varTransferIntentGetFailureReason); err == nil {
		*o = TransferIntentGetFailureReason(varTransferIntentGetFailureReason)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "error_type")
		delete(additionalProperties, "error_code")
		delete(additionalProperties, "error_message")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTransferIntentGetFailureReason struct {
	value *TransferIntentGetFailureReason
	isSet bool
}

func (v NullableTransferIntentGetFailureReason) Get() *TransferIntentGetFailureReason {
	return v.value
}

func (v *NullableTransferIntentGetFailureReason) Set(val *TransferIntentGetFailureReason) {
	v.value = val
	v.isSet = true
}

func (v NullableTransferIntentGetFailureReason) IsSet() bool {
	return v.isSet
}

func (v *NullableTransferIntentGetFailureReason) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransferIntentGetFailureReason(val *TransferIntentGetFailureReason) *NullableTransferIntentGetFailureReason {
	return &NullableTransferIntentGetFailureReason{value: val, isSet: true}
}

func (v NullableTransferIntentGetFailureReason) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransferIntentGetFailureReason) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


