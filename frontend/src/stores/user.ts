import { ref, computed, watch, getCurrentInstance, onMounted } from "vue";
import { defineStore } from "pinia";
import type { User } from "@/types/go/database";
import { api } from "@/utils/helpers";
import router from "@/router";

export const useUsers = defineStore(
    "user",
    () => {
        const user = ref<User | null>(null);
        const token = ref<string>("");

        const isLoggedIn = computed(
            () => user.value !== null && token.value !== "",
        );

        if (getCurrentInstance()) {
            onMounted(() => {
                console.log("store is mounted");
            });
        }

        watch(user, (newVal) => {
            if (newVal !== null) {
            }
        });

        // Functions
        const logout = () => {
            // api.post("/users/logout")
            user.value = null;
            token.value = "";

            router.push({ name: "Login" });
        };

        return { user, token, isLoggedIn, logout };
    },
    { persist: true },
);
