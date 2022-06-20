<!--员工信息变更页面-->
<template>
  <!--  外层容器设置了高度和overflow实现滚动-->
  <div style="height: 1000px; overflow: scroll">
    <!--    注意：chen-test-box已经被注解掉，整体无效-->
    <div class="chen-test-box">
      <div class="chen-test-title" v-for="fit in fits" :key="fit" >
        <el-image style="width: 50px; height: 50px;" :src="url" :fit="fit" />
        河南天盛景观规划设计院
        <!--      <span class="chen-test-title">河南天盛景观规划设计院</span>-->
      </div>
    </div>
    <!--    第一个标题框员工信息变更-->
    <div>
      <div class="tableTitle2"><span class="midText2" style="font-family: QingNiaoHuaGuangXingKai; background-color: #377ebb">详单信息</span></div>
    </div>
    <!--    员工信息变更页面-->
    <div class="example-block" style="margin-top: 25px">
      <div class="example-block0">
        <span class="example-demonstration; el-cascader-panel; input:focus"></span>
        <!--1-1小标题-->
        <img
            class="css-img31"
            src="@/image/31.png"
            alt="元素加载失败"
            style="position: relative; left: 170px; top: 0%"
        >
        <!--1-2小标题-->
        <img
            class="css-img32"
            src="@/image/32.png"
            alt="元素加载失败"
            style="position: relative; left: 170px; top: -5%"
        >
        <!--1-3小标题-->
        <img
            class="css-img32"
            src="@/image/32.png"
            alt="元素加载失败"
            style="position: relative; left: 170px; top: -5%"
        >
        <!--1-4小标题-->
        <img
            class="css-img32"
            src="@/image/32.png"
            alt="元素加载失败"
            style="position: relative; left: 170px; top: -5%"
        >
        <!--1-5小标题-->
        <img
            class="css-img32"
            src="@/image/32.png"
            alt="元素加载失败"
            style="position: relative; left: 170px; top: -5%"
        >
        <!--1-6小标题-->
        <img
            class="css-img32"
            src="@/image/32.png"
            alt="元素加载失败"
            style="position: relative; left: 170px; top: -5%"
        >
        <!--1-7小标题-->
        <img
            class="css-img32"
            src="@/image/32.png"
            alt="元素加载失败"
            style="position: relative; left: 170px; top: -5%"
        >
        <br>
        <!--1-2-1小标题-->
        <img
            class="css-img33"
            src="@/image/33.png"
            alt="元素加载失败"
            style="position: relative; left: 332px; top: -15%"
        >
        <!--1-2-2小标题-->
        <img
            class="css-img33"
            src="@/image/33.png"
            alt="元素加载失败"
            style="position: relative; left: 280px; top: -15%"
        >
        <br>
        <!--2-1小标题-->
        <img
            class="css-img31"
            src="@/image/31.png"
            alt="元素加载失败"
            style="position: relative; left: 170px; top: -16%"
        >
        <!--2-2小标题-->
        <img
            class="css-img34"
            src="@/image/34.png"
            alt="元素加载失败"
            style="position: relative; left: 170px; top: -16%"
        >
        <!--2-3小标题-->
        <img
            class="css-img34"
            src="@/image/34.png"
            alt="元素加载失败"
            style="position: relative; left: 170px; top: -16%"
        >
        <!--2-4小标题-->
        <img
            class="css-img34"
            src="@/image/34.png"
            alt="元素加载失败"
            style="position: relative; left: 170px; top: -16%"
        >
        <br>
        <!--3-1小标题-->
        <img
            class="css-img31"
            src="@/image/31.png"
            alt="元素加载失败"
            style="position: relative; left: 170px; top: -16%"
        >
        <!--3-2小标题-->
        <img
            class="css-img35"
            src="@/image/35.png"
            alt="元素加载失败"
            style="position: relative; left: 165px; top: -16%"
        >
        <br>
        <!--4-1小标题-->
        <img
            class="css-img31"
            src="@/image/31.png"
            alt="元素加载失败"
            style="position: relative; left: 170px; top: -16%"
        >
        <!--4-2小标题-->
        <img
            class="css-img35"
            src="@/image/35.png"
            alt="元素加载失败"
            style="position: relative; left: 165px; top: -16%"
        >
      </div>
    </div>
    <div class="example-block1">
      <!--        添加按钮样式-->
      <img
          class="css-img36"
          src="@/image/36.png"
          alt="元素加载失败"
          style="position: relative; left: 300px; top: -25%"
      >
      <!--        添加按钮样式-->
      <img
          class="css-img37"
          src="@/image/37.png"
          alt="元素加载失败"
          style="position: relative; left: 600px; top: -20%"
      >
      <!--        添加按钮样式-->
      <img
          class="css-img36"
          src="@/image/36.png"
          alt="元素加载失败"
          style="position: relative; left: 900px; top: -25%"
      >
    </div>
    <!--    添加页面第二条横线-->
    <div class="tableTitle0"><span class="midText0"></span></div>
    <!--    表格-->
    <div class="example-block" style="background: #e0e8fc;">
      <div class="example-block0">
        <!--表格-->
        <img
            class="css-img39"
            src="@/image/39.png"
            alt="元素加载失败"
            style="position: relative; left: 0px; top: 0%"
        >
      </div>
    </div>
    <!--    添加页面第三条横线-->
    <div class="tableTitle0"><span class="midText0"></span></div>
    <div class="example-block1">
      <!--        主页按钮样式-->
      <img
          class="css-img24"
          src="@/image/24.png"
          alt="元素加载失败"
          style="position: relative; left: 60px; top: 0%"
      >
      <!--        上一页按钮样式-->
      <img
          class="css-img24"
          src="@/image/24.png"
          alt="元素加载失败"
          style="position: relative; left: 60px; top: 0%"
      >
      <!--        下一页按钮样式-->
      <img
          class="css-img24"
          src="@/image/24.png"
          alt="元素加载失败"
          style="position: relative; left: 60px; top: 0%"
      >
      <!--        取消按钮样式-->
      <img
          class="css-img24"
          src="@/image/24.png"
          alt="元素加载失败"
          style="position: relative; left: 900px; top: 0%"
      >
      <!--        确定按钮样式-->
      <img
          class="css-img24"
          src="@/image/24.png"
          alt="元素加载失败"
          style="position: relative; left: 900px; top: 0%"
      >
    </div>
    <!--最底栏-->
    <div style="width: 100%;height: 110px;background-color: #377ebb;margin-top: 10px;">
    </div>
  </div>
