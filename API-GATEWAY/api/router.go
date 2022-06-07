package api

import (
	"github.com/gorilla/mux"
	v1 "github.com/venomuz/project5/API-GATEWAY/api/handlers/v1"
	"github.com/venomuz/project5/API-GATEWAY/config"
	"github.com/venomuz/project5/API-GATEWAY/pkg/logger"
	"github.com/venomuz/project5/API-GATEWAY/services"
)

type Option struct {
	Conf           config.Config
	Logger         logger.Logger
	ServiceManager services.IServiceManager
}

func New(option Option) *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	handlerV1 := v1.New(&v1.HandlerV1Config{
		Logger:         option.Logger,
		ServiceManager: option.ServiceManager,
		Cfg:            option.Conf,
	})
	router.HandleFunc("/get", handlerV1.GetUser).Methods("GET")
	router.HandleFunc("/creat", handlerV1.CreateUser).Methods("POST")

	// api.PUT("/users/:id", handlerV1.UpdateUser)
	// api.DELETE("/users/:id", handlerV1.DeleteUser)

	return router

}
