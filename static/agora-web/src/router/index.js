import Vue from 'vue'
import Router from 'vue-router'
// import HelloWorld from '@/components/HelloWorld'
import Navigation from '@/components/Web/Navigation'
import PostDetail from '@/components/Web/PostDetail'
import PostList from '@/components/Web/PostList'

Vue.use(Router)

export default new Router({
// const router = new Router({
  routes: [
    {
      path: '/',
      name: 'PostList',
      component: PostList
    },
    {
      path: '/post/:postId',
      name: 'PostDetail',
      component: PostDetail
    },
    {
      path: '/postlist/:type',
      name: 'PostList',
      component: PostList
    },
    {
      path: '/cate',
      name: 'Category',
      component: Navigation
    }
  ]
})

// router.beforeEach((to, from, next) => {
//   const auth = Storage.state.auth
// })
