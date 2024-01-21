<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import Sidebar from '@/components/Sidebar.vue'
import axios from 'axios'

const url: string = import.meta.env.VITE_ENDPOINT
const INVALID_USER_ID: number = 0
const isLoggedIn = ref<number>(0)
const router = useRouter()

onMounted( async () => {
    try {
    const response = await axios.get<{ id: number }>(`${url}/api/auth/user`, {
      withCredentials: true,
    })

    if (response.data.id !== INVALID_USER_ID) {
      isLoggedIn.value = 1;
    }
  } catch (error) {
    if (axios.isAxiosError(error) && error.response?.status === 401) {
      console.log("pushing to /login")
      router.push("/login");
    }
  }
})
</script>

<template>
<div class="flex">
    <Sidebar class=" w-2/12"></Sidebar>
    <RouterView class="w-10/12"></RouterView>
</div>
</template>

<style scoped style="scss">

</style>