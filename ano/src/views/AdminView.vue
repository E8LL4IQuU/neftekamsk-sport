<script setup lang="ts">
import { onMounted } from 'vue'
import axios, { type AxiosResponse } from 'axios'
import { useRouter } from 'vue-router'

const router = useRouter()

const url: string = import.meta.env.VITE_ENDPOINT

let bebra: string = ""

onMounted(async () => {
  await axios
    .get(`${url}/api/user`, {
      withCredentials: true,
      // 200 OK
    })
    .then((response) => {
      if (response.data.id != 0) {
        bebra = response.data.name
      }
    })
    .catch((error) => {
      if (401 == error.response.status) {
        router.push("/login")
      }
    });
});

</script>

<template>
<h2 class="text-black text-3xl">
  Admin panel
</h2>
<h2>
  Hello, {{ bebra }}
</h2>
</template>

<style>

</style>
