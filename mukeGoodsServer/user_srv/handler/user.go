package handler

import (
	"context"
	"goods/user_srv/model"
	userpb "goods/user_srv/proto/v1"

	"google.golang.org/protobuf/types/known/emptypb"
)

type UserServer struct {
	MysqlDB *model.UserModel
}

func userToUserinfo(user model.User) *userpb.UserInfoResponse {
	userInfoRep := userpb.UserInfoResponse{
		Id:       user.ID,
		Password: user.Password,
		Mobile:   user.Mobile,
		NickName: user.NickName,
		Gender:   user.Gender,
		Role:     int32(user.Role),
	}
	if user.Birthday != nil {
		userInfoRep.BirthDay = uint64(user.Birthday.Unix())
	}
	return &userInfoRep
}

func (s *UserServer) GetUserList(c context.Context, pi *userpb.PageInfo) (*userpb.UserListResponse, error) {
	total, users, err := s.MysqlDB.GetList(int(pi.Pn), int(pi.PSize))
	if err != nil {
		return nil, err
	}
	// userRes := global.DB.Scopes(model.Paginate(int(pi.Pn), int(pi.PSize))).Find(&users)

	rsp := &userpb.UserListResponse{}
	rsp.Totle = total
	for _, user := range users {
		userInfores := userToUserinfo(user)
		rsp.Userinfo = append(rsp.Userinfo, userInfores)
	}

	return rsp, nil
}

func (s *UserServer) GetUserByMobile(c context.Context, mobile *userpb.MobileRequest) (*userpb.UserInfoResponse, error) {
	return &userpb.UserInfoResponse{
		Id:       1,
		NickName: "明世隐",
	}, nil
}
func (s *UserServer) GetUserById(c context.Context, id *userpb.IDRequest) (*userpb.UserInfoResponse, error) {
	return nil, nil
}
func (s *UserServer) CreateUser(c context.Context, cr *userpb.CreateRequest) (*userpb.UserInfoResponse, error) {
	return nil, nil
}
func (s *UserServer) UpdateUser(c context.Context, up *userpb.UpdateRequest) (*emptypb.Empty, error) {
	return nil, nil
}
func (s *UserServer) CheckAPassword(c context.Context, check *userpb.CheckPasswordRequest) (*userpb.CheckResponse, error) {

	return &userpb.CheckResponse{
		Success: true,
	}, nil
}
