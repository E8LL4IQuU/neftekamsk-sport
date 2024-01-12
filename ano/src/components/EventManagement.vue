<template>
    <form @submit.prevent="submit" action="#">
    <!-- FIXME: white text on white background on firefox -->
    <input type="text" v-model="title" placeholder="Название мероприятия" />
    <input type="text" v-model="description" placeholder="Описание мероприятия" />
    <input type="file" @change="onFileChange">
    <button type="submit">Create</button>
    </form>
</template>

<script setup lang="ts">
import { ref } from 'vue';
import axios from 'axios';

const file = ref<File | null>(null);
let title = '';
let description = '';

const onFileChange = (event: Event) => {
    const input = event.target as HTMLInputElement;
    if (input.files && input.files.length) {
        file.value = input.files[0];
    }
};

const submit = async () => {
    // FIXME: 413 on big file name I'd guess
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
            // FIXME: this error is pretty big in console like (g, yh, Rt, qe)
            throw new Error('No image for event was choosen')
        }

        // Make POST request using Axios
        const response = await axios.post('/api/events', formData, {
            headers: {
                'Content-Type': 'multipart/form-data'
            }
        });

        console.log(formData)
        console.log('Response:', response.data);
    } catch (error) {
        console.error('Error', error);
    }
};
</script>

<style>

</style>