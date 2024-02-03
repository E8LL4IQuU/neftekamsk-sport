<script setup lang="ts">
import { ref, onMounted } from "vue";
import { useRoute } from "vue-router";
import axios from "axios";
import { type IRLEvent } from "@/types/apiTypes";

const url: string = import.meta.env.VITE_ENDPOINT;
const IRLEvent = ref<IRLEvent>();
const route = useRoute();
const EventId: number = Number(route.params.id);
const isRegistrationPopupOpen = ref<boolean>(false);
const FirstName = ref<string>('');
const LastName = ref<string>('');
const PhoneNumber = ref<string>('');
const Email = ref<string>('');

const openRegistrationPopup = () => {
  isRegistrationPopupOpen.value = true;
};

const closeRegistrationPopup = () => {
  isRegistrationPopupOpen.value = false;
};

const submitRegistrationForm = () => {
  console.log('Submitting Registration Form');
  closeRegistrationPopup();
};

const fetchIRLEvents = async (): Promise<void> => {
  try {
    const response = await axios.get<IRLEvent>(`${url}/api/events/${EventId}`);
    IRLEvent.value = response.data;
  } catch (error) {
    console.error("Error fetching IRLevents:", error);
  }
};



onMounted(() => {
  fetchIRLEvents();
});
</script>

<template>
  <div class="px-1 sm:px-0 sm:container mx-auto pb-20">
    <div class="sm:flex gap-x-12 justify-between" v-if="IRLEvent">
      <img class="h-full" :src="`${url}/uploads/${IRLEvent.img}`" alt="">
      <div>
        <h1 class="text-black text-3xl font-bold pb-3">{{ IRLEvent.title }}</h1>
        <p class="text-black max-w-7xl">{{ IRLEvent.description }}</p>
        <div>
          <button class="px-4 py-3 mt-3 rounded bg-red-400 text-white" @click="openRegistrationPopup">Записаться</button>

          <!-- Registration Popup -->
          <div v-if="isRegistrationPopupOpen"
            class="fixed inset-0 bg-gray-800 bg-opacity-75 flex items-center justify-center"
            @click.self="closeRegistrationPopup">
            <div @click.stop class="relative bg-white rounded-md sm:w-5/12 flex flex-col gap-y-3 items-center text-center pt-2 pb-6 px-3 sm:px-0">
              <button @click="closeRegistrationPopup"
                class="text-4xl absolute top-0 right-2 text-gray-600 hover:text-gray-800">&times;</button>
              <h2 class="text-2xl font-bold">ЗАПИСЬ</h2>
              <!-- Registration form content -->
              <form @submit.prevent="submitRegistrationForm" class="md:w-8/12 pb-6">
                <div class="mb-4">
                  <input required v-model="FirstName" type="text" id="FirstName"
                    class="mt-1 p-2 w-full border-none rounded-md bg-stone-200 placeholder-black placeholder-shown:font-bold"
                    placeholder="Имя">
                </div>

                <div class="mb-4">
                  <input required v-model="LastName" type="text" id="LastName"
                    class="mt-1 p-2 w-full border-none rounded-md bg-stone-200 placeholder-black placeholder-shown:font-bold"
                    placeholder="Фамилия">
                </div>

                <div class="mb-4">
                  <input required v-model="PhoneNumber" type="text" id="PhoneNumber"
                    class="mt-1 p-2 w-full border-none rounded-md bg-stone-200 placeholder-black placeholder-shown:font-bold"
                    placeholder="Номер телефона">
                </div>

                <!-- FIXME: make it optional -->
                <div class="mb-4">
                  <input v-model="Email" type="text" id="Email"
                    class="mt-1 p-2 w-full border-none rounded-md bg-stone-200 placeholder-black placeholder-shown:font-bold"
                    placeholder="Почта(необязательно)">
                </div>

                <button type="submit" class="bg-red-400 text-white px-4 py-2 mt-4 rounded-md w-full">Записаться</button>
              </form>
              <hr class="w-10/12 border-t-2 border-gray-300">
              <p class="font-bold">или вы можете записаться через whatsapp</p>
              <div class="flex items-center gap-x-2">
                <img class="w-10" src="@/assets/wats.svg">
                <p>+ 7 (917) 352 20 80</p>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<style scoped></style>
