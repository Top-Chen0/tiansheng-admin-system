<!--主页面-->
<template >
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
          style="position: relative; left: 12.1%; top: 15%"
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
          style="position: relative; left: 17.1%; top: 15%"

      />
      <el-cascader
          v-model="value"
          :options="options"
          :props="props"
          @change="handleChange"
          placeholder="Cascader 级联选择器"
          clearable
          class="el-cascader-panel"
          style="position: relative; left: 22.1%; top: 15%"
      />
      <el-cascader
          v-model="value"
          :options="options"
          :props="props"
          @change="handleChange"
          placeholder="Cascader 级联选择器"
          clearable
          class="el-cascader-panel"
          style="position: relative; left: 27.1%; top: 15%"
      />
      <el-cascader
          v-model="value"
          :options="options"
          :props="props"
          @change="handleChange"
          placeholder="Cascader 级联选择器"
          clearable
          class="el-cascader-panel"
          style="position: relative; left: 32.1%; top: 15%"

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
      <span class="example-block1-font" style="left: 360px;top: 523px">
        工号
        <el-input v-model="searchInfo.worknumber" class="input-el-input"/>
      </span>
      <!--1-1横线元素-->
      <img
          class="css-img2"
          src="@/image/2.png"
          alt="元素加载失败"
          style="position: relative; left: 280px; top: 18%"
      >
      <span class="example-block1-font" style="left: 681px;top: 523px">
        姓名
        <el-input v-model="searchInfo.name" class="input-el-input"/>
      </span>
      <!--      1-2放字体的元素-->
      <img
          class="css-img1"
          src="@/image/1.png"
          alt="元素加载失败"
          style="position: relative; left: 320px; top: 10%"
      >
      <span class="example-block1-font" style="left: 1065px;top: 523px">
        年龄
        <el-input v-model="searchInfo.age" class="input-el-input"/>
      </span>
      <!--1-2横线元素-->
      <img
          class="css-img2"
          src="@/image/2.png"
          alt="元素加载失败"
          style="position: relative; left: 300px; top: 18%"
      >
      <span class="example-block1-font" style="left: 1390px;top: 523px">
        性别
        <el-select v-model="searchInfo.sex" clearable class="input-el-select">
            <el-option v-for="(item,key) in genderOptions" :key="key" :label="item.label" :value="item.value" />
        </el-select>
      </span>
      <!--      1-3放字体的元素-->
      <img
          class="css-img1"
          src="@/image/1.png"
          alt="元素加载失败"
          style="position: relative; left: 400px; top: 10%"
      >
      <!--      <span class="example-block1-font" style="left: 340px;top: 615px;letter-spacing:-14px;">-->
      <!--        生日月份-->
      <!--      </span>-->
      <span class="example-block1-font" style="left: 356px;top: 617px">
        部门
        <el-select v-model="searchInfo.department" placeholder="请选择" clearable class="input-el-select">
            <el-option v-for="(item,key) in departmentOptions" :key="key" :label="item.label" :value="item.value" />
        </el-select>
      </span>
      <!--1-3横线元素-->
      <img
          class="css-img2"
          src="@/image/2.png"
          alt="元素加载失败"
          style="position: relative; left: 380px; top: 18%"
      >
      <span class="example-block1-font" style="left: 681px;top: 617px">
        岗位
          <el-select v-model="searchInfo.job" placeholder="请选择" clearable class="input-el-select">
            <el-option v-for="(item,key) in jobOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
      </span>
      <!--      1-4放字体的元素-->
      <img
          class="css-img1"
          src="@/image/1.png"
          alt="元素加载失败"
          style="position: relative; left: 420px; top: 10%"
      >
      <span class="example-block1-font" style="left: 1065px;top: 622px">
        学历
         <el-select v-model="searchInfo.degree" placeholder="请选择" clearable class="input-el-select">
            <el-option v-for="(item,key) in degreeOptions" :key="key" :label="item.label" :value="item.value" />
         </el-select>
      </span>
      <!--1-4横线元素-->
      <img
          class="css-img2"
          src="@/image/2.png"
          alt="元素加载失败"
          style="position: relative; left: 400px; top: 18%"
      >
      <span class="example-block1-font" style="left: 1383px;top: 622px;letter-spacing:-14px;">
        生日月份
        <el-select v-model="searchInfo.birth_month" placeholder="请选择" clearable class="input-el-select">
            <el-option v-for="(item,key) in birth_monthOptions" :key="key" :label="item.label" :value="item.value" />
        </el-select>
      </span>
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
      <el-form-item>
        <div style="text-align: right">

          <el-button size="small" type="primary" icon="search" @click="onSubmit" >查询</el-button>

          <el-button size="small" icon="refresh" @click="onReset">重置</el-button>

        </div>
      </el-form-item>
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
<!--      数据表格1-->
      <el-table
          ref="multipleTable"
          :data="tableData"
          row-key="ID"
          style="position: absolute;width: 19%; height: 25.6%; left: 161px; top: 106%;background-color: #ffffff"
      >
        <el-table-column align="center" label="工号" prop="worknumber" width="60" />
        <el-table-column align="center" label="姓名" prop="name" width="70" />
        <el-table-column align="center" label="年龄" prop="age" width="60" />
        <el-table-column align="center" label="性别" prop="sex" width="55">
          <template #default="scope">
            {{ filterDict(scope.row.sex,genderOptions) }}
          </template>
        </el-table-column>
        <el-table-column align="center" label="岗位" prop="job" width="120">
          <template #default="scope">
            {{ filterDict(scope.row.job,jobOptions) }}
          </template>
        </el-table-column>
      </el-table>
