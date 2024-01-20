<template>
    <h3 class="text-4xl text-center pt-4 pb-3 text-gray-500 font-medium mobile:text-4xl">Создание мероприятия</h3>
    <form @submit.prevent="submit" class="flex gap-x-2" action="#">
    <!-- FIXME: white text on white background on firefox -->
    <!-- TODO: add toast notification on errors --> 
    <input type="text" required v-model="title" placeholder="Название мероприятия" />
    <input type="text" required v-model="description" placeholder="Описание мероприятия" />
    <input type="file" ref="fileInputRef" @change="onFileChange" class="hidden">
    <button @click.prevent="openFileInput" class="text-black">+ Add a file</button>
    <button type="submit" class="rounded-md bg-indigo-600 bg- px-3.5 py-2.5 text-sm font-semibold text-white shadow-sm hover:bg-indigo-500 focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-indigo-600">Create</button>
    </form>
</template>

<script setup lang="ts">
import { ref } from 'vue';
import axios from 'axios';

const emit = defineEmits(['reloadSlider'])

const url: string = import.meta.env.VITE_ENDPOINT;
const file = ref<File | null>(null);
// FIXME: I think these 2 variables aren't used
let title = '';
let description = '';
const fileInputRef = ref<HTMLInputElement | null>(null)

const openFileInput = (): void => {
    // Trigger the hidden file input
    if (fileInputRef.value instanceof HTMLInputElement) {
        fileInputRef.value.click();
    }
}

const onFileChange = (event: Event): void => {
    const input = event.target as HTMLInputElement;
    if (input.files && input.files.length) {
        file.value = input.files[0];
    }
};

const submit = async (): Promise<void> => {
    try {
        const jsonData = {
            title: title,
            description: description
        };

        const formData = new FormData();

        // Append JSON data
        formData.append('jsonData', JSON.stringify(jsonData));

        // Append file
        if (file.value !== null) {
            formData.append('image', file.value);
        } else {
            // FIXME: this error is pretty big in console, prints out a lot of gibberish like "g, yh, Rt, qe"
            throw new Error('No image for event was choosen')
        }

        // Make POST request using Axios
        const response = await axios.post(`${url}/api/events`, formData, {
            headers: {
                'Content-Type': 'multipart/form-data'
            }
        })
        .finally(() => {
            emit('reloadSlider')
        })

        console.log('Response:', response.data);
    } catch (error) {
        console.error('Error', error);
    }
};
</script>

<style>

</style>