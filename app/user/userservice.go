package user

import (
	"github.com/hmuhammadazeem/user-service/app/persistance"
	"github.com/hmuhammadazeem/user-service/app/bloom"
)

type UserService struct {
	userrepo persistance.UserRepo
	filter   bloom.Filter
}


func (u *UserService) UsernameExists(username string) (bool, error) {
	if exists := u.filter.MayExist(username); !exists {
		return false, nil
	}
	
	exists, err := u.userrepo.IsUsernameExist(username)
	if err != nil {
		return false, err
	}

	// update bloom filter
	if !exists {
		u.filter.Add(username)
	}

	return exists, nil
}

func NewUserService(repo persistance.UserRepo, filter bloom.Filter) UserService {
	return UserService{repo, filter}
}