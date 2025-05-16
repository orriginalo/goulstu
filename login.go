package ulstu

import (
	"fmt"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strings"
)

type Ulstu struct {
	UserLogin string
	Password  string
	client    *http.Client
	jar       *cookiejar.Jar
}

func New(userLogin, password string) *Ulstu {
	jar := &cookiejar.Jar{}
	ulstu := &Ulstu{
		UserLogin: userLogin,
		Password:  password,
		client:    &http.Client{Jar: jar},
		jar:       jar,
	}
	return ulstu
}

func (u *Ulstu) Login() error {
	form := url.Values{}
	form.Add("login", u.UserLogin)
	form.Add("password", u.Password)

	loginReq, _ := http.NewRequest("POST", "https://lk.ulstu.ru/?q=auth/login", strings.NewReader(form.Encode()))
	loginReq.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	loginReq.Header.Set("User-Agent", "Mozilla/5.0")

	_, err := u.client.Do(loginReq)
	if err != nil {
		return err
	}
	return nil
}

func PrintTimetable(timetable TimetableResponse) {
	if timetable.Error != "" {
		fmt.Println("Ошибка:", timetable.Error)
		return
	}

	for weekName, week := range timetable.Response.Weeks {
		fmt.Printf("Неделя: %s\n", weekName)

		for _, day := range week.Days {
			fmt.Printf("\tДень: %d\n", day.Day)

			for lessonIndex, lessonGroup := range day.Lessons {
				fmt.Printf("\t\tУрок #%d\n", lessonIndex+1)

				if len(lessonGroup) == 0 {
					fmt.Printf("\t\t\t<нет занятий>\n")
					continue
				}

				for _, lesson := range lessonGroup {
					fmt.Printf("\t\t\tГруппа: %s\n", lesson.Group)
					fmt.Printf("\t\t\tПредмет: %s\n", lesson.NameOfLesson)
					fmt.Printf("\t\t\tПреподаватель: %s\n", lesson.Teacher)
					fmt.Printf("\t\t\tАудитория: %s\n", lesson.Room)
					fmt.Println()
				}
			}
		}
	}
}
