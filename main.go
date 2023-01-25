package main

import (
	"github.com/sirupsen/logrus"
	"github.com/tanya-lyubimaya/date_difference_calculator/cmd"
)

func main() {
	err := cmd.Execute()
	if err != nil {
		logrus.Fatalln(err)
	}
}
