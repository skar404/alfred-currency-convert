package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/deanishe/awgo"

	"alfred_course/course"
)

func getQuery(wf aw.Workflow) string {
	var query string

	// Use wf.Args() to enable Magic Actions
	if args := wf.Args(); len(args) > 0 {
		query = args[0]

		log.Printf("[main] query=%s", query)
	}

	return query
}

func main() {
	wf := aw.New()

	query := getQuery(*wf)

	var userMoney float32
	raw := course.GetCBRCourse()
	if len(query) != 0 {
		i, err := strconv.Atoi(query)
		if err == nil {
			userMoney = float32(i)
		}
	}

	for _, item := range raw.Data {
		courseInf := item.Value
		if userMoney != 0 {
			courseInf = userMoney / item.Value
		}

		wf.NewItem(fmt.Sprintf("%.2f %s", courseInf, item.CharCode)).
			Icon(&aw.Icon{Value: fmt.Sprintf("flags/%s.png", item.CharCode)})
	}

	wf.SendFeedback()
}
