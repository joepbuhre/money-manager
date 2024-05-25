import { defineStore } from "pinia";

export interface NotObject {
    id: string;
    text: string;
    clickFnc?: (obj: NotObject) => void;
}

interface NotificationsState {
    notifications: {
        [key: string]: NotObject;
    };
}

export const useNotifications = defineStore("Notifications", {
    state: (): NotificationsState => {
        return {
            notifications: {},
        };
    },
    getters: {
        get(state) {
            return state.notifications;
        },
    },
    actions: {
        add(
            text: string,
            type = "info",
            timeout = 10000,
            clickFnc?: NotObject["clickFnc"],
        ) {
            const randId = `ID_${Math.round(Math.random() * 1000).toString()}`;
            const obj: NotObject = {
                id: randId,
                text: text,
                clickFnc: clickFnc,
            };
            this.notifications[randId] = obj;

            setTimeout(() => {
                this.delete(randId);
            }, timeout);
        },
        delete(randId: string) {
            try {
                delete this.notifications[randId];
            } catch (error) {
                console.log(error);
            }
        },
    },
});
