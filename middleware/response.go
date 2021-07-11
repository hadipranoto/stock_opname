package middleware

import (
	"net/http"

	"github.com/labstack/echo"
)



const (
	TextError             string = `error`
	TextOk                string = `ok`
	CodeSuccess           int    = 200
	CodeCreated           int    = 201
	CodeAccepted          int    = 202
	CodeNoContent         int    = 204
	CodeBadRequestError   int    = 400
	CodeUnauthorizedError int    = 401
	CodeDatabaseError     int    = 402
	CodeValidationError   int    = 403
	CodeNotFound          int    = 404
	CodeDuplicateError    int    = 405
	CodeUnProcessable     int    = 422
)

// ResponseHelper ...
type ResponseHelper struct {	
	Context  echo.Context	
	Message  string
	Data     interface{}
	Code     int 
}	

type HTTPResponseHelper struct {
 	//some validator
}

func (h *HTTPResponseHelper) SendResponse(res ResponseHelper) error {		
	return res.Context.JSON(http.StatusOK, map[string]interface{}{
		"code":         res.Code,		
		"code_message": res.Message,
		"data":         res.Data,
	})
}

func (h *HTTPResponseHelper) SendSuccess(c echo.Context, message string, data interface{}) error {
	res := ResponseHelper{c, message, data, CodeSuccess}	
	return h.SendResponse(res)
}

func (h *HTTPResponseHelper) SendBadRequest(c echo.Context, message string, data interface{}) error {
	res := ResponseHelper{c, message, data, CodeBadRequestError}	
	return h.SendResponse(res)
}

func (h *HTTPResponseHelper) SendNoContent(c echo.Context, message string, data interface{}) error {
	res := ResponseHelper{c, message, data, CodeNoContent}	
	return h.SendResponse(res)
}