<template>
  <div class="wrapper">
    <Hero />
    <h3 class="text-6xl text-center pt-16 pb-6 text-gray-500 font-medium mobile:text-4xl">Мероприятия</h3>
    <Slider :SliderData="IRLEvents" />
    <h3 class="text-6xl pt-16 text-center text-black font-normal">Новости</h3>
    <div class="flex flex-wrap justify-center">
      <NewsComponent :NewsData="News" />
    </div>
    <h3 class="text-6xl pt-10 pb-10 text-center text-black font-normal">Галерея</h3>
    <div>
      <Gallery :GalleryData="GalleryData" />
    </div>
    <h3 class="text-6xl pt-16 pb-10 text-center text-black font-normal">Истории людей</h3>
    <div>
      <StoriesPeople :PeopleData="PeopleData" />
    </div>
    <div>
      <Connect />
    </div>
  </div>
</template>

<script setup lang="ts">
import Hero from "@/components/Hero.vue";
import Slider from "@/components/Slider.vue";
import NewsComponent from "@/components/News.vue";
import Gallery from "@/components/Gallery.vue";
import StoriesPeople from "@/components/StoriesPeople.vue";
import Connect from "@/components/Connect.vue";
import { onMounted, ref } from "vue";
import axios from 'axios'
import { type IRLEvent, type News } from '@/types/apiTypes'

const url: string = import.meta.env.VITE_ENDPOINT;
const IRLEvents = ref<IRLEvent[]>([]);
const News = ref<News[]>([])

const fetchIRLEvents = async (): Promise<void> => {
  try {
    const response = await axios.get<IRLEvent[]>(`${url}/api/events`)
    IRLEvents.value = response.data
  } catch (error) {
    console.log('Error fetching IRLEvents: ', error)
  }
}

const fetchNews = async (): Promise<void> => {
  try {
    const response = await axios.get<News[]>(`${url}/api/news?limit=3`)
    News.value = response.data
  } catch (error) {
    console.error('Error fetching news:', error)
  }
}

onMounted(async () => {
  fetchIRLEvents()
  fetchNews()
})

// mock data
const GalleryData = {
  "linkImgOne": "https://thumb.tildacdn.com/tild6538-3765-4639-b764-656637656238/-/format/webp/_1.JPG",
  "linkImgTwo": "https://thumb.tildacdn.com/tild6538-3765-4639-b764-656637656238/-/format/webp/_1.JPG",
  "linkImgThree": "https://thumb.tildacdn.com/tild6538-3765-4639-b764-656637656238/-/format/webp/_1.JPG",
  "videos": "https://res.cloudinary.com/demo/video/upload/q_auto,f_auto/dog.mp4"
}
const PeopleData = [
  {
    "img": "https://static.tildacdn.com/tild3235-3065-4462-a662-666363353834/i.jpg",
    "name": "Dmitry Shishkin",
    "activity": "Предприниматель",
    "description": "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Phasellus nisi metus, cursus et diam in, feugiat porttitor sapien. Mauris vel fringilla quam",
  },
  {
    "img": "https://static.tildacdn.com/tild3235-3065-4462-a662-666363353834/i.jpg",
    "name": "Dmitry Shishkin",
    "activity": "Предприниматель",
    "description": "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Phasellus nisi metus, cursus et diam in, feugiat porttitor sapien. Mauris vel fringilla quam",
  },
  {
    "img": "https://static.tildacdn.com/tild3235-3065-4462-a662-666363353834/i.jpg",
    "name": "Dmitry Shishkin",
    "activity": "Предприниматель",
    "description": "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Phasellus nisi metus, cursus et diam in, feugiat porttitor sapien. Mauris vel fringilla quam",
  },
]
</script>
<style scoped>
.wrapper {
  max-width: 1380px;
  padding: 18px 0px;
  margin: 0 auto;
}
</style>