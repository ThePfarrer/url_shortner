<script setup lang="ts">
import { ref } from "vue";
import axios from "axios"

interface URL {
  key: string;
  long_url: string;
  short_url: string
}

let id =0

const longUrl = ref("");
const shortUrl = ref("")
const toggle = ref(true)
const hideMyURLs = ref(false)
const myURLs = ref<URL[]>([])

const getURLs = async () => {
    try {
        const response = await axios.get("http://127.0.0.1:5000/api/urls/")
        myURLs.value =  response.data
    } catch (error) {
        console.error("Error sending data:",error)
        return error
    }
}
const removeURL = async (key: string) => {
    try {
        await axios.delete("http://127.0.0.1:5000/"+key)
        myURLs.value = myURLs.value.filter((url: URL) => url.key !== key)
    } catch (error){
        console.error(error)
    }

}
const shortenURL = async () => {
    try {
        const response = await axios.post("http://127.0.0.1:5000/api/urls/", {url: longUrl.value})
        shortUrl.value = response.data.short_url
       if(response.status === 201){
        myURLs.value.push(response.data)
       }
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

const showAllURLs = () => {
    hideMyURLs.value = !hideMyURLs.value

    if (hideMyURLs.value && myURLs.value.length === 0){
        getURLs()
    }

}
</script>

<template>
  <div id="app">
    <h1>URL Shortener</h1>
    <div v-if="!hideMyURLs">
      <form @submit.prevent="shortenURL">
        <input v-model="longUrl" type="text" placeholder="Enter a long URL here" required />
        <p v-if="shortUrl">
          Short URL: <a :href="shortUrl" target="_blank">{{ shortUrl }}</a>
        </p>
        <div class="container">
          <div class="box" v-if="toggle">
            <button type="submit" class="">Shorten URL</button>
          </div>
          <div class="box">
            <button @click="showAllURLs">My URLs</button>
          </div>
          <div class="box" v-if="!toggle">
            <button @click="shortenAnotherURL">Shorten another</button>
          </div>
        </div>
      </form>
    </div>

    <div v-if="hideMyURLs">
      <button @click="showAllURLs">X</button>
      <ul>
        <li v-for="url in myURLs" :key="url" class="list-url">
          <span class="url-name"><a :href="url.short_url" target="_blank">{{ url.short_url }}</a></span>
          <button class="delete-button" @click="removeURL(url.key)">X</button>
        </li>
      </ul>
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
ul {
  list-style-type: none;
  padding: 0;
}
/* List item layout */
.list-url {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 10px;
  border-bottom: 1px solid #ccc;
  text-align: left;
}

/* Name of the item */
.url-name {
  flex: 1;
  font-size: 16px;
}

/* Delete button styling */
.delete-button {
  background-color: red;
  color: white;
}
</style>
