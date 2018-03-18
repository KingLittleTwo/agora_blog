<template>
  <div class="side-list">
    <div class="side-list-title">
      <span>
        最热文章
      </span>
    </div>
    <div class="side-list-main">
      <ul>
        <li v-for="(item, key) in postlist" :key="key" v-if="key < 10">
          <a href="">{{ key+1 }}. {{ item.title }}</a>
        </li>
      </ul>
    </div>
  </div>
</template>

<script>
export default {
  data: function () {
    return {}
  },
  props: [
    'postlist'
  ]
}
</script>

<style>
.side-list .side-list-main {
  margin-top: 10px;
}
.side-list .side-list-title span, .side-list .side-list-main li {
  height: 35px;
  line-height: 35px;
  overflow: hidden;
}
.side-list .side-list-title{
  padding-bottom: 10px;
  border-bottom: 1px solid #F2F6FC;
}
</style>
