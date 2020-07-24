package shared

import "net/http"

type AppError struct {
	Error   error
	Message string
	Code    int
}

func newInternalError(e error, message string) *AppError {
	return &AppError {
		e,
		message,
		http.StatusInternalServerError,}
}

// ----------------------------------------------------------------------------
// Database Errors factories
// ----------------------------------------------------------------------------

func ErrorInsertingRecord(e error) *AppError {
	return newInternalError(e, "Unable to insert record")
}

func ErrorParsingRecord(e error) *AppError {
	return newInternalError(e, "Unable to parse a record")
}

func ErrorRetrievingRecord(e error) *AppError {
	return newInternalError(e, "Unable to retrieve record")
}

func ErrorUpdatingRecord(e error) *AppError {
	return newInternalError(e, "Unable to update record")
}

func ErrorDeletingRecord(e error) *AppError {
	return newInternalError(e, "Unable to delete record",)
}

// ----------------------------------------------------------------------------
// REST errors
// ----------------------------------------------------------------------------

func ErrorProcessingJSON(e error) *AppError {
	return newInternalError(e, "Unable to process JSON")
}

func ErrorProcessingParamater(e error) *AppError {
	return newInternalError(e, "Unable to process paramater")
}

func ErrorProcessingBody(e error) *AppError {
	return newInternalError(e, "Unable to read request body")
}

func ErrorFormingResponse(e error) *AppError {
	return newInternalError(e, "Unable to form response")
}

