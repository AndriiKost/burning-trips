<template>
    <el-form-item :label="label">
        <div class="el-input">
            <input 
                type="text" 
                ref="autocomplete" 
                class="el-input__inner"
            >
        </div>
    </el-form-item>
</template>

<script lang="ts">
import Vue from 'vue';
import { Component, Prop, Watch } from 'vue-property-decorator';

@Component({
   name: 'AddressAutocomplete'
})
export default class AddressAutocomplete extends Vue {

    /* Props */
    @Prop({ type: String, default: 'Address' })
    label: string;
    /* Computed */

    /* Data */
   	autocomplete: any = null;

    /* Methods */
	initAutocomplete() {
		//@ts-ignore
		this.autocomplete = new google.maps.places.Autocomplete((this.$refs.autocomplete), {types: ['geocode']});
		this.autocomplete.addListener('place_changed', () => {
			let place = this.autocomplete.getPlace();
            this.$emit('location-click', place);
		});
	}

    /* Lifecycle Hooks */
   	mounted() {
		this.initAutocomplete();
	}
}
</script>

<style lang='scss'>

</style>