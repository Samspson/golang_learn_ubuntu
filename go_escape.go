package main

type UserData struct {
	Name string
}

/*
	escape to heap
func main() {
	var info UserData
	info.Name = "WiburXU"
	_ = GetUserInfo(info)
}

func GetUserInfo(userInfo UserData) *UserData {
	return &userInfo
}
*/

func main() {
	var info UserData
	info.Name = "WiburXU"
	_ = GetUserInfo(&info)
}

func GetUserInfo(userInfo *UserData) *UserData {
	return userInfo
}
