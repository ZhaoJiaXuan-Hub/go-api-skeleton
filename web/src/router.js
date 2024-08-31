import { createRouter, createWebHistory } from 'vue-router';
import Home from './views/Home.vue';
import Login from './views/Login.vue';
import Dashboard from './views/Dashboard.vue';
import {useUserStore} from "./stores/userStore.js";
import Layout from "./components/layout/Layout.vue";
import SimpleLayout from "./components/layout/SimpleLayout.vue";
import cache from "./utils/cache.js";
import {TOKEN_KEY} from "./enums/cacheEnums.js";
import {SUCCESS} from "./enums/requestCodeEnum.js";
import {message} from "ant-design-vue";

const routes = [
    {
        path: '/',
        name: 'Home',
        component: Home,
        meta: {
            layout: SimpleLayout,
            requiresAuth: false
        }
    },
    {
        path: '/dashboard',
        name: 'Dashboard',
        component: Dashboard,
        meta: {
            layout: Layout,
            requiresAuth: true, // 表示这个路由需要认证
        },
    },
    {
        path: '/login',
        name: 'Login',
        component: Login,
        meta: {
            layout: SimpleLayout,
            requiresAuth: false
        }
    },
];

const router = createRouter({
    history: createWebHistory(),
    routes,
});


router.beforeEach((to, from, next) => {
    // 在这里进行权限检查、加载指示器等操作
    const state = useUserStore();
    state.token = cache.get(TOKEN_KEY);

    // 优化 token 为空或不存在的判断方式
    const isTokenEmpty =!state.token || state.token.trim() === '';

    // 当路由需要登录且 token 为空
    if (to.meta.requiresAuth && isTokenEmpty) {
        console.log('/login');
        next('/login');
        return;
    }

    if (to.meta.requiresAuth) {
        state.getUserInfo()
            .then(res => {
                if (res.code === SUCCESS) {
                    next();
                }
            })
            .catch(error => {
                message.error(error.message);
                next();
            });
        return;
    }

    next();
});

// 重置路由
export function resetRouter() {
}

export default router;