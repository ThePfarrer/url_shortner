<script setup lang="ts">
import { ref } from "vue";
import axios from "axios"

const longUrl = ref("");
const shortUrl = ref("")

const shortenURL = async () => {
    try {
        const response = await axios.post("http://127.0.0.1:5000/api/urls/", {url: longUrl.value})
        shortUrl.value = response.data.short_url

    } catch (error) {
        console.error("Error sending data:",error)

    }
}
</script>

<template>
  <div id="app">
    <h1>URL Shortener</h1>
    <form @submit.prevent="shortenURL">
      <input v-model="longUrl" type="text" placeholder="Enter a long URL here" required />
      <button type="submit">Shorten URL</button>
    </form>
    <p v-if="shortUrl">
      Short URL: <a :href="shortUrl" target="_blank">{{ shortUrl }}</a>
    </p>
  </div>
</template>

<style scoped>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  text-align: center;
  margin-top: 50px;
}
input {
  margin: 5px;
  padding: 8px;
  width: 300px;
}
button {
  padding: 8px 15px;
  background-color: #007bff;
  color: white;
  border: none;
  cursor: pointer;
}
button:hover {
  background-color: #0056b3;
}
</style>
