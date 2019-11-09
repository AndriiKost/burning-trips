<template>
	<el-dialog
		:title="title"
		:visible.sync="visible"
		width="30%"
		:before-close="handleClose"
	>
	<slot></slot>
	<span slot="footer" class="dialog-footer">
		<el-button 
			@click="handleClose"
		>
			Cancel
		</el-button>
		<el-button 
			type="primary" 
			@click="$emit('submit')"
		>
			Submit
		</el-button>
	</span>
	</el-dialog>
</template>

<script lang="ts">
import Vue from 'vue';
import { Component, Prop, Watch } from 'vue-property-decorator';

@Component({
   name: 'Modal'
})
export default class Modal extends Vue {

   /* Props */
   @Prop({ type: Boolean, default: false })
   readonly visible: Boolean;

   @Prop({ type: String, default: 'Modal Title' })
   readonly title: String;

   @Prop({ type: Function })
   readonly beforeClose: Function;

   /* Computed */

   /* Data */

   /* Methods */
   handleClose() {
	   if(this.beforeClose) this.beforeClose();
	   this.$nextTick(() => this.$emit('cancel'));
   }

   /* Lifecycle Hooks */

}
</script>

<style lang='scss' scoped>

</style>