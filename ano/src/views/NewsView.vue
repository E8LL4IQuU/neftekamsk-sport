<script setup lang="ts">
import { ref, onMounted, watch } from "vue";
import { useRoute } from "vue-router";
import axios from "axios";
import { type News } from "@/types/apiTypes";

const url: string = import.meta.env.VITE_ENDPOINT;
const news = ref<News>();
const newsBulk = ref<News[]>([])
const route = useRoute();
let newsID: number = Number(route.params.id);

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

watch(() => route.params.id, (newID) => {
  newsID = Number(newID);
  fetchNews();
  fetchNewsBulk()
})
</script>

<template>
  <div class="lg:grid lg:grid-cols-12 sm:px-0 xl:container mx-auto pb-16">
    <div class="sm:col-span-9 mb-16" v-if="news">
      <img class="sm:w-full sm:h-[600px] object-cover" :src="`${url}/uploads/${news.ImagePath}`" alt="">
      <div class="px-5 pt-3">
        <h1 class="leading-8 text-black text-3xl font-bold mb-3 line-clamp-5">{{ news.Title }}</h1>
        <p class="leading-6 text-black max-w-7xl">{{ news.Description }}</p>
      </div>
    </div>
    <div class="flex flex-col justify-center sm:col-span-3 lg:mt-0 mx-5 lg:mx-2">
      <p class="font-bold text-black mb-0.5">Новости</p>
      <hr class="mb-1">
      <router-link v-for="item in newsBulk" :to="`/news/${item.ID}`">
        <div class="hover:bg-slate-100 transition-colors">
          <div class="flex items-center pb-2 gap-3">
            <img class="h-10 w-10 object-cover flex-shrink-0" :src="`${url}/uploads/${item.ImagePath}`" />
            <!-- FIXME: Clamp text at a reasonable point, for example after a dot -->
            <h4 class="text-black line-clamp-2">{{ item.Title }}</h4>
          </div>
        </div>
      </router-link>
    </div>
  </div>
</template>

<style scoped></style>
