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
		try
		{
			const response = await this.vue.$http({
				url: '/api/v1/auth/login',
				data: {username, password},
				method: 'POST'
			});

			if (response.data.code === 200)
			{
				Cookies.set(cookieToken, response.data.access_token);

				await this.store.dispatch('user/status', User.statuses.CHECKING);
				await this.store.dispatch('user/token', response.data.access_token);

				await this.check();

				return true;
			}
		} catch (e)
		{
			this.store.commit('user/token', null);

			return e.response.data.message || 'Internal server error';
		}
	}

	async logout()
	{
		Cookies.remove(cookieToken);

		await this.store.commit('user/token', null);

		await this.check();
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

		try
		{
			let {data: activeData} = await this.vue.$http({
				url: '/api/v1/auth/active',
				method: 'GET'
			});

			await this.store.dispatch('user/whoami', {
				id: activeData.id,
				username: activeData.username,
				userType: activeData.user_type,
				firstName: activeData.first_name,
				lastName: activeData.last_name
			});
		} catch (e)
		{
			await this.store.dispatch('user/logged', false);
		}

		await this.store.dispatch('user/status', User.statuses.CHECKED);
	}
}

export default User;
