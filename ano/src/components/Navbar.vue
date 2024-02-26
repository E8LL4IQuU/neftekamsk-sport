<script setup lang="ts">
import Logo from "./icons/Logo.vue";

import { onMounted, ref } from "vue";
import { RouterLink } from "vue-router";
import {
  Dialog,
  DialogPanel,
  Disclosure,
  DisclosureButton,
  DisclosurePanel,
  Popover,
  PopoverButton,
  PopoverGroup,
  PopoverPanel,
} from "@headlessui/vue";
import { Bars3Icon, XMarkIcon } from "@heroicons/vue/24/outline";
import { ChevronDownIcon } from "@heroicons/vue/20/solid";
import axios from "axios";
import { type IRLEvent } from "@/types/apiTypes";

const props = defineProps<{
  navbarItems: string[][];
}>();
const url: string = import.meta.env.VITE_ENDPOINT;
const IRLEvents = ref<IRLEvent[]>([]);
const mobileMenuOpen = ref<boolean>(false);

const fetchIRLEvents = async (): Promise<void> => {
  try {
    const response = await axios.get<IRLEvent[]>(`${url}/api/events?limit=6`);
    IRLEvents.value = response.data;
  } catch (error) {
    console.log("Error fetching IRLEvents: ", error);
  }
};

const formatTimestamp = (timestamp: string): string => {
  const date = new Date(timestamp);

  const options: Intl.DateTimeFormatOptions = {
    month: '2-digit',
    day: 'numeric',
    hour: 'numeric',
    minute: 'numeric'

  };

  return date.toLocaleDateString("ru-RU", options);
}

onMounted(() => {
  fetchIRLEvents();
});
</script>

<template>
  <header class="bg-white">
    <nav class="container mx-auto flex items-center justify-between p-6 sm:p-0 lg:px-8" aria-label="Global">
      <div class="flex lg:flex-1">
        <router-link :to="'/'" class="-m-1.5 p-1.5">
          <span class="sr-only">АНО "Нефтекамск спортивный"</span>
          <Logo class="sm:p-3"></Logo>
        </router-link>
      </div>
      <div class="flex lg:hidden">
        <button type="button" class="-m-2.5 inline-flex items-center justify-center rounded-md p-2.5 text-gray-700"
          @click="mobileMenuOpen = true">
          <span class="sr-only">Открыть основное меню</span>
          <Bars3Icon class="h-6 w-6" aria-hidden="true" />
        </button>
      </div>
      <PopoverGroup class="hidden lg:flex lg:gap-x-12">
        <Popover class="relative">
          <PopoverButton class="flex items-center gap-x-1 text-sm font-semibold leading-6 text-gray-900">
            Мероприятия
            <ChevronDownIcon class="h-5 w-5 flex-none text-gray-400" aria-hidden="true" />
          </PopoverButton>

          <transition enter-active-class="transition ease-out duration-200" enter-from-class="opacity-0 translate-y-1"
            enter-to-class="opacity-100 translate-y-0" leave-active-class="transition ease-in duration-150"
            leave-from-class="opacity-100 translate-y-0" leave-to-class="opacity-0 translate-y-1">
            <PopoverPanel
              class="absolute -left-8 top-full z-10 mt-3 w-screen max-w-md overflow-hidden rounded-3xl bg-white shadow-lg ring-1 ring-gray-900/5">
              <div class="p-4">
                <div v-for="item in IRLEvents" :key="item.Title"
                  class="group relative flex items-center gap-x-6 rounded-lg p-4 text-sm leading-6 hover:bg-gray-50">
                  <div class="flex-auto">
                    <router-link :to="`/events/${item.ID}`" class="block text-gray-900">
                      <h2 class="font-semibold">{{ item.Title }}</h2>
                      <h3>{{ formatTimestamp(item.Date) }}</h3>
                      <span class="absolute inset-0" />
                    </router-link>
                  </div>
                </div>
                <router-link to="/events">
                  <div
                    class="group relative flex justify-center gap-x-6 rounded-lg p-4 text-sm leading-6 hover:bg-gray-50 text-black">
                    Все мероприятия
                  </div>
                </router-link>
              </div>
            </PopoverPanel>
          </transition>
        </Popover>

        <div v-for="item in props.navbarItems">
          <router-link :to="item[1]" class="text-sm font-semibold leading-6 text-gray-900">{{ item[0] }}</router-link>
        </div>
      </PopoverGroup>
    </nav>
    <Dialog as="div" class="lg:hidden" @close="mobileMenuOpen = false" :open="mobileMenuOpen">
      <div class="fixed inset-0 z-10" />
      <DialogPanel
        class="fixed inset-y-0 right-0 z-10 w-full overflow-y-auto bg-white px-6 py-6 sm:max-w-sm sm:ring-1 sm:ring-gray-900/10">
        <div class="flex items-center justify-between">
          <router-link to="/" class="-m-1.5 p-1.5">
            <span class="sr-only">АНО "Нефтекамск Спортивный"</span>
            <Logo></Logo>
          </router-link>
          <button type="button" class="-m-2.5 rounded-md p-2.5 text-gray-700" @click="mobileMenuOpen = false">
            <span class="sr-only">Закрыть меню</span>
            <XMarkIcon class="h-6 w-6" aria-hidden="true" />
          </button>
        </div>
        <div class="mt-6 flow-root">
          <div class="-my-6 divide-y divide-gray-500/10">
            <div class="space-y-2 py-6">

              <Disclosure as="div" class="-mx-3" v-slot="{ open }">
                <DisclosureButton
                  class="flex w-full items-center justify-between rounded-lg py-2 pl-3 pr-3.5 text-base font-semibold leading-7 text-gray-900 hover:bg-gray-50">
                  Мероприятия
                  <ChevronDownIcon :class="[open ? 'rotate-180' : '', 'h-5 w-5 flex-none']" aria-hidden="true" />
                </DisclosureButton>
                <DisclosurePanel class="mt-2 space-y-2">
                  <router-link v-for="item in [...IRLEvents]" :key="item.Title" :to="`/events/${item.ID}`"
                    @click="mobileMenuOpen = false">
                    <DisclosureButton as="div"
                      class="block rounded-lg py-0.5 my-1 pl-6 pr-3 text-sm leading-7 text-gray-900 bg-gray-100 hover:bg-gray-300">
                      <h2 class="font-semibold">{{ item.Title }}</h2>
                      <h3>{{ formatTimestamp(item.Date) }}</h3>
                    </DisclosureButton>
                  </router-link>
                  <router-link to="/events" @click="mobileMenuOpen = false">
                    <DisclosureButton as="div" 
                      class="block text-center rounded-lg py-2 pl-6 pr-3 text-sm leading-7 bg-gray-200 hover:bg-gray-300">
                      Все мероприятия
                    </DisclosureButton>
                  </router-link>
                </DisclosurePanel>
              </Disclosure>

              <div v-for="item in props.navbarItems">
                <router-link :to="item[1]" class="text-sm font-semibold leading-6 text-gray-900"
                  @click="mobileMenuOpen = false">{{ item[0] }}</router-link>
              </div>
            </div>
          </div>
        </div>
      </DialogPanel>
    </Dialog>
  </header>
</template>

<style></style>
