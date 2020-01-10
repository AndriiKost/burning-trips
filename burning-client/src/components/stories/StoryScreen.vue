<template>
   <div>
        <h1>Story Screen</h1>
        <div v-for="story in stories" :key="story.id">
           {{ story.title }}
        </div>
   </div>
</template>

<script lang="ts">
import Vue from 'vue';
import { Component, Prop, Watch } from 'vue-property-decorator';
import { IStory } from '@/types/Story';
import { Sync, Get } from 'vuex-pathify';

@Component({
   name: 'StoryScreen'
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

   /* Lifecycle Hooks */
   async created() {
      await this.init();
   }

}
</script>

<style lang='scss' scoped>

</style>