<template>hi</template>
<script setup lang="ts">
import { computed } from "vue";

import { useWindowSize } from "@vueuse/core";

const { width, height } = useWindowSize();
</script>
