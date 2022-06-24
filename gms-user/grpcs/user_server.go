package grpcs

import (
	"context"

	userpb "github.com/sy-yoon/krealtors/protos/v1/user"
	"github.com/sy-yoon/krealtors/utils"
	"google.golang.org/protobuf/types/known/emptypb"
	"gorm.io/gorm"
)

type UserServer struct {
	userpb.UserServiceServer
	orm *gorm.DB
}

func (me *UserServer) GetUser(ctx context.Context, req *userpb.GetUserRequest) (*userpb.User, error) {
	user := userpb.User{
		Id: req.Id,
	}
	if err := utils.CheckError(me.orm.Table("user").First(&user)); err != nil {
		return nil, err
	}

	return &user, nil
}

func (me *UserServer) GetUserByEmail(ctx context.Context, req *userpb.GetUserRequest) (*userpb.User, error) {
	user := userpb.User{
		Email: req.Email,
	}
	if err := utils.CheckError(me.orm.Table("user").First(&user)); err != nil {
		return nil, err
	}

	return &user, nil
}

func (me *UserServer) ListUsers(ctx context.Context, req *userpb.ListUsersRequest) (*userpb.ListUsersResponse, error) {
	users := []*userpb.User{}
	if err := utils.CheckError(me.orm.Table("user").Find(&users)); err != nil {
		return nil, err
	}

	response := userpb.ListUsersResponse{
		Users: users,
	}
	return &response, nil
}
func (me *UserServer) CreateUser(ctx context.Context, user *userpb.User) (*userpb.User, error) {
	if err := utils.CheckError(me.orm.Table("user").Omit("created_date", "updated_date").Create(user)); err != nil {
		return nil, err
	}

	return user, nil
}
func (me *UserServer) UpdateUser(ctx context.Context, user *userpb.User) (*userpb.User, error) {
	if err := utils.CheckError(me.orm.Table("user").Save(user)); err != nil {
		return nil, err
	}

	return user, nil
}
func (me *UserServer) DeleteUser(ctx context.Context, req *userpb.DeleteUserRequest) (*emptypb.Empty, error) {
	var user userpb.User
	user.Id = req.Id
	if err := utils.CheckError(me.orm.Table("user").Delete(&user)); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (me *UserServer) AddDBContext(orm interface{}) {
	me.orm = orm.(*gorm.DB)
}
