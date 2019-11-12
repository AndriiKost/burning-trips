<template>
	<div>
		<el-form 
			:model="newStop" 
			:rules="rules" 
			ref="new_stop_form" 
			label-width="120px" 
			class="new-stop-form"
		>
			<el-form-item label="Stop Name" prop="name">
				<el-input v-model="newStop.name"></el-input>
			</el-form-item>
			<el-form-item label="Address" prop="address.address">
				<el-input v-model="newStop.address"></el-input>
			</el-form-item>
			<el-form-item label="Image Url" prop="imageUrl">
				<el-input v-model="newStop.imageUrl"></el-input>
			</el-form-item>
			<el-form-item label="Description" prop="description">
				<el-input type="textarea" v-model="newStop.description"></el-input>
			</el-form-item>
			<el-form-item label="Preview My Stop">
				<el-switch :value="preview" @input="handlePreview"></el-switch>
			</el-form-item>
			<stop-summary-card
				v-if="preview && valid"
				:stop="newStop"
			/>
			<el-form-item>
				<el-button type="primary" @click="validate">Create</el-button>
				<el-button @click="resetForm">Reset</el-button>
			</el-form-item>
		</el-form>
	</div>
</template>

<script lang="ts">
import Vue from 'vue';
import { Component, Prop, Watch } from 'vue-property-decorator';
import { IStop } from '../../../types/Stop';
import Stop from '../../../models/Stop';
import StopSummaryCard from '../StopSummaryCard.vue';

@Component({
	name: 'CreateStopSection',
	components: {
		StopSummaryCard
	}
})
export default class CreateStopSection extends Vue {

	/* Props */

	/* Computed */
	newStop: IStop = new Stop();
	preview: boolean = false;
	previewImage: boolean = false;
	valid: boolean = false;

	rules = {
		name: [
			{ required: true, message: 'Please enter a Stop Name', trigger: 'blur' },
			{ min: 3, max: 60, message: 'Length should be 3 to 60 characters', trigger: 'blur' }
		],
		address: [
			{ required: true, message: 'Please enter a Stop Name', trigger: 'blur' },
			{ min: 3, message: 'Please enter at least 3 characters', trigger: 'blur' }
		],
		imageUrl: [
			{ required: true, message: 'Please enter a Stop Name', trigger: 'blur' },
			{ min: 3, message: 'Please enter at least 3 characters', trigger: 'blur' }
		],
		description: [
			{ required: true, message: 'Please enter a Stop Name', trigger: 'blur' },
			{ min: 3, message: 'Please enter at least 3 characters', trigger: 'blur' }
		],
	};

	/* Methods */

	validate() {
		// @ts-ignore
        this.$refs.new_stop_form.validate(valid => this.valid = valid);
	}

	handlePreview(val: boolean) {
		this.validate();
		if (val && this.valid) {
			this.preview = true;
		} else if (!val) {
			this.preview = val;
		}
	}
	  
    resetForm() {
		// @ts-ignore
    	this.$refs.new_stop_form.resetFields();
    }

	/* Lifecycle Hooks */
	beforeMount() {

	}

}
</script>

<style lang='scss' scoped>
.new-stop-form {

	.stop-summary {
		margin: 2rem auto;
	}
	
	.image-wrapper {
		
		.image {
			width: 100%;
			max-width: 100%;
			height: auto;
			object-fit: cover;
		}
	}
}
</style>