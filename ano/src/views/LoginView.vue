<script setup lang="ts">
import { onMounted, ref } from "vue";
// TODO: We have duplicate src/assets/logo.svg and src/components/icons/Logo.vue
import Logo from "@/components/icons/Logo.vue";
import axios from "axios";
import { useRouter } from "vue-router";
import { useToast } from "vue-toastification"
import ToastError from "@/components/ToastError.vue"

const url: string = import.meta.env.VITE_ENDPOINT
const router = useRouter();
const toast = useToast();

// Login
const isLoginRequired = ref<number>(0)
const data = {
    email: "",
    password: "",
};

// TODO: hide axios errors in console 
// https://dev.to/zelig880/how-to-catch-the-body-of-an-axios-error-4lk0
onMounted(async () => {
    await axios
        .get(`${url}/api/auth/user`, {
            withCredentials: true,
        })
        // 200 OK
        .then((response) => {
            if (response.data.id != 0) {
                // TODO: push to where the user came from
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
        .post(`${url}/api/auth/login`, data, { withCredentials: true })
        .then((response) => {
            router.push("/admin");
        })
        .catch(function (error) {
            // TODO: maybe make field "email" red for visibility
            // make field normal back on input
            if (404 == error.response.status) {
                toast(StatusNotFound)
            } else if (400 == error.response.status) {
                toast(StatusBadRequest)
            } else {
                toast(StatusInternalServerError)
                console.log(error)
            }
        });
}

// Toast
const StatusNotFound = {
    component: ToastError,

    props: {
        text: "Пользователь не найден",
        description: "Проверьте правильность написания логина",
    }
}
const StatusBadRequest = {
    component: ToastError,

    props: {
        text: "Неверный пароль",
        description: "Проверьте правильность логина или пароля",
    }
}
const StatusInternalServerError = {
    component: ToastError,

    props: {
        text: "Ошибка сервера",
        description: "Консоль может содержать подробную информацию",
    }
}


// TODO: add logout button somewhere
// Logout
// GET /api/auth/logout
</script>


<template>
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
    <body class="h-full bg-white">
        <!--
      This example requires updating your template:
  
      ```
      <html class="h-full bg-white">
      <body class="h-full">
      ```
    -->
        <div class="flex min-h-full flex-1 flex-col justify-center px-6 py-12 lg:px-8" v-if="isLoginRequired">
            <div class="sm:mx-auto sm:w-full sm:max-w-sm">
                <div class="flex justify-center">
                    <Logo></Logo>
                </div>
                <h2 class="mt-10 text-center text-2xl font-bold leading-9 tracking-tight text-gray-900">
                    Sign in to your account
                </h2>
            </div>

            <div class="mt-10 sm:mx-auto sm:w-full sm:max-w-sm">
                <form @submit.prevent="submit" class="space-y-6" action="#" method="POST">
                    <div>
                        <label for="email" class="block text-sm font-medium leading-6 text-gray-900">Email address</label>
                        <div class="mt-2">
                            <input v-model="data.email" id="email" name="email" autocomplete="email"
                                class="block w-full rounded-md border-0 py-1.5 text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 focus:ring-2 focus:ring-inset focus:ring-indigo-600 sm:text-sm sm:leading-6" />
                        </div>
                    </div>

                    <div>
                        <div class="flex items-center justify-between">
                            <label for="password" class="block text-sm font-medium leading-6 text-gray-900">Password</label>
                        </div>
                        <div class="mt-2">
                            <input v-model="data.password" id="password" name="password" type="password"
                                autocomplete="current-password"
                                class="block w-full rounded-md border-0 py-1.5 text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 focus:ring-2 focus:ring-inset focus:ring-indigo-600 sm:text-sm sm:leading-6" />
                        </div>
                    </div>

                    <div>
                        <!-- TODO: add animation on click -->
                        <button type="submit"
                            class="flex w-full justify-center rounded-md bg-indigo-600 px-3 py-1.5 text-sm font-semibold leading-6 text-white shadow-sm hover:bg-indigo-500 focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-indigo-600">
                            Sign in
                        </button>
                    </div>
                </form>
            </div>
        </div>
    </body>
</template>

<style></style>