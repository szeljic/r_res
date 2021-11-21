<template>
	<v-container>
		<v-row dense>
			<v-col>
				<v-toolbar dense elevation="2">
					<v-toolbar-title>Resursi</v-toolbar-title>
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
				<resource-form-component
					:id="form.id"
					@success="success"
					@close="form.show = false"
					:pre-category="category"
				></resource-form-component>
			</v-col>
		</v-row>

		<v-row dense v-if="!form.show">
			<v-col md="3">
				<select-category
					label="Kategorija"
					v-model="category"
					:disabled="loading"
					@change="fetch()"
					hide-details
					solo
					:categories.sync="categories"
					clearable
				></select-category>
			</v-col>
			<v-col md="3">
				<v-text-field
					label="Trazi"
					solo
					clearable
					hide-details
					v-model="search"
					@keyup="fetch()"
				></v-text-field>
			</v-col>
		</v-row>

		<v-row dense v-if="!form.show && items.length > 0">
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
					:must-sort="false"
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
							<td>{{ item.user.first_name + ' ' + item.user.last_name }}</td>
							<td>{{ $dateFormatL18n(new Date(item.created_at * 1000), true) }}</td>

							<template v-for="f in selectedSpecificFields">
								<td v-if="f.data_type === 'boolean'" :key="f.sc_name">
									<status-icon v-model="item[f.sc_name]"></status-icon>
								</td>
								<td v-else :key="f.sc_name">{{ item[f.sc_name] || '' }}</td>
							</template>

							<td class="text-center">
								<table-menu-btn :disabled="!showAction(item)">
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
			:url="'/api/v1/resources/' + deleteDialog.id"
			@success="fetch()"
		></delete-dialog>

	</v-container>
</template>

<script>
	import ResourceFormComponent from '@/views/resources/ResourceFormComponent';
	import SelectCategory from '@/components/SelectCategory';
	import {mapGetters} from 'vuex';

	export default {
		name: 'ResourcesTable',
		components: {
			ResourceFormComponent,
			SelectCategory
		},
		data()
		{
			return {
				options: {
					itemsPerPage: 10,
					page: 1
				},
				staticHeaders: [{
					text: '#',
					value: 'id',
					width: 60,
					align: 'center'
				}, {
					text: 'Naziv',
					value: 'name',
					width: 180
				}, {
					text: 'Napravio',
					value: 'user',
					width: 180
				}, {
					text: 'Datum',
					value: 'created_at',
					width: 180
				}],
				items: [],
				total: null,
				category: null,
				categories: [],
				search: null,
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
		computed: {
			...mapGetters({
				whoami: 'user/whoami'
			}),
			headers()
			{
				let headers = [...this.staticHeaders];

				if (this.category !== null)
				{
					if (this.selectedCategory)
					{
						this.selectedSpecificFields.forEach(f =>
							headers.push({
								text: f.name,
								value: f.sc_name
							}));
					}
				}

				headers.push({
					text: '',
					value: 'action',
					sortable: false,
					filterable: false,
					width: 60,
					align: 'center'
				});

				return headers;
			},
			selectedCategory()
			{
				return this.categories.find(item => item.id === this.category);
			},
			selectedSpecificFields()
			{
				const c = this.selectedCategory;

				return c ? c.specific_fields : [];
			}
		},
		methods: {
			async fetch()
			{
				if (this.category === null)
				{
					return;
				}

				this.loading = true;

				let url = new URLSearchParams();

				url.append('category_id', this.category);

				if (this.search)
				{
					url.append('q', this.search);
				}

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
						url: '/api/v1/resources/?' + url.toString()
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
				this.deleteDialog.id = item.id;

				this.deleteDialog.show = true;
			},
			showAction(item)
			{
				if (this.whoami.userType === 'admin')
				{
					return true;
				}

				return item.user.id === this.whoami.id;
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
