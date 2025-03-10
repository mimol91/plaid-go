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

// InstitutionStatus The status of an institution is determined by the health of its Item logins, Transactions updates, Investments updates, Liabilities updates, Auth requests, Balance requests, Identity requests, Investments requests, and Liabilities requests. A login attempt is conducted during the initial Item add in Link. If there is not enough traffic to accurately calculate an institution's status, Plaid will return null rather than potentially inaccurate data.  Institution status is accessible in the Dashboard and via the API using the `/institutions/get_by_id` endpoint with the `include_status` option set to true. Note that institution status is not available in the Sandbox environment. 
type InstitutionStatus struct {
	ItemLogins ProductStatus `json:"item_logins"`
	TransactionsUpdates ProductStatus `json:"transactions_updates"`
	Auth ProductStatus `json:"auth"`
	Identity ProductStatus `json:"identity"`
	InvestmentsUpdates ProductStatus `json:"investments_updates"`
	LiabilitiesUpdates *ProductStatus `json:"liabilities_updates,omitempty"`
	Liabilities *ProductStatus `json:"liabilities,omitempty"`
	Investments *ProductStatus `json:"investments,omitempty"`
	// Details of recent health incidents associated with the institution.
	HealthIncidents []HealthIncident `json:"health_incidents,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _InstitutionStatus InstitutionStatus

// NewInstitutionStatus instantiates a new InstitutionStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInstitutionStatus(itemLogins ProductStatus, transactionsUpdates ProductStatus, auth ProductStatus, identity ProductStatus, investmentsUpdates ProductStatus) *InstitutionStatus {
	this := InstitutionStatus{}
	this.ItemLogins = itemLogins
	this.TransactionsUpdates = transactionsUpdates
	this.Auth = auth
	this.Identity = identity
	this.InvestmentsUpdates = investmentsUpdates
	return &this
}

// NewInstitutionStatusWithDefaults instantiates a new InstitutionStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInstitutionStatusWithDefaults() *InstitutionStatus {
	this := InstitutionStatus{}
	return &this
}

// GetItemLogins returns the ItemLogins field value
func (o *InstitutionStatus) GetItemLogins() ProductStatus {
	if o == nil {
		var ret ProductStatus
		return ret
	}

	return o.ItemLogins
}

// GetItemLoginsOk returns a tuple with the ItemLogins field value
// and a boolean to check if the value has been set.
func (o *InstitutionStatus) GetItemLoginsOk() (*ProductStatus, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ItemLogins, true
}

// SetItemLogins sets field value
func (o *InstitutionStatus) SetItemLogins(v ProductStatus) {
	o.ItemLogins = v
}

// GetTransactionsUpdates returns the TransactionsUpdates field value
func (o *InstitutionStatus) GetTransactionsUpdates() ProductStatus {
	if o == nil {
		var ret ProductStatus
		return ret
	}

	return o.TransactionsUpdates
}

// GetTransactionsUpdatesOk returns a tuple with the TransactionsUpdates field value
// and a boolean to check if the value has been set.
func (o *InstitutionStatus) GetTransactionsUpdatesOk() (*ProductStatus, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.TransactionsUpdates, true
}

// SetTransactionsUpdates sets field value
func (o *InstitutionStatus) SetTransactionsUpdates(v ProductStatus) {
	o.TransactionsUpdates = v
}

// GetAuth returns the Auth field value
func (o *InstitutionStatus) GetAuth() ProductStatus {
	if o == nil {
		var ret ProductStatus
		return ret
	}

	return o.Auth
}

// GetAuthOk returns a tuple with the Auth field value
// and a boolean to check if the value has been set.
func (o *InstitutionStatus) GetAuthOk() (*ProductStatus, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Auth, true
}

// SetAuth sets field value
func (o *InstitutionStatus) SetAuth(v ProductStatus) {
	o.Auth = v
}

// GetIdentity returns the Identity field value
func (o *InstitutionStatus) GetIdentity() ProductStatus {
	if o == nil {
		var ret ProductStatus
		return ret
	}

	return o.Identity
}

// GetIdentityOk returns a tuple with the Identity field value
// and a boolean to check if the value has been set.
func (o *InstitutionStatus) GetIdentityOk() (*ProductStatus, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Identity, true
}

// SetIdentity sets field value
func (o *InstitutionStatus) SetIdentity(v ProductStatus) {
	o.Identity = v
}

// GetInvestmentsUpdates returns the InvestmentsUpdates field value
func (o *InstitutionStatus) GetInvestmentsUpdates() ProductStatus {
	if o == nil {
		var ret ProductStatus
		return ret
	}

	return o.InvestmentsUpdates
}

// GetInvestmentsUpdatesOk returns a tuple with the InvestmentsUpdates field value
// and a boolean to check if the value has been set.
func (o *InstitutionStatus) GetInvestmentsUpdatesOk() (*ProductStatus, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.InvestmentsUpdates, true
}

// SetInvestmentsUpdates sets field value
func (o *InstitutionStatus) SetInvestmentsUpdates(v ProductStatus) {
	o.InvestmentsUpdates = v
}

// GetLiabilitiesUpdates returns the LiabilitiesUpdates field value if set, zero value otherwise.
func (o *InstitutionStatus) GetLiabilitiesUpdates() ProductStatus {
	if o == nil || o.LiabilitiesUpdates == nil {
		var ret ProductStatus
		return ret
	}
	return *o.LiabilitiesUpdates
}

// GetLiabilitiesUpdatesOk returns a tuple with the LiabilitiesUpdates field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstitutionStatus) GetLiabilitiesUpdatesOk() (*ProductStatus, bool) {
	if o == nil || o.LiabilitiesUpdates == nil {
		return nil, false
	}
	return o.LiabilitiesUpdates, true
}

// HasLiabilitiesUpdates returns a boolean if a field has been set.
func (o *InstitutionStatus) HasLiabilitiesUpdates() bool {
	if o != nil && o.LiabilitiesUpdates != nil {
		return true
	}

	return false
}

// SetLiabilitiesUpdates gets a reference to the given ProductStatus and assigns it to the LiabilitiesUpdates field.
func (o *InstitutionStatus) SetLiabilitiesUpdates(v ProductStatus) {
	o.LiabilitiesUpdates = &v
}

// GetLiabilities returns the Liabilities field value if set, zero value otherwise.
func (o *InstitutionStatus) GetLiabilities() ProductStatus {
	if o == nil || o.Liabilities == nil {
		var ret ProductStatus
		return ret
	}
	return *o.Liabilities
}

// GetLiabilitiesOk returns a tuple with the Liabilities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstitutionStatus) GetLiabilitiesOk() (*ProductStatus, bool) {
	if o == nil || o.Liabilities == nil {
		return nil, false
	}
	return o.Liabilities, true
}

// HasLiabilities returns a boolean if a field has been set.
func (o *InstitutionStatus) HasLiabilities() bool {
	if o != nil && o.Liabilities != nil {
		return true
	}

	return false
}

// SetLiabilities gets a reference to the given ProductStatus and assigns it to the Liabilities field.
func (o *InstitutionStatus) SetLiabilities(v ProductStatus) {
	o.Liabilities = &v
}

// GetInvestments returns the Investments field value if set, zero value otherwise.
func (o *InstitutionStatus) GetInvestments() ProductStatus {
	if o == nil || o.Investments == nil {
		var ret ProductStatus
		return ret
	}
	return *o.Investments
}

// GetInvestmentsOk returns a tuple with the Investments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstitutionStatus) GetInvestmentsOk() (*ProductStatus, bool) {
	if o == nil || o.Investments == nil {
		return nil, false
	}
	return o.Investments, true
}

// HasInvestments returns a boolean if a field has been set.
func (o *InstitutionStatus) HasInvestments() bool {
	if o != nil && o.Investments != nil {
		return true
	}

	return false
}

// SetInvestments gets a reference to the given ProductStatus and assigns it to the Investments field.
func (o *InstitutionStatus) SetInvestments(v ProductStatus) {
	o.Investments = &v
}

// GetHealthIncidents returns the HealthIncidents field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InstitutionStatus) GetHealthIncidents() []HealthIncident {
	if o == nil  {
		var ret []HealthIncident
		return ret
	}
	return o.HealthIncidents
}

// GetHealthIncidentsOk returns a tuple with the HealthIncidents field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InstitutionStatus) GetHealthIncidentsOk() (*[]HealthIncident, bool) {
	if o == nil || o.HealthIncidents == nil {
		return nil, false
	}
	return &o.HealthIncidents, true
}

// HasHealthIncidents returns a boolean if a field has been set.
func (o *InstitutionStatus) HasHealthIncidents() bool {
	if o != nil && o.HealthIncidents != nil {
		return true
	}

	return false
}

// SetHealthIncidents gets a reference to the given []HealthIncident and assigns it to the HealthIncidents field.
func (o *InstitutionStatus) SetHealthIncidents(v []HealthIncident) {
	o.HealthIncidents = v
}

func (o InstitutionStatus) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["item_logins"] = o.ItemLogins
	}
	if true {
		toSerialize["transactions_updates"] = o.TransactionsUpdates
	}
	if true {
		toSerialize["auth"] = o.Auth
	}
	if true {
		toSerialize["identity"] = o.Identity
	}
	if true {
		toSerialize["investments_updates"] = o.InvestmentsUpdates
	}
	if o.LiabilitiesUpdates != nil {
		toSerialize["liabilities_updates"] = o.LiabilitiesUpdates
	}
	if o.Liabilities != nil {
		toSerialize["liabilities"] = o.Liabilities
	}
	if o.Investments != nil {
		toSerialize["investments"] = o.Investments
	}
	if o.HealthIncidents != nil {
		toSerialize["health_incidents"] = o.HealthIncidents
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *InstitutionStatus) UnmarshalJSON(bytes []byte) (err error) {
	varInstitutionStatus := _InstitutionStatus{}

	if err = json.Unmarshal(bytes, &varInstitutionStatus); err == nil {
		*o = InstitutionStatus(varInstitutionStatus)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "item_logins")
		delete(additionalProperties, "transactions_updates")
		delete(additionalProperties, "auth")
		delete(additionalProperties, "identity")
		delete(additionalProperties, "investments_updates")
		delete(additionalProperties, "liabilities_updates")
		delete(additionalProperties, "liabilities")
		delete(additionalProperties, "investments")
		delete(additionalProperties, "health_incidents")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableInstitutionStatus struct {
	value *InstitutionStatus
	isSet bool
}

func (v NullableInstitutionStatus) Get() *InstitutionStatus {
	return v.value
}

func (v *NullableInstitutionStatus) Set(val *InstitutionStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableInstitutionStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableInstitutionStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInstitutionStatus(val *InstitutionStatus) *NullableInstitutionStatus {
	return &NullableInstitutionStatus{value: val, isSet: true}
}

func (v NullableInstitutionStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInstitutionStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


