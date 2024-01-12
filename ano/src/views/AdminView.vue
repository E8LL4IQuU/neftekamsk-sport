<script setup lang="ts">
import { onMounted, ref } from 'vue'
import axios, { type AxiosResponse } from 'axios'
import { useRouter } from 'vue-router'
import EventManagement from '../components/EventManagement.vue'
import EventsList from '../components/EventsList.vue'
import AdminSlider from '../components/AdminSlider.vue'

const router = useRouter()

const isLoggedIn = ref<number>(0)
const events = ref([])

onMounted(async () => {
  await axios
    .get('/api/auth/user', {
      withCredentials: true,
    })
    // 200 OK
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

  await axios
    .get('/api/events', {
      withCredentials: true,
    })
    .then((response) => {
      events.value = response.data
      console.log(response.data)
    })
    .catch((error) => {
      console.log("Error while fetching events:", error)
    })
});

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
<AdminSlider :SliderData="events" />
</main>
</body>
</template>

<style>

</style>
