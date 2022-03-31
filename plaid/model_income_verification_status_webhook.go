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

// IncomeVerificationStatusWebhook Fired when the status of an income verification instance has changed. It will typically take several minutes for this webhook to fire after the end user has uploaded their documents in the Document Income flow.
type IncomeVerificationStatusWebhook struct {
	// `\"INCOME\"`
	WebhookType string `json:"webhook_type"`
	// `INCOME_VERIFICATION`
	WebhookCode string `json:"webhook_code"`
	// The Item ID associated with the verification.
	ItemId string `json:"item_id"`
	// `VERIFICATION_STATUS_PROCESSING_COMPLETE`: The income verification status processing has completed. If the user uploaded multiple documents, this webhook will fire when all documents have finished processing. Call the `/income/verification/paystubs/get` endpoint and check the document metadata to see which documents were successfully parsed.  `VERIFICATION_STATUS_PROCESSING_FAILED`: A failure occurred when attempting to process the verification documentation.  `VERIFICATION_STATUS_PENDING_APPROVAL`: The income verification has been sent to the user for review.
	VerificationStatus string `json:"verification_status"`
	AdditionalProperties map[string]interface{}
}

type _IncomeVerificationStatusWebhook IncomeVerificationStatusWebhook

// NewIncomeVerificationStatusWebhook instantiates a new IncomeVerificationStatusWebhook object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIncomeVerificationStatusWebhook(webhookType string, webhookCode string, itemId string, verificationStatus string) *IncomeVerificationStatusWebhook {
	this := IncomeVerificationStatusWebhook{}
	this.WebhookType = webhookType
	this.WebhookCode = webhookCode
	this.ItemId = itemId
	this.VerificationStatus = verificationStatus
	return &this
}

// NewIncomeVerificationStatusWebhookWithDefaults instantiates a new IncomeVerificationStatusWebhook object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIncomeVerificationStatusWebhookWithDefaults() *IncomeVerificationStatusWebhook {
	this := IncomeVerificationStatusWebhook{}
	return &this
}

// GetWebhookType returns the WebhookType field value
func (o *IncomeVerificationStatusWebhook) GetWebhookType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WebhookType
}

// GetWebhookTypeOk returns a tuple with the WebhookType field value
// and a boolean to check if the value has been set.
func (o *IncomeVerificationStatusWebhook) GetWebhookTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.WebhookType, true
}

// SetWebhookType sets field value
func (o *IncomeVerificationStatusWebhook) SetWebhookType(v string) {
	o.WebhookType = v
}

// GetWebhookCode returns the WebhookCode field value
func (o *IncomeVerificationStatusWebhook) GetWebhookCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WebhookCode
}

// GetWebhookCodeOk returns a tuple with the WebhookCode field value
// and a boolean to check if the value has been set.
func (o *IncomeVerificationStatusWebhook) GetWebhookCodeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.WebhookCode, true
}

// SetWebhookCode sets field value
func (o *IncomeVerificationStatusWebhook) SetWebhookCode(v string) {
	o.WebhookCode = v
}

// GetItemId returns the ItemId field value
func (o *IncomeVerificationStatusWebhook) GetItemId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ItemId
}

// GetItemIdOk returns a tuple with the ItemId field value
// and a boolean to check if the value has been set.
func (o *IncomeVerificationStatusWebhook) GetItemIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ItemId, true
}

// SetItemId sets field value
func (o *IncomeVerificationStatusWebhook) SetItemId(v string) {
	o.ItemId = v
}

// GetVerificationStatus returns the VerificationStatus field value
func (o *IncomeVerificationStatusWebhook) GetVerificationStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VerificationStatus
}

// GetVerificationStatusOk returns a tuple with the VerificationStatus field value
// and a boolean to check if the value has been set.
func (o *IncomeVerificationStatusWebhook) GetVerificationStatusOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.VerificationStatus, true
}

// SetVerificationStatus sets field value
func (o *IncomeVerificationStatusWebhook) SetVerificationStatus(v string) {
	o.VerificationStatus = v
}

func (o IncomeVerificationStatusWebhook) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["webhook_type"] = o.WebhookType
	}
	if true {
		toSerialize["webhook_code"] = o.WebhookCode
	}
	if true {
		toSerialize["item_id"] = o.ItemId
	}
	if true {
		toSerialize["verification_status"] = o.VerificationStatus
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *IncomeVerificationStatusWebhook) UnmarshalJSON(bytes []byte) (err error) {
	varIncomeVerificationStatusWebhook := _IncomeVerificationStatusWebhook{}

	if err = json.Unmarshal(bytes, &varIncomeVerificationStatusWebhook); err == nil {
		*o = IncomeVerificationStatusWebhook(varIncomeVerificationStatusWebhook)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "webhook_type")
		delete(additionalProperties, "webhook_code")
		delete(additionalProperties, "item_id")
		delete(additionalProperties, "verification_status")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableIncomeVerificationStatusWebhook struct {
	value *IncomeVerificationStatusWebhook
	isSet bool
}

func (v NullableIncomeVerificationStatusWebhook) Get() *IncomeVerificationStatusWebhook {
	return v.value
}

func (v *NullableIncomeVerificationStatusWebhook) Set(val *IncomeVerificationStatusWebhook) {
	v.value = val
	v.isSet = true
}

func (v NullableIncomeVerificationStatusWebhook) IsSet() bool {
	return v.isSet
}

func (v *NullableIncomeVerificationStatusWebhook) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIncomeVerificationStatusWebhook(val *IncomeVerificationStatusWebhook) *NullableIncomeVerificationStatusWebhook {
	return &NullableIncomeVerificationStatusWebhook{value: val, isSet: true}
}

func (v NullableIncomeVerificationStatusWebhook) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIncomeVerificationStatusWebhook) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


