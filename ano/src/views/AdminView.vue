<script setup lang="ts">
import { onMounted, ref } from 'vue'
import axios, { type AxiosResponse } from 'axios'
import { useRouter } from 'vue-router'
import EventManagement from '../components/EventManagement.vue'
import EventsList from '../components/EventsList.vue'

const router = useRouter()
const url: string = import.meta.env.VITE_ENDPOINT

const isLoggedIn = ref<number>(0)

onMounted(async () => {
  await axios
    .get(`${url}/api/auth/user`, {
      withCredentials: true,
      // 200 OK
    })
    .then((response) => {
      if (response.data.id !== 0) {
        isLoggedIn.value = 1
      }
    })
    .catch((error) => {
      if (401 == error.response.status) {
        router.push("/login")
      }
    });
});


// TODO: add event(meropriyatye) create method
// POST igater.burger.moe/api/events
// name         string
// description  string
// image        file

</script>

<template>
  <body class="h-full bg-white">
<main v-if="isLoggedIn">
  <h1 class="text-black text-3xl">
  Admin panel
</h1>
<EventManagement></EventManagement>
<EventsList></EventsList>
<h2 class="text-black text-2xl">

</h2>
</main></body>
</template>

<style>

</style>
