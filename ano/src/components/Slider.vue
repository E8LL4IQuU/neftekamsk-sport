<template>
  <div class="w-[100%] h-[100%]">
    <Swiper :spaceBetween="30" :centeredSlides="true" :autoplay="{
      delay: 2500,
      disableOnInteraction: true,
    }" :navigation="true" :modules="modules">
      <SwiperSlide v-for="(slide, index) in SliderData" :key="index">
        <div class="flex flex-col justify-center">
          <img class="w-auto h-72 lg:h-screen object-cover lg:max-h-[700px] brightness-[60%]"
            :src="`${url}/uploads/${slide.ImagePath}`" alt="slider image" />
          <div class="w-full text-white absolute text-center justify-center">
            <h2 class="text-white text-2xl lg:text-7xl font-bold pb-4 lg:pb-10 line-clamp-1">{{ slide.Title }}</h2>
            <!-- <h3 class="text-white font-normal pb-3 text-[20px]">{{ formatTimestamp(slide.Date) }}</h3> -->
            <router-link :to="`/events/${slide.ID}`" class="">
              <a
                class="py-3 px-6 font-bold rounded-lg bg-gray-700 hover:bg-gray-600 hover:scale-110 duration-200 ease-out">
                Подробнее
              </a>
            </router-link>
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

// const formatTimestamp = (timestamp: string): string => {
//   const date = new Date(timestamp);

//   const options: Intl.DateTimeFormatOptions = {
//     year: 'numeric',
//     month: '2-digit',
//     day: 'numeric',
//     hour: 'numeric',
//     minute: 'numeric'

//   };

//   return date.toLocaleDateString("ru-RU", options);
// }
</script>

<style scoped lang="sass">

</style>