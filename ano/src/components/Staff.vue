<template>
  <div>
    <div class="grid grid-cols-2 mx-auto max-w-3xl lg:max-w-5xl">
      <div class="" v-for="(people, index) in PeopleData" :key="index">
        <div class="text-center">
          <div>
            <img class="w-36 lg:w-48 rounded-full mb-3 mx-auto" :src="people.img" />
          </div>
          <div class="text-black">
            <h4 class="text-2xl lg:text-4xl pb-[3px] font-bold ">{{ people.name }}</h4>
            <span class="lg:text-2xl font-medium  text-neutral-600">{{
        people.activity
      }}</span>
            <p class="text-2xl font-medium">{{ people.description }}</p>
          </div>
        </div>
      </div>
    </div>
    <div class="text-center mb-[40px]">
      <!-- <router-link to="/people">
        <span
          class="text-white bg-black pt-[13px] pb-[13px] pr-[20px] pl-[20px] rounded-[10px]  hover:bg-[#474747] duration-300 cursor-pointer">
          Больше историй людей</span>
      </router-link> -->
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
  PeopleData: IPeople[];
}>();
</script>

<style scoped></style>
