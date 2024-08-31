import axios from 'axios';
import {FAIL, LOGIN_FAILURE, SUCCESS} from "../enums/requestCodeEnum.js";
import router from "../router.js";
import * as PageEnum from "../enums/pageEnums.js";
import cache from "./cache.js";
import {TOKEN_KEY} from "../enums/cacheEnums.js";
import {message} from "ant-design-vue";

const request = axios.create({
    baseURL: 'http://127.0.0.1:9501/api', // 设置基础 URL
});

// 请求拦截器
request.interceptors.request.use(
    (config) => {
        // 可以在这里添加请求头、请求参数处理等逻辑
        const token = cache.get(TOKEN_KEY);
        if (token) {
            config.headers.Authorization = `Bearer ${token}`;
        }
        return config;
    },
    (error) => {
        return Promise.reject(error);
    }
);

// 响应拦截器
request.interceptors.response.use(
    (response) => {
        switch (response.data.code) {
            case SUCCESS:
                return response.data;
            case FAIL:
                return Promise.reject(response.data);
            case LOGIN_FAILURE:
                message.error(response.data.message);
                router.push(PageEnum.LOGIN)
                return Promise.reject()
            default:
                return response.data
        }
    },
    (error) => {
        // 可以在这里处理响应错误，如统一的错误提示等
        return Promise.reject(error);
    }
);

export default request;