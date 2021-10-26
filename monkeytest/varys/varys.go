package varys

type UserInfo struct {
	Name     string
	Birthday string
}

func GetInfoByUID(uid int64) (*UserInfo, error) {
	return &UserInfo{}, nil
}
