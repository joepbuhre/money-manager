import { useUsers } from "@/stores/user";
import axios, { InternalAxiosRequestConfig } from "axios";

export const getUrl = (path?: string, sparams?: URLSearchParams) => {
    const url = new URL(window.location.href);
    url.pathname = "";
    url.hash = "";
    if (path) {
        url.pathname = path;
    }
    if (sparams) {
        url.search = sparams.toString();
    }
    return url;
};

export const getRelativeDateTime = (dt: Date) => {
    // in miliseconds
    let units = {
        year: 24 * 60 * 60 * 1000 * 365,
        month: (24 * 60 * 60 * 1000 * 365) / 12,
        day: 24 * 60 * 60 * 1000,
        hour: 60 * 60 * 1000,
        minute: 60 * 1000,
        second: 1000,
    };

    let rtf = new Intl.RelativeTimeFormat("en", { numeric: "auto" });

    let elapsed = dt.getTime() - new Date().getTime();

    // "Math.abs" accounts for both "past" & "future" scenarios
    let u: keyof typeof units;
    for (u in units) {
        if (Math.abs(elapsed) > units[u] || u == "second")
            return rtf.format(Math.round(elapsed / units[u]), u);
    }
};

export const api = axios.create({
    baseURL: "/api/",
});

api.interceptors.request.use((config: InternalAxiosRequestConfig) => {
    // debugger;
    if (!("api-key" in config.headers)) {
        const user = useUsers();
        config.headers["api-key"] = user.token;
    }
    return config;
});
