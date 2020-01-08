<template>
	<div class="stop-summary">
		<el-card :body-style="{ padding: '0px' }" shadow="always">
			<a class="image-wrapper no-dec" :href="`/stops/${stop.id}`">
				<img :src="stop.imageUrl" class="image" />
			</a>
			<div class="content-wrapper relative">
				<a class="content no-dec" href="#">
					<h2>{{ stop.name }}</h2>
					<p>{{ stop.content }}</p>
				</a>
				<div class="bottom clearfix flex flex-row space-between">
					<vote-section 
						:total-votes="totalVotes" 
						:trending="stop.trending"
						:cur-user-votes="curUserVoteCount"
						@save-votes="save"
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
import { IStopVote } from '../../types/Vote';
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
	@Get('auth/loggedInUser')
	readonly user: IUser;

	get totalVotes(): Number {
		if (!this.stop.votes || this.stop.votes.length < 1) return 0;
		return this.stop.votes.reduce((acc, cur) => acc += cur.count, 0);
	}

	get curUserVote(): IStopVote {
		return this.stop.votes.find(v => v.userId === this.user.id) as IStopVote;
	}

	get curUserVoteCount(): Number {
		return this.curUserVote ? this.curUserVote.count : 0;
	}

	async save(count: number) {
		const stopVote: IStopVote = { userId: this.user.id, count, id: 0, stopId: this.stop.id };
		if (this.curUserVote) {
			// update existing votes
			stopVote.id = this.curUserVote.id;
		}
		const result = await this.$store.dispatch('stop/updateStopVote', stopVote);
		return this.$emit('update-votes', result);
	}

	/* Lifecycle Hooks */
	
}
</script>

<style lang='scss'>
@import 'src/scss/variables.scss';
.stop-summary {
	max-width: 450px;

	&.short-summary {
		.image-wrapper {
			.image {

			}
		}
	}

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
		padding: 1rem;
	}
}
</style>