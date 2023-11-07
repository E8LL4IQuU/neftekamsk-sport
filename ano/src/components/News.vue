<template>
  <div class="pt-5" v-for="(news, index) in NewsData" :key="index">
    <div class="w-[406px] h-[445px] hover:-translate-y-5 duration-300 mobile:w-[300px]">
      <a  :href="news.links">
        <img class="w-[350px]" :src="news.img">
      </a>
      <div class="w-[350px] mobile:w-[300px]">
        <h5 class="text-black font-semibold text-[24px]">{{news.title}}</h5>
        <p class="text-gray-500 font-medium text-[20px]">{{news.description}}</p>
        <span class="text-gray-400 font-normal text-[20px]">Дата: <span>{{news.date}}</span> </span>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
interface INews {
  title: string;
  description: string;
  links: string;
  img: string;
  date: string;
}

const props = defineProps<{
  NewsData: INews[]
}>()
</script>

<style scoped>

</style>