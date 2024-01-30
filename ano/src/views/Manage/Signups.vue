<script setup lang="ts">
import { onMounted, ref } from 'vue'
import axios from 'axios'
import { RouterLink } from 'vue-router'
import { type Signup } from '@/types/apiTypes'
import TrashButton from '@/components/TrashButton.vue'

const url: string = import.meta.env.VITE_ENDPOINT
const signups = ref<Signup[]>([])

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
        minute: 'numeric',
        second: 'numeric',
    };
    
    return localDate.toLocaleString('en-US', options);
};


const fetchSignups = async (): Promise<void> => {
    try {
        const response = await axios.get<Signup[]>(`${url}/api/signups`)
        signups.value = response.data
    } catch (error) {
        console.error('Error fetching signups:', error)
    }
}

onMounted(() => {
    fetchSignups()
})
</script>

<template>
    <div class="py-8 px-12">
        <div class="flex justify-between">
            <h1 class="text-3xl text-black font-bold">Новости</h1>
            <router-link :to="'/manage/signups/create'">
                <button class="bg-black rounded text-white py-2 px-3 ml-auto">Создать</button>
            </router-link>
        </div>
        <div class="signups-list">
            <router-link v-for="signupItem in signups" :key="signupItem.ID" :to="{ name: 'signupsEdit', params: { id: signupItem.ID } }">
                <div class="mb-4 p-4 bg-white shadow-md rounded-md cursor-pointer">
                    <div class="flex justify-between items-center mb-4">
                        <h2 class="text-2xl font-bold text-gray-700">{{ signupItem.FirstName }}  {{ signupItem.LastName }}</h2>
                        <!-- FIXME: timestamp is incorrect -->
                        <span class="text-gray-500">{{ formatTimestamp(signupItem.CreatedAt) }}</span>
                    </div>
                    <div class="flex justify-between items-center">
                        <p class="text-gray-700">{{ signupItem.PhoneNumber }}</p>
                    <TrashButton @reloadItems="fetchSignups" :id="signupItem.ID" type="signups"></TrashButton>
                    </div>
                </div>
            </router-link>
        </div>
    </div>
</template>

<style scoped>
.signups-list {
    max-width: 800px;
    margin: 0 auto;
}
</style>