import { defineStore } from 'pinia';
import {getUserInfo, login} from "../api/login/index.js";
import cache from "../utils/cache.js";
import {TOKEN_KEY} from "../enums/cacheEnums.js";

export const useUserStore = defineStore('user', {
    state: () => ({
        token: '',
        userInfo: {},
    }),
    actions: {
        resetState() {
            this.token = ''
            this.userInfo = {}
        },
        login(username, password) {
            const credentials = { username: username, password: password };
            return new Promise((resolve, reject) => {
                login(credentials)
                    .then((data) => {
                        this.token = data.data.token
                        cache.set(TOKEN_KEY, data.data.token)
                        resolve(data)
                    })
                    .catch((error) => {
                        reject(error)
                    })
            })
        },
        getUserInfo() {
            return new Promise((resolve, reject) => {
                getUserInfo()
                    .then(res => {
                        this.userInfo = res.data;
                        resolve(res)
                    })
                    .catch((error) => {
                        reject(error)
                    })
            })
        }
    },
    getters: {
    },
});