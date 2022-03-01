/*
 * The Plaid API
 *
 * The Plaid REST API. Please see https://plaid.com/docs/api for more details.
 *
 * API version: 2020-09-14_1.78.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package plaid

import (
	"encoding/json"
)

// ConnectedApplication Describes the connected application for a particular end user.
type ConnectedApplication struct {
	// This field will map to the application ID that is returned from /item/applications/list, or provided to the institution in an oauth redirect.
	ApplicationId string `json:"application_id"`
	// The name of the application
	Name string `json:"name"`
	// A URL that links to the application logo image (will be deprecated in the future, please use logo_url).
	Logo NullableString `json:"logo"`
	// A URL that links to the application logo image.
	LogoUrl NullableString `json:"logo_url"`
	// The URL for the application's website
	ApplicationUrl NullableString `json:"application_url"`
	// A string provided by the connected app stating why they use their respective enabled products.
	ReasonForAccess NullableString `json:"reason_for_access"`
	// The date this application was linked in [ISO 8601](https://wikipedia.org/wiki/ISO_8601) (YYYY-MM-DD) format in UTC.
	CreatedAt string `json:"created_at"`
	// The date this application was granted production access at Plaid in [ISO 8601](https://wikipedia.org/wiki/ISO_8601) (YYYY-MM-DD) format in UTC.
	JoinDate string `json:"join_date"`
	// (Deprecated) A list of enums representing the data collected and products enabled for this connected application.
	ProductDataTypes []string `json:"product_data_types"`
	Scopes NullableScopesNullable `json:"scopes,omitempty"`
	RequestedScopes *RequestedScopes `json:"requested_scopes,omitempty"`
}

// NewConnectedApplication instantiates a new ConnectedApplication object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConnectedApplication(applicationId string, name string, logo NullableString, logoUrl NullableString, applicationUrl NullableString, reasonForAccess NullableString, createdAt string, joinDate string, productDataTypes []string) *ConnectedApplication {
	this := ConnectedApplication{}
	this.ApplicationId = applicationId
	this.Name = name
	this.Logo = logo
	this.LogoUrl = logoUrl
	this.ApplicationUrl = applicationUrl
	this.ReasonForAccess = reasonForAccess
	this.CreatedAt = createdAt
	this.JoinDate = joinDate
	this.ProductDataTypes = productDataTypes
	return &this
}

// NewConnectedApplicationWithDefaults instantiates a new ConnectedApplication object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConnectedApplicationWithDefaults() *ConnectedApplication {
	this := ConnectedApplication{}
	return &this
}

// GetApplicationId returns the ApplicationId field value
func (o *ConnectedApplication) GetApplicationId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ApplicationId
}

// GetApplicationIdOk returns a tuple with the ApplicationId field value
// and a boolean to check if the value has been set.
func (o *ConnectedApplication) GetApplicationIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ApplicationId, true
}

// SetApplicationId sets field value
func (o *ConnectedApplication) SetApplicationId(v string) {
	o.ApplicationId = v
}

// GetName returns the Name field value
func (o *ConnectedApplication) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ConnectedApplication) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ConnectedApplication) SetName(v string) {
	o.Name = v
}

// GetLogo returns the Logo field value
// If the value is explicit nil, the zero value for string will be returned
func (o *ConnectedApplication) GetLogo() string {
	if o == nil || o.Logo.Get() == nil {
		var ret string
		return ret
	}

	return *o.Logo.Get()
}

// GetLogoOk returns a tuple with the Logo field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ConnectedApplication) GetLogoOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Logo.Get(), o.Logo.IsSet()
}

// SetLogo sets field value
func (o *ConnectedApplication) SetLogo(v string) {
	o.Logo.Set(&v)
}

// GetLogoUrl returns the LogoUrl field value
// If the value is explicit nil, the zero value for string will be returned
func (o *ConnectedApplication) GetLogoUrl() string {
	if o == nil || o.LogoUrl.Get() == nil {
		var ret string
		return ret
	}

	return *o.LogoUrl.Get()
}

// GetLogoUrlOk returns a tuple with the LogoUrl field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ConnectedApplication) GetLogoUrlOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.LogoUrl.Get(), o.LogoUrl.IsSet()
}

// SetLogoUrl sets field value
func (o *ConnectedApplication) SetLogoUrl(v string) {
	o.LogoUrl.Set(&v)
}

// GetApplicationUrl returns the ApplicationUrl field value
// If the value is explicit nil, the zero value for string will be returned
func (o *ConnectedApplication) GetApplicationUrl() string {
	if o == nil || o.ApplicationUrl.Get() == nil {
		var ret string
		return ret
	}

	return *o.ApplicationUrl.Get()
}

// GetApplicationUrlOk returns a tuple with the ApplicationUrl field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ConnectedApplication) GetApplicationUrlOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ApplicationUrl.Get(), o.ApplicationUrl.IsSet()
}

// SetApplicationUrl sets field value
func (o *ConnectedApplication) SetApplicationUrl(v string) {
	o.ApplicationUrl.Set(&v)
}

// GetReasonForAccess returns the ReasonForAccess field value
// If the value is explicit nil, the zero value for string will be returned
func (o *ConnectedApplication) GetReasonForAccess() string {
	if o == nil || o.ReasonForAccess.Get() == nil {
		var ret string
		return ret
	}

	return *o.ReasonForAccess.Get()
}

// GetReasonForAccessOk returns a tuple with the ReasonForAccess field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ConnectedApplication) GetReasonForAccessOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ReasonForAccess.Get(), o.ReasonForAccess.IsSet()
}

// SetReasonForAccess sets field value
func (o *ConnectedApplication) SetReasonForAccess(v string) {
	o.ReasonForAccess.Set(&v)
}

// GetCreatedAt returns the CreatedAt field value
func (o *ConnectedApplication) GetCreatedAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *ConnectedApplication) GetCreatedAtOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *ConnectedApplication) SetCreatedAt(v string) {
	o.CreatedAt = v
}

// GetJoinDate returns the JoinDate field value
func (o *ConnectedApplication) GetJoinDate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.JoinDate
}

// GetJoinDateOk returns a tuple with the JoinDate field value
// and a boolean to check if the value has been set.
func (o *ConnectedApplication) GetJoinDateOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.JoinDate, true
}

// SetJoinDate sets field value
func (o *ConnectedApplication) SetJoinDate(v string) {
	o.JoinDate = v
}

// GetProductDataTypes returns the ProductDataTypes field value
func (o *ConnectedApplication) GetProductDataTypes() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.ProductDataTypes
}

// GetProductDataTypesOk returns a tuple with the ProductDataTypes field value
// and a boolean to check if the value has been set.
func (o *ConnectedApplication) GetProductDataTypesOk() (*[]string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ProductDataTypes, true
}

// SetProductDataTypes sets field value
func (o *ConnectedApplication) SetProductDataTypes(v []string) {
	o.ProductDataTypes = v
}

// GetScopes returns the Scopes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ConnectedApplication) GetScopes() ScopesNullable {
	if o == nil || o.Scopes.Get() == nil {
		var ret ScopesNullable
		return ret
	}
	return *o.Scopes.Get()
}

// GetScopesOk returns a tuple with the Scopes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ConnectedApplication) GetScopesOk() (*ScopesNullable, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Scopes.Get(), o.Scopes.IsSet()
}

// HasScopes returns a boolean if a field has been set.
func (o *ConnectedApplication) HasScopes() bool {
	if o != nil && o.Scopes.IsSet() {
		return true
	}

	return false
}

// SetScopes gets a reference to the given NullableScopesNullable and assigns it to the Scopes field.
func (o *ConnectedApplication) SetScopes(v ScopesNullable) {
	o.Scopes.Set(&v)
}
// SetScopesNil sets the value for Scopes to be an explicit nil
func (o *ConnectedApplication) SetScopesNil() {
	o.Scopes.Set(nil)
}

// UnsetScopes ensures that no value is present for Scopes, not even an explicit nil
func (o *ConnectedApplication) UnsetScopes() {
	o.Scopes.Unset()
}

// GetRequestedScopes returns the RequestedScopes field value if set, zero value otherwise.
func (o *ConnectedApplication) GetRequestedScopes() RequestedScopes {
	if o == nil || o.RequestedScopes == nil {
		var ret RequestedScopes
		return ret
	}
	return *o.RequestedScopes
}

// GetRequestedScopesOk returns a tuple with the RequestedScopes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectedApplication) GetRequestedScopesOk() (*RequestedScopes, bool) {
	if o == nil || o.RequestedScopes == nil {
		return nil, false
	}
	return o.RequestedScopes, true
}

// HasRequestedScopes returns a boolean if a field has been set.
func (o *ConnectedApplication) HasRequestedScopes() bool {
	if o != nil && o.RequestedScopes != nil {
		return true
	}

	return false
}

// SetRequestedScopes gets a reference to the given RequestedScopes and assigns it to the RequestedScopes field.
func (o *ConnectedApplication) SetRequestedScopes(v RequestedScopes) {
	o.RequestedScopes = &v
}

func (o ConnectedApplication) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["application_id"] = o.ApplicationId
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["logo"] = o.Logo.Get()
	}
	if true {
		toSerialize["logo_url"] = o.LogoUrl.Get()
	}
	if true {
		toSerialize["application_url"] = o.ApplicationUrl.Get()
	}
	if true {
		toSerialize["reason_for_access"] = o.ReasonForAccess.Get()
	}
	if true {
		toSerialize["created_at"] = o.CreatedAt
	}
	if true {
		toSerialize["join_date"] = o.JoinDate
	}
	if true {
		toSerialize["product_data_types"] = o.ProductDataTypes
	}
	if o.Scopes.IsSet() {
		toSerialize["scopes"] = o.Scopes.Get()
	}
	if o.RequestedScopes != nil {
		toSerialize["requested_scopes"] = o.RequestedScopes
	}
	return json.Marshal(toSerialize)
}

type NullableConnectedApplication struct {
	value *ConnectedApplication
	isSet bool
}

func (v NullableConnectedApplication) Get() *ConnectedApplication {
	return v.value
}

func (v *NullableConnectedApplication) Set(val *ConnectedApplication) {
	v.value = val
	v.isSet = true
}

func (v NullableConnectedApplication) IsSet() bool {
	return v.isSet
}

func (v *NullableConnectedApplication) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConnectedApplication(val *ConnectedApplication) *NullableConnectedApplication {
	return &NullableConnectedApplication{value: val, isSet: true}
}

func (v NullableConnectedApplication) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConnectedApplication) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


