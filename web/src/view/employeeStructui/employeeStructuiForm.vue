<template>
  <div>
    <div class="gva-form-box">
      <!--    :inline="true"数据并列-->
      <el-form :inline="true" :model="formData" label-position="right" label-width="90px">
       <el-form :inline="true" :model="formData" label-position="right" label-width="120px">
        <el-form-item label="工号:">
          <el-input v-model.number="formData.worknumber" clearable placeholder="请输入" style="width: 220px"/>
        </el-form-item>
        <el-form-item label="姓名:">
          <el-input v-model="formData.name" clearable placeholder="请输入" style="width: 220px"/>
        </el-form-item>
        <el-form-item label="年龄:">
          <el-input v-model.number="formData.age" clearable placeholder="请输入" style="width: 220px"/>
        </el-form-item>
        <el-form-item label="性别:">
          <el-select v-model="formData.sex" placeholder="请选择" style="width:220px" clearable>
            <el-option v-for="(item,key) in genderOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item label="部门:">
          <el-select v-model="formData.department" placeholder="请选择" clearable style="width: 220px">
            <el-option v-for="(item,key) in departmentOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item label="岗位:">
          <el-select v-model="formData.job" placeholder="请选择" clearable style="width: 220px">
            <el-option v-for="(item,key) in jobOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item label="入职日期:">
          <el-date-picker v-model="formData.dateofentry" type="date" placeholder="选择日期" clearable style="width: 220px"></el-date-picker>
        </el-form-item>
        <el-form-item label="身份证号:">
          <el-input v-model="formData.socialsecuritynumber" clearable placeholder="请输入" style="width: 220px"/>
        </el-form-item>
        <el-form-item label="出生日期:">
          <el-date-picker v-model="formData.dateofbirth" type="date" placeholder="选择日期" clearable style="width: 220px"></el-date-picker>
        </el-form-item>
        <el-form-item label="生日月份:">
          <el-select v-model="formData.birth_month" placeholder="请选择" clearable style="width: 220px">
            <el-option v-for="(item,key) in birth_monthOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item label="籍贯:">
          <el-input v-model="formData.origin" clearable placeholder="请输入" style="width: 220px"/>
        </el-form-item>
        <el-form-item label="户口所在地:">
          <el-input v-model="formData.householdregistration" clearable placeholder="请输入" style="width: 220px"/>
        </el-form-item>
        <el-form-item label="身份证详细地址:">
          <el-input v-model="formData.idaddress" clearable placeholder="请输入" style="width: 220px"/>
        </el-form-item>
        <el-form-item label="现住详细地址:">
          <el-input v-model="formData.address" clearable placeholder="请输入" style="width: 220px"/>
        </el-form-item>
        <el-form-item label="民族:">
          <el-input v-model="formData.ethnicgroup" clearable placeholder="请输入" style="width: 220px"/>
        </el-form-item>
        <el-form-item label="婚姻状况:">
          <el-select v-model="formData.marriage" placeholder="请选择" clearable style="width: 220px">
            <el-option v-for="(item,key) in marriageOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item label="政治面貌:">
          <el-select v-model="formData.politicalface" placeholder="请选择" clearable style="width: 220px">
            <el-option v-for="(item,key) in politicalfaceOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item label="紧急联系人:">
          <el-input v-model="formData.emergencycontacts" clearable placeholder="请输入" style="width: 220px"/>
        </el-form-item>
        <el-form-item label="紧急联系人电话:">
          <el-input v-model.number="formData.ec_number" clearable placeholder="请输入" style="width: 220px"/>
        </el-form-item>
        <el-form-item label="手机号码:">
          <el-input v-model.number="formData.number" clearable placeholder="请输入" style="width: 220px"/>
        </el-form-item>
        <el-form-item label="与联系人的关系:">
          <el-input v-model="formData.ec_relationship" clearable placeholder="请输入" style="width: 220px"/>
        </el-form-item>
        <el-form-item label="学历:">
          <el-select v-model="formData.degree" placeholder="请选择" clearable style="width: 220px">
            <el-option v-for="(item,key) in degreeOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item label="毕业院校:">
          <el-input v-model="formData.graduatingschool" clearable placeholder="请输入" style="width: 220px"/>
        </el-form-item>
        <el-form-item label="毕业时间:">
          <el-date-picker v-model="formData.graduationtime" type="date" placeholder="选择日期" clearable style="width: 220px"></el-date-picker>
        </el-form-item>
        <el-form-item label="专业:">
          <el-input v-model="formData.specialized" clearable placeholder="请输入" style="width: 220px"/>
        </el-form-item>
        <el-form-item label="特长:">
          <el-input v-model="formData.specialty" clearable placeholder="请输入" style="width: 220px"/>
        </el-form-item>
        <el-form-item label="个人技能:">
          <el-input v-model="formData.personalskills" clearable placeholder="请输入" style="width: 220px"/>
        </el-form-item>
        <el-form-item label="专业职称:">
          <el-input v-model="formData.professionaltitle" clearable placeholder="请输入" style="width: 220px"/>
        </el-form-item>
        <el-form-item label="邮编:">
          <el-input v-model.number="formData.zip" clearable placeholder="请输入" style="width: 220px"/>
        </el-form-item>
        <el-form-item label="部门经理:">
          <el-select v-model="formData.departmentmanager" placeholder="请选择" clearable style="width: 220px">
            <el-option v-for="(item,key) in departmentmanagerOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item label="社保状态:">
          <el-select v-model="formData.socialsecurity" placeholder="请选择" clearable style="width: 220px">
            <el-option v-for="(item,key) in socialsecurityOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
       </el-form>
        <el-form-item>
          <el-button size="mini" type="primary" @click="save">保存</el-button>
          <el-button size="mini" type="primary" @click="back">返回</el-button>
        </el-form-item>
      </el-form>
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
  updateEmployeeStructui,
  findEmployeeStructui
} from '@/api/employeeStructui'

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref } from 'vue'
const route = useRoute()
const router = useRouter()
const type = ref('')
const departmentmanagerOptions = ref([])
const genderOptions = ref([])
const jobOptions = ref([])
const birth_monthOptions = ref([])
const degreeOptions = ref([])
const departmentOptions = ref([])
const marriageOptions = ref([])
const politicalfaceOptions = ref([])
const socialsecurityOptions = ref([])
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

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findEmployeeStructui({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.reemployeeStructui
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
    departmentmanagerOptions.value = await getDictFunc('departmentmanager')
    genderOptions.value = await getDictFunc('gender')
    jobOptions.value = await getDictFunc('job')
    birth_monthOptions.value = await getDictFunc('birth_month')
    degreeOptions.value = await getDictFunc('degree')
    departmentOptions.value = await getDictFunc('department')
    marriageOptions.value = await getDictFunc('marriage')
    politicalfaceOptions.value = await getDictFunc('politicalface')
    socialsecurityOptions.value = await getDictFunc('socialsecurity')
}

init()
// 保存按钮
const save = async() => {
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
      }
      if (res.code === 0) {
        ElMessage({
          type: 'success',
          message: '创建/更改成功'
        })
      }
}

// 返回按钮
const back = () => {
    router.go(-1)
}

</script>

<style>
</style>
