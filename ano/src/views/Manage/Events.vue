<script setup lang="ts">
import { ref, onMounted } from 'vue';
import axios from 'axios'
import { useRouter, RouterLink } from 'vue-router'
import AdminSlider from '@/components/AdminSlider.vue'
import { Event } from '@/types/apiTypes'

const INVALID_USER_ID: number = 0;
const url: string = import.meta.env.VITE_ENDPOINT
const router = useRouter()
const isLoggedIn = ref<number>(0)
const events = ref<Event[]>([]);

const fetchEvents = async (): Promise<void> => {
  try {
    const response = await axios.get<Event[]>(`${url}/api/events`, {
      withCredentials: true,
    })
    events.value = response.data
  } catch (error) {
    console.error('Error fetching events:', error)
  }
}

onMounted(async () => {
  fetchEvents()
})
</script>

<template>
  <!-- FIXME: template spans more than the screen width -->
  <div class="py-8 px-12">
    <div class="flex justify-between">
      <h1 class="text-3xl text-black font-bold">Мероприятия</h1>
      <router-link :to="'/manage/events/create'">
        <button class="bg-black rounded text-white py-2 px-3 ml-auto">Создать</button>
      </router-link>
    </div>
    <AdminSlider :SliderData="events" @reloadSlider="fetchEvents" />
  </div>
</template>

<style scoped></style>
