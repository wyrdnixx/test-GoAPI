package tcpserver

import (
	"fmt"
	"onboarding/lib"
	"onboarding/models"
	"regexp"
	"strings"
)

func parseHL7(message string) {
	fmt.Println("HL7 parser starting...")
	message = removeFirstVTab(message)
	s := strings.Split(message, "\r")
	fmt.Printf("Full messge : %q\n", s)

	//var hexVerticalTabbyte byte = 0xB
	//for , a := range s {  -> i = index

	msg := models.HL7message{}

	for _, a := range s {
		//mshfound, err := regexp.MatchString("MSH|", a)
		//s1 := strings.Split(a, "|")
		//for _, x := range s1 {
		//	header.fields = append(header.fields, x)
		//}
		fmt.Println("-> ", a)

		mshfound, err := regexp.MatchString(`^MSH\|`, a)
		if err != nil {
			fmt.Println("Not an MSH segment")
		}
		pidfound, err := regexp.MatchString(`^PID\|`, a)
		if err != nil {
			fmt.Println("Not an PID segment")
		}

		if mshfound {
			//fmt.Println("MSH found")
			msg.Segments = append(msg.Segments, a)

		}

		if pidfound {
			//fmt.Println("PID found")
			msg.Segments = append(msg.Segments, a)
			elements := strings.Split(a, "|")
			//fmt.Println("Name: ", string(elements[5]))
			msg.PIDPER = elements[4]
			msg.PIDGIVENNAME = strings.Split(elements[5], "^")[0]
			msg.PIDSURNAME = strings.Split(elements[5], "^")[1]

		}
	}

	//fmt.Println("Header-Fields: ", header.fields[8])

	fmt.Println("Message-Object after parser: ", msg)
	fmt.Println("Message-Object given: ", msg.PIDGIVENNAME)
	fmt.Println("Message-Object sure: ", msg.PIDSURNAME)
	fmt.Println("Message-Object per: ", msg.PIDPER)

	// test - add as company
	f := models.NewCompany{}
	f.Name = msg.PIDGIVENNAME
	fmt.Println("newfirma name ", f.Name)
	lib.InsertFirmen(f)

}

func removeFirstVTab(s string) string {
	var hexVerticalTabbyte byte = 0xB

	b := []byte(s)
	var res []byte

	for i, v := range b {
		if v == hexVerticalTabbyte {
			fmt.Println("removing vTAB on pos:", i)
		} else {
			res = append(res, v)
		}
	}
	return string(res)

}
