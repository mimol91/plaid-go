/*
 * The Plaid API
 *
 * The Plaid REST API. Please see https://plaid.com/docs/api for more details.
 *
 * API version: 2020-09-14_1.77.1
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package plaid

import (
	"encoding/json"
)

// InstitutionsGetRequestOptions An optional object to filter `/institutions/get` results.
type InstitutionsGetRequestOptions struct {
	// Filter the Institutions based on which products they support. 
	Products []Products `json:"products,omitempty"`
	// Specify an array of routing numbers to filter institutions. The response will only return institutions that match all of the routing numbers in the array. Routing number records used for this matching are not comprehensive; failure to match a given routing number to an institution does not mean that the institution is unsupported by Plaid.
	RoutingNumbers []string `json:"routing_numbers,omitempty"`
	// Limit results to institutions with or without OAuth login flows.
	Oauth NullableBool `json:"oauth,omitempty"`
	// When `true`, return the institution's homepage URL, logo and primary brand color.  Note that Plaid does not own any of the logos shared by the API, and that by accessing or using these logos, you agree that you are doing so at your own risk and will, if necessary, obtain all required permissions from the appropriate rights holders and adhere to any applicable usage guidelines. Plaid disclaims all express or implied warranties with respect to the logos.
	IncludeOptionalMetadata *bool `json:"include_optional_metadata,omitempty"`
	// When `true`, returns metadata related to the Auth product indicating which auth methods are supported.
	IncludeAuthMetadata *bool `json:"include_auth_metadata,omitempty"`
	// When `true`, returns metadata related to the Payment Initiation product indicating which payment configurations are supported.
	IncludePaymentInitiationMetadata *bool `json:"include_payment_initiation_metadata,omitempty"`
}

// NewInstitutionsGetRequestOptions instantiates a new InstitutionsGetRequestOptions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInstitutionsGetRequestOptions() *InstitutionsGetRequestOptions {
	this := InstitutionsGetRequestOptions{}
	var includeAuthMetadata bool = false
	this.IncludeAuthMetadata = &includeAuthMetadata
	var includePaymentInitiationMetadata bool = false
	this.IncludePaymentInitiationMetadata = &includePaymentInitiationMetadata
	return &this
}

// NewInstitutionsGetRequestOptionsWithDefaults instantiates a new InstitutionsGetRequestOptions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInstitutionsGetRequestOptionsWithDefaults() *InstitutionsGetRequestOptions {
	this := InstitutionsGetRequestOptions{}
	var includeAuthMetadata bool = false
	this.IncludeAuthMetadata = &includeAuthMetadata
	var includePaymentInitiationMetadata bool = false
	this.IncludePaymentInitiationMetadata = &includePaymentInitiationMetadata
	return &this
}

// GetProducts returns the Products field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InstitutionsGetRequestOptions) GetProducts() []Products {
	if o == nil  {
		var ret []Products
		return ret
	}
	return o.Products
}

// GetProductsOk returns a tuple with the Products field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InstitutionsGetRequestOptions) GetProductsOk() (*[]Products, bool) {
	if o == nil || o.Products == nil {
		return nil, false
	}
	return &o.Products, true
}

// HasProducts returns a boolean if a field has been set.
func (o *InstitutionsGetRequestOptions) HasProducts() bool {
	if o != nil && o.Products != nil {
		return true
	}

	return false
}

// SetProducts gets a reference to the given []Products and assigns it to the Products field.
func (o *InstitutionsGetRequestOptions) SetProducts(v []Products) {
	o.Products = v
}

// GetRoutingNumbers returns the RoutingNumbers field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InstitutionsGetRequestOptions) GetRoutingNumbers() []string {
	if o == nil  {
		var ret []string
		return ret
	}
	return o.RoutingNumbers
}

// GetRoutingNumbersOk returns a tuple with the RoutingNumbers field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InstitutionsGetRequestOptions) GetRoutingNumbersOk() (*[]string, bool) {
	if o == nil || o.RoutingNumbers == nil {
		return nil, false
	}
	return &o.RoutingNumbers, true
}

// HasRoutingNumbers returns a boolean if a field has been set.
func (o *InstitutionsGetRequestOptions) HasRoutingNumbers() bool {
	if o != nil && o.RoutingNumbers != nil {
		return true
	}

	return false
}

// SetRoutingNumbers gets a reference to the given []string and assigns it to the RoutingNumbers field.
func (o *InstitutionsGetRequestOptions) SetRoutingNumbers(v []string) {
	o.RoutingNumbers = v
}

// GetOauth returns the Oauth field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InstitutionsGetRequestOptions) GetOauth() bool {
	if o == nil || o.Oauth.Get() == nil {
		var ret bool
		return ret
	}
	return *o.Oauth.Get()
}

// GetOauthOk returns a tuple with the Oauth field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InstitutionsGetRequestOptions) GetOauthOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Oauth.Get(), o.Oauth.IsSet()
}

// HasOauth returns a boolean if a field has been set.
func (o *InstitutionsGetRequestOptions) HasOauth() bool {
	if o != nil && o.Oauth.IsSet() {
		return true
	}

	return false
}

// SetOauth gets a reference to the given NullableBool and assigns it to the Oauth field.
func (o *InstitutionsGetRequestOptions) SetOauth(v bool) {
	o.Oauth.Set(&v)
}
// SetOauthNil sets the value for Oauth to be an explicit nil
func (o *InstitutionsGetRequestOptions) SetOauthNil() {
	o.Oauth.Set(nil)
}

// UnsetOauth ensures that no value is present for Oauth, not even an explicit nil
func (o *InstitutionsGetRequestOptions) UnsetOauth() {
	o.Oauth.Unset()
}

// GetIncludeOptionalMetadata returns the IncludeOptionalMetadata field value if set, zero value otherwise.
func (o *InstitutionsGetRequestOptions) GetIncludeOptionalMetadata() bool {
	if o == nil || o.IncludeOptionalMetadata == nil {
		var ret bool
		return ret
	}
	return *o.IncludeOptionalMetadata
}

// GetIncludeOptionalMetadataOk returns a tuple with the IncludeOptionalMetadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstitutionsGetRequestOptions) GetIncludeOptionalMetadataOk() (*bool, bool) {
	if o == nil || o.IncludeOptionalMetadata == nil {
		return nil, false
	}
	return o.IncludeOptionalMetadata, true
}

// HasIncludeOptionalMetadata returns a boolean if a field has been set.
func (o *InstitutionsGetRequestOptions) HasIncludeOptionalMetadata() bool {
	if o != nil && o.IncludeOptionalMetadata != nil {
		return true
	}

	return false
}

// SetIncludeOptionalMetadata gets a reference to the given bool and assigns it to the IncludeOptionalMetadata field.
func (o *InstitutionsGetRequestOptions) SetIncludeOptionalMetadata(v bool) {
	o.IncludeOptionalMetadata = &v
}

// GetIncludeAuthMetadata returns the IncludeAuthMetadata field value if set, zero value otherwise.
func (o *InstitutionsGetRequestOptions) GetIncludeAuthMetadata() bool {
	if o == nil || o.IncludeAuthMetadata == nil {
		var ret bool
		return ret
	}
	return *o.IncludeAuthMetadata
}

// GetIncludeAuthMetadataOk returns a tuple with the IncludeAuthMetadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstitutionsGetRequestOptions) GetIncludeAuthMetadataOk() (*bool, bool) {
	if o == nil || o.IncludeAuthMetadata == nil {
		return nil, false
	}
	return o.IncludeAuthMetadata, true
}

// HasIncludeAuthMetadata returns a boolean if a field has been set.
func (o *InstitutionsGetRequestOptions) HasIncludeAuthMetadata() bool {
	if o != nil && o.IncludeAuthMetadata != nil {
		return true
	}

	return false
}

// SetIncludeAuthMetadata gets a reference to the given bool and assigns it to the IncludeAuthMetadata field.
func (o *InstitutionsGetRequestOptions) SetIncludeAuthMetadata(v bool) {
	o.IncludeAuthMetadata = &v
}

// GetIncludePaymentInitiationMetadata returns the IncludePaymentInitiationMetadata field value if set, zero value otherwise.
func (o *InstitutionsGetRequestOptions) GetIncludePaymentInitiationMetadata() bool {
	if o == nil || o.IncludePaymentInitiationMetadata == nil {
		var ret bool
		return ret
	}
	return *o.IncludePaymentInitiationMetadata
}

// GetIncludePaymentInitiationMetadataOk returns a tuple with the IncludePaymentInitiationMetadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstitutionsGetRequestOptions) GetIncludePaymentInitiationMetadataOk() (*bool, bool) {
	if o == nil || o.IncludePaymentInitiationMetadata == nil {
		return nil, false
	}
	return o.IncludePaymentInitiationMetadata, true
}

// HasIncludePaymentInitiationMetadata returns a boolean if a field has been set.
func (o *InstitutionsGetRequestOptions) HasIncludePaymentInitiationMetadata() bool {
	if o != nil && o.IncludePaymentInitiationMetadata != nil {
		return true
	}

	return false
}

// SetIncludePaymentInitiationMetadata gets a reference to the given bool and assigns it to the IncludePaymentInitiationMetadata field.
func (o *InstitutionsGetRequestOptions) SetIncludePaymentInitiationMetadata(v bool) {
	o.IncludePaymentInitiationMetadata = &v
}

func (o InstitutionsGetRequestOptions) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Products != nil {
		toSerialize["products"] = o.Products
	}
	if o.RoutingNumbers != nil {
		toSerialize["routing_numbers"] = o.RoutingNumbers
	}
	if o.Oauth.IsSet() {
		toSerialize["oauth"] = o.Oauth.Get()
	}
	if o.IncludeOptionalMetadata != nil {
		toSerialize["include_optional_metadata"] = o.IncludeOptionalMetadata
	}
	if o.IncludeAuthMetadata != nil {
		toSerialize["include_auth_metadata"] = o.IncludeAuthMetadata
	}
	if o.IncludePaymentInitiationMetadata != nil {
		toSerialize["include_payment_initiation_metadata"] = o.IncludePaymentInitiationMetadata
	}
	return json.Marshal(toSerialize)
}

type NullableInstitutionsGetRequestOptions struct {
	value *InstitutionsGetRequestOptions
	isSet bool
}

func (v NullableInstitutionsGetRequestOptions) Get() *InstitutionsGetRequestOptions {
	return v.value
}

func (v *NullableInstitutionsGetRequestOptions) Set(val *InstitutionsGetRequestOptions) {
	v.value = val
	v.isSet = true
}

func (v NullableInstitutionsGetRequestOptions) IsSet() bool {
	return v.isSet
}

func (v *NullableInstitutionsGetRequestOptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInstitutionsGetRequestOptions(val *InstitutionsGetRequestOptions) *NullableInstitutionsGetRequestOptions {
	return &NullableInstitutionsGetRequestOptions{value: val, isSet: true}
}

func (v NullableInstitutionsGetRequestOptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInstitutionsGetRequestOptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


