<template>
   <div :class="`add-btn-wrapper ${position}`">
	   <div class="icon-add"></div>
   </div>
</template>

<script lang="ts">
import Vue from 'vue';
import { Component, Prop, Watch } from 'vue-property-decorator';

@Component({
   name: 'AddButton'
})
export default class AddButton extends Vue {

   /* Props */
   @Prop({ type: String, default: 'bottom-right' })
   readonly position: String;

   /* Computed */

   /* Data */

   /* Methods */

   /* Lifecycle Hooks */

}
</script>

<style lang='scss' scoped>
.add-btn-wrapper {
	display: block;
	&.bottom-right {
		position: fixed;
		bottom: 1rem;
		right: 1rem;
	}
}
</style>