package tcpserver

import (
	"fmt"
	"strings"
)

type msh struct {
	fields []string
}

func parseHL7(message string) {
	fmt.Println("HL7 parser starting...")
	message = removeFirstVTab(message)
	s := strings.Split(message, "\r")
	fmt.Printf("Full messge : %q\n", s)

	//var hexVerticalTabbyte byte = 0xB
	//for , a := range s {  -> i = index
	var fields msh

	for _, a := range s {
		//mshfound, err := regexp.MatchString("MSH|", a)
		//s1 := strings.Split(a, "|")
		//for _, x := range s1 {
		//	header.fields = append(header.fields, x)
		//}

		fields.fields = append(fields.fields, a)
		fmt.Println("-> ", a)

	}

	//fmt.Println("Header-Fields: ", header.fields[8])

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
