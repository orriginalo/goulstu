package goulstu

type TimetableResponse struct {
	Response TimetableData `json:"response"`
	Error    string        `json:"error"`
}

type TimetableData struct {
	Weeks map[string]Week `json:"weeks"`
}

type Week struct {
	Days []Day `json:"days"`
}

type Day struct {
	Day     int        `json:"day"`
	Lessons [][]Lesson `json:"lessons"` // 8 уроков, каждый может содержать 0 или более занятий (подгруппы и т.п.)
}

type Lesson struct {
	Group        string `json:"group"`
	NameOfLesson string `json:"nameOfLesson"`
	Teacher      string `json:"teacher"`
	Room         string `json:"room"`
}
