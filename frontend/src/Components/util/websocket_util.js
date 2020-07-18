import ReconnectingWebSocket from "reconnecting-websocket"



// "proxy": "http://localhost:5000"

let connect = () => {
  let socket
  console.log(process.env)
  if (process.env.NODE_ENV === "production") {
    socket = new ReconnectingWebSocket(window.location.protocol.replace("http", "ws") + "//" + window.location.host + "/ws");;
  } else {
    socket = new ReconnectingWebSocket("ws://localhost:5000/ws")
  }

  console.log("Attempting Connection...");

  socket.onopen = () => {
    console.log("Successfully Connected");
    socket.send("test")
  };

  socket.onclose = event => {
    console.log("Socket Closed Connection: ", event);
  };

  socket.onerror = error => {
    console.log("Socket Error: ", error);
  };

  return socket
};




export { connect };