<!--数据表格2-->
      <el-table
          ref="multipleTable"
          :data="tableData"
          row-key="ID"
          style="position: absolute;width: 19%; height: 10%; left: 161px; top: 115.9%;background-color: #ffffff"

      >
        <el-table-column align="center" label="部门" prop="department" width="190 ">
          <template #default="scope">
            {{ filterDict(scope.row.department,departmentOptions) }}
          </template>
        </el-table-column>
        <el-table-column align="center" label="部门经理" prop="departmentmanager" width="100">
          <template #default="scope">
            {{ filterDict(scope.row.departmentmanager,departmentmanagerOptions) }}
          </template>
        </el-table-column>
      </el-table>
      <!--1-2详单框-->
      <img
          class="css-img8"
          src="@/image/8.png"
          alt="元素加载失败"
          style="position: relative; left: 200px; top: 0%"
      >
<!--      数据表格1-->
      <el-table
          ref="multipleTable"
          :data="tableData"
          row-key="ID"
          style="position: absolute;width: 19%; height: 25.6%; left: 751px; top: 106%;background-color: #ffffff"
      >
        <el-table-column align="center" label="入职日期" prop="dateofentry" width="120"/>
        <el-table-column align="center" label="籍贯" prop="origin" width="120" />
        <el-table-column align="center" label="户口所在地" prop="householdregistration" width="120" />
      </el-table>
<!--      数据表格2-->
      <el-table
          ref="multipleTable"
          :data="tableData"
          row-key="ID"
          style="position: absolute;width: 19%; height: 10%; left: 751px; top: 115.9%;background-color: #ffffff"
      >
        <el-table-column align="center" label="身份证详细地址" prop="idaddress" width="360" />
      </el-table>
      <!--1-3详单框-->
      <img
          class="css-img8"
          src="@/image/8.png"
          alt="元素加载失败"
          style="position: relative; left: 300px; top: 0%"
      >
<!--      数据表格1-->
      <el-table
          ref="multipleTable"
          :data="tableData"
          row-key="ID"
          style="position: absolute;width: 19%; height: 25.6%; left: 1341px; top: 106%;background-color: #ffffff"
      >
        <el-table-column align="center" label="民族" prop="ethnicgroup" width="120" />
        <el-table-column align="center" label="婚姻状况" prop="marriage" width="120" >
            <template #default="scope">
            {{ filterDict(scope.row.marriage,marriageOptions) }}
            </template>
        </el-table-column>
        <el-table-column align="center" label="生日月份" prop="birth_month" width="120">
            <template #default="scope">
            {{ filterDict(scope.row.birth_month,birth_monthOptions) }}
            </template>
        </el-table-column>
      </el-table>
<!--      数据表格2-->
      <el-table
          ref="multipleTable"
          :data="tableData"
          row-key="ID"
          style="position: absolute;width: 19%; height: 10%; left: 1341px; top: 115.9%;background-color: #ffffff"
      >
        <el-table-column align="center" label="现住详细地址" prop="address" width="360" />
      </el-table>
      <!--2-1详单框-->
      <img
          class="css-img8"
          src="@/image/8.png"
          alt="元素加载失败"
          style="position: relative; left: 370px; top: -16%"
      >
