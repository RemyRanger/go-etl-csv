package http_error

// /!\ THIS ERRORS ARE ONLY SHOWED IN APP LOGS (DEBUGGING PURPOSES), SHOULD NEVER BE SEND TO CLIENT (ONLY HTTP STANDARD ERRORS) /!\

// ********** FIREBASE MIDDLEWARE ERRORS **********

// ErrNoUserFirebaseInContextError : cdv ErrNoUserFirebaseInContextError
var ErrNoUserFirebaseInContextError = "No user in context"

// ErrUnauthorisedEmptyToken : cdv ErrUnauthorisedEmptyToken
var ErrUnauthorisedEmptyToken = "No bearer token provided"

// ErrUnauthorisedUnableVerifyJwt : cdv ErrUnauthorisedUnableVerifyJwt
var ErrUnauthorisedUnableVerifyJwt = "Unable to verify bearer token"

// ErrUnauthorisedUserRole : cdv ErrUnauthorisedUnableVerifyJwt
var ErrUnauthorisedUserRole = "User role unauthorised"

// ********** INTERNAL ERRORS **********

// ErrInternalProcess : cdv ErrInternalProcess
var ErrInternalProcess = "Internal error during process"

// ErrNotFound : cdv ErrNotFound
var ErrNotFound = "Objet not found"

// ErrNotFound : cdv ErrNotFound
var ErrProgramAlreadyLinked = "New Program already linked"

// ErrNotFound : cdv ErrNotFound
var ErrAlreadyExist = "Objet already exist"

// ErrRenderProcess : cdv ErrRenderProcess
var ErrRenderProcess = "Render error during process"

// ********** JWT MIDDLEWARE ERROR **********

// ErrUnauthorisedDateWrongFormat : cdv ErrUnauthorisedDateRequired
var ErrUnauthorisedDateWrongFormat = "Wrong API Date format"

// ErrUnauthorisedWrongAuth : cdv ErrUnauthorisedWrongAuth
var ErrUnauthorisedWrongAuth = "Wrong Auth"

// ErrUnauthorisedWrongOrigin : cdv ErrUnauthorisedWrongOrigin
var ErrUnauthorisedWrongOrigin = "Wrong Origin"

// ErrUnauthorisedExpiredToken : cdv ErrUnauthorisedExpiredToken
var ErrUnauthorisedExpiredToken = "Expired Token"

// ErrUnauthorisedAccessObject : cdv ErrUnauthorisedAccessObject
var ErrUnauthorisedAccessObject = "Object access unauthorised"
