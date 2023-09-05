<template>
  <div class="system-dic-container layout-padding">
    <el-card shadow="hover" class="layout-padding-auto">
      <div class="system-user-search mb15">
        <el-input size="default" placeholder="请输入名称" style="max-width: 180px"> </el-input>
        <el-button size="default" type="primary" class="ml10" @click="getTableData()">
          <el-icon>
            <ele-Search />
          </el-icon>
          查询
        </el-button>
        <el-button size="default" type="success" class="ml10" @click="onOpenAdd('add')">
          <el-icon>
            <ele-FolderAdd />
          </el-icon>

        </el-button>
      </div>
      <el-table :data="state.tableData.data" v-loading="state.tableData.loading" style="width: 100%">
        <el-table-column type="index" label="序号" width="50" />
        <el-table-column prop="clientId" label="应用模块" show-overflow-tooltip></el-table-column>
        <el-table-column prop="jobId" label="任务ID" show-overflow-tooltip></el-table-column>
        <el-table-column prop="name" label="任务名称" show-overflow-tooltip>
          <template #default="scope">
            <el-input v-model="scope.row.name" placeholder="无" @change="updateTaskInfo(scope.row)"></el-input>
          </template>
        </el-table-column>
        <el-table-column prop="status" label="状态" show-overflow-tooltip>
          <template #default="scope">
            <el-tag type="success" v-if="scope.row.status == 10" @click="onKillTask(scope.row)">启用</el-tag>
            <el-tag type="warning" v-else @click="onStartTask(scope.row)">未启用</el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="executorHandler" label="任务标识" show-overflow-tooltip></el-table-column>
        <el-table-column prop="spec" label="定时器" >
          <template #default="scope">
            <el-input v-model="scope.row.spec" placeholder="无" @change="changeTaskSpec(scope.row)"></el-input>
          </template>
        </el-table-column>
        <el-table-column prop="prevTime" label="上次执行时间" show-overflow-tooltip></el-table-column>
        <el-table-column prop="nextTime" label="下次执行时间" show-overflow-tooltip></el-table-column>
        <el-table-column prop="instanceNum" label="当前实例量" show-overflow-tooltip></el-table-column>
        <el-table-column prop="maxInstanceNum" label="实例限制" show-overflow-tooltip>
          <template #default="scope">
            <el-input v-model="scope.row.maxInstanceNum" placeholder="0" @change="updateTaskInfo(scope.row)"></el-input>
          </template>
        </el-table-column>
        <el-table-column label="操作" width="100">
          <template #default="scope">
            <el-button size="small" text type="primary" @click="onExecTask(scope.row)">立即执行</el-button>
            <!-- <el-button size="small" text type="primary" @click="onOpenEdit('edit', scope.row)">修改</el-button>
            <el-button size="small" text type="primary" @click="onRowDel(scope.row)">删除</el-button> -->
          </template>
        </el-table-column>
      </el-table>
      <el-pagination @size-change="onHandleSizeChange" @current-change="onHandleCurrentChange" class="mt15"
        :pager-count="5" v-model:current-page="state.tableData.param.pageNum" background
        v-model:page-size="state.tableData.param.pageSize" layout="total, prev, pager, next, jumper"
        :page-count="state.tableData.total">
      </el-pagination>
    </el-card>
    <!-- <TaskDialog ref="dicDialogRef" @refresh="getTableData()" /> -->
  </div>
</template>

<script setup lang="ts" name="systemDic">
import { defineAsyncComponent, reactive, onMounted, ref } from 'vue';
import { ElMessageBox, ElMessage } from 'element-plus';
import { useTaskApi } from '/@/api/task/index';

// 后端控制路由

// 引入 api 请求接口
const taskApi = useTaskApi();
// 引入组件
// const TaskDialog = defineAsyncComponent(() => import('/@/views/task/edit.vue'));

// 定义变量内容
const dicDialogRef = ref();
const state = reactive({
  tableData: {
    data: [],
    total: 0,
    loading: false,
    param: {
      pageNum: 1,
      pageSize: 10,
    },
  },
});

