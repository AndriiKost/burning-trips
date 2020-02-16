<template>
	<div>

	</div>
</template>

<script lang="ts">
import Vue from 'vue';
import { Component, Prop, Watch } from 'vue-property-decorator';
import { ISearchQuery } from '@/types/ISearchQuery';

@Component({
	name: 'ExploreResultScreen'
})
export default class ExploreResultScreen extends Vue {

	/* Props */

	/* Computed */

	/* Data */
	searchQuery: ISearchQuery = null;

	/* Methods */
	async search() {
		await this.$store.dispatch('explore/searchStops', this.searchQuery);
	}

	/* Lifecycle Hooks */
	async created() {
		const { lng, lat } = this.$route.query;
		if (lng && lat) {
			this.searchQuery = {
				longtitude: lng,
				latitude: lat
			}
			await this.search();
		}
	}

}
</script>

<style lang='scss' scoped>

</style>