<!--      数据表1-->
      <el-table
          ref="multipleTable"
          :data="tableData"
          row-key="ID"
          style="position: absolute;width: 19%; height: 25.6%; left: 432px; top: 142%;background-color: #ffffff"
      >
        <el-table-column align="center" label="政治面貌" prop="politicalface" width="120">
            <template #default="scope">
            {{ filterDict(scope.row.politicalface,politicalfaceOptions) }}
            </template>
        </el-table-column>
        <el-table-column align="center" label="手机号码" prop="number" width="120" />
        <el-table-column align="center" label="邮编" prop="zip" width="120" />
      </el-table>
<!--      数据表2-->
      <el-table
          ref="multipleTable"
          :data="tableData"
          row-key="ID"
          style="position: absolute;width: 19%; height: 10%; left: 432px; top: 151.9%;background-color: #ffffff"
      >
        <el-table-column align="center" label="学历" prop="degree" width="120">
            <template #default="scope">
            {{ filterDict(scope.row.degree,degreeOptions) }}
            </template>
        </el-table-column>
        <el-table-column align="center" label="毕业院校" prop="graduatingschool" width="120" />
        <el-table-column align="center" label="毕业时间" prop="graduationtime" width="120" />
      </el-table>
      <!--2-2详单框-->
      <img
          class="css-img8"
          src="@/image/8.png"
          alt="元素加载失败"
          style="position: relative; left: 490px; top: -16%"
      >
<!--      数据表1-->
      <el-table
          ref="multipleTable"
          :data="tableData"
          row-key="ID"
          style="position: absolute;width: 19%; height: 25.6%; left: 1038px; top: 142%;background-color: #ffffff"
      >
        <el-table-column align="center" label="专业" prop="specialized" width="120" />
        <el-table-column align="center" label="特长" prop="specialty" width="120" />
        <el-table-column align="center" label="个人技能" prop="personalskills" width="120" />
      </el-table>
<!--数据表2-->
      <el-table
          ref="multipleTable"
          :data="tableData"
          row-key="ID"
          style="position: absolute;width: 19%; height: 10%; left: 1038px; top: 151.9%;background-color: #ffffff"
      >
        <el-table-column align="center" label="专业职称" prop="professionaltitle" width="120" />
      </el-table>
    </div>
    <!--    添加页面第三条横线-->
    <div class="tableTitle0"><span class="midText0"></span></div>
    <!--最底栏-->
    <div style="width: 100%;height: 110px;background-color: #377ebb;margin-top: 10px;">
    </div>

  </div>
</template>

<script>
export default {
  name: 'EmployeeStructui'
}
</script>

<script setup>
import {
  createEmployeeStructui,
  deleteEmployeeStructui,
  deleteEmployeeStructuiByIds,
  updateEmployeeStructui,
  findEmployeeStructui,
  getEmployeeStructuiList
} from '@/api/employeeStructui'
// 设置logo
const fits = [ 'cover' ]
//logo地址
const url =
    '/src/assets/logo.png'
    // 'https://fuss10.elemecdn.com/e/5d/4a731a90594a4af544c0c25941171jpeg.jpeg'

const value = ref([])

const props = {
  expandTrigger: 'hover',
}

const handleChange = (value) => {
  console.log(value)
}

//导入图像包
import ChooseImg from '@/components/chooseImg/index.vue'

// 全量引入格式化工具 请按需保留
import { getDictFunc, formatDate, formatBoolean, filterDict } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref } from 'vue'
import { Search } from '@element-plus/icons-vue'
import { useUserStore } from '@/pinia/modules/user'

// 自动化生成的字典（可能为空）以及字段
const jobOptions = ref([])
const birth_monthOptions = ref([])
const marriageOptions = ref([])
const genderOptions = ref([])//和字典引入有关
const politicalfaceOptions = ref([])
const degreeOptions = ref([])
const departmentmanagerOptions = ref([])
const socialsecurityOptions = ref([])
const departmentOptions = ref([])
const formData = ref({
  worknumber: 0,
  name: '',
  age: 0,
  sex: undefined,
  department: undefined,
  job: undefined,
  dateofentry: new Date(),
  socialsecuritynumber: '',
  dateofbirth: new Date(),
  birth_month: undefined,
  origin: '',
  householdregistration: '',
  idaddress: '',
  address: '',
  ethnicgroup: '',
  marriage: undefined,
  politicalface: undefined,
  emergencycontacts: '',
  ec_number: 0,
  number: 0,
  ec_relationship: '',
  degree: undefined,
  graduatingschool: '',
  graduationtime: new Date(),
  specialized: '',
  specialty: '',
  personalskills: '',
  professionaltitle: '',
  zip: 0,
  departmentmanager: undefined,
  socialsecurity: undefined,
})

