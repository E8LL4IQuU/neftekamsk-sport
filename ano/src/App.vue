<script setup lang="ts">
import { RouterView } from 'vue-router'

const navbarData = [
  {
    name: "Новости",
    path: "/news"
  },
  {
    name: "Расписание",
    path: "/schedule"
  },
  {
    name: "Направления",
    path: "/sections"
  },
  {
    name: "Галерея",
    path: "/gallery"
  },
  {
    name: "Люди спорта",
    path: "/people"
  },
  {
    name: "Документы",
    path: "/documents"
  },
]
</script>

<template>
  <div class="bg-white">
    <!-- @/components/Navbar.vue -->
    <RouterView name="navbar" :navbarData="navbarData"></RouterView>
    <RouterView />
    <!-- @/components/Footer.vue -->
    <RouterView name="footer" :navbarData="navbarData"></RouterView>
  </div>
</template>

<style scoped></style>
