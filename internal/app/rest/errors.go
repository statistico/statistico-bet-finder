package rest

import "errors"

var errBadRequestBody = errors.New("unable to parse request body")
var errBadRequestPath = errors.New("unable to parse request path parameter")
