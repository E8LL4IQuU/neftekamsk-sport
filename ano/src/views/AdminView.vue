<script setup lang="ts">
import { onMounted } from 'vue'
import axios, { type AxiosResponse } from 'axios'
import { useRouter } from 'vue-router'

const router = useRouter()

const url: string = import.meta.env.VITE_ENDPOINT

let bebra: string = ""


</script>

<template>
<h2 class="text-black text-3xl">
  Admin panel
</h2>
</template>

<style>

</style>
