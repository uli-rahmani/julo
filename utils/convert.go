package utils

import (
	"fmt"
	"julo/constants"
	"strconv"
	"time"
)

func GetInt(x string) int {
	i, err := strconv.Atoi(x)
	if err != nil {
		fmt.Println("utils -> GentInt : error, ", err)
		fmt.Println("Can't convert into Integer")
		fmt.Println("Please re-check .env, you recently input")
		fmt.Println(x)
	}

	return i
}

func ToFormatTime(datetime string) (string, error) {
	t, err := time.Parse(constants.DBTimeLayout, datetime)
	if err != nil {
		return datetime, err
	}

	tString := t.Format(constants.ResponseTimeLayout)

	return tString, nil
}

func GetTimeString() string {
	t := time.Now()
	tString := t.Format(constants.DBTimeLayout)

	return tString
}
