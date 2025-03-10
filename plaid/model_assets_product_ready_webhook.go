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

// AssetsProductReadyWebhook Fired when the Asset Report has been generated and `/asset_report/get` is ready to be called.  If you attempt to retrieve an Asset Report before this webhook has fired, you’ll receive a response with the HTTP status code 400 and a Plaid error code of `PRODUCT_NOT_READY`.
type AssetsProductReadyWebhook struct {
	// `ASSETS`
	WebhookType string `json:"webhook_type"`
	// `PRODUCT_READY`
	WebhookCode string `json:"webhook_code"`
	// The `asset_report_id` that can be provided to `/asset_report/get` to retrieve the Asset Report.
	AssetReportId string `json:"asset_report_id"`
	AdditionalProperties map[string]interface{}
}

type _AssetsProductReadyWebhook AssetsProductReadyWebhook

// NewAssetsProductReadyWebhook instantiates a new AssetsProductReadyWebhook object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAssetsProductReadyWebhook(webhookType string, webhookCode string, assetReportId string) *AssetsProductReadyWebhook {
	this := AssetsProductReadyWebhook{}
	this.WebhookType = webhookType
	this.WebhookCode = webhookCode
	this.AssetReportId = assetReportId
	return &this
}

// NewAssetsProductReadyWebhookWithDefaults instantiates a new AssetsProductReadyWebhook object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssetsProductReadyWebhookWithDefaults() *AssetsProductReadyWebhook {
	this := AssetsProductReadyWebhook{}
	return &this
}

// GetWebhookType returns the WebhookType field value
func (o *AssetsProductReadyWebhook) GetWebhookType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WebhookType
}

// GetWebhookTypeOk returns a tuple with the WebhookType field value
// and a boolean to check if the value has been set.
func (o *AssetsProductReadyWebhook) GetWebhookTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.WebhookType, true
}

// SetWebhookType sets field value
func (o *AssetsProductReadyWebhook) SetWebhookType(v string) {
	o.WebhookType = v
}

// GetWebhookCode returns the WebhookCode field value
func (o *AssetsProductReadyWebhook) GetWebhookCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WebhookCode
}

// GetWebhookCodeOk returns a tuple with the WebhookCode field value
// and a boolean to check if the value has been set.
func (o *AssetsProductReadyWebhook) GetWebhookCodeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.WebhookCode, true
}

// SetWebhookCode sets field value
func (o *AssetsProductReadyWebhook) SetWebhookCode(v string) {
	o.WebhookCode = v
}

// GetAssetReportId returns the AssetReportId field value
func (o *AssetsProductReadyWebhook) GetAssetReportId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AssetReportId
}

// GetAssetReportIdOk returns a tuple with the AssetReportId field value
// and a boolean to check if the value has been set.
func (o *AssetsProductReadyWebhook) GetAssetReportIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AssetReportId, true
}

// SetAssetReportId sets field value
func (o *AssetsProductReadyWebhook) SetAssetReportId(v string) {
	o.AssetReportId = v
}

func (o AssetsProductReadyWebhook) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["webhook_type"] = o.WebhookType
	}
	if true {
		toSerialize["webhook_code"] = o.WebhookCode
	}
	if true {
		toSerialize["asset_report_id"] = o.AssetReportId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *AssetsProductReadyWebhook) UnmarshalJSON(bytes []byte) (err error) {
	varAssetsProductReadyWebhook := _AssetsProductReadyWebhook{}

	if err = json.Unmarshal(bytes, &varAssetsProductReadyWebhook); err == nil {
		*o = AssetsProductReadyWebhook(varAssetsProductReadyWebhook)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "webhook_type")
		delete(additionalProperties, "webhook_code")
		delete(additionalProperties, "asset_report_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAssetsProductReadyWebhook struct {
	value *AssetsProductReadyWebhook
	isSet bool
}

func (v NullableAssetsProductReadyWebhook) Get() *AssetsProductReadyWebhook {
	return v.value
}

func (v *NullableAssetsProductReadyWebhook) Set(val *AssetsProductReadyWebhook) {
	v.value = val
	v.isSet = true
}

func (v NullableAssetsProductReadyWebhook) IsSet() bool {
	return v.isSet
}

func (v *NullableAssetsProductReadyWebhook) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssetsProductReadyWebhook(val *AssetsProductReadyWebhook) *NullableAssetsProductReadyWebhook {
	return &NullableAssetsProductReadyWebhook{value: val, isSet: true}
}

func (v NullableAssetsProductReadyWebhook) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssetsProductReadyWebhook) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


