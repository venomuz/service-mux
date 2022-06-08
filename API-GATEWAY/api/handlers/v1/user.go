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
	"strconv"
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
	w.WriteHeader(http.StatusCreated)

	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		if err != nil {
			fmt.Fprintf(w, "Error while marshiling to json")
		}
		h.log.Error("failed to bind json", l.Error(err))
		return
	}
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
	page, err := strconv.ParseInt(queryParams["page"][0], 10, 64)
	if err != nil {
		if err != nil {
			fmt.Fprintf(w, "Error While converting to integer")
		}
	}
	limit, err := strconv.ParseInt(queryParams["limit"][0], 10, 64)
	if err != nil {
		if err != nil {
			fmt.Fprintf(w, "Error While converting to integer")
		}
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()
	response, err := h.serviceManager.UserService().GetList(ctx, &pb.LimitRequest{
		Page:  page,
		Limit: limit,
	})
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		if err != nil {
			fmt.Fprintf(w, "Error while marshiling to json")
		}
		h.log.Error("failed to bind json", l.Error(err))
		return
	}
}
