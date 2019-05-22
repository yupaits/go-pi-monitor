<template>
  <a-card :loading="loading" class="monitor-card">
    <a-progress type="dashboard" :percent="virtualPercent" :status="virtualStatus"></a-progress>
    <a-progress type="dashboard" :percent="swapPercent" :status="swapStatus" class="ml-3"></a-progress>
    <a-card-meta title="Mem">
      <template slot="description">
        <div>虚拟内存：
          <span>使用中 <b>{{`${$utils.bytesToSize(virtualMem.used)} / ${$utils.bytesToSize(virtualMem.total)}`}}</b></span> 
          <span class="ml-1">可用 <b>{{`${$utils.bytesToSize(virtualMem.available)}`}}</b></span>
        </div>
        <div>交换内存：
          <span>已提交 <b>{{`${$utils.bytesToSize(swapMem.used)} / ${$utils.bytesToSize(swapMem.total)}`}}</b></span>
          <span class="ml-1">可用 <b>{{`${$utils.bytesToSize(swapMem.free)}`}}</b></span>
        </div>
      </template>
    </a-card-meta>
  </a-card>
</template>

<script>
export default {
  name: "Mem",
  props: {
    loading: {
      type: Boolean,
      default: true
    },
    swapMem: {
      type: Object,
      default: {}
    },
    virtualMem: {
      type: Object,
      default: {}
    }
  },
  computed: {
    swapPercent() {
      return Math.round(this.swapMem.usedPercent * 100) || 0;
    },
    virtualPercent() {
      return Math.round(this.virtualMem.usedPercent) || 0;
    },
    swapStatus() {
      return this.$utils.percentStatus(this.swapPercent);
    },
    virtualStatus() {
      return this.$utils.percentStatus(this.virtualPercent);
    }
  }
}
</script>

<style scoped>

</style>