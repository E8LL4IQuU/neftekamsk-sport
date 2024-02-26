<template>
  <div class="w-[100%] h-[100%]">
    <Swiper :spaceBetween="30" :centeredSlides="true" :autoplay="{
      delay: 2500,
      disableOnInteraction: true,
    }" :navigation="true" :modules="modules">
      <SwiperSlide v-for="(slide, index) in SliderData" :key="index">
        <div class="flex flex-col justify-center">
          <img class="w-auto h-72 lg:h-screen object-cover lg:max-h-[700px] brightness-75"
            :src="`${url}/uploads/${slide.ImagePath}`" alt="slider image" />
          <div class="w-full text-white absolute text-center justify-center">
            <h2 class="text-white text-5xl pb-3 mobile:text-3xl line-clamp-1">{{ slide.Title }}</h2>
            <h3 class="text-white font-normal pb-3 text-[20px]">{{ formatTimestamp(slide.Date) }}</h3>
            <router-link disabled class="p-2 rounded-[4px] bg-gray-400 hover:bg-gray-600 duration-300"
              :to="`/events/${slide.ID}`">Подробнее</router-link>
          </div>
        </div>
      </SwiperSlide>
    </Swiper>
  </div>
</template>
<script setup lang="ts">
import { Swiper, SwiperSlide } from "swiper/vue";
import { Autoplay, Pagination, Navigation } from 'swiper/modules';
import 'swiper/swiper-bundle.css';
import 'swiper/scss/navigation';
import 'swiper/scss/pagination';
import type { IRLEvent } from "@/types/apiTypes";

const url: string = import.meta.env.VITE_ENDPOINT
const props = defineProps<{
  SliderData: IRLEvent[]
}>()
const modules = [
  Autoplay,
  Pagination,
  Navigation
]

const formatTimestamp = (timestamp: string): string => {
  const date = new Date(timestamp);

  const options: Intl.DateTimeFormatOptions = {
    year: 'numeric',
    month: '2-digit',
    day: 'numeric',
    hour: 'numeric',
    minute: 'numeric'

  };

  return date.toLocaleDateString("ru-RU", options);
}
</script>

<style scoped lang="sass">

</style>