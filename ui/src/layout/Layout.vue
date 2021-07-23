<template>
	<component :is="component" :id="component"></component>
</template>

<script>
	import Blank from '@/layout/Blank';
	import Public from '@/layout/Public';
	import Application from '@/layout/Application';

	const layouts = {
		Blank: 'Blank',
		Public: 'Public',
		Application: 'Application'
	};

	export default {
		name: 'Layout',
		components: {
			Blank,
			Public,
			Application
		},
		async created()
		{
			await this.$user.check();
		},
		computed: {
			component()
			{
				const store = this.$store,
					token = store.getters['user/token'],
					status = store.getters['user/status'],
					logged = store.getters['user/logged'],
					user = this.$user;

				if (token === null)
				{
					return layouts.Public;
				}

				if (status === user.$statuses.INIT || status === user.$statuses.CHECKING)
				{
					return layouts.Blank;
				}

				if (status === user.$statuses.CHECKED && logged === true)
				{
					return layouts.Application;
				}

				return layouts.Public;
			}
		}
	};
</script>
