<script setup>

import {MenuFoldOutlined, MenuUnfoldOutlined, UserOutlined} from "@ant-design/icons-vue";
import {useLayoutStore} from "../../stores/layoutStore.js";
import {useUserStore} from "../../stores/userStore.js";

const store = useLayoutStore()
const user = useUserStore()
</script>

<template>
  <a-layout-header>
    <menu-unfold-outlined
        v-if="store.collapsed"
        class="trigger"
        @click="store.changeCollapsed()"
    />
    <menu-fold-outlined v-else class="trigger" @click="store.changeCollapsed()" />

    <div class="right-action">
      <a-dropdown>
        <div class="userinfo" @click.prevent>
          <a-avatar size="30" :src="user.userInfo.avatar">
            <template #icon><UserOutlined /></template>
          </a-avatar>
          <span class="username">
              {{user.userInfo.username}}
            </span>
        </div>
        <template #overlay>
          <a-menu>
            <a-menu-item>
              <a href="javascript:;">1st menu item</a>
            </a-menu-item>
            <a-menu-item>
              <a href="javascript:;">2nd menu item</a>
            </a-menu-item>
            <a-menu-item>
              <a href="javascript:;">3rd menu item</a>
            </a-menu-item>
          </a-menu>
        </template>
      </a-dropdown>
    </div>
  </a-layout-header>
</template>

<style scoped>

</style>