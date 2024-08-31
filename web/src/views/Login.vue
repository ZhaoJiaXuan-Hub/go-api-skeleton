<script setup>
import {reactive, ref} from 'vue';
import {LockOutlined,UserOutlined,MobileOutlined,MailOutlined,SafetyCertificateOutlined,SmileOutlined} from '@ant-design/icons-vue';
import cache from "../utils/cache.js";
import {useUserStore} from "../stores/userStore.js";
import {useRoute, useRouter} from "vue-router";
import {ACCOUNT_KEY} from "../enums/cacheEnums.js";
import * as PageEnum from "../enums/pageEnums.js";
import {message} from "ant-design-vue";
import {SUCCESS} from "../enums/requestCodeEnum.js";
import {reg} from "../api/login/index.js";
import {useSettingStore} from "../stores/settingStore.js";
const data = reactive(['用户登陆', '用户注册']);
const value = ref(data[0]);
const route = useRoute()
const router = useRouter()
const setting = useSettingStore();
const loginFormState = reactive({
  username: '',
  password: '',
  remember: false,
});

const regFormState = reactive({
  username: 'admin',
  email: 'ocink@qq.com',
  mobile: '13115204080',
  captcha: '666666',
  password: '123456',
  confirm_password: '123456',
  agreement: false,
});

const loginRemember = cache.get(ACCOUNT_KEY)
if (loginRemember){
  loginFormState.username = loginRemember.username
  loginFormState.password = loginRemember.password
  loginFormState.remember = loginRemember.remember
}

const loading = ref(false)
const visibleCaptcha = ref(false)
const captchaBtnText = ref("获取验证码")
const captchaTime = ref(0)
// 设定倒计时的总时长（以秒为单位）
const totalSeconds = 60;

const sendCaptcha = () => {
  if(regFormState.mobile === '') {
    message.warn("请填写手机号")
    return
  }
  if(captchaTime.value > 0){
    message.warn("发送验证码频繁")
    return
  }
  captchaTime.value = totalSeconds
  visibleCaptcha.value = true
  message.success("发送验证码成功")
  const timer = setInterval(() => {
    if (captchaTime.value > 0) {
      captchaTime.value--;
      captchaBtnText.value = captchaTime.value + "秒"
    } else {
      captchaBtnText.value = "获取验证码"
      visibleCaptcha.value = false
      clearInterval(timer);
    }
  }, 1000);
}

const onLoginFinish = values => {
  handleLogin()
};

const onRegFinish = values => {
  handleReg()
};

const handleReg = async () => {
  loading.value = true
  if(!regFormState.agreement) {
    message.warn("请勾选用户协议")
    loading.value = false
    return
  }
  await reg(regFormState).then(res => {
    loading.value = false
    if(res.code === SUCCESS){
      value.value = data[0]
      message.success("注册成功")
      return
    }
    message.error(res.message)
  }).catch(err => {
    loading.value = false
    console.log('注册失败:', err)
    message.error(err.message);
  });
}

const handleLogin = async () => {
  loading.value = true
  const state = useUserStore();
  if(!loginFormState.remember) {
    cache.remove(ACCOUNT_KEY)
  }
  await state.login(loginFormState.username, loginFormState.password).then(res=>{
    loading.value = false
    if(res.code === SUCCESS){
      message.success("登陆成功")
      // 记住账号，缓存
      cache.set(ACCOUNT_KEY, {
        remember: loginFormState.remember,
        username: loginFormState.remember ? loginFormState.username : '',
        password: loginFormState.remember ? loginFormState.password : ''
      })
      const {
        query: { redirect }
      } = route
      const path = typeof redirect === 'string' ? redirect : PageEnum.INDEX
      router.push(path)
      return
    }
    message.error(res.message);
  }).catch(err=>{
    loading.value = false
    console.log('登陆失败:', err)
    message.error(err.message);
  })
}
</script>

