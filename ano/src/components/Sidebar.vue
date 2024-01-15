<script setup lang="ts">
import { RouterLink } from 'vue-router'

</script>

<template>
<h1>Управление</h1>
<ul>
    <li>
        <router-link :to="'/events'">Мероприятия</router-link>
    </li>
</ul>
</template>

<style scoped>

</style>