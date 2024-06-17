<script setup lang="ts">
import { ref, onMounted, computed } from 'vue';
import axios from 'axios';
import { useRoute, useRouter } from 'vue-router';
import { ChevronLeftIcon, PlusIcon } from '@heroicons/vue/24/solid';
import { type IRLEvent, type News } from '@/types/apiTypes';
import '@vuepic/vue-datepicker/dist/main.css';

// Props
const ItemForm = defineProps<{
  requestType: string;
  url: string;
  breadcrumb: string;
  title?: string;
}>();

// Constants
const baseUrl: string = import.meta.env.VITE_ENDPOINT;

// Router & Route
const router = useRouter();
const route = useRoute();
const itemId: number = Number(route.params.id);

// Refs
const file = ref<File | null>(null);
const fileInputRef = ref<HTMLInputElement | null>(null);
const filename = ref<string>('+ Добавить изображение');
const title = ref<string>('');
const description = ref<string>('');
const date = ref<Date | null>(null);

// Computed
const mode = computed(() => {
  switch (ItemForm.requestType) {
    case 'post':
      return 'новая запись';
    case 'put':
      return 'редактирование';
    default:
      return 'custom action';
  }
});

// Methods
const openFileInput = (): void => {
  fileInputRef.value?.click();
};

const onFileChange = (event: Event): void => {
  const input = event.target as HTMLInputElement;
  if (input.files && input.files.length) {
    file.value = input.files[0];
    filename.value = file.value.name;
  }
};

const submit = async (): Promise<void> => {
  try {
    const formData = new FormData();
    formData.append('title', title.value);
    formData.append('description', description.value);
    if (date.value !== null) {
      formData.append('date', date.value.toISOString());
    }
    if (file.value !== null) {
      formData.append('image', file.value);
    }

    const config = {
      headers: {
        'Content-Type': 'multipart/form-data',
      },
    };

    const requestConfig = ItemForm.requestType === 'post'
      ? { method: 'post', url: `${baseUrl}/api/${ItemForm.url}`, data: formData, config }
      : { method: 'put', url: `${baseUrl}/api/${ItemForm.url}/${itemId}`, data: formData, config };

    await axios(requestConfig);
    router.push(`/manage/${ItemForm.url}`);
  } catch (error) {
    console.error('Error updating Item:', error);
    // TODO: Show user-friendly error message
  }
};

const getFields = async (): Promise<void> => {
  if (isNaN(itemId)) return;

  try {
    const response = await axios.get<News | IRLEvent>(`${baseUrl}/api/${ItemForm.url}/${itemId}`);
    title.value = response.data.Title;
    description.value = response.data.Description;
    if ('Date' in response.data) {
      date.value = new Date(response.data.Date);
    }
  } catch (error) {
    if (axios.isAxiosError(error) && error.response?.status === 401) {
      router.push('/login');
    } else {
      console.error('Error fetching item:', error);
      // TODO: Show user-friendly error message
    }
  }
};

// Lifecycle Hooks
onMounted(() => {
  if (ItemForm.requestType === 'put') {
    getFields();
  }
});
</script>

<template>

  <body class="bg-white">
    <header class="pe-auto">
      <div class="flex items-center pe-auto md:p-3 h-100">
        <router-link :to="`/manage/${ItemForm.url}`" class="flex p-4 gap-x-3 text-black items-center">
          <ChevronLeftIcon class="h-4 pt-1" />
          <p class="font-bold text-sm text-gray-700 capitalize">{{ ItemForm.breadcrumb }}</p>
        </router-link>
        <p class="text-gray-400 ms-3">{{ mode }}</p>
      </div>
    </header>
    <form @submit.prevent="submit" @keydown.enter.prevent="submit" class="flex flex-col ps-5 md:ps-28 pt-8">
      <div class="flex gap-x-5">
        <input type="file" ref="fileInputRef" @change="onFileChange" class="hidden" />
        <div class="inline-flex items-center group">
          <button @click.prevent="openFileInput" :class="{ 'text-black': file, 'text-gray-400': !file }"
            class="ms-1 group-hover:text-gray-600">
            {{ filename }}
          </button>
        </div>
      </div>
      <input v-if="ItemForm.title" v-model="title" required
        class="text-3xl md:text-6xl placeholder-gray-300 text-black font-bold border-none tracking-tight -ms-3 focus:ring-0"
        :placeholder="ItemForm.title" />
      <textarea v-if="ItemForm.url !== 'photos'" v-model="description" required
        class="md:text-xl placeholder-gray-400 text-black border-none -ms-2 focus:ring-0"
        placeholder="Начните писать описание..." rows="15"></textarea>
    </form>
    <button @click="submit"
      class="bg-red-600 hover:bg-red-400 w-full text-xl font-bold py-3 rounded-xl absolute bottom-0 lg:w-auto lg:right-4 lg:bottom-3 lg:px-7 text-white duration-300 ease-out hover:scale-105">Применить</button>
  </body>
</template>

<style>
.dp__input {
  border: 0;
}
</style>
