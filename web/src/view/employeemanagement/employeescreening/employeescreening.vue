<!--员工信息录入页面-->
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
    <!--    第一个标题框筛选员工信息条件-->
    <div>
      <div class="tableTitle2"><span class="midText2" style="font-family: QingNiaoHuaGuangXingKai; background-color: #377ebb">员工信息录入</span></div>
    </div>
    <!--    员工信息录入页面-->
    <el-form ref="form">
    <div class="example-block">
      <div class="example-block0"  style="margin-top: 25px;border:5px solid rgb(55,126,187);">
        <span class="example-demonstration; el-cascader-panel; input:focus"></span>
        <!--1-1小标题-->
        <img
            class="css-img9"
            src="@/image/12.png"
            alt="元素加载失败"
            style="position: relative; left: 0px; top: 0%"
        >
<!--        1-1输入框-->
        <img
            class="css-img10"
            src="@/image/13.png"
            alt="元素加载失败"
            style="position: relative; left: -30px; top: 0%"
        >
        <el-form-item label="工号:" class="example-block1-font" style="left: 75px;top: 179px">
          <el-input v-model.number="formData.worknumber" clearable class="input-el-input"/>
        </el-form-item>
        <!--1-2小标题-->
        <img
            class="css-img9"
            src="@/image/12.png"
            alt="元素加载失败"
            style="position: relative; left: -30px; top: 0%"
        >
        <!--        1-2输入框-->
        <img
            class="css-img10"
            src="@/image/13.png"
            alt="元素加载失败"
            style="position: relative; left: -60px; top: 0%"
        >
        <el-form-item label="姓名:" class="example-block1-font" style="left: 351px;top: 179px">
          <el-input v-model="formData.name" clearable class="input-el-input"/>
        </el-form-item>
        <!--1-3小标题-->
        <img
            class="css-img9"
            src="@/image/12.png"
            alt="元素加载失败"
            style="position: relative; left: -60px; top: 0%"
        >
        <!--        1-3输入框-->
        <img
            class="css-img10"
            src="@/image/13.png"
            alt="元素加载失败"
            style="position: relative; left: -90px; top: 0%"
        >
        <el-form-item label="年龄:" class="example-block1-font" style="left: 627px;top: 179px">
          <el-input v-model.number="formData.age" clearable class="input-el-input"/>
        </el-form-item>
        <!--1-4小标题-->
        <img
            class="css-img9"
            src="@/image/12.png"
            alt="元素加载失败"
            style="position: relative; left: -90px; top: 0%"
        >
        <!--        1-4输入框-->
        <img
            class="css-img10"
            src="@/image/13.png"
            alt="元素加载失败"
            style="position: relative; left: -120px; top: 0%"
        >
        <el-form-item label="性别:" class="example-block1-font" style="left: 903px;top: 179px">
          <el-select v-model="formData.sex" class="input-el-select" clearable >
            <el-option v-for="(item,key) in genderOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <!--1-5小标题-->
        <img
            class="css-img9"
            src="@/image/12.png"
            alt="元素加载失败"
            style="position: relative; left: -120px; top: 0%"
        >
        <!--        1-5输入框-->
        <img
            class="css-img10"
            src="@/image/13.png"
            alt="元素加载失败"
            style="position: relative; left: -150px; top: 0%"
        >
        <el-form-item label="部门:" class="example-block1-font" style="left: 1179px;top: 179px">
          <el-select v-model="formData.department" clearable class="input-el-select">
            <el-option v-for="(item,key) in departmentOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <!--1-6小标题-->
        <img
            class="css-img9"
            src="@/image/12.png"
            alt="元素加载失败"
            style="position: relative; left: -150px; top: 0%"
        >
        <!--        1-6长输入框-->
        <img
            class="css-img11"
            src="@/image/13.png"
            alt="元素加载失败"
            style="position: relative; left: 1480px; top: -23.5%"
        >
        <el-form-item label="入职日期" class="example-block1-font0" style="left: 1441px;top: 179px">
          <el-date-picker v-model="formData.dateofentry" type="date" clearable class="input-el-input0" style="width: 255px"></el-date-picker>
        </el-form-item>
        <!--2-1小标题-->
        <img
            class="css-img9"
            src="@/image/12.png"
            alt="元素加载失败"
            style="position: relative; left: -350px; top: -8%"
        >
        <!--        2-1输入框-->
        <img
            class="css-img10"
            src="@/image/13.png"
            alt="元素加载失败"
            style="position: relative; left: -380px; top: -8%"
        >
        <el-form-item label="岗位:" class="example-block1-font" style="left: 75px;top: 246px">
          <el-select v-model="formData.job" clearable class="input-el-select">
            <el-option v-for="(item,key) in jobOptions" :key="key" :label="item.label" :value="item.value"/>
          </el-select>
        </el-form-item>
        <!--2-2小标题-->
        <img
            class="css-img9"
            src="@/image/12.png"
            alt="元素加载失败"
            style="position: relative; left: -380px; top: -8%"
        >
        <!--        2-2输入框-->
        <img
            class="css-img10"
            src="@/image/13.png"
            alt="元素加载失败"
            style="position: relative; left: -410px; top: -8%"
        >
        <el-form-item label="生日月份" class="example-block1-font0" style="left: 331px;top: 246px">
          <el-select v-model="formData.birth_month" clearable class="input-el-select">
            <el-option v-for="(item,key) in birth_monthOptions" :key="key" :label="item.label" :value="item.value"  />
          </el-select>
        </el-form-item>
        <!--2-3小标题-->
        <img
            class="css-img9"
            src="@/image/12.png"
            alt="元素加载失败"
            style="position: relative; left: -410px; top: -8%"
        >
        <!--        2-3输入框-->
        <img
            class="css-img10"
            src="@/image/13.png"
            alt="元素加载失败"
            style="position: relative; left: -440px; top: -8%"
        >
        <el-form-item label="籍贯:"  class="example-block1-font" style="left: 627px;top: 246px">
          <el-input v-model="formData.origin" clearable class="input-el-input"/>
        </el-form-item>
        <!--2-4小标题-->
        <img
            class="css-img9"
            src="@/image/12.png"
            alt="元素加载失败"
            style="position: relative; left: -440px; top: -8%"
        >
        <!--        2-4输入框-->
        <img
            class="css-img10"
            src="@/image/13.png"
            alt="元素加载失败"
            style="position: relative; left: -470px; top: -8%"
        >
        <el-form-item label="民族:" class="example-block1-font" style="left: 903px;top: 246px">
          <el-input v-model="formData.ethnicgroup" clearable class="input-el-input"/>
        </el-form-item>
        <!--2-5小标题-->
        <img
            class="css-img9"
            src="@/image/12.png"
            alt="元素加载失败"
            style="position: relative; left: -470px; top: -8%"
        >
        <!--        2-5输入框-->
        <img
            class="css-img10"
            src="@/image/13.png"
            alt="元素加载失败"
            style="position: relative; left: 1228px; top: -31%"
        >
        <el-form-item label="婚姻状况"  class="example-block1-font0" style="left: 1164px;top: 245px">
          <el-select v-model="formData.marriage" clearable class="input-el-select">
            <el-option v-for="(item,key) in marriageOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <!--2-6小标题-->
        <img
            class="css-img9"
            src="@/image/12.png"
            alt="元素加载失败"
            style="position: relative; left: 1228px; top: -31%"
        >
        <!--        2-6长输入框-->
        <img
            class="css-img11"
            src="@/image/13.png"
            alt="元素加载失败"
            style="position: relative; left: 1175px; top: -31%"
        >
        <el-form-item label="身份证号" class="example-block1-font0" style="left: 1440px;top: 246px">
          <el-input v-model="formData.socialsecuritynumber" clearable class="input-el-input0" style="width: 255px"/>
        </el-form-item>
        <!--3-1小标题-->
        <img
            class="css-img9"
            src="@/image/12.png"
            alt="元素加载失败"
            style="position: relative; left: -656px; top: -16%"
        >
        <!--        3-1输入框-->
        <img
            class="css-img10"
            src="@/image/13.png"
            alt="元素加载失败"
            style="position: relative; left: -686px; top: -16%"
        >
        <el-form-item label="政治面貌" class="example-block1-font0" style="left: 58px;top: 311px">
          <el-select v-model="formData.politicalface" clearable class="input-el-select">
            <el-option v-for="(item,key) in politicalfaceOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <!--3-2小标题-->
        <img
            class="css-img9"
            src="@/image/12.png"
            alt="元素加载失败"
            style="position: relative; left: -686px; top: -16%"
        >
        <!--        3-2输入框-->
        <img
            class="css-img10"
            src="@/image/13.png"
            alt="元素加载失败"
            style="position: relative; left: -716px; top: -16%"
        >
        <el-form-item label="学历:" class="example-block1-font" style="left: 351px;top: 311px">
          <el-select v-model="formData.degree" clearable class="input-el-select">
            <el-option v-for="(item,key) in degreeOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <!--3-3小标题-->
        <img
            class="css-img9"
            src="@/image/12.png"
            alt="元素加载失败"
            style="position: relative; left: -716px; top: -16%"
        >
        <!--        3-3输入框-->
        <img
            class="css-img10"
            src="@/image/13.png"
            alt="元素加载失败"
            style="position: relative; left: -746px; top: -16%"
        >
        <el-form-item label="毕业院校" class="example-block1-font0" style="left: 611px;top: 311px">
          <el-input v-model="formData.graduatingschool" clearable class="input-el-input2"/>
        </el-form-item>
        <!--3-4小标题-->
        <img
            class="css-img9"
            src="@/image/12.png"
            alt="元素加载失败"
            style="position: relative; left: -746px; top: -16%"
        >
        <!--        3-4输入框-->
        <img
            class="css-img10"
            src="@/image/13.png"
            alt="元素加载失败"
            style="position: relative; left: 950px; top: -38.8%"
        >
        <el-form-item label="专业:" class="example-block1-font" style="left: 903px;top: 311px">
          <el-input v-model="formData.specialized" clearable class="input-el-input"/>
        </el-form-item>
        <!--3-5小标题-->
        <img
            class="css-img9"
            src="@/image/12.png"
            alt="元素加载失败"
            style="position: relative; left: 950px; top: -38.8%"
        >
        <!--        3-5输入框-->
        <img
            class="css-img10"
            src="@/image/13.png"
            alt="元素加载失败"
            style="position: relative; left: 923px; top: -38.8%"
        >
        <el-form-item label="特长:" class="example-block1-font" style="left: 1179px;top: 311px">
          <el-input v-model="formData.specialty" clearable class="input-el-input"/>
        </el-form-item>
        <!--3-6小标题-->
        <img
            class="css-img9"
            src="@/image/12.png"
            alt="元素加载失败"
            style="position: relative; left: 922px; top: -38.8%"
        >
        <!--        3-6长输入框-->
        <img
            class="css-img11"
            src="@/image/13.png"
            alt="元素加载失败"
            style="position: relative; left: 870px; top: -38.8%"
        >
        <el-form-item label="出生日期" class="example-block1-font0" style="left: 1441px;top: 311px">
          <el-date-picker v-model="formData.dateofbirth" type="date" clearable class="input-el-input0" style="width: 255px"></el-date-picker>
        </el-form-item>
        <!--4-1小标题-->
        <img
            class="css-img9"
            src="@/image/12.png"
            alt="元素加载失败"
            style="position: relative; left: -962px; top: -24%"
        >
        <!--        4-1输入框-->
        <img
            class="css-img10"
            src="@/image/13.png"
            alt="元素加载失败"
            style="position: relative; left: -992px; top: -24%"
        >
        <el-form-item label="专业职称" class="example-block1-font0" style="left: 58px;top: 376px">
          <el-input v-model="formData.professionaltitle" clearable class="input-el-input2"/>
        </el-form-item>
        <!--4-2小标题-->
        <img
            class="css-img9"
            src="@/image/12.png"
            alt="元素加载失败"
            style="position: relative; left: -992px; top: -24%"
        >
        <!--        4-2输入框-->
        <img
            class="css-img10"
            src="@/image/13.png"
            alt="元素加载失败"
            style="position: relative; left: -1022px; top: -24%"
        >
        <el-form-item label="邮编:" class="example-block1-font" style="left: 351px;top: 376px">
          <el-input v-model.number="formData.zip" clearable class="input-el-input"/>
        </el-form-item>
        <!--4-3小标题-->
        <img
            class="css-img9"
            src="@/image/12.png"
            alt="元素加载失败"
            style="position: relative; left: -1022px; top: -24%"
        >
        <!--        4-3输入框-->
        <img
            class="css-img10"
            src="@/image/13.png"
            alt="元素加载失败"
            style="position: relative; left: 675px; top: -47%"
        >
        <el-form-item label="部门经理" class="example-block1-font0" style="left: 611px;top: 376px">
          <el-select v-model="formData.departmentmanager" clearable class="input-el-select">
            <el-option v-for="(item,key) in departmentmanagerOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <!--4-4小标题-->
        <img
            class="css-img9"
            src="@/image/12.png"
            alt="元素加载失败"
            style="position: relative; left: 675px; top: -47%"
        >
        <!--        4-4输入框-->
        <img
            class="css-img10"
            src="@/image/13.png"
            alt="元素加载失败"
            style="position: relative; left: 644px; top: -47%"
        >
        <el-form-item label="社保状态" class="example-block1-font0" style="left: 886px;top: 376px">
          <el-select v-model="formData.socialsecurity" clearable class="input-el-select">
            <el-option v-for="(item,key) in socialsecurityOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <!--4-5小标题-->
        <img
            class="css-img9"
            src="@/image/12.png"
            alt="元素加载失败"
            style="position: relative; left: 645px; top: -47%"
        >
        <!--        4-5输入框-->
        <img
            class="css-img10"
            src="@/image/13.png"
            alt="元素加载失败"
            style="position: relative; left: 617px; top: -47%"
        >
        <el-form-item label="个人技能" class="example-block1-font0" style="left: 1164px;top: 376px">
          <el-input v-model="formData.personalskills" clearable class="input-el-input2"/>
        </el-form-item>
        <!--4-6小标题-->
        <img
            class="css-img9"
            src="@/image/12.png"
            alt="元素加载失败"
            style="position: relative; left: 615px; top: -47%"
        >
        <!--        4-6长输入框-->
        <img
            class="css-img11"
            src="@/image/13.png"
            alt="元素加载失败"
            style="position: relative; left: 565px; top: -47%"
        >
        <el-form-item label="毕业时间" class="example-block1-font0" style="left: 1440px;top: 376px">
          <el-date-picker v-model="formData.graduationtime" clearable type="date" class="input-el-input0" style="width: 255px"></el-date-picker>
        </el-form-item>
        <!--5-1小标题-->
        <img
            class="css-img12"
            src="@/image/12.png"
            alt="元素加载失败"
            style="position: relative; left: -1302px; top: -30%"
        >
        <!--        5-1输入框-->
        <img
            class="css-img11"
            src="@/image/13.png"
            alt="元素加载失败"
            style="position: relative; left: 144px; top: -53%"
        >
        <el-form-item label="户口所在地：" class="example-block1-font0" style="left: 50px;top: 451px">
          <el-input v-model="formData.householdregistration" clearable class="input-el-input1" style="width: 255px"/>
        </el-form-item>
        <!--5-2小标题-->
        <img
            class="css-img12"
            src="@/image/12.png"
            alt="元素加载失败"
            style="position: relative; left: 80px; top: -52.5%"
        >
        <!--        5-2输入框-->
        <img
            class="css-img11"
            src="@/image/13.png"
            alt="元素加载失败"
            style="position: relative; left: 8px; top: -52.5%"
        >
        <el-form-item label="身份证详细地址" class="example-block1-font0" style="left: 499px;top: 451px">
          <el-input v-model="formData.idaddress" clearable class="input-el-input1" style="width: 255px"/>
        </el-form-item>
        <!--5-3小标题-->
        <img
            class="css-img12"
            src="@/image/12.png"
            alt="元素加载失败"
            style="position: relative; left: -60px; top: -52.5%"
        >
        <!--        5-3输入框-->
        <img
            class="css-img11"
            src="@/image/13.png"
            alt="元素加载失败"
            style="position: relative; left: -130px; top: -52.5%"
        >
        <el-form-item label="现住详细地址:" class="example-block1-font0" style="left: 969px;top: 451px">
          <el-input v-model="formData.address" clearable class="input-el-input1" style="width: 255px"/>
        </el-form-item>
        <!--5-4小标题-->
        <img
            class="css-img12"
            src="@/image/12.png"
            alt="元素加载失败"
            style="position: relative; left: -190px; top: -52.5%"
        >
        <!--        5-4输入框-->
        <img
            class="css-img11"
            src="@/image/13.png"
            alt="元素加载失败"
            style="position: relative; left: 1535px; top: -75.5%"
        >
        <el-form-item label="紧急联系人:" class="example-block1-font0" style="left: 1458px;top: 451px">
          <el-input v-model="formData.emergencycontacts" clearable class="input-el-input0" style="width: 255px"/>
        </el-form-item>
        <br>
        <!--6-1小标题-->
        <img
            class="css-img12"
            src="@/image/12.png"
            alt="元素加载失败"
            style="position: relative; left: -35px; top: -84%"
        >
        <!--        6-1输入框-->
        <img
            class="css-img11"
            src="@/image/13.png"
            alt="元素加载失败"
            style="position: relative; left: -105px; top: -84%"
        >
        <el-form-item label="紧急联系人电话" class="example-block1-font0" style="left: 34px;top: 517px">
          <el-input v-model.number="formData.ec_number" clearable class="input-el-input1" style="width: 255px"/>
        </el-form-item>
        <!--6-2小标题-->
        <img
            class="css-img12"
            src="@/image/12.png"
            alt="元素加载失败"
            style="position: relative; left: -170px; top: -84%"
        >
        <!--        6-2输入框-->
        <img
            class="css-img11"
            src="@/image/13.png"
            alt="元素加载失败"
            style="position: relative; left: -243px; top: -84%"
        >
        <el-form-item label="与联系人的关系" class="example-block1-font0" style="left: 499px;top: 517px">
          <el-input v-model="formData.ec_relationship" clearable class="input-el-input1" style="width: 255px"/>
        </el-form-item>
        <!--6-3小标题-->
        <img
            class="css-img12"
            src="@/image/12.png"
            alt="元素加载失败"
            style="position: relative; left: -310px; top: -84%"
        >
        <!--        6-3输入框-->
        <img
            class="css-img11"
            src="@/image/13.png"
            alt="元素加载失败"
            style="position: relative; left: -381px; top: -84%"
        >
        <el-form-item label="手机号码：" class="example-block1-font0" style="left: 1000px;top: 517px">
          <el-input v-model.number="formData.number" clearable class="input-el-input1" style="width: 255px"/>
        </el-form-item>
        <!--6-4小标题-->
