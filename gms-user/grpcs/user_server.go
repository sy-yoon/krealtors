package grpcs

import (
	"context"

	"github.com/sy-yoon/krealtors/gms"
	userpb "github.com/sy-yoon/krealtors/protos/v1/user"
	"google.golang.org/protobuf/types/known/emptypb"
	"xorm.io/xorm"
)

type UserServer struct {
	userpb.UserServiceServer
	orm *xorm.Engine
}

func (me *UserServer) GetUser(ctx context.Context, req *userpb.GetUserRequest) (*userpb.User, error) {
	user := userpb.User{}
	user.UserId = req.UserId
	_, err := me.orm.Get(&user)
	if err != nil {
		gms.Logger.Error("DB", "SQL", err)
		return nil, err
	}
	return &user, nil
}

func (me *UserServer) ListUsers(ctx context.Context, req *userpb.ListUsersRequest) (*userpb.ListUsersResponse, error) {
	users := []*userpb.User{}
	err := me.orm.Find(&users)
	if err != nil {
		gms.Logger.Error("DB", "SQL", err)
		return nil, err
	}

	response := userpb.ListUsersResponse{
		Users: users,
	}
	return &response, nil
}
func (me *UserServer) CreateUser(ctx context.Context, user *userpb.User) (*userpb.User, error) {
	_, err := me.orm.InsertOne(user)
	if err != nil {
		gms.Logger.Error("DB", "SQL", err)
		return nil, err
	}
	return user, nil
}
func (me *UserServer) UpdateUser(ctx context.Context, user *userpb.User) (*userpb.User, error) {
	_, err := me.orm.Update(user)
	if err != nil {
		gms.Logger.Error("DB", "SQL", err)
		return nil, err
	}
	return user, nil
}
func (me *UserServer) DeleteUser(ctx context.Context, req *userpb.DeleteUserRequest) (*emptypb.Empty, error) {
	var user userpb.User
	_, err := me.orm.Where("cntry_id", req.UserId).Delete(&user)
	if err != nil {
		gms.Logger.Error("DB", "SQL", err)
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

func (me *UserServer) AddDBContext(orm interface{}) {
	me.orm = orm.(*xorm.Engine)
}
