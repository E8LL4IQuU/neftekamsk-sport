<script setup lang="ts">
import { onMounted, ref } from 'vue'
import axios from 'axios'
import { useRouter } from 'vue-router'
import AddEvent from '@/components/AddEvent.vue'
import AdminSlider from '@/components/AdminSlider.vue'

const url: string = import.meta.env.VITE_ENDPOINT
const router = useRouter()

const isLoggedIn = ref<number>(0)
const events = ref([])

const getEvents = async (): Promise<void> =>
  await axios
    .get(`${url}/api/events`, {
      withCredentials: true,
    })
    .then((response) => {
      events.value = response.data
    })
    .catch((error) => {
      console.log("Error while fetching events:", error)
    })

onMounted(async (): Promise<void> => {
  await axios
    .get(`${url}/api/auth/user`, {
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

  getEvents()
});

</script>

<template>
  <body class="h-full bg-white">
    <main v-if="isLoggedIn">
      <AddEvent @reloadSlider="getEvents"></AddEvent>
      <AdminSlider :SliderData="events" @reloadSlider="getEvents" />
    </main>
  </body>
</template>

<style></style>
