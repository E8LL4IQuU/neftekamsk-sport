<script setup lang="ts">

</script>

<template>
  <div class="containers">
    <div class="flex justify-between py-[88px] gap-3 laptop:block laptop:text-center mobile:py-[35px]" >
      <div>
        <router-link to="/news-item">
          <img class="w-[650px] h-[520px] rounded-[30px] mac:w-[500px] mac:h-[400px] laptop:mx-auto mobile:w-full" src="../assets/images.jpeg"/>
        </router-link>
      </div>
      <div class="max-w-[665px] mac:max-w-[500px] laptop:mx-auto mobile:w-full">
        <h4 class="text-black text-center text-[36px] pb-[14px] mac:text-[30px]">Якубович в Нефтекамске</h4>
        <hr class="max-w-[540px] bg-[#E4E4E4] mx-auto">
        <p class="text-[#6F6F6F] text-[28px] pt-[14px] text-right  mac:text-[20px] laptop:text-center">Давно выяснено, что при оценке дизайна и композиции читаемый текст мешает сосредоточиться. Lorem Ipsum используют потому, что тот обеспечивает более или менее стандартное заполнение шаблона, а также реальное распределение букв и пробелов в абзацах, которое не получается при простой дубликации "Здесь ваш текст..</p>
      </div>
    </div>
  </div>
</template>

<style scoped>
.containers{
  max-width: 1450px;
  margin: 0 auto;
}
</style>