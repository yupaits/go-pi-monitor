<template>
  <a-card :loading="loading" class="monitor-card">
    <a-card-meta :title="host.platform || 'Platform'">
      <template slot="description">
        <div class="device-item"><b>版本：{{host.platformVersion || 'Version'}}</b></div>
        <div class="device-item"><b>已运行：{{dayjs().subtract(host.uptime || 0, 'second').fromNow(true)}}</b></div>
        <div class="device-item"><b>主机名：{{host.hostname || 'Hostname'}}</b></div>
      </template>
    </a-card-meta>
  </a-card>
</template>

<script>
import dayjs from 'dayjs'
import zhCN from 'dayjs/locale/zh-cn'
import relativeTime from 'dayjs/plugin/relativeTime'
dayjs.locale(zhCN)
dayjs.extend(relativeTime)
export default {
  name: "Device",
  props: {
    loading: {
      type: Boolean,
      default: true,
    },
    host: {
      type: Object,
      default: {
        platform: 'Platform',
        platformVersion: 'Version',
        uptime: 0,
        hostname: 'Hostname'
      }
    }
  },
  data() {
    return {
      dayjs
    }
  }
}
</script>

<style scoped>
.device-item {
  line-height: 32px;
}
</style>