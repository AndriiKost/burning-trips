<template>
   <div class="editor-wrapper" v-if="activeStory">
		<el-input style="margin-bttom: 1rem;" placeholder="Please input" v-model="activeStory.title"></el-input>
		<ckeditor :editor="editor" v-model="editorData" :config="editorConfig"></ckeditor>
   </div>
</template>

<script lang="ts">
import Vue from 'vue';
import { Component, Prop, Watch } from 'vue-property-decorator';
import ClassicEditor from '@ckeditor/ckeditor5-build-classic';
import { IStory } from '@/types/Story';
import { IUser } from '@/types/User';
import { Get } from 'vuex-pathify';
import Story from '@/models/Story';

@Component({
   name: 'EditStoryScreen',
})
export default class EditStoryScreen extends Vue {

   /* Props */

   /* Computed */
   @Get('user')
   user: IUser;

   /* Data */
	editor = ClassicEditor;
	editorData:string = '<p>Content of the editor.</p>';
	editorConfig = {};
	activeStory: IStory = null;

   /* Methods */
   	async init() {
		const storyId: number = Number(this.$route.params.storyId);
		if (!storyId) {
			this.$router.push('/');
		}
		const result = await this.$store.dispatch('story/getStory', storyId);
		if (result) {
			this.activeStory = result;
		}
	}

   /* Lifecycle Hooks */
   created() {
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