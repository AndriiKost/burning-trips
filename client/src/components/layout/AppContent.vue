<template>
  <div id="app">
    <section id="header">
      <div class="flex flex-row">
        <img src="https://burningbucket.s3.us-east-2.amazonaws.com/logo.png" alt="burning trips logo" />
        <span>Burning Trips</span>
      </div>
    </section>
    <router-view />
    <!-- <navigation-menu /> -->
    <app-footer />
  </div>
</template>

<script lang="ts">
import Vue from 'vue';
import { Component, Prop, Watch } from 'vue-property-decorator';
import NavigationMenu from './navigation/NavigationMenu.vue';
import AppFooter from './AppFooter.vue';

@Component({
   name: 'AppContent',
   components: {
      NavigationMenu,
      AppFooter
   }
})
export default class AppContent extends Vue {

   /* Props */

   /* Computed */

   /* Lifecycle Hooks */

}
</script>

<style lang='scss'>
@import 'src/scss/variables.scss';
@import 'src/scss/mixins.scss';
@import 'src/scss/global.scss';
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
  border-bottom: 2px solid $blue;
  img {
    width: 60px;
    height: 60px;
  }
  span {
    font-size: 2rem;
    margin-left: .5rem;
    font-weight: 500;
    color: $blue;
  }
}
</style>