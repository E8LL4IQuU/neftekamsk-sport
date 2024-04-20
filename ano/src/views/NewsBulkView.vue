<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { type News } from '@/types/apiTypes'
import axios from 'axios'

const url: string = import.meta.env.VITE_ENDPOINT
const news = ref<News[]>([])

const fetchNewsBulk = async (): Promise<void> => {
  try {
    // FIXME: pagination and lazy load
    const response = await axios.get<News[]>(`${url}/api/news?limit=50`)
    news.value = response.data
  } catch (error) {
    console.error('Error fetching news:', error)
  }
}

onMounted(() => {
  fetchNewsBulk()
})
</script>

<template>
<body class="bg-white">
  <router-link v-for="item in news" :to="`/news/${item.ID}`">
    <div class="containers hover:bg-slate-200 transition-colors">
      <div class="flex justify-between mt-5 py-4 laptop:block laptop:text-center mobile:py-[35px]">
        <div>
          <img
            class="w-[350px] h-[260px] md:w-[600px] md:h-[440px] rounded laptop:mx-auto object-cover"
            :src="`${url}/uploads/${item.ImagePath}`" />
        </div>
        <div class="max-w-[665px] mac:max-w-[500px] laptop:mx-auto mobile:w-full">
          <h4 class="text-black text-center text-[36px] mb-[14px] mac:text-[30px] line-clamp-2 leading-10">{{ item.Title }}</h4>
          <hr class="max-w-[540px] bg-[#E4E4E4] mx-auto">
          <p
            class="text-[#6F6F6F] text-[28px] pt-[14px] text-right  mac:text-[20px] laptop:text-center line-clamp-[11] leading-7">
            {{ item.Description }}</p>
        </div>
      </div>
    </div>
  </router-link>
</body>
</template>

<style scoped>
.containers {
  max-width: 1450px;
  margin: 0 auto;
  padding: 0px 30px;
  border-radius: 10px;
}
</style>