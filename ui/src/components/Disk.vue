<template>
  <a-card :loading="loading" class="disk-card">
    <a-progress type="dashboard" :percent="diskPercent" :status="diskStatus"></a-progress>
    <span class="ml-3">
      <span>已使用：<b>{{`${$utils.bytesToSize(diskUsage.used)} / ${$utils.bytesToSize(diskUsage.total)}`}}</b></span>
      <span class="ml-2">可用：<b>{{$utils.bytesToSize(diskUsage.free)}}</b></span>      
    </span>
    <a-card-meta title="Disk">
      <template slot="description">
        <span v-for="part in diskParts" :key="part.device" class="part-item mb-1">
          <b>
            <a-icon type="hdd"/> {{part.fstype}} 
            <span class="ml-2">{{part.device}} / {{part.mountpoint}}</span>
          </b>
        </span>
      </template>
    </a-card-meta>
  </a-card>
</template>

<script>
export default {
  name: "Disk",
  props: {
    loading: {
      type: Boolean,
      default: true
    },
    parts: {
      type: Array,
      default: []
    },
    usage: {
      type: Object,
      default: {}
    },
  },
  computed: {
    diskUsage() {
      return this.usage || {};
    },
    diskPercent() {
      return Math.round(this.diskUsage.usedPercent) || 0;
    },
    diskStatus() {
      return this.$utils.percentStatus(this.diskPercent);
    },
    diskParts() {
      return this.parts || [];
    }
  }
}
</script>

<style scoped>
.disk-card {
  min-height: 256px;
}
.part-item {
  display: block;
  line-height: 32px;
  padding: 2px 4px;
  border: 1px solid #e5e5e5;
  border-radius: 4px;
}
</style>