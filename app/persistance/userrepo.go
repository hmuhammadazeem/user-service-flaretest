package persistance

type UserRepo interface {
	IsUsernameExist(username string) (bool, error)
}


type FakeDb struct {
	users []string
}


func (f *FakeDb) GetUserByUsername(username string) (bool, error) {
	for _, val := range f.users {
		if val == username {
			return true, nil
		}
	}
	return false, nil
}