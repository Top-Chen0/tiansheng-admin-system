<!--主页面-->
<template>
<!--  外层容器设置了高度和overflow实现滚动-->
  <div style="height: 1000px; overflow: scroll; position: absolute">
<!--    注意：chen-test-box已经被注解掉，整体无效-->
    <div class="chen-test-box">
      <div class="chen-test-title" v-for="fit in fits" :key="fit" >
        <el-image style="width: 50px; height: 50px;" :src="url" :fit="fit" />
        河南天盛景观规划设计院
        <!--      <span class="chen-test-title">河南天盛景观规划设计院</span>-->
      </div>
    </div>
<!--    添加页面第一条横线-->
    <div class="tableTitle"><span class="midText"></span></div>
<!--    页面下拉栏Cascader 级联选择器-->
    <div class="example-block">
      <span class="example-demonstration; el-cascader-panel; input:focus"></span>
<!--      排版思路：每层先确定最左侧的cascader再确定其它的cascader-->
<!--      横排1-1cascader-->
      <el-cascader
          v-model="value"
          :options="options"
          :props="props"
          @change="handleChange"
          placeholder="Cascader 级联选择器"
          clearable
          class="el-cascader-panel"
          style="position: relative; left: 12%; top: 15%"
      />
      <!--横排1-2cascader-->
      <el-cascader
          v-model="value"
          :options="options"
          :props="props"
          @change="handleChange"
          placeholder="Cascader 级联选择器"
          clearable
          class="el-cascader-panel"
          style="position: relative; left: 17%; top: 15%"

      />
      <el-cascader
          v-model="value"
          :options="options"
          :props="props"
          @change="handleChange"
          placeholder="Cascader 级联选择器"
          clearable
          class="el-cascader-panel"
          style="position: relative; left: 22%; top: 15%"
      />
      <el-cascader
          v-model="value"
          :options="options"
          :props="props"
          @change="handleChange"
          placeholder="Cascader 级联选择器"
          clearable
          class="el-cascader-panel"
          style="position: relative; left: 27%; top: 15%"
      />
      <el-cascader
          v-model="value"
          :options="options"
          :props="props"
          @change="handleChange"
          placeholder="Cascader 级联选择器"
          clearable
          class="el-cascader-panel"
          style="position: relative; left: 32%; top: 15%"

      />
<!--      竖排2-1cascader-->
      <el-cascader
          v-model="value"
          :options="options"
          :props="props"
          @change="handleChange"
          placeholder="Cascader 级联选择器"
          clearable
          class="el-cascader-panel"
          style="position: relative; left: -841px; top: 30%"
      />
      <!--      竖排2-2cascader-->
      <el-cascader
          v-model="value"
          :options="options"
          :props="props"
          @change="handleChange"
          placeholder="Cascader 级联选择器"
          clearable
          class="el-cascader-panel"
          style="position: relative; left: -745px; top: 30%"
      />
      <el-cascader
          v-model="value"
          :options="options"
          :props="props"
          @change="handleChange"
          placeholder="Cascader 级联选择器"
          clearable
          class="el-cascader-panel"
          style="position: relative; left: -650px; top: 30%"
      />
      <el-cascader
          v-model="value"
          :options="options"
          :props="props"
          @change="handleChange"
          placeholder="Cascader 级联选择器"
          clearable
          class="el-cascader-panel"
          style="position: relative; left: 1158px; top: 17%"
      />
      <el-cascader
          v-model="value"
          :options="options"
          :props="props"
          @change="handleChange"
          placeholder="Cascader 级联选择器"
          clearable
          class="el-cascader-panel"
          style="position: relative; left: 1255px; top: 17%"
      />
      <!--      竖排3-1cascader-->
      <el-cascader
          v-model="value"
          :options="options"
          :props="props"
          @change="handleChange"
          placeholder="Cascader 级联选择器"
          clearable
          class="el-cascader-panel"
          style="position: relative; left: -200px; top: 32%"
      />
      <!--      竖排3-2cascader-->
      <el-cascader
          v-model="value"
          :options="options"
          :props="props"
          @change="handleChange"
          placeholder="Cascader 级联选择器"
          clearable
          class="el-cascader-panel"
          style="position: relative; left: -104px; top: 32%"
      />
      <el-cascader
          v-model="value"
          :options="options"
          :props="props"
          @change="handleChange"
          placeholder="Cascader 级联选择器"
          clearable
          class="el-cascader-panel"
          style="position: relative; left: -09px; top: 32%"
      />
      <el-cascader
          v-model="value"
          :options="options"
          :props="props"
          @change="handleChange"
          placeholder="Cascader 级联选择器"
          clearable
          class="el-cascader-panel"
          style="position: relative; left: 88px; top: 32%"
      />
      <el-cascader
          v-model="value"
          :options="options"
          :props="props"
          @change="handleChange"
          placeholder="Cascader 级联选择器"
          clearable
          class="el-cascader-panel"
          style="position: relative; left: 185px; top: 32%"
      />
      <!--      竖排4-1cascader-->
      <el-cascader
          v-model="value"
          :options="options"
          :props="props"
          @change="handleChange"
          placeholder="Cascader 级联选择器"
          clearable
          class="el-cascader-panel"
          style="position: relative; left: -1270px; top: 47%"
      />
      <!--      竖排4-1cascader-->
      <el-cascader
          v-model="value"
          :options="options"
          :props="props"
          @change="handleChange"
          placeholder="Cascader 级联选择器"
          clearable
          class="el-cascader-panel"
          style="position: relative; left: 538px; top: 34%"
      />
      <el-cascader
          v-model="value"
          :options="options"
          :props="props"
          @change="handleChange"
          placeholder="Cascader 级联选择器"
          clearable
          class="el-cascader-panel"
          style="position: relative; left: 633px; top: 34%"
      />
      <el-cascader
          v-model="value"
          :options="options"
          :props="props"
          @change="handleChange"
          placeholder="Cascader 级联选择器"
          clearable
          class="el-cascader-panel"
          style="position: relative; left: 730px; top: 34%"
      />
      <el-cascader
          v-model="value"
          :options="options"
          :props="props"
          @change="handleChange"
          placeholder="Cascader 级联选择器"
          clearable
          class="el-cascader-panel"
          style="position: relative; left: 827px; top: 34%"
      />
    </div>


       <!--    添加页面第二条横线-->
    <div class="tableTitle0"><span class="midText0"></span></div>

