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

// AssetReportAuditCopyRemoveResponse AssetReportAuditCopyRemoveResponse defines the response schema for `/asset_report/audit_copy/remove`
type AssetReportAuditCopyRemoveResponse struct {
	// `true` if the Audit Copy was successfully removed.
	Removed bool `json:"removed"`
	// A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive.
	RequestId string `json:"request_id"`
	AdditionalProperties map[string]interface{}
}

type _AssetReportAuditCopyRemoveResponse AssetReportAuditCopyRemoveResponse

// NewAssetReportAuditCopyRemoveResponse instantiates a new AssetReportAuditCopyRemoveResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAssetReportAuditCopyRemoveResponse(removed bool, requestId string) *AssetReportAuditCopyRemoveResponse {
	this := AssetReportAuditCopyRemoveResponse{}
	this.Removed = removed
	this.RequestId = requestId
	return &this
}

// NewAssetReportAuditCopyRemoveResponseWithDefaults instantiates a new AssetReportAuditCopyRemoveResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssetReportAuditCopyRemoveResponseWithDefaults() *AssetReportAuditCopyRemoveResponse {
	this := AssetReportAuditCopyRemoveResponse{}
	return &this
}

// GetRemoved returns the Removed field value
func (o *AssetReportAuditCopyRemoveResponse) GetRemoved() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Removed
}

// GetRemovedOk returns a tuple with the Removed field value
// and a boolean to check if the value has been set.
func (o *AssetReportAuditCopyRemoveResponse) GetRemovedOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Removed, true
}

// SetRemoved sets field value
func (o *AssetReportAuditCopyRemoveResponse) SetRemoved(v bool) {
	o.Removed = v
}

// GetRequestId returns the RequestId field value
func (o *AssetReportAuditCopyRemoveResponse) GetRequestId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value
// and a boolean to check if the value has been set.
func (o *AssetReportAuditCopyRemoveResponse) GetRequestIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RequestId, true
}

// SetRequestId sets field value
func (o *AssetReportAuditCopyRemoveResponse) SetRequestId(v string) {
	o.RequestId = v
}

func (o AssetReportAuditCopyRemoveResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["removed"] = o.Removed
	}
	if true {
		toSerialize["request_id"] = o.RequestId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *AssetReportAuditCopyRemoveResponse) UnmarshalJSON(bytes []byte) (err error) {
	varAssetReportAuditCopyRemoveResponse := _AssetReportAuditCopyRemoveResponse{}

	if err = json.Unmarshal(bytes, &varAssetReportAuditCopyRemoveResponse); err == nil {
		*o = AssetReportAuditCopyRemoveResponse(varAssetReportAuditCopyRemoveResponse)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "removed")
		delete(additionalProperties, "request_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAssetReportAuditCopyRemoveResponse struct {
	value *AssetReportAuditCopyRemoveResponse
	isSet bool
}

func (v NullableAssetReportAuditCopyRemoveResponse) Get() *AssetReportAuditCopyRemoveResponse {
	return v.value
}

func (v *NullableAssetReportAuditCopyRemoveResponse) Set(val *AssetReportAuditCopyRemoveResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAssetReportAuditCopyRemoveResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAssetReportAuditCopyRemoveResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssetReportAuditCopyRemoveResponse(val *AssetReportAuditCopyRemoveResponse) *NullableAssetReportAuditCopyRemoveResponse {
	return &NullableAssetReportAuditCopyRemoveResponse{value: val, isSet: true}
}

func (v NullableAssetReportAuditCopyRemoveResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssetReportAuditCopyRemoveResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


