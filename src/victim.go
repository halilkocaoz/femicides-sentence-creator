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
	return false
}

func (v Victim) CreateSentence() string {
	if v.IsSuitToCreateSentence() {

	}
	return "Test : " + v.FullName
}

func (v Victim) SelectByWhomStatement() string {
	return ""
}

func (v Victim) SelectMethodStatement() string {
	return ""
}

func (v Victim) SelectCauseStatement() string {
	return ""
}
