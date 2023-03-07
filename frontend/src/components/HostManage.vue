<template>
  <main>
    <el-button :type="btnType" @click="dialogVisible = true">{{title}}</el-button>
    <el-dialog
        v-model="dialogVisible"
        :title="title"
        width="60%"
    >
      <el-form :model="form" label-width="120px">

        <el-form-item label="名称">
          <el-input placeholder="名称" v-model="form.name" />
        </el-form-item>
        <el-form-item>
          <el-button v-if="data === undefined" type="primary" @click="createHost">创建</el-button>
          <el-button v-else type="primary" @click="editHost">保存</el-button>
          <el-button @click="dialogVisible = false">取消</el-button>
        </el-form-item>
      </el-form>
    </el-dialog>

  </main>
</template>

<script setup>
import {ref, reactive, watch, watchEffect} from "vue";
import {HostCreate} from "../../wailsjs/go/main/App.js";
import {HostEdit} from "../../wailsjs/go/main/App.js";
import {ElNotification} from "element-plus";

let props = defineProps(['title', 'btnType', 'data'])
const emits = defineEmits(['emit-host-list'])
const dialogVisible = ref(false)

let form = reactive({})
if (props.data !== undefined) {
  form = props.data
}

function createHost() {
  HostCreate(form).then(res => {
    dialogVisible.value = false
    if (res.code !== 200) {
      ElNotification({
        title:res.msg,
        type: "error",
      })
      return
    }
    ElNotification({
      title:res.msg,
      type: "success",
    })
    emits('emit-host-list')
  })
}

function editHost() {
  console.log(form);
  HostEdit(form).then(res => {
    dialogVisible.value = false
    if (res.code !== 200) {
      ElNotification({
        title:res.msg,
        type: "error",
      })
      return
    }
    ElNotification({
      title:res.msg,
      type: "success",
    })
    emits('emit-host-list')
  })
}
</script>

<style scoped>
</style>