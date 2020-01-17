<template>
   <div class="story-summary-wrapper">
       <div>
           <img :src="story.imageUrl" :alt="`${story.title} image`">
       </div>
       <div class="content">
           <h4>{{ story.title }}</h4>
           <!-- <p>{{ story.content }}</p> -->
           <div v-html="story.content"></div>
       </div>
        <vote-section  
            :total-votes="totalVotes" 
            :trending="story.trending"
            :cur-user-votes="curUserVoteCount"
            @save-votes="saveVotes"
        />
   </div>
</template>

<script lang="ts">
import Vue from 'vue';
import { Component, Prop, Watch } from 'vue-property-decorator';
import { IStory } from '@/types/Story';
import VoteSection from '@/components/global/VoteSection.vue';
import { IStoryVote } from '../../types/Vote';
import { Get } from 'vuex-pathify';
import { IUser } from '../../types/User';

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

   /* Data */

   /* Methods */

   async saveVotes() {
    //    TODO
   }

   /* Lifecycle Hooks */

}
</script>

<style lang='scss' scoped>

</style>