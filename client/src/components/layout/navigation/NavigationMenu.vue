<template>
  <div class="menu-wrapper">
	<el-menu 
	  :default-active="activeIndex" 
	  class="nav-menu flex flex-row space-between" 
	  mode="horizontal" 
	  @select="handleSelect"
	  router
	>
		<el-menu-item 
			v-for="menuItem in navMenuItems" 
			:key="menuItem.label" 
			:index="menuItem.link"
		>
			<div class="flex flex-center flex-column">
				<i :class="`el-icon el-icon-sm el-icon-${menuItem.icon}`"></i>
				<span class="nav-label">{{ menuItem.label }}</span>
			</div>
		</el-menu-item>
	</el-menu>
  </div>
</template>

<script lang="ts">
import Vue from 'vue';
import { Component, Prop, Watch } from 'vue-property-decorator';
import { INavigationMenuItem } from '@/types/UiTypes';
import { IUser } from '@/types/User';
import { Get } from 'vuex-pathify';

@Component({
   name: 'NavigationMenu'
})
export default class NavigationMenu extends Vue {

   	/* Props */

	/* Computed */
	@Get('auth/loggedInUser')
	loggedInUser: IUser;

	/* Data */
	activeIndex: string = null;
	navMenuItems: INavigationMenuItem[] = [
		{
			label: 'Stops',
			icon: 'place',
			link: 'stops'
		},
		{
			label: 'Routes',
			icon: 'guide',
			link: 'routes'
		},
		{
			label: 'Stories',
			icon: 'reading',
			link: 'stories'
		},
	];

   	/* Methods */
  	handleOpen(key, keyPath) {
		// console.log(key, keyPath);
  	}

	handleClose(key, keyPath) {
		// console.log(key, keyPath);
  	}

  	handleSelect(key, keyPath) {
		// console.log(key, keyPath);
	  }
	  
	preselectMenuIndex() {
		const routePath = this.$route.path;
		const activeIndex = routePath.replace('/', ' ').trim().split(' ')[0];
		this.activeIndex = activeIndex;
	}

	/* Lifecycle Hooks */
	beforeMount() {
		this.preselectMenuIndex();
		if (this.loggedInUser) {
			this.navMenuItems.push({
				icon: 'user',
				link: 'account',
				label: 'Account'
			})
		} else {
			this.navMenuItems.push({
				icon: 'user',
				link: 'signin',
				label: 'Sign In'
			})
		}
	}

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