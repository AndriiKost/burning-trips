<template>
    <div v-if="story" class="story-content-wrapper container">
        <div class="story-content">
            <h1>{{ story.title }}</h1>
            <div v-html="storyContent"></div>
            <div class="spacer-md"></div>
            <vote-section  
                :total-votes="totalVotes" 
                :trending="story.trending"
                :cur-user-votes="curUserVoteCount"
                @save-votes="saveVotes"
                icon="el-icon-reading"
            />
        </div>
    </div>
</template>

<script lang="ts">
import Vue from 'vue';
import { Component, Prop, Watch } from 'vue-property-decorator';
import { IStory } from '@/types/Story';
import { IUser } from '@/types/User';
import { IStoryVote } from '@/types/Vote';
import { Get } from 'vuex-pathify';

@Component({
   name: 'StoryDetails'
})
export default class StoryDetails extends Vue {

   /* Props */

   /* Computed */
   @Get('auth/loggedInUser')
   user: IUser;

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

   get storyContent() {
       const contentHtml = this.story.content.trim();
       return contentHtml;
   }

   /* Data */
   story: IStory = null;

   /* Methods */
   async init() {
       const storyId = this.$route.params.storyId;
       const story: IStory = await this.$store.dispatch('story/getStory', storyId);
       this.story = story;
   }

	async saveVotes(count: number) {
        if (!this.user) this.$router.push('/login');
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
   created() {
       this.init();
   }

}
</script>

<style lang='scss'>
.story-content-wrapper {

    .story-content {
        @include mobile {
            margin: 3rem 1rem 3rem 1rem;
        }

        h1 {
            font-size: 2rem;
            margin: 0 0 2rem 0;
        }

        h2 {
            margin: 2rem 0 1rem 0;
            color: rgba(0, 0, 0, 0.8);
            font-size: 1.8rem;
        }

        p {
            color: lighten($dark-grey, 15%);
            font-size: 1.1rem;
            line-height: 1.8rem;
        }

        img {
            width: 100%;
            height: auto;
            object-fit: cover;
        }

        .image {
            width: 100%;
            margin: 2rem 0;
            padding: 0;
        }
    
    }
}
</style>