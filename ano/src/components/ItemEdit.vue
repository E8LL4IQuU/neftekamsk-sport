<script setup lang="ts">
import { ref, onMounted } from 'vue';
import axios from 'axios';
import { RouterLink, useRoute, useRouter } from 'vue-router';
import { ChevronLeftIcon, PlusIcon } from '@heroicons/vue/24/solid';
import { type News } from '@/types/apiTypes';
import VueDatePicker from '@vuepic/vue-datepicker';
import '@vuepic/vue-datepicker/dist/main.css';

const ItemForm = defineProps(['requestType', 'url']);
const url: string = import.meta.env.VITE_ENDPOINT;
const router = useRouter();
const route = useRoute();
const file = ref<File | null>(null);
const fileInputRef = ref<HTMLInputElement | null>(null);
const title = ref<string>('');
const description = ref<string>('');
const itemId: number = Number(route.params.id)
const selectedDate = ref<Date>();

const openFileInput = (): void => {
  if (fileInputRef.value instanceof HTMLInputElement) {
    fileInputRef.value.click();
  }
};

const onFileChange = (event: Event): void => {
  const input = event.target as HTMLInputElement;
  if (input.files && input.files.length) {
    file.value = input.files[0];
  }
};

const submit = async (): Promise<void> => {
  try {
    const formData = new FormData();

    formData.append('title', title.value);
    formData.append('description', description.value);

    if (file.value !== null) {
      formData.append('image', file.value);
    }

    const config = {
      headers: {
        'Content-Type': 'multipart/form-data',
      },
    };

    const requestConfig =
      ItemForm.requestType === 'post'
        ? { method: 'post', url: `${url}/api/${ItemForm.url}`, data: formData, config }
        : { method: 'put', url: `${url}/api/${ItemForm.url}/${itemId}`, data: formData, config };

    const response = await axios(requestConfig).then(() => {
      router.push(`/manage/${ItemForm.url}`);
    });

    console.log(response);
  } catch (error) {
    console.error('Error updating Item:', error);
  }
};

const getFields = async (): Promise<void> => {
  if (isNaN(itemId)) {
    return
  }

  try {
    const response = await axios.get<News>(`${url}/api/${ItemForm.url}/${itemId}`)
    title.value = response.data.Title
    description.value = response.data.Description
  } catch (error) {
    if (axios.isAxiosError(error) && error.response?.status === 401) {
      router.push("/login")
    } else {
      console.error('Error fetching news: ', error)
    }
  }

}

onMounted(() => {
  getFields()
}
)
</script>

<template>
  <!-- TODO: add image upload progress -->

  <body class="bg-white">
    <header class="pe-auto">
      <div class="flex items-center pe-auto p-3 h-100">
        <router-link :to="`/manage/${ItemForm.url}`" class="flex p-4 gap-x-3 text-black items-center">
          <ChevronLeftIcon class="h-4 pt-1"></ChevronLeftIcon>
          <p class="font-bold text-sm text-gray-700">Мероприятия</p>
        </router-link>
        <p class="text-gray-400 ms-3">Новое</p>
      </div>
    </header>
    <form @submit.prevent="submit" @keydown.enter.prevent="submit" class="flex flex-col ps-28 pt-8">
      <div class="flex items-center gap-x-5">
        <input required type="file" ref="fileInputRef" @change="onFileChange" class="hidden">
        <div class="inline-flex group">
          <PlusIcon class="h-5 pt-1 text-gray-500 group-hover:text-gray-600"></PlusIcon>
          <button @click.prevent="openFileInput" class="text-gray-400 ms-1 group-hover:text-gray-600">Добавить
            изображение</button>
        </div>
        <div class=" w-64">
          <VueDatePicker v-model="selectedDate" placeholder="Выберите дату" text-input></VueDatePicker>
        </div>
      </div>
      <input required v-model="title"
        class="text-6xl placeholder-gray-300 font-bold border-none tracking-tight -ms-3 focus:ring-0"
        placeholder="Название мероприятия" />
      <textarea required v-model="description" class="text-xl placeholder-gray-400 border-none -ms-2 focus:ring-0"
        placeholder="Начните писать описание..." rows="15"></textarea>
    </form>
  </body>
</template>

<style>

.dp__input {
  border: 0px;
}
</style>
