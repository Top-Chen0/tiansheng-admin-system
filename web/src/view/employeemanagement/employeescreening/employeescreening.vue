<!--员工信息录入页面-->
<template>
  <!--  外层容器设置了高度和overflow实现滚动-->
  <div style="height: 1000px; overflow: scroll">
    <!--    注意：chen-test-box已经被注解掉，整体无效-->
    <div class="chen-test-box">
      <div class="chen-test-title" v-for="fit in fits" :key="fit" >
        <el-image style="width: 60px; height: 60px;" :src="url" :fit="fit" />
        河南天盛景观规划设计院
        <!--      <span class="chen-test-title">河南天盛景观规划设计院</span>-->
      </div>
    </div>
    <!--    第一个标题框筛选员工信息条件-->
    <div>
      <div class="tableTitle2"><span class="midText2" style="font-family: QingNiaoHuaGuangXingKai; background-color: #377ebb">员工信息录入</span></div>
    </div>
    <!--    员工信息录入页面-->
    <div class="example-block" style="margin-top: 25px;border:5px solid rgb(55,126,187);">
      <span class="example-demonstration; el-cascader-panel; input:focus"></span>

    </div>


    <!--    员工信息查询页面-->
    <div class="example-block" style="background: #f94949;margin-top: 110px;border:5px solid rgb(55,126,187);">

    </div>
    <!--    添加页面第三条横线-->
    <div class="tableTitle0"><span class="midText0">这里是横线</span></div>
    <!--最底栏-->
    <div style="width: 100%;height: 110px;background-color: #377ebb;margin-top: 10px;">
    </div>

  </div>
</template>

<!--图像加载-->
<script lang="ts" setup>
const fits = [ 'cover' ]
const url =
    'https://fuss10.elemecdn.com/e/5d/4a731a90594a4af544c0c25941171jpeg.jpeg'

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
  height: 300px;
  background: #f94949;

  /*margin: 20rem;*/
  /*margin-top: 50px;*/
  /*margin-bottom: 20px;*/
  /*margin-left:124px;*/
  /*margin-right:20px;*/
}
.example-demonstration {
  margin: 1rem;
}
/*级联选择器背景颜色*/
.el-cascader-panel{
  background: #f9c449 !important;
}




</style>
