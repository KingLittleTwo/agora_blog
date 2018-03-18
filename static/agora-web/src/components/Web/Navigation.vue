<template>
  <el-menu
    :default-active="activeIndex2"
    class="el-menu"
    id="navigation"
    mode="horizontal"
    @select="handleSelect"
    background-color="#545c64"
    text-color="#fff"
    active-text-color="#ffd04b">
    <el-menu-item index="1">
      <router-link :to="{ name: 'PostList', params: { type: 'hot' }}">社区</router-link>
    </el-menu-item>
    <el-menu-item index="3">消息中心</el-menu-item>
    <el-submenu index="" style="float: right">
      <span slot="title">我的工作台</span>
      <el-menu-item index="2-1">选项1</el-menu-item>
      <el-menu-item index="2-2">选项2</el-menu-item>
      <el-menu-item index="2-3">选项3</el-menu-item>
    </el-submenu>
    <el-menu-item index="4" style="float: right">
      <i class="el-icon-message"></i>
    </el-menu-item>
  </el-menu>
</template>

<script>
export default {
  data () {
    return {
      activeIndex: '1',
      activeIndex2: '1'
    }
  },
  methods: {
    handleSelect (key, keyPath) {
      console.log(key, keyPath)
    }
  }
}
</script>

<style>
#navigation {
  margin-bottom: 20px;
}
</style>
