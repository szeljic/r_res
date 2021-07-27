class Common {
	install(Vue)
	{
		Vue.component('message-box', () => import('../components/MessageBox'));
		Vue.component('delete-dialog', () => import('../components/DeleteDialog'));
		Vue.component('status-icon', () => import('../components/StatusIcon'));
	}
}

export default Common;
