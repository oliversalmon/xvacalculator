package model

import (
	_ "time"
)

type LegalDocument interface {
}

type StandardCSA struct {
}

type StandardCSA2013EnglishLaw struct{}

type StandardCSA2013NewYorkLaw struct{}

type StandardCSA2014EnglishLaw struct{}

type StandardCSA2014NewYorkLaw struct{}

type Party struct{}

type DocumentHeader struct {
	AgreementDate            string
	EffectiveDate            string
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
}
