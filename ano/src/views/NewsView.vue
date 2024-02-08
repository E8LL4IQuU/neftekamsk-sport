<script setup lang="ts">
import { ref, onMounted } from "vue";
import { useRoute } from "vue-router";
import axios from "axios";
import { type News } from "@/types/apiTypes";

const url: string = import.meta.env.VITE_ENDPOINT;
const news = ref<News>();
const newsBulk = ref<News[]>([])
const route = useRoute();
const newsID: number = Number(route.params.id);

const fetchNewsBulk = async (): Promise<void> => {
  try {
    const response = await axios.get<News[]>(`${url}/api/news?limit=4&exclude=${newsID}`)
    newsBulk.value = response.data
  } catch (error) {
    console.error('Error fetching news:', error)
  }
}

const fetchNews = async (): Promise<void> => {
  try {
    const response = await axios.get<News>(`${url}/api/news/${newsID}`)
    news.value = response.data
  } catch (error) {
    console.error('Error fetching news:', error)
  }
}

onMounted(() => {
  fetchNews()
  fetchNewsBulk()
})
</script>

<template>
  <div class="lg:grid lg:grid-cols-12 sm:px-0 sm:container mx-auto pb-20">
    <div class="sm:col-span-8 gap-x-12 justify-between" v-if="news">
      <img class="sm:w-full sm:h-[600px] object-cover" :src="`${url}/uploads/${news.ImagePath}`" alt="">
      <div class="px-3 pt-3">
        <h1 class="text-black text-3xl font-bold mb-3 line-clamp-5">{{ news.Title }}</h1>
        <p class="leading-6 text-black max-w-7xl">{{ news.Description }}</p>
      </div>
    </div>
    <div class="sm:col-span-4 mt-32 lg:mt-0">
      <p class="text-center text-2xl font-bold text-slate-700 mb-12">Последние новости</p>
      <router-link v-for="item in newsBulk" :to="`${url}/news/${item.ID}`">
        <div class="hover:bg-slate-100 transition-colors">
          <div class="sm:grid sm:grid-cols-12 mt-1 py-4 px-3 gap-3 laptop:block laptop:text-center mobile:py-[35px]">
            <div class="sm:col-span-4">
              <img
                class="sm:w-48 sm:h-32 rounded-[30px] laptop:mx-auto object-cover"
                :src="`${url}/uploads/${item.ImagePath}`" />
            </div>
            <div class="sm:col-span-8 sm:w-fit">
              <h4 class="text-black mb-4 font-bold mac:text-[30px] line-clamp-2">{{ item.Title }}</h4>
              <p
                class="text-[#6F6F6F] mac:text-[20px] laptop:text-center line-clamp-5">
                {{ item.Description }}</p>
            </div>
          </div>
        </div>
      </router-link>
    </div>
  </div>
</template>

<style scoped></style>
