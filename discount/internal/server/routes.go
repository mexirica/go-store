package server

import (
	pb "go-store/discount/pkg/grpc"

	"github.com/sirupsen/logrus"
)

var Logger logrus.Logger

type Server struct {
	pb.DiscountProtoServiceServer
}

