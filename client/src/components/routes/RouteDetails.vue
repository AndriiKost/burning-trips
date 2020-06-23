<template>
    <div class="container-sm">
        <a class="image-wrapper no-dec" :href="`/routes/${route.id}`">
            <div class="flex flex-row flex-wrap">
                <img 
                    v-for="(imgUrl, index) in imageList" 
                    :key="index" 
                    :src="imgUrl" 
                    style="width: 100%; height: 400px; object-fit: cover;"
                />
            </div>
        </a>
        <div class="content-wrapper relative">
            <a class="content no-dec" href="#">
                <h2>{{ route.name }}</h2>
                <p>{{ route.content }}</p>
                <div class="stop-count-section" v-for="stop in route.stops" :key="stop.id">
                    <i class="el-icon el-icon-sm el-icon-place" />
                    <span>{{ stop.name }}</span>
                    <span class="sm" style="margin-left: 1rem;color: grey;font-size: .85rem;">({{ stop.address }})</span>
                </div>
            </a>
            <div class="bottom clearfix flex flex-row space-between">
                <vote-section 
                    :total-votes="totalVotes" 
                    :trending="route.trending"
                    :cur-user-votes="curUserVoteCount"
                    @save-votes="saveVotes"
                />
            </div>
        </div>
    </div>
</template>

<script lang="ts">
import Vue from 'vue';
import { Component, Prop, Watch } from 'vue-property-decorator';
import { IRoute } from '../../types/Route';
import { Get } from 'vuex-pathify';
import { IUser } from '@/types/User';
import { IRouteVote } from '@/types/Vote';

@Component({
   name: 'RouteDetails'
})
export default class RouteDetails extends Vue {

   /* Props */
   @Prop({ type: Object as () => IRoute, required: true })
   readonly route: IRoute;

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

   /* Data */

   /* Methods */

   /* Lifecycle Hooks */

}
</script>

<style lang='scss' scoped>

</style>