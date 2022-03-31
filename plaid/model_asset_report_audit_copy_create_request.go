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

// AssetReportAuditCopyCreateRequest AssetReportAuditCopyCreateRequest defines the request schema for `/asset_report/audit_copy/get`
type AssetReportAuditCopyCreateRequest struct {
	// Your Plaid API `client_id`. The `client_id` is required and may be provided either in the `PLAID-CLIENT-ID` header or as part of a request body.
	ClientId *string `json:"client_id,omitempty"`
	// Your Plaid API `secret`. The `secret` is required and may be provided either in the `PLAID-SECRET` header or as part of a request body.
	Secret *string `json:"secret,omitempty"`
	// A token that can be provided to endpoints such as `/asset_report/get` or `/asset_report/pdf/get` to fetch or update an Asset Report.
	AssetReportToken string `json:"asset_report_token"`
	// The `auditor_id` of the third party with whom you would like to share the Asset Report.
	AuditorId string `json:"auditor_id"`
}

// NewAssetReportAuditCopyCreateRequest instantiates a new AssetReportAuditCopyCreateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAssetReportAuditCopyCreateRequest(assetReportToken string, auditorId string) *AssetReportAuditCopyCreateRequest {
	this := AssetReportAuditCopyCreateRequest{}
	this.AssetReportToken = assetReportToken
	this.AuditorId = auditorId
	return &this
}

// NewAssetReportAuditCopyCreateRequestWithDefaults instantiates a new AssetReportAuditCopyCreateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssetReportAuditCopyCreateRequestWithDefaults() *AssetReportAuditCopyCreateRequest {
	this := AssetReportAuditCopyCreateRequest{}
	return &this
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *AssetReportAuditCopyCreateRequest) GetClientId() string {
	if o == nil || o.ClientId == nil {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetReportAuditCopyCreateRequest) GetClientIdOk() (*string, bool) {
	if o == nil || o.ClientId == nil {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *AssetReportAuditCopyCreateRequest) HasClientId() bool {
	if o != nil && o.ClientId != nil {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *AssetReportAuditCopyCreateRequest) SetClientId(v string) {
	o.ClientId = &v
}

// GetSecret returns the Secret field value if set, zero value otherwise.
func (o *AssetReportAuditCopyCreateRequest) GetSecret() string {
	if o == nil || o.Secret == nil {
		var ret string
		return ret
	}
	return *o.Secret
}

// GetSecretOk returns a tuple with the Secret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetReportAuditCopyCreateRequest) GetSecretOk() (*string, bool) {
	if o == nil || o.Secret == nil {
		return nil, false
	}
	return o.Secret, true
}

// HasSecret returns a boolean if a field has been set.
func (o *AssetReportAuditCopyCreateRequest) HasSecret() bool {
	if o != nil && o.Secret != nil {
		return true
	}

	return false
}

// SetSecret gets a reference to the given string and assigns it to the Secret field.
func (o *AssetReportAuditCopyCreateRequest) SetSecret(v string) {
	o.Secret = &v
}

// GetAssetReportToken returns the AssetReportToken field value
func (o *AssetReportAuditCopyCreateRequest) GetAssetReportToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AssetReportToken
}

// GetAssetReportTokenOk returns a tuple with the AssetReportToken field value
// and a boolean to check if the value has been set.
func (o *AssetReportAuditCopyCreateRequest) GetAssetReportTokenOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AssetReportToken, true
}

// SetAssetReportToken sets field value
func (o *AssetReportAuditCopyCreateRequest) SetAssetReportToken(v string) {
	o.AssetReportToken = v
}

// GetAuditorId returns the AuditorId field value
func (o *AssetReportAuditCopyCreateRequest) GetAuditorId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AuditorId
}

// GetAuditorIdOk returns a tuple with the AuditorId field value
// and a boolean to check if the value has been set.
func (o *AssetReportAuditCopyCreateRequest) GetAuditorIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AuditorId, true
}

// SetAuditorId sets field value
func (o *AssetReportAuditCopyCreateRequest) SetAuditorId(v string) {
	o.AuditorId = v
}

func (o AssetReportAuditCopyCreateRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ClientId != nil {
		toSerialize["client_id"] = o.ClientId
	}
	if o.Secret != nil {
		toSerialize["secret"] = o.Secret
	}
	if true {
		toSerialize["asset_report_token"] = o.AssetReportToken
	}
	if true {
		toSerialize["auditor_id"] = o.AuditorId
	}
	return json.Marshal(toSerialize)
}

type NullableAssetReportAuditCopyCreateRequest struct {
	value *AssetReportAuditCopyCreateRequest
	isSet bool
}

func (v NullableAssetReportAuditCopyCreateRequest) Get() *AssetReportAuditCopyCreateRequest {
	return v.value
}

func (v *NullableAssetReportAuditCopyCreateRequest) Set(val *AssetReportAuditCopyCreateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAssetReportAuditCopyCreateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAssetReportAuditCopyCreateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssetReportAuditCopyCreateRequest(val *AssetReportAuditCopyCreateRequest) *NullableAssetReportAuditCopyCreateRequest {
	return &NullableAssetReportAuditCopyCreateRequest{value: val, isSet: true}
}

func (v NullableAssetReportAuditCopyCreateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssetReportAuditCopyCreateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


