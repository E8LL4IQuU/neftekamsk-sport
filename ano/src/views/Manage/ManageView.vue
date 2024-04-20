<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import Sidebar from '@/components/Manage/Sidebar.vue'
import axios from 'axios'
import { ArchiveBoxIcon, CalendarIcon, MegaphoneIcon, PhotoIcon, TrophyIcon } from '@heroicons/vue/24/outline'
import { ArchiveBoxIcon as ArchiveBoxSolidIcon, CalendarIcon as CalendarSolidIcon, MegaphoneIcon as MegaphoneSolidIcon, PhotoIcon as PhotoSolidIcon, TrophyIcon as TrophySolidIcon } from '@heroicons/vue/24/solid'

const url: string = import.meta.env.VITE_ENDPOINT
const INVALID_USER_ID: number = 0
const isLoggedIn = ref<number>(0)
const router = useRouter()

// FIXME: The event vnode-unmounted is not declared in the emits option. It will leak into the component's attributes ($attrs)

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

const sidebarData = [
    { name: 'Мероприятия', path: '/manage/events', icon: CalendarIcon, iconActive: CalendarSolidIcon},
    { name: 'Новости', path: '/manage/news', icon: MegaphoneIcon, iconActive: MegaphoneSolidIcon},
    { name: 'Записи', path: '/manage/signups', icon: ArchiveBoxIcon, iconActive: ArchiveBoxSolidIcon},
    { name: 'Галлерея', path: '/manage/photos', icon: PhotoIcon, iconActive: PhotoSolidIcon},
    { name: 'Спортсмены', path: '/manage/athletes', icon: TrophyIcon, iconActive: TrophySolidIcon}
]
</script>

<template>
  <body class="bg-white">
    <Sidebar v-if="isLoggedIn" :sidebarData=sidebarData></Sidebar>
    <RouterView v-if="isLoggedIn" class="pb-20 lg:pb-0"></RouterView>
  </body>
</template>

<style scoped style="scss"></style>