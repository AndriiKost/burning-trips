<template>
	<div class="route-summary">
		<el-card :body-style="{ padding: '0px' }" shadow="always">
			<a class="image-wrapper no-dec" :href="`/routes/${route.id}`">
				<div class="flex flex-row flex-wrap">
					<img 
						v-for="(imgUrl, index) in imageList" 
						:key="index" 
						:src="imgUrl" 
						style="width: auto; height: 120px;"
					/>
				</div>
			</a>
			<div class="content-wrapper relative">
				<a class="content no-dec" href="#">
					<h2>{{ route.name }}</h2>
					<p>{{ route.content }}</p>
					<div class="stop-count-section">
						<i class="el-icon el-icon-sm el-icon-place" />
						<span v-if="route.stops">{{ route.stops.length }} Stops</span>
					</div>
				</a>
				<div class="bottom clearfix flex flex-row space-between">
					<vote-section 
						:total-votes="totalVotes" 
						:trending="route.trending"
						:cur-user-votes="curUserVoteCount"
						@save-votes="saveVotes"
					/>
					<el-button type="text" class="button-primary" @click="viewDetails">
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
import { IRoute } from '../../types/Route';
import { IRouteVote } from '../../types/Vote';
import { IUser } from '../../types/User';
import { Get } from 'vuex-pathify';
import VoteSection from '@/components/global/VoteSection.vue';

@Component({
   name: 'RouteSummaryCard',
   components: {
       VoteSection,
   }
})
export default class RouteSummaryCard extends Vue {

   /* Props */
   @Prop({ type: Object as () => IRoute, required: true })
   route: IRoute;

   	/* Computed */
	@Get('auth/loggedInUser')
	readonly user: IUser;

	get totalVotes(): number {
		if (!this.route.votes || this.route.votes.length < 1) return 0;
		return this.route.votes.reduce((acc, cur) => acc += cur.count, 0);
	}

	get curUserVote(): IRouteVote {
		if (!this.user) return null;
		return this.route.votes.find(v => v.userId === this.user.id) as IRouteVote;
	}

	get curUserVoteCount(): number {
		return this.curUserVote ? this.curUserVote.count : 0;
	}

	get imageList(): String[] {
		if (!this.route.stops || this.route.stops.length < 1) return [];
		let imageList = this.route.stops.map(stop => stop.imageUrl);
		imageList.unshift(this.route.imageUrl);
		return imageList.filter(imgUrl => imgUrl !== '');
	}

	/* Methods */
	async saveVotes(count: number) {
		if (!this.user) {
			localStorage.setItem('redirectTo', window.location.pathname + window.location.search);
			return this.$router.push('/login');
		}
		const routeVote: IRouteVote = { userId: this.user.id, count, id: 0, routeId: this.route.id };
		if (this.curUserVote) {
			// update existing votes
			routeVote.id = this.curUserVote.id;
		}
		const result = await this.$store.dispatch('route/updateRouteVote', routeVote);
		return this.$emit('update-votes', result);
	}

	viewDetails() {
		this.$router.push(`routes/${this.route.id}`)
	}

}
</script>

<style lang='scss'>
@import 'src/scss/variables.scss';
.stop-count-section {
	margin: .5rem 0;
	.el-icon {
		margin-right: .3rem;
		color: $dark-grey;
	}
}

.route-summary {
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