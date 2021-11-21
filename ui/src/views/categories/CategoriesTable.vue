<template>
	<v-container>
		<v-row dense>
			<v-col>
				<v-toolbar dense elevation="2">
					<v-toolbar-title>Kategorije</v-toolbar-title>
					<v-spacer></v-spacer>
					<v-btn icon @click="showForm()" v-if="!form.show">
						<v-icon>mdi-plus</v-icon>
					</v-btn>
					<v-btn icon @click.prevent="fetch" v-if="!form.show">
						<v-icon>mdi-refresh</v-icon>
					</v-btn>
				</v-toolbar>
			</v-col>
		</v-row>

		<v-row dense v-if="form.show">
			<v-col>
				<category-form-component
					:id="form.id"
					@success="success"
					@close="form.show = false"
				></category-form-component>
			</v-col>
		</v-row>

		<v-row dense v-if="!form.show">
			<v-col>
				<v-data-table
					:headers="headers"
					:items="items"
					:server-items-length="total"
					class="elevation-2"
					no-data-text="Nema podataka"
					no-results-text="Nema rezultata"
					:loading="loading"
					:multi-sort="false"
					:footer-props="{
						itemsPerPageText: 'Redova po stranici',
						pageText: '{0}-{1} od {2}'
					}"
					:options.sync="options"
					@update:options="onChangeOptions"
				>
					<template v-slot:item="{item}">
						<tr>
							<td class="text-center">{{ item.id }}</td>
							<td>{{ item.name }}</td>
							<td>{{ item.description }}</td>
							<td>{{ item.user.first_name + ' ' + item.user.last_name }}</td>
							<td>{{ $dateFormatL18n(new Date(item.created_at), true) }}</td>
							<td class="text-center">
								<table-menu-btn>
									<v-list dense>
										<v-list-item-group>
											<v-list-item @click.prevent="showForm(item)">
												<v-list-item-icon>
													<v-icon>mdi-pencil</v-icon>
												</v-list-item-icon>
												<v-list-item-content>
													<v-list-item-title>Uredi</v-list-item-title>
												</v-list-item-content>
											</v-list-item>

											<v-list-item @click.prevent="showDelete(item)">
												<v-list-item-icon>
													<v-icon>mdi-delete-forever-outline</v-icon>
												</v-list-item-icon>
												<v-list-item-content>
													<v-list-item-title>Briši</v-list-item-title>
												</v-list-item-content>
											</v-list-item>
										</v-list-item-group>
									</v-list>
								</table-menu-btn>
							</td>
						</tr>
					</template>
				</v-data-table>
			</v-col>
		</v-row>

		<delete-dialog
			:show.sync="deleteDialog.show"
			:url="'/api/v1/categories/' + deleteDialog.id"
			@success="fetch()"
		></delete-dialog>

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
				options: {
					itemsPerPage: 10,
					page: 1
				},
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
					value: 'action',
					sortable: false,
					filterable: false,
					width: 60,
					align: 'center'
				}],
				items: [],
				total: null,
				loading: false,
				initLoad: false,
				form: {
					show: false,
					id: null
				},
				deleteDialog: {
					show: false,
					id: null
				}
			};
		},
		async created()
		{
			this.initLoad = true;
			await this.fetch();
			this.initLoad = false;
		},
		methods: {
			async fetch()
			{
				this.loading = true;

				let url = new URLSearchParams();

				if (Object.keys(this.options).length > 0)
				{
					url.append('paginate-by', this.options.itemsPerPage);
					url.append('page', this.options.page);

					if (this.options.sortBy && this.options.sortBy.length > 0)
					{
						url.append('sort-by', this.options.sortBy[0]);
						url.append('order', this.options.sortDesc[0] === true ? 'desc' : 'asc');
					}
				}

				try
				{
					const {data} = await this.$http({
						url: '/api/v1/categories/?' + url.toString()
					});

					this.total = data.total;

					this.items = data.items || [];
				} catch (e)
				{
					console.warn(e);
				} finally
				{
					this.loading = false;
				}
			},
			showForm(item)
			{
				this.form.id = item ? item.id : null;
				this.form.show = true;
			},
			success()
			{
				this.fetch();

				this.form.show = false;
			},
			showDelete(item)
			{
				this.deleteDialog.id = item ? item.id : null;

				this.deleteDialog.show = true;
			},
			onChangeOptions()
			{
				if (!this.initLoad)
				{
					this.fetch();
				}
			}
		}
	};
</script>
