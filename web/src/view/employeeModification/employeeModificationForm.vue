<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" label-position="right" label-width="80px">
        <el-form-item label="工号:">
          <el-input v-model.number="formData.worknumber" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="姓名:">
          <el-input v-model="formData.name" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="年龄:">
          <el-input v-model.number="formData.age" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="性别:">
          <el-input v-model.number="formData.sex" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="部门:">
          <el-input v-model.number="formData.department" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="岗位名称:">
          <el-input v-model.number="formData.job" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="入职日期:">
          <el-date-picker v-model="formData.dateofentry" type="date" placeholder="选择日期" clearable></el-date-picker>
        </el-form-item>
        <el-form-item label="身份证号:">
          <el-input v-model="formData.socialsecuritynumber" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="出生日期:">
          <el-date-picker v-model="formData.dateofbirth" type="date" placeholder="选择日期" clearable></el-date-picker>
        </el-form-item>
        <el-form-item label="生日月份:">
          <el-input v-model.number="formData.birthMonth" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="籍贯:">
          <el-input v-model="formData.origin" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="户口所在地:">
          <el-input v-model="formData.householdregistration" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="身份证详细地址:">
          <el-input v-model="formData.idaddress" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="现住详细地址:">
          <el-input v-model="formData.address" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="民族:">
          <el-input v-model="formData.ethnicgroup" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="婚姻状况:">
          <el-input v-model.number="formData.marriage" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="政治面貌:">
          <el-input v-model.number="formData.politicalface" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="紧急联系人:">
          <el-input v-model="formData.emergencycontacts" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="紧急联系人电话:">
          <el-input v-model.number="formData.ecNumber" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="手机号码:">
          <el-input v-model.number="formData.number" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="与联系人的关系:">
          <el-input v-model="formData.ecRelationship" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="学历:">
          <el-input v-model.number="formData.degree" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="毕业院校:">
          <el-input v-model="formData.graduatingschool" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="毕业时间:">
          <el-date-picker v-model="formData.graduationtime" type="date" placeholder="选择日期" clearable></el-date-picker>
        </el-form-item>
        <el-form-item label="专业:">
          <el-input v-model="formData.specialized" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="特长:">
          <el-input v-model="formData.specialty" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="个人技能:">
          <el-input v-model="formData.personalskills" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="专业职称:">
          <el-input v-model="formData.professionaltitle" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="邮编:">
          <el-input v-model.number="formData.zip" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="部门经理:">
          <el-input v-model.number="formData.departmentmanager" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="社保状态:">
          <el-input v-model.number="formData.socialsecurity" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="员工照片:">
          <el-input v-model="formData.employeeImg" clearable placeholder="请输入" />
        </el-form-item>
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
  name: 'EmployeeModification'
}
</script>

<script setup>
import {
  createEmployeeModification,
  updateEmployeeModification,
  findEmployeeModification
} from '@/api/employeeModification'

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref } from 'vue'
const route = useRoute()
const router = useRouter()
const type = ref('')
const formData = ref({
        worknumber: 0,
        name: '',
        age: 0,
        sex: 0,
        department: 0,
        job: 0,
        dateofentry: new Date(),
        socialsecuritynumber: '',
        dateofbirth: new Date(),
        birthMonth: 0,
        origin: '',
        householdregistration: '',
        idaddress: '',
        address: '',
        ethnicgroup: '',
        marriage: 0,
        politicalface: 0,
        emergencycontacts: '',
        ecNumber: 0,
        number: 0,
        ecRelationship: '',
        degree: 0,
        graduatingschool: '',
        graduationtime: new Date(),
        specialized: '',
        specialty: '',
        personalskills: '',
        professionaltitle: '',
        zip: 0,
        departmentmanager: 0,
        socialsecurity: 0,
        employeeImg: '',
        })

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findEmployeeModification({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.reemployeeModification
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
}

init()
// 保存按钮
const save = async() => {
      let res
      switch (type.value) {
        case 'create':
          res = await createEmployeeModification(formData.value)
          break
        case 'update':
          res = await updateEmployeeModification(formData.value)
          break
        default:
          res = await createEmployeeModification(formData.value)
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