<!--        <img-->
<!--            class="css-img12"-->
<!--            src="@/image/12.png"-->
<!--            alt="元素加载失败"-->
<!--            style="position: relative; left: 1350px; top: -137%"-->
<!--        >-->
        <!--        6-4输入框-->
<!--        <img-->
<!--            class="css-img11"-->
<!--            src="@/image/13.png"-->
<!--            alt="元素加载失败"-->
<!--            style="position: relative; left: 1287px; top: -137%"-->
<!--        >-->
      </div>
    </div>
    </el-form>
    <div class="example-block1">
      <!--        保存按钮样式-->
      <img
          class="css-img15"
          src="@/image/15.png"
          alt="元素加载失败"
          style="position: relative; left: 1650px; top: -10%"
      >
      <el-button size="large" type="primary" @click="save" class="button0" plain style="left: 1433px; top: 600px; --el-button-text-color: #ffffff; --el-button-hover-bg-color: #0a3d98">保存</el-button>
<!--      取消按钮样式-->
      <img
          class="css-img15"
          src="@/image/15.png"
          alt="元素加载失败"
          style="position: relative; left: 1160px; top: -10%"
      >
      <el-button size="large" type="primary" @click="back" class="button0" plain style="left: 1683px; top: 600px; --el-button-text-color: #ffffff; --el-button-hover-bg-color: #0a3d98">返回</el-button>
    </div>


    <!--    员工信息显示页面-->
    <div class="example-block2">
      <div class="example-block0" style="margin-top: 25px;border:5px solid rgb(55,126,187);">
        <el-table
            ref="multipleTable"
            :data="tableData"
            height="200"
            tooltip-effect="dark"
            row-key="ID"
            style="position: absolute;width: 96.45%; height: 51.6%;background-color: #f9c449"
            :cell-style="{borderColor:'#C0C0C0'}"
            :header-cell-style="{borderColor:'#C0C0C0'}"
            border
        >
