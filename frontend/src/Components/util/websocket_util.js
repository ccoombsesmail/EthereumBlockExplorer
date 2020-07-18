



// "proxy": "http://localhost:5000"

let connect = () => {
  let socket
  if (process.env.NODE_ENV === "production") {
    socket = new WebSocket("wss://eth-blockexplorer-go.uc.r.appspot.com/ws");
  } else {
    socket = new WebSocket("ws://localhost:5000/ws")
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