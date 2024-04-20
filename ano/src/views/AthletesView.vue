<template>
  <div>
    <div class="container m-auto">
      <div>
        <h1 class="text-black text-[40px] font-light pb-[24px]">Истории людей</h1>
      </div>
      <div v-for="athlete in Athletes" class="flex gap-[17px] pb-6 pc:block">
        <div>
          <img class="max-w-[365px] max-h-[560px] rounded-xl pc:mx-auto mobile:max-w-full"
            :src="`${url}/uploads/${athlete.ImagePath}`">
        </div>
        <div class="text-black pc:text-center">
          <h3 class="pb-[12px] text-[36px] font-bold mobile:text-[24px]">{{ athlete.Title }}</h3>
          <!-- <span class="pb-[12px] text-[#727272] font-medium text-[24px] mobile:text-[18px]">Предприниматель</span> -->
          <p class="max-w-[950px] text-[24px] pc:mx-auto mobile:text-[16px]">{{ athlete.Description }}</p>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from "vue";
import axios from 'axios';
import { type Athlete } from '@/types/apiTypes';

const url: string = import.meta.env.VITE_ENDPOINT
const Athletes = ref<Athlete[]>([])

const fetchAthletes = async (): Promise<void> => {
  try {
    // FIXME: pagination and lazy load
    const response = await axios.get<Athlete[]>(`${url}/api/athletes?limit=50`)
    Athletes.value = response.data
  } catch (error) {
    console.error('Error fetching Athletes: ', error)
  }
}

onMounted(() => {
  fetchAthletes()
})
</script>

<style scoped></style>