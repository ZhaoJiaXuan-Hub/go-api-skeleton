<script setup>
import {useLayoutStore} from "../../stores/layoutStore.js";
import {reactive, ref, watch, h, onMounted, onUnmounted} from 'vue';
import {
  AppstoreOutlined,
  SettingOutlined
} from '@ant-design/icons-vue';
import {useSettingStore} from "../../stores/settingStore.js";
const store = useLayoutStore()
const setting = useSettingStore();
console.log(setting)
const selectedKeys = ref(['1']);
const openKeys = ref(['sub1']);
function getItem(label, key, icon, children, type) {
  return {
    key,
    icon,
    children,
    label,
    type,
  };
}
const items = reactive([
  getItem('Navigation Two', 'sub1', () => h(AppstoreOutlined), [
    getItem('Option 5', '1'),
    getItem('Option 6', '6'),
    getItem('Submenu', 'sub3', null, [getItem('Option 7', '7'), getItem('Option 8', '8')]),
  ]),
  getItem('Navigation Three', 'sub2', () => h(SettingOutlined), [
    getItem('Option 9', '9'),
    getItem('Option 10', '10'),
    getItem('Option 11', '11'),
    getItem('Option 12', '12'),
  ]),
]);
const handleClick = e => {
  console.log('click', e);
};
watch(openKeys, val => {
  console.log('openKeys', val);
});
const isMobile = ref(false)
const handleResize = () => {
  const screenWidth = window.innerWidth || document.documentElement.clientWidth || document.body.clientWidth;
  isMobile.value = screenWidth < 600;
}
handleResize()

onMounted(() => {
  window.addEventListener('resize', handleResize);
});

onUnmounted(() => {
  window.removeEventListener('resize', handleResize);
});
</script>

<template>
  <a-layout-sider v-if="!isMobile" v-model:collapsed="store.collapsed" :class="{show:!store.collapsed}" width="250px" style="height: 100vh;" :trigger="null" collapsible>
    <div class="logo" :style="{width:store.collapsed?'80px':'250px'}">
      <img width="auto" height="22" :src="setting.logo" alt="logo">
      <span v-if="!store.collapsed">{{setting.siteName}}</span>
    </div>
    <a-menu
        id="sider-menu"
        v-model:openKeys="openKeys"
        v-model:selectedKeys="selectedKeys"
        mode="inline"
        :items="items"
        @click="handleClick"
    ></a-menu>
  </a-layout-sider>

  <a-layout-sider v-if="isMobile" :class="{show:store.collapsed}" width="250px" style="height: 100vh;" :trigger="null" collapsible>
    <div class="logo" :style="{width:'250px'}">
      <img width="auto" height="22" :src="setting.logo" alt="logo">
      <span>{{setting.siteName}}</span>
    </div>
    <button v-if="store.collapsed" class="sider-collapsed-button">
      <span @click="store.changeCollapsed()"><svg width="1em" height="1em" viewBox="0 0 12 12" fill="currentColor" aria-hidden="true"><path d="M6.432 7.967a.448.448 0 01-.318.133h-.228a.46.46 0 01-.318-.133L2.488 4.85a.305.305 0 010-.43l.427-.43a.293.293 0 01.42 0L6 6.687l2.665-2.699a.299.299 0 01.426 0l.42.431a.305.305 0 010 .43L6.432 7.967z"></path></svg></span>
    </button>
    <a-menu
        id="sider-menu"
        v-model:openKeys="openKeys"
        v-model:selectedKeys="selectedKeys"
        mode="inline"
        :items="items"
        @click="handleClick"
    ></a-menu>
  </a-layout-sider>
</template>

<style scoped>
.sider-collapsed-button {
  display: none;
  position: absolute;
  inset-block-start: 15px;
  z-index: 101;
  width: 24px;
  height: 24px;
  text-align: center;
  border-radius: 40px;
  inset-inline-end: -13px;
  transition: transform 0.3s;
  align-items: center;
  justify-content: center;
  cursor: pointer;
  color: rgba(0, 0, 0, 0.25);
  border: none;
  background-color: #ffffff;
  box-shadow: 0 2px 8px -2px rgba(0, 0, 0, 0.05), 0 1px 4px -1px rgba(25, 15, 15, 0.07), 0 0 1px 0 rgba(0, 0, 0, 0.08);
}
.sider-collapsed-button svg{
  transition: transform 0.3s;
  transform: rotate(90deg);
}
.sider-collapsed-button span{
  height: 100%;
  display: flex;
  align-items: center;
}
</style>