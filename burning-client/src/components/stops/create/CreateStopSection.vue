<template>
	<div>
		<el-form :model="newStop" :rules="rules" ref="newStop" label-width="120px" class="new-stop-form">
			<el-form-item label="Stop Name" prop="name">
				<el-input v-model="newStop.name"></el-input>
			</el-form-item>
			<el-form-item label="Include Address">
				<el-switch v-model="showAddressFields"></el-switch>
			</el-form-item>
			<div v-if="showAddressFields">
				<el-form-item label="Address" prop="address.address">
					<el-input v-model="newStop.address.address"></el-input>
				</el-form-item>
				<el-form-item label="City" prop="address.city">
					<el-input v-model="newStop.address.city"></el-input>
				</el-form-item>
				<el-form-item label="State/Province" prop="address.state">
					<el-input v-model="newStop.address.state"></el-input>
				</el-form-item>
				<el-form-item label="Zip Code" prop="address.zipcode">
					<el-input v-model="newStop.address.zipcode"></el-input>
				</el-form-item>
				<el-form-item label="Country" prop="address.country">
					<el-input v-model="newStop.address.country"></el-input>
				</el-form-item>
			</div>
			<el-form-item label="Image Url" prop="imageUrl">
				<el-input v-model="newStop.imageUrl"></el-input>
			</el-form-item>
			<el-form-item label="Description" prop="description">
				<el-input type="textarea" v-model="newStop.description"></el-input>
			</el-form-item>
			<el-form-item>
				<el-button type="primary" @click="submitForm('newStop')">Create</el-button>
				<el-button @click="resetForm('newStop')">Reset</el-button>
			</el-form-item>
		</el-form>
	</div>
</template>

<script lang="ts">
import Vue from 'vue';
import { Component, Prop, Watch } from 'vue-property-decorator';
import { IStop } from '../../../types/Stop';
import Stop from '../../../models/Stop';

@Component({
	name: 'CreateStopSection'
})
export default class CreateStopSection extends Vue {

	/* Props */

	/* Computed */

	/* Data */
	newStop: IStop = new Stop();
	showAddressFields: boolean = false;
	ruleForm = {
          name: '',
          region: '',
          date1: '',
          date2: '',
          delivery: false,
          type: [],
          resource: '',
          desc: ''
	};

	rules = {
		name: [
			{ required: true, message: 'Please input Activity name', trigger: 'blur' },
			{ min: 3, max: 5, message: 'Length should be 3 to 5', trigger: 'blur' }
		],
		region: [
			{ required: true, message: 'Please select Activity zone', trigger: 'change' }
		],
		type: [
			{ type: 'array', required: true, message: 'Please select at least one activity type', trigger: 'change' }
		],
		resource: [
			{ required: true, message: 'Please select activity resource', trigger: 'change' }
		],
		desc: [
			{ required: true, message: 'Please input activity form', trigger: 'blur' }
		]
	};

	/* Methods */

	submitForm(formName) {
        this.$refs[formName].validate(valid => {
          if (valid) {
            alert('submit!');
          } else {
            return false;
          }
        });
	}
	  
    resetForm(formName) {
    	this.$refs[formName].resetFields();
    }

	/* Lifecycle Hooks */
	beforeMount() {

	}

}
</script>

<style lang='scss' scoped>

</style>