<template>
    <el-card :body-style="{ padding: '0px' }" shadow="always">
        <a 
            class="image-wrapper no-dec block ss-image-wrapper" 
            :href="`/story/${story.id}`" 
            @click.prevent="goToStoryDetails" 
        >
            <img 
                class="ss-image"
                :src="story.imageUrl" 
                :alt="`${story.title} image`"
            />
        </a>
        <div class="ss-card">
            <div class="content-wrapper relative">
                <a class="image-wrapper no-dec" :href="`/story/${story.id}`" @click.prevent="goToStoryDetails">
                    <h2>{{ story.title }}</h2>
                    <p>{{ briefContent }}</p>
                    <div class="stop-count-section" v-if="story.stops && story.stops.length">
                        <i class="el-icon el-icon-sm el-icon-place" />
                        <span v-if="story.stops">{{ story.stops.length }} Featured Stops</span>
                    </div>
                </a>
                <div class="bottom clearfix flex flex-row space-between" style="margin-top: 1rem;">
                    <vote-section  
                        :total-votes="totalVotes" 
                        :trending="story.trending"
                        :cur-user-votes="curUserVoteCount"
                        @save-votes="saveVotes"
                        icon="el-icon-reading"
                    />
                    <el-button type="text" class="button-primary" @click.prevent="goToStoryDetails">
                        Read More
                    </el-button>
                </div>
            </div>
        </div>
    </el-card>
</template>

<script lang="ts">
import Vue from 'vue';
import { Component, Prop, Watch } from 'vue-property-decorator';
import { IStory } from '@/types/Story';
import VoteSection from '@/components/global/VoteSection.vue';
import { IStoryVote } from '../../types/Vote';
import { Get } from 'vuex-pathify';
import { IUser } from '../../types/User';
import { stripHtml } from '@/utils/functions';

@Component({
   name: 'StorySummaryCard',
   components: {
       VoteSection
   }
})
export default class StorySummaryCard extends Vue {

   /* Props */
   @Prop({ type: Object as () => IStory, required: true })
   readonly story: IStory;

   @Prop({ type: Number, default: 150 })
   readonly briefContentLimit: number;

   /* Computed */
	@Get('auth/loggedInUser')
    readonly user: IUser;
    
   	get totalVotes(): Number {
		if (!this.story.votes || this.story.votes.length < 1) return 0;
		return this.story.votes.reduce((acc, cur) => acc += cur.count, 0);
    }
    
    get curUserVote(): IStoryVote {
        if (!this.user) return null;
		return this.story.votes.find(v => v.userId === this.user.id) as IStoryVote;
    }
    
    get curUserVoteCount(): Number {
		return this.curUserVote ? this.curUserVote.count : 0;
    }
    
    get briefContent(): string {
        const contentString: string = stripHtml(this.story.content);
        const briefContent = contentString.substring(0, this.briefContentLimit);
        return `${briefContent}...`;
    }

   /* Data */

   /* Methods */
   async goToStoryDetails() {
       await this.$router.push(`/story/details/${this.story.id}`);
   }

	/* Methods */
	async saveVotes(count: number) {
        if (!this.user) return this.$router.push('/login');
        const id = this.curUserVote ? this.curUserVote.id : 0;
		let storyVote: IStoryVote = {
            userId: this.user.id,
            count,
            id,
            storyId: this.story.id
        };
		const result = await this.$store.dispatch('story/updateStoryVote', storyVote);
		return this.$emit('update-votes', result);
	}

   /* Lifecycle Hooks */

}
</script>

<style lang='scss'>
.ss-image-wrapper {
    max-height: 200px;
    width: auto;

    .ss-image {
        width: 100%;
        height: auto;
        max-height: 100%;
        object-fit: cover;
    }
}
.ss-card {
    padding: 1rem;

}
</style>