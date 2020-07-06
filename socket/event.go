package socket

import (
	"fmt"
	engineio "github.com/googollee/go-engine.io"
	"github.com/googollee/go-engine.io/transport"
	"github.com/googollee/go-engine.io/transport/polling"
	socketio "github.com/googollee/go-socket.io"
	tws "github.com/googollee/go-engine.io/transport/websocket"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"log"
	"net/http"
)

//初始化websocket
func InitWsSocket(app *iris.Application) {
	mvc.Configure(app.Party("/websocket"), configureWs)
}


func InitSocketIo(app *iris.Application)  {
	pt := polling.Default
	wt := tws.Default
	wt.CheckOrigin = func(req *http.Request) bool {
		return true
	}
	server, err := socketio.NewServer(&engineio.Options{
		Transports: []transport.Transport{
			pt,
			wt,
		},
	})

	if err != nil {
		log.Fatal(err)
	}
	server.OnConnect("/", func(s socketio.Conn) error {
		s.SetContext("")
		fmt.Println("connected:", s.ID())
		return nil
	})

	server.OnEvent("/", "notice", func(s socketio.Conn, msg string) {
		fmt.Println("notice:", msg)
		s.Emit("reply", "have "+msg)
	})
	server.OnEvent("/chat", "msg", func(s socketio.Conn, msg string) string {
		s.SetContext(msg)
		return "recv " + msg
	})
	server.OnEvent("/", "bye", func(s socketio.Conn) string {
		last := s.Context().(string)
		s.Emit("bye", last)
		s.Close()
		return last
	})
	server.OnError("/", func(s socketio.Conn, e error) {
		fmt.Print(e)
	})
	server.OnDisconnect("/", func(s socketio.Conn, msg string) {
		fmt.Println("closed", msg)
	})


	go server.Serve()

	//defer server.Close()
	app.HandleMany("GET POST", "/socket.io/{any:path}", iris.FromStd(server) )
}


