<template>
  <div class="max-w-[1250px]  m-auto mb-8 ">
    <div class="flex justify-between items-center p-5">
      <div>
        <h2 class="text-black pt-[14px] pb-[18px] font-light text-4xl mobile:text-3xl">Мероприятия</h2>
      </div>
      <div class="text-black">
      </div>
    </div>
    <div v-if="IRLEvents.length > 0" class="flex flex-wrap gap-5 justify-center tablet:p-5">
      <router-link v-for="(event, index) in IRLEvents" to="#" :key="index" >
      <div style="position: relative;" class="hover:opacity-[85%]">
        <img class="w-[600px] h-[350px] object-cover" :src="`${url}/uploads/${event.img}`">
      </div>
      <div class="text-black flex justify-between max-w-[600px] pb-2 pt-[10px]">
        <div>
          <h3 class="text-xl font-bold uppercase mobile:text-[16px]">{{ event.title }}</h3>
        </div>
      </div>
      <div class="text-gray-500">
        <p class="max-w-[600px]">{{ event.description }}</p>
      </div>
    </router-link>
    </div>
    <div v-else class="text-gray-500 text-6xl">Мероприятия не найдены</div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from "vue";
import axios from 'axios';
import { type IRLEvent } from '@/types/apiTypes'

const url: string = import.meta.env.VITE_ENDPOINT
const IRLEvents = ref<IRLEvent[]>([])

const fetchIRLEvents = async (): Promise<void> => {
  try {
    const response = await axios.get<IRLEvent[]>(`${url}/api/events`)
    IRLEvents.value = response.data
  } catch (error) {
    console.error('Error fetching IRLEvents: ', error)
  }
}

onMounted(() => {
  fetchIRLEvents()
})
</script>

<style scoped>

</style>