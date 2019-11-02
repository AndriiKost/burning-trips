<template>
	<div class="stop-summary">
		<el-card :body-style="{ padding: '0px' }" shadow="always">
			<a class="image-wrapper no-dec" href="#">
				<img :src="stop.imageUrl" class="image" />
			</a>
			<div class="content-wrapper relative">
				<a class="content no-dec" href="#">
					<h2>{{ stop.name }}</h2>
					<p>{{ stop.description }}</p>
				</a>
				<div class="bottom clearfix flex-row space-between">
					<vote-section 
						:total-votes="totalVotes" 
						:trending="stop.trending" 
						@update-votes="updateVotes"
					/>
					<el-button type="text" class="button-primary">
						Read More
					</el-button>
				</div>
			</div>
		</el-card>
	</div>
</template>

<script lang="ts">
import Vue from 'vue';
import { Component, Prop, Watch } from 'vue-property-decorator';
import { Get, Sync } from 'vuex-pathify';
import { IStop } from '../../types/Stop';
import { IUser } from '../../types/User';
import { IVote } from '../../types/Vote';
import VoteSection from '@/components/global/VoteSection.vue';

@Component({
	name: 'StopSummary',
	components: {
		VoteSection
	}
})
export default class StopSummaryCard extends Vue {

	/* Props */
	@Prop({ type: Object as () => IStop, required: true })
	readonly stop: IStop;

	/* Computed */
	@Get('user')
	readonly user: IUser;

	@Sync('stop/stops')
	stops: IStop[];

	get totalVotes() {
		return this.stop.votes.reduce((acc, cur) => acc += cur.userVotes, 0);
	}

	updateVotes(userVotes: number) {
		const userVote: IVote = { userID: this.user.uid, userVotes };
		// update votes
	}

	/* Lifecycle Hooks */
	
}
</script>

<style lang='scss'>
@import 'src/scss/variables.scss';
.stop-summary {
	max-width: 450px;
	.image-wrapper {
		.image {
			width: 100%;
			max-width: 100%;
			height: auto;
			object-fit: cover;
		}
	}

	.content {
		color: $blue;
		p {
			opacity: .8;
		}
	}

	.icon-brng {
		&.top-left {
			position: absolute;
			top: 0;
			left: 0;
		}
	}

	.button-primary {
		color: $blue;
	}

	.bottom {
		margin-top: 2rem;
	}

	.content-wrapper {
		padding: 1.4rem;
	}
}
</style>