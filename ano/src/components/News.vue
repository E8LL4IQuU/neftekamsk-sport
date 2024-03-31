<template>
  <div v-for="(news, index) in props.NewsData" :key="index">
    <!-- FIXME: When hovering the cursor over the bottom of the div, there's jittery movement -->
    <router-link :to="`/news/${news.ID}`">
      <div class="hover:-translate-y-5 duration-300 lg:mx-3 mb-8">
        <img class="w-full lg:w-[350px] h-[260px] object-cover" :src='`${url}/uploads/${news.ImagePath}`'>
        <div class="w-full lg:w-[350px] px-2 lg:px-0.5">
          <div class="flex items-center justify-between">
            <h5 class="text-start text-black font-semibold text-[24px] line-clamp-2 leading-6">{{ news.Title }}</h5>
            <span class="text-gray-400 font-normal text-[20px]">
              {{ formatTimestamp(news.CreatedAt) }}
            </span>
          </div>
          <p class="text-start text-gray-500 font-medium text-[20px] line-clamp-3 leading-6">{{ news.Description }}</p>
        </div>
      </div>
    </router-link>
  </div>
</template>

<script setup lang="ts">
import { type News } from '@/types/apiTypes'

const url: string = import.meta.env.VITE_ENDPOINT
const props = defineProps<{
  NewsData: News[]
}>()

const formatTimestamp = (timestamp: number): string => {
  const date = new Date(timestamp * 1000);

  // Get the time zone offset in minutes
  const timeZoneOffset = date.getTimezoneOffset();

  // Use the offset to create a new date with the local time zone
  const localDate = new Date(date.getTime() - timeZoneOffset * 60 * 1000);

  const options: Intl.DateTimeFormatOptions = {
    year: 'numeric',
    month: '2-digit',
    day: 'numeric',
  };

  return localDate.toLocaleString('ru-RU', options);
};

</script>

<style scoped></style>