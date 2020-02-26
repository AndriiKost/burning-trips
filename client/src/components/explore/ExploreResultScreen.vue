<template>
	<div>
		<div 
			v-for="stop in stops"
			:key="stop.id"
			class="stop-summary-wrapper"
		>
			<stop-summary-card :stop="stop" @update-votes="updateUserVote" />
		</div>
	</div>
</template>

<script lang="ts">
import Vue from 'vue';
import { Component, Prop, Watch } from 'vue-property-decorator';
import StopSummaryCard from '@/components/stops/StopSummaryCard.vue';
import { ISearchQuery, ISearchResult } from '@/types/Explore';
import { Get } from 'vuex-pathify';
import { IStop } from '../../types/Stop';

@Component({
	name: 'ExploreResultScreen',
	components: {
		StopSummaryCard
	}
})
export default class ExploreResultScreen extends Vue {

	/* Props */

	/* Computed */
	@Get('explore/stops')
	stops: IStop[];

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
		const longitude = parseFloat(lng);
		this.searchQuery = { longitude, latitude }
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