<script setup lang="ts">
import { ref, onMounted, watch } from "vue";
import { useRoute } from "vue-router";
import axios from "axios";
import { type IRLEvent } from "@/types/apiTypes";

const url: string = import.meta.env.VITE_ENDPOINT;
const IRLEvent = ref<IRLEvent>();
const route = useRoute();
const isRegistrationPopupOpen = ref<boolean>(false);
const FirstName = ref<string>('');
const LastName = ref<string>('');
const PhoneNumber = ref<string>('');
const Email = ref<string>('');
let EventID: number = Number(route.params.id);

const openRegistrationPopup = (): void => {
  isRegistrationPopupOpen.value = true;
};

const closeRegistrationPopup = (): void => {
  isRegistrationPopupOpen.value = false;
};

const submitRegistrationForm = async (): Promise<void> => {
  try {
    const signupObject = {
      FirstName: FirstName.value,
      LastName: LastName.value,
      PhoneNumber: PhoneNumber.value,
      Email: Email.value,
    }

    const requestConfig = { method: 'post', url: `${url}/api/signups?id=${EventID}`, data: signupObject }

    await axios(requestConfig).then(() => {
      closeRegistrationPopup();
    })

  } catch (error) {
    console.error('Error signing up for event: ', error);
  }
};

const fetchIRLEvent = async (): Promise<void> => {
  try {
    const response = await axios.get<IRLEvent>(`${url}/api/events/${EventID}`);
    IRLEvent.value = response.data;
  } catch (error) {
    console.error("Error fetching IRLevents:", error);
  }
};

onMounted(() => {
  fetchIRLEvent();
});

watch(() => route.params.id, (newID) => {
  EventID = Number(newID)
  fetchIRLEvent();
})
</script>

<template>
  <body class="px-1 sm:px-0 sm:container mx-auto pb-20 bg-white">
    <div class="lg:flex gap-x-12" v-if="IRLEvent">
      <img class="sm:w-full md:w-1/2" :src="`${url}/uploads/${IRLEvent.ImagePath}`" alt="">
      <div class="mobile:text-center">
        <h1 class="text-black text-3xl font-bold pb-3">{{ IRLEvent.Title }}</h1>
        <p class="text-black max-w-7xl">{{ IRLEvent.Description }}</p>
        <div>
          <button class="px-2 py-2 mt-3 rounded bg-red-400 text-white hover:bg-red-500" @click="openRegistrationPopup">Записаться</button>

          <!-- Registration Popup -->
          <!-- FIXME: Too small bottom padding -->
          <div v-if="isRegistrationPopupOpen"
            class="fixed inset-0 bg-gray-800 bg-opacity-75 flex items-center justify-center"
            @click.self="closeRegistrationPopup">
            <div @click.stop class="relative bg-white rounded-[20px] w-full xl:w-1/2 flex flex-col gap-y-3 items-center text-center p-[40px] sm:px-0 mac:p-[10px]">
              <button @click="closeRegistrationPopup"
                class="text-7xl absolute top-0 right-2 pt-5 pr-4 text-gray-600 hover:text-[#F24E1E] mac:pt-0">&times;</button>
              <h2 class="text-black text-[40px] font-bold mobile:pr-5">ЗАПИСЬ</h2>
              <!-- Registration form content -->
              <form @submit.prevent="submitRegistrationForm" class="w-full md:w-8/12 pb-6">
                <div class="">
                  <input required v-model="FirstName" type="text" id="FirstName"
                         class="w-full text-black mt-[30px] px-[15px] py-[15px] border-none rounded-[20px] bg-stone-200 placeholder-black placeholder-shown:font-bold text-[24px] mac:mt-[10px] mobile:text-[16px]"
                    placeholder="Имя">
                </div>

                <div class="">
                  <input required v-model="LastName" type="text" id="LastName"
                         class="w-full text-black mt-[30px] px-[15px] py-[15px] border-none rounded-[20px] bg-stone-200 placeholder-black placeholder-shown:font-bold text-[24px] mac:mt-[10px] mobile:text-[16px]"
                    placeholder="Фамилия">
                </div>

                <div class="">
                  <input required v-model="PhoneNumber" type="text" id="PhoneNumber"
                         class="w-full text-black mt-[30px] px-[15px] py-[15px] border-none rounded-[20px] bg-stone-200 placeholder-black placeholder-shown:font-bold text-[24px] mac:mt-[10px] mobile:text-[16px]"
                    placeholder="Номер телефона">
                </div>

                <div class="">
                  <input v-model="Email" type="text" id="Email"
                    class="w-full text-black mt-[30px] px-[15px] py-[15px] border-none rounded-[20px] bg-stone-200 placeholder-black placeholder-shown:font-bold text-[24px] mac:mt-[10px] mobile:text-[16px]"
                    placeholder="Почта(необязательно)">
                </div>
                <button type="submit" class="bg-red-400 font-bold text-white text-[24px] px-[10px] py-[10px] mt-[40px] rounded-[20px] w-full hover:bg-red-500 hover:text-gray-100 mac:mt-[10px] mobile:text-[16px]">Записаться</button>
              </form>
              <hr class="w-10/12 border-t-1 border-[#CCC;]">
              <p class="font-bold text-[24px] text-[#797979] pb-3 mac:text-[20px] mac:pb-[10px] mobile:12px">Или вы можете записаться
                <br>
                через whatsapp</p>
              <div class="flex items-center gap-2 mobile:pb-2">
                <img class="w-[60px] mobile:w-[35px]" src="@/assets/wats.svg">
                <p class="text-black text-[32px] font-light mac:text-[25px] mobile:text-[20px]">+ 7 (917) 352 20 80</p>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </body>
</template>

<style scoped></style>
