<template>
  <div class="bg-[#242424]">
    <div class="flex text-center justify-center gap-[54px] pt-[72px] pb-[14px] tablet:gap-2 mobile:block mobile:pt-[10px]">
      <div v-for="(nav, index) in NavData" :key="index">
        <div>
          <a class="tablet:text-[13px]" :href="nav.links">{{ nav.name }}</a>
        </div>
      </div>
    </div>
    <hr class="w-[680px] justify-center text-center m-auto  mb-[14px] bg-[#989898] tablet:w-[380px] mobile:w-[180px]">
    <div class="flex items-center justify-center gap-10 pb-6 tablet:gap-5 mobile:block">
      <div class="flex items-center gap-5 tablet:gap-2 mobile:justify-center">
        <img class="w-[45px] tablet:w-[25px]" src="../assets/mail.svg" alt="mail">
        <a href="mailto: neftekamsk.sport@mail.ru">neftekamsk.sport@mail.ru</a>
      </div>
      <div class="flex items-center gap-5 tablet:gap-2 mobile:justify-center">
        <img class="w-[45px] tablet:w-[25px]" src="../assets/wats.svg" alt="wats">
        <a href="https://api.whatsapp.com/send/?phone=79173522080&text&type=phone_number&app_absent=0">+7 (917) 352 20 80</a>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
interface INav {
  links: string;
  name: string
}

const props = defineProps<{
  NavData: INav[]
}>()

</script>

<style scoped>

</style>