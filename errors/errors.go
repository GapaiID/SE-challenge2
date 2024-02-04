package errors

import "errors"

var ErrDbRecordNotFound = errors.New("[database] Record Not Found")
var ErrDbInternalError = errors.New("[database] Internal Error")

var ErrPolicyUnauthorized = errors.New("[policy] you do not have permission to access this resource")
var ErrPolicyForbidden = errors.New("[policy] you do not have permission to access this resource")
