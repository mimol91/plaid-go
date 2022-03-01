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

// Application Metadata about the application
type Application struct {
	// This field will map to the application ID that is returned from /item/applications/list, or provided to the institution in an oauth redirect.
	ApplicationId string `json:"application_id"`
	// The name of the application
	Name string `json:"name"`
	// The date this application was linked in [ISO 8601](https://wikipedia.org/wiki/ISO_8601) (YYYY-MM-DD) format in UTC.
	CreatedAt *string `json:"created_at,omitempty"`
	// The date this application was granted production access at Plaid in [ISO 8601](https://wikipedia.org/wiki/ISO_8601) (YYYY-MM-DD) format in UTC.
	JoinDate string `json:"join_date"`
	// A URL that links to the application logo image.
	LogoUrl NullableString `json:"logo_url"`
	// The URL for the application's website
	ApplicationUrl NullableString `json:"application_url"`
	// A string provided by the connected app stating why they use their respective enabled products.
	ReasonForAccess NullableString `json:"reason_for_access"`
}

// NewApplication instantiates a new Application object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplication(applicationId string, name string, joinDate string, logoUrl NullableString, applicationUrl NullableString, reasonForAccess NullableString) *Application {
	this := Application{}
	this.ApplicationId = applicationId
	this.Name = name
	this.JoinDate = joinDate
	this.LogoUrl = logoUrl
	this.ApplicationUrl = applicationUrl
	this.ReasonForAccess = reasonForAccess
	return &this
}

// NewApplicationWithDefaults instantiates a new Application object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplicationWithDefaults() *Application {
	this := Application{}
	return &this
}

// GetApplicationId returns the ApplicationId field value
func (o *Application) GetApplicationId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ApplicationId
}

// GetApplicationIdOk returns a tuple with the ApplicationId field value
// and a boolean to check if the value has been set.
func (o *Application) GetApplicationIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ApplicationId, true
}

// SetApplicationId sets field value
func (o *Application) SetApplicationId(v string) {
	o.ApplicationId = v
}

// GetName returns the Name field value
func (o *Application) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Application) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Application) SetName(v string) {
	o.Name = v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *Application) GetCreatedAt() string {
	if o == nil || o.CreatedAt == nil {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Application) GetCreatedAtOk() (*string, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *Application) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *Application) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetJoinDate returns the JoinDate field value
func (o *Application) GetJoinDate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.JoinDate
}

// GetJoinDateOk returns a tuple with the JoinDate field value
// and a boolean to check if the value has been set.
func (o *Application) GetJoinDateOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.JoinDate, true
}

// SetJoinDate sets field value
func (o *Application) SetJoinDate(v string) {
	o.JoinDate = v
}

// GetLogoUrl returns the LogoUrl field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Application) GetLogoUrl() string {
	if o == nil || o.LogoUrl.Get() == nil {
		var ret string
		return ret
	}

	return *o.LogoUrl.Get()
}

// GetLogoUrlOk returns a tuple with the LogoUrl field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Application) GetLogoUrlOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.LogoUrl.Get(), o.LogoUrl.IsSet()
}

// SetLogoUrl sets field value
func (o *Application) SetLogoUrl(v string) {
	o.LogoUrl.Set(&v)
}

// GetApplicationUrl returns the ApplicationUrl field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Application) GetApplicationUrl() string {
	if o == nil || o.ApplicationUrl.Get() == nil {
		var ret string
		return ret
	}

	return *o.ApplicationUrl.Get()
}

// GetApplicationUrlOk returns a tuple with the ApplicationUrl field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Application) GetApplicationUrlOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ApplicationUrl.Get(), o.ApplicationUrl.IsSet()
}

// SetApplicationUrl sets field value
func (o *Application) SetApplicationUrl(v string) {
	o.ApplicationUrl.Set(&v)
}

// GetReasonForAccess returns the ReasonForAccess field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Application) GetReasonForAccess() string {
	if o == nil || o.ReasonForAccess.Get() == nil {
		var ret string
		return ret
	}

	return *o.ReasonForAccess.Get()
}

// GetReasonForAccessOk returns a tuple with the ReasonForAccess field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Application) GetReasonForAccessOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ReasonForAccess.Get(), o.ReasonForAccess.IsSet()
}

// SetReasonForAccess sets field value
func (o *Application) SetReasonForAccess(v string) {
	o.ReasonForAccess.Set(&v)
}

func (o Application) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["application_id"] = o.ApplicationId
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.CreatedAt != nil {
		toSerialize["created_at"] = o.CreatedAt
	}
	if true {
		toSerialize["join_date"] = o.JoinDate
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
	return json.Marshal(toSerialize)
}

type NullableApplication struct {
	value *Application
	isSet bool
}

func (v NullableApplication) Get() *Application {
	return v.value
}

func (v *NullableApplication) Set(val *Application) {
	v.value = val
	v.isSet = true
}

func (v NullableApplication) IsSet() bool {
	return v.isSet
}

func (v *NullableApplication) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplication(val *Application) *NullableApplication {
	return &NullableApplication{value: val, isSet: true}
}

func (v NullableApplication) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplication) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


