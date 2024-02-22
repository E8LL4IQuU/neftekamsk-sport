<template>
  <div class="max-w-[1250px]  m-auto mb-8 ">
    <div class="flex justify-between items-center p-5">
      <div>
        <h2 class="text-black pt-[14px] pb-[18px] font-light text-4xl mobile:text-3xl">Мероприятия</h2>
      </div>
      <div class="text-black">
        <select v-model="selectedFilter" class="focus-visible:none">
          <option value="all" selected>Все</option>
          <option value="current">Текущие</option>
          <option value="previous">Прошедшие</option>
        </select>
      </div>
    </div>
    <div v-if="IRLEvents.length > 0 || isLoading" class="flex flex-wrap gap-5 justify-center tablet:p-5">
      <router-link v-for="event in IRLEvents" :to="`/events/${event.ID}`" :key="event.ID">
        <div
          v-if="selectedFilter === 'all' || (isStale(event.Date) === false && selectedFilter === 'current') || (isStale(event.Date) && selectedFilter === 'previous')">
          <div style="position: relative;" class="hover:opacity-[85%]">
            <img class="w-[600px] h-[350px] object-cover" :src="`${url}/uploads/${event.ImagePath}`"
              :class="{ 'opacity-50': isStale(event.Date) }">
            <div class="bg-white bg-opacity-[80%] absolute top-2 left-2 rounded-2xl">
              <span class="font-light text-black text-base pb-[5px] pt-[5px] pr-[18px] pl-[18px]">{{ isStale(event.Date) ?
                'Прошедшее' : 'Текущее' }}</span>
            </div>
          </div>
          <div class="text-black flex justify-between max-w-[600px] pb-2 pt-[10px]">
            <div>
              <h3 class="text-xl font-bold uppercase mobile:text-[16px]">{{ event.Title }}</h3>
            </div>
            <div>
              <span class="text-gray-600 text-base font-light mobile:text-[14px]">{{ formatTimestamp(event.Date) }}</span>
            </div>
          </div>
          <div class="text-gray-500">
            <p class="max-w-[600px]">{{ event.Description }}</p>
          </div>
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
const selectedFilter = ref<string>("all")
let isLoading: boolean = true

const formatTimestamp = (timestamp: string): string => {
  const date = new Date(timestamp);

  const options: Intl.DateTimeFormatOptions = {
    year: 'numeric',
    month: '2-digit',
    day: 'numeric',
    hour: 'numeric',
    minute: 'numeric'

  };

  return date.toLocaleDateString("ru-RU", options);
}

const isStale = (date: string): boolean => {
  const yesterday = new Date()
  yesterday.setDate(yesterday.getDate() - 1)

  return yesterday > new Date(date)
}

const fetchIRLEvents = async (): Promise<void> => {
  try {
    const response = await axios.get<IRLEvent[]>(`${url}/api/events`)
    IRLEvents.value = response.data
  } catch (error) {
    console.error('Error fetching IRLEvents: ', error)
  }
  isLoading = false
}

onMounted(() => {
  fetchIRLEvents()
})
</script>

<style scoped></style>