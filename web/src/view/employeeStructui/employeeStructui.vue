<template>
  <div>
    <div class="gva-search-box">
<!--      :inline="true"数据并列-->
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
        <el-form-item label="工号:">
          <el-input v-model="searchInfo.worknumber" placeholder="搜索条件" :suffix-icon="Search"  style="width: 220px"/>
        </el-form-item>
        <el-form-item label="姓名:">
          <el-input v-model="searchInfo.name" placeholder="搜索条件" style="width: 220px"/>
        </el-form-item>
        <el-form-item label="年龄:">
          <el-input v-model="searchInfo.age" placeholder="搜索条件" style="width: 220px"/>
        </el-form-item>
<!--        <el-form-item label="性别">-->
<!--          <el-input v-model="searchInfo.sex" placeholder="搜索条件" />-->
<!--        </el-form-item>-->
        <el-form-item label="性别:">
          <el-select v-model="searchInfo.sex" placeholder="搜索条件" clearable style="width: 220px">
            <el-option v-for="(item,key) in genderOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>

<!--        <el-form-item label="部门">-->
<!--          <el-input v-model="searchInfo.department" placeholder="搜索条件" />-->
<!--        </el-form-item>-->
        <el-form-item label="部门:">
          <el-select v-model="searchInfo.department" placeholder="请选择" clearable style="width: 220px">
            <el-option v-for="(item,key) in departmentOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>

<!--        <el-form-item label="岗位">-->
<!--          <el-input v-model="searchInfo.job" placeholder="搜索条件" />-->
<!--        </el-form-item>-->
        <el-form-item label="岗位:">
          <el-select v-model="searchInfo.job" placeholder="请选择" clearable style="width: 220px">
            <el-option v-for="(item,key) in jobOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
<!--        <el-form-item label="入职日期">-->
<!--          <el-input v-model="searchInfo.dateofentry" placeholder="搜索条件" />-->
<!--        </el-form-item>-->
<!--        <el-form-item label="身份证号">-->
<!--          <el-input v-model="searchInfo.socialsecuritynumber" placeholder="搜索条件" />-->
<!--        </el-form-item>-->
<!--        <el-form-item label="出生日期">-->
<!--          <el-input v-model="searchInfo.dateofbirth" placeholder="搜索条件" />-->
<!--        </el-form-item>-->

<!--        <el-form-item label="生日月份">-->
<!--          <el-input v-model="searchInfo.birth_month" placeholder="搜索条件" />-->
<!--        </el-form-item>-->

        <el-form-item label="籍贯:">
          <el-input v-model="searchInfo.origin" placeholder="搜索条件" style="width: 220px"/>
        </el-form-item>
<!--        <el-form-item label="户口所在地">-->
<!--          <el-input v-model="searchInfo.householdregistration" placeholder="搜索条件" />-->
<!--        </el-form-item>-->
<!--        <el-form-item label="身份证详细地址">-->
<!--          <el-input v-model="searchInfo.idaddress" placeholder="搜索条件" />-->
<!--        </el-form-item>-->
        <el-form-item label="民族:">
          <el-input v-model="searchInfo.ethnicgroup" placeholder="搜索条件" style="width: 220px"/>
        </el-form-item>
<!--        <el-form-item label="婚姻状况">-->
<!--          <el-input v-model="searchInfo.marriage" placeholder="搜索条件" />-->
<!--        </el-form-item>-->
<!--        <el-form-item label="政治面貌">-->
<!--          <el-input v-model="searchInfo.politicalface" placeholder="搜索条件" />-->
<!--        </el-form-item>-->
        <el-form-item label="学历:">
          <el-select v-model="searchInfo.degree" placeholder="请选择" clearable style="width: 220px">
            <el-option v-for="(item,key) in degreeOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item label="专业:">
          <el-input v-model="searchInfo.specialized" placeholder="搜索条件" style="width: 220px"/>
        </el-form-item>
        <el-form-item label="生日月份:">
          <el-select v-model="searchInfo.birth_month" placeholder="请选择" clearable style="width: 192px">
            <el-option v-for="(item,key) in birth_monthOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item label="现住地址:">
          <el-input v-model="searchInfo.address" placeholder="搜索条件" style="width: 192px"/>
        </el-form-item>
        <el-form-item label="婚姻状况:">
          <el-select v-model="searchInfo.marriage" placeholder="请选择" clearable style="width: 192px">
            <el-option v-for="(item,key) in marriageOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item label="政治面貌:">
          <el-select v-model="searchInfo.politicalface" placeholder="请选择" clearable style="width: 192px">
            <el-option v-for="(item,key) in politicalfaceOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>

