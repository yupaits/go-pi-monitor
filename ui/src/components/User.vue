<template>
  <a-card :loading="loading" class="users-card">
    <a-card-meta title="Users">
      <template slot="description">
        <div v-for="(user, index) in loggedInUsers" :key="index" class="user-item mb-1">
          <b>
            <a-icon type="user"/> {{user.user}} <a-icon type="cluster" class="ml-1"/> {{user.host}}（{{user.terminal}}）
            <span class="pull-right">{{dayjs(user.started * 1000).format('YYYY-MM-DD HH:mm:ss')}}</span>
          </b>
        </div>
      </template>
    </a-card-meta>
  </a-card>
</template>

<script>
import dayjs from 'dayjs'
export default {
  name: "User",
  props: {
    loading: {
      type: Boolean,
      default: true
    },
    users: {
      type: Array,
      default: []
    }
  },
  computed: {
    loggedInUsers() {
      return this.users || [];
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
.users-card {
  min-height: 256px;
}
.user-item {
  display: block;
  line-height: 32px;
  padding: 2px 4px;
  border: 1px solid #e5e5e5;
  border-radius: 4px;
}
</style>