<template>
  <a-card :loading="loading" class="net-card">
    <a-card-meta :title="interfaceName">
      <template slot="description">
        <div class="net-item"><b>状态：{{status}}</b></div>
        <div class="net-item" v-for="(addr, index) in addrs" :key="addr"><b>地址{{index}}：{{addr}}</b></div>
        <div class="net-item"><b>连接数：{{connCounts}}</b></div>
        <div class="net-item"><b>数据包：{{`已发送 ${netIO.packetsSent} 个 已接收 ${netIO.packetsRecv} 个`}}</b></div>
        <div class="net-item"><b>数据：{{`已发送 ${netIO.bytesSent} 字节 已接收 ${netIO.bytesRecv} 字节`}}</b></div>
      </template>
    </a-card-meta>
  </a-card>
</template>

<script>
export default {
  name: "Net",
  props: {
    loading: {
      type: Boolean,
      default: true
    },
    connections: {
      type: Array,
      default: 0,
    },
    interfaces: {
      type: Array,
      default: []
    },
    io: {
      type: Array,
      default: []
    },
  },
  computed: {
    connCounts() {
      if (!this.connections) {
        return 0;
      }
      return this.connections.length || 0;
    },
    interface() {
      let result = {addrs: [], flags: []};
      if (!this.interfaces) {
        return result;
      } 
      this.interfaces.forEach(iface => {
        if (iface.mtu && iface.mtu > 0) {
          result = iface;
        }
      });
      return result;
    },
    addrs() {
      let addrList = [];
      this.interface.addrs.forEach(addr => {
        addrList.push(addr.addr);
      });
      return addrList;
    },
    interfaceName() {
      return this.interface.name || '未知';
    },
    status() {
      if (this.interface.flags.indexOf('up') !== -1) {
        return 'online';
      } else {
        return 'offline';
      }
    },
    netIO() {
      return this.io[0] || {bytesSent: 0, bytesRecv: 0, packetsSent: 0, packetsRecv: 0};
    }
  }
}
</script>

<style scoped>
.net-card {
  min-height: 256px;
}
.net-item {
  line-height: 32px;
}
</style>