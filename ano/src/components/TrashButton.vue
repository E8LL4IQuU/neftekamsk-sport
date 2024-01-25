<script setup lang="ts">
import axios from 'axios'
import { TrashIcon } from '@heroicons/vue/24/outline'

const emit = defineEmits(['reloadNews'])
const props = defineProps(['id', 'type'])
const url: string = import.meta.env.VITE_ENDPOINT

const handleClick = async () => {
    try {
        const response = await axios.delete(`${url}/api/${props.type}/${props.id}`)
        .then((response) => {
            console.log(response)
        })
        .finally(() => {
            emit('reloadNews')
        })
    } catch (error) {
        console.error('Error deleting news: ', error)
    }
}
</script>

<template>
<TrashIcon @click.stop="handleClick" class="text-white bg-red-500 h-7 p-1 rounded"></TrashIcon>
</template>

<style scoped>

</style>