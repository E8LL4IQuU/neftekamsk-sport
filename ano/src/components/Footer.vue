<template>
  <!-- FIXME: Footer is flickering on some pages while loading -->
  <div class="flex flex-col justify-center bg-[#242424] p-10">
    <div class="flex text-center text-neutral-400 justify-center gap-x-5 pb-2 tablet:gap-2 mobile:block mobile:pt-[10px]">
      <div v-for="(nav, index) in NavData" :key="index">
        <div>
          <router-link class="tablet:text-[13px]" :to="nav.links">{{ nav.name }}</router-link>
        </div>
      </div>
    </div>

    <hr class="w-[30rem] justify-center text-center mx-auto mb-2 border-neutral-500">
    <div class="flex items-center justify-center gap-10 tablet:gap-5 mobile:block text-neutral-200">
      <div class="flex items-center gap-5 tablet:gap-2 mobile:justify-center">
        <img class="w-7" src="@/assets/mail.svg" alt="mail">
        <a href="mailto: neftekamsk.sport@mail.ru">neftekamsk.sport@mail.ru</a>
      </div>
      <div class="flex items-center gap-5 tablet:gap-2 mobile:justify-center">
        <img class="w-7" src="@/assets/wats.svg" alt="wats">
        <a rel='noopener noreferrer'
          href="https://api.whatsapp.com/send/?phone=79173522080&text&type=phone_number&app_absent=0">+7 (917) 352 20
          80</a>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { RouterLink } from 'vue-router';

interface INav {
  links: string;
  name: string
}

const props = defineProps<{
  NavData: INav[]
}>()

</script>

<style scoped></style>