</template>

<!--图像加载-->
<script lang="ts" setup>
const fits = [ 'cover' ]
const url =
    '/src/assets/logo.png'

// 页面下拉栏Cascader 级联选择器
import { ref } from 'vue'

const value = ref([])

const props = {
  expandTrigger: 'hover',
}

const handleChange = (value) => {
  console.log(value)
}

const options = [
  {
    value: 'guide',
    label: 'Guide',
    children: [
      {
        value: 'disciplines',
        label: 'Disciplines',
        children: [
          {
            value: 'consistency',
            label: 'Consistency',
          },
          {
            value: 'feedback',
            label: 'Feedback',
          },
          {
            value: 'efficiency',
            label: 'Efficiency',
          },
          {
            value: 'controllability',
            label: 'Controllability',
          },
        ],
      },
      {
        value: 'navigation',
        label: 'Navigation',
        children: [
          {
            value: 'side nav',
            label: 'Side Navigation',
          },
          {
            value: 'top nav',
            label: 'Top Navigation',
          },
        ],
      },
    ],
  },
  {
    value: 'component',
    label: 'Component',
    children: [
      {
        value: 'basic',
        label: 'Basic',
        children: [
          {
            value: 'layout',
            label: 'Layout',
          },
          {
            value: 'color',
            label: 'Color',
          },
          {
            value: 'typography',
            label: 'Typography',
          },
          {
            value: 'icon',
            label: 'Icon',
          },
          {
            value: 'button',
            label: 'Button',
          },
        ],
      },
      {
        value: 'form',
        label: 'Form',
        children: [
          {
            value: 'radio',
            label: 'Radio',
          },
          {
            value: 'checkbox',
            label: 'Checkbox',
          },
          {
            value: 'input',
            label: 'Input',
          },
          {
            value: 'input-number',
            label: 'InputNumber',
          },
          {
            value: 'select',
            label: 'Select',
          },
          {
            value: 'cascader',
            label: 'Cascader',
          },
          {
            value: 'switch',
            label: 'Switch',
          },
          {
            value: 'slider',
            label: 'Slider',
          },
          {
            value: 'time-picker',
            label: 'TimePicker',
          },
          {
            value: 'date-picker',
            label: 'DatePicker',
          },
          {
            value: 'datetime-picker',
            label: 'DateTimePicker',
          },
          {
            value: 'upload',
            label: 'Upload',
          },
          {
            value: 'rate',
            label: 'Rate',
          },
          {
            value: 'form',
            label: 'Form',
          },
        ],
      },
      {
        value: 'data',
        label: 'Data',
        children: [
          {
            value: 'table',
            label: 'Table',
          },
          {
            value: 'tag',
            label: 'Tag',
          },
          {
            value: 'progress',
            label: 'Progress',
          },
          {
            value: 'tree',
            label: 'Tree',
          },
          {
            value: 'pagination',
            label: 'Pagination',
          },
          {
            value: 'badge',
            label: 'Badge',
          },
        ],
      },
      {
        value: 'notice',
        label: 'Notice',
        children: [
          {
            value: 'alert',
            label: 'Alert',
          },
          {
            value: 'loading',
            label: 'Loading',
          },
          {
            value: 'message',
            label: 'Message',
          },
          {
            value: 'message-box',
            label: 'MessageBox',
          },
          {
            value: 'notification',
            label: 'Notification',
          },
        ],
      },
      {
        value: 'navigation',
        label: 'Navigation',
        children: [
          {
            value: 'menu',
            label: 'Menu',
          },
          {
            value: 'tabs',
            label: 'Tabs',
          },
          {
            value: 'breadcrumb',
            label: 'Breadcrumb',
          },
          {
            value: 'dropdown',
            label: 'Dropdown',
          },
          {
            value: 'steps',
            label: 'Steps',
          },
        ],
      },
      {
        value: 'others',
        label: 'Others',
        children: [
          {
            value: 'dialog',
            label: 'Dialog',
          },
          {
            value: 'tooltip',
            label: 'Tooltip',
          },
          {
            value: 'popover',
            label: 'Popover',
          },
          {
            value: 'card',
            label: 'Card',
          },
          {
            value: 'carousel',
            label: 'Carousel',
          },
          {
            value: 'collapse',
            label: 'Collapse',
          },
        ],
      },
    ],
  },
  {
    value: 'resource',
    label: 'Resource',
    children: [
      {
        value: 'axure',
        label: 'Axure Components',
      },
      {
        value: 'sketch',
        label: 'Sketch Templates',
      },
      {
        value: 'docs',
        label: 'Design Documentation',
      },
    ],
  },
]

