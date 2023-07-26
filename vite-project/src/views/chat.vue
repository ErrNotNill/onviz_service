<template>
  <main id="chat">
    <h1>Chat</h1>
    <p>This is chat page</p>
  </main>
  <div>{{.data}} </div>
</template>
<script>
import axios from "axios";

const data = axios.get("http://localhost:9090/leads_get")

</script>