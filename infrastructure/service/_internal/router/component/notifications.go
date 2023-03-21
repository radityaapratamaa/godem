package component

import (
	"net/http"

	context2 "github.com/kodekoding/phastos/go/context"
	"github.com/kodekoding/phastos/go/notifications"
)

func WrapNotif(next http.Handler, obj notifications.Platforms) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		context2.SetNotif(request, obj)
		next.ServeHTTP(writer, request)
	})
}
