package v1

import (
	"context"
	"encoding/json"
	"fmt"
	pb "github.com/venomuz/project5/API-GATEWAY/genproto"
	l "github.com/venomuz/project5/API-GATEWAY/pkg/logger"
	"google.golang.org/protobuf/encoding/protojson"
	"io/ioutil"
	"net/http"
	"time"
)

func (h *handlerV1) CreateUser(w http.ResponseWriter, r *http.Request) {
	var (
		jspbMarshal protojson.MarshalOptions
	)
	body := pb.Useri{}
	jspbMarshal.UseProtoNames = true
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		if err != nil {
			fmt.Fprintf(w, "Kindly enter data with the event title and description only in order to update")
		}
		h.log.Error("failed to bind json", l.Error(err))
		return
	}
	json.Unmarshal(reqBody, &body)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	response, err := h.serviceManager.UserService().Create(ctx, &body)
	if err != nil {
		if err != nil {
			fmt.Fprintf(w, "Kindly enter data with the event title and description only in order to update")
		}
		h.log.Error("failed to bind json", l.Error(err))
		return
	}
	marshal, err := json.Marshal(response)
	if err != nil {
		return
	}
	w.Write(marshal)
}

// GetUser gets user by id
// route /v1/users/{id} [get]
func (h *handlerV1) GetUser(w http.ResponseWriter, r *http.Request) {
	var (
		jspbMarshal protojson.MarshalOptions
	)

	jspbMarshal.UseProtoNames = true
	queryParams := r.URL.Query()
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Got parameter id:%s!\n", queryParams["id"][0])
	fmt.Fprintf(w, "Got parameter category:%s!", queryParams["category"][0])
}
