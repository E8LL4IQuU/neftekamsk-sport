<template>
  <div>
    <div class="flex justify-center flex-wrap">
      <div class="m-auto" v-for="(people, index) in PeopleData" :key="index">
        <div class="text-center max-w-[328px] pt-5 pb-[60px]">
          <div>
            <img class="w-[150px] h-[150px] rounded-full mb-[20px] m-auto" :src="people.img">
          </div>
          <div class="text-black">
            <h4 class="pb-[3px] font-bold text-[36px]">{{ people.name }}</h4>
            <span class="font-medium text-[24px] text-[#727272]">{{ people.activity }}</span>
            <p class=" font-medium text-[24px]">{{ people.description }}</p>
          </div>
        </div>
      </div>
    </div>
    <div class="text-center mb-[40px]">
      <router-link to="/people">
        <span
          class="text-white bg-black pt-[13px] pb-[13px] pr-[20px] pl-[20px] rounded-[10px]  hover:bg-[#474747] duration-300 cursor-pointer">
          Больше историй людей</span>
      </router-link>

    </div>
  </div>
</template>

<script setup lang="ts">
interface IPeople {
  img: string;
  name: string;
  activity: string;
  description: string;
}

const props = defineProps<{
  PeopleData: IPeople[]
}>()
</script>

<style scoped></style>