</script>

<!--下拉栏样式Cascader 级联选择器-->
<style>
/*位置：margin: 20rem是上下左右距离*/
/*margin-top: 90px;是上外边距*/
/*margin-bottom: 20px;是下外边距*/
/*margin-left:20px;是左外边距*/
/*margin-right:20px;是右外边距*/
.example-block {
  /*容器大小*/
  height: 500px;
  background: #e0e8fc;
  /*给其父元素添加display: flex*/
  /*实现子元素水平居中*/
  /*给其父元素添加display: flex;弹性布局内容对齐（justify-content）属性：应用在弹性容器上，即父元素上。把弹性项沿着弹性容器的主轴线（main axis）对齐，如果横向有多个div，则X轴就为当前弹性容器的主轴线，Y轴为副轴线，反之亦然。*/
  display: flex;
  justify-content: center;
  align-items: center;
  /*margin: 20rem;*/
  /*margin-top: 50px;*/
  /*margin-bottom: 20px;*/
  /*margin-left:124px;*/
  /*margin-right:20px;*/
}
.example-block0{
  height: 95%;
  width: 95%;
  background: #e0e8fc;
}
.example-block1{
  height: 90px;
  width: auto;
  background: #e0e8fc;
}
.example-block2{
  /*容器大小*/
  height: 400px;
  /*给其父元素添加display: flex*/
  /*实现子元素水平居中*/
  /*给其父元素添加display: flex;弹性布局内容对齐（justify-content）属性：应用在弹性容器上，即父元素上。把弹性项沿着弹性容器的主轴线（main axis）对齐，如果横向有多个div，则X轴就为当前弹性容器的主轴线，Y轴为副轴线，反之亦然。*/
  display: flex;
  justify-content: center;
  align-items: center;
}
.example-demonstration {
  margin: 1rem;
}
.css-img24{
  width: 173.33px;
  height: 111.33px;
}
.css-img31{
  width: 172.77px;
  height: 113.33px;
}
.css-img32{
  width: 207.33px;
  height: 86.66px;
}
.css-img33{
  width: 665px;
  height: 81.33px;
}
.css-img34{
  width: 420px;
  height: 113.33px;
}
.css-img35{
  width: 1259px;
  height: 130px;
}
.css-img36{
  width: 173.33px;
  height: 111.33px;
}
.css-img37{
  width: 173.33px;
  height: 135px;
}
.css-img39{
  width: 1830px;
  height: 485px;
}
</style>
