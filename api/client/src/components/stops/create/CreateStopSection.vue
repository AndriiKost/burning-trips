<template>
	<div>
		<div v-if="error" style="color: darkred; font-size: 16px;">
			{{ error }}
		</div>
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

			<el-form-item label="Stop Address" prop="address">
				<div class="el-input">
					<input 
						type="text" 
						ref="autocomplete" 
						class="el-input__inner"
					>
				</div>
			</el-form-item>

			<el-form-item label="content" prop="content">
				<el-input type="textarea" v-model="newStop.content"></el-input>
			</el-form-item>

			<el-upload
				v-if="newStop.name"
				class="upload-demo"
				:action="`http://localhost:8080/file-upload/upload-image/${newStop.name}`"
				:on-preview="handlePreview"
				:on-remove="handleRemove"
				:before-upload="beforeUpload"
				:file-list="fileList"
				list-type="picture"
				:on-success="onUploadSuccess"
			>
				<el-button size="small" type="primary">Upload Image</el-button>
			</el-upload>

			<el-form-item label="Preview My Stop">
				<el-switch :value="preview" @input="handlePreview"></el-switch>
			</el-form-item>
			<stop-summary-card
				v-if="preview && valid"
				:stop="newStop"
			/>
			<el-form-item>
				<el-button 
					type="primary" 
					@click="create" 
					:loading="uploading"
				>
				Create
			</el-button>
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
import { Get } from 'vuex-pathify';
import { IUser } from '../../../types/User';
import config from '@/config/index';
import axios from 'axios';
import { IObjectUploadResponse } from '../../../types/Upload';

@Component({
	name: 'CreateStopSection',
	components: {
		StopSummaryCard
	}
})
export default class CreateStopSection extends Vue {

	/* Props */

	/* Computed */
	@Get('auth/loggedInUser')
	loggedInUser: IUser;
 
	/* Data */
	newStop: IStop = new Stop();
	preview: boolean = false;
	previewImage: boolean = false;
	valid: boolean = false;
	error: any = null;
	autocomplete: any = null;
	file: File;
	fileList: File[] = [];
	uploading: boolean = false;
	
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
		content: [
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

	beforeUpload() {
		this.uploading = true;
	}

	onUploadSuccess(res: IObjectUploadResponse, file: File, fileList: File[]) {
		// if (!res.success) // Add failure message to the notification stack
		this.newStop.imageUrl = res.objectUrl;
		this.uploading = false;
	}
	
	async create() {
		this.validate();
		if (!this.valid) return;
		this.newStop.authorID = this.loggedInUser.id;
		const result = await this.$store.dispatch('stop/createStop', this.newStop);
		if (!result) this.error = 'Oops, something went wrong. Please try again!';
		// if success, add message to a notif stack with success info
		this.$router.push({ name: 'Stop List' });
	}

	initAutocomplete() {
		//@ts-ignore
		this.autocomplete = new google.maps.places.Autocomplete((this.$refs.autocomplete), {types: ['geocode']});
		this.autocomplete.addListener('place_changed', () => {
			let place = this.autocomplete.getPlace();
			let { formatted_address } = place;
			let lat = place.geometry.location.lat();
			let lon = place.geometry.location.lng();
			
			this.newStop.longtitude = lon;
			this.newStop.latitude = lat;
			this.newStop.address = formatted_address;
		});
	}

	handleRemove() {
		// console.log('Needs to remove image')
		// TODO
	}

	mounted() {
		this.initAutocomplete();
	}

	beforeDestroy() {
		this.resetForm();
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