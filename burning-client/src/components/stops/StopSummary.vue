<template>
	<div class="stop-summary">
		<el-card :body-style="{ padding: '0px' }" shadow="always">
			<a class="image-wrapper no-dec" href="#">
				<img :src="stop.imageUrl" class="image" />
			</a>
			<div class="content-wrapper">
				<a class="content no-dec" href="#">
					<h2>{{ stop.name }}</h2><div class="icon-brng inline-block"></div>
					<p>{{ stop.description }}</p>
				</a>
				<div class="bottom clearfix flex-row space-between">
					<div class="vote-section flex-row space-between">
						<div class="icon-wrapper relative">
							<i class="icon icon-btn el-icon-camera-solid" @click="handleVote" />
							<transition name="el-fade-in">
								<span v-show="showUserVotes" class="user-votes bg">
									+{{ userVotes }}
								</span>
							</transition>
						</div>
						<span class="vote-amount bg">{{ totalVotes }}</span>
					</div>
					<el-button type="text" class="button-primary">Read More</el-button>
				</div>
			</div>
		</el-card>
	</div>
</template>

<script lang="ts">
import Vue from 'vue';
import { Component, Prop, Watch } from 'vue-property-decorator';
import { Get, Sync } from 'vuex-pathify';
import { Stop, Vote } from '@/types/stops.type';
import { User } from '@/types/User.type';

let timer: any;

@Component({
	name: 'StopSummary'
})
export default class  extends Vue {

	/* Props */
	@Prop({ type: Stop, required: true })
	readonly stop: Stop;

	/* Computed */
	@Get('user')
	readonly user: User;

	@Get('stop/voteLimit')
	readonly voteLimit: number;

	@Sync('stop/stops')
	stops: Array<Stop>

	get totalVotes() {
		return this.stop.votes.reduce((acc, cur) => acc += cur.votes, 0);
	}

	/* Data */
	userVotes: number = 0;
	showUserVotes: boolean = false;


	/* Methods */
	handleVote(): void {
		clearTimeout(timer);
		this.showUserVotes = true;
		if (this.userVotes < this.voteLimit) this.userVotes++;
		timer = setTimeout(() => {
			this.showUserVotes = false;
			this.updateVotes();
		}, 1500);
	}

	updateVotes() {
		const userVote: Vote = {
			uid: this.user.uid,
			votes: this.userVotes
		};
	}

	/* Lifecycle Hooks */
	
}
</script>

<style lang='scss'>
@import 'src/sass/variables.scss';
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

	.vote-section {
		.icon-btn {
			color: darken($red, 10%);
			font-size: 1.75rem;
			margin-right: .25rem;
			cursor: pointer;
		}
		.vote-amount {
			color: $blue;
		}
		.user-votes {
			position: absolute;
			top: -2rem;
			left: 0;
			color: $red;
			padding: .25rem;
			border-radius: 50%;
			font-size: .85rem;
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