// =========== 表格控制部分 ===========
const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])
const searchInfo = ref({})

// 重置
const onReset = () => {
  searchInfo.value = {}

}

// 搜索
const onSubmit = () => {
  page.value = 1
  pageSize.value = 10
  if (searchInfo.value.age === ""){
    searchInfo.value.age=null
  }
  getTableData()
}

// 分页
const handleSizeChange = (val) => {
  pageSize.value = val
  getTableData()
}

// 修改页面容量
const handleCurrentChange = (val) => {
  page.value = val
  getTableData()
}

// 查询
const getTableData = async() => {
  const table = await getEmployeeStructuiList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
  if (table.code === 0) {
    tableData.value = table.data.list
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize
  }
}

//getTableData()//控制搜索页不自动加载数据

// ============== 表格控制部分结束 ===============

// 获取需要的字典 可能为空 按需保留
const setOptions = async () =>{
  departmentmanagerOptions.value = await getDictFunc('departmentmanager')
  genderOptions.value = await getDictFunc('gender')
  jobOptions.value = await getDictFunc('job')
  birth_monthOptions.value = await getDictFunc('birth_month')
  marriageOptions.value = await getDictFunc('marriage')
  politicalfaceOptions.value = await getDictFunc('politicalface')
  degreeOptions.value = await getDictFunc('degree')
  departmentmanagerOptions.value = await getDictFunc('departmentmanager')
  socialsecurityOptions.value = await getDictFunc('socialsecurity')
  departmentOptions.value = await getDictFunc('department')
}

// 获取需要的字典 可能为空 按需保留
setOptions()


// 多选数据
const multipleSelection = ref([])
// 多选
const handleSelectionChange = (val) => {
  multipleSelection.value = val
}

