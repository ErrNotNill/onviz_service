
<template>
  <main id="chat">
    <h1>Chat</h1>
    <p>This is chat page</p>
  </main>
</template>
<script>

import axios from "axios";

const url = `http://localhost:9090/leads_get`;
axios.get(url).then(response => (this.ID = response.data.ID));

const socketUrl = `ws://localhost:9090/chat`;

const socket = new WebSocket(socketUrl);

socket.onopen = connectionOpen;
socket.onmessage = messageReceived;
//socket.onerror = errorOccurred;
socket.onopen = connectionClosed;

function connectionOpen() {
  socket.send("UserName:ZXCZXCZXC@gmail.com");
}

function messageReceived(e) {
  var messageLog = document.getElementById("messageLog");
  messageLog.innerHTML += "<br>" + "Ответ сервера: " + e.data;
}

function connectionClosed(e) {
  var messageLog = document.getElementById("messageLog");
  messageLog.innerHTML += "<br>" + "Соединение закрыто";
  socket.disconnect();
}
</script>

