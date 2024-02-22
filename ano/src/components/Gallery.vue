<template>
  <div class="lg:grid lg:grid-flow-col lg:gap-x-5 mobile:gap-2">
    <div class="pt-[20%] pc:pt-2">
      <img
        class="rounded-[30px] h-[495px] pc:m-auto w-[700px] tablet:w-[500px] tablet:h-[313px] mobile:w-[83%]  mobile:h-98vh]"
        :src="props.GalleryData.linkImgOne">
      <div class="pt-4 pc:pt-2 ">
        <img
          class="rounded-[30px] w-[700px] h-[426px] pc:m-auto tablet:w-[500px] tablet:h-[313px] mobile:w-[83%]  mobile:h-98vh]"
          :src="props.GalleryData.linkImgTwo">
      </div>
    </div>
    <div class=" pc:pt-2">
      <img
        class="rounded-[30px] w-[700px] h-[426px] pc:m-auto tablet:w-[500px] tablet:h-[313px] mobile:w-[83%] mobile:h-98vh]"
        :src="props.GalleryData.linkImgThree">
      <video class=" w-[700px] h-[426px] pc:m-auto tablet:w-[500px] tablet:h-[313px] mobile:w-[83%] mobile:h-98vh]"
        :src="props.GalleryData.videos" controls autoplay muted="true"></video>
    </div>
  </div>
</template>

<script setup lang="ts">
interface IGallery {
  // FIXME: how about link[]
  linkImgOne: string;
  linkImgTwo: string;
  linkImgThree: string;
  videos: string
}
const props = defineProps<{
  GalleryData: IGallery
}>()
</script>

<style scoped></style>