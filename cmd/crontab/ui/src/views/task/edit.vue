<template>
	<div class="system-dic-dialog-container">
		<el-dialog :title="state.dialog.title" v-model="state.dialog.isShowDialog" width="769px">
			<!-- <el-alert title="半成品，交互过于复杂，请自行扩展！" type="warning" :closable="false" class="mb20"> </el-alert> -->
			<el-form ref="dicDialogFormRef" :model="state.ruleForm" size="default" label-width="90px">
				<el-row :gutter="35">
          <el-col :xs="24" :sm="12" :md="12" :lg="12" :xl="12" class="mb20">
						<el-form-item label="任务名称">
							<el-input v-model="state.ruleForm.name" placeholder="请输入任务名称" clearable></el-input>
						</el-form-item>
					</el-col>
          <el-col :xs="24" :sm="12" :md="12" :lg="12" :xl="12" class="mb20">
						<el-form-item label="任务名称">
							<el-input v-model="state.ruleForm.name" placeholder="请输入任务名称" clearable></el-input>
						</el-form-item>
					</el-col>
				</el-row>
			</el-form>
			<template #footer>
				<span class="dialog-footer">
					<el-button @click="onCancel" size="default">取 消</el-button>
					<el-button type="primary" @click="onSubmit" size="default">{{ state.dialog.submitTxt }}</el-button>
				</span>
			</template>
		</el-dialog>
	</div>
</template>

<script setup lang="ts" name="taskDialog">
import { reactive, ref } from 'vue';

// 定义子组件向父组件传值/事件
const emit = defineEmits(['refresh']);

// 定义变量内容
const dicDialogFormRef = ref();
const state = reactive({
	ruleForm: {
		clientId: '', 
		name: '', 
	},
	dialog: {
		isShowDialog: false,
		type: '',
		title: '',
		submitTxt: '',
	},
});

// 打开弹窗
const openDialog = (type: string, row: any) => {
	if (type === 'edit') {
		state.ruleForm = row;
		state.dialog.title = '修改';
		state.dialog.submitTxt = '修 改';
    state.dialog.type = 'edit'
	} else {
		state.dialog.title = '新增';
		state.dialog.submitTxt = '新 增';
    state.dialog.type = 'add'
		// 清空表单，此项需加表单验证才能使用
		// nextTick(() => {
		// 	dicDialogFormRef.value.resetFields();
		// });
	}
	state.dialog.isShowDialog = true;
};
// 关闭弹窗
const closeDialog = () => {
	state.dialog.isShowDialog = false;
};
// 取消
const onCancel = () => {
	closeDialog();
};
// 提交
const onSubmit = () => {
	closeDialog();
	emit('refresh');
	if (state.dialog.type === 'edit') {
    
  } else {

  }
};
// 新增行
// const onAddRow = () => {
// 	state.ruleForm.list.push({
// 		id: Math.random(),
// 		label: '',
// 		value: '',
// 	});
// };
// 删除行
// const onDelRow = (k: number) => {
// 	state.ruleForm.list.splice(k, 1);
// };

// 暴露变量
defineExpose({
	openDialog,
});
</script>
