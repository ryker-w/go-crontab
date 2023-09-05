import request from '/@/utils/request';

const api = "/api/admin"
/**
 * @method updateInfo 更新任务信息
 * @method start 启动
 * @method kill 终止
 */
export function useTasklogApi() {
	return {
		list: (params?: object) => {
			return request({
				url: api +'/tasklog',
				method: 'get',
				params,
			});
		},
	};
}