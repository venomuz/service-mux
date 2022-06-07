package repo

import (
	pb "github.com/venomuz/project5/UserService/genproto"
)

//UserStorageI ...
type UserStorageI interface {
	Create(user *pb.Useri) (*pb.Useri, error)
	GetByID(ID string) (*pb.Useri, error)
	DeleteByID(ID string) (*pb.GetIdFromUserID, error)
	GetAllUserFromDb(empty *pb.Empty) (*pb.AllUser, error)
	GetList(page, limit int64) (*pb.LimitResponse, error)
}