// 删除行
const deleteRow = (row) => {
  ElMessageBox.confirm('确定要删除吗?', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(() => {
    deleteEmployeeStructuiFunc(row)
  })
}


// 批量删除控制标记
const deleteVisible = ref(false)

// 多选删除
const onDelete = async() => {
  const ids = []
  if (multipleSelection.value.length === 0) {
    ElMessage({
      type: 'warning',
      message: '请选择要删除的数据'
    })
    return
  }
  multipleSelection.value &&
  multipleSelection.value.map(item => {
    ids.push(item.ID)
  })
  const res = await deleteEmployeeStructuiByIds({ ids })
  if (res.code === 0) {
    ElMessage({
      type: 'success',
      message: '删除成功'
    })
    if (tableData.value.length === ids.length && page.value > 1) {
      page.value--
    }
    deleteVisible.value = false
    getTableData()
  }
}

// 行为控制标记（弹窗内部需要增还是改）
const type = ref('')

// 更新行
const updateEmployeeStructuiFunc = async(row) => {
  const res = await findEmployeeStructui({ ID: row.ID })
  type.value = 'update'
  if (res.code === 0) {
    formData.value = res.data.reemployeeStructui
    dialogFormVisible.value = true
  }
}


// 删除行
const deleteEmployeeStructuiFunc = async (row) => {
  const res = await deleteEmployeeStructui({ ID: row.ID })
  if (res.code === 0) {
    ElMessage({
      type: 'success',
      message: '删除成功'
    })
    if (tableData.value.length === 1 && page.value > 1) {
      page.value--
    }
    getTableData()
  }
}

// 弹窗控制标记
const dialogFormVisible = ref(false)

// 打开弹窗
const openDialog = () => {
  type.value = 'create'
  dialogFormVisible.value = true
}

// 关闭弹窗
const closeDialog = () => {
  dialogFormVisible.value = false
  formData.value = {
    worknumber: 0,
    name: '',
    age: 0,
    sex: undefined,
    department: undefined,
    job: undefined,
    dateofentry: new Date(),
    socialsecuritynumber: '',
    dateofbirth: new Date(),
    birth_month: undefined,
    origin: '',
    householdregistration: '',
    idaddress: '',
    address: '',
    ethnicgroup: '',
    marriage: undefined,
    politicalface: undefined,
    emergencycontacts: '',
    ec_number: 0,
    number: 0,
    ec_relationship: '',
    degree: undefined,
    graduatingschool: '',
    graduationtime: new Date(),
    specialized: '',
    specialty: '',
    personalskills: '',
    professionaltitle: '',
    zip: 0,
    departmentmanager: undefined,
    socialsecurity: undefined,
    employeeImg: '',

  }
}
// 弹窗确定
const enterDialog = async () => {
  let res
  switch (type.value) {
    case 'create':
      res = await createEmployeeStructui(formData.value)
      break
    case 'update':
      res = await updateEmployeeStructui(formData.value)
      break
    default:
      res = await createEmployeeStructui(formData.value)
      break
  }``
  if (res.code === 0) {
    ElMessage({
      type: 'success',
      message: '创建/更改成功'
    })
    closeDialog()
    getTableData()
  }
}

//头像加载代替图像加载demo
// const userStore = useUserStore()
// const path = ref(import.meta.env.VITE_BASE_API)
// const openChooseImg = () => {
//   chooseImgRef.value.open()
// }


//图像加载
const chooseImg = ref(null)
const openHeaderChange = () => {
  chooseImg.value.open()
}

// 图像加载弹窗相关
const employeeimgInfo = ref({
  employeeImg: '',
})

const userInfo = ref({
  headerImg: '',
})
const path = ref(import.meta.env.VITE_BASE_API)

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
<!--scoped用于将样式只作用于当下模块下-->
<style>
</style>

<!--头像加载代替图像加载demo-->
<!--<style lang="scss">-->
<!--.user-headpic-update {-->
<!--width: 120px;-->
<!--height: 120px;-->
<!--line-height: 120px;-->
<!--margin: 0 auto;-->
<!--display: flex;-->
<!--justify-content: center;-->
<!--border-radius: 20px;-->
<!--&:hover {-->
<!--color: #fff;-->
<!--background: linear-gradient(-->
<!--to bottom,-->
<!--rgba(255, 255, 255, 0.15) 0%,-->
<!--rgba(0, 0, 0, 0.15) 100%-->
<!--),-->
<!--radial-gradient(-->
<!--at top center,-->
<!--rgba(255, 255, 255, 0.4) 0%,-->
<!--rgba(0, 0, 0, 0.4) 120%-->
<!--)-->
<!--#989898;-->
<!--background-blend-mode: multiply, multiply;-->
<!--.update {-->
<!--color: #fff;-->
<!--}-->
<!--}-->
<!--.update {-->
<!--height: 120px;-->
<!--width: 120px;-->
<!--text-align: center;-->
<!--color: transparent;-->
<!--}-->
<!--}-->
<!--</style>-->


<!--图片加载-->
<style lang="scss">
.user-dialog {
  .header-img-box {
    width: 200px;
    height: 200px;
    border: 1px dashed #ccc;
    border-radius: 20px;
    text-align: center;
    line-height: 200px;
    cursor: pointer;
  }
  .avatar-uploader .el-upload:hover {
    border-color: #409eff;
  }
  .avatar-uploader-icon {
    border: 1px dashed #d9d9d9 !important;
    border-radius: 6px;
    font-size: 28px;
    color: #8c939d;
    width: 178px;
    height: 178px;
    line-height: 178px;
    text-align: center;
  }
  .avatar {
    width: 178px;
    height: 178px;
    display: block;
  }
}
.nickName{
  display: flex;
  justify-content: flex-start;
  align-items: center;
}
.pointer{
  cursor: pointer;
  font-size: 16px;
  margin-left: 2px;
}



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
  position: absolute;
  z-index: 1;
  /*display: inline;*/
  font-size: 35px;
  color: #FFFFFF;
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
  background: #f9c449;
}
.input-el-input{
  width: 25%;
  --el-input-bg-color: transparent;
  --el-input-border: none;
  font-size: 35px;
  font-family: QingNiaoHuaGuangXingKai;
}
.input-el-select{
  width: 50%;
  border:none;
  outline:none;
  --el-color-white: transparent;
  --el-border-base: transparent;
  --el-text-color-placeholder: transparent;
  top: -10px;
}
//数据表加边框
.table0{
  border: 3px solid #353755;
  border-right: none;
  border-left: none;
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
