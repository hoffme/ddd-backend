package apps

import (
	"log"
	"sync"

	"github.com/hoffme/ddd-backend/internal/shared/bus"

	"github.com/hoffme/ddd-backend/internal/apps/app"
	"github.com/hoffme/ddd-backend/internal/apps/http"
	"github.com/hoffme/ddd-backend/internal/apps/websocket"
)

func Setup(config Config, buses *bus.Buses) {
	appHttp := http.New(config.Http, buses)
	appWebsocket := websocket.New(config.Websocket, buses)

	var wg sync.WaitGroup

	wg.Add(2)

	go runApp(appHttp, &wg)
	go runApp(appWebsocket, &wg)

	log.Println("Starting apps ...")
	wg.Wait()
	log.Println("Apps closed!")
}

func runApp(app app.App, wg *sync.WaitGroup) {
	defer wg.Done()

	log.Printf("running app (%s) on %s\n", app.Name(), app.Entrypoint())
	err := app.Run()
	if err != nil {
		log.Printf("error in server %s\n", err.Error())
	}
}
