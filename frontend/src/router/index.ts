import { createRouter, createWebHistory } from "vue-router";
import { useUsers } from "@/stores/user";
import HomeView from "@/views/HomeView.vue";

const router = createRouter({
    history: createWebHistory(),
    routes: [
        {
            path: "/",
            name: "home",
            component: HomeView,
        },
    ],
});

// router.beforeEach((to, from) => {
//     const u = useUsers();

//     if (to.name === "Login") {
//         return;
//     } else {
//         if (u.isLoggedIn) {
//             return;
//         } else {
//             return {
//                 name: "Login",
//             };
//         }
//     }
// });

export default router;
