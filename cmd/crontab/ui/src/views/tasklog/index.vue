<template>
  <div class="system-dic-container layout-padding">
    <!-- <el-card shadow="hover" class="layout-padding-auto"> -->
      <div class="system-user-search mb15">
        <el-input v-model="state.searchField.clientId" size="default" placeholder="客户端"
          style="max-width: 180px"></el-input>
        <el-input v-model="state.searchField.jobId" size="default" placeholder="任务ID" style="max-width: 180px"></el-input>
        <el-input v-model="state.searchField.name" size="default" placeholder="任务名称" style="max-width: 180px"></el-input>
        <el-input v-model="state.searchField.executorHandler" size="default" placeholder="任务标识"
          style="max-width: 180px"></el-input>
        <el-select v-model="state.searchField.result" placeholder="执行结果" clearable size="mini">
          <el-option label="success" :value="200"></el-option>
          <el-option label="error" :value="500"></el-option>
        </el-select>
        <el-button size="default" type="primary" class="ml10" @click="getTableData()">
          <el-icon>
            <ele-Search />
          </el-icon>
          查询
        </el-button>
      </div>
      <el-table :data="state.tableData.data" v-loading="state.tableData.loading" style="width: 100%">
        <el-table-column type="index" label="序号" width="50" />
        <el-table-column prop="clientId" label="客户端" show-overflow-tooltip></el-table-column>
        <el-table-column prop="jobId" label="任务ID" show-overflow-tooltip></el-table-column>
        <el-table-column prop="name" label="任务名称" show-overflow-tooltip></el-table-column>
        <el-table-column prop="executorHandler" label="任务标识" show-overflow-tooltip></el-table-column>
        <el-table-column prop="prevTime" label="上次执行时间" show-overflow-tooltip></el-table-column>
        <el-table-column prop="nextTime" label="下次执行时间" show-overflow-tooltip></el-table-column>
        <el-table-column prop="result" label="执行结果" show-overflow-tooltip>
          <template #default="scope">
            <el-tag type="success" v-if="scope.row.result == 200">Success</el-tag>
            <el-tag type="danger" v-else>Error</el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="msg" label="执行详情" show-overflow-tooltip></el-table-column>
        <el-table-column prop="ctime" label="创建时间" show-overflow-tooltip></el-table-column>
      </el-table>
    <!-- </el-card> -->
    <el-pagination @size-change="onHandleSizeChange" @current-change="onHandleCurrentChange" class="mt15"
        :pager-count="5" :page-sizes="[10, 50, 100]" v-model:current-page="state.tableData.param.pageNum" background
        v-model:page-size="state.tableData.param.pageSize" layout="total, sizes, prev, pager, next, jumper"
        :page-count="state.tableData.total">
      </el-pagination>
    <!-- <TaskDialog ref="dicDialogRef" @refresh="getTableData()" /> -->
  </div>
</template>

<script setup lang="ts" name="systemDic">
import { reactive, onMounted } from 'vue';
import { useTasklogApi } from '/@/api/tasklog/index';

// 引入 api 请求接口
const tasklogApi = useTasklogApi();

const state = reactive({
  searchField: {
    result: '',
    clientId: '',
    jobId: '',
    name: '',
    executorHandler: '',
  },
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
  tasklogApi.list({
    result: state.searchField.result,
    clientId: state.searchField.clientId,
    jobId: state.searchField.jobId,
    name: state.searchField.name,
    executorHandler: state.searchField.executorHandler,
    pageNum: state.tableData.param.pageNum,
    pageSize: state.tableData.param.pageSize
  }).then((res) => {
    console.log(res)
    if (res.code == 200 && res.items && res.items.length > 0) {
      state.tableData.data = res.items;
      state.tableData.total = res.totalPage;
    } else {
      state.tableData.data = []
    }
  }).finally(() => {
    setTimeout(() => {
      state.tableData.loading = false;
    }, 500);
  })
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
