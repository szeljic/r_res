<template>
	<v-container>
		<v-row dense>
			<v-col>
				<v-toolbar dense elevation="2">
					<v-toolbar-title>Kategorije</v-toolbar-title>
					<v-spacer></v-spacer>
					<v-btn icon @click="showForm()">
						<v-icon>mdi-plus</v-icon>
					</v-btn>
					<v-btn icon @click.prevent="fetch">
						<v-icon>mdi-refresh</v-icon>
					</v-btn>
				</v-toolbar>
			</v-col>
		</v-row>
		<v-row dense>
			<v-col>
				<v-data-table
					:headers="headers"
					:items="items"
					:server-items-length="total"
					:items-per-page="10"
					class="elevation-2"
					no-data-text="Nema podataka"
					no-results-text="Nema rezultata"
					:loading="loading"
				>
					<template v-slot:item="{item}">
						<tr>
							<td class="text-center">{{ item.id }}</td>
							<td>{{ item.name }}</td>
							<td>{{ item.description }}</td>
							<td>{{ item.created_by }}</td>
							<td>{{ item.created_at }}</td>
							<td class="text-center">
								<v-btn small :to="`/kategorije/uredi/${item.id}`" icon>
									<v-icon>mdi-pencil</v-icon>
								</v-btn>
							</td>
						</tr>
					</template>
				</v-data-table>
			</v-col>
		</v-row>

		<v-dialog v-model="form.show" persistent max-width="640">
			<category-form-component></category-form-component>
		</v-dialog>
	</v-container>
</template>

<script>
	import CategoryFormComponent from '@/views/categories/CategoryFormComponent';

	export default {
		name: 'UsersTable',
		components: {
			CategoryFormComponent
		},
		data()
		{
			return {
				headers: [{
					text: '#',
					value: 'ID',
					width: 100,
					align: 'center'
				}, {
					text: 'Naziv',
					value: 'name'
				}, {
					text: 'Kratak opis',
					value: 'description'
				}, {
					text: 'Napravio',
					value: 'created_by',
					width: 300
				}, {
					text: 'Datum pravljenja',
					value: 'created_at'
				}, {
					text: '',
					value: null,
					sortable: false,
					filterable: false,
					width: 104,
					align: 'center'
				}],
				items: [],
				total: null,
				loading: false,
				form: {
					show: false,
					id: null
				}
			};
		},
		created()
		{
			this.fetch();
		},
		methods: {
			async fetch()
			{
				this.loading = true;

				const {data} = await this.$http({
					url: '/api/v1/categories'
				});

				this.total = data.total;
				this.items = data.items || [];

				this.loading = false;
			},
			showForm()
			{
				this.form.id = null;
				this.form.show = true;
			}
		}
	};
</script>
