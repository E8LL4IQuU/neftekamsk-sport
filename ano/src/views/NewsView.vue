<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { type News } from '@/types/apiTypes'
import axios from 'axios'

const url: string = import.meta.env.VITE_ENDPOINT
const news = ref<News[]>([])

const fetchNews = async (): Promise<void> => {
    try {
        const response = await axios.get<News[]>(`${url}/api/news`)
        news.value = response.data
    } catch (error) {
        console.error('Error fetching news:', error)
    }
}

onMounted(() => {
    fetchNews()
})
</script>

<template>
  <router-link v-for="item in news" :to="`${url}/news/${item.ID}`">
  <div class="containers hover:bg-slate-200 transition-colors">
    <div class="flex justify-between mt-5 py-9 gap-3 laptop:block laptop:text-center mobile:py-[35px]" >
      <div>
          <img class="w-[650px] h-[520px] rounded-[30px] mac:w-[500px] mac:h-[400px] laptop:mx-auto mobile:w-full object-contain" :src="`${url}/uploads/${item.ImagePath}`"/>
        </div>
        <div class="max-w-[665px] mac:max-w-[500px] laptop:mx-auto mobile:w-full">
          <h4 class="text-black text-center text-[36px] pb-[14px] mac:text-[30px]">{{ item.Title }}</h4>
          <hr class="max-w-[540px] bg-[#E4E4E4] mx-auto">
          <p class="text-[#6F6F6F] text-[28px] pt-[14px] text-right  mac:text-[20px] laptop:text-center line-clamp-[10]">{{ item.Description }}</p>
        </div>
      </div>
    </div>
  </router-link>
</template>

<style scoped>
.containers{
  max-width: 1450px;
  margin: 0 auto;
  padding: 0px 30px;
  border-radius: 10px;
}
</style>