<template>
  <main id="Home-page">
    <h1>Home</h1>
    <p>This is the home page</p>
    <!-- Filter dropdown -->
    <div>
      <label for="filter">Filter by:</label>
      <select id="filter" v-model="selectedFilter">
        <option value="all">All</option>
        <option value="Евгений">Евгений</option>
        <!-- Add more filter options here as needed -->
      </select>
      <button @click="applyFilter">Apply Filter</button>
    </div>

    <div class="table-container">
      <ul class="table">
        <li v-for="item in jsonArray" :key="item.id" class="list-item">
          <div class="button-container">
            <button @click="toggleMenu(item.ID) " class="table-button" :class="{ 'highlighted': item.name === 'Евгений' }">
              {{ item.ID }}
            </button>
          </div>
          <div v-if="activeItem === item.ID" class="item-details">
            <p>Name: {{ item.name }}</p>
            <p>Phone: {{ item.phone }}</p>
            <p>Assigned By Lead: {{ item.assignedByLead }}</p>
          </div>
        </li>
      </ul>
    </div>
  </main>
</template>

<script>
import axios from "axios";

export default {
  data() {
    return {
      jsonArray: [],
      activeItem: null,
      selectedFilter: "all", // Default to "All"
    };
  },
  created() {
    axios
      .get("http://localhost:9090/leads_get")
      .then((response) => {
        this.jsonArray = response.data; // Assign the JSON array to a data property
      })
      .catch((error) => {
        console.error("Error fetching data:", error);
      });
  },
  computed: {
    filteredItems() {
      // Filter items based on the selected filter criteria
      if (this.selectedFilter === "all") {
        return this.jsonArray; // Return all items
      } else {
        return this.jsonArray.filter((item) => item.name === this.selectedFilter);
      }
    },
  },
  methods: {
    toggleMenu(ID) {
      // Toggle the active item based on its ID
      this.activeItem = this.activeItem === ID ? null : ID;
    },
    applyFilter() {
      // Apply the selected filter
      if (this.selectedFilter === "Евгений") {
        // Set the selected filter to "Евгений" to display all "Евгений" items
        this.selectedFilter = "Евгений";
      }
      // You can update the displayed items here if needed
    },
  },
};
</script>


<style>
/* Style for the table button */
.table-button {
  background-color: green;
  border: none;
  color: white;
  padding: 5px 10px;
  margin: 5px;
  cursor: pointer;
  text-align: center;
}

/* Style for the item details */
.item-details {
  margin-top: 10px;
  background-color: lightgray;
  padding: 5px;
  border-radius: 5px;
}

/* Remove bullet points from list items */
.list-item {
  list-style-type: none;
  margin-left: 0;
  padding-left: 0;
  background-color: transparent; /* Ensure consistent background */
}

/* Highlighted background color for 'Евгений' */
.highlighted {
  background-color: red;
}

/* Position the table on the left side */
/* Position the table on the left side */
.table-container {
  text-align: center ;
  width: 50%;
  margin: 0 auto;
}
.table li {
  float: left;
}

</style>

