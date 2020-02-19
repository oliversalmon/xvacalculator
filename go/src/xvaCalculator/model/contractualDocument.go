package model

import (
	"time"
)

type MasterAgreement struct {
	NettingSet NettingSet
}

type LegalDocument struct {
}

type StandardCSA struct {
}

type DocumentType struct {
	Name      string // i.e. Credit Support Annex
	Publisher string //i.e. ISDA
	Style     string //i.e. 2013 English Law
}

type Party struct{}

type DocumentHeader struct {
	AgreementDate            time.Time
	EffectiveDate            time.Time
	PartyRoles               []*PartyRole
	PartyDocumentIdentifiers []*PartyDocumentIdentifier
}

type Account struct{}

type PartyRole struct{}

type PartyDocumentIdentifier struct {
	Id              string
	PartyReference  string
	DocumentId      []*string
	DocumentVersion string
	AmendedDocument string
}
