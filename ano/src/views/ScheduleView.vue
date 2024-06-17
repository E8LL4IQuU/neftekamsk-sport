<template>
  <!-- FIXME: white text on the white background -->

  <body class="bg-white">
    <!-- TODO: mobile layout -->
    <div class="container mx-auto mb-8 ">
      <!-- FIXME: inconsistent header text -->
      <h2 class="text-black pt-[14px] text-4xl font-bold mb-8 mobile:text-3xl">Расписание</h2>
      <p class="text-[24px] text-neutral-600 font-light leading-6">
        Для записи на занятия, используйте наш WhatsApp
        <!-- TODO: Change href to copy to clipboard on click -->
        <a href="tel:+79173522080" class="text-[#585858]"> +7 (917) 352 20 80</a>
      </p>
      <p class="text-[24px] text-neutral-600 font-light">Занятия проводятся по адресу:
        <span class="text-[#585858]">г. Нефтекамск, ул.Трактовая 14 а, 2 эт. (здание "Прогресс") </span>
      </p>
      <div class="flex flex-col justify-center pt-1">
        <table class="w-full table-auto border-collapse border-gray-200">
          <thead>
            <tr>
              <th class="px-6 py-3 text-gray-500 text-left font-bold bg-slate-300 text-sm uppercase">День недели</th>
              <th class="px-6 py-3 text-gray-500 text-left font-bold bg-slate-300 text-sm uppercase">Занятия</th>
            </tr>
          </thead>
          <tbody>
            <tr v-for="(item, weekIndex) in weeklySchedule" :key="weekIndex"
              :class="{ 'bg-slate-300': (weekIndex % 2 == 1) }">
              <td class="px-6 py-3 text-left font-bold">{{ item.day }}</td>
              <td>
                <div :class="{ 'bg-slate-300': (weekIndex % 2 == 1) }" class="w-full">
                  <template v-for="(event, eventIndex) in item.events">
                    <div class="px-6 py-3 text-left">{{ event }}</div>
                    <!-- <span class="text-gray-500 text-sm">({{ getDuration(event) }})</span> -->
                  </template>
                </div>
              </td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>
  </body>
</template>

<style></style>

<script setup lang="ts">
import { ref } from 'vue'

interface ScheduleItem {
  day: string
  events: string[]
}

const scheduleItems = [
  {
    day: 'Понедельник',
    events: ['11:00 - 11:45 - ЛФК', '12:00 - 13:30 - ОФП. Силовая подготовка - Стадион "Торпедо"']
  },
  {
    day: 'Вторник',
    events: ['11:00 - 11:45 - Пилатес']
  },
  {
    day: 'Среда',
    events: ['12:00 - 13:30 - Творческая мастерская', '12:00 - 13:30 - ОФП. Силовая подготовка - Стадион "Торпедо"', '15:00 - 16:30 - Шашки', '16:30 - 18:30 - Шахматы']
  },
  {
    day: 'Четверг',
    events: ['11:00 - 11:45 - Пилатес']
  },
  {
    day: 'Пятница',
    events: ['11:00 - 11:45 - ЛФК']
  },
  {
    day: 'Суббота',
    events: ['12:00 - 15:00 - Мастер класс по живописи', '15:00 - 16:30 - Шашки', '16:30 - 18:30 - Шахматы']
  },
  {
    day: 'Воскресенье',
    events: ['15:00 - 17:00 - Волейболл - Спортзал СОШ 6']
  },
]

const weeklySchedule = ref(scheduleItems)

// TODO: implement duration of event
// function getDuration(event: string) {
//   const parts = event.split(' - ');
//   return `${parts[0]} (${parts[1]})`;
// }

</script>
