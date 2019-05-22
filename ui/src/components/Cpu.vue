<template>
  <a-card :loading="loading" class="monitor-card">
    <a-progress type="dashboard" :percent="cpuPercent" :status="cpuStatus"></a-progress>
    <a-card-meta title="CPU">
      <template slot="description">
        <div><b>{{cpuInfo.modelName}}</b></div>
        <div><b>{{cpuInfo.cores}}</b> Threads <b class="ml-2">{{cpuInfo.mhz}}</b> MHz</div>
      </template>
    </a-card-meta>
  </a-card>
</template>

<script>
export default {
  name: "Cpu",
  props: {
    loading: {
      type: Boolean,
      default: true,
    },
    info: {
      type: Array,
      default: []
    },
    percent: {
      type: Array,
      default: []
    }
  },
  computed: {
    cpuInfo() {
      return this.info[0] || {cores: NaN, mhz: NaN};
    },
    cpuPercent() {
      return Math.round(this.percent[0]) || 0;
    },
    cpuStatus() {
      return this.$utils.percentStatus(this.cpuPercent);
    }
  }
}
</script>

<style scoped>

</style>