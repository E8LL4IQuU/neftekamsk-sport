<template>
  <div>
    <h3 class="text-6xl text-center pt-4 pb-3 text-gray-500 font-medium mobile:text-4xl">Мероприятия</h3>
    <div class="w-[100%] h-[100%]">
      <Swiper
          :spaceBetween="30"
          :centeredSlides="true"
          :autoplay="{
      delay: 2500,
      disableOnInteraction: false,
    }" :navigation="true"
          :modules="modules"
      >
        <SwiperSlide v-for="(slide, index) in SliderData" :key="index">
          <div class="flex flex-col justify-center">
            <div class="w-[100%] text-white absolute text-center justify-center top-[36%]">
              <h2 class="text-5xl pb-3 mobile:text-3xl">{{ slide.title }}</h2>
              <p class="text-2xl pb-3 mobile:text-sm">{{ slide.description }}</p>
              <!-- TODO: imlement, make router-link -->
              <a disabled class="p-2 rounded-[4px] bg-gray-400 hover:bg-gray-600 duration-300" :href="'#'">Подробнее</a>
            </div>
            <img class="w-auto h-72 lg:h-screen object-cover" :src="`${url}/uploads/${slide.img}`" alt="slider image"/>
          </div>
        </SwiperSlide>
      </Swiper>
    </div>
  </div>
</template>
<script setup lang="ts">
import {Swiper, SwiperSlide} from "swiper/vue";
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
</script>

<style scoped lang="sass">

</style>