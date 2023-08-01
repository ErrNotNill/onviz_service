
<template>
  <main id="Home-page">
    <h1>Home</h1>
    <p>TThis is the home page</p>
  </main>
  <div class="test">
    <td>{{ rows }}</td>
  </div>
</template>

<script>
import axios from 'axios'
export default {
  data() {
    return {
      loading: false,
      rows: []
    }
  },
  created() {
    this.getDataFromApi()
  },
  methods: {
    getDataFromApi() {
      this.loading = true
      axios.get('http://localhost:9090/leads_get')
        .then(response => {
          this.loading = false
          this.rows = response.data
        })
        .catch(error => {
          this.loading = false
          console.log(error)
        })
    }
  }
}
</script>