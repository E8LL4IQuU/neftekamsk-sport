<script setup lang="ts">
import { ref, onMounted } from 'vue';
import axios from 'axios'
import { useRouter, RouterLink } from 'vue-router'

const INVALID_USER_ID: number = 0;
const url: string = import.meta.env.VITE_ENDPOINT
const router = useRouter()
const isLoggedIn = ref<number>(0)

onMounted(async () => {
  try {
    const response = await axios.get<{ id: number }>(`${url}/api/auth/user`, {
      withCredentials: true,
    })

    if (response.data.id !== INVALID_USER_ID) {
      isLoggedIn.value = 1;
    }
  } catch (error) {
    if (axios.isAxiosError(error) && error.response?.status === 401) {
      router.push("/login");
    }
  }
})
</script>

<template>
  <div class="py-8 px-12 w-full">
    <div class="flex justify-between">
      <h1 class="text-3xl text-black font-bold">Мероприятия</h1>
      <router-link :to="'/manage/events/create'">
        <button class="bg-black rounded text-white py-2 px-3 ml-auto">Создать</button>
      </router-link>
    </div>
  </div>
</template>

<style scoped></style>
