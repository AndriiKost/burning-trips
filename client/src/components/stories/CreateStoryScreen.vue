<template>
   <div class="editor-wrapper" v-if="activeStory">
       <h2>
           Create Your Story
       </h2>
		<el-input style="margin-bottom: 1rem;" placeholder="Story Headline" v-model="activeStory.title"></el-input>
        <el-input style="margin-bottom: 1rem;" placeholder="Main Image Url" v-model="activeStory.imageUrl"></el-input>
		<ckeditor :editor="editor" v-model="editorData" :config="editorConfig"></ckeditor>
        <el-row style="margin-top:2rem;">
            <el-button icon="el-icon-upload" type="primary" @click="createStory">
                Create
            </el-button>
        </el-row>
   </div>
</template>

<script lang="ts">
import Vue from 'vue';
import { Component, Prop, Watch } from 'vue-property-decorator';
import ClassicEditor from '@ckeditor/ckeditor5-build-classic';
import { IStory } from '@/src/types/Story';
import { IUser } from '@/src/types/User';
import { Get } from 'vuex-pathify';
import Story from '@/models/Story';
import { jsonCopy } from '../../utils/functions';

@Component({
   name: 'CreateStoryScreen',
})
export default class CreateStoryScreen extends Vue {

   /* Props */

   /* Computed */
   @Get('auth/loggedInUser')
   loggedInUser: IUser;

   /* Data */
	editor = ClassicEditor;
	editorData: string = '<p>The best vacation of my life started the day I ...</p>';
	editorConfig = {};
	activeStory: IStory = null;

   /* Methods */

   	init(): void {
		this.activeStory = new Story(this.loggedInUser.id);
    }
    
    async createStory(): Promise<void> {
        let story: IStory = jsonCopy(this.activeStory);
        story.content = this.editorData;
        const result = await this.$store.dispatch('story/createStory', story);
        if (result) {
            this.$router.push('/');
        }
    }

   /* Lifecycle Hooks */
   created(): void {
	  this.init();
   }

}
</script>

<style lang='scss'>
.editor-wrapper {
	.ck-content {
		min-height: 25rem;
	}
}
</style>