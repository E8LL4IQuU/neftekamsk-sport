<template>
  <div class="flex gap-10 content-center items-center tablet:gap-2 mobile:block">
    <div class="hidden text-black text-center tablet:w-[100%] mobile:block">
      <h4 class="text-5xl laptop:text-3xl pb-6 tablet:text-2xl mobile:pb-2 font-semibold">
        АНО “Нефтекамск Спортивный”
      </h4>
      <p class="text-2xl laptop:text-xl tablet:text-sm pb-3">
        Тренировочная, физкультурная и туристическая деятельности
      </p>
    </div>
    <div>
      <!-- TODO: maybe add white fade to the right side of the image -->
      <img class="tablet:w-[100%]" src="@/assets/heroBackground.png" alt="image-people">
    </div>
    <div class="text-black text-center tablet:w-[100%] mobile:hidden">
      <h4 class="text-5xl laptop:text-3xl pb-6 tablet:text-2xl font-semibold">
        АНО “Нефтекамск Спортивный”
      </h4>
      <p class="text-2xl laptop:text-xl tablet:text-sm m-0">
        Тренировочная, физкультурная и туристическая деятельности
      </p>
    </div>
  </div>
</template>

<script setup lang="ts">
</script>

<style scoped></style>