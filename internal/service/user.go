package service

import "github.com/chubin518/kestrel-layout-basic/internal/model"

type UserService struct {
}

func NewUserService() *UserService {
	return &UserService{}
}

func (us *UserService) List() []*model.User {
	list := make([]*model.User, 0)
	list = append(list, &model.User{
		Id:   1,
		Name: "tom",
	})

	list = append(list, &model.User{
		Id:   2,
		Name: "lily",
	})

	list = append(list, &model.User{
		Id:   3,
		Name: "lucy",
	})
	return list
}
