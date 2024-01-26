<template>
  <div class="pt-5" v-for="(news, index) in props.NewsData" :key="index">
    <div class="w-[406px] h-[445px] hover:-translate-y-5 duration-300 mobile:w-[300px]">
      <!-- TODO: implement, change to router-link -->
      <a  :href="'#'">
        <img class="w-[350px]" :src='`${url}/uploads/${news.ImagePath}`'>
      </a>
      <div class="w-[350px] mobile:w-[300px]">
        <h5 class="text-black font-semibold text-[24px]">{{news.Title}}</h5>
        <p class="text-gray-500 font-medium text-[20px]">{{news.Description}}</p>
        <!-- FIXME: change date display to conform to '@/views/Manage/News.vue' -->
        <span class="text-gray-400 font-normal text-[20px]">Дата: <span>{{news.CreatedAt}}</span> </span>
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

</script>

<style scoped>

</style>