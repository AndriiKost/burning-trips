<template>
  <div id="app">
    <navigation-menu />
    <div id="nav">
      <router-link to="/">Home</router-link>
      <!-- <router-link to="/about">About</router-link> -->
    </div>
    <router-view />
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

</style>