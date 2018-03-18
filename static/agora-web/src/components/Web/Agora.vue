<template>
  <el-container>
    <el-main class="">
      <router-view></router-view>
    </el-main>
    <el-aside class="box-side">
      <div class="box box-margin box-padding">
        <vuserinfo></vuserinfo>
      </div>
      <div class="box box-margin box-padding">
        <vhotlist :postlist="postlist"></vhotlist>
      </div>
      <div class="box box-margin box-padding">
        Aside22222222 <br>
        Aside22222222 <br>
        Aside22222222 <br>
        Aside22222222 <br>
        Aside22222222 <br>
        Aside22222222 <br>
        Aside22222222 <br>
      </div>
    </el-aside>
  </el-container>
</template>

<script>
import postList from './PostList'
import sideUserinfo from './SideUserinfo'
import sideList from './SideList'

export default {
  components: {
    vpostlist: postList,
    vuserinfo: sideUserinfo,
    vhotlist: sideList
  },
  data: function () {
    return {
      postlist: [
        {avator: '../../assets/logo.png', title: '这是一大坨测试数据111111111111！！！', id: 1},
        {avator: '../../assets/logo.png', title: '这是一大坨测试数据！！！', id: 1},
        {avator: '../../assets/logo.png', title: '这是一大坨测试数据！！！', id: 1},
        {avator: '../../assets/logo.png', title: '这是一大坨测试数据！！！', id: 1},
        {avator: '../../assets/logo.png', title: '这是一大坨测试数据！！！', id: 1},
        {avator: '../../assets/logo.png', title: '这是一大坨测试数据！！！', id: 1},
        {avator: '../../assets/logo.png', title: '这是一大坨测试数据2222222222222！！！', id: 1},
        {avator: '../../assets/logo.png', title: '这是一大坨测试数据！！！', id: 1},
        {avator: '../../assets/logo.png', title: '这是一大坨测试数据！！！', id: 1},
        {avator: '../../assets/logo.png', title: '这是一大坨测试数据！！！', id: 1},
        {avator: '../../assets/logo.png', title: '这是一大坨测试数据！！！', id: 1},
        {avator: '../../assets/logo.png', title: '这是一大坨测试数据！！！', id: 1},
        {avator: '../../assets/logo.png', title: '这是一大坨测试数据！！！', id: 1},
        {avator: '../../assets/logo.png', title: '这是一大坨测试数据333333333333！！！', id: 1},
        {avator: '../../assets/logo.png', title: '这是一大坨测试数据！！！', id: 1}
      ]
    }
  }
}
</script>

<style>
.el-main{
  margin: 0;
  padding: 0;
}
</style>
