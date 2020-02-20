package model

import (
	"time"
)

type MasterAgreement struct {
	nettingSets []*NettingSet
}

type StandardCSA struct {
	DocumentType
	partyDocIdentifiers []*PartyDocumentIdentifier
	settlementDay       string
	baseCurrency        string
	independentAmounts  []*IndependentAmount
}

type SCSA2013NewYorkLaw struct {
	StandardCSA
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

type PartyRole struct{}

type PartyDocumentIdentifier struct {
	id              string
	partyReference  string
	documentIds     []*string
	documentVersion string
	amendedDocument string
}

type IndependentAmount struct{}
