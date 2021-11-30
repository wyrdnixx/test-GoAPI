package models

type HL7messageTest struct {
	Segments     []string
	PIDPER       string
	PIDSURNAME   string
	PIDGIVENNAME string
}

type HL7Message struct {
	MSH MSH
	PID PID
}

type MSH struct {
	Time          string
	Type          string
	Event         string
	ControlNumber string
}

type PID struct {
	PER       string
	PAT       string
	SURNAME   string
	GIVENNAME string
	SEX       string
}

/*
Full messge : ["" "MSH|^~\\&|DPS||PHILIPS||202107101353||ADT^A08|0035648|P|2.3|||AL|NE" "EVN|A08|20210710135203" "PID||SMW10104511|10702967|200900021|Sonne^Marie||19451010|F|||Untere Hauptstra\xdfe 34^^Walsheim^^76833^D|07337082|21323423||deutsch|||||||||N||D" "PV1||I|M-1^M-1-102^102-2^M-IN||200900021|||||||||||||S|200900021||K||||||||||||||||||9201|||||20091208161100||||||200900021" "PV2|||||||||||||||||||||0|N" "GT1|1|70900021|Sonne^Marie||Untere Hauptstra\xdfe 34^^Walsheim^^76833^D|||19451010000000|||||19010101" "\x1c"]
->
->  MSH|^~\&|DPS||PHILIPS||202107101353||ADT^A08|0035648|P|2.3|||AL|NE
->  EVN|A08|20210710135203
->  PID||SMW10104511|10702967|200900021|Sonne^Marie||19451010|F|||Untere Hauptstra�e 34^^Walsheim^^76833^D|07337082|21323423||deutsch|||||||||N||D
->  PV1||I|M-1^M-1-102^102-2^M-IN||200900021|||||||||||||S|200900021||K||||||||||||||||||9201|||||20091208161100||||||200900021
->  PV2|||||||||||||||||||||0|N
->  GT1|1|70900021|Sonne^Marie||Untere Hauptstra�e 34^^Walsheim^^76833^D|||19451010000000|||||19010101
*/
