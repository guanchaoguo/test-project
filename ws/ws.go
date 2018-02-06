package main

import (
	"net/http"
	"golang.org/x/net/websocket"
	"html/template"
	"fmt"
)

const tpl=`
<html>
<head></head>
<body>
	<script type="text/javascript">
		var sock = null;
		var wsuri = "ws://127.0.0.1:12345";

		window.onload = function() {

			console.log("onload");

			sock = new WebSocket(wsuri);

			sock.onopen = function() {
				console.log("connected to " + wsuri);
			}

			sock.onclose = function(e) {
				console.log("connection closed (" + e.code + ")");
			}

			sock.onmessage = function(e) {
				console.log("message received: " + e.data);
			}
		};

		function send() {
			var msg = document.getElementById('message').value;
			sock.send(msg);
		};
	</script>
	<h1>WebSocket Echo Test</h1>
	<form>
		<p>
			Message: <input id="message" type="text" value="Hello, world!">
		</p>
	</form>
	<button onclick="send();">Send Message</button>
</body>
</html>
`

func Myhandle(w http.ResponseWriter,r *http.Request)  {
	temp,_:= template.New("tpl").Parse(tpl)
	temp.Execute(w,nil)

}

// Echo the data received on the WebSocket.
func EchoServer(ws *websocket.Conn) {
	//io.Copy(ws, ws)
	var err error

	for {
		var reply string

		if err = websocket.Message.Receive(ws, &reply); err != nil {
			fmt.Println("Can't receive")
			break
		}

		fmt.Println("Received back from client: " + reply)

		msg := "Received:  " + reply
		fmt.Println("Sending to client: " + msg)

		if err = websocket.Message.Send(ws, msg); err != nil {
			fmt.Println("Can't send")
			break
		}
	}
}


// This example demonstrates a trivial echo server.
func main() {
	http.Handle("/", websocket.Handler(EchoServer))
	http.Handle("/ws", http.HandlerFunc(Myhandle))
	err := http.ListenAndServe(":12345", nil)
	if err != nil {
		panic("ListenAndServe: " + err.Error())
	}
}