<!--    第一个标题框筛选员工信息条件-->
    <div>
      <div class="tableTitle1"><span class="midText1" style="font-family: QingNiaoHuaGuangXingKai; background-color: #377ebb">筛选员工信息条件</span></div>
    </div>
<!--    添加筛选员工信息条件里的内容-->
    <div class="example-block1" style="background: #e0e8fc;">
<!--      1-1放字体的元素-->
      <img
          class="css-img1"
          src="@/image/1.png"
          alt="元素加载失败"
          style="position: relative; left: 300px; top: 10%"
      >
      <span class="example-block1-font">
        姓名
      </span>
<!--1-1横线元素-->
      <img
          class="css-img2"
          src="@/image/2.png"
          alt="元素加载失败"
          style="position: relative; left: 280px; top: 18%"
      >
<!--      1-2放字体的元素-->
      <img
          class="css-img1"
          src="@/image/1.png"
          alt="元素加载失败"
          style="position: relative; left: 320px; top: 10%"
      >
      <!--1-2横线元素-->
      <img
          class="css-img2"
          src="@/image/2.png"
          alt="元素加载失败"
          style="position: relative; left: 300px; top: 18%"
      >
      <!--      1-3放字体的元素-->
      <img
          class="css-img1"
          src="@/image/1.png"
          alt="元素加载失败"
          style="position: relative; left: 400px; top: 10%"
      >
      <!--1-3横线元素-->
      <img
          class="css-img2"
          src="@/image/2.png"
          alt="元素加载失败"
          style="position: relative; left: 380px; top: 18%"
      >
      <!--      1-4放字体的元素-->
      <img
          class="css-img1"
          src="@/image/1.png"
          alt="元素加载失败"
          style="position: relative; left: 420px; top: 10%"
      >
      <!--1-4横线元素-->
      <img
          class="css-img2"
          src="@/image/2.png"
          alt="元素加载失败"
          style="position: relative; left: 400px; top: 18%"
      >
      <!--      2-1放字体的元素-->
      <img
          class="css-img3"
          src="@/image/3.png"
          alt="元素加载失败"
          style="position: relative; left: -940px; top: 35%"
      >
      <!--2-1横线元素-->
      <img
          class="css-img4"
          src="@/image/4.png"
          alt="元素加载失败"
          style="position: relative; left: -972px; top: 42.5%"
      >
      <!--      2-2放字体的元素-->
      <img
          class="css-img3"
          src="@/image/3.png"
          alt="元素加载失败"
          style="position: relative; left: -956px; top: 35%"
      >
      <!--2-2横线元素-->
      <img
          class="css-img4"
          src="@/image/4.png"
          alt="元素加载失败"
          style="position: relative; left: 748px; top: 20%"
      >
      <!--      2-3放字体的元素-->
      <img
          class="css-img3"
          src="@/image/3.png"
          alt="元素加载失败"
          style="position: relative; left: 818px; top: 12.6%"
      >
      <!--2-3横线元素-->
      <img
          class="css-img4"
          src="@/image/4.png"
          alt="元素加载失败"
          style="position: relative; left: 788px; top: 19.5%"
      >
      <!--      2-4放字体的元素-->
      <img
          class="css-img3"
          src="@/image/3.png"
          alt="元素加载失败"
          style="position: relative; left: 800px; top: 12.6%"
      >
      <!--2-4横线元素-->
      <img
          class="css-img4"
          src="@/image/4.png"
          alt="元素加载失败"
          style="position: relative; left: 767px; top: 19.5%"
      >