<!--        <el-form-item label="紧急联系人">-->
<!--          <el-input v-model="searchInfo.emergencycontacts" placeholder="搜索条件" />-->
<!--        </el-form-item>-->
<!--        <el-form-item label="紧急联系人电话">-->
<!--          <el-input v-model="searchInfo.ec_number" placeholder="搜索条件" />-->
<!--        </el-form-item>-->
        <el-form-item label="手机号码:">
          <el-input v-model="searchInfo.number" placeholder="搜索条件" style="width: 192px"/>
        </el-form-item>
<!--        <el-form-item label="与联系人的关系">-->
<!--          <el-input v-model="searchInfo.ec_relationship" placeholder="搜索条件" />-->
<!--        </el-form-item>-->
<!--        <el-form-item label="学历">-->
<!--          <el-input v-model="searchInfo.degree" placeholder="搜索条件" />-->
<!--        </el-form-item>-->

        <el-form-item label="毕业院校:">
          <el-input v-model="searchInfo.graduatingschool" placeholder="搜索条件" style="width: 192px"/>
        </el-form-item>
<!--        <el-form-item label="毕业时间:">-->
<!--          <el-input v-model="searchInfo.graduationtime" placeholder="搜索条件" style="width: 192px"/>-->
<!--        </el-form-item>-->
<!--        <el-form-item label="特长">-->
<!--          <el-input v-model="searchInfo.specialty" placeholder="搜索条件" />-->
<!--        </el-form-item>-->
        <el-form-item label="个人技能:">
          <el-input v-model="searchInfo.personalskills" placeholder="搜索条件" style="width: 192px"/>
        </el-form-item>
        <el-form-item label="专业职称:">
          <el-input v-model="searchInfo.professionaltitle" placeholder="搜索条件" style="width: 192px"/>
        </el-form-item>
<!--        <el-form-item label="邮编">-->
<!--          <el-input v-model="searchInfo.zip" placeholder="搜索条件" />-->
<!--        </el-form-item>-->
<!--        <el-form-item label="部门经理">-->
<!--          <el-input v-model="searchInfo.departmentmanager" placeholder="搜索条件" />-->
<!--        </el-form-item>-->
<!--        <el-form-item label="社保状态">-->
<!--          <el-input v-model="searchInfo.socialsecurity" placeholder="搜索条件" />-->
<!--        </el-form-item>-->
        <el-form-item label="部门经理:">
          <el-select v-model="searchInfo.departmentmanager" placeholder="请选择" clearable style="width: 192px">
            <el-option v-for="(item,key) in departmentmanagerOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item label="社保状态:">
          <el-select v-model="searchInfo.socialsecurity" placeholder="请选择" clearable style="width: 192px">
            <el-option v-for="(item,key) in socialsecurityOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>

        <el-form-item>
          <div style="text-align: right">

            <el-button size="small" type="primary" icon="search" @click="onSubmit" >查询</el-button>

            <el-button size="small" icon="refresh" @click="onReset">重置</el-button>

          </div>
        </el-form-item>
      </el-form>
    </div>
    <div class="gva-table-box">
        <div class="gva-btn-list">
            <el-button size="small" type="primary" icon="plus" @click="openDialog">新增</el-button>
            <el-popover v-model:visible="deleteVisible" placement="top" width="160">
            <p>确定要删除吗？</p>
            <div style="text-align: right; margin-top: 8px;">
                <el-button size="small" type="text" @click="deleteVisible = false">取消</el-button>
                <el-button size="small" type="primary" @click="onDelete">确定</el-button>
            </div>
            <template #reference>
                <el-button icon="delete" size="small" style="margin-left: 10px;" :disabled="!multipleSelection.length" @click="deleteVisible = true">删除</el-button>
            </template>
            </el-popover>
        </div>
        <el-table
        ref="multipleTable"
        style="width: 100%"
        tooltip-effect="dark"
        :data="tableData"
        row-key="ID"
        @selection-change="handleSelectionChange"
        >
          //数据表格
        <el-table-column type="selection" width="55" />
