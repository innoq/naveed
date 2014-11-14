package userindex

import "encoding/json"
import "errors"

type Root struct {
	Member []Member `json:"member"`
}

type Member struct {
	Id string `json:"uid"`
	Name string `json:"displayName"`
	Email string `json:"mail"`
}

type User struct {
	Name string `json:"name"`
	Email string `json:"email"`
}

func Convert(memberData []byte) (userData []byte, err error) {
	root := new(Root)
	err = json.Unmarshal(memberData, &root)
	if err != nil {
		return nil, errors.New("failed to decode JSON data")
	}

	users := map[string]User{}
	for _, member := range root.Member {
		user := new(User)
		user.Name = member.Name
		user.Email = member.Email
		users[member.Id] = *user
	}

	userData, err = json.Marshal(users)
	if err != nil {
		return nil, errors.New("failed to encode JSON data")
	}

	return
}