<!--      3-1放字体的元素-->
      <img
          class="css-img5"
          src="@/image/5.png"
          alt="元素加载失败"
          style="position: relative; left: -595px; top: 150px"
      >
      <!--3-1横线元素-->
      <img
          class="css-img6"
          src="@/image/6.png"
          alt="元素加载失败"
          style="position: relative; left: -640px; top: 44.3%"
      >
      <!--      3-2放字体的元素-->
      <img
          class="css-img5"
          src="@/image/5.png"
          alt="元素加载失败"
          style="position: relative; left: -622px; top: 150px"
      >
      <!--3-2 横线元素-->
      <img
          class="css-img6"
          src="@/image/6.png"
          alt="元素加载失败"
          style="position: relative; left: -675px; top: 44.3%"
      >
      <!--      3-3放字体的元素-->
      <img
          class="css-img5"
          src="@/image/5.png"
          alt="元素加载失败"
          style="position: relative; left: -599px; top: 150px"
      >
      <!--3-3 横线元素-->
      <img
          class="css-img6"
          src="@/image/6.png"
          alt="元素加载失败"
          style="position: relative; left: 1133px; top: 20%"
      >
      <!--      3-4放字体的元素-->
      <img
          class="css-img5"
          src="@/image/5.png"
          alt="元素加载失败"
          style="position: relative; left: 1150px; top: 60px"
      >
      <!--3-4 横线元素-->
      <img
          class="css-img6"
          src="@/image/6.png"
          alt="元素加载失败"
          style="position: relative; left: 1100px; top: 22%"
      >
    </div>
<!--    添加页面第三条横线-->
    <div class="tableTitle0"><span class="midText0"></span></div>
<!--第二个标题框员工信息详单-->
    <div>
      <div class="tableTitle1"><span class="midText1" style="font-family: QingNiaoHuaGuangXingKai; background-color: #377ebb">员工信息详单</span></div>
    </div>
    <!--    添加筛选员工信息详单里的内容-->
    <div class="example-block2" style="background: #e0e8fc;">
      <!--1-1详单框-->
      <img
          class="css-img8"
          src="@/image/8.png"
          alt="元素加载失败"
          style="position: relative; left: 100px; top: 0%"
      >
      <!--1-2详单框-->
      <img
          class="css-img8"
          src="@/image/8.png"
          alt="元素加载失败"
          style="position: relative; left: 200px; top: 0%"
      >
      <!--1-3详单框-->
      <img
          class="css-img8"
          src="@/image/8.png"
          alt="元素加载失败"
          style="position: relative; left: 300px; top: 0%"
      >
      <!--2-1详单框-->
      <img
          class="css-img8"
          src="@/image/8.png"
          alt="元素加载失败"
          style="position: relative; left: 370px; top: -16%"
      >
      <!--2-2详单框-->
      <img
          class="css-img8"
          src="@/image/8.png"
          alt="元素加载失败"
          style="position: relative; left: 490px; top: -16%"
      >
    </div>
    <!--    添加页面第三条横线-->
    <div class="tableTitle0"><span class="midText0"></span></div>
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
    // 'https://fuss10.elemecdn.com/e/5d/4a731a90594a4af544c0c25941171jpeg.jpeg'

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
  background: #e0e8fc;
  /*margin: 20rem;*/
  /*margin-top: 50px;*/
  /*margin-bottom: 20px;*/
  /*margin-left:124px;*/
  /*margin-right:20px;*/
}
.example-block1 {
  /*容器大小*/
  height: 370px;
  background: #f94949;
  /*置于底层*/
  z-index: -1;
}
.example-block1-font{
  font-family: QingNiaoHuaGuangXingKai;
  background-color: #377ebb;
}
.example-block2 {
  /*容器大小*/
  height: 800px;
  background: #f94949;
}
.example-demonstration {
  margin: 1rem;
}
/*级联选择器背景颜色*/
.el-cascader-panel{
  background: #f9c449 !important;
}
.css-img1{
  width: 153px;
  height: 90px;
}
.css-img2{
  width: 153px;
  height: 40px;
}
.css-img3{
  width: 170px;
  height: 90px;
}
.css-img4{
  width: 178px;
  height: 40px;
}
.css-img5{
  width: 185px;
  height: 90px;
}
.css-img6{
  width: 175px;
  height: 90px;
}
.css-img8{
  width: 487.27px;
  height: 487.27px;
}



</style>
