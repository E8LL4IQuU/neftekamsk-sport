<script setup lang="ts">
import { RouterLink } from "vue-router";
</script>

<template>
    <div class="min-h-screen p-4 border-r">
      <h1 class="text-3xl mb-5 text-black">Управление</h1>
      <nav>
        <ul>
          <li class="flex flex-col">
            <router-link :to="'/manage/events'" :class="{ 'text-black font-bold' : $route.path === '/manage/events', 'text-gray-700' : $route.path !== '/manage/events'}">Мероприятия</router-link>
            <router-link :to="'/manage/news'" :class="{ 'text-black font-bold' : $route.path === '/manage/news', 'text-gray-700' : $route.path !== '/manage/news'}">Новости</router-link>
          </li>
        </ul>
      </nav>
    </div>
</template>

<style scoped></style>
