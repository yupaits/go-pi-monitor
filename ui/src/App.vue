<template>
  <div id="app">
    <a-locale-provider :locale="locale">

    </a-locale-provider>
  </div>
</template>

<script>
import zhCN from 'ant-design-vue/lib/locale-provider/zh_CN'
import axios from 'axios'
import Device from './components/Device'
import Cpu from './components/Cpu'
import Disk from './components/Disk'
import Mem from './components/Mem'
import Net from './components/Net'
export default {
  components: {
    Device, Cpu, Disk, Mem, Net
  },
  data() {
    return {
      locale: zhCN,
      metrics: {
        host: {},
        temperatures: [],
        users: [],
        "cpu.counts": undefined,
        "cpu.logicCounts": undefined,
        "cpu.info": [],
        "cpu.percent": [],
        "cpu.percentPerCpu": [],
        "cpu.times": [],
        "cpu.timesPerCpu": [],
        "mem.swap": {},
        "mem.virtual": {},
        "net.connections": [],
        "net.interfaces": [],
        "net.io": [],
        "net.ioPerNic": [],
        "disk.parts": [],
        "disk.usage": {},
        "disk.io": {}
      }
    }
  },
  created() {
    this.fetchInfo();
  },
  methods: {
    fetchInfo() {
      axios.get('/info').then(res => {
        this.metrics = res.data;
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