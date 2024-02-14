<script setup lang="ts">
import { ref, onMounted } from 'vue';
import axios from 'axios'
import { RouterLink } from 'vue-router'
import AdminSlider from '@/components/Manage/Slider.vue'
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
  <div class="lg:ps-52 lg:pe-1">
    <div class="p-4 lg:py-5 lg:px-12 flex justify-between">
      <h1 class="text-3xl text-black font-bold mb-3">Мероприятия</h1>
      <router-link :to="'/manage/events/create'">
        <button class="bg-black rounded-xl text-xl lg:text-base lg:rounded text-white py-2 px-3 ml-auto">Создать</button>
      </router-link>
    </div>
    <AdminSlider :SliderData="IRLEvents" @reloadSlider="fetchIRLEvents" />
  </div>
</template>

<style scoped></style>
