<template>
	<div>
		<h2 class="container bordered-heading">
			{{ locationName }}
		</h2>
		<p class="grey" v-if="locationDescription">
			{{ locationDescription }}
		</p>
		<div v-if="stops.length > 0">
			<div
				v-for="stop in stops"
				:key="stop.id"
				class="stop-summary-wrapper"
			>
				<stop-summary-card
					:stop="stop"
					@update-votes="updateUserVote"
				/>
			</div>
		</div>
		<div v-else>
			<background-label
				text="Can't find any stops for this location"
				color-preset="light-grey"
			/>
		</div>
	</div>
</template>

<script lang="ts">
import Vue from 'vue';
import { Component, Prop, Watch } from 'vue-property-decorator';
import StopSummaryCard from '@/components/stops/StopSummaryCard.vue';
import BackgroundLabel from '@/components/global/BackgroundLabel.vue';
import { ISearchQuery, ISearchResult } from '@/types/Explore';
import { Get } from 'vuex-pathify';
import { IStop } from '../../types/Stop';

@Component({
	name: 'ExploreResultScreen',
	components: {
		StopSummaryCard,
		BackgroundLabel
	}
})
export default class ExploreResultScreen extends Vue {

	/* Computed */
	@Get('explore/stops')
	stops: IStop[];

	/* Data */
	searchQuery: ISearchQuery = null;
	locationName: string = null;
	locationDescription: string = null;

	/* Methods */
	async search() {
		await this.$store.dispatch('explore/searchStops', this.searchQuery);
	}

	updateUserVote(votes) {
		console.log(votes);
	}

	async init() {
		this.locationName = this.$route.query.loc as string;
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