<!--          <el-table-column type="selection" width="40" />-->
          <el-table-column align="center" label="操作" width="200">
            <template #default="scope">
              <el-button type="text" icon="edit" size="small" class="table-button" @click="goPkgs">详情-变更</el-button>
              <el-button type="text" icon="delete" size="small" @click="deleteRow(scope.row)">删除</el-button>
            </template>
          </el-table-column>
          <el-table-column align="center" label="日期" width="180">
            <template #default="scope">{{ formatDate(scope.row.CreatedAt) }}</template>
          </el-table-column>
          <el-table-column align="center" label="工号" prop="worknumber" width="60" />
          <el-table-column align="center" label="姓名" prop="name" width="70" />
          <el-table-column align="center" label="年龄" prop="age" width="60" />
          <el-table-column align="center" label="性别" prop="sex" width="55">
            <template #default="scope">
              {{ filterDict(scope.row.sex,genderOptions) }}
            </template>
          </el-table-column>
          <el-table-column align="center" label="部门" prop="department" width="330">
            <template #default="scope">
              {{ filterDict(scope.row.department,departmentOptions) }}
            </template>
          </el-table-column>
          <el-table-column align="center" label="入职日期" prop="dateofentry" width="120" />
          <el-table-column align="center" label="岗位" prop="job" width="150">
            <template #default="scope">
              {{ filterDict(scope.row.job,jobOptions) }}
            </template>
          </el-table-column>
          <el-table-column align="center" label="生日月份" prop="birth_month" width="120">
              <template #default="scope">
              {{ filterDict(scope.row.birth_month,birth_monthOptions) }}
              </template>
          </el-table-column>
          <el-table-column align="center" label="籍贯" prop="origin" width="120" />
          <el-table-column align="center" label="民族" prop="ethnicgroup" width="120" />
          <el-table-column align="center" label="婚姻状况" prop="marriage" width="120">
              <template #default="scope">
              {{ filterDict(scope.row.marriage,marriageOptions) }}
              </template>
          </el-table-column>
          <el-table-column align="center" label="身份证号" prop="socialsecuritynumber" width="120" />
          <el-table-column align="center" label="政治面貌" prop="politicalface" width="120">
              <template #default="scope">
              {{ filterDict(scope.row.politicalface,politicalfaceOptions) }}
              </template>
          </el-table-column>
          <el-table-column align="center" label="学历" prop="degree" width="120">
              <template #default="scope">
              {{ filterDict(scope.row.degree,degreeOptions) }}
              </template>
          </el-table-column>
          <el-table-column align="center" label="毕业院校" prop="graduatingschool" width="120" />
          <el-table-column align="center" label="专业" prop="specialized" width="120" />
          <el-table-column align="center" label="特长" prop="specialty" width="120" />
          <el-table-column align="center" label="出生日期" prop="dateofbirth" width="120" />
          <el-table-column align="center" label="专业职称" prop="professionaltitle" width="120" />
          <el-table-column align="center" label="邮编" prop="zip" width="120" />
          <el-table-column align="center" label="部门经理" prop="departmentmanager" width="120">
            <template #default="scope">
              {{ filterDict(scope.row.departmentmanager,departmentmanagerOptions) }}
            </template>
          </el-table-column>
          <el-table-column align="center" label="社保状态" prop="socialsecurity" width="120">
            <template #default="scope">
              {{ filterDict(scope.row.socialsecurity,socialsecurityOptions) }}
            </template>
          </el-table-column>
          <el-table-column align="center" label="个人技能" prop="personalskills" width="120" />
          <el-table-column align="center" label="毕业时间" prop="graduationtime" width="120" />
          <el-table-column align="center" label="户口所在地" prop="householdregistration" width="120" />
          <el-table-column align="center" label="身份证详细地址" prop="idaddress" width="120" />
          <el-table-column align="center" label="现住详细地址" prop="address" width="120" />
          <el-table-column align="center" label="紧急联系人" prop="emergencycontacts" width="120" />
          <el-table-column align="center" label="紧急联系人电话" prop="ec_number" width="120" />
          <el-table-column align="center" label="与联系人的关系" prop="ec_relationship" width="120" />
          <el-table-column align="center" label="手机号码" prop="number" width="120" />
        </el-table>
      </div>
    </div>
