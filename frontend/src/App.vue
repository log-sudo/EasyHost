<script setup>
import Host from './components/Host.vue'
import HostManage from "./components/HostManage.vue";
import {ref} from "vue";


let flushFlag = ref(true)

function flushHostList() {
  flushFlag.value = !flushFlag.value
}
</script>

<template>
  
  <el-row>
    <el-col :span="24" style="height: 100vh; padding: 12px">
      <div style="margin-bottom: 12px; padding-left: 8px;" >
        <HostManage title="新增" btn-type="primary" @emit-host-list="flushHostList"/>
      </div>
      <Host :flush="flushFlag"/>
    </el-col>
  </el-row>
</template>
<style>

</style>
