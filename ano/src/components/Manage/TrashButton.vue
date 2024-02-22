<script setup lang="ts">
import axios from 'axios'
import { TrashIcon } from '@heroicons/vue/24/outline'

const emit = defineEmits(['reloadItems'])
const props = defineProps(['id', 'type'])
const url: string = import.meta.env.VITE_ENDPOINT

const handleClick = async () => {
    // TODO: add "Do you really want to delete the item"
    try {
        const response = await axios.delete(`${url}/api/${props.type}/${props.id}`)
            .then((response) => {
                console.log(response)
            })
            .finally(() => {
                emit('reloadItems')
            })
    } catch (error) {
        console.error('Error deleting news: ', error)
    }
}
</script>

<template>
    <!-- We use `click.prevent` to prevent redirection since the TrashButton is typically placed inside a `router-link` component -->
    <TrashIcon @click.prevent="handleClick" class="text-white bg-red-500 w-full h-9 mt-2 lg:w-7 lg:h-7 lg:mt-0 p-1 rounded hover:bg-red-800"></TrashIcon>
</template>

<style scoped></style>