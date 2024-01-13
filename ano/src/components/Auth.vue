<!--
  This example requires some changes to your config:
  
  ```
  // tailwind.config.js
  module.exports = {
    // ...
    plugins: [
      // ...
      require('@tailwindcss/forms'),
    ],
  }
  ```
-->
<template>
  <!--
      This example requires updating your template:
  
      ```
      <html class="h-full bg-white">
      <body class="h-full">
      ```
    -->
  <div
    class="flex min-h-full flex-1 flex-col justify-center px-6 py-12 lg:px-8"
    v-if="isLoginRequired"
  >
    <div class="sm:mx-auto sm:w-full sm:max-w-sm">
      <div class="flex justify-center">
        <Logo></Logo>
      </div>
      <h2
        class="mt-10 text-center text-2xl font-bold leading-9 tracking-tight text-gray-900"
      >
        Sign in to your account
      </h2>
    </div>

    <div class="mt-10 sm:mx-auto sm:w-full sm:max-w-sm">
      <form @submit.prevent="submit" class="space-y-6" action="#" method="POST">
        <div>
          <label
            for="email"
            class="block text-sm font-medium leading-6 text-gray-900"
            >Email address</label
          >
          <div class="mt-2">
            <input
              v-model="data.email"
              id="email"
              name="email"
              autocomplete="email"
              class="block w-full rounded-md border-0 py-1.5 text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 focus:ring-2 focus:ring-inset focus:ring-indigo-600 sm:text-sm sm:leading-6"
            />
          </div>
        </div>

        <div>
          <div class="flex items-center justify-between">
            <label
              for="password"
              class="block text-sm font-medium leading-6 text-gray-900"
              >Password</label
            >
          </div>
          <div class="mt-2">
            <input
              v-model="data.password"
              id="password"
              name="password"
              type="password"
              autocomplete="current-password"
              class="block w-full rounded-md border-0 py-1.5 text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 focus:ring-2 focus:ring-inset focus:ring-indigo-600 sm:text-sm sm:leading-6"
            />
          </div>
        </div>

        <div>
          <button
            type="submit"
            class="flex w-full justify-center rounded-md bg-indigo-600 px-3 py-1.5 text-sm font-semibold leading-6 text-white shadow-sm hover:bg-indigo-500 focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-indigo-600"
          >
            Sign in
          </button>
        </div>
      </form>
    </div>
  </div>
</template>

<script setup lang="ts">
import { onMounted, ref } from "vue";
import Logo from "./icons/Logo.vue";
import axios from "axios";
import { useRouter } from "vue-router";

const router = useRouter();

const isLoginRequired = ref<number>(0)

const data = {
  email: "",
  password: "",
};

onMounted(async () => {
  await axios
    .get('/api/auth/user', {
      withCredentials: true,
    })
    // 200 OK
    .then((response) => {
      if (response.data.id != 0) {
        router.push("/admin");
      }
    })
    .catch((error) => {
      if (401 == error.response.status) {
        isLoginRequired.value = 1
      }
    });
});

async function submit() {

  await axios
    .post('/api/auth/login', data, { withCredentials: true })
    .then((response) => {
      router.push("/admin");
    })
    .catch(function (error) {
      // TODO: add visual presentation that an action is required
      if (404 == error.response.status) {
        console.log("User not found");
      } else if (400 == error.response.status) {
        console.log("Incorrect password");
      }
    });
}

// Register
// POST /api/auth/register
// {
//     username:   string
//     email:      string
//     password:   string
// }

// Login
// POST /api/auth/login
// {
//     email:      string
//     password:   string
// }

// Logout
// POST /api/auth/logout

// Routes to be implemented
// /api/auth/user returns user json,
// /api/auth/healthcheck return if you are still logged in
</script>
