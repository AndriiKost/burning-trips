<template>
   <div class="stop-screen-modal-wrapper">
      <create-content-from-stops v-if="isCreateContentFromStops" />
   </div>
</template>

<script lang="ts">
import Vue from 'vue';
import { Component, Prop, Watch } from 'vue-property-decorator';
import { ModalType } from '@/types/components/Modal';
import CreateContentFromStops from './CreateContentFromStops.vue';
import { Sync } from 'vuex-pathify';

@Component({
   name: 'ModalWrapper',
   components: {
      CreateContentFromStops
   }
})
export default class ModalWrapper extends Vue {

   /* Props */

   /* Computed */
   @Sync('stop/activeModal')
   activeModal: ModalType;

   get isCreateContentFromStops(): boolean {
      return this.activeModal === ModalType.CreateContentFromStops;
   }

   /* Data */

   /* Methods */
   closeModal() {
      this.activeModal = null;
   }

   /* Lifecycle Hooks */

}
</script>

<style lang='scss'>

</style>