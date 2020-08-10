package rest

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/zacharyworks/huddle-data/shared"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

// Handler type to allow errors to be returned.
// (It's usually a mistake to pass back the concrete type of an error
// rather than error, but it's the right thing to do here because ServeHTTP
// is the only place that sees the value and uses its contents.)
type Handler func(http.ResponseWriter, *http.Request) *shared.AppError

// ServeHttp function for Handler impliments http.Handler interface's 'ServeHTTP'
// method, while also handling errors.
func (fn Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	e := fn(w, r) // e is *shared.AppError, not os.Error.

	if e != nil {
		log.Printf("%s : %v", e.Message, e.Error)
		http.Error(w, e.Message, e.Code)
	}
}

type appError struct {
	Error   error
	Message string
	Code    int
}

// Respond is a helper function to simplify REST endpoints.
// It combines serialising data and writing to a response writer, which is
// a common combination of calls.
func respond (w http.ResponseWriter, data interface{}) *shared.AppError {
	responseJSON, err := json.Marshal(data)
	if err != nil {
		return shared.ErrorProcessingJSON(err)
	}

	if _, err = w.Write(responseJSON); err != nil {
		return shared.ErrorFormingResponse(err)
	}

	return nil
}

// GetVarAsInt performs conversion of a url paramater to an int
func getVarAsInt(r *http.Request, v string) (int, *shared.AppError) {
	param, err := strconv.Atoi(mux.Vars(r)[v])
	if err != nil {
		return 0, shared.ErrorProcessingParameter(err)
	}
	return param, nil
}

func readBodyIntoType(requestBody io.ReadCloser, structure interface{}) *shared.AppError{
	body, err := ioutil.ReadAll(requestBody)
	if err != nil {
		return shared.ErrorProcessingBody(err)
	}

	if err = json.Unmarshal(body, &structure); err != nil {
		return shared.ErrorProcessingJSON(err)
	}

	return nil
}