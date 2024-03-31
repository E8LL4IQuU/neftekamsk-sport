<template>
  <div class="wrapper">
    <Hero />
    <h3
      class="text-6xl text-center pt-16 pb-6 text-gray-500 font-medium mobile:text-4xl"
    >
      Мероприятия
    </h3>
    <Slider :SliderData="IRLEvents" />
    <h3 class="text-6xl pt-16 text-center text-black font-normal">Новости</h3>
    <div class="flex flex-wrap justify-center">
      <NewsComponent :NewsData="News" />
    </div>
    <h3 class="text-6xl pt-10 pb-10 text-center text-black font-normal">
      Галерея
    </h3>
    <div>
      <Gallery :galleryData="galleryData" />
    </div>
    <h3 class="text-6xl pt-16 pb-10 text-center text-black font-normal">
      Руководители
    </h3>
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
import axios from "axios";
import { type IRLEvent, type News } from "@/types/apiTypes";

import HeroPhoto1 from "@/assets/heroPhoto1.jpeg"
import HeroPhoto2 from "@/assets/heroPhoto2.jpeg"
import HeroPhoto3 from "@/assets/heroPhoto3.jpeg"
import HeroVideo from "@/assets/heroVideo.webm";
import StaffPhoto1 from "@/assets/staff1.jpeg";
import StaffPhoto2 from "@/assets/staff2.jpeg";

const url: string = import.meta.env.VITE_ENDPOINT;
const IRLEvents = ref<IRLEvent[]>([]);
const News = ref<News[]>([]);

const fetchIRLEvents = async (): Promise<void> => {
  try {
    const response = await axios.get<IRLEvent[]>(`${url}/api/events`);
    IRLEvents.value = response.data;
  } catch (error) {
    console.log("Error fetching IRLEvents: ", error);
  }
};

const fetchNews = async (): Promise<void> => {
  try {
    const response = await axios.get<News[]>(`${url}/api/news?limit=3`);
    News.value = response.data;
  } catch (error) {
    console.error("Error fetching news:", error);
  }
};

onMounted(async () => {
  fetchIRLEvents();
  fetchNews();
});

const galleryData = {
  photos: [
    HeroPhoto1,
    HeroPhoto3,
    HeroPhoto2,
  ],
  video: HeroVideo
};

const PeopleData = [
  {
    img: StaffPhoto1,
    name: "Андрей Филимонов",
    activity: "Директор",
    description: "",
  },
  {
    img: StaffPhoto2,
    name: "Елена Приходько",
    activity: "Учредитель",
    description: "",
  },
];
</script>

<style scoped>
.wrapper {
  max-width: 1380px;
  padding: 18px 0px;
  margin: 0 auto;
}
</style>
