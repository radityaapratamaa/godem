package component

import (
	"net/http"

	"github.com/kodekoding/phastos/go/apps"

	"godem/infrastructure/context"
)

func WrapApp(next http.Handler, obj apps.Slacks) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		context.SetSlackApps(request, obj)
		next.ServeHTTP(writer, request)
	})
}
