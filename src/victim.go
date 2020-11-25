package main

//femicidesinturkey.com/api/statistic/killer
//femicidesinturkey.com/api/statistic/cause
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

func (v Victim) IsSuitToCreateSentence() bool {
	return true
	//todo: fullname length, skip suicide cases.
}

func (v Victim) CreateSentence() string {
	if v.IsSuitToCreateSentence() {
		sentence := v.FullName + " was murdered "

		sentence += v.ByWhomPart()
		sentence += "in " + v.City + ", " + v.Year
		return sentence + "."
	}

	return ""
}

/*
	killer definitions: femicidesinturkey.com/api/statistic/killer
*/
func (v Victim) ByWhomPart() string {
	ByWhom := "by "

	if v.Killer.Definition == "Someone she knowns" { //by someone she knows
		ByWhom += v.Killer.Definition
	} else if v.Killer.Definition == "Unknown" || v.Killer.Definition == "Foreigner" { // by unknown people, by foreigner people
		ByWhom += v.Killer.Definition + " people"
	} else { // by her husband, boyfriend, kinsman and others..
		ByWhom += "her " + v.Killer.Definition
	}

	return ByWhom + " "
}

func (v Victim) SelectCauseStatement() string {
	return ""
}

func (v Victim) SelectMethodStatement() string {
	return ""
}
