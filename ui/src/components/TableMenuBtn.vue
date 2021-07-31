<template>
	<v-menu transition="slide-y-transition" :disabled="disabled">
		<template v-slot:activator="{ on }">
			<v-btn small text icon class="mr-0 ml-0" v-on="on" :disabled="disabled">
				<v-icon>mdi-dots-vertical</v-icon>
			</v-btn>
		</template>
		<slot></slot>
	</v-menu>
</template>

<script>
	export default {
		name: 'table-menu-btn',
		props: {
			disabled: {
				type: Boolean,
				default: false
			}
		}
	};
</script>
