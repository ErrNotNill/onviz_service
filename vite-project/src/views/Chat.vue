
<template>
  <main id="chat">
    <h1>Chat</h1>
    <p>This is chat page</p>
  </main>
  <body>
  <div>
    <input type="text" id="messageInput" />
    <button onclick="sendMessage()">Send Message</button>
  </div>
  </body>
</template>

<script>
const socket = new WebSocket("ws://localhost:9090/chat");

socket.onopen = function () {
  console.log("WebSocket connection established.");
};

socket.onmessage = function (event) {
  const messageDiv = document.getElementById("message");
  messageDiv.innerText = event.data;
};

socket.onclose = function (event) {
  console.log("WebSocket connection closed:", event);
};

socket.onerror = function (error) {
  console.error("WebSocket error:", error);
};

function sendMessage() {
  const inputElement = document.getElementById("messageInput");
  const message = inputElement.value;

  if (message.trim() !== "") {
    socket.send(message);
    inputElement.value = "";
  }
}
</script>

