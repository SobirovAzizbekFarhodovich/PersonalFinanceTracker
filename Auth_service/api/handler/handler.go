package handler

import (
	pb "auth/genprotos"
	"auth/kafka"
	"auth/storage"
	r "auth/storage/redis"
)

type Handler struct {
	UserStorage storage.StorageI
	User     pb.UserServiceClient
	redis    r.InMemoryStorageI
	producer kafka.KafkaProducer
}

func NewHandler(us pb.UserServiceClient, rdb r.InMemoryStorageI, pr kafka.KafkaProducer, userStorage storage.StorageI) *Handler {
	return &Handler{userStorage, us, rdb, pr }
}
