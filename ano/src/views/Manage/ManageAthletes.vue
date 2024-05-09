<script setup lang="ts">
import { onMounted, ref } from 'vue'
import axios from 'axios'
import { RouterLink } from 'vue-router'
import { type Athlete } from '@/types/apiTypes'
import TrashButton from '@/components/Manage/TrashButton.vue'

const url: string = import.meta.env.VITE_ENDPOINT
const athletes = ref<Athlete[]>([])

const fetchAthletes = async (): Promise<void> => {
    try {
        const response = await axios.get<Athlete[]>(`${url}/api/athletes`)
        athletes.value = response.data
    } catch (error) {
        // FIXME: prints out gibberish instead of server's response
        console.error('Error fetching athletes:', error)
    }
}

onMounted(() => {
    fetchAthletes()
})
</script>

<template>
    <div class="lg:ps-52">
        <div class="p-4 lg:py-5 lg:px-12">
            <div class="flex justify-between">
                <h1 class="text-3xl text-black font-bold mb-3">Люди спорта</h1>
                <router-link :to="'/manage/athletes/create'">
                    <button
                        class="bg-black rounded-xl text-xl lg:text-base lg:rounded text-white py-2 px-3 ml-auto">Создать</button>
                </router-link>
            </div>
            <div class="athletes-list">
                <router-link v-for="athlete in athletes" :key="athlete.ID"
                    :to="{ name: 'athletesEdit', params: { id: athlete.ID } }">
                    <!-- FIXME: enforce fixed image size here and in AthletesView -->
                    <div class="mb-4 p-4 bg-white shadow-md rounded-md cursor-pointer flex gap-x-2">
                        <img :src="`${url}/uploads/${athlete.ImagePath}`" class="h-32 w-24 object-cover" alt="">
                        <div class="basis-full">
                            <h2 class="text-2xl font-bold text-gray-700 line-clamp-1">{{ athlete.Title }}</h2>
                            <div class="lg:flex lg:justify-between lg:items-center">
                                <p class="text-gray-700 line-clamp-2">{{ athlete.Description }}</p>
                                <div class="min-w-fit">
                                    <TrashButton @reloadItems="fetchAthletes" :id="athlete.ID" type="athletes">
                                    </TrashButton>
                                </div>
                            </div>
                        </div>
                    </div>
                </router-link>
            </div>
        </div>
    </div>
</template>

<style scoped>
.athletes-list {
    max-width: 900px;
    margin: 0 auto;
}
</style>