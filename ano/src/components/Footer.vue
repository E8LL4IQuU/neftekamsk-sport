<template>
  <!-- FIXME: Footer is flickering on some pages while loading -->
  <!-- FIXME: it only shows 4 links although navbar has 5 -->

  <!-- Desktop -->
  <div class="hidden md:flex flex-col justify-center bg-neutral-800 p-10">
    <div class="flex text-center text-neutral-400 justify-center gap-x-5 pb-2">
      <router-link v-for="nav in props.navbarData" :to="nav.path">
        {{ nav.name }}
      </router-link>
    </div>

    <hr class="w-[30rem] justify-center text-center mx-auto mb-2 border-neutral-500">
    <div class="flex items-center justify-center gap-10 text-neutral-200">
      <div class="flex items-center gap-x-3">
        <img class="w-7" src="@/assets/mail.svg" alt="mail">
        <a href="mailto: neftekamsk.sport@mail.ru">neftekamsk.sport@mail.ru</a>
      </div>
      <div class="flex items-center gap-x-3">
        <img class="w-7" src="@/assets/wats.svg" alt="wats">
        <a rel='noopener noreferrer'
          href="https://api.whatsapp.com/send/?phone=79173522080&text&type=phone_number&app_absent=0">
          +7 (917) 352 20 80
        </a>
      </div>
    </div>
  </div>

  <!-- Mobile -->
  <div class="md:hidden flex flex-col gap-y-9 p-6 bg-neutral-800 text-white">
    <div>
      <h2 class="text-gray-300 font-bold">Информация</h2>
      <hr class="w-20 my-2">
      <div v-for="index in Array(2).keys()">
        <router-link :to="props.navbarData[index].path">
          <span class="leading-6">{{ props.navbarData[index].name }}</span>
        </router-link>
      </div>
    </div>
    <div>
      <h2 class="text-gray-300 font-bold">О нас</h2>
      <hr class="w-20 my-2">
      <!-- FIXME: make dynamic from number of elements in navbarData -->
      <div v-for="index in Array(2, 3)">
        <router-link :to="props.navbarData[index].path">
          <span class="leading-6">{{ props.navbarData[index].name }}</span>
        </router-link>
      </div>
    </div>
    <div>
      <div class="flex gap-x-1">
        <img class="w-7" src="@/assets/mail.svg" alt="mail">
        <a href="mailto: neftekamsk.sport@mail.ru">neftekamsk.sport@mail.ru</a>
      </div>
      <div class="flex gap-x-1">
        <img class="w-7" src="@/assets/wats.svg" alt="wats">
        <a rel='noopener noreferrer'
          href="https://api.whatsapp.com/send/?phone=79173522080&text&type=phone_number&app_absent=0">
          +7 (917) 352 20 80
        </a>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { RouterLink } from 'vue-router';

const props = defineProps(['navbarData']);
</script>

<style scoped></style>