<!--          <el-table-column align="left" label="日期" width="180">-->
<!--            <template #default="scope">{{ formatDate(scope.row.CreatedAt) }}</template>-->
<!--          </el-table-column>-->
          <el-table-column align="left" label="工号" prop="worknumber" width="120" />
          <el-table-column align="left" label="姓名" prop="name" width="120" />
          <el-table-column align="left" label="年龄" prop="age" width="120" />
          <el-table-column align="left" label="性别" prop="sex" width="120">
            <template #default="scope">
              {{ filterDict(scope.row.sex,genderOptions) }}
            </template>
          </el-table-column>
          <el-table-column align="left" label="部门" prop="department" width="120">
            <template #default="scope">
            {{ filterDict(scope.row.department,departmentOptions) }}
            </template>
        </el-table-column>
        <el-table-column align="left" label="岗位" prop="job" width="120">
            <template #default="scope">
            {{ filterDict(scope.row.job,jobOptions) }}
            </template>
        </el-table-column>
        <el-table-column align="left" label="入职日期" prop="dateofentry" width="120" />
        <el-table-column align="left" label="身份证号" prop="socialsecuritynumber" width="120" />
        <el-table-column align="left" label="出生日期" prop="dateofbirth" width="120" />
        <el-table-column align="left" label="生日月份" prop="birth_month" width="120">
            <template #default="scope">
            {{ filterDict(scope.row.birth_month,birth_monthOptions) }}
            </template>
        </el-table-column>
        <el-table-column align="left" label="籍贯" prop="origin" width="120" />
        <el-table-column align="left" label="户口所在地" prop="householdregistration" width="120" />
        <el-table-column align="left" label="身份证详细地址" prop="idaddress" width="120" />
        <el-table-column align="left" label="现住详细地址" prop="address" width="120" />
        <el-table-column align="left" label="民族" prop="ethnicgroup" width="120" />
        <el-table-column align="left" label="婚姻状况" prop="marriage" width="120">
            <template #default="scope">
            {{ filterDict(scope.row.marriage,marriageOptions) }}
            </template>
        </el-table-column>
        <el-table-column align="left" label="政治面貌" prop="politicalface" width="120">
            <template #default="scope">
            {{ filterDict(scope.row.politicalface,politicalfaceOptions) }}
            </template>
        </el-table-column>
        <el-table-column align="left" label="紧急联系人" prop="emergencycontacts" width="120" />
        <el-table-column align="left" label="紧急联系人电话" prop="ec_number" width="120" />
        <el-table-column align="left" label="手机号码" prop="number" width="120" />
        <el-table-column align="left" label="与联系人的关系" prop="ec_relationship" width="120" />
        <el-table-column align="left" label="学历" prop="degree" width="120">
            <template #default="scope">
            {{ filterDict(scope.row.degree,degreeOptions) }}
            </template>
        </el-table-column>
        <el-table-column align="left" label="毕业院校" prop="graduatingschool" width="120" />
        <el-table-column align="left" label="毕业时间" prop="graduationtime" width="120" />
        <el-table-column align="left" label="专业" prop="specialized" width="120" />
        <el-table-column align="left" label="特长" prop="specialty" width="120" />
        <el-table-column align="left" label="个人技能" prop="personalskills" width="120" />
        <el-table-column align="left" label="专业职称" prop="professionaltitle" width="120" />
        <el-table-column align="left" label="邮编" prop="zip" width="120" />
        <el-table-column align="left" label="部门经理" prop="departmentmanager" width="120">
            <template #default="scope">
            {{ filterDict(scope.row.departmentmanager,departmentmanagerOptions) }}
            </template>
        </el-table-column>
        <el-table-column align="left" label="社保状态" prop="socialsecurity" width="120">
            <template #default="scope">
            {{ filterDict(scope.row.socialsecurity,socialsecurityOptions) }}
            </template>
        </el-table-column>
        <el-table-column align="left" label="按钮组" width="130px">
            <template #default="scope">
              <el-button type="text" icon="edit" size="small" class="table-button" @click="updateEmployeeStructuiFunc(scope.row)">详情-变更</el-button>
              <el-button type="text" icon="delete" size="small" @click="deleteRow(scope.row)">删除</el-button>
            </template>
        </el-table-column>

        </el-table>
        <div class="gva-pagination">
            <el-pagination
            layout="total, sizes, prev, pager, next, jumper"
            :current-page="page"
            :page-size="pageSize"
            :page-sizes="[10, 30, 50, 100]"
            :total="total"
            @current-change="handleCurrentChange"
            @size-change="handleSizeChange"
            />
        </div>
    </div>
    <!--    Dialog对话框-->
    <!--    :inline="true"数据并列-->
    <el-dialog v-model="dialogFormVisible" :before-close="closeDialog" title="详情-变更" width="80%">
