import { TOKEN_KEY } from '../enums/cacheEnums'
import { resetRouter } from '../router.js'
import useUserStore from '../stores/userStore.js'

import cache from './cache'

export function getToken() {
    return cache.get(TOKEN_KEY)
}

export function clearAuthInfo() {
    const userStore = useUserStore()
    userStore.resetState()
    cache.remove(TOKEN_KEY)
    resetRouter()
}