<template>
  <div id="app">
    <a-locale-provider :locale="locale">

    </a-locale-provider>
  </div>
</template>

<script>
  import zhCN from 'ant-design-vue/lib/locale-provider/zh_CN'
  import axios from 'axios'
  export default {
    data() {
      return {
        locale: zhCN,
        info: {}
      }
    },
    created() {
      this.fetchInfo();
    },
    methods: {
      fetchInfo() {
        axios.get('/info').then(res => {

        }).catch(() => {
          this.$message.error('获取监控信息失败');
        });
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