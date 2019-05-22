<template>
  <div id="app">
    <a-locale-provider :locale="locale">
      <a-row :gutter="24">
        <a-col :span="8" class="mb-2">
          <device :loading="loading" :host="metrics.host"/>
        </a-col>
        <a-col :span="8" class="mb-2">
          <cpu :loading="loading" :info="metrics['cpu.info']" :percent="metrics['cpu.percent']"/>
        </a-col>
        <a-col :span="8" class="mb-2">
          <mem :loading="loading" :swapMem="metrics['mem.swap']" :virtualMem="metrics['mem.virtual']"/>
        </a-col>
        <a-col :span="8" class="mb-2">
          <net :loading="loading" :connections="metrics['net.connections']" :interfaces="metrics['net.interfaces']" :io="metrics['net.io']"/>
        </a-col>
        <a-col :span="16" class="mb-2">
          <disk :loading="loading" :parts="metrics['disk.parts']" :usage="metrics['disk.usage']"></disk>
        </a-col>
        <a-col :span="8">
          <user :loading="loading" :users="metrics.users"/>
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
import User from './components/User'
export default {
  components: {
    Device, Cpu, Disk, Mem, Net, User
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
  padding: 24px 256px;
}
.monitor-card {
  height: 256px;
}
.pull-right {
  float: right;
}
.ml-1 {
  margin-left: 8px;
}
.ml-2 {
  margin-left: 16px;
}
.ml-3 {
  margin-left: 24px;
}
.mb-1 {
  margin-bottom: 8px;
}
.mb-2 {
  margin-bottom: 16px;
}
.mr-2 {
  margin-right: 16px;
}
</style>