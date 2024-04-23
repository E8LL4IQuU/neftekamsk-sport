<script setup lang="ts">
import { ref, onMounted } from 'vue'
import axios from 'axios'
import { type Photo } from '@/types/apiTypes'

const url: string = import.meta.env.VITE_ENDPOINT
const photos = ref<Photo[]>([])

const fetchPhotos = async (): Promise<void> => {
  try {
    const response = await axios.get<Photo[]>(`${url}/api/photos`)
    photos.value = response.data
  } catch (error) {
    console.error('Error fetching photos: ', error)
  }
}

const getFileType = (url: string): string => {
  const extension = url.split('.').pop()?.toLowerCase()
  if (extension === 'mp4' || extension === 'mov' || extension === 'wmv' || extension === 'avi' || extension === 'ogg' || extension === 'webm' || extension === '3GP') {
    return 'video'
  }
  else return 'image'
}

onMounted(() => {
  fetchPhotos()
})

</script>

<template>
  <body class="container mx-auto pb-10 bg-white">
    <h1 class="text-3xl font-bold mb-4 text-black">Галерея</h1>
    <div class="grid lg:grid-cols-2 gap-4">
      <!-- TODO: enlarge photo on click -->
      <div v-for="photo in photos" :key="photo.ID">
        <img v-if="getFileType(photo.ImagePath) === 'image'" :src="`${url}/uploads/${photo.ImagePath}`"
          class="w-full h-72 object-cover rounded-lg">
        <video autoplay muted loop controls v-else :src="`${url}/uploads/${photo.ImagePath}`"
          class="w-full h-72 object-cover rounded-lg"></video>
      </div>
    </div>
  </body>
</template>

<style></style>