<template></template>

<script setup lang="ts">
import { onMounted } from 'vue'
import axios, { type AxiosResponse } from 'axios'

const url: string = import.meta.env.VITE_ENDPOINT

onMounted(async () => {
    await axios.get(`${url}/api/events`)
        .then((response) => {
            console.log(response)
        })
})
</script>

<style></style>