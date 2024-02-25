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
    }   catch (error) {
        console.error('Error fetching photos: ', error)
    }
}

onMounted(() => {
  fetchPhotos()
})

</script>

<template>
    <div class="container mx-auto pb-10">
      <h1 class="text-3xl font-bold mb-4">Галлерея</h1>
      <div class="grid lg:grid-cols-2 gap-4">
        <!-- TODO: enlarge photo on click -->
        <div v-for="photo in photos" :key="photo.ID">
          <img :src="`${url}/uploads/${photo.ImagePath}`" class="w-full h-72 object-cover rounded-lg">
        </div>
      </div>
    </div>
  </template>

<style>

</style>