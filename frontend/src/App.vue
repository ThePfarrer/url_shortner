<script setup lang="ts">
import { ref } from 'vue'
import axios from 'axios'

interface URL {
  key: string
  long_url: string
  short_url: string
}

const longUrl = ref('')
const shortUrl = ref('')
const toggle = ref(true)
const hideMyURLs = ref(false)
const myURLs = ref<URL[]>([])

const backendServer = ref(import.meta.env.VITE_API_URL)
const getURLs = async () => {
  try {
    const response = await axios.get(`${backendServer.value}/api/urls/`)
    myURLs.value = response.data
  } catch (error) {
    console.error('Error sending data:', error)
    return error
  }
}
const removeURL = async (key: string) => {
  try {
    await axios.delete(`${backendServer.value}/${key}`)
    myURLs.value = myURLs.value.filter((url: URL) => url.key !== key)
  } catch (error) {
    console.error(error)
  }
}
const shortenURL = async () => {
  try {
    const response = await axios.post(`${backendServer.value}/api/urls/`, { url: longUrl.value })
    shortUrl.value = response.data.short_url
    if (response.status === 201) {
      myURLs.value.push(response.data)
    }
  } catch (error) {
    console.error('Error sending data:', error)
  }
  toggle.value = !toggle.value
}

const shortenAnotherURL = () => {
  longUrl.value = ''
  shortUrl.value = ''
  toggle.value = !toggle.value
}

const showAllURLs = () => {
  hideMyURLs.value = !hideMyURLs.value

  if (hideMyURLs.value && myURLs.value.length === 0) {
    getURLs()
  }
}
</script>

<template>
  <div class="container">
    <h1>URL Shortener</h1>
    <div class="card rounded-20px bg-gray-100 text-dark mb-0">
      <div class="card-body p-0" v-if="!hideMyURLs">
        <div class="py-3 px-3">
          <form @submit.prevent="shortenURL">
            <div class="mb-3">
              <label class="form-label">Shorten a long URL</label>
              <input
                v-model="longUrl"
                type="text"
                class="form-control"
                placeholder="Enter a long URL here"
                required
              />
              <p v-if="shortUrl">
                Short URL: <a :href="shortUrl" target="_blank">{{ shortUrl }}</a>
              </p>
            </div>
            <div
              class="btn-group d-flex gap-2"
              role="group"
              aria-label="Basic mixed styles example"
            >
              <div v-if="toggle">
                <button type="submit" class="btn btn-primary">Shorten URL</button>
              </div>
              <div>
                <button type="button" @click="showAllURLs" class="btn btn-info">My URLs</button>
              </div>
              <div v-if="!toggle">
                <button type="button" @click="shortenAnotherURL" class="btn btn-success">
                  Shorten another
                </button>
              </div>
            </div>
          </form>
        </div>
      </div>
      <div class="card-body" v-if="hideMyURLs">
        <button type="button" class="btn-close" aria-label="Close" @click="showAllURLs"></button>
        <ul class="list-group">
          <li
            v-for="url in myURLs"
            :key="url.short_url"
            class="list-group-item d-flex justify-content-center"
          >
            <span class="url-name"
              ><a :href="url.short_url" target="_blank">{{ url.short_url }}</a></span
            >
            <button type="button" class="btn btn-danger" @click="removeURL(url.key)">X</button>
          </li>
        </ul>
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