<template>
<div class="login-content">
  <div class="login-card">
    <div class="logo">
      <div class="logo-text">
        <span>{{setting.siteName}}</span>
        <div class="desc">
          用户控制端
        </div>
      </div>
      <SmileOutlined />
    </div>
  <a-segmented v-model:value="value" block :options="data" />
  <template v-if="value === data[0]">
    <a-form
        :model="loginFormState"
        name="normal_login"
        class="login-form"
        @finish="onLoginFinish"
    >
      <a-form-item
          name="username"
          :rules="[{ required: true, message: '请输入用户名' }]"
      >
        <a-input v-model:value="loginFormState.username" placeholder="请输入用户名">
          <template #prefix>
            <UserOutlined />
          </template>
        </a-input>
      </a-form-item>

      <a-form-item
          name="password"
          :rules="[{ required: true, message: '请输入密码' }]"
      >
        <a-input-password v-model:value="loginFormState.password" placeholder="请输入密码">
          <template #prefix>
            <LockOutlined />
          </template>
        </a-input-password>
      </a-form-item>

      <a-form-item>
        <a-form-item name="remember" no-style>
          <a-checkbox v-model:checked="loginFormState.remember">记住密码</a-checkbox>
        </a-form-item>
        <a class="login-form-forgot" href="">找回密码</a>
      </a-form-item>

      <a-form-item>
        <a-button :loading="loading" type="primary" html-type="submit" class="login-form-button">
          立即登陆
        </a-button>
      </a-form-item>
    </a-form>
  </template>
  <template v-if="value === data[1]">
    <a-form
        :model="regFormState"
        name="normal_reg"
        class="login-form"
        @finish="onRegFinish"
    >
      <a-form-item
          name="username"
          :rules="[{ required: true, message: '请输入用户名' }]"
      >
        <a-input v-model:value="regFormState.username" placeholder="请输入用户名">
          <template #prefix>
            <UserOutlined />
          </template>
        </a-input>
      </a-form-item>

      <a-form-item
          name="email"
          :rules="[{ required: true, message: '请输入邮箱账号' }]"
      >
        <a-input v-model:value="regFormState.email" placeholder="请输入邮箱账号">
          <template #prefix>
            <MailOutlined />
          </template>
        </a-input>
      </a-form-item>

      <a-form-item
          name="mobile"
          :rules="[{ required: true, message: '请输入手机号' }]"
      >
        <a-space>
          <a-input v-model:value="regFormState.mobile" placeholder="请输入手机号">
            <template #prefix>
              <MobileOutlined />
            </template>
          </a-input>
          <a-button @click="sendCaptcha" class="login-form-captcha-button" :disabled="visibleCaptcha">{{ captchaBtnText }}</a-button>
        </a-space>
      </a-form-item>

      <a-form-item
          name="captcha"
          :rules="[{ required: true, message: '请输入短信验证码' }]"
      >
        <a-input v-model:value="regFormState.captcha" placeholder="请输入短信验证码">
          <template #prefix>
            <SafetyCertificateOutlined />
          </template>
        </a-input>
      </a-form-item>

      <a-form-item
          name="password"
          :rules="[{ required: true, message: '请输入密码' }]"
      >
        <a-input-password v-model:value="regFormState.password" placeholder="请输入密码">
          <template #prefix>
            <LockOutlined />
          </template>
        </a-input-password>
      </a-form-item>

      <a-form-item
          name="confirm_password"
          :rules="[{ required: true, message: '请再次输入密码' }]"
      >
        <a-input-password v-model:value="regFormState.confirm_password" placeholder="请再次输入密码">
          <template #prefix>
            <LockOutlined />
          </template>
        </a-input-password>
      </a-form-item>

      <a-form-item>
        <a-form-item name="remember" no-style>
          <a-checkbox v-model:checked="regFormState.agreement">
            <a href="">用户协议、法律声明和隐私政策</a>
          </a-checkbox>
        </a-form-item>
      </a-form-item>

      <a-form-item>
        <a-button :loading="loading" type="primary" html-type="submit" class="login-form-button">
          立即注册
        </a-button>
      </a-form-item>
    </a-form>
  </template>
  </div>
</div>
</template>

<style scoped>
.login-content{
  width: 100%;
  height: 100%;
  display: flex;
  justify-content: center;
  align-items: center;
  flex-direction: column;
  background: #f8f8f8;
}

.logo{
  display: flex;
  align-items: center;
  margin-bottom: 20px;
  justify-content: space-between;
}

.logo img{
  margin-right: 15px;
  width: 35px;
  height: 35px;
}

.logo span{
  font-size: 20px;
  font-weight: 600;
}

.desc{
  color:rgba(0, 0, 0, 0.65);
  font-size: 14px;
  display: flex;
  align-items: center;
  margin-top: 5px;
}

.login-form{
  margin-top: 20px;
  max-width: 350px;
  width: 100%;
}
.login-form-forgot {
  float: right;
}
.login-form-button {
  width: 100%;
}

.login-form-captcha-button{
  width:100px;
}
.login-card{
  background:  #fff;
  width: 100%;
  max-width: 350px;
  padding: 20px 20px 0;
  border-radius: 10px;
}
</style>