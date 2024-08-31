import {defineStore} from "pinia";

export const useSettingStore = defineStore('setting', {
    state: () => ({
        siteName: "我爱云加速",
        logo: "https://gw.alipayobjects.com/zos/rmsportal/KDpgvguMpGfqaHPjicRK.svg"
    }),
    actions: {
        changeSetting(setting) {
            this.siteName = setting.siteName;
            this.logo = setting.logo;
        },
    },
    getters: {
    },
});