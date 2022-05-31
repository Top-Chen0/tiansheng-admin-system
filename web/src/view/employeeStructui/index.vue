<template>
  <div id="uploadlist">
    <el-upload
      :action="`${path}/fileUploadAndDownload/upload`"
      :headers="{ 'x-token': userStore.token }"
      :on-error="uploadError"
      :file-list="fileList"
      :on-success="handleImageSuccess"
      :before-upload="beforeImageUpload"
      :on-remove="handleRemove"
      :limit="limit"
      list-type="picture"
      :auto-upload="true"
    >
      <el-tooltip
        class="box-item"
        effect="dark"
        :content="tips"
        placement="right-start"
      >
        <el-button size="small" type="primary" plain>压缩上传</el-button>
      </el-tooltip>
    </el-upload>
  </div>
</template>
<script setup>
import ImageCompress from '@/utils/image'
import { ref, reactive } from 'vue'
import { deleteFile } from '@/api/fileUploadAndDownload'
import { ElMessage } from 'element-plus'
import { useUserStore } from '@/pinia/modules/user'

const emit = defineEmits(['change', 'upload-type'])
const props = defineProps({
  imageUrl: {
    type: String,
    default: ''
  },
  uploadType: {
    type: String,
    default: ''
  },
  fileSize: {
    type: Number,
    default: 2048 // 2M 超出后执行压缩
  },
  maxWH: {
    type: Number,
    default: 1920 // 图片长宽上限
  },
  tips: {
    type: String,
    default: ''
  },
  showList: {
    type: Array,
    default: function() {
      return []
    }
  },
  limit: {
    type: Number,
    default: 4
  },
})

const path = ref(import.meta.env.VITE_BASE_API)
const fileList = ref([])
const obj = reactive({})
const userStore = useUserStore()

// 检查上传图片的格式、尺寸
const beforeImageUpload = (file) => {
  const isJPG = file.type === 'image/jpeg'
  const isPng = file.type === 'image/png'
  if (!isJPG && !isPng) {
    ElMessage.error('上传头像图片只能是 jpg或png 格式!')
    return false
  }
  const isRightSize = file.size / 1024 < props.fileSize
  if (!isRightSize) {
    // 压缩
    const compress = new ImageCompress(file, props.fileSize, props.maxWH)
    return compress.compress()
  }
  return isRightSize
}
const handleImageSuccess = (res) => {
  const { data } = res
  if (data.file) {
    console.log('handleImageSuccess ', data)
    fileList.value.push({
      name: data.file.name,
      url: data.file.url,
      id: data.file.ID
    })
    // fileList.value.push(data.file.ID)
    obj.value = {
      uploadType: props.uploadType,
      fileList: fileList.value// JSON.parse(JSON.stringify(fileList))
    }
    emit('change', obj)
    console.log('handleImageSuccess》》:', fileList.value)
    console.log('handleImageSuccess》》》:', obj)
  }
}

const handleRemove = async(file, fileList) => {
  console.log('removed fileList>>> ', file, fileList)
  const res = await deleteFile(file)
  console.log('deleteFile>>> res>>>', res)
  if (res.code === 0) {
    ElMessage({
      type: 'success',
      message: '删除成功!'
    })
    obj.value = {
      uploadType: props.uploadType,
      fileList: JSON.parse(JSON.stringify(fileList))
    }
    emit('change', obj)
    // const key = fileList.findIndex(file)
    // console.log('555:', key)
    // fileList.splice(key, 1)
    console.log('666:', fileList)
    console.log('666:', obj)
    // this.fileList = fileList
    // if (this.tableData.length === 1 && this.page > 1) {
    //   this.page--
    // }
    // this.getTableData()
  }
  // }).catch(() => {
  //   ElMessage({
  //     type: 'info',
  //     message: '已取消删除'
  //   })
  // })
}
// 获取图片列表
// const getimgdata = () => {
//   fileList = props.showList
//   for (const i in fileList) {
//     this.picList.push(this.fileList[i].url)
//   }
// }
const uploadError = () => {
  this.$message({
    type: 'error',
    message: '上传失败'
  })
  this.fullscreenLoading = false
}
</script>
<script>

export default {
  name: 'UploadList',
  methods: {

  }
}
</script>
<style lang="scss" scoped>
</style>
