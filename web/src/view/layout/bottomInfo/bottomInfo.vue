<template>
  <div class="bottom-info">
    <div>
      <span>Powered by</span>
      <span>
        <a href="https://top-chen.com">{{ $GIN_VUE_ADMIN.appNamee }}</a>
      </span>
      <el-divider direction="vertical" />
      <span>Copyright</span>
      <span>
        <a href="https://top-chen.com">陈俊豪-李广奇团队-开发环境</a>
      </span>
    </div>
  </div>
</template>

<script>

export default {
  name: 'BottomInfo'
}
</script>

<style lang="scss">
.bottom-info {
  color: #888;
  height: 30px;
  line-height: 12px;
  a {
    color: #888;
  }
  div {
    display: flex;
    justify-content: center;
    span{
      margin: 0 3px;
    }
  }
}
</style>
