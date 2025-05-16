package goulstu

import "encoding/json"

func (u *Ulstu) GetGroups() GroupsResponse {
	groupsUrl := "https://time.ulstu.ru/api/1.0/groups"
	resp1, err := u.client.Get(groupsUrl)
	if err != nil {
		panic(err)
	}
	resp1.Body.Close()

	resp2, err := u.client.Get(groupsUrl)
	if err != nil {
		panic(err)
	}
	defer resp2.Body.Close()

	var groups GroupsResponse
	if err := json.NewDecoder(resp2.Body).Decode(&groups); err != nil {
		panic(err)
	}
	return groups
}
