<template>
  <div class="pt-5" v-for="(news, index) in props.NewsData" :key="index">
    <!-- FIXME: the whole block shold probably be a link -->
    <!-- FIXME: mobile layout: too big y-gap between news -->
    <div class="w-[406px] h-[445px] hover:-translate-y-5 duration-300 mobile:w-[300px]">
      <router-link  :to="`/news/${news.ID}`">
        <img class="w-[350px] h-[260px] object-cover" :src='`${url}/uploads/${news.ImagePath}`'>
      </router-link>
      <div class="w-[350px] mobile:w-[300px]">
        <div class="flex items-center justify-between">
        <h5 class="text-black font-semibold text-[24px] line-clamp-2">{{news.Title}}</h5>
        <span class="text-gray-400 font-normal text-[20px]"><span>{{ formatTimestamp(news.CreatedAt) }}</span></span>
      </div>
        <p class="text-gray-500 font-medium text-[20px] line-clamp-3">{{news.Description}}</p>
      </div>
    </div>
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

<style scoped>

</style>