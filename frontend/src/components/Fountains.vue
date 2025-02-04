<script setup>
import { ref, onMounted } from 'vue'
import axios from 'axios'

const fountains = ref([])
const newFountain = ref({ state: 'good', latitude: '', longitude: '' })
const isLoading = ref(false)
const error = ref(null)

const fetchFountains = async () => {
  isLoading.value = true
  error.value = null
  try {
    const response = await axios.get('http://localhost:8080/fountains')
    fountains.value = response.data || []
  } catch (e) {
    console.error('Error fetching fountains:', e)
    error.value = 'Failed to fetch fountains. Please try again.'
  } finally {
    isLoading.value = false
  }
}

const addFountain = async () => {
  isLoading.value = true
  error.value = null
  try {
    const response = await axios.post('http://localhost:8080/fountains', newFountain.value)
    if (response.data) {
      if (!Array.isArray(fountains.value)) {
        fountains.value = []
      }
      fountains.value.push(response.data)
      newFountain.value = { state: 'good', latitude: '', longitude: '' }
    } else {
      throw new Error('No data received from server')
    }
  } catch (e) {
    console.error('Error adding fountain:', e)
    error.value = 'Failed to add fountain. Please try again.'
  } finally {
    isLoading.value = false
  }
}

const updateFountainState = async (fountain) => {
  isLoading.value = true
  error.value = null
  try {
    const updatedFountain = { ...fountain, state: fountain.state === 'good' ? 'faulty' : 'good' }
    const response = await axios.put(`http://localhost:8080/fountains/${fountain.id}`, updatedFountain)
    if (response.data) {
      const index = fountains.value.findIndex(f => f.id === fountain.id)
      if (index !== -1) {
        fountains.value[index] = response.data
      }
    } else {
      throw new Error('No data received from server')
    }
  } catch (e) {
    console.error('Error updating fountain:', e)
    error.value = 'Failed to update fountain. Please try again.'
  } finally {
    isLoading.value = false
  }
}

const deleteFountain = async (id) => {
  isLoading.value = true
  error.value = null
  try {
    await axios.delete(`http://localhost:8080/fountains/${id}`)
    fountains.value = fountains.value.filter(f => f.id !== id)
  } catch (e) {
    console.error('Error deleting fountain:', e)
    error.value = 'Failed to delete fountain. Please try again.'
  } finally {
    isLoading.value = false
  }
}

onMounted(fetchFountains)
</script>

<template>
  <div class="fountains">
    <div v-if="error" class="text-red-500 mb-4">{{ error }}</div>
    <form @submit.prevent="addFountain" class="mb-4">
      <div class="flex space-x-2">
        <select v-model="newFountain.state" class="border p-2 rounded">
          <option value="good">Good</option>
          <option value="faulty">Faulty</option>
        </select>
        <hr>
        <input v-model="newFountain.latitude" type="number" step="0.000001" placeholder="Latitude" class="border p-2 rounded" required>
        <hr>
        <input v-model="newFountain.longitude" type="number" step="0.000001" placeholder="Longitude" class="border p-2 rounded" required>
        <hr>
        <button id="adding" type="submit" class="bg-blue-500 text-white p-2 rounded" :disabled="isLoading">
          {{ isLoading ? 'Adding...' : 'Add Fountain' }}
        </button>
      </div>
    </form>
    <button id="refresh" @click="fetchFountains" class="bg-green-500 text-white p-2 rounded mb-4" :disabled="isLoading">
      {{ isLoading ? 'Loading...' : 'Refresh Fountains' }}
    </button>
    <br>
    <h2 id="list" class="text-xl font-semibold mb-4">Fountain List:</h2>
    <ul v-if="fountains.length" class="space-y-2">
      <li v-for="fountain in fountains" :key="fountain.id" class="border p-2 rounded flex justify-between items-center">
        <span>
            ID: {{ fountain.id }} | State: {{ fountain.state }} | Lat: {{ fountain.latitude }} | Lng: {{ fountain.longitude }}
        </span>
        <br>
        <button id="state" @click="updateFountainState(fountain)" class="bg-yellow-500 text-white p-1 rounded" :disabled="isLoading">
          Toggle State
        </button>
        <br>
        <button id="delete" @click="deleteFountain(fountain.id)" class="bg-red-500 text-white p-1 rounded" :disabled="isLoading">
            Delete
        </button>
      </li>
    </ul>
    <p v-else-if="!isLoading" class="text-gray-500">No fountains found.</p>
  </div>
</template>

<style>
#state{
    color: rgb(237, 121, 200);
    font-size: 17px;
    font-weight: bold;
    font-family: Verdana, Geneva, Tahoma, sans-serif;
}

#delete{
    color: rgb(245, 15, 57);
    font-size: 17px;
    font-weight: bold;
    font-family: Verdana, Geneva, Tahoma, sans-serif;
}

#adding{
    color: rgb(237, 121, 200);
    font-size: 30px;
    font-weight: bold;
    font-family: Verdana, Geneva, Tahoma, sans-serif;
}

#refresh{
    color: rgb(121, 242, 137);
    font-size: 30px;
    font-weight: bold;
    font-family: Verdana, Geneva, Tahoma, sans-serif;   
}

#list{
    color: black;
    font-size: 30px;
    font-weight: bold;
    font-family: Verdana, Geneva, Tahoma, sans-serif;   
}

</style>