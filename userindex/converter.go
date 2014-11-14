package userindex

import "encoding/json"
import "errors"

type User struct {
	Name string `json:"name"`
	Email string `json:"email"`
}

type registry struct {
	Member []member `json:"member"`
}

type member struct {
	Id string `json:"uid"`
	Name string `json:"displayName"`
	Email string `json:"mail"`
}

func Convert(memberData []byte) (userData []byte, err error) {
	reg := new(registry)
	err = json.Unmarshal(memberData, &reg)
	if err != nil {
		return nil, errors.New("failed to decode JSON data")
	}

	users := map[string]User{}
	for _, member := range reg.Member {
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
