import Cookies from 'js-cookie';

const cookieToken = 'c-token';

class User
{
	static statuses = {
		INIT: 'INIT',
		CHECKING: 'CHECKING',
		CHECKED: 'CHECKED'
	};

	install(Vue)
	{
		this.vue = new Vue();
		this.store = this.vue.$store;

		this.store.dispatch('user/token', Cookies.get(cookieToken));

		this.$statuses = User.statuses;

		Vue.prototype.$user = this;
	}

	async login(username, password)
	{
		const response = await this.vue.$http({
			url: '/api/v1/auth/login',
			data: {username, password},
			method: 'POST'
		});

		if (response.data.code === 200)
		{
			Cookies.set(cookieToken, response.data.access_token);

			await this.store.dispatch('user/token', response.data.access_token);
		} else if (response.data.code === 403)
		{
			this.store.commit('user/token', null);
		}
	}

	async check()
	{
		await this.store.dispatch('user/status', User.statuses.CHECKING);

		const token = this.store.getters['user/token'];

		if (!token)
		{
			await this.store.dispatch('user/logged', false);
			await this.store.dispatch('user/status', User.statuses.CHECKED);

			return;
		}

		this.vue.$http({
			url: '/api/v1/auth/check',
			method: 'GET'
		}).then(() =>
		{
		}).catch(async () =>
		{
			await this.store.dispatch('user/logged', false);
		}).finally(async () =>
		{
			await this.store.dispatch('user/status', User.statuses.CHECKED);
		});
	}
}

export default User;
