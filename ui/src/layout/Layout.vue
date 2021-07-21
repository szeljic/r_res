<template>
	<component :is="component" :id="component"></component>
</template>

<script>
	import Blank from '@/layout/Blank';
	import Public from '@/layout/Public';
	import Application from '@/layout/Application';

	export default {
		name: 'Layout',
		components: {
			Blank,
			Public,
			Application
		},
		computed: {
			component()
			{
				return 'Public';
			}
		}
	};
</script>
