import request from '../../utils/request.js'

export const login = (credentials) => request.post('/account/login', credentials);

export const reg = (credentials) => request.post('/account/reg', credentials);

export const getUserInfo = () => request.get('/account/getDetail')