<template>
  <main id="Home-page">
    <h1>Home</h1>
    <p>This is the home page</p>
  </main>



</template>


<script lang="ts">
import { defineComponent } from 'vue'
import axios from 'axios'
interface AnimalFacts {
  text: string
}
export default defineComponent({
  name: 'AnimalFacts',
  data() {
    return {
      catFacts: [] as AnimalFacts[],
      fetchingFacts: false
    }
  },
  methods: {
    async fetchCatFacts() {
      const catFactsResponse = await axios.get<AnimalFacts[]>('https://cat-fact.herokuapp.com/facts/random?animal_type=cat&amount=5')
      this.catFacts = catFactsResponse.data
    },
    async loadMoreFacts() {
      this.fetchingFacts = true
      const catFactsResponse = await axios.get<AnimalFacts[]>('https://cat-fact.herokuapp.com/facts/random?animal_type=cat&amount=5')
      this.catFacts.push(...(catFactsResponse.data || []))

      this.fetchingFacts = false
    }
  },
  async mounted() {
    await this.fetchCatFacts()
  }
})
</script>