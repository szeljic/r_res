<template>
	<v-container>
		<v-row dense>
			<v-col>
				<v-toolbar dense elevation="2">
					<v-toolbar-title>Rezervacije</v-toolbar-title>
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
				<reservation-form-component
					:id="form.id"
					@success="success"
					@close="form.show = false"
				></reservation-form-component>
			</v-col>
		</v-row>

		<v-row dense v-if="!form.show">
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
					:footer-props="{
						itemsPerPageText: 'Redova po stranici',
						pageText: '{0}-{1} od {2}'
					}"
				>
					<template v-slot:item="{item}">
						<tr>
							<td class="text-center">{{ item.id }}</td>
							<td>{{ item.resource.name }}</td>
							<td>{{ item.resource.user.first_name + ' ' + item.resource.user.last_name }}</td>
							<td>{{ item.user.first_name + ' ' + item.user.last_name }}</td>
							<td>{{ $dateFormatL18n(new Date(item.from_date * 1000)) }}</td>
							<td>{{ $dateFormatL18n(new Date(item.to_date * 1000)) }}</td>
							<td>{{ $dateFormatL18n(new Date(item.created_at * 1000), true) }}</td>
							<td class="text-right">

								<v-chip color="green" v-if="item.status === 0" dark label>Neizvrsen</v-chip>
								<v-chip color="orange" v-if="item.status === 1" dark label>U toku</v-chip>
								<v-chip color="red" v-if="item.status === 2" dark label>Zavrsen</v-chip>

							</td>

							<td class="text-center">
								<table-menu-btn :disabled="!showAction(item)">
									<v-list dense>
										<v-list-item-group>
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
			:url="'/api/v1/reservations/' + deleteDialog.id"
			@success="fetch()"
		></delete-dialog>

	</v-container>
</template>

<script>
	import ReservationFormComponent from '@/views/reservations/ReservationFormComponent';
	import {mapGetters} from 'vuex';

	export default {
		name: 'ReservationsTable',
		components: {
			ReservationFormComponent
		},
		data()
		{
			return {
				headers: [{
					text: '#',
					value: 'id',
					width: 60,
					align: 'center',
					sortable: false
				}, {
					text: 'Resurs',
					value: 'resource',
					width: 180
				}, {
					text: 'Vlasnik resursa',
					value: 'resource.user'
				}, {
					text: 'Iznajmljuje',
					value: 'user'
				}, {
					text: 'Rezervisano od',
					value: 'from_date',
					width: 150
				}, {
					text: 'Rezervisano do',
					value: 'to_date',
					width: 150
				}, {
					text: 'Datum rezervisanja',
					value: 'created_at'
				}, {
					text: 'Status',
					value: 'status',
					width: 60,
					align: 'right',
					sortable: false,
					filterable: false
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
				search: null,
				loading: false,
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
		created()
		{
			this.fetch();
		},
		computed:{
			...mapGetters({
				whoami: 'user/whoami'
			})
		},
		methods: {
			async fetch()
			{
				if (this.category === null)
				{
					return;
				}

				this.loading = true;

				try
				{
					const {data} = await this.$http({
						url: '/api/v1/reservations' + (this.search ? '/?q=' + this.search : '')
					});

					this.total = data.total;

					this.items = data.items || [];

					this.items.map(item =>
					{
						item.status = 0;

						const now = new Date().valueOf();

						if (now < item.from_date * 1000)
						{
							item.status = 0;
						} else if (now > item.from_date * 1000 && now < item.to_date * 1000)
						{
							item.status = 1;
						} else if (now > item.to_date * 1000)
						{
							item.status = 2;
						}
					});
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
			}
		}
	};
</script>
