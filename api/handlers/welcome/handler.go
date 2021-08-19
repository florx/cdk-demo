package welcome

import (
	"encoding/json"
	"net/http"

	"github.com/florx/cdk-demo/api/respond"
	"go.uber.org/zap"
)

type WelcomeRequest struct {
	Email string `json:"email,omitempty"`
}

type WelcomeResponse struct {
	HelloWorld string `json:"helloWorld,omitempty"`
}

func NewHandler(log *zap.Logger) (h Handler) {
	h.Log = log
	return
}

type Handler struct {
	Log *zap.Logger
}

func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		h.Post(w, r)
		return
	}
}

func (h Handler) Post(w http.ResponseWriter, r *http.Request) {
	d := json.NewDecoder(r.Body)
	var req WelcomeRequest
	err := d.Decode(&req)
	if err != nil {
		errmsg := "invalid request"
		h.Log.Error(errmsg, zap.Error(err))
		respond.WithError(w, errmsg, http.StatusBadRequest)
		return
	}

	valid, err := Validate(req)

	if !valid {
		errmsg := "invalid request"
		h.Log.Error(errmsg, zap.Error(err))
		respond.WithError(w, errmsg, http.StatusBadRequest)
		return
	}

	response := WelcomeResponse{
		HelloWorld: "testtest!",
	}

	respond.WithJSON(w, response)
}

func Validate(req WelcomeRequest) (valid bool, err error) {

	if req.Email == "" {
		return false, err
	}

	return true, err
}
