<template>
  <div v-if="isLogin !== true">
    <span>
      您还没有登录，请登录！！！
    </span>
  </div>
  <div v-else>
    <template>
      <el-carousel :interval='5000' indicator-position="outside">
        <el-carousel-item>
          <div class="user-info">
            <div class="user-avator">
              <img src="../../assets/logo.png" alt="">
            </div>
            <div class="user-name">
              <span class="ft-danger">大王</span>
            </div>
          </div>
        </el-carousel-item>
        <el-carousel-item>
          1
        </el-carousel-item>
        <el-carousel-item>
          2
        </el-carousel-item>
      </el-carousel>
    </template>
  </div>
</template>

<script>
export default {
  data () {
    return {
      isLogin: this.getUserInfo()
    }
  },
  methods: {
    getUserInfo: function () {
      return true
    }
  }
}
</script>
<style>
.el-carousel__container, .el-carousel__item {
  height: 200px;
}
.el-carousel__item h3 {
  color: #475669;
  font-size: 18px;
  opacity: 0.75;
  line-height: 200px;
  margin: 0;
}
.el-carousel__indicator {
  display: none;
}
.el-carousel__item:nth-child(2n) {
  background-color: #99a9bf;
}
.el-carousel__item:nth-child(2n+1) {
  background-color: #d3dce6;
}
.user-info .user-avator {
  border-bottom: #114e9e;
  height: 140px;
  padding: 20px 0 0 0;
}
.user-info .user-avator img {
  width: 100px;
  height: 100px;
  margin: 0 auto;
  border: 1px solid #cfcfcf;
  border-radius: 50px;
  overflow: hidden;
}
.user-info .user-name {
  height: 30px;
  line-height: 30px;
  text-align: center;
}
</style>