// 初始化表格数据
const getTableData = () => {
  state.tableData.loading = true;
  taskApi.list({
    pageNum: state.tableData.param.pageNum,
    pageSize: state.tableData.param.pageSize
  }).then((res) => {
    console.log(res)
    if (res.code == 200 && res.items && res.items.length > 0) {
      state.tableData.data = res.items;
      state.tableData.total = res.totalPage;
    }
  }).finally(() => {
    setTimeout(() => {
      state.tableData.loading = false;
    }, 500);
  })
};
// 打开新增弹窗
const onOpenAdd = (type: string) => {

};
// 打开修改弹窗
const onOpenEdit = (type: string, row: any) => {

};
// 删除
const onRowDel = (row: any) => {
  ElMessageBox.confirm(`此操作将删除任务：“${row.name}”，是否继续?`, '提示', {
    confirmButtonText: '确认',
    cancelButtonText: '取消',
    type: 'warning',
  })
    .then(() => {
      ElMessage.success('删除成功');
      getTableData();
    })
    .catch(() => { });
};

const changeTaskSpec = (row: any) => {
  console.log(row.spec, row);
  ElMessageBox.confirm(`此操作将修改任务：“${row.name}” 的定时器, 是否终止任务并继续?`, '提示', {
    confirmButtonText: '确认',
    cancelButtonText: '取消',
    type: 'warning',
  })
    .then(() => {
      taskApi.updateSpec({
        id: row.id,
        spec: row.spec,
      }).then((res) => {
        if (res.code == 200) {
          ElMessage.success("修改成功")
        } else {
          ElMessage.error("格式错误")
        }
      }).catch((err) => {
        console.log(err)
      }).finally(() => {
        getTableData();
      });
    })
    .catch(() => { getTableData(); });
};

const updateTaskInfo = (row: any) => {
  console.log(row.spec, row);
  ElMessageBox.confirm(`此操作将修改任务信息：“${row.name}” , 是否终止任务并继续?`, '提示', {
    confirmButtonText: '确认',
    cancelButtonText: '取消',
    type: 'warning',
  })
    .then(() => {
      taskApi.update({
        id: row.id,
        name: row.name,
        maxInstanceNum: parseInt(row.maxInstanceNum || ''),
      }).then((res) => {
        if (res.code == 200) {
          ElMessage.success("修改成功")
        } else {
          ElMessage.error("格式错误")
        }
      }).catch((err) => {
        console.log(err)
      }).finally(() => {
        getTableData();
      });
    })
    .catch(() => { getTableData(); });
};
const onStartTask = (row: any) => {
  console.log(row.spec, row);
  ElMessageBox.confirm(`此操作将启动任务：“${row.name}”, 是否继续?`, '提示', {
    confirmButtonText: '确认',
    cancelButtonText: '取消',
    type: 'warning',
  })
    .then(() => {
      taskApi.start({ id: row.id }).then((res) => {
        if (res.code == 200) {
          ElMessage.success("启动成功")
        } else {
          ElMessage.error("操作失败")
        }
      })
        .catch((err) => { console.log(err) })
        .finally(() => {
          getTableData();
        });
    })
    .catch(() => { getTableData(); });
};

const onKillTask = (row: any) => {
  console.log(row.spec, row);
  ElMessageBox.confirm(`此操作将终止任务：“${row.name}”, 是否继续?`, '提示', {
    confirmButtonText: '确认',
    cancelButtonText: '取消',
    type: 'warning',
  })
    .then(() => {
      taskApi.kill({ id: row.id }).then((res) => {
        if (res.code == 200) {
          ElMessage.success("操作成功")
        } else {
          ElMessage.error("操作失败")
        }
      })
        .catch((err) => { console.log(err) })
        .finally(() => {
          getTableData();
        });
    })
    .catch(() => { getTableData(); });
};

const onExecTask = (row: any) => {
  console.log(row.spec, row);
  if (row.status == 20) {
    ElMessage.warning("任务未启用，不可执行！")
    return
  }
  ElMessageBox.prompt(`此操作将执行一次任务：“${row.name}”, 是否继续?(可填参数)`, '提示', {
    confirmButtonText: '确认',
    cancelButtonText: '取消',
    type: 'warning',
  })
    .then((value) => {
      taskApi.exec({ id: row.id, params: value.value }).then((res) => {
        if (res.code == 200) {
          ElMessage.success("操作成功")
        } else {
          ElMessage.error("操作失败")
        }
      })
        .catch((err) => { console.log(err) })
        .finally(() => {
          getTableData();
        });
    })
    .catch(() => { getTableData(); });
};

// 分页改变
const onHandleSizeChange = (val: number) => {
  state.tableData.param.pageSize = val;
  getTableData();
};
// 分页改变
const onHandleCurrentChange = (val: number) => {
  state.tableData.param.pageNum = val;
  getTableData();
};
// 页面加载时
onMounted(() => {
  getTableData();
});
</script>
