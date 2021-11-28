package tcpserver

import (
	"fmt"
	"regexp"
	"strings"
)

type msh struct {
	fields []string
}

func parseHL7(message string) {
	fmt.Println("HL7 parser starting...")
	s := strings.Split(message, "\r")
	fmt.Printf("Full messge : %q\n", s)

	//var hexVerticalTabbyte byte = 0xB
	//for , a := range s {  -> i = index
	var header msh
	for _, a := range s {
		MSHmatched, _ := regexp.MatchString("MSH|", a)
		if MSHmatched {
			fmt.Println("MSHHeader found")

			s1 := strings.Split(a, "|")
			for _, x := range s1 {
				header.fields = append(header.fields, x)
			}
		}
	}

	fmt.Println("Header-Fields: ", header.fields)

}
