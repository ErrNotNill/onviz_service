$(function () {
    let websocket = newWebSocket("ws://localhost:9090/ws");
    let room = $("#chat-text");
    websocket.addEventListener("message", function (e) {
        let data = JSON.parse(e.data);
        let chatContent = `<p><strong>${data.username}</strong>: ${data.text}</p>`;
        room.append(chatContent);
        room.scrollTop = room.scrollHeight; // Auto scroll to the bottom
    });});