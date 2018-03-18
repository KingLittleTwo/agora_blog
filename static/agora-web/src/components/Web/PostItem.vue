<template>
  <div>
    <div class="post-list">
        <div v-for="(item, key) in postlist" :key="key" class="post-item cl">
          <div class="avator-list">
            <a href="http://">
              <img :src="item.avator" alt="">
            </a>
          </div>
          <div class="post-item-title">
            <span class="tag tag-sucess">
              表情
            </span>
            <p>
              {{ item.title }}
            </p>
          </div>
          <span class="date-time pull-right">
              <a title="33" href="">33</a>
              <span>/</span>
              <a href="">2</a>
              <span>/</span>
              <a href="">1</a>
              <span>|</span>
              <a class="date-time-recent">5小时前</a>
            </span>
        </div>
    </div>
  </div>
</template>

<script>
export default {
  data () {
    return {}
  },
  props: [
    'postlist'
  ]
}
</script>

<style>
.post-bread {
  height: 30px;
  line-height: 30px;
  padding-bottom: 20px;
  border-bottom: 1px solid #F2F6FC;
}
.post-list .post-item {
  padding: 10px;
  border-bottom: 1px solid #F2F6FC;
}
.post-list .post-item:hover {
  background: #F2F6FC;
}
.avator-list {
  width: 46px;
  height: 46px;
  background: red;
  border-radius: 23px;
  overflow: hidden;
  display: block;
  float: left;
}
.avator-list img {
  width: 46px;
  height: 46px;
}
.post-item .post-item-title, .post-item .date-time {
  line-height: 46px;
  display: inline-block;
  padding-left: 20px;

}
.post-item .post-item-title p {
  display: inline-block;
  color: #303133;
}
.post-item .date-time a {
  color: #606266;
}
.post-item .date-time span {
  color: #C0C4CC;
}
.post-item .date-time .date-time-recent {
  font-size: 14px;
}
</style>
