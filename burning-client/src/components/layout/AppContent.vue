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
  </div>
</template>

<script lang="ts">
import Vue from 'vue';
import { Component, Prop, Watch } from 'vue-property-decorator';
import NavigationMenu from './navigation/NavigationMenu.vue';

@Component({
   name: 'AppContent',
   components: {
      NavigationMenu
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
body {
  height: 100%;
  padding: 0;
  margin: 0;
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