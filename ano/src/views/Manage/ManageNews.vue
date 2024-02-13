<script setup lang="ts">
import { onMounted, ref } from 'vue'
import axios from 'axios'
import { RouterLink } from 'vue-router'
import { type News } from '@/types/apiTypes'
import TrashButton from '@/components/Manage/TrashButton.vue'

const url: string = import.meta.env.VITE_ENDPOINT
const news = ref<News[]>([])

const formatTimestamp = (timestamp: number): string => {
    const date = new Date(timestamp * 1000);

    // Get the time zone offset in minutes
    const timeZoneOffset = date.getTimezoneOffset();
    
    // Use the offset to create a new date with the local time zone
    const localDate = new Date(date.getTime() - timeZoneOffset * 60 * 1000);

    const options: Intl.DateTimeFormatOptions = {
        year: 'numeric',
        month: 'long',
        day: 'numeric',
        hour: 'numeric',
        minute: 'numeric'
    };
    
    return localDate.toLocaleString('ru-RU', options);
};


const fetchNews = async (): Promise<void> => {
    try {
        const response = await axios.get<News[]>(`${url}/api/news`)
        news.value = response.data
    } catch (error) {
        console.error('Error fetching news:', error)
    }
}

onMounted(() => {
    fetchNews()
})
</script>

<template>
    <div class="py-8 px-12">
        <div class="flex justify-between">
            <h1 class="text-3xl text-black font-bold">Новости</h1>
            <router-link :to="'/manage/news/create'">
                <button class="bg-black rounded text-white py-2 px-3 ml-auto">Создать</button>
            </router-link>
        </div>
        <div class="news-list">
            <router-link v-for="newsItem in news" :key="newsItem.ID" :to="{ name: 'newsEdit', params: { id: newsItem.ID } }">
                <div class="mb-4 p-4 bg-white shadow-md rounded-md cursor-pointer">
                    <div class="flex justify-between items-center mb-4">
                        <h2 class="text-2xl font-bold text-gray-700 line-clamp-1">{{ newsItem.Title }}</h2>
                        <span class="text-gray-500 min-w-fit">{{ formatTimestamp(newsItem.CreatedAt) }}</span>
                    </div>
                    <div class="flex justify-between items-center">
                        <p class="text-gray-700 line-clamp-2">{{ newsItem.Description }}</p>
                        <!-- TODO: deleting doesn't update the page -->
                        <div class="min-w-fit">
                            <TrashButton @reloadNews="fetchNews" :id="newsItem.ID" type="news"></TrashButton>
                        </div>
                    </div>
                </div>
            </router-link>
        </div>
    </div>
</template>

<style scoped>
.news-list {
    max-width: 800px;
    margin: 0 auto;
}
</style>