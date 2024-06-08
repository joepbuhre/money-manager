import { RouteRecordRaw, createRouter, createWebHistory } from "vue-router";
import routes from "./routes";

const router = createRouter({
    history: createWebHistory(),
    routes: routes,
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
