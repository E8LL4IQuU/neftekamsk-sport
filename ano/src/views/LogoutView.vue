<script setup lang="ts">
import { onMounted } from 'vue'
import axios from 'axios'
import { useRouter } from 'vue-router'

const router = useRouter()

const url: string = import.meta.env.VITE_ENDPOINT

onMounted(async () => {
  await axios.post(`${url}/api/auth/logout`, {
    withCredentials: true
  }).then((response) => {
      console.log(response.data)
  })
})
</script>

<template>

</template>

<style>

</style>