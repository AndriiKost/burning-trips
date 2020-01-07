<template>
  <div class="menu-wrapper">
	<el-menu 
	  :default-active="activeIndex" 
	  class="nav-menu" 
	  mode="horizontal" 
	  @select="handleSelect"
	  router
	>
	  <el-menu-item index="stops">Stops</el-menu-item>
	  <el-menu-item index="routes">Routes</el-menu-item>
	  <el-menu-item index="profile">Profile</el-menu-item>
	</el-menu>
  </div>
</template>

<script lang="ts">
import Vue from 'vue';
import { Component, Prop, Watch } from 'vue-property-decorator';

@Component({
   name: 'NavigationMenu'
})
export default class NavigationMenu extends Vue {

   	/* Props */
   	isCollapse: boolean = true;
	activeIndex= '1';
	activeIndex2= '1';

   	/* Computed */

   	/* Data */

   	/* Methods */
  	handleOpen(key, keyPath) {
		console.log(key, keyPath);
  	}

	handleClose(key, keyPath) {
		console.log(key, keyPath);
  	}

  	handleSelect(key, keyPath) {
		console.log(key, keyPath);
  	}

   	/* Lifecycle Hooks */

}
</script>

<style lang='scss'>
.menu-wrapper {
 position:fixed;
 bottom:0;
 left:0;
 right:0;
  .el-menu--horizontal > .el-menu-item.is-active {
	border-bottom: none;
	border-top: 2px solid #3a1f5d;
  }
}
</style>