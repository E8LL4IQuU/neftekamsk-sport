<script setup lang="ts">
import { RouterLink, useRouter } from "vue-router";
import { ChevronLeftIcon, PlusIcon } from "@heroicons/vue/24/solid";

import { ref } from 'vue';
import axios from 'axios';

const url: string = import.meta.env.VITE_ENDPOINT;
const router = useRouter()
const file = ref<File | null>(null);
const fileInputRef = ref<HTMLInputElement | null>(null)
const title = ref()
const description = ref()

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
}

// TODO: add toasts or something on errors
const submit = async (): Promise<void> => {
  try {
    const formData = new FormData();

    formData.append('title', title.value)
    formData.append('description', description.value);

    if (file.value !== null) {
      formData.append('image', file.value)
    }

    const response = await axios.post(`${url}/api/events/`, formData, {
      headers: {
        'Content-Type': 'multipart/form-data'
      }
    })
      .then(() => {
        router.push("/manage/events")
      })
    console.log(response)
  } catch (error) {
    console.error('Error updating event:', error)
  }
}

</script>

<template>
  <body class="bg-white">
    <header class="pe-auto">
      <div class="flex items-center pe-auto p-3 h-100">
        <!-- TODO: animation on hover -->
        <router-link :to="'/manage/events'" class="flex p-4 gap-x-3 text-black items-center">
          <ChevronLeftIcon class="h-4 pt-1"></ChevronLeftIcon>
          <p class="font-bold text-sm text-gray-700">Мероприятия</p>
        </router-link>
        <p class="text-gray-400 ms-3">Новое</p>
        <!-- TODO: implement event creation -->
      </div>
    </header>
    <form @submit.prevent="submit" @keydown.enter.prevent="submit" class="flex flex-col ps-28 pt-8">
      <!-- TODO: check if fields are filled before submiting -->
      <input required type="file" ref="fileInputRef" @change="onFileChange" class="hidden">
      <div class="inline-flex group">
        <PlusIcon class="h-5 pt-1 text-gray-500 group-hover:text-gray-600"></PlusIcon>
        <button @click.prevent="openFileInput" class="text-gray-400 ms-1 group-hover:text-gray-600">Добавить
          изображение</button>
      </div>
      <!-- FIXME: white text on white background in firefox dark mode -->
      <input required v-model="title" class="text-6xl placeholder-gray-300 font-bold border-none tracking-tight -ms-3 focus:ring-0"
        placeholder="Название мероприятия" />
      <textarea required v-model="description" class="text-xl placeholder-gray-400 border-none -ms-2 focus:ring-0"
        placeholder="Начните писать описание..." rows="15"></textarea>
    </form>

  </body>
</template>

<style scoped></style>
