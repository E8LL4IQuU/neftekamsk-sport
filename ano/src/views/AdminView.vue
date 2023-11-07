<script setup lang="ts">
import { onMounted } from 'vue'
import axios, { type AxiosResponse } from 'axios'
import { useRouter } from 'vue-router'

const router = useRouter()

const url: string = import.meta.env.VITE_ENDPOINT

let bebra: string = ""

onMounted(async () => {
  await axios.get(`${url}/api/healthcheck`, {
    withCredentials: true
  }).then((response) => {
      console.log(response.data)
      bebra = response.data.name
  })
})
</script>

<template>
<h2>
  Hello, {{ bebra }}
</h2>
</template>

<style>

</style>