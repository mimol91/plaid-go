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
	"fmt"
)

// InvestmentAccountSubtype Valid account subtypes for investment accounts. For a list containing descriptions of each subtype, see [Account schemas](https://plaid.com/docs/api/accounts/#StandaloneAccountType-investment).
type InvestmentAccountSubtype string

// List of InvestmentAccountSubtype
const (
	INVESTMENTACCOUNTSUBTYPE__529 InvestmentAccountSubtype = "529"
	INVESTMENTACCOUNTSUBTYPE__401A InvestmentAccountSubtype = "401a"
	INVESTMENTACCOUNTSUBTYPE__401K InvestmentAccountSubtype = "401k"
	INVESTMENTACCOUNTSUBTYPE__403_B InvestmentAccountSubtype = "403B"
	INVESTMENTACCOUNTSUBTYPE__457B InvestmentAccountSubtype = "457b"
	INVESTMENTACCOUNTSUBTYPE_BROKERAGE InvestmentAccountSubtype = "brokerage"
	INVESTMENTACCOUNTSUBTYPE_CASH_ISA InvestmentAccountSubtype = "cash isa"
	INVESTMENTACCOUNTSUBTYPE_EDUCATION_SAVINGS_ACCOUNT InvestmentAccountSubtype = "education savings account"
	INVESTMENTACCOUNTSUBTYPE_FIXED_ANNUITY InvestmentAccountSubtype = "fixed annuity"
	INVESTMENTACCOUNTSUBTYPE_GIC InvestmentAccountSubtype = "gic"
	INVESTMENTACCOUNTSUBTYPE_HEALTH_REIMBURSEMENT_ARRANGEMENT InvestmentAccountSubtype = "health reimbursement arrangement"
	INVESTMENTACCOUNTSUBTYPE_HSA InvestmentAccountSubtype = "hsa"
	INVESTMENTACCOUNTSUBTYPE_IRA InvestmentAccountSubtype = "ira"
	INVESTMENTACCOUNTSUBTYPE_ISA InvestmentAccountSubtype = "isa"
	INVESTMENTACCOUNTSUBTYPE_KEOGH InvestmentAccountSubtype = "keogh"
	INVESTMENTACCOUNTSUBTYPE_LIF InvestmentAccountSubtype = "lif"
	INVESTMENTACCOUNTSUBTYPE_LIFE_INSURANCE InvestmentAccountSubtype = "life insurance"
	INVESTMENTACCOUNTSUBTYPE_LIRA InvestmentAccountSubtype = "lira"
	INVESTMENTACCOUNTSUBTYPE_LRIF InvestmentAccountSubtype = "lrif"
	INVESTMENTACCOUNTSUBTYPE_LRSP InvestmentAccountSubtype = "lrsp"
	INVESTMENTACCOUNTSUBTYPE_MUTUAL_FUND InvestmentAccountSubtype = "mutual fund"
	INVESTMENTACCOUNTSUBTYPE_NON_TAXABLE_BROKERAGE_ACCOUNT InvestmentAccountSubtype = "non-taxable brokerage account"
	INVESTMENTACCOUNTSUBTYPE_OTHER InvestmentAccountSubtype = "other"
	INVESTMENTACCOUNTSUBTYPE_OTHER_ANNUITY InvestmentAccountSubtype = "other annuity"
	INVESTMENTACCOUNTSUBTYPE_OTHER_INSURANCE InvestmentAccountSubtype = "other insurance"
	INVESTMENTACCOUNTSUBTYPE_PERSON InvestmentAccountSubtype = "person"
	INVESTMENTACCOUNTSUBTYPE_PRIF InvestmentAccountSubtype = "prif"
	INVESTMENTACCOUNTSUBTYPE_PROFIT_SHARING_PLAN InvestmentAccountSubtype = "profit sharing plan"
	INVESTMENTACCOUNTSUBTYPE_QSHR InvestmentAccountSubtype = "qshr"
	INVESTMENTACCOUNTSUBTYPE_RDSP InvestmentAccountSubtype = "rdsp"
	INVESTMENTACCOUNTSUBTYPE_RESP InvestmentAccountSubtype = "resp"
	INVESTMENTACCOUNTSUBTYPE_RETIREMENT InvestmentAccountSubtype = "retirement"
	INVESTMENTACCOUNTSUBTYPE_RLIF InvestmentAccountSubtype = "rlif"
	INVESTMENTACCOUNTSUBTYPE_ROTH InvestmentAccountSubtype = "roth"
	INVESTMENTACCOUNTSUBTYPE_ROTH_401K InvestmentAccountSubtype = "roth 401k"
	INVESTMENTACCOUNTSUBTYPE_RRIF InvestmentAccountSubtype = "rrif"
	INVESTMENTACCOUNTSUBTYPE_RRSP InvestmentAccountSubtype = "rrsp"
	INVESTMENTACCOUNTSUBTYPE_SARSEP InvestmentAccountSubtype = "sarsep"
	INVESTMENTACCOUNTSUBTYPE_SEP_IRA InvestmentAccountSubtype = "sep ira"
	INVESTMENTACCOUNTSUBTYPE_SIMPLE_IRA InvestmentAccountSubtype = "simple ira"
	INVESTMENTACCOUNTSUBTYPE_SIPP InvestmentAccountSubtype = "sipp"
	INVESTMENTACCOUNTSUBTYPE_STOCK_PLAN InvestmentAccountSubtype = "stock plan"
	INVESTMENTACCOUNTSUBTYPE_TFSA InvestmentAccountSubtype = "tfsa"
	INVESTMENTACCOUNTSUBTYPE_TRUST InvestmentAccountSubtype = "trust"
	INVESTMENTACCOUNTSUBTYPE_UGMA InvestmentAccountSubtype = "ugma"
	INVESTMENTACCOUNTSUBTYPE_UTMA InvestmentAccountSubtype = "utma"
	INVESTMENTACCOUNTSUBTYPE_VARIABLE_ANNUITY InvestmentAccountSubtype = "variable annuity"
	INVESTMENTACCOUNTSUBTYPE_ALL InvestmentAccountSubtype = "all"
)

