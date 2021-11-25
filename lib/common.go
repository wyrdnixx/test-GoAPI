package lib

import (
	"fmt"
	// "log"
)

func CheckErr(err error) {
	if err != nil {
		//panic(err)
		//log.Fatal(err)
		fmt.Println("an error occourred: %S\n", err.Error())

	}
}
