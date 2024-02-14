<script setup lang="ts">
import { RouterView } from 'vue-router'

const navbarItems: string[][] = [
  // [Navbar item name, Navbar item link]
  ["Новости", "/news"],
  ["Расписание", "/schedule"],
  ["Направления", "/sections"],
  ["Галерея", "/gallery"],
  ["Истории людей", "/people"],
  ["Документы", "/documents"]
]
const NavData = [
  {
    "links": "/news",
    "name": "Новости"
  },
  {
    "links": "/schedule",
    "name": "Расписание"
  },
  {
    "links": "/sections",
    "name": "Направления"
  },
  {
    "links": "/gallery",
    "name": "Галерея"
  },
  {
    "links": "/people",
    "name": "Истории людей"
  },
  {
    "links": "/documents",
    "name": "Документы"
  },
]
</script>

<template>
  <div class="bg-white">
    <!-- Management -->
    <RouterView name="manageSidebar"></RouterView>

    <RouterView name="navbar" :navbarItems="navbarItems"></RouterView>
    <RouterView />
    <RouterView name="footer" :NavData="NavData"></RouterView>
  </div>
</template>

<style scoped></style>
