<!--员工筛选页面-->
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
    <!--    第一个标题框员工信息变更-->
    <div>
      <div class="tableTitle2"><span class="midText2" style="font-family: QingNiaoHuaGuangXingKai; background-color: #377ebb">员工信息变更</span></div>
    </div>
    <!--    员工信息变更页面-->
    <el-form ref="form">
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
        <el-form-item label="工号:" class="example-block1-font" style="left: 90px;top: 205px">
          <el-input v-model.number="formData.worknumber" clearable class="input-el-input"/>
        </el-form-item>
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
        <el-form-item label="姓名:" class="example-block1-font" style="left: 383px;top: 205px">
          <el-input v-model="formData.name" clearable class="input-el-input"/>
        </el-form-item>
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
        <el-form-item label="年龄:" class="example-block1-font" style="left: 676px;top: 205px">
          <el-input v-model.number="formData.age" clearable class="input-el-input"/>
        </el-form-item>
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
        <el-form-item label="性别:" class="example-block1-font" style="left: 969px;top: 205px">
          <el-select v-model="formData.sex" class="input-el-select" clearable >
            <el-option v-for="(item,key) in genderOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
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
        <el-form-item label="部门:" class="example-block1-font" style="left: 1262px;top: 205px">
          <el-select v-model="formData.department" clearable class="input-el-select">
            <el-option v-for="(item,key) in departmentOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
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
        <el-form-item label="岗位:" class="example-block1-font" style="left: 90px;top: 267px">
          <el-select v-model="formData.job" clearable class="input-el-select">
            <el-option v-for="(item,key) in jobOptions" :key="key" :label="item.label" :value="item.value"/>
          </el-select>
        </el-form-item>
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
        <el-form-item label="生日月份" class="example-block1-font0" style="left: 368px;top: 267px">
          <el-select v-model="formData.birth_month" clearable class="input-el-select1">
            <el-option v-for="(item,key) in birth_monthOptions" :key="key" :label="item.label" :value="item.value"  />
          </el-select>
        </el-form-item>
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
        <el-form-item label="籍贯:"  class="example-block1-font" style="left: 676px;top: 267px">
          <el-input v-model="formData.origin" clearable class="input-el-input"/>
        </el-form-item>
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
        <el-form-item label="民族:" class="example-block1-font" style="left: 969px;top: 267px">
          <el-input v-model="formData.ethnicgroup" clearable class="input-el-input"/>
        </el-form-item>
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
        <el-form-item label="婚姻状况"  class="example-block1-font0" style="left: 1246px;top: 267px">
          <el-select v-model="formData.marriage" clearable class="input-el-select1">
            <el-option v-for="(item,key) in marriageOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
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
        <el-form-item label="政治面貌" class="example-block1-font0" style="left: 75px;top: 333px">
          <el-select v-model="formData.politicalface" clearable class="input-el-select1">
            <el-option v-for="(item,key) in politicalfaceOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
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
        <el-form-item label="学历:" class="example-block1-font" style="left: 383px;top: 331px">
          <el-select v-model="formData.degree" clearable class="input-el-select">
            <el-option v-for="(item,key) in degreeOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
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
        <el-form-item label="毕业院校" class="example-block1-font0" style="left: 662px;top: 331px">
          <el-input v-model="formData.graduatingschool" clearable class="input-el-input2"/>
        </el-form-item>
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
        <el-form-item label="专业:" class="example-block1-font" style="left: 969px;top: 331px">
          <el-input v-model="formData.specialized" clearable class="input-el-input"/>
        </el-form-item>
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
        <el-form-item label="特长:" class="example-block1-font" style="left: 1262px;top: 331px">
          <el-input v-model="formData.specialty" clearable class="input-el-input"/>
        </el-form-item>
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
        <el-form-item label="专业职称" class="example-block1-font0" style="left: 75px;top: 397px">
          <el-input v-model="formData.professionaltitle" clearable class="input-el-input2"/>
        </el-form-item>
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
        <el-form-item label="邮编:" class="example-block1-font" style="left: 383px;top: 397px">
          <el-input v-model.number="formData.zip" clearable class="input-el-input"/>
        </el-form-item>
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
        <el-form-item label="部门经理" class="example-block1-font0" style="left: 662px;top: 397px">
          <el-select v-model="formData.departmentmanager" clearable class="input-el-select1">
            <el-option v-for="(item,key) in departmentmanagerOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
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
        <el-form-item label="社保状态" class="example-block1-font0" style="left: 954px;top: 397px">
          <el-select v-model="formData.socialsecurity" clearable class="input-el-select1">
            <el-option v-for="(item,key) in socialsecurityOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
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
        <el-form-item label="个人技能" class="example-block1-font0" style="left: 1246px;top: 397px">
          <el-input v-model="formData.personalskills" clearable class="input-el-input2"/>
        </el-form-item>
      </div>
    </div>
    </el-form>
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
    <el-form ref="form">
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
        <el-form-item label="入职日期" class="example-block1-font0" style="left: 263px;top: 673px">
          <el-date-picker v-model="formData.dateofentry" type="date" clearable class="input-el-input0" style="width: 400px"></el-date-picker>
        </el-form-item>
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
        <el-form-item label="身份证号" class="example-block1-font0" style="left: 1073px;top: 673px">
          <el-input v-model="formData.socialsecuritynumber" clearable class="input-el-input0" style="width: 400px"/>
        </el-form-item>
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
        <el-form-item label="出生日期" class="example-block1-font0" style="left: 263px;top: 734px">
          <el-date-picker v-model="formData.dateofbirth" type="date" clearable class="input-el-input0" style="width: 400px"></el-date-picker>
        </el-form-item>
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
        <el-form-item label="毕业时间" class="example-block1-font0" style="left: 1073px;top: 734px">
          <el-date-picker v-model="formData.graduationtime" clearable type="date" class="input-el-input0" style="width: 400px"></el-date-picker>
        </el-form-item>
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
        <el-form-item label="户口所在地" class="example-block1-font0" style="left: 250px;top: 797px">
          <el-input v-model="formData.householdregistration" clearable class="input-el-input1" style="width: 430px"/>
        </el-form-item>
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
        <el-form-item label="紧急联系人" class="example-block1-font0" style="left: 1055px;top: 797px">
          <el-input v-model="formData.emergencycontacts" clearable class="input-el-input1" style="width: 430px"/>
        </el-form-item>
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
        <el-form-item label="手机号码" class="example-block1-font0" style="left: 270px;top: 858px">
          <el-input v-model.number="formData.number" clearable class="input-el-input1" style="width: 430px"/>
        </el-form-item>
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
    </el-form>
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
    <el-form ref="form">
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
        <el-form-item label="身份证详细地址" class="example-block1-font0" style="left: 490px;top: 1240px">
          <el-input v-model="formData.idaddress" clearable class="input-el-input4" style="width: 780px"/>
        </el-form-item>
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
        <el-form-item label="现住详细地址" class="example-block1-font0" style="left: 510px;top: 1305px">
          <el-input v-model="formData.address" clearable class="input-el-input4" style="width: 780px"/>
        </el-form-item>
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
        <el-form-item label="紧急联系人电话" class="example-block1-font0" style="left: 490px;top: 1370px">
          <el-input v-model.number="formData.ec_number" clearable class="input-el-input4" style="width: 780px"/>
        </el-form-item>
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
        <el-form-item label="与联系人的关系" class="example-block1-font0" style="left: 490px;top: 1434px">
          <el-input v-model="formData.ec_relationship" clearable class="input-el-input4" style="width: 780px"/>
        </el-form-item>
      </div>
    </div>
    </el-form>
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
          style="position: relative; left: 60px; top: 0%"
      >
      <!--        取消按钮样式-->
      <img
          class="css-img24"
          src="@/image/24.png"
          alt="元素加载失败"
          style="position: relative; left: 1255px; top: 0%"
      >
      <div class="dialog-footer">
        <el-button size="small" @click="closeDialog">取 消</el-button>
        <el-button size="small" type="primary" @click="enterDialog">确 定</el-button>
      </div>
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

