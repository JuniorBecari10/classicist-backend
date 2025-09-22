<script setup lang="ts">
    import { ref } from "vue";

    let isHello = ref(true);
    const flip = () => isHello.value ^= true;
</script>

<template>
    <button class="
        bg-sky-500 hover:bg-white
        text-white hover:text-black
        rounded-full py-2 px-4 m-4 cursor-pointer
        duration-200 ease-in-out"
        @click="flip">
        Change
    </button>

    <h1 class="text-3xl font-fredoka font-bold"
        v-if="isHello">Hello</h1>
    <h1 class="text-3xl font-fredoka font-bold"
        v-else>World</h1>
</template>
