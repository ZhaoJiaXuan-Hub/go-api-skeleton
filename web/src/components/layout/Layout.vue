<template>
  <a-layout id="sider-content">
    <Sider></Sider>
    <a-layout>
      <Header></Header>
      <a-layout-content
          :style="{ padding: '16px 16px' }"
      >
        <router-view></router-view>
      </a-layout-content>
    </a-layout>
  </a-layout>
</template>
<script setup>
import Sider from "./Sider.vue";
import Header from "./Header.vue";
</script>
<style>
#sider-content .trigger {
  font-size: 18px;
  line-height: 54px;
  padding: 0 24px;
  cursor: pointer;
  transition: color 0.3s;
}

#sider-content .trigger:hover {
  color: #1890ff;
}

#sider-content .site-layout .site-layout-background {
  background: #fff;
}

#sider-content .ant-layout{
  height: 100%;
}

#sider-content .logo {
  width: 250px;
  height: 54px;
  display: flex;
  justify-content: center;
  align-items: center;
  white-space: nowrap;
}

#sider-content .logo span {
  font-size: 20px;
  line-height: 54px;
  margin-left: 16px;
  white-space: nowrap;
}

#sider-content .ant-layout-sider{
  background: rgba(255, 255, 255, 0.6) !important;
  border-inline-end: 1px solid rgba(5, 5, 5, 0.06);
  -webkit-backdrop-filter: blur(8px);
  backdrop-filter: blur(8px);
  transition: background 0.3s cubic-bezier(0.645, 0.045, 0.355, 1);
}
#sider-content .ant-menu-light{
  background: transparent !important;
}
#sider-content .ant-menu-light .ant-menu-item-selected{
  background: rgba(0, 0, 0, 0.04) !important;
  color:rgba(0, 0, 0, 0.95);
}
#sider-content .ant-layout .ant-layout-header{
  height: 54px !important;
  line-height: 54px !important;
}
#sider-content .ant-menu-light:not(.ant-menu-horizontal) .ant-menu-item:not(.ant-menu-item-selected):active {
  background-color: rgba(0, 0, 0, 0.15);
}

#sider-content .ant-menu-light.ant-menu-inline .ant-menu-sub.ant-menu-inline {
  background: rgba(0, 0, 0, 0.02);
  border-radius: 7px;
}

#sider-content .ant-menu-light .ant-menu-submenu-selected >.ant-menu-submenu-title {
  color: #000000;
}

#sider-content .ant-menu-light.ant-menu-root.ant-menu-inline ,.ant-menu-light.ant-menu-root.ant-menu-vertical{
  border:none !important;
  padding: 0 10px !important;
}
#sider-content .ant-layout-header{
  padding: 0 16px !important;
  border-block-end: 1px solid rgba(5, 5, 5, 0.06);
  background: rgba(255, 255, 255, 0.6) !important;
  -webkit-backdrop-filter: blur(8px);
  backdrop-filter: blur(8px);
  transition: background 0.3s cubic-bezier(0.645, 0.045, 0.355, 1);
  position: sticky;
  inset-block-start: 0;
  width: 100%;
  z-index: 100;
  inset-inline-end: 0;
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding-inline:0 !important;
}

#sider-content .ant-row{
  margin-bottom: 16px;
}

#sider-content .ant-layout-content {
  height: 100%;
  overflow: scroll;
}

#sider-content .ant-layout-sider{
  height: 100vh;
  position: sticky;
  top: 0;
}

#sider-content .right-action{
  display: flex;
  align-items: center;
  height: 54px;
  width: 15rem;
  flex-direction: row-reverse;
}

#sider-content .right-action .userinfo{
  margin: 0 12px;
  padding: 0 12px;
  display: flex;
  align-items: center;
}

#sider-content .right-action .userinfo:hover {
  background-color: rgba(0, 0, 0, 0.03);
  border-radius: 7px;
  height: 44px;
}

#sider-content .right-action .userinfo .username{
  margin-left: 7px;
  color: rgba(0, 0, 0, 0.45);
  font-size: 16px;
}

@media screen and (max-width: 600px) {
  #sider-content .ant-layout-sider{
    position: fixed;
    z-index: 999;
    left: -250px;
  }
  #sider-content .ant-layout-sider.show{
    position: fixed;
    z-index: 999;
    left: 0;
  }
  .sider-collapsed-button {
    display: flex !important;
  }
}
</style>