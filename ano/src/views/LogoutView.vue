<!-- FIXME: do we even need this page? maybe just redirect to "/" -->
<script setup lang="ts">
import { onMounted } from 'vue'
import axios from 'axios'

const url: string = import.meta.env.VITE_ENDPOINT

onMounted(async () => {
  await axios.post(`${url}/api/auth/logout`, {
    withCredentials: true
  })
    .catch((error) => {
      console.error(error)
    })
})
</script>

<template>
  <div class="flex items-center justify-center h-screen bg-gradient-to-r from-purple-500 to-pink-500 saturate-[.75]">
    <div class="text-center text-white">
      <h1 class="text-4xl font-bold mb-4">🎉 Успешный выход</h1>
      <p class="text-lg mb-8">Вы успешно вышли из личного кабинета. Спасибо за работу!</p>
      <router-link to="/"
        class="text-xl bg-white text-gray-800 py-2 px-4 rounded-full hover:bg-gray-200 transition duration-300 ease-in-out">На
        главную</router-link>
    </div>
  </div>
</template>

<style scoped></style>