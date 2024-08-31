import {defineStore} from "pinia";

export const useLayoutStore = defineStore('layout', {
    state: () => ({
        collapsed: false,
    }),
    actions: {
        changeCollapsed() {
            this.collapsed = !this.collapsed;
        },
    },
    getters: {
    },
});