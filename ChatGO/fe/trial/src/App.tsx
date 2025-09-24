import { useEffect, useState } from "react";
import "./App.css";

function App() {
  const [count, setCount] = useState("WELC");
  useEffect(() => {
    const socket = new WebSocket("ws://localhost:5001/ws");
    if (socket) {
      socket.onopen = () => {
        setCount("Connected");
        socket.send("ping");
      };
      socket.onmessage = (ev) => {
        setCount(ev.data);
      };
      
    }
  });
  return <div>{count}</div>;
}

export default App;
