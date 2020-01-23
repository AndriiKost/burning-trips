<template>
   <div class="editor-wrapper">
        <ckeditor :editor="editor" v-model="editorData" :config="editorConfig"></ckeditor>
   </div>
</template>

<script lang="ts">
import Vue from 'vue';
import { Component, Prop, Watch } from 'vue-property-decorator';
import ClassicEditor from '@ckeditor/ckeditor5-build-classic';

@Component({
   name: 'EditStoryScreen',
})
export default class EditStoryScreen extends Vue {

   /* Props */

   /* Computed */

   /* Data */
    editor = ClassicEditor;
    editorData:string = '<p>Content of the editor.</p>';
    editorConfig = {};

   /* Methods */

   /* Lifecycle Hooks */

}
</script>

<style lang='scss'>
.editor-wrapper {
    .ck-content {
        min-height: 25rem;
    }
}
</style>