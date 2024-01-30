<script setup lang="ts">
import { ref, onMounted } from 'vue';
import axios from 'axios'
import { RouterLink } from 'vue-router'
import AdminSlider from '@/components/AdminSlider.vue'
import { type IRLEvent } from '@/types/apiTypes'

const url: string = import.meta.env.VITE_ENDPOINT
const IRLEvents = ref<IRLEvent[]>([]);

const fetchIRLEvents = async (): Promise<void> => {
  try {
    const response = await axios.get<IRLEvent[]>(`${url}/api/events`)
    IRLEvents.value = response.data
  } catch (error) {
    console.error('Error fetching IRLevents:', error)
  }
}

onMounted(() => {
  fetchIRLEvents()
})
</script>

<template>
  <div class="py-8 px-12">
    <div class="flex justify-between">
      <h1 class="text-3xl text-black font-bold">Мероприятия</h1>
      <router-link :to="'/manage/events/create'">
        <button class="bg-black rounded text-white py-2 px-3 ml-auto">Создать</button>
      </router-link>
    </div>
    <AdminSlider :SliderData="IRLEvents" @reloadSlider="fetchIRLEvents" />
  </div>
</template>

<style scoped></style>
