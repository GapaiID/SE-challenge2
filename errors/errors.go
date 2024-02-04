package errors

import "errors"

var ErrDbRecordNotFound = errors.New("record Not Found")
var ErrDbInternalError = errors.New("internal Error")

var ErrPolicyUnauthorized = errors.New("you do not have permission to access this resource")
var ErrPolicyForbidden = errors.New("you do not have permission to access this resource")
