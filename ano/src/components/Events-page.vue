<template>
  <div class="max-w-[1250px]  m-auto mb-8 ">
    <div class="flex justify-between items-center p-5">
      <div>
        <h2 class="text-black pt-[14px] pb-[18px] font-light text-4xl mobile:text-3xl">Мероприятия</h2>
      </div>
      <div class="text-black">
        <select v-model="selectedFilter" class="focus-visible:none" @change="filterEvents">
          <option value="0">Все</option>
          <option value="1">Текущие</option>
          <option value="2">Прошедшие</option>
          <option value="3">Предстоящие</option>
        </select>
      </div>
    </div>
    <div v-if="filteredEvents.length > 0" class="flex flex-wrap gap-5 justify-center tablet:p-5">
      <router-link v-for="(event, index) in filteredEvents" :to="'events-page-item/' + index " :key="index" >
      <div style="position: relative;" class="hover:opacity-[85%]" :class="{ 'opacity-50': event.status === 'Прошедшее' }">
        <img class="w-[600px] h-[350px]" :src="event.image">
        <div class="bg-white bg-opacity-[80%] absolute top-2 left-2 rounded-2xl">
          <span class="font-light text-black text-base pb-[5px] pt-[5px] pr-[18px] pl-[18px]">{{ event.status }}</span>
        </div>
      </div>
      <div class="text-black flex justify-between max-w-[600px] pb-2 pt-[10px]">
        <div>
          <h3 class="text-xl font-bold uppercase mobile:text-[16px]">{{ event.name }}</h3>
        </div>
        <div>
          <span class="text-gray-600 text-base font-light mobile:text-[14px]">{{ event.date }}</span>
        </div>
      </div>
      <div class="text-gray-500">
        <p class="max-w-[600px]">{{ event.description }}</p>
      </div>
    </router-link>
    </div>
    <div v-else class="text-gray-500 text-6xl">Мероприятия не найдены</div>
  </div>
</template>

<script setup lang="ts">
import {ref} from "vue";
const events = [
  {
    image: '/src/components/icons/Background.png',
    status: 'Текущее',
    name: 'ПОСИДЕЛКИ С ЯКУБОВИЧЕМ',
    date: '12.12.2023',
    description: 'Давно выяснено, что при оценке дизайна и композиции читаемый текст мешает сосредоточиться. Lorem Ipsum используют потому, что тот...'
  },
  {
    image: '/src/components/icons/Background.png',
    status: 'Прошедшее',
    name: 'ПОСИДЕЛКИ С ЯКУБОВИЧЕМ',
    date: '12.12.2023',
    description: 'Давно выяснено, что при оценке дизайна и композиции читаемый текст мешает сосредоточиться. Lorem Ipsum используют потому, что тот...'
  },
  {
    image: '/src/components/icons/Background.png',
    status: 'Текущее',
    name: 'ПОСИДЕЛКИ С ЯКУБОВИЧЕМ',
    date: '12.12.2023',
    description: 'Давно выяснено, что при оценке дизайна и композиции читаемый текст мешает сосредоточиться. Lorem Ipsum используют потому, что тот...'
  },
  {
    image: '/src/components/icons/Background.png',
    status: 'Текущее',
    name: 'ПОСИДЕЛКИ С ЯКУБОВИЧЕМ',
    date: '12.12.2023',
    description: 'Давно выяснено, что при оценке дизайна и композиции читаемый текст мешает сосредоточиться. Lorem Ipsum используют потому, что тот...'
  }
];
let selectedFilter = "0";

const filteredEvents = ref(events);

const filterEvents = () => {
  switch (selectedFilter) {
    case "1":
      filteredEvents.value = events.filter(event => event.status === "Текущее");
      break;
    case "2":
      filteredEvents.value = events.filter(event => event.status === "Прошедшее");
      break;
    case "3":
      filteredEvents.value = events.filter(event => event.status === "Текущее");
      break;
    default:
      filteredEvents.value = events;
      break;
  }
};
</script>

<style scoped>

</style>