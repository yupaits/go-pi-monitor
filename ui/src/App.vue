<template>
  <div id="app">
    <a-locale-provider :locale="locale">
      <a-row :gutter="24">
        <a-col :span="6">
          <device :loading="loading" :host="metrics.host"/>
        </a-col>
        <a-col :span="6">
          <cpu :loading="loading" :info="metrics['cpu.info']" :percent="metrics['cpu.percent']"/>
        </a-col>
      </a-row>
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
      loading: false,
      isFirst: true,
      metrics: {
        host: {},
        temperatures: [],
        users: [],
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
    this.schedule();
  },
  methods: {
    fetchInfo() {
      if (this.isFirst) {
        this.loading = true;
      }
      axios.get('/info').then(res => {
        this.metrics = res.data;
        if (this.isFirst) {
          this.loading = false;
          this.isFirst = false;
        }
      }).catch(() => {
        this.$message.error('获取监控信息失败');
        if (this.isFirst) {
          this.loading = false;
          this.isFirst = false;
        }
      });
    },
    schedule() {
      setInterval(this.fetchInfo, 5000);
    }
  }
}
</script>

<style>
#app {
  font-family: 'Avenir', Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  padding: 24px 64px;
}
.monitor-card {
  height: 240px;
}
</style>