package response

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/AthanatiusC/pizza-hub/helper/logger"
	"github.com/AthanatiusC/pizza-hub/model"
)

func Success(ctx context.Context, w http.ResponseWriter, data interface{}, message ...interface{}) {
	defer logger.InfoContext(ctx, "response success returned", data)
	var response model.Response
	response.Data = data
	for _, msg := range message {
		response.Message += msg.(string)
	}

	if response.StatusCode == 0 {
		response.StatusCode = http.StatusOK
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(response.StatusCode)
	json.NewEncoder(w).Encode(response)
}

func Error(ctx context.Context, w http.ResponseWriter, httpCode int, message ...interface{}) {
	defer logger.InfoContext(ctx, "response error returned", message)
	var response model.Response
	for _, msg := range message {
		response.Message += msg.(string)
	}

	if httpCode == 0 {
		httpCode = http.StatusInternalServerError
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpCode)
	json.NewEncoder(w).Encode(response)
}
