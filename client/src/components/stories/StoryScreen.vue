<template>
   <div class="story-screen-wrapper">
      <div
         v-for="story in stories" 
         :key="story.id"
         class="ssc-wrapper"
      >
         <story-summary-card :story="story" @update-votes="updateUserVote" />
      </div>
   </div>
</template>

<script lang="ts">
import Vue from 'vue';
import { Component, Prop, Watch } from 'vue-property-decorator';
import { IStory } from '@/types/Story';
import { Sync, Get } from 'vuex-pathify';
import StorySummaryCard from './StorySummaryCard.vue';
import { IStoryVote } from '@/types/Vote';

@Component({
   name: 'StoryScreen',
   components: {
      StorySummaryCard
   }
})
export default class StoryScreen extends Vue {

   /* Props */

   /* Computed */
   @Get('story/stories')
   stories: IStory[];

   /* Data */

   /* Methods */
   async init() {
      await this.$store.dispatch('story/getAllStories');
   }

   updateUserVote(userVote: IStoryVote) {
      this.stories.forEach(stop => {
         stop.votes.some(v => v.id === userVote.id)
            ? stop.votes.forEach(v => v.id === userVote.id ? v.count = userVote.count : null)
            : stop.votes.push(userVote);
      });
   }

   /* Lifecycle Hooks */
   async created() {
      await this.init();
   }

}
</script>

<style lang='scss'>
.story-screen-wrapper {
   .ssc-wrapper {
      margin-top: 2rem;
   }
}
</style>