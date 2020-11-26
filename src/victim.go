package main

import (
	"errors"
	"strings"
)

//femicidesinturkey.com/api/victim
type FemicidesInTurkeyJson struct {
	Message     interface{} `json:"message"`
	Information struct {
		Count int         `json:"count"`
		Pages int         `json:"pages"`
		Next  string      `json:"next"`
		Prev  interface{} `json:"prev"`
	} `json:"information"`
	Data []Victim
}

type Victim struct {
	FullName string `json:"fullName"`
	City     string `json:"city"`
	Killer   struct {
		Definition string `json:"definition"`
		Status     string `json:"status"`
	} `json:"killer"`
	Methods []struct {
		Method string `json:"method"`
	} `json:"methods"`
	Causes []struct {
		Cause string `json:"cause"`
	} `json:"causes"`
	Adult             bool        `json:"adult"`
	ProtectionRequest interface{} `json:"protectionRequest"`
	Year              string      `json:"year"`
	URL               string      `json:"url"`
}

func (v Victim) IsFullNameSuit() bool {
	fnLen := len(v.FullName)
	fnLowered := strings.ToLower(v.FullName)

	if fnLen > 6 && !(strings.Contains(fnLowered, "ismi") || strings.Contains(fnLowered, "bilinmiyor") || strings.Contains(fnLowered, "undefined") || strings.Contains(fnLowered, "unknown")) {
		return true
	}
	return false
}

func (v Victim) IsCauseSuit() bool {
	if v.Causes[0].Cause != "Unknown" {
		return true
	}
	return false
}

func (v Victim) IsSuitToCreateSentence() bool {
	if v.IsFullNameSuit() && v.IsCauseSuit() {
		return true
	}
	return false
}

func (v Victim) CreateSentence() (string, error) {
	if v.IsSuitToCreateSentence() {
		sentence := v.FullName + " was murdered " + v.ByWhomPart() + v.CausePart() + v.YearCityPart()
		return sentence + ".", nil
	}
	return "", errors.New("Regular sentence can't create")
}

func (v Victim) ByWhomPart() string {
	ByWhom := "by "
	justBy := v.Killer.Definition == "Someone she knowns"
	byPeople := v.Killer.Definition == "Unknown" || v.Killer.Definition == "Foreigner"

	if justBy {
		ByWhom += v.Killer.Definition
	} else if byPeople { // by unknown people, by foreigner people
		ByWhom += v.Killer.Definition + " people"
	} else { // by her husband, boyfriend, kinsman and others..
		ByWhom += "her " + v.Killer.Definition
	}

	return strings.ToLower(ByWhom + " ")
}

func (v Victim) CausePart() string {
	Cause := ""

	causeAfterThe := v.Causes[0].Cause == "Sexual assault" || v.Causes[0].Cause == "Break up"
	causeBecause := strings.Contains(v.Causes[0].Cause, "Because")
	causeWhile := strings.Contains(v.Causes[0].Cause, "Protecting")

	if causeWhile { // murdered while doing
		Cause += "while " + v.Causes[0].Cause
	} else if causeBecause { // murdered because x
		Cause += v.Causes[0].Cause
	} else if causeAfterThe { // murdered after the break up
		Cause += "after the " + v.Causes[0].Cause
	} else { // murdered because of the money, envy, honor
		Cause += "because of the " + v.Causes[0].Cause
	}

	return strings.ToLower(Cause)
}

func (v Victim) YearCityPart() string {
	YearCity := ", in " + v.Year

	if v.City != "Unknown" {
		YearCity += ", " + v.City
	}
	return YearCity
}
