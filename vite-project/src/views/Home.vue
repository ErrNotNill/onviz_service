<template>
  <main id="Home-page">
    <h1>Home</h1>
    <p>This is the home page</p>
  </main>
</template>
<script>

import axios from "axios";

const url = `http://localhost:9090/leads_get`;
//write function to get data from url
export default {
  name: 'Home',
  data() {
    return {
      leads: []
    }
  },
  methods: {
    getLeads() {
      axios.get(url)
      .then(response => {
          this.leads = response.data;
        })
      .catch(error => {
          console.log(error);
        })
    }
  },
  mounted() {
    this.getLeads();
  }
}

</script>