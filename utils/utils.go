package utils

import (
	"gopls-workspace/models"
	"log"
	"os"
	"strconv"
	"time"
)

func CheckNewUserTimer(u *models.GetUser) int {
	duration := time.Duration(60) * time.Second
	tk:= time.NewTicker(duration)
	i := 0
	for range tk.C {
		if i > 9{
		}
	}

	return 0
}
func MustGetenv(k string) string {
	v := os.Getenv(k)
	if v == "" {
		log.Fatalf("%s environment variable not set.", k)
	}
	return v
}
func AddToSingleDigit(t string) string {
	if len(t) == 1 {
		t = "0" + t
	}
	return t
}

func IntToString(f int) string {
	// to convert a float number to a string
	return strconv.Itoa(f)
}


func StringToInt(f string) int {
	n, err := strconv.Atoi(f)
	if err != nil {
		panic(err)
	}
	return n
}

