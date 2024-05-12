<script setup lang="ts">

</script>

<template>
  <!-- TODO: mobile layout -->
  <div class="container mx-auto mb-8 ">
    <h2 class="text-black pt-[14px] pb-[18px] font-light text-4xl mobile:text-3xl">Расписание</h2>
    <p class="text-[24px] text-neutral-600 font-light leading-6">
      Для записи на занятия, используйте наш WhatsApp
      <!-- TODO: Change href to copy to clipboard on click -->
      <a href="tel:+79173522080" class="text-[#585858]"> +7 (917) 352 20 80</a>
    </p>
    <p class="text-[24px] text-neutral-600 font-light">Занятия проводятся по адресу:
      <span class="text-[#585858]">г. Нефтекамск, ул.Трактовая 14 а, 2 эт. (здание "Прогресс") </span>
    </p>
    <!-- FIXME: fix schedule layout -->
    <div class="flex justify-between gap-0.5 pt-9">
      <div class="grid grid-rows-7 gap-0.5">
        <!-- <div class="bg-[#B5B5B5]">
          <p class="text-white font-bold text-[20px] p-[14px]">Название</p>
        </div> -->
        <div class="bg-[#B5B5B5]">
          <p class="text-white font-bold text-[20px] p-[14px] w-[165px] pl-[9px]">Пн</p>
        </div>
        <div class="bg-[#B5B5B5]">
          <p class="text-white font-bold text-[20px] p-[14px] w-[165px] pl-[9px]">Вт</p>
        </div>
        <div class="bg-[#B5B5B5]">
          <p class="text-white font-bold text-[20px] p-[14px] w-[165px] pl-[9px]">Ср</p>
        </div>
        <div class="bg-[#B5B5B5]">
          <p class="text-white font-bold text-[20px] p-[14px] w-[165px] pl-[9px]">Чт</p>
        </div>
        <div class="bg-[#B5B5B5]">
          <p class="text-white font-bold text-[20px] p-[14px] w-[165px] pl-[9px]">Пт</p>
        </div>
        <div class="bg-[#B5B5B5]">
          <p class="text-white font-bold text-[20px] p-[14px] w-[165px] pl-[9px]">Сб</p>
        </div>
        <div class="bg-[#B5B5B5]">
          <p class="text-white font-bold text-[20px] p-[14px] w-[165px] pl-[9px]">Вс</p>
        </div>
      </div>
      <div class="grid grid-rows-7 flex-grow gap-0.5">
        <!-- <div class="bg-[#B5B5B5] ">
          <p class="text-white font-bold text-[20px] p-[14px]">Занятие 1</p>
        </div> -->
        <div class="bg-[#B5B5B5]">
          <p class="text-white font-bold text-[20px] p-[14px] w-[165px] pl-[9px]">

            11:00 - 11:45 - ЛФК
            12:00 - 13:30 - ОФП. Силовая подготовка -стадион "Торпедо"
          </p>
        </div>
        <div class="bg-[#B5B5B5]">
          <p class="text-white font-bold text-[20px] p-[14px] w-[165px] pl-[9px]">

            11:00 - 11:45 - пилатес
          </p>
        </div>
        <div class="bg-[#B5B5B5]">
          <p class="text-white font-bold text-[20px] p-[14px] w-[165px] pl-[9px]">
            12.00-13.30- Творческая мастерская
            12:00 - 13:30 - ОФП. Силовая подготовка -стадион "Торпедо"
            15:00 - 16.30 - Шашки
            16:30 - 18.30 - Шахматы
          </p>
        </div>
        <div class="bg-[#B5B5B5]">
          <p class="text-white font-bold text-[20px] p-[14px] w-[165px] pl-[9px]">
            11:00 - 11:45 - пилатес
          </p>
        </div>
        <div class="bg-[#B5B5B5]">
          <p class="text-white font-bold text-[20px] p-[14px] w-[165px] pl-[9px]">
            11:00-11:45 - ЛФК
          </p>
        </div>
        <div class="bg-[#B5B5B5]">
          <p class="text-white font-bold text-[20px] p-[14px] w-[165px] pl-[9px]">
            12.00-15.00 - мастер-класс по живописи
            15:00 - 16:30 - Шашки
            16:30 - 18:30 - Шахматы
          </p>
        </div>
        <div class="bg-[#B5B5B5]">
          <p class="text-white font-bold text-[20px] p-[14px] w-[165px] pl-[9px]">
            15:00 - 17:00 - волейбол (спортзал СОШ 6)

            Запись на занятия проводится по телефону : 8-917-352-20-80
          </p>
        </div>
      </div>
    </div>
  </div>
</template>

<style></style>