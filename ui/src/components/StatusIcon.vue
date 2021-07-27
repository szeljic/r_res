<template>
	<v-icon>{{v ? 'mdi-check' : 'mdi-remove'}}</v-icon>
</template>

<script>
	export default {
		name: 'StatusIcon',
		props: {
			value: {
				default: false,
				required: true
			}
		},
		computed: {
			v()
			{
				return this.value === true || this.value === 1 || this.value !== '';
			}
		}
	};
</script>
