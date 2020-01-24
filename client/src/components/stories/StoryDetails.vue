<template>
    <div v-if="story" class="story-content-wrapper">
        <div class="story-content">
            <h1>{{ story.title }}</h1>
            <div v-html="storyContent"></div>
            <!-- <div>
                {{ story.content }}
            </div> -->
        </div>
        <!-- <vote-section  
            :total-votes="totalVotes" 
            :trending="story.trending"
            :cur-user-votes="curUserVoteCount"
            @save-votes="saveVotes"
        /> -->
    </div>
</template>

<script lang="ts">
import Vue from 'vue';
import { Component, Prop, Watch } from 'vue-property-decorator';
import { IStory } from '../../types/Story';

@Component({
   name: 'StoryDetails'
})
export default class StoryDetails extends Vue {

   /* Props */

   /* Computed */
   get storyContent() {
       const contentHtml = this.story.content.trim();
       console.log(contentHtml);
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

   /* Lifecycle Hooks */
   created() {
       this.init();
   }

}
</script>

<style lang='scss'>
.story-content-wrapper {
    margin: 3rem 1.5rem;
    .story-content {
        h1 {
            font-size: 2rem;
            margin: 0 0 2rem 0;
        }
        h2 {
            margin: 2rem 0 .5rem 0;
            color: rgba(0, 0, 0, 0.8);
            font-size: 1.8rem;
        }
        p {
            color: rgba(0, 0, 0, 0.6);
            font: 1.6rem;
        }
        img {
            width: 100%;
            height: auto;
            object-fit: cover;
        }
    }
}
</style>