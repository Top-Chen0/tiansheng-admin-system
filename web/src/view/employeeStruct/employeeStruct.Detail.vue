<!--员工查看跳转页面-->
<template>
  <div>
    <div class="chen-test-box">
      <div class="demo-image">
        <div v-for="fit in fits" :key="fit" class="block">
          <el-image style="width: 60px; height: 60px" :src="url" :fit="fit" />
        </div>
      </div>
      <div class="chen-test-title">河南天盛景观规划设计院</div>


    </div>

  </div>
</template>

<!--图像加载-->
<script lang="ts" setup>
const fits = [ 'cover' ]
const url =
    'https://fuss10.elemecdn.com/e/5d/4a731a90594a4af544c0c25941171jpeg.jpeg'
</script>



<!--图像加载-->
<style>
/*页面头部样式*/
.chen-test-box {
  padding:  10px;
  background-color: #377ebb;
  border-radius: 2px;
}
/*logo样式*/
.demo-image .block {
  padding: 10px 0;
  /*text-align: center;*/
  /*display: inline-block;*/
  width: 60%;
  /*box-sizing: border-box;*/
  /*vertical-align: top;*/
}
/*logo未加载时的背景图样式*/
.demo-image .block:last-child {
  border-right: none;
}
/*字体样式*/
.chen-test-title {
  /*text-align: center;*/
  width: 180%;
  height: 1px;
  /*字体大小*/
  font-size: 50px;
  /*字体粗细*/
  font-weight: 500;
  color: #FFFFFF;
  /*引入外部字体*/
  font-family: 'QingNiaoHuaGuangXingKai';
  display: inline-block;
}

</style>
