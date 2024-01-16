package rest

//  errPkg "board/utils/errors"
// "net/http"

// func (srv server) err2code(err error) int {

// errPkg.DatabaseError.

//TODO: change err status code
// errorsList := map[error]int{
// 	errPkg.ErrBadRequest:           400,
// 	errPkg.ErrStreamingUnsupported: 401,
// 	errPkg.ErrTeaPot:               402,
// 	errPkg.ErrOauthTimeout:         403,
// 	errPkg.ErrEmailNotVerified:     404,
// 	errPkg.ErrEmailNotProvided:     405,
// 	errPkg.ErrServiceUnavailable:   406,
// }

// val, ok := errorsList[err]

// if ok {
// 	return val
// }

// 	return 200

// }

// func (srv server) httpResponder() {

// writer.Header().Set("Content-Type", "application/json")

// }

// func (srv server) responderErr(w http.ResponseWriter, err error) {
// 	statusCode := srv.err2code(err)
// 	if statusCode == http.StatusInternalServerError {
// 		// if !errors.Is(err, context.Canceled) {
// 		// 	_ = h.logger.Log("err", err)
// 		// }
// 		http.Error(w, "internal server error", http.StatusInternalServerError)
// 		return
// 	}

// 	http.Error(w, err.Error(), statusCode)
// }
