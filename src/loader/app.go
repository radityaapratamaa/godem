package loader

import (
	"embed"
	"os"
	"strconv"

	"github.com/kodekoding/phastos/v2/go/api"
)

type (
	AppOptions func(*app)
	app        struct {
		*api.App
		templateFolder embed.FS
	}
)

func NewApp(options ...AppOptions) *app {
	port, _ := strconv.Atoi(os.Getenv("SERVER_PORT"))
	writeTimeout, _ := strconv.Atoi(os.Getenv("SERVER_WRITE_TIMEOUT"))
	readTimeout, _ := strconv.Atoi(os.Getenv("SERVER_READ_TIMEOUT"))
	apiTimeout, _ := strconv.Atoi(os.Getenv("SERVER_API_TIMEOUT"))
	godemApp := api.NewApp(
		api.WithAppPort(port),
		api.WriteTimeout(writeTimeout),
		api.ReadTimeout(readTimeout),
		api.WithAPITimeout(apiTimeout),
	)

	initApp := &app{App: godemApp}
	initApp.Init()

	for _, option := range options {
		option(initApp)
	}
	initApp.loadResource()
	return initApp
}

func WithFolderTemplate(tmpl embed.FS) AppOptions {
	return func(a *app) {
		a.templateFolder = tmpl
	}
}

func (a *app) loadResource() {
	loadDataSources()
	a.loadModules()
}
