<script setup lang="ts">
import { onMounted, ref } from 'vue'
import axios from 'axios'
import { RouterLink } from 'vue-router'
import { type Signup } from '@/types/apiTypes'
import TrashButton from '@/components/Manage/TrashButton.vue'

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
            <h1 class="text-3xl text-black font-bold mb-3">Записи</h1>
        </div>
        <div class="signups-list">
            <div v-for="signupItem in signups" key="signupItem.ID" class="mb-4 p-4 bg-white shadow-md rounded-md">
                <div class="lg:flex lg:justify-between lg:items-center mb-4">
                    <h2 class="text-2xl font-bold text-gray-700 line-clamp-1">{{ signupItem.FirstName }} {{ signupItem.LastName }}</h2>
                    <!-- FIXME: timestamp is incorrect -->
                    <span class="text-gray-500">{{ formatTimestamp(signupItem.CreatedAt) }}</span>
                </div>
                <div class="lg:flex lg:justify-between lg:items-center">
                    <div class="lg:flex lg:gap-x-5"> 
                        <span class="text-gray-700 line-clamp-1">{{ signupItem.PhoneNumber }}</span>
                        <span class="text-gray-700 line-clamp-1">{{ signupItem.Email }}</span>
                    </div>
                    <!-- FIXME: 405 Method not allowed -->
                    <TrashButton @reloadItems="fetchSignups" :id="signupItem.ID" type="signups"></TrashButton>
                </div>
            </div>
        </div>
    </div>
</template>

<style scoped>
.signups-list {
    max-width: 800px;
    margin: 0 auto;
}
</style>