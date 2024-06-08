import { RouteRecordRaw } from "vue-router";
import HomeView from "../views/HomeView.vue";
import AccountSettingsView from "../views/AccountSettingsView.vue";
import {
    ArrowLeftRightIcon,
    WalletCardsIcon,
    WalletMinimalIcon,
} from "lucide-vue-next";
import TransactionsView from "@/views/TransactionsView.vue";
import BudgetView from "@/views/BudgetView.vue";
import { Component } from "vue";

interface RouteMeta {
    excludeInMenu?: boolean;
    lucideIcon?: Component;
    [key: string]: any;
}

type RouteRecordCustom = Omit<RouteRecordRaw, "meta"> & {
    meta?: RouteMeta;
};

const routes: RouteRecordCustom[] = [
    {
        path: "/",
        name: "home",
        component: HomeView,
        meta: {
            excludeInMenu: true,
        },
    },
    {
        path: "/accounts",
        name: "Accounts",
        component: AccountSettingsView,
        meta: {
            lucideIcon: WalletCardsIcon,
        },
    },
    {
        path: "/budgets",
        name: "Budgets",
        component: BudgetView,
        meta: {
            lucideIcon: WalletMinimalIcon,
        },
    },
    {
        path: "/transactions",
        name: "Transactions",
        component: TransactionsView,
        meta: {
            lucideIcon: ArrowLeftRightIcon,
        },
    },
];

export default routes;
