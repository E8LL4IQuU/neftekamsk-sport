<template>
  <div>
    <h3 class="text-4xl text-center pt-4 pb-3 text-gray-500 font-medium mobile:text-4xl">Редактирование мероприятий</h3>
    <div class="w-[100%] h-[100%]">
      <Swiper :spaceBetween="30" :centeredSlides="true" :autoplay="false" :navigation="true" :modules="modules">
        <SwiperSlide v-for="(slide, index) in SliderData" :key="index">
          <div class="flex flex-col justify-center">
            <div class="w-[100%] text-white absolute text-center justify-center top-[36%]">
              <form @submit.prevent="updateEvent(slide)">
                <input v-model="slide.title"
                  class="text-5xl mb-3 mobile:text-3xl block mx-auto text-center bg-black bg-opacity-60 border-0" />
                <input v-model="slide.description"
                  class="text-2xl block mx-auto text-center mb-3 mobile:text-sm bg-black bg-opacity-60 border-0" />
                <button type="submit"
                  class="p-2 rounded-[4px] bg-slate-700 hover:bg-gray-900 duration-300 mr-3">Применить</button>
                <button class="p-2 rounded-[4px] bg-red-500 hover:bg-red-900 duration-300 mr-3"
                  @click="deleteEvent(slide.id)">Удалить</button>
                <!-- TODO: visualize changing background image -->
                <input type="file" @change="onFileChange" accept="image/*" />
              </form>
            </div>
            <img class="w-auto h-72 lg:h-[75vh] object-cover" :src="`${url}/uploads/${slide.img}`" alt="slider image" />
          </div>
        </SwiperSlide>
      </Swiper>
    </div>
  </div>
</template>
<script setup lang="ts">
import { ref } from 'vue';
import axios from 'axios';
import { Swiper, SwiperSlide } from "swiper/vue";
import { Autoplay, Pagination, Navigation } from 'swiper/modules';
import { type IRLEvent } from "@/types/apiTypes"
import 'swiper/swiper-bundle.css';
import 'swiper/scss/navigation';
import 'swiper/scss/pagination';

const url: string = import.meta.env.VITE_ENDPOINT;
const emit = defineEmits(['reloadSlider'])
const file = ref<File | null>(null)
const props = defineProps<{
  SliderData: IRLEvent[]
}>()
const modules = [
  Autoplay,
  Pagination,
  Navigation
]

const onFileChange = (event: Event) => {
  const input = event.target as HTMLInputElement;

  if (input.files && input.files.length) {
    file.value = input.files[0];
  }
}

const updateEvent = async (slide: IRLEvent) => {
  try {
    const formData = new FormData();

    formData.append('title', slide.title)
    formData.append('description', slide.description)

    if (file.value !== null) {
      formData.append('image', file.value)
    }

    const response = await axios.put(`${url}/api/events/${slide.id}`, formData, {
      headers: {
        'Content-Type': 'multipart/form-data'
      }
    });
    console.log(response)
  } catch (error) {
    console.error('Error updating event:', error)
  }
}

const deleteEvent = async (id: number) => {
  try {
    const response = await axios.delete(`${url}/api/events/${id}`)
      .then((response) =>
        console.log(response))
      .finally(() => {
        emit('reloadSlider')
      })
  } catch (error) {
    console.error('Error deleting record:', error)
  }
}
</script>

<style scoped lang="sass">

</style>