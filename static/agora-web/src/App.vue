<template>
  <div id='app'>
    <vnav></vnav>
    <agora></agora>
    <vfooter></vfooter>
  </div>
</template>

<script>
import nav from './components/Web/Navigation'
import agora from './components/Web/Agora'
import footer from './components/Web/Footer'

export default {
  components: {
    vnav: nav,
    agora: agora,
    vfooter: footer
  },
  name: 'App',
  data () {
    return {
      userinfo: ''
    }
  },
  created: function () {
    this.userinfo = this.getCookie('gosessionid')
    console.log(1111111111111111)
    console.log(this.userinfo)
  },
  methods: {
    // 设置cookie
    setCookie: function (cname, cvalue, exdays) {
      var d = new Date()
      d.setTime(d.getTime() + (exdays * 24 * 60 * 60 * 1000))
      var expires = 'expires=' + d.toUTCString()
      console.info(cname + '=' + cvalue + '; ' + expires)
      document.cookie = cname + '=' + cvalue + '; ' + expires
      console.info(document.cookie)
    },
    // 获取cookie
    getCookie2: function (cname) {
      if (document.cookie.length > 0) {
        var cstart = document.cookie.indexOf(cname + '=')
        if (cstart !== -1) {
          cstart = cstart + cname.length + 1
          var cend = document.cookie.indexOf(';', cstart)
          if (cend === -1) {
            cend = document.cookie.length
          }
          return unescape(document.cookie.substring(cstart, cend))
        }
      }
      return ''
    },
    getCookie: function (name) {
      var reg = new RegExp('(^| )' + name + ' = ([^;]*)(;|$)')
      if (reg === document.cookie.match(reg)) {
        return unescape(reg[2])
      } else {
        return null
      }
    },
    // 清除cookie
    clearCookie: function () {
      this.setCookie('username', '', -1)
    },
    checkCookie: function () {
      var user = this.getCookie('username')
      if (user !== '') {
        alert('Welcome again ' + user)
      } else {
        user = prompt('Please enter your name:', '')
        if (user !== '' && user != null) {
          this.setCookie('username', user, 365)
        }
      }
    }
  }
}
</script>

<style>
#app {
  font-family: 'Avenir', Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
}
</style>
