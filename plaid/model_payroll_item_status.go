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

// PayrollItemStatus Details about the status of the payroll item.
type PayrollItemStatus struct {
	// Denotes the processing status for the verification.  `UNKNOWN`: The processing status could not be determined.  `PROCESSING_COMPLETE`: The processing has completed and the data is available to be retrieved.  `PROCESSING`: The verification is still processing. The data is not available yet.  `UPLOADED`: The user uploaded a document. The document has not yet begun processing.  `CREATED`: The verification has been created but no data is associated with it yet.  `FAILED`: The processing failed to complete successfully.  `APPROVAL_STATUS_PENDING`: The user has not yet approved the sharing of the data.  `APPROVAL_STATUS_APPROVED`: The user has approved the sharing of the data.
	ProcessingStatus NullableString `json:"processing_status,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PayrollItemStatus PayrollItemStatus

// NewPayrollItemStatus instantiates a new PayrollItemStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPayrollItemStatus() *PayrollItemStatus {
	this := PayrollItemStatus{}
	return &this
}

// NewPayrollItemStatusWithDefaults instantiates a new PayrollItemStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPayrollItemStatusWithDefaults() *PayrollItemStatus {
	this := PayrollItemStatus{}
	return &this
}

// GetProcessingStatus returns the ProcessingStatus field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PayrollItemStatus) GetProcessingStatus() string {
	if o == nil || o.ProcessingStatus.Get() == nil {
		var ret string
		return ret
	}
	return *o.ProcessingStatus.Get()
}

// GetProcessingStatusOk returns a tuple with the ProcessingStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PayrollItemStatus) GetProcessingStatusOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ProcessingStatus.Get(), o.ProcessingStatus.IsSet()
}

// HasProcessingStatus returns a boolean if a field has been set.
func (o *PayrollItemStatus) HasProcessingStatus() bool {
	if o != nil && o.ProcessingStatus.IsSet() {
		return true
	}

	return false
}

// SetProcessingStatus gets a reference to the given NullableString and assigns it to the ProcessingStatus field.
func (o *PayrollItemStatus) SetProcessingStatus(v string) {
	o.ProcessingStatus.Set(&v)
}
// SetProcessingStatusNil sets the value for ProcessingStatus to be an explicit nil
func (o *PayrollItemStatus) SetProcessingStatusNil() {
	o.ProcessingStatus.Set(nil)
}

// UnsetProcessingStatus ensures that no value is present for ProcessingStatus, not even an explicit nil
func (o *PayrollItemStatus) UnsetProcessingStatus() {
	o.ProcessingStatus.Unset()
}

func (o PayrollItemStatus) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ProcessingStatus.IsSet() {
		toSerialize["processing_status"] = o.ProcessingStatus.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *PayrollItemStatus) UnmarshalJSON(bytes []byte) (err error) {
	varPayrollItemStatus := _PayrollItemStatus{}

	if err = json.Unmarshal(bytes, &varPayrollItemStatus); err == nil {
		*o = PayrollItemStatus(varPayrollItemStatus)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "processing_status")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePayrollItemStatus struct {
	value *PayrollItemStatus
	isSet bool
}

func (v NullablePayrollItemStatus) Get() *PayrollItemStatus {
	return v.value
}

func (v *NullablePayrollItemStatus) Set(val *PayrollItemStatus) {
	v.value = val
	v.isSet = true
}

func (v NullablePayrollItemStatus) IsSet() bool {
	return v.isSet
}

func (v *NullablePayrollItemStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePayrollItemStatus(val *PayrollItemStatus) *NullablePayrollItemStatus {
	return &NullablePayrollItemStatus{value: val, isSet: true}
}

func (v NullablePayrollItemStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePayrollItemStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


