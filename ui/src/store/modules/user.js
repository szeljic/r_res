import User from '@/plugins/user';

export default {
	namespaced: true,
	state: {
		status: User.statuses.INIT,
		token: null,
		logged: false
	},
	mutations: {
		token(state, v = null)
		{
			state.token = v;
			state.logged = v !== null;
		},
		status(state, v)
		{
			state.status = v;
		},
		logged(state, v)
		{
			state.logged = v;
		}
	},
	actions: {
		token({commit}, v)
		{
			commit('token', v);
		},
		status({commit}, v)
		{
			commit('status', v);
		},
		logged({commit}, v)
		{
			commit('logged', v);
		}
	},
	getters: {
		token(state)
		{
			return state.token;
		},
		logged(state)
		{
			return state.logged;
		},
		status(state)
		{
			return state.status;
		}
	}
};
