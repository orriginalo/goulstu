package goulstu

import (
	"encoding/json"
	"fmt"
)

func (u *Ulstu) GetTimetable(group string) TimetableResponse {
	timetableUrl := fmt.Sprintf("https://time.ulstu.ru/api/1.0/timetable?filter=%s", group)
	resp1, err := u.client.Get(timetableUrl)
	if err != nil {
		panic(err)
	}
	resp1.Body.Close()

	resp2, err := u.client.Get(timetableUrl)
	if err != nil {
		panic(err)
	}
	defer resp2.Body.Close()

	var timetable TimetableResponse
	if err := json.NewDecoder(resp2.Body).Decode(&timetable); err != nil {
		panic(err)
	}
	return timetable
}

func (u *Ulstu) GetTimetables(groups []string) []TimetableResponse {
	var timetables []TimetableResponse
	for _, group := range groups {
		timetable := u.GetTimetable(group)
		timetables = append(timetables, timetable)
	}
	return timetables
}
