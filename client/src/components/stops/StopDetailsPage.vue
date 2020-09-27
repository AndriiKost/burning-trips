<template>
	<div v-if="!activeStop">
		Loading ...
	</div>
	<div v-else class="stop-details-screen">
		<div class="grid-content bg-purple stop-detailes-wrapper">
			<div class="sd-head">
				<img 
					:src="activeStop.imageFullUrl" 
					:alt="`Image of ${activeStop.name}`" 
				/>
			</div>
			<h1>{{ activeStop.name }}</h1>
			<div class="sd-content">
				{{ activeStop.content }}
			</div>

			<div class="route-carousel" v-if="featuredLandmarks">
				<h2>Burning Landmarks in the area</h2>
				<landmark-carousel :stops="featuredLandmarks" />
			</div>

			<div class="landmark-carousel" v-if="upsellLandmarks">
				<h2>Similar Landmarks to visit</h2>
				<landmark-carousel :stops="upsellLandmarks" />
			</div>

			<div class="sd-footer">
				<vote-section 
					:total-votes="totalVotes" 
					:trending="activeStop.trending"
					:cur-user-votes="curUserVoteCount"
					@save-votes="save"
				/>
			</div>
		</div>
	</div>
</template>

<script lang="ts">
import Vue from 'vue';
import { Component, Prop, Watch } from 'vue-property-decorator';
import { IStop } from '../../types/Stop';
import StopSummaryCard from './StopSummaryCard.vue';
import VoteSection from '@/components/global/VoteSection.vue';
import { Get } from 'vuex-pathify';
import { IUser } from '../../types/User';
import { IStopVote } from '../../types/Vote';
import { imageExists } from '@/utils/functions';
import { constructImageUrlForWikipedia } from '@/helpers/ImageHelper';

@Component({
   name: 'StopDetailsPage',
   components: {
	   StopSummaryCard,
	   VoteSection
   }
})
export default class StopDetailsPage extends Vue {

	/* Computed */
	@Get('auth/loggedInUser')
	readonly user: IUser;

	@Get('stop/featuredLandmarks')
	readonly featuredLandmarks: IStop[];

	@Get('stop/upsellLandmarks')
	readonly upsellLandmarks: IStop[];

	/* Data */
	activeStop: IStop = null;

	/* Methods */
	async init() {
		const { id } = this.$route.params;
		const stop = await this.$store.dispatch('stop/getStop', id);
		if (!stop) {
			// add error to the notif stack
			this.$router.push('/stops');
		}
		this.activeStop = stop;
		await this.$store.dispatch('stop/getStopViewModel', stop.id);
	}

	// async initImage() {
	// 	this.stopImageUrl = this.activeStop.imageUrl;
	// 	const biggerImageUrl = constructImageUrlForWikipedia(this.stopImageUrl);
	// 	const hasBiggerImage = await imageExists(biggerImageUrl);
	// 	if (hasBiggerImage) {
	// 		this.stopImageUrl = biggerImageUrl;
	// 	}
	// }

	get totalVotes(): Number {
		if (this.activeStop.votes && this.activeStop.votes.length > 0) {
			return this.activeStop.votes.reduce((acc, cur) => acc += cur.count, 0);
		}
		return 0;
	}

	get curUserVote(): IStopVote {
		if (!this.user) return null;
		if (this.activeStop.votes && this.activeStop.votes.length > 0) {
			return this.activeStop.votes.find(v => v.userId === this.user.id) as IStopVote;
		}
		return null;
	}

	get curUserVoteCount(): Number {
		return this.curUserVote ? this.curUserVote.count : 0;
	}

	async save(count: number) {
		if (!this.user) {
			localStorage.setItem('redirectTo', window.location.pathname + window.location.search);
			this.$router.push('/login');
		}
		const userId = this.user ? this.user.id : 0;
		const userVote: IStopVote = { 
			userId, 
			count, 
			id: 0, 
			stopId: this.activeStop.id 
		};
		if (this.curUserVote) {
			// update existing votes
			userVote.id = this.curUserVote.id;
		}
		const result = await this.$store.dispatch('stop/updateStopVote', userVote);
		this.updateUserVote(userVote);
	}

	updateUserVote(userVote: IStopVote) {
		const hasUserVotes = this.activeStop.votes.some(v => v.id === userVote.id);
		if (hasUserVotes) {
			this.activeStop.votes.forEach(v => v.id === userVote.id ? v.count = userVote.count : null)
		} else {
			this.activeStop.votes.push(userVote);
		}
	}

	/* Lifecycle Hooks */

	created() {
		this.init();
	}

}
</script>

<style lang='scss' scoped>
.stop-details-screen {

	.sd-head {
		margin-bottom: 1rem;
		
		img {
			width: 100%;
		}
	}

	h1 {
		margin-bottom: 1rem;
	}

	.stop-detailes-wrapper {
		margin: 5rem;
	}

	.sd-footer {
		margin: 2rem 0;
	}
}
</style>