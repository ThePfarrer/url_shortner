<script setup lang="ts">
import { ref } from "vue";
import axios from "axios"

const longUrl = ref("");
const shortUrl = ref("")
const toggle = ref(true)

const shortenURL = async () => {
    try {
        const response = await axios.post("http://127.0.0.1:5000/api/urls/", {url: longUrl.value})
        shortUrl.value = response.data.short_url

    } catch (error) {
        console.error("Error sending data:",error)

    }
    toggle.value = !toggle.value
}

const shortenAnotherURL = () => {
    longUrl.value = ""
    shortUrl.value = ""
    toggle.value = !toggle.value

}
</script>

<template>
  <div id="app">
    <h1>URL Shortener</h1>
    <form @submit.prevent="shortenURL">
      <input v-model="longUrl" type="text" placeholder="Enter a long URL here" required />
      <p v-if="shortUrl">
        Short URL: <a :href="shortUrl" target="_blank">{{ shortUrl }}</a>
      </p>

      <div class="box" v-if="toggle">
        <button type="submit" class="">Shorten URL</button>
      </div>
    </form>
    <div class="container" v-if="!toggle">
      <div class="box">
        <a href="/app/my-urls"> My URLs </a>
      </div>
      <div class="box">
        <button @click="shortenAnotherURL">Shorten another</button>
      </div>
    </div>
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
.container {
  display: flex;
  justify-content: center;
  align-items: center;
  gap: 10px;
}
.box {
  text-align: center;
}
</style>
