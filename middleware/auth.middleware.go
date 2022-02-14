package middleware

import (
	"anshoryihsan/simple_rest/helper"
	"anshoryihsan/simple_rest/model/web"
	"net/http"
)

type AuthMiddleware struct {
	Handler http.Handler
}

func NewAuthMiddleware(handler http.Handler) *AuthMiddleware {
	return &AuthMiddleware{Handler: handler}
}

func (middleware *AuthMiddleware) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	if "RAHASIA" == request.Header.Get("X-API-Key") {
		//ok
		middleware.Handler.ServeHTTP(writer, request)
	} else {
		//err
		writer.Header().Set("Conten-Type", "application/json")
		writer.WriteHeader(http.StatusUnauthorized)

		webResponse := web.WebResponse{
			Code:   http.StatusUnauthorized,
			Status: false,
		}

		helper.WriteToResponseBody(writer, webResponse)
	}
}
