package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func ParserBy(r *http.Request) Input {
	Value := r.FormValue("value")
	from := r.FormValue("From")
	to := r.FormValue("To")
	if Value == "" || from == "" || to == "" {
		fmt.Errorf("VALUE cannot be empty")
		return Input{}
	}
	num, err := strconv.ParseFloat(Value, 64)
	if err != nil {
		fmt.Errorf("VALUE is not % valid", err)
		return Input{}
	}
	return Input{
		Value: num,
		From:  from,
		To:    to,
	}
}
