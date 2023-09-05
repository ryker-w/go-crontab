import request from '/@/utils/request';

const api = "/api/admin"
/**
 * @method updateInfo 更新任务信息
 * @method start 启动
 * @method kill 终止
 */
export function useTaskApi() {
	return {
		list: (params?: object) => {
			return request({
				url: api +'/task',
				method: 'get',
				params,
			});
		},
		update: (data?: object) => {
			return request({
				url: api + '/task',
				method: 'put',
				data,
			});
		},
		updateSpec: (data?: object) => {
			return request({
				url: api + '/task/spec',
				method: 'post',
				data,
			});
		},
		start: (data?: object) => {
			return request({
				url: api +'/task/start',
				method: 'post',
				data,
			});
		},
    kill: (data?: object) => {
			return request({
				url: api +'/task/kill',
				method: 'post',
				data,
			});
		},
		exec: (data?: object) => {
			return request({
				url: api +'/task/exec',
				method: 'post',
				data,
			});
		}
	};
}