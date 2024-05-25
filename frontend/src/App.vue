<template>
    <router-view />
    <Notifications />
</template>

<script setup lang="ts">
import { RouterView } from "vue-router";
import Notifications from "@/components/Notifications.vue";
</script>
