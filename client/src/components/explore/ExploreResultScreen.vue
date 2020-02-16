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

	async init() {
		const lng = this.$route.query.lng as string;
		const lat = this.$route.query.lat as string;
		if (!lat || !lng) return;
		const latitude = parseFloat(lat)
		const longtitude = parseFloat(lng);
		this.searchQuery = { longtitude, latitude }
		await this.search();
	}

	/* Lifecycle Hooks */
	async created() {
		await this.init();
	}

}
</script>

<style lang='scss' scoped>

</style>