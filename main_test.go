package main

import (
	"fmt"
	"onboarding/lib"
	"onboarding/tcpserver"
	"testing"

	"github.com/NaySoftware/go-fcm"
)

var testmsg = `MSH|^~\&|DPS||PHILIPS||202107101353||ADT^A08|0035648|P|2.3|||AL|NE
EVN|A08|20210710135203
PID||SMW10104511|10702967|200900021|Sonne^Marie||19451010|F|||Untere Hauptstra�e 34^^Walsheim^^76833^D|07337082|21323423||deutsch|||||||||N||D
PV1||I|M-1^M-1-102^102-2^M-IN||200900021|||||||||||||S|200900021||K||||||||||||||||||9201|||||20091208161100||||||200900021
PV2|||||||||||||||||||||0|N
GT1|1|70900021|Sonne^Marie||Untere Hauptstra�e 34^^Walsheim^^76833^D|||19451010000000|||||19010101
`

// TestHelloName calls greetings.Hello with a name, checking
// for a valid return value.
func TestHelloName(t *testing.T) {

	msg, err := tcpserver.ParseHL7(testmsg)
	if err != nil {
		t.Fatalf(`ParseHL7 : want err= "nil", got: %v`, err.Error())
	} else {
		t.Logf("MSG OK: %s", msg.PID)
	}
	/*
		name := "Gladys"
		want := regexp.MustCompile(`\b` + name + `\b`)
		msg, err := Hello("Gladys")
		if !want.MatchString(msg) || err != nil {
			t.Fatalf(`Hello("Gladys") = %q, %v, want match for %#q, nil`, msg, err, want)
		}
	*/
}

func TestCookieID(t *testing.T) {

	// insert into cookieIds (Id) values ("de070071-0b1a-45a5-84d0-cc89d631a960");

	goodID := "de070071-0b1a-45a5-84d0-cc89d631a960"

	//pendingID := "de070071-0b1a-45a5-84d0-pending"
	badID := "FALSE-COOKIE-45a5-84d0-cc89d631a960"

	resGood, err := lib.CheckCookie(goodID)
	if err != nil {
		t.Fatalf(`TestCookieId - Error: %v`, err.Error())
	} else {
		if resGood.Status == "valid" {
			t.Log("TestCookieId - goodId Ok")
		} else {
			t.Fatalf("TestCookieId - goodId failed - didn't got $true")
		}
	}

	resBad, err := lib.CheckCookie(badID)
	if err.Error() != "CookieNotFound" {
		t.Fatalf("TestCookieId - BadID failed -  got : %v ", resBad.Id)

	}

	// TestHelloEmpty calls greetings.Hello with an empty string,
	// checking for an error.
	/*
	   func TestHelloEmpty(t *testing.T) {
	   	msg, err := Hello("")
	   	if msg != "" || err == nil {
	   		t.Fatalf(`Hello("") = %q, %v, want "", error`, msg, err)
	   	}
	   }
	*/
}

const (
	apiKey = "AIzaSyBg6v9qNqSFPFJyOJef5zi8c8jPdfKouy8"
	/*
		authDomain: "testdb-6b732.firebaseapp.com",
		projectId: "testdb-6b732",
		storageBucket: "testdb-6b732.appspot.com",
		messagingSenderId: "147079647684",
		appId: "1:147079647684:web:3c463ba006487b95edd491"
	*/

)

func testFirebird() {

	data := map[string]string{
		"msg": "Hello World1",
		"sum": "Happy Day",
	}

	c := fcm.NewFcmClient(apiKey)
	c.NewFcmMsgTo(topic, data)

	status, err := c.Send()

	if err == nil {
		status.PrintResults()
	} else {
		fmt.Println(err)
	}
}
