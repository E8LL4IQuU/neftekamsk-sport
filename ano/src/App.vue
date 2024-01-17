<script setup lang="ts">
import {RouterView} from 'vue-router'
import NavBar from "./components/NavBar.vue"
import Footer from "@/components/Footer.vue"
import Sidebar from  "@/components/Sidebar.vue"
import ManageView from "@/views/Manage/ManageView.vue"

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
    <ManageView v-if="$route.meta.managementRoute"></ManageView>
    <NavBar :navbarItems="navbarItems" v-if="!$route.meta.hideNavbar"/>
    <RouterView v-if="!$route.meta.managementRoute" />
    <Footer :NavData="NavData" v-if="!$route.meta.hideFooter" />
  </div>
</template>

<style scoped>

</style>