<!--    分页选择器-->
    <div class="gva-pagination">
      <el-pagination
          layout="total, sizes, prev, pager, next, jumper"
          :current-page="page"
          :page-size="pageSize"
          :page-sizes="[5, 30, 50, 100]"
          :total="total"
          @current-change="handleCurrentChange"
          @size-change="handleSizeChange"
      />
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

<!--图像加载-->
<script setup>
import {
  createEmployeeStructui,
  deleteEmployeeStructui,
  deleteEmployeeStructuiByIds,
  updateEmployeeStructui,
  findEmployeeStructui,
  getEmployeeStructuiList
} from '@/api/employeeStructui'
// 自动获取字典
import { getDictFunc, formatDate, formatBoolean, filterDict } from '@/utils/format'
import { Search } from '@element-plus/icons-vue'
import { useRoute, useRouter } from "vue-router"
import { ElMessage, ElMessageBox } from 'element-plus'
import { useUserStore } from '@/pinia/modules/user'

// import { ref } from 'vue'
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
  worknumber: '',
  name: '',
  age: '',
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
  ec_number: '',
  number: '',
  ec_relationship: '',
  degree: undefined,
  graduatingschool: '',
  graduationtime: new Date(),
  specialized: '',
  specialty: '',
  personalskills: '',
  professionaltitle: '',
  zip: '',
  departmentmanager: undefined,
  socialsecurity: undefined,
})

