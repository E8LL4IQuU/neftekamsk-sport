<script setup lang="ts">
import { RouterLink } from "vue-router";
import { ArchiveBoxIcon, CalendarIcon, MegaphoneIcon } from '@heroicons/vue/24/solid'
</script>

<template>
  <!-- Desktop -->
  <nav class="hidden lg:block min-h-screen p-4 border-r">
    <h1 class="text-3xl mb-5 text-black">Управление</h1>
    <ul>
      <li class="flex flex-col">
        <router-link :to="'/manage/events'"
          :class="{ 'text-black font-bold': $route.path === '/manage/events', 'text-gray-700': $route.path !== '/manage/events' }">Мероприятия</router-link>
        <router-link :to="'/manage/news'"
          :class="{ 'text-black font-bold': $route.path === '/manage/news', 'text-gray-700': $route.path !== '/manage/news' }">Новости</router-link>
        <router-link :to="'/manage/signups'"
          :class="{ 'text-black font-bold': $route.path === '/manage/signups', 'text-gray-700': $route.path !== '/manage/signups' }">Записи</router-link>
      </li>
    </ul>
    <!-- TODO: add logout button
      POST /api/auth/logout -->
  </nav>

  <!-- Mobile -->
  <nav class="lg:hidden text-start fixed bottom-0 w-full bg-white p-2">
    <ul class="grid grid-cols-3 text-center justify-between gap-x-3">
      <router-link class="h-20 rounded-lg flex flex-col justify-center" :to="'/manage/events'"
        :class="{ 'text-black font-bold bg-red-300': $route.path === '/manage/events', 'text-gray-700 bg-slate-300': $route.path !== '/manage/events' }">
        <CalendarIcon class="h-11"></CalendarIcon>
        <span>Мероприятия</span>
      </router-link>
      <router-link class="h-20 rounded-lg flex flex-col justify-center" :to="'/manage/news'"
        :class="{ 'text-black font-bold bg-red-300': $route.path === '/manage/news', 'text-gray-700 bg-slate-300': $route.path !== '/manage/news' }">
        <MegaphoneIcon class="h-11"></MegaphoneIcon>
        <span>Новости</span>
      </router-link>
      <router-link class="h-20 rounded-lg flex flex-col justify-center" :to="'/manage/signups'"
        :class="{ 'text-black font-bold bg-red-300': $route.path === '/manage/signups', 'text-gray-700 bg-slate-300': $route.path !== '/manage/signups' }">
        <ArchiveBoxIcon class="h-11"></ArchiveBoxIcon>
        <span>Записи</span>
      </router-link>
    </ul>
    <!-- TODO: add logout button
      POST /api/auth/logout -->
  </nav>
</template>

<style scoped></style>
