<script setup>
import {h, ref, watch,onMounted,getCurrentInstance } from "vue"
import {HostList,HostCreate,HostEdit,WriteHost, HostDelete} from '../../wailsjs/go/main/App'
import {ElNotification} from "element-plus"
import { basicSetup, EditorView } from "codemirror";
import { EditorState } from "@codemirror/state";
let list = ref()
let props = defineProps(['flush'])
watch(props, (newFlush)=>{
  hostList()
})
let clickIndex=ref(0);
let host=ref(
  {
    identity:"",
    name:"",
    info:"",
    use:""
  }
);

function  hostList (){
  HostList().then(res => {
    if (res.code !== 200) {
      ElNotification({
        title:res.msg,
        type: "success",
      })
    }
    list.value = res.data;

    for(var i=0;i<list.value.length;i++){
        if(list.value[i].use=="1"){
          host.value=list.value[i];
          clickIndex.value=i;
          updateValue(list.value[i].info)
        }
    }
    
  })
}

function click(index,item) {
  item.use="1";
  host.value=item;
  HostEdit(host.value).then(res => {
    if (res.code !== 200) {
      ElNotification({
        title:res.msg,
        type: "error",
      })
      return
    }
    ElNotification({
      title:"切换成功",
      type: "success",
    })
  })
  clickIndex.value=index;
  WriteHost(host.value);
  updateValue(host.value.info)
}
 hostList()

 function hostDelete(item) {
    if(item.identity=='1'){
      ElNotification({
      title:"系统默认不能删除",
      type: "warn",
      })
      return;
    }
    if(item.use=='1'){
      ElNotification({
      title:"正在使用不能删除",
      type: "warn",
    })
    return;
  }
  HostDelete(item.identity).then(res => {
    hostList()
    ElNotification({
      title:"删除成功",
      type: "success",
    })

  })
}


function onchange() {
  host.value.info=editorView.value.state.doc.toString();
  HostEdit(host.value).then(res => {
    console.log(res)
  })
  WriteHost(host.value);
}



onMounted(() => {
    initEditor();
});


const editorRef = ref();

const editorView = ref();

const initEditor = () => {

if (typeof editorView.value !== "undefined") {
   editorView.value.destroy();
}

const fixedHeightEditor = EditorView.theme({
  "&": {height: "100%"},
  ".cm-scroller": {overflow: "auto"}
})

let startState = EditorState.create({
  doc:"",
  extensions: [
    basicSetup,
    fixedHeightEditor,
    EditorView.updateListener.of(update => onchange())
  ],
});




if (editorRef.value) {
    editorView.value = new EditorView({
      state: startState,
      parent: editorRef.value,
    });
  }
}
  
const updateValue = (value)=>{
  editorView.value.dispatch({changes: {from: 0, to:editorView.value.state.doc.length, insert: value}})
}

</script>

<template>
  <main>
    <el-container>
      <el-container>
        <el-aside width="200px">
          <el-row >
            <el-col :span="24" class="lightgreen-box"  style="padding:6px">
              <el-card v-for="(item,index) in list" :key="item.identity"
                style="background-color:  rgb(215 204 204); padding: 0px"
                shadow="hover" 
                @click.native="click(index,item)"
                :class="{active:clickIndex==index}">
                <div>
                    <div class="bottom clearfix">
                        <span class="name">{{ item.name }}</span>
                        <el-popconfirm title="确认删除?"  @confirm="hostDelete(item)">
                          <template #reference>
                            <el-button link  type="txt" :icon="Delete" class="button" @click.stop></el-button>
                          </template>
                        </el-popconfirm>
                    </div>
                </div>
              </el-card>
            </el-col>
          </el-row>   


        </el-aside>
        <el-container>
          <el-main style="padding: 5px;">
            <el-card id="editor" style="padding: 0px" >
              <div ref="editorRef" class="editor-main"></div>
            </el-card>
          </el-main>
        </el-container>
      </el-container>
    </el-container>
  </main>
</template>

<style scoped>
.name {
    font-size: 13px;
    color: #020202;
}
.button {
    padding: 0;
    float: right;
    color: #0e5194;
  }
.active {
background-color:#ffffff!important;
border: 1px solid #0e5194;
}
.el-card ::v-deep .el-card__body {
  padding: 3px;
}
.el-card {
  margin: 3px;
}
.el-main{
  padding: 10px;
}

.w-e-menu
{
display: none;
}

.editor-main{
  width: 100vw;
  height: 90vh;
}



</style>

<!-- <style lang="scss">
#code{
  overflow-y: scroll !important;
  .CodeMirror{
      overflow-y: scroll !important;
      height: auto !important;
    }

}
</style> -->
