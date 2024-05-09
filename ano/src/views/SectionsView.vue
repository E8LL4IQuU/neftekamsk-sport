<script setup lang="ts">
import {ref} from "vue";
import Tabs from "@/components/Tabs.vue";

// FIXME: enforce fixed aspect ratio to images
// FIXME: Mobile layout
// FIXME: Fix gaps between categories
// FIXME: Extraneous non-props attributes (modelValue) were passed to component but could not be automatically inherited because component renders fragment or text root nodes.
const selectedTab = ref('tab-br1')
</script>

<template>
  <div>
  <div class="max-w-[1200px] m-auto">
    <div>
      <h1 class="text-black text-[40px] font-light pb-[24px]">Направления</h1>
    </div>
  </div>
  <div>
    <tabs v-model="selectedTab">
      <template #tab-br1>
        <p class="text-black text-[24px] max-w-[1430px] text-center font-light py-[15px] mx-auto">Более 100 детей и подростков воспользовались физкультурно-оздоровительной услугой, посещая тренировочные занятия по детской лёгкой атлетике, волейболу, баскетболу, общей физической подготовке, подготовке к сдаче норм ГТО.В 2023 году запустили новые направления для людей старше 55 лет: «Йога + суставная гимнастика», ОФП, «Шахматы», «Шашки»</p>
        <div class="flex justify-center gap-[84px] max-w-[1200px] m-auto pc:block pc: w-full pc: mx-auto">
          <div>
            <img class="max-w-[600px] max-h-[400px] rounded-[30px] pc:mx-auto pc:pt-[20px] pc:w-full pc:h-full tablet:rounded-0" src="../assets/tabs-one.jpeg" alt="">
          </div>
          <div>
            <img class="max-w-[600px] max-h-[400px] rounded-[30px] pc:mx-auto pc:pt-[20px] pc:w-full pc:h-full tablet:rounded-0" src="../assets/tabs-one2.jpeg" alt="">
          </div>
        </div>
        <div class="py-[30px] text-center">
        <router-link to="/gallery" class="text-white bg-black pt-[13px] pb-[13px] pr-[20px] pl-[20px] rounded-[10px]  hover:bg-[#474747] duration-300 cursor-pointer">
          Смотреть еще фото
        </router-link>
        </div>
      </template>
      <template #tab-br2>
        <p class="text-black text-[24px] max-w-[1430px] text-center font-light py-[15px] mx-auto">АНО «Нефтекамск спортивный» является одной из немногих организаций, проводящих физкультурно-оздоровительные и туристские мероприятия для коллективов организаций и предприятий Нефтекамска.
          <br>
          <br>
          Среди наших мероприятий: зимние и летние Дни здоровья, туристические мероприятия, фестивали и слеты, спартакиады по видам спорта, квесты, физкультурно-оздоровительные мероприятия для школьников и дошкольников.
          <br>
          <br>
          Среди наших клиентов: ООО «Центр отопления», ЦГБ г. Нефтекамск, АО «Белкамнефть», АО «Транснефть Урал», ПАО «Сбербанк-Нефтекамск», КГРЭС, БашРТС-Нефтекамск, УЖКХ-Нефтекамск.</p>
        <div class="flex justify-center gap-[84px] max-w-[1200px] m-auto pc:block pc: w-full pc: mx-auto">
          <div>
            <img class="max-w-[600px] max-h-[400px] rounded-[30px] pc:mx-auto pc:pt-[20px] pc:w-full pc:h-full tablet:rounded-0" src="../assets/tabs-two.jpeg" alt="">
          </div>
          <div>
            <img class="max-w-[600px] max-h-[400px] rounded-[30px] pc:mx-auto pc:pt-[20px] pc:w-full pc:h-full tablet:rounded-0" src="../assets/tabs-two2.jpeg" alt="">
          </div>
        </div>
        <div class="py-[30px] text-center">
          <router-link to="/gallery" class="text-white bg-black pt-[13px] pb-[13px] pr-[20px] pl-[20px] rounded-[10px]  hover:bg-[#474747] duration-300 cursor-pointer">
            Смотреть еще фото
          </router-link>
        </div>
      </template>
      <template #tab-br3>
        <p class="text-black text-[24px] max-w-[1430px] text-center font-light py-[15px] mx-auto">Благодаря наличию собственного туристского оборудования мы организуем походы и водные сплавы различной продолжительности и сложности:
          Многодневные походы и водные сплавы по Южному Уралу ( реки Белая, Инзер, Зелим, Ай, Юрюзань, озеро Тургаяк, Зюраткуль, национальный парк «Таганай»,
          Мурадымовское ущелье и др.
          Однодневные водные сплавы по р. Буй.</p>
        <div class="flex justify-center gap-[84px] max-w-[1200px] m-auto pc:block pc: w-full pc: mx-auto">
          <div>
            <img class="max-w-[600px] max-h-[400px] rounded-[30px] pc:mx-auto pc:pt-[20px] pc:w-full pc:h-full tablet:rounded-0" src="../assets/tabs-three.jpeg" alt="">
          </div>
          <div>
            <img class="max-w-[600px] max-h-[400px] rounded-[30px] pc:mx-auto pc:pt-[20px] pc:w-full pc:h-full tablet:rounded-0" src="../assets/tabs-three2.jpeg" alt="">
          </div>
        </div>
        <div class="py-[30px] text-center">
          <router-link to="/gallery" class="text-white bg-black pt-[13px] pb-[13px] pr-[20px] pl-[20px] rounded-[10px]  hover:bg-[#474747] duration-300 cursor-pointer">
            Смотреть еще фото
          </router-link>
        </div>
      </template>
      <template #tab-br4>
        <p class="text-black text-[24px] max-w-[1430px] text-center font-light py-[15px] mx-auto">Оказываем информационную, организационную помощь ветеранам спорта в организации поездок на соревнования, организации физкультурно-оздоровительных занятий.</p>
        <div class="flex justify-center gap-[84px] max-w-[1200px] m-auto pc:block pc: w-full">
          <div>
            <img class="max-w-[600px] max-h-[400px] rounded-[30px] pc:mx-auto pc:pt-[20px] pc:w-full pc:h-full tablet:rounded-0" src="../assets/tabs-four.jpeg" alt="">
          </div>
          <div>
            <img class="max-w-[600px] max-h-[400px] rounded-[30px] pc:mx-auto pc:pt-[20px] pc:w-full pc:h-full tablet:rounded-0" src="../assets/tabs-four2.jpeg" alt="">
          </div>
        </div>
        <div class="py-[30px] text-center">
          <router-link to="/gallery" class="text-white bg-black pt-[13px] pb-[13px] pr-[20px] pl-[20px] rounded-[10px]  hover:bg-[#474747] duration-300 cursor-pointer">
            Смотреть еще фото
          </router-link>
        </div>
      </template>
      <template #tab-br5>
        <p class="text-black text-[24px] max-w-[1430px] text-center font-light py-[15px] mx-auto">ЖДЕМС</p>
      </template>
    </tabs>
  </div>
  </div>
</template>

<style scoped>

</style>