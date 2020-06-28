<template>
	<div id="app">
		<section id="header">
			<div class="flex flex-row space-between">
				<router-link to="/">
					<div class="flex flex-row">
						<img src="https://burningbucket.s3.us-east-2.amazonaws.com/logo.png" alt="burning trips logo" />
						<span>Burning Trips</span>
					</div>
				</router-link>
				<!-- <top-navigation /> -->
			</div>
		</section>
		<section id="content">
			<router-view />
		</section>
		<!-- <navigation-menu /> -->
		<bottom-circle-menu v-if="isMobile" />
		<app-footer />
	</div>
</template>

<script lang="ts">
import Vue from 'vue';
import { Component, Prop, Watch } from 'vue-property-decorator';
import NavigationMenu from './navigation/NavigationMenu.vue';
import AppFooter from './AppFooter.vue';
import TopNavigation from './navigation/TopNavigation.vue';
import BottomCircleMenu from './navigation/BottomCircleMenu.vue';

@Component({
	name: 'AppContent',
	components: {
		NavigationMenu,
		BottomCircleMenu,
		TopNavigation,
		AppFooter
	}
})
export default class AppContent extends Vue {

	/* Props */

	/* Computed */
	get isMobile(): boolean {
		return true;
	}

	/* Data */

	/* Lifecycle Hooks */

}
</script>

<style lang='scss'>
html {
	position: relative;
	min-height: 100%;
}
body {
	height: 100%;
	padding: 0;
	 margin: 0 0 3rem; /* bottom = footer height */
}

#header {
	-webkit-box-shadow: 0px 0px 9px 3px rgba(41,41,41,.25);
	-moz-box-shadow: 0px 0px 9px 3px rgba(41,41,41,.25);
	box-shadow: 0px 0px 9px 3px rgba(41,41,41,.25);
	padding: .5rem 1rem;

	img {
		width: 60px;
		height: 60px;
	}
	span {
		font-size: 2rem;
		margin-left: .7rem;
		font-weight: 500;
		color: darken($red, 50%);
	}
	a {
		text-decoration: none;
	}
}
#content {
	margin: 1.5rem .75rem 3.5rem .75rem;
}
</style>