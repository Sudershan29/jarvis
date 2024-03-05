package main

import (
	"backend/src/algorithm/planner"
	"backend/src/models"
	"fmt"
)

const MY_UUID = "f3d43975-84a2-4af3-b0b4-3c8f3fa13af0"

func main() {
	me, _ := models.UserFind(MY_UUID)

	fmt.Println(me)
	planner.PrepareTimeTable(MY_UUID)
}
