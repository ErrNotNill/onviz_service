<template>
  <main id="chat">
    <h1>Chat</h1>
    <p>This is chat page</p>

    <!-- New block for sent messages -->
    <div class="sent-messages">
      <p
        v-for="(message, index) in messages"
        :key="index"
        :class="{ 'sent-message': message.sent, 'cancelled-message': !message.sent }"
      >
        {{ message.text }}
      </p>
    </div>

    <div class="chat-container">
      <div v-for="(message, index) in messages" :key="index" class="input-block">
        <textarea
          rows="5"
          v-model="message.text"
          placeholder="Enter your text here..."
          @keyup.enter="sendData(index)"
        ></textarea>
        <div class="button-block">
          <div class="button-group">
            <button class="button send-button" @click="sendData(index)">Send</button>
            <button class="button cancel-button" @click="cancel(index)">Cancel</button>
          </div>
        </div>
        <p v-if="message.sent">Sent message: {{ message.text }}</p>
      </div>
    </div>
  </main>
</template>

<script>
export default {
  data() {
    return {
      messages: [
        { text: '', sent: false }
        // Add more initial messages as needed
      ]
    };
  },
  methods: {
    async sendData(index) {
      const message = this.messages[index];
      if (message.text.trim() === '') {
        return; // Don't send empty messages
      }

      // Simulate sending data to the server
      try {
        // Replace with your actual server URL and API endpoint
        const response = await fetch('https://localhost:9090/text_collect', {
          method: 'POST',
          headers: {
            'Content-Type': 'application/json'
          },
          body: JSON.stringify({ message: message.text })
        });

        if (response.ok) {
          message.sent = true;
          message.text = '';
        } else {
          console.error('Error sending message:', response.statusText);
        }
      } catch (error) {
        console.error('Error sending message:', error);
      }
    },
    cancel(index) {
      this.messages[index].text = '';
      this.messages[index].sent = false;
    }
  }
};
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
