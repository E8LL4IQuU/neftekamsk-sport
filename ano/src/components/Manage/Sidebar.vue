<script setup lang="ts">
import { RouterLink } from "vue-router";

const props = defineProps(['sidebarData'])
</script>

<template>
  <div class="lg:fixed lg:left-0">
    <!-- Desktop -->
    <nav class="hidden lg:block min-h-screen p-4 border-r">
      <h1 class="text-3xl mb-5 text-black">Управление</h1>
      <ul>
        <li class="flex flex-col">
          <router-link v-for="item in props.sidebarData" :to="item.path"
            :class="{ 'text-black font-bold': $route.path === item.path, 'text-gray-700': $route.path !== item.path }">
            {{ item.name }}
          </router-link>
        </li>
      </ul>
      <!-- TODO: add logout button
      POST /api/auth/logout -->
    </nav>

    <!-- Mobile -->
    <nav class="lg:hidden fixed bottom-0 w-full bg-stone-300 py-3 z-10">
      <ul class="grid grid-cols-4 text-center justify-between">
        <div v-for="item in props.sidebarData">
          <router-link v-if="$route.path === item.path" class="h-12 rounded-lg flex flex-col justify-between items-center"
            :to="item.path"
            :class="{ 'text-black': $route.path === item.path, 'text-gray-700': $route.path !== item.path }">
            <div>
              <component :class="{ 'bg-orange-300 saturate-50 rounded-2xl': $route.path === item.path }" class="h-7 w-16"
                :is="item.iconActive"></component>
            </div>
            <span class="font-bold text-sm">{{ item.name }}</span>
          </router-link>

          <router-link v-else class="h-12 rounded-lg flex flex-col justify-between items-center" :to="item.path"
            :class="{ 'text-black': $route.path === item.path, 'text-gray-700': $route.path !== item.path }">
            <div>
              <component :class="{ 'bg-orange-300 saturate-50 rounded-2xl': $route.path === item.path }" class="h-7 w-16"
                :is="item.icon"></component>
            </div>
            <span class="font-bold text-sm">{{ item.name }}</span>
          </router-link>
        </div>
      </ul>
      <!-- TODO: add logout button
      POST /api/auth/logout -->
    </nav>
  </div>
</template>

<style scoped></style>
