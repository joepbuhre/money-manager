<template>
    <div
        style="z-index: 99999"
        class="fixed top-0 w-screen"
        id="toastPlacement"
    >
        <div
            id="toast-default"
            class="m-auto mb-3 flex max-w-xs items-center rounded-lg bg-white p-4 text-gray-500 shadow"
            role="alert"
            v-for="not in notifications.get"
        >
            <button
                @click="handleClick(not)"
                class="inline-flex h-8 w-8 flex-shrink-0 items-center justify-center rounded-lg bg-blue-100 text-blue-500 dark:bg-blue-800 dark:text-blue-200"
            >
                <FlameIcon />
                <span class="sr-only">Fire icon</span>
            </button>
            <div class="ms-3 text-sm font-normal">{{ not.text }}</div>
            <button
                type="button"
                class="-mx-1.5 -my-1.5 ms-auto inline-flex h-8 w-8 items-center justify-center rounded-lg bg-white p-1.5 text-gray-400 hover:bg-gray-100 hover:text-gray-900 focus:ring-2 focus:ring-gray-300 dark:bg-gray-800 dark:text-gray-500 dark:hover:bg-gray-700 dark:hover:text-white"
                data-dismiss-target="#toast-default"
                aria-label="Close"
                @click="notifications.delete(not.id)"
            >
                <span class="sr-only">Close</span>
                <XIcon />
            </button>
        </div>
    </div>
</template>

<script setup lang="ts">
import { NotObject, useNotifications } from "@/stores/notifications";
import { FlameIcon, XIcon } from "lucide-vue-next";

const notifications = useNotifications();

const handleClick = (not: NotObject) => {
    if (not.clickFnc && typeof not.clickFnc === "function") {
        not.clickFnc(not);
    }
};
</script>
