package server

import (
	pb "go-store/discount/grpc"

	"github.com/sirupsen/logrus"
)

var Logger logrus.Logger

type Server struct {
	pb.DiscountProtoServiceServer
}

