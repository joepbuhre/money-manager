<template>
    <div class="flex h-screen w-screen">
        <TheMenu />
        <div class="h-screen w-full overflow-auto p-3">
            <router-view />
        </div>
        <Notifications />
    </div>
</template>

<script setup lang="ts">
import { RouterView } from "vue-router";
import Notifications from "@/components/Notifications.vue";
import TheMenu from "@/components/TheMenu.vue";
</script>
