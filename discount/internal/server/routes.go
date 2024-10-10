package server

import (
	"context"
	"errors"
	"go-store/discount/internal/database/repository"
	"go-store/discount/internal/types"
	pb "go-store/discount/pkg/grpc"

	"github.com/sirupsen/logrus"
)

var Logger logrus.Logger
var repo repository.Repository

type Server struct {
	pb.DiscountProtoServiceServer
}

func (s *Server) GetDiscount(ctx context.Context, req *pb.GetDiscountRequest) (*pb.CouponModel, error) {
	Logger.Info("Getting discount")
	if req == nil {
		return nil, errors.New("request is nil")
	}
	coupon, err := repo.GetDiscount(req.ProductName)
	if err != nil {
		return nil, err
	}

	res := &pb.CouponModel{
		Id:          coupon.ID,
		ProductName: coupon.ProductName,
		Description: coupon.Description,
		Amount:      coupon.Amount,
	}
	return res, nil
}

func (s *Server) CreateDiscount(ctx context.Context, req *pb.CreateDiscountRequest) (*pb.CouponModel, error) {
	Logger.Info("Creating discount")
	if req == nil {
		return nil, errors.New("request is nil")
	}
	coupon := &types.Coupon{
		ProductName: req.Coupon.ProductName,
		Description: req.Coupon.Description,
		Amount:      req.Coupon.Amount,
	}
	coupon, err := repo.CreateDiscount(coupon)
	if err != nil {
		return nil, err
	}

	res := &pb.CouponModel{
		Id:          coupon.ID,
		ProductName: coupon.ProductName,
		Description: coupon.Description,
		Amount:      coupon.Amount,
	}

	return res, nil
}

func (s *Server) UpdateDiscount(ctx context.Context, req *pb.UpdateDiscountRequest) (*pb.CouponModel, error) {
	Logger.Info("Updating discount")
	if req == nil {
		return nil, errors.New("request is nil")
	}
	coupon := &types.Coupon{
		ProductName: req.Coupon.ProductName,
		Description: req.Coupon.Description,
		Amount:      req.Coupon.Amount,
	}
	err := repo.UpdateDiscount(coupon)
	if err != nil {
		return nil, err
	}
	return &pb.CouponModel{
		Id:          coupon.ID,
		ProductName: coupon.ProductName,
		Description: coupon.Description,
		Amount:      coupon.Amount,
	}, nil
}

func (s *Server) DeleteDiscount(ctx context.Context, req *pb.DeleteDiscountRequest) (*pb.DeleteDiscountResponse, error) {
	Logger.Info("Deleting discount")
	if req == nil {
		return &pb.DeleteDiscountResponse{Success: false}, errors.New("request is nil")
	}
	err := repo.DeleteDiscount(req.ProductName)
	if err != nil {
		return &pb.DeleteDiscountResponse{Success: false}, err
	}
	return &pb.DeleteDiscountResponse{Success: true}, nil
}