//导入图像包
import ChooseImg from '@/components/chooseImg/index.vue'

// 全量引入格式化工具 请按需保留
import { getDictFunc, formatDate, formatBoolean, filterDict } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref } from 'vue'
import { Search } from '@element-plus/icons-vue'
import { useUserStore } from '@/pinia/modules/user'
// 页面跳转
import { useRouter } from 'vue-router'

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

getTableData()//控制搜索页不自动加载数据

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







const fits = [ 'cover' ]
const url =
    '/src/assets/logo.png'



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
  //--el-input-bg-color: #f94949;
  --el-input-border: none;
  font-size: 22px;
  left: 16px;
  padding-left: 16px;
  --el-text-color-regular: #353755;
  top: -4px;
}
.input-el-input0{
  --el-input-bg-color: transparent;
  //--el-input-bg-color: #f94949;
  --el-input-border: none;
  font-size: 20px;
  left: 38px;
  padding-left: 0px;
  --el-text-color-regular: #353755;
  top: -4px;
  --el-text-color-placeholder: transparent;
}
.input-el-input1{
  --el-input-bg-color: transparent;
  //--el-input-bg-color: #f94949;
  --el-input-border: none;
  font-size: 20px;
  padding-left: 25px;
  --el-text-color-regular: #353755;
  top: -4px;
}
.input-el-input2{
  width: 38.6%;
  --el-input-bg-color: transparent;
  //--el-input-bg-color: #f94949;
  --el-input-border: none;
  font-size: 22px;
  padding-left: 16px;
  --el-text-color-regular: #353755;
  top: -4px;
}
.input-el-input3{
  --el-input-bg-color: transparent;
  //--el-input-bg-color: #f94949;
  --el-input-border: none;
  font-size: 20px;
  padding-left: 20px;
  --el-text-color-regular: #353755;
  top: -4px;
}
.input-el-input4{
  --el-input-bg-color: transparent;
  //--el-input-bg-color: #f94949;
  --el-input-border: none;
  font-size: 20px;
  padding-left: 115px;
  --el-text-color-regular: #353755;
  top: -4px;
}
.input-el-select{
  width: 55%;
  border:none;
  outline:none;
  --el-color-white: transparent;
  //--el-color-white: #f94949;
  --el-border-base: transparent;
  --el-text-color-placeholder: transparent;
  top: -4px;
  left: 16px;
  padding-left: 16px;
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
  --el-color-white: transparent;
  //--el-color-white: #f94949;
  --el-border-base: transparent;
  --el-text-color-placeholder: transparent;
  top: -2px;
  padding-left: 28px;
}
.input-el-select1{
  width: 55%;
  border:none;
  outline:none;
  --el-color-white: transparent;
  //--el-color-white: #f94949;
  --el-border-base: transparent;
  --el-text-color-placeholder: transparent;
  top: -4px;
  left: 16px;
  padding-left: 0px;
  --el-text-color-regular: #353755;
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


