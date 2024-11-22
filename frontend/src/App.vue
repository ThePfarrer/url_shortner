<script setup lang="ts">
import { ref } from "vue";

const longUrl = ref("");
</script>

<template>
  <div id="app">
    <h1>URL Shortener</h1>
    <form>
      <input v-model="longUrl" type="text" placeholder="Enter a long URL here" required />
      <button type="submit">Shorten URL</button>
    </form>
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
