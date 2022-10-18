package service

import (
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"web_cheese/internal/biz"

	pb "web_cheese/api/digital_collection/v1"
)

type UserService struct {
	pb.UnimplementedUserServer
	user *biz.UserUsecase
	log  *log.Helper
}

func NewUserService(user *biz.UserUsecase, logger log.Logger) *UserService {
	return &UserService{
		user: user,
		log:  log.NewHelper(logger),
	}
}

func (s *UserService) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CommonReply, error) {
	s.log.Infof("create")
	if req.Password != req.ConfirmPassword {
		return &pb.CommonReply{Code: 500, Msg: "please check password"}, nil
	}
	user := &biz.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	}
	err := s.user.Create(ctx, user)
	if err != nil {
		return &pb.CommonReply{Code: 500, Msg: "create failed"}, nil
	}
	return &pb.CommonReply{Msg: "ok"}, nil
}
func (s *UserService) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.CommonReply, error) {
	s.log.Infof("update")
	user := &biz.User{
		ID:       int32(req.Id),
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	}
	err := s.user.Update(ctx, user)
	if err != nil {
		return &pb.CommonReply{Code: 500, Msg: "update failed"}, nil
	}
	return &pb.CommonReply{Msg: "ok"}, nil
}
func (s *UserService) DeleteUser(ctx context.Context, req *pb.DeleteUserRequest) (*pb.CommonReply, error) {
	s.log.Infof("service-DeleteUser")
	user := &biz.User{
		ID: int32(req.Id),
	}
	err := s.user.Delete(ctx, user)
	if err != nil {
		return &pb.CommonReply{Code: 500, Msg: "delete failed"}, nil
	}
	return &pb.CommonReply{Msg: "ok"}, nil
}
func (s *UserService) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserReply, error) {
	fmt.Println("service-user")
	u, err := s.user.Get(ctx, int32(req.Id))
	if err != nil {
		return nil, err
	}
	return &pb.GetUserReply{
		Id:    u.ID,
		Name:  u.Name,
		Email: u.Email,
	}, nil
}

func (s *UserService) ListUser(ctx context.Context, req *pb.ListUserRequest) (*pb.ListUserReply, error) {
	u, err := s.user.List(ctx)
	if err != nil {
		return nil, err
	}
	res := make([]*pb.UserInfo, 0)
	for i := 0; i < len(u); i++ {
		fmt.Println(u[i])
		user := &pb.UserInfo{
			Id:    u[i].ID,
			Name:  u[i].Name,
			Email: u[i].Email,
		}

		res = append(res, user)
	}

	return &pb.ListUserReply{
		Results: res,
	}, nil
}
