<template>
  <main id="Home-page">
    <h1>Home</h1>
    <p>This is the home page</p>
  </main>
  <div class="test">
    <pre>{{ rows }}</pre>
  </div>
</template>

<script>
import axios from 'axios'
export default {
  data() {
    return {
      loading: false,
      rows: [{
        ID: "ID",
      }]
    }
  },
  created() {
    this.getDataFromApi()
  },
  methods: {
    getDataFromApi() {
      this.loading = true
      axios.get('https://onviz-api.ru/leads_get')
        .then(response => {
          this.loading = false
          this.rows = response.data
          this.rows.ID = response.data.ID
        })
        .catch(error => {
          this.loading = false
          console.log(error)
        })
    }
  }
}
</script>