// =========== 表格控制部分 ===========
const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])
const searchInfo = ref({})

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

getTableData()// 控制搜索页不自动加载数据

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


// 点击详情-变更跳转到员工信息变更页面
const goPkgs = () => {
  router.push({ name: 'EmployeeModificationhr' })
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

// 设计页面底层
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
  height: 465px;
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
  margin-left:-20px;
  margin-right:-20px;
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
.example-block2 {
  /*容器大小*/
  height: 550px;
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
  margin-left:-20px;
  margin-right:-20px;
}
.example-demonstration {
  margin: 1rem;
}
.example-block1-font{
  font-family: QingNiaoHuaGuangXingKai;
  position: absolute;
  z-index: 1;
  --font-size: 30px;
  --el-text-color-regular: #ffffff;
  margin: auto 0;
}
.example-block1-font0{
  font-family: QingNiaoHuaGuangXingKai;
  position: absolute;
  z-index: 1;
  --font-size: 25px;
  --el-text-color-regular: #ffffff;
}
.input-el-input{
  width: 38.6%;
  --el-input-bg-color: transparent;
  /*--el-input-bg-color: #f94949;*/
  --el-input-border: none;
  font-size: 22px;
  padding-left: 19px;
  --el-text-color-regular: #353755;
  top: -4px;
}
.input-el-input0{
  --el-input-bg-color: transparent;
  /*--el-input-bg-color: #f94949;*/
  --el-input-border: none;
  font-size: 20px;
  padding-left: 15px;
  --el-text-color-regular: #353755;
  top: -4px;
  --el-text-color-placeholder: transparent;
}
.input-el-input1{
  --el-input-bg-color: transparent;
  /*--el-input-bg-color: #f94949;*/
  --el-input-border: none;
  font-size: 20px;
  padding-left: 6px;
  --el-text-color-regular: #353755;
  top: -4px;
}
.input-el-input2{
  width: 38.6%;
  --el-input-bg-color: transparent;
  /*--el-input-bg-color: #f94949;*/
  --el-input-border: none;
  font-size: 22px;
  padding-left: 6px;
  --el-text-color-regular: #353755;
  top: -4px;
}
.input-el-input3{
  /*--el-input-bg-color: transparent;*/
  --el-input-bg-color: #f94949;
  --el-input-border: none;
  font-size: 20px;
  padding-left: 20px;
  --el-text-color-regular: #353755;
  top: -4px;
}

.input-el-select{
  width: 55%;
  border:none;
  outline:none;
  --el-color-white: transparent;
  /*--el-color-white: #f94949;*/
  --el-border-base: transparent;
  --el-text-color-placeholder: transparent;
  top: -3px;
  padding-left: 19px;
  --el-text-color-regular: #353755;
}
/*下拉栏默认字体控制*/
/*input::placeholder{*/
/*  color: transparent !important;*/
/*}*/

.input-el-select0{
  width: 275px;
  border:none;
  outline:none;
  /*--el-color-white: transparent;*/
  --el-color-white: #f94949;
  --el-border-base: transparent;
  --el-text-color-placeholder: transparent;
  top: -2px;
  padding-left: 28px;
}
/*保存、返回按钮*/
.button0{
  width: 176px;
  height: 54px;
  background: transparent;
  position: absolute;
  z-index: 1;
  font-size: 40px;
  font-family: QingNiaoHuaGuangXingKai;
}


.css-img9{
  width: 153px;
  height: 90px;
}
.css-img10{
  width: 153px;
  height: 100px;
}
.css-img11{
  width: 350px;
  height: 100px;
}
/*5-*小标题专用*/
.css-img12{
  width: 250px;
  height: 90px;
}
.css-img15{
  width: 240px;
  height: 99.2px;
}
</style>
