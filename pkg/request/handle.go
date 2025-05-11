package request

import (
	"net/http"
	"order-server/pkg"
	"order-server/pkg/response"
)

func HandleBody[T any](w http.ResponseWriter, r *http.Request) (*T, error) {
	body, err := pkg.Decode[T](r.Body)
	if err != nil {
		response.WriteError(w, http.StatusBadRequest, err.Error())
		return nil, err
	}
	err = pkg.IsValid(body)
	if err != nil {
		response.WriteError(w, http.StatusBadRequest, err.Error())
		return nil, err
	}
	return &body, nil
}
