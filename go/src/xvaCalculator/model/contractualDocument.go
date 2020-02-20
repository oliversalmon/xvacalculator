package model

import (
	"time"
)

type MasterAgreement struct {
	nettingSets []*NettingSet
	DocumentIdentity
}

type DocumentIdentity struct {
	id string
	DocumentType
	agreementDate            time.Time
	partyDocumentIdentifiers []*PartyDocumentIdentifier
}

type StandardCSA struct {
	DocumentType
	partyDocIdentifiers []*PartyDocumentIdentifier
	settlementDay       string // i.e. T+1 or T+2
	baseCurrency        string // ISO code for currency
	independentAmounts  [2]*IndependentAmount
}

type SCSA2013NewYorkLaw struct {
	StandardCSA
	DisputeResolution
	demandsAndNotices []*DemandAndNotice
	OtherProvisions
	transportCurrency [2]*TransportCurrency
	DayCount
	IndependentAmountInterestRate
	independentAmountEligibleCollateral [2]*IndependentAmountEligibleCollateral
	holdingAndUsingPostedCollateral     [2]*HoldingAndUsingPostedCollateral
}

type SCSA2014NewYorkLaw struct {
	StandardCSA
	demandsAndNotices                      []*DemandAndNotice
	independentAmountEligibleCreditSupport [2]*IndependentAmountEligibleCollateral
	useOfPostedCreditSupport               [2]*useOfPostedCreditSupport
	valuationDateCity                      []*string // i.e. {USNY,GBLO}
	notificationTimeCity                   string    // i.e. USNY

}

type SCSA2013EnglishLaw struct {
	StandardCSA
	DisputeResolution
	demandsAndNotices []*DemandAndNotice
	OtherProvisions
	transportCurrency [2]*TransportCurrency
	DayCount
	IndependentAmountInterestRate
	independentAmountEligibleCreditSupport [2]*IndependentAmountEligibleCollateral
	exchangeDate                           string // i.e. T, T+1, T+2 etc
}

type SCSA2014EnglishLaw struct {
	StandardCSA
	demandsAndNotices                      []*DemandAndNotice
	independentAmountEligibleCreditSupport [2]*IndependentAmountEligibleCollateral
	valuationDateCity                      []*string // i.e. {USNY,GBLO}
	notificationTimeCity                   string    // i.e. USNY
}

type DocumentType struct {
	name      string // i.e. Credit Support Annex
	publisher string //i.e. ISDA
	style     string //i.e. 2013 English Law
}

type Party struct{}

type DocumentHeader struct {
	agreementDate            time.Time
	effectiveDate            time.Time
	partyRoles               []*PartyRole
	partyDocumentIdentifiers []*PartyDocumentIdentifier
}

type Account struct{}

type PartyRole struct {
	relatedParties []*RelatedParty
}

type RelatedParty struct {
	partyReference   string
	accountReference string
	role             string
}

type PartyDocumentIdentifier struct {
	id              string
	partyReference  string
	documentIds     []*string
	documentVersion string
	amendedDocument string
}

type IndependentAmount struct{}

type DisputeResolution struct {
	resolutionTime string // Hour minute time i.e. 14:00:00
	businessCentre string // i.e. USNY
}

type DemandAndNotice struct {
	partyReference string
	ContactInfo
	businessUnit []*BusinessUnit
	person       []*Person
}

type Person struct {
}

type BusinessUnit struct {
}

type ContactInfo struct {
	telephone []*string
	email     []*string
	Address
}

type Address struct {
	streetAddress1 string
	streetAddress2 string
	streetAddress3 string
	city           string
	state          string
	country        string // ISO Code for country
	postalCode     string
}

type TransportCurrency struct {
	partyReference           string
	electedTransportCurrency string // ISO Currency code
}

type CurrencySpecificDayCount struct {
	dayCountValue string // i.e. 360 or 365
	currency      string // ISO currency code
}

type DayCount struct {
	defaultDayCount           string // i.e. 360 or 365
	currencySpecificDayCounts []*CurrencySpecificDayCount
}

type IndependentAmountInterestRate struct {
	InitialMarginInterestRateTerms
	specifiedRate []*SpecifiedRate
}

type InitialMarginInterestRateTerms struct {
	term string
}

type SpecifiedRate struct {
	currency           string
	fixedRate          float32
	floatingRate       string // name of floating rate i.e. EUR-LIBOR-BBA
	floatingRateSpread float32
}

type IndependentAmountEligibleCollateral struct {
	partyReference               string
	eligibleCollateral           []*EligibleCollateral
	independentAmountEligibility string
}

type EligibleCollateral struct {
	eligibleAsset       string
	lowerMaturity       eligibleCollateralMaturity
	higherMaturity      eligibleCollateralMaturity
	minimumRating       string
	valuationPercentage float32
}

type eligibleCollateralMaturity struct {
	periodMultiplier int
	period           string
}

type HoldingAndUsingPostedCollateral struct {
	partyReference string
	eligibilityToHoldCollateral
	useOfPostedCollateral bool
}

type eligibilityToHoldCollateral struct {
	holdingPostedCollateral string
	custodianTerms
}

type custodianTerms struct {
	minimumAssets string
	minimumRating
	initialDesignation string
}

type minimumRating struct {
	condition      string
	creditNotation []*creditNotation
}

type creditNotation struct {
	agency   string
	notation string
	debt
}

type debt struct {
	condition string
	debtType  []*string
}

type OtherProvisions struct {
	provisions map[string]string
}

type useOfPostedCreditSupport struct {
	partyReference        string
	useOfPostedCollateral bool
}
