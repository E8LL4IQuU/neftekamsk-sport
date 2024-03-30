<template>
  <div class="lg:grid lg:grid-flow-col lg:gap-x-5">
    <div class="flex flex-col gap-y-3 lg:pt-20">
      <img class="w-full md:w-10/12 mx-auto md:rounded-3xl object-cover lg:w-[704px] lg:h-[396px]" :src="props.galleryData.photos[0]">
      <img class="w-full md:w-10/12 mx-auto md:rounded-3xl object-cover lg:w-[704px] lg:h-[396px]" :src="props.galleryData.photos[1]">
    </div>
    <div class="flex flex-col gap-y-3 mt-3 lg:mt-0">
      <img class="w-full md:w-10/12 mx-auto md:rounded-3xl object-cover lg:w-[704px] lg:h-[396px]" :src="props.galleryData.photos[2]">
      <video class="w-full md:w-10/12 mx-auto md:rounded-3xl object-cover lg:w-[704px] lg:h-[396px]" :src="props.galleryData.video" controls
        autoplay muted></video>
    </div>
  </div>
</template>

<script setup lang="ts">
const props = defineProps(['galleryData'])
</script>

<style scoped></style>