<!--      设置数据距离左侧宽度label-width="120px"-->
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
          <el-select v-model="formData.department" placeholder="请选择" style="width:220px" clearable>
            <el-option v-for="(item,key) in departmentOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item label="岗位:">
          <el-select v-model="formData.job" placeholder="请选择" style="width:220px" clearable>
            <el-option v-for="(item,key) in jobOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item label="入职日期:">
          <el-date-picker v-model="formData.dateofentry" type="date" style="width:220px" placeholder="选择日期" clearable />
        </el-form-item>
        <el-form-item label="身份证号:">
          <el-input v-model="formData.socialsecuritynumber" clearable placeholder="请输入" style="width:220px"/>
        </el-form-item>
        <el-form-item label="出生日期:">
          <el-date-picker v-model="formData.dateofbirth" type="date" style="width:220px" placeholder="选择日期" clearable />
        </el-form-item>
        <el-form-item label="生日月份:">
          <el-select v-model="formData.birth_month" placeholder="请选择" style="width:220px" clearable>
            <el-option v-for="(item,key) in birth_monthOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item label="籍贯:">
          <el-input v-model="formData.origin" clearable placeholder="请输入" style="width:220px"/>
        </el-form-item>
        <el-form-item label="户口所在地:" >
          <el-input v-model="formData.householdregistration" clearable placeholder="请输入" style="width:220px"/>
        </el-form-item>
        <el-form-item label="身份证详细地址:">
          <el-input v-model="formData.idaddress" clearable placeholder="请输入" style="width:220px"/>
        </el-form-item>
        <el-form-item label="现住详细地址:">
          <el-input v-model="formData.address" clearable placeholder="请输入" style="width:220px"/>
        </el-form-item>
        <el-form-item label="民族:">
          <el-input v-model="formData.ethnicgroup" clearable placeholder="请输入" style="width:220px"/>
        </el-form-item>
        <el-form-item label="婚姻状况:">
          <el-select v-model="formData.marriage" placeholder="请选择" style="width:220px" clearable>
            <el-option v-for="(item,key) in marriageOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item label="政治面貌:">
          <el-select v-model="formData.politicalface" placeholder="请选择" style="width:220px" clearable>
            <el-option v-for="(item,key) in politicalfaceOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item label="紧急联系人:">
          <el-input v-model="formData.emergencycontacts" clearable placeholder="请输入" style="width:220px"/>
        </el-form-item>
        <el-form-item label="紧急联系人电话:">
          <el-input v-model.number="formData.ec_number" clearable placeholder="请输入" style="width:220px"/>
        </el-form-item>
        <el-form-item label="手机号码:">
          <el-input v-model.number="formData.number" clearable placeholder="请输入" style="width:220px" />
        </el-form-item>
        <el-form-item label="与联系人的关系:">
          <el-input v-model="formData.ec_relationship" clearable placeholder="请输入" style="width:220px"/>
        </el-form-item>
        <el-form-item label="学历:">
          <el-select v-model="formData.degree" placeholder="请选择" style="width:220px" clearable>
            <el-option v-for="(item,key) in degreeOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item label="毕业院校:">
          <el-input v-model="formData.graduatingschool" clearable placeholder="请输入" style="width:220px"/>
        </el-form-item>
        <el-form-item label="毕业时间:">
          <el-date-picker v-model="formData.graduationtime" type="date" style="width:220px" placeholder="选择日期" clearable />
        </el-form-item>
        <el-form-item label="专业:">
          <el-input v-model="formData.specialized" clearable placeholder="请输入" style="width:220px"/>
        </el-form-item>
        <el-form-item label="特长:">
          <el-input v-model="formData.specialty" clearable placeholder="请输入" style="width:220px"/>
        </el-form-item>
        <el-form-item label="个人技能:">
          <el-input v-model="formData.personalskills" clearable placeholder="请输入" style="width:220px"/>
        </el-form-item>
        <el-form-item label="专业职称:">
          <el-input v-model="formData.professionaltitle" clearable placeholder="请输入" style="width:220px"/>
        </el-form-item>
        <el-form-item label="邮编:">
          <el-input v-model.number="formData.zip" clearable placeholder="请输入" style="width:220px"/>
        </el-form-item>
        <el-form-item label="部门经理:">
          <el-select v-model="formData.departmentmanager" placeholder="请选择" style="width:220px" clearable>
            <el-option v-for="(item,key) in departmentmanagerOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item label="社保状态:">
          <el-select v-model="formData.socialsecurity" placeholder="请选择" style="width:220px" clearable>
            <el-option v-for="(item,key) in socialsecurityOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
      </el-form>
      <template #footer>
        <div class="dialog-footer">
          <el-button size="small" @click="closeDialog">取 消</el-button>
          <el-button size="small" type="primary" @click="enterDialog">确 定</el-button>
        </div>
      </template>
    </el-dialog>
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

// 全量引入格式化工具 请按需保留
import { getDictFunc, formatDate, formatBoolean, filterDict } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref } from 'vue'
import { Search } from '@element-plus/icons-vue'

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
      }
      if (res.code === 0) {
        ElMessage({
          type: 'success',
          message: '创建/更改成功'
        })
        closeDialog()
        getTableData()
      }
}
</script>

<style>
</style>
