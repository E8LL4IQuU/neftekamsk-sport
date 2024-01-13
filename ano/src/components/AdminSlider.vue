<template>
  <div>
    <h3 class="text-4xl text-center pt-4 pb-3 text-gray-500 font-medium mobile:text-4xl">Редактирование мероприятий</h3>
    <div class="w-[100%] h-[100%]">
      <Swiper
          :spaceBetween="30"
          :centeredSlides="true"
          :autoplay="false"
          :navigation="true"
          :modules="modules"
      >
        <SwiperSlide v-for="(slide, index) in SliderData" :key="index">
          <div class="flex flex-col justify-center">
            <div class="w-[100%] text-white absolute text-center justify-center top-[36%]">
              <input v-model="slide.title"
              class="text-5xl mb-3 mobile:text-3xl block mx-auto text-center bg-black bg-opacity-60 border-0" />
              <input v-model="slide.description"
              class="text-2xl block mx-auto text-center mb-3 mobile:text-sm bg-black bg-opacity-60 border-0" />
              <button class="p-2 rounded-[4px] bg-slate-700 hover:bg-gray-900 duration-300 mr-3">Применить</button>
              <!-- TODO: Implement event removal -->
              <button class="p-2 rounded-[4px] bg-red-500 hover:bg-red-900 duration-300" @click="deleteEvent(slide.ID)">Удалить</button>
            </div>
            <img class="w-auto h-[100vh]  mobile:h-72" :src="`/uploads/${slide.img}`" alt="slider image"/>
            <!-- TODO: background image selector -->
            <input type="file" />
          </div>
        </SwiperSlide>
      </Swiper>
    </div>
  </div>
</template>
<script setup lang="ts">
import axios from 'axios';
import {Swiper, SwiperSlide} from "swiper/vue";
import { Autoplay, Pagination, Navigation } from 'swiper/modules';
import 'swiper/swiper-bundle.css';
import 'swiper/scss/navigation';
import 'swiper/scss/pagination';

interface ISlider {
  ID: number;
  title: string;
  description: string;
  links: string;
  img: string
}

const props = defineProps<{
  SliderData: ISlider[]
}>()
const modules = [
  Autoplay,
  Pagination,
  Navigation
]

const deleteEvent = async (id: number) => {
  try {
    const response = await axios.delete(`/api/events/${id}`)
    .then((response) =>
    console.log(response));
  } catch (error) {
    console.error('Error deleting record:', error)
  }
}
</script>

<style scoped lang="sass">

</style>