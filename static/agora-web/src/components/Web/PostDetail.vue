<template>
  <div>
    <div class="post-detail box box-margin box-padding">
        <h1>
          {{ post.title }}
        </h1>
        <div class="post-info">
          <span>
            {{ post.create_time }}
          </span>
          <span>
            333
            <!-- {{ post.zan_ids.count() }} -->
          </span>
        </div>
        <div class="post-content">
          {{ post.content }}
        </div>
    </div>
    <vapply></vapply>
  </div>
</template>

<script>
import apply from './Apply'
export default {
  components: {
    vapply: apply
  },
  data () {
    return {
      post: {
        id: 1,
        title: '这是一个标题',
        sub_title: '大副标题',
        author: '大王',
        create_time: '2018-02-20 01:12:43',
        source: '自创',
        img: '',
        content: '范德萨发你上开机女方家的快感怒放的是哪看见处女就科学城那几款是女方家看到过hiu',
        zan_ids: [1, 2, 3, 5]
      }
    }
  }
}
</script>

<style>
.post-bread {
  height: 30px;
  line-height: 30px;
  padding-bottom: 20px;
  border-bottom: 1px solid #F2F6FC;
}
.post-detail h1{
  font-size: 28px;
  font-weight: 200;
  margin: 20px;
  text-align: center;
}
.post-detail .post-info {
  text-align: center;
}
.post-detail .post-info span {
  padding: 0 10px;
}
.post-detail .post-content {
  margin: 20px 0;
  padding-top: 20px;
  border-top: 1px solid #F2F6FC;
}
</style>
