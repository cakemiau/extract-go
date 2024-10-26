package extract

import "fmt"

var isVerbose = false

func SetVerbose(value bool) {
	isVerbose = value
}

func verbose(msg any) {
	if isVerbose {
		fmt.Println(msg)
	}
}
