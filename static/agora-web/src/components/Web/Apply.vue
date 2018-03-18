<template>
<div>
  <div class="zan-detail box box-margin box-padding">
    <!-- <button class="btn btn-danger" type="button">
      <span class="glyphicon glyphicon-star-empty" aria-hidden="true"></span>
    </button> -->
    <span>赞过的人</span>
    <div class="zan-box">
        <img v-for="(item, key) in 6" src="../../assets/logo.png" :alt="item" :key="key" class="img-circle">
    </div>
  </div>
  <div class="comment box box-margin box-padding" style="padding-bottom: 0">
    <div class="comment-area">
      <el-form :model="commentForm" :rules="rules" ref="commentForm" label-width="0" class="commentForm">
        <el-form-item prop="desc">
          <el-input type="textarea" v-model="commentForm.desc" size="medium"></el-input>
        </el-form-item>
        <el-form-item>
          <el-button class="pull-right" type="danger" @click="submitForm('commentForm')">回复</el-button>
        </el-form-item>
      </el-form>
    </div>
  </div>
  <div class="box box-margin box-padding comment-list">
    <el-breadcrumb class="post-bread" separator="/">
      <el-breadcrumb-item :to="{ path: '/' }">
      <span class="ft-danger">活跃</span>
      </el-breadcrumb-item>
      <el-breadcrumb-item>精华</el-breadcrumb-item>
      <el-breadcrumb-item>最近</el-breadcrumb-item>
    </el-breadcrumb>
    <div class="comment-item" v-for="(item, key) in 6" :key="key">
      <div class="comment-item-header">
        <img class="pull-left" src="" alt="">
        <p>fbdshjfndjsnfkdsnfkmdsmfkldsmflkds</p>
      </div>
      <div class="comment-item-bottom">
        <span>点赞</span>
        <span @click="repply($event)">回复</span>
      </div>
      <div class="apply-form">
        <el-form :model="applyForm" :rules="applyRules" ref="applyForm" label-width="0" class="applyForm">
          <el-form-item prop="desc">
            <el-input type="textarea" v-model="applyForm.desc" size="medium"></el-input>
          </el-form-item>
          <el-form-item>
            <el-button size="mini" class="pull-right" type="danger" @click="submitForm('applyForm')">回复</el-button>
          </el-form-item>
        </el-form>
      </div>
    </div>
  </div>
</div>
</template>

<script>
export default {
  data () {
    return {
      commentForm: {
        desc: ''
      },
      rules: {
        desc: [
          { type: 'date', required: true, message: '', trigger: 'change' }
        ]
      },
      applyForm: {
        desc: ''
      },
      applyRules: {
        desc: [
          { type: 'date', required: true, message: '', trigger: 'change' }
        ]
      }
    }
  },
  methods: {
    repply (event) {
      console.log(event.currentTarget)
      event.currentTarget.parent.parent.find('.apply-form').style.display = 'block'
    }
  }
}
</script>
<style>
.zan-detail{
  text-align: center;
}
.zan-detail .img-circle{
  width: 30px;
  height: 30px;
  border-radius: 25px;
  border: 1px solid #cfcfcf;
  display: inline-block;
  margin: 20px 5px 10px;
  padding: 6px;
}
.comment textarea {
  height: 150px;
}
.comment-list .comment-item {
  padding: 10px 0;
  border-bottom: 1px dotted #cfcfcf;
}
.comment-list .comment-item .comment-item-header {
  height: 46px;
  margin: 10px 0;
}
.comment-list .comment-item .comment-item-header img{
  width: 46px;
  height: 100%;
  background: red;
  border-radius: 23px;
  overflow: hidden;
  display: inline-block;
}
.comment-list .comment-item .comment-item-header p {
  display: inline-block;
  height: 100%;
  padding-left: 20px;
}
.comment-list .comment-item .comment-item-bottom {
  text-align: right;
  font-size: 14px;
  margin-bottom: 20px;
}
.comment-list .comment-item .comment-item-bottom span {
  margin: 0 5px;
  cursor: pointer;
}
.comment-list .comment-item .comment-item-bottom span:hover {
  color: #F56C6C;
}
.apply-form {
  display: none;
}
</style>
