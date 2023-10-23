<template>
  <main id="chat">
    <h1>Chat</h1>
    <p>This is chat page</p>

    <div class="sent-messages">
      <div :class="{ 'sent-message': sent, 'cancelled-message': !sent }">
        <p v-for="receivedMsg in receivedMessages" :key="sendMessage">{{ receivedMsg }}</p>
      </div>
    </div>


    <div class="chat-container">
      <form :action="sendMessage" @click.prevent="onSubmit">
      </form>
      <textarea rows="5" v-model="message"></textarea>
      <div class="button-block">
        <div class="button-group">
          <input class="button send-button" type="submit" value="Send" @click="sendMessage" />
          <button class="button cancel-button" @click="cancelMessage">Cancel</button>
        </div>
      </div>
    </div>

    <!-- New block for sent messages -->
<!--    <div>
      <h1> Request </h1>
      <p> {{ receivedMsg }}</p>
    </div>-->
  </main>
</template>
<script>
export default {
  name: 'App',
  data() {
    return {
      message: '',
      socket: null,
      receivedMessages: [], // Array to store received messages
      sent: false,
      error: false
    }
  },


  mounted() {
    this.$nextTick(function () {
      this.socket = new WebSocket("ws://localhost:9090/chat")
      this.socket.onmessage = (msg) => {
        this.receivedMessages.push(msg.data); // Push the received message into the array
      }
    });
  },


  methods: {
    sendMessage() {
      let msg = {
        "greeting": this.message
      }
      this.socket.send(JSON.stringify(msg))
      this.sent = true
    },
    //socketError(){
    // this.socketError()
    //},
   // quitFromChat(){
    //  this.socket.disconnect()
    //  this.sent = false
   // },
    cancelMessage(){
      this.message = ""
      this.sent = false
    }
  }
}
</script>

<style scoped>
#chat {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  height: 100vh; /* Three times larger than actual content */
}

.sent-messages {
  width: 80%;
  max-width: 800px;
  height: 90vh; /* Three times larger than main block */
  border: 1px solid #ccc;
  border-radius: 5px;
  padding: 10px;
  margin-bottom: 20px;
  overflow-y: auto;
}

.sent-message {
  color: green; /* Sent message color */
  margin-bottom: 5px;
}

.cancelled-message {
  color: red; /* Canceled message color */
  margin-bottom: 5px;
}

.chat-container {
  width: 80%;
  max-width: 800px;
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
}

.input-block {
  width: 100%;
  padding: 10px;
  margin-top: 10px;
  border: 1px solid #ccc;
  border-radius: 5px;
  display: flex;
  flex-direction: column;
  align-items: center;
}

.button-block {
  width: 100%;
  display: flex;
  justify-content: space-between;
  margin-top: 10px;
}

.button-group {
  display: flex;
  align-items: center;
}

.button {
  padding: 10px;
  border: none;
  border-radius: 5px;
  cursor: pointer;
  font-weight: bold;
}

.send-button {
  background-color: #4caf50;
  color: white;
  margin-right: 5px;
}

.cancel-button {
  background-color: #f44336;
  color: white;
}

textarea {
  width: 100%;
}

p {
  margin-top: 10px;
}
</style>