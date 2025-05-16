package goulstu

import "encoding/json"

func (u *Ulstu) GetTeachers() []string {
	teachersUrl := "https://time.ulstu.ru/api/1.0/teachers"
	resp1, err := u.client.Get(teachersUrl)
	if err != nil {
		panic(err)
	}
	resp1.Body.Close()

	resp2, err := u.client.Get(teachersUrl)
	if err != nil {
		panic(err)
	}
	defer resp2.Body.Close()

	var teachers []string
	if err := json.NewDecoder(resp2.Body).Decode(&teachers); err != nil {
		panic(err)
	}
	return teachers
}
