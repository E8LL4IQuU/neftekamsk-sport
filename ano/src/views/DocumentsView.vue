<script setup lang="ts">
</script>

<template>
    <body class="bg-white text-black mx-auto container text-xl pt-4">
        <h2>ИНН 0264079701</h2>
        <h2>ОГРН 1210200007401</h2>
        <div class="flex">
            <object class="w-96 h-[600px]" :data="`/inn.pdf`" type="application/pdf">
                <a href="/inn.pdf">ИНН</a>
            </object>
            <object class="w-96 h-[600px]" :data="`/ogrn.pdf`" type="application/pdf">
                <a href="/ogrn.pdf">ИНН</a>
            </object>
            <!-- FIXME: dude it's like an image of terms but no terms actually💀💀 -->
            <img class="w-96" src="@/assets/terms.jpg" alt="Устав" />
        </div>
    </body>
</template>

<style></style>