var allowedInvestmentAccountSubtypeEnumValues = []InvestmentAccountSubtype{
	"529",
	"401a",
	"401k",
	"403B",
	"457b",
	"brokerage",
	"cash isa",
	"education savings account",
	"fixed annuity",
	"gic",
	"health reimbursement arrangement",
	"hsa",
	"ira",
	"isa",
	"keogh",
	"lif",
	"life insurance",
	"lira",
	"lrif",
	"lrsp",
	"mutual fund",
	"non-taxable brokerage account",
	"other",
	"other annuity",
	"other insurance",
	"person",
	"prif",
	"profit sharing plan",
	"qshr",
	"rdsp",
	"resp",
	"retirement",
	"rlif",
	"roth",
	"roth 401k",
	"rrif",
	"rrsp",
	"sarsep",
	"sep ira",
	"simple ira",
	"sipp",
	"stock plan",
	"tfsa",
	"trust",
	"ugma",
	"utma",
	"variable annuity",
	"all",
}

func (v *InvestmentAccountSubtype) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := InvestmentAccountSubtype(value)
	for _, existing := range allowedInvestmentAccountSubtypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid InvestmentAccountSubtype", value)
}

// NewInvestmentAccountSubtypeFromValue returns a pointer to a valid InvestmentAccountSubtype
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewInvestmentAccountSubtypeFromValue(v string) (*InvestmentAccountSubtype, error) {
	ev := InvestmentAccountSubtype(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for InvestmentAccountSubtype: valid values are %v", v, allowedInvestmentAccountSubtypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v InvestmentAccountSubtype) IsValid() bool {
	for _, existing := range allowedInvestmentAccountSubtypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to InvestmentAccountSubtype value
func (v InvestmentAccountSubtype) Ptr() *InvestmentAccountSubtype {
	return &v
}

type NullableInvestmentAccountSubtype struct {
	value *InvestmentAccountSubtype
	isSet bool
}

func (v NullableInvestmentAccountSubtype) Get() *InvestmentAccountSubtype {
	return v.value
}

func (v *NullableInvestmentAccountSubtype) Set(val *InvestmentAccountSubtype) {
	v.value = val
	v.isSet = true
}

func (v NullableInvestmentAccountSubtype) IsSet() bool {
	return v.isSet
}

func (v *NullableInvestmentAccountSubtype) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInvestmentAccountSubtype(val *InvestmentAccountSubtype) *NullableInvestmentAccountSubtype {
	return &NullableInvestmentAccountSubtype{value: val, isSet: true}
}

func (v NullableInvestmentAccountSubtype) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInvestmentAccountSubtype) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

