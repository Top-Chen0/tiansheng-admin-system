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
      <div class="tableTitle2"><span class="midText2" style="font-family: QingNiaoHuaGuangXingKai; background-color: #377ebb">员工信息变更</span></div>
    </div>
    <!--    员工信息变更页面-->
    <div class="example-block" style="margin-top: 25px">
      <div class="example-block0">
        <span class="example-demonstration; el-cascader-panel; input:focus"></span>
        <!--1-1小标题-->
        <img
            class="css-img21"
            src="@/image/21.png"
            alt="元素加载失败"
            style="position: relative; left: 0px; top: -30%"
        >
        <!--        1-1输入框-->
        <img
            class="css-img22"
            src="@/image/22.png"
            alt="元素加载失败"
            style="position: relative; left: -10px; top: -30%"
        >
        <!--1-2小标题-->
        <img
            class="css-img21"
            src="@/image/21.png"
            alt="元素加载失败"
            style="position: relative; left: 0px; top: -30%"
        >
        <!--        1-2输入框-->
        <img
            class="css-img22"
            src="@/image/22.png"
            alt="元素加载失败"
            style="position: relative; left: -10px; top: -30%"
        >
        <!--1-3小标题-->
        <img
            class="css-img21"
            src="@/image/21.png"
            alt="元素加载失败"
            style="position: relative; left: 0px; top: -30%"
        >
        <!--        1-3输入框-->
        <img
            class="css-img22"
            src="@/image/22.png"
            alt="元素加载失败"
            style="position: relative; left: -10px; top: -30%"
        >
        <!--1-4小标题-->
        <img
            class="css-img21"
            src="@/image/21.png"
            alt="元素加载失败"
            style="position: relative; left: 0px; top: -30%"
        >
        <!--        1-4输入框-->
        <img
            class="css-img22"
            src="@/image/22.png"
            alt="元素加载失败"
            style="position: relative; left: -10px; top: -30%"
        >
        <!--1-5小标题-->
        <img
            class="css-img21"
            src="@/image/21.png"
            alt="元素加载失败"
            style="position: relative; left: 0px; top: -30%"
        >
        <!--        1-5输入框-->
        <img
            class="css-img22"
            src="@/image/22.png"
            alt="元素加载失败"
            style="position: relative; left: -10px; top: -30%"
        >
        <!--1-6头像框-->
        <img
            class="css-img89"
            src="@/image/89.png"
            alt="元素加载失败"
            style="position: relative; left: 0px; top: 0%"
        >
        <!--2-1小标题-->
        <img
            class="css-img21"
            src="@/image/21.png"
            alt="元素加载失败"
            style="position: relative; left: 0px; top: -75%"
        >
        <!--        2-1输入框-->
        <img
            class="css-img22"
            src="@/image/22.png"
            alt="元素加载失败"
            style="position: relative; left: -10px; top: -75%"
        >
        <!--2-2小标题-->
        <img
            class="css-img21"
            src="@/image/21.png"
            alt="元素加载失败"
            style="position: relative; left: 0px; top: -75%"
        >
        <!--        2-2输入框-->
        <img
            class="css-img22"
            src="@/image/22.png"
            alt="元素加载失败"
            style="position: relative; left: -10px; top: -75%"
        >
        <!--2-3小标题-->
        <img
            class="css-img21"
            src="@/image/21.png"
            alt="元素加载失败"
            style="position: relative; left: 0px; top: -75%"
        >
        <!--        2-3输入框-->
        <img
            class="css-img22"
            src="@/image/22.png"
            alt="元素加载失败"
            style="position: relative; left: -10px; top: -75%"
        >
        <!--2-4小标题-->
        <img
            class="css-img21"
            src="@/image/21.png"
            alt="元素加载失败"
            style="position: relative; left: 0px; top: -75%"
        >
        <!--        2-4输入框-->
        <img
            class="css-img22"
            src="@/image/22.png"
            alt="元素加载失败"
            style="position: relative; left: -10px; top: -75%"
        >
        <!--2-5小标题-->
        <img
            class="css-img21"
            src="@/image/21.png"
            alt="元素加载失败"
            style="position: relative; left: 0px; top: -75%"
        >
        <!--        2-5输入框-->
        <img
            class="css-img22"
            src="@/image/22.png"
            alt="元素加载失败"
            style="position: relative; left: -10px; top: -75%"
        >
        <!--3-1小标题-->
        <img
            class="css-img21"
            src="@/image/21.png"
            alt="元素加载失败"
            style="position: relative; left: -1465px; top: -55%"
        >
        <!--        3-1输入框-->
        <img
            class="css-img22"
            src="@/image/22.png"
            alt="元素加载失败"
            style="position: relative; left: -1475px; top: -55%"
        >
        <!--3-2小标题-->
        <img
            class="css-img21"
            src="@/image/21.png"
            alt="元素加载失败"
            style="position: relative; left: 293px; top: -83%"
        >
        <!--        3-2输入框-->
        <img
            class="css-img22"
            src="@/image/22.png"
            alt="元素加载失败"
            style="position: relative; left: 283px; top: -83%"
        >
        <!--3-3小标题-->
        <img
            class="css-img21"
            src="@/image/21.png"
            alt="元素加载失败"
            style="position: relative; left: 293px; top: -83%"
        >
        <!--        3-3输入框-->
        <img
            class="css-img22"
            src="@/image/22.png"
            alt="元素加载失败"
            style="position: relative; left: 283px; top: -83%"
        >
        <!--3-4小标题-->
        <img
            class="css-img21"
            src="@/image/21.png"
            alt="元素加载失败"
            style="position: relative; left: 293px; top: -83%"
        >
        <!--        3-4输入框-->
        <img
            class="css-img22"
            src="@/image/22.png"
            alt="元素加载失败"
            style="position: relative; left: 283px; top: -83%"
        >
        <!--3-5小标题-->
        <img
            class="css-img21"
            src="@/image/21.png"
            alt="元素加载失败"
            style="position: relative; left: 293px; top: -83%"
        >
        <!--        3-5输入框-->
        <img
            class="css-img22"
            src="@/image/22.png"
            alt="元素加载失败"
            style="position: relative; left: 283px; top: -83%"
        >
        <!--4-1小标题-->
        <img
            class="css-img21"
            src="@/image/21.png"
            alt="元素加载失败"
            style="position: relative; left: -1172px; top: -63%"
        >
        <!--        4-1输入框-->
        <img
            class="css-img22"
            src="@/image/22.png"
            alt="元素加载失败"
            style="position: relative; left: -1182px; top: -63%"
        >
        <!--4-2小标题-->
        <img
            class="css-img21"
            src="@/image/21.png"
            alt="元素加载失败"
            style="position: relative; left: -1172px; top: -63%"
        >
        <!--        4-2输入框-->
        <img
            class="css-img22"
            src="@/image/22.png"
            alt="元素加载失败"
            style="position: relative; left: -1182px; top: -63%"
        >
        <!--4-3小标题-->
        <img
            class="css-img21"
            src="@/image/21.png"
            alt="元素加载失败"
            style="position: relative; left: 585px; top: -90%"
        >
        <!--        4-3输入框-->
        <img
            class="css-img22"
            src="@/image/22.png"
            alt="元素加载失败"
            style="position: relative; left: 576px; top: -90%"
        >
        <!--4-4小标题-->
        <img
            class="css-img21"
            src="@/image/21.png"
            alt="元素加载失败"
            style="position: relative; left: 585px; top: -90%"
        >
        <!--        4-4输入框-->
        <img
            class="css-img22"
            src="@/image/22.png"
            alt="元素加载失败"
            style="position: relative; left: 576px; top: -90%"
        >
        <!--4-5小标题-->
        <img
            class="css-img21"
            src="@/image/21.png"
            alt="元素加载失败"
            style="position: relative; left: 585px; top: -90%"
        >
        <!--        4-5输入框-->
        <img
            class="css-img22"
            src="@/image/22.png"
            alt="元素加载失败"
            style="position: relative; left: 576px; top: -90%"
        >
      </div>
    </div>
    <div class="example-block1">
      <!--        添加按钮样式-->
      <img
          class="css-img24"
          src="@/image/24.png"
          alt="元素加载失败"
          style="position: relative; left: 1600px; top: -15%"
      >
    </div>
    <!--    添加页面第二条横线-->
    <div class="tableTitle0"><span class="midText0"></span></div>
    <!--    添加筛选员工信息条件里的内容-->
    <div class="example-block2" style="background: #e0e8fc;margin-top: 30px">
      <div class="example-block0">
        <!--1-1小标题-->
        <img
            class="css-img25"
            src="@/image/25.png"
            alt="元素加载失败"
            style="position: relative; left: 150px; top: 0%"
        >
        <!--        1-1输入框-->
        <img
            class="css-img26"
            src="@/image/26.png"
            alt="元素加载失败"
            style="position: relative; left: 110px; top: 0%"
        >
        <!--1-2小标题-->
        <img
            class="css-img25"
            src="@/image/25.png"
            alt="元素加载失败"
            style="position: relative; left: 270px; top: 0%"
        >
        <!--        1-2输入框-->
        <img
            class="css-img26"
            src="@/image/26.png"
            alt="元素加载失败"
            style="position: relative; left: 230px; top: 0%"
        >
        <br>
        <!--2-1小标题-->
        <img
            class="css-img25"
            src="@/image/25.png"
            alt="元素加载失败"
            style="position: relative; left: 150px; top: -15%"
        >
        <!--        2-1输入框-->
        <img
            class="css-img26"
            src="@/image/26.png"
            alt="元素加载失败"
            style="position: relative; left: 110px; top: -15%"
        >
        <!--2-2小标题-->
        <img
            class="css-img25"
            src="@/image/25.png"
            alt="元素加载失败"
            style="position: relative; left: 270px; top: -15%"
        >
        <!--        2-2输入框-->
        <img
            class="css-img26"
            src="@/image/26.png"
            alt="元素加载失败"
            style="position: relative; left: 230px; top: -15%"
        >
        <br>
        <!--3-1小标题-->
        <img
            class="css-img25"
            src="@/image/25.png"
            alt="元素加载失败"
            style="position: relative; left: 150px; top: -30%"
        >
        <!--        3-1输入框-->
        <img
            class="css-img26"
            src="@/image/26.png"
            alt="元素加载失败"
            style="position: relative; left: 110px; top: -30%"
        >
        <!--3-2小标题-->
        <img
            class="css-img25"
            src="@/image/25.png"
            alt="元素加载失败"
            style="position: relative; left: 270px; top: -30%"
        >
        <!--        3-2输入框-->
        <img
            class="css-img26"
            src="@/image/26.png"
            alt="元素加载失败"
            style="position: relative; left: 230px; top: -30%"
        >
        <br>
        <!--4-1小标题-->
        <img
            class="css-img25"
            src="@/image/25.png"
            alt="元素加载失败"
            style="position: relative; left: 150px; top: -45%"
        >
        <!--        4-1输入框-->
        <img
            class="css-img26"
            src="@/image/26.png"
            alt="元素加载失败"
            style="position: relative; left: 110px; top: -45%"
        >
        <!--4-2小标题-->
        <img
            class="css-img25"
            src="@/image/25.png"
            alt="元素加载失败"
            style="position: relative; left: 270px; top: -45%"
        >
        <!--        4-2输入框-->
        <img
            class="css-img26"
            src="@/image/26.png"
            alt="元素加载失败"
            style="position: relative; left: 230px; top: -45%"
        >
        <br>
        <!--5-1小标题-->
        <img
            class="css-img25"
            src="@/image/25.png"
            alt="元素加载失败"
            style="position: relative; left: 150px; top: -60%"
        >
        <!--        5-1输入框-->
        <img
            class="css-img26"
            src="@/image/26.png"
            alt="元素加载失败"
            style="position: relative; left: 110px; top: -60%"
        >
        <!--5-2小标题-->
        <img
            class="css-img25"
            src="@/image/25.png"
            alt="元素加载失败"
            style="position: relative; left: 270px; top: -60%"
        >
        <!--        5-2输入框-->
        <img
            class="css-img26"
            src="@/image/26.png"
            alt="元素加载失败"
            style="position: relative; left: 230px; top: -60%"
        >
      </div>
    </div>
    <div class="example-block1">
      <!--        添加按钮样式-->
      <img
          class="css-img24"
          src="@/image/24.png"
          alt="元素加载失败"
          style="position: relative; left: 1600px; top: -40%"
      >
    </div>
    <!--    添加页面第三条横线-->
    <div class="tableTitle0"><span class="midText0"></span></div>
    <!--第二个标题框员工信息详单-->
    <div>
      <div class="tableTitle1"><span class="midText1" style="font-family: QingNiaoHuaGuangXingKai; background-color: #377ebb">员工信息详单</span></div>
    </div>
    <!--    添加筛选员工信息详单里的内容-->
    <div class="example-block" style="background: #e0e8fc;">
      <div class="example-block0">
        <!--1-1小标题-->
        <img
            class="css-img27"
            src="@/image/27.png"
            alt="元素加载失败"
            style="position: relative; left: 380px; top: 0%"
        >
        <!--        1-1输入框-->
        <img
            class="css-img33"
            src="@/image/33.png"
            alt="元素加载失败"
            style="position: relative; left: 380px; top: 0%"
        >
        <br>
        <!--2-1小标题-->
        <img
            class="css-img27"
            src="@/image/27.png"
            alt="元素加载失败"
            style="position: relative; left: 380px; top: -15%"
        >
        <!--        2-1输入框-->
        <img
            class="css-img33"
            src="@/image/33.png"
            alt="元素加载失败"
            style="position: relative; left: 380px; top: -15%"
        >
        <br>
        <!--3-1小标题-->
        <img
            class="css-img27"
            src="@/image/27.png"
            alt="元素加载失败"
            style="position: relative; left: 380px; top: -30%"
        >
        <!--        3-1输入框-->
        <img
            class="css-img33"
            src="@/image/33.png"
            alt="元素加载失败"
            style="position: relative; left: 380px; top: -30%"
        >
        <br>
        <!--4-1小标题-->
        <img
            class="css-img27"
            src="@/image/27.png"
            alt="元素加载失败"
            style="position: relative; left: 380px; top: -45%"
        >
        <!--        4-1输入框-->
        <img
            class="css-img33"
            src="@/image/33.png"
            alt="元素加载失败"
            style="position: relative; left: 380px; top: -45%"
        >
      </div>
    </div>
    <div class="example-block1">
      <!--        添加按钮样式-->
      <img
          class="css-img24"
          src="@/image/24.png"
          alt="元素加载失败"
          style="position: relative; left: 1600px; top: -60%"
      >
    </div>
    <!--    添加页面第三条横线-->
    <div class="tableTitle0"><span class="midText0"></span></div>
    <div class="example-block1">
      <!--        主页按钮样式-->
      <img
          class="css-img24"
          src="@/image/24.png"
          alt="元素加载失败"
          style="position: relative; left: 0px; top: 0%"
      >
      <!--        取消按钮样式-->
      <img
          class="css-img24"
          src="@/image/24.png"
          alt="元素加载失败"
          style="position: relative; left: 1255px; top: 0%"
      >
      <!--        确定按钮样式-->
      <img
          class="css-img24"
          src="@/image/24.png"
          alt="元素加载失败"
          style="position: relative; left: 1255px; top: 0%"
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
  height: 350px;
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
.css-img21{
  width: 153px;
  height: 90px;
}
.css-img22{
  width: 140px;
  height: 65px;
}
.css-img89{
  width: 270.71px;
  height: 333.57px;
}
.css-img24{
  width: 173.33px;
  height: 111.33px;
}
.css-img25{
  width: 230px;
  height: 118.66px;
}
.css-img26{
  width: 456.92px;
  height: 76.92px;
}
.css-img27{
  width: 299.33px;
  height: 114px;
}
.css-img33{
  width: 800px;
  height: 85px;
}
</style>
