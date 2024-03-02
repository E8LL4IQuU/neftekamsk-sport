<script setup lang="ts">
import { ref, onMounted } from 'vue'
import axios from 'axios'
import { type Photo } from '@/types/apiTypes'
import TrashButton from '@/components/Manage/TrashButton.vue'


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
    <div class="lg:ps-52">
        <div class="p-4 lg:py-5 lg:px-12">
            <div class="flex justify-between">
                <h1 class="text-3xl text-black font-bold mb-3">Галлерея</h1>
                <router-link :to="'/manage/photos/create'">
                    <button
                        class="bg-black rounded-xl text-xl lg:text-base lg:rounded text-white py-2 px-3 ml-auto">Создать</button>
                </router-link>
            </div>
            <div>
                <div class="grid grid-cols-2 gap-4">
                    <div v-for="photo in photos" :key="photo.ID" class="relative">
                        <img v-if="getFileType(photo.ImagePath) === 'image'" :src="`${url}/uploads/${photo.ImagePath}`"
                            class="w-full h-72 object-cover rounded-lg">
                        <video autoplay muted loop controls v-else :src="`${url}/uploads/${photo.ImagePath}`"
                            class="w-full h-72 object-cover rounded-lg"></video>
                        <div class="min-w-fit absolute right-2 bottom-2">
                            <TrashButton @reloadItems="fetchPhotos" :id="photo.ID" type="photos"></TrashButton>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    </div>
</template>

<style></style>