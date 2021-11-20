import User from '@/plugins/user';

export default {
	namespaced: true,
	state: {
		status: User.statuses.INIT,
		token: null,
		logged: false,
		whoami: null
	},
	mutations: {
		token(state, v = null)
		{
			state.token = v;
		},
		status(state, v)
		{
			state.status = v;
		},
		logged(state, v)
		{
			state.logged = v;
		},
		whoami(state, v)
		{
			state.whoami = v;
		}
	},
	actions: {
		token({commit, dispatch}, v)
		{
			commit('token', v);
			dispatch('logged', v !== null && v !== undefined);
		},
		status({commit}, v)
		{
			commit('status', v);
		},
		logged({commit}, v)
		{
			commit('logged', v);
		},
		whoami({commit}, v)
		{
			commit('whoami', v);
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
		},
		whoami(state)
		{
			return state.whoami;
		}
	}
};
