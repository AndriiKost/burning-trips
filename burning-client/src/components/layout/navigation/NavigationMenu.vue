<template>
  <div class="menu-wrapper">
	<el-menu 
	  :default-active="activeIndex" 
	  class="nav-menu flex flex-row space-between" 
	  mode="horizontal" 
	  @select="handleSelect"
	  router
	>
		<el-menu-item index="stops">
			<div class="flex flex-center flex-column">
				<i class="el-icon el-icon-sm el-icon-place"></i>
				<span class="nav-label">Stops</span>
			</div>
		</el-menu-item>
		<el-menu-item index="routes">
			<div class="flex flex-center flex-column">
				<i class="el-icon el-icon-sm el-icon-guide"></i>
				<span class="nav-label">Routes</span>
			</div>
		</el-menu-item>
		<el-menu-item index="stories">
			<div class="flex flex-center flex-column">
				<i class="el-icon el-icon-sm el-icon-reading"></i>
				<span class="nav-label">Stories</span>
			</div>
		</el-menu-item>
		<el-menu-item index="account">
			<div class="flex flex-center flex-column">
				<i class="el-icon el-icon-sm el-icon-user"></i>
				<span class="nav-label">Account</span>
			</div>
		</el-menu-item>
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

   	/* Computed */

	/* Data */
	activeIndex = '1';

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
@import 'src/scss/variables.scss';
.menu-wrapper {
	position:fixed;
	bottom:0;
	left:0;
	right:0;
  	.el-menu--horizontal {

		.el-menu-item {
			height: 100%;
			padding-top: 7px;
			padding-bottom: 5px;
			&.is-active {
				border-bottom: none;
				border-top: 2px solid #3a1f5d;
			}
			.el-icon {
				margin-bottom: 3px;
				color: $blue;
			}
		}
  	}
  	.nav-label {
		line-height: 12px;
		font-size: 12px;
		color: $blue;
  	}
}
</style>