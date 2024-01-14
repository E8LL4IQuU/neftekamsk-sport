<script setup lang="ts">
import { ref, onMounted } from 'vue'
import axios from 'axios'
import { useRouter } from 'vue-router'

const router = useRouter()
const url: string = import.meta.env.VITE_ENDPOINT

const message = ref<string>()

onMounted(async () => {
  await axios.post(`${url}/api/auth/logout`, {
    withCredentials: true
  })
  .then(() => {
    message.value = "You've succesfully logged out"
  })
  .catch((error) => {
    message.value = `Error occurred: ${error}`
  })
})
</script>

<template>
<h2>{{ message }}</h2>
</template>

<style>

</style>