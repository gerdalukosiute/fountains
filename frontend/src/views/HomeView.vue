<script>
export default {
  data: function() {
    return {
      fountains: [],
    }
  },
  methods: {
    async loadFountainsList() {
      try {
        let response = await this.$axios.get("/fountains/", {
          params: {
            latitude: 1,
            longitude: 2,
          }
        });
        this.fountains = response.data;
      } catch (e) {
        alert("Error: " + e);
      }
    }
  }
}
</script>

<template>
  <button @click="loadFountainsList">
    Load Fountains 
  </button>
  <hr />
  <p>List of fountains:</p>
  <ul>
    <li v-for="f in fountains">
      Lat: {{ f.latitude }}
      Lng: {{ f.longitude }}
    </li>
  </ul>
</template>
