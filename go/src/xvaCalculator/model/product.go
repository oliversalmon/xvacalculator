package model

import (
	"time"
)

// Strategy
type Strategy struct {
	strategyComponentIdentifiers []*strategyComponentIdentifier
	premiumProductReference
	product []*Product
}

type strategyComponentIdentifier struct{}

type premiumProductReference struct{}

type Product struct{}

// FX Products
type FXSingleLeg struct {
	id string
	ProductModel
	FXCoreDetailsModel
}

// Models
type ProductModel struct{}

// FX Core Details Model
type FXCoreDetailsModel struct {
	exchangeTradedCurrency1 exchangeTradedCurrency
	exchangeTradedCurrency2 exchangeTradedCurrency
	dealtCurrency           string
	fxTenorModel
	valueDate          time.Time // Use either valueDate or the ones for each currency
	currency1ValueDate time.Time
	currency2ValueDate time.Time
	exchangeRate
}

type exchangeTradedCurrency struct {
	id string
	PayerReceiverModel
	paymentAmount
	paymentDate time.Time
	paymentType
	settlementInformation
	discountFactor float32
	presentValueAmount
}

type PayerReceiverModel struct {
	payer    partyReference
	receiver partyReference
}

type partyReference struct {
	partyReference   string
	accountReference string
}

type paymentAmount struct {
}

type paymentType struct {
}

type settlementInformation struct {
	standardSettlementStyle string
	settlementInstruction
}

type settlementInstruction struct {
	settlementMethod string
	correspondentInformation
	intermediaryInformation []*intermediaryInformation
	beneficiaryBank         beneficiary
	beneficiary
	depositoryPartyReference string
	splitSettlement          []*splitSettlement
}

type correspondentInformation struct {
	routingIdentificationModel
	correspondentPartyReference string
}

type routingIdentificationModel struct {
	routingIds []*string
	routingExplicitDetails
}

type routingExplicitDetails struct {
	routingName          string
	routingAddress       Address
	routingAccountNumber string
	routingReferenceText []*string
}

type intermediaryInformation struct {
	routingIdentificationModel
	intermediarySequenceNumber int32
	intermediaryPartyReference string
}

type beneficiary struct {
	routingIdentificationModel
	beneficiaryPartyReference string
}

type splitSettlement struct {
	splitSettlementAmount float64
	beneficiaryBank       beneficiary
	beneficiary
}

type presentValueAmount struct {
	currency string // ISO currency code
	amount   float64
}

type fxTenorModel struct {
	tenorName string // Enumeration of "Broken", "Today", "Tomorrow", "TomorrowNext", "Spot", "SpotNext"
	tenorPeriod
}

type tenorPeriod struct {
	id               string
	periodMultiplier int32
	period           string // Enumeration of "D", "W", "M", "Y"
}

type exchangeRate struct {
	quotedCurrencyPair
	rate          float32
	spotRate      float32
	forwardPoints float32
	pointValue    float32
	crossRate     []*crossRate
}

type quotedCurrencyPair struct {
	currency1  string
	currency2  string
	quoteBasis string // Enumeration:"Currency1PerCurrency2"-The amount of currency1 for one unit of currency2 	"Currency2PerCurrency1"-The amount of currency2 for one unit of currency1
}

type crossRate struct {
	currency1     string
	currency2     string
	rate          float32
	quoteBasis    string // Enumeration:"Currency1PerCurrency2"-The amount of currency1 for one unit of currency2 	"Currency2PerCurrency1"-The amount of currency2 for one unit of currency1
	spotRate      float32
	forwardPoints float32
}
type nonDeliverableSettlement struct {
	settlementCurrency string
	referenceCurrency  string
	notionalAmount     positiveMoney
	fixing             []*fxFixing
}

type positiveMoney struct {
	id       string
	currency string
	amount   float32
}
type fxFixing struct {
	quotedCurrencyPair
	fxFixingDate time.Time
	fxSportRateSource
}

type fxSportRateSource struct {
	rateSource            string
	rateSourcePage        string
	rateSourcePageHeading string
}

type fxRateSourceFixing struct {
}

type fxSettlementRateSource struct {
	settlementRateOption      string
	nonStandardSettlementRate fxInformationSource
}

type fxInformationSource struct {
	rateSource        string
	rateSourcePage    string
	rateSourceHeading string
	fixingTime        businessCentreTime
}

type businessCentreTime struct {
	hourMinuteTime time.Time
	businessCentre string
}
