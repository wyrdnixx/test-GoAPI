package tcpserver

import (
	"errors"
	"fmt"
	"onboarding/models"
	"regexp"
	"strings"
)

func ParseHL7(message string) error {
	fmt.Println("HL7 parser starting...")
	message = removeFirstVTab(message)
	s := strings.Split(message, "\r")
	fmt.Printf("Full messge : %q\n", s)

	//var hexVerticalTabbyte byte = 0xB
	//for , a := range s {  -> i = index

	//msg := models.HL7message{}
	msg := models.HL7Message{}

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
			//msg.Segments = append(msg.Segments, a)
			elements := strings.Split(a, "|")
			//fmt.Println("Name: ", string(elements[5]))
			msg.MSH.Type = strings.Split(elements[8], "^")[0]
			msg.MSH.Event = strings.Split(elements[8], "^")[1]

			//fmt.Println("MSH: ", msg.MSH)
		}

		if pidfound {

			elements := strings.Split(a, "|")
			msg.PID.PAT = elements[4]
			msg.PID. = strings.Split(elements[5], "^")[0]


			/* Tests
			//fmt.Println("PID found")
			msg.Segments = append(msg.Segments, a)
			elements := strings.Split(a, "|")
			//fmt.Println("Name: ", string(elements[5]))
			msg.PIDPER = elements[4]
			msg.PIDGIVENNAME = strings.Split(elements[5], "^")[0]
			msg.PIDSURNAME = strings.Split(elements[5], "^")[1]
			*/

		}
	}

	//fmt.Println("Header-Fields: ", header.fields[8])

	//lib.InsertFirmen(f)

	if len(msg.MSH.Type) > 0 {
		return nil
		//TEST : return errors.New("error: no MSH Type")
	} else {
		return errors.New("error: no MSH Type")
	}

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
