package helpers

import "os/user"

type UserInfo struct {
	Home     string
	Username string
}

func GetCurrentUserInfo() (UserInfo, error) {
	user, err := user.Current()
	if err != nil {
		return UserInfo{}, err
	}

	return UserInfo{
		Home:     user.HomeDir,
		Username: user.Username,
	}, nil
}
