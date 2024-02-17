package main

import (
	"backend/src/algorithm/planner"
	"backend/src/models"
	"fmt"
)

const MY_UUID = "bafdaa75-0360-4b81-b4ed-2f8e098968d8"

func main() {
	me, _ := models.UserFind(MY_UUID)

	fmt.Print(me)

	planner.PrepareTimeTable(MY_UUID)
}
