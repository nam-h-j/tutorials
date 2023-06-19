<template>
  <div ref="formArea">
    <DynamicForms
      v-for="(formData, index) in formPublishData"
      :key="index"
      :form-data="formData"
      @getTextData="getTextData"
      @getNumData="getNumData"
      @getChkData="getChkData"
      @getSelData="getSelData"
    />
    <button @click="submitForm">Submit</button>
  </div>
</template>

<script>
import { formDataComputed } from '~/store/formDataHelper'
import DynamicForms from '~/components/formaker/modules/DynamicForms'

export default {
  name: 'MakeFormResult',
  components: { DynamicForms },
  data: () => ({
    submitData: {},
  }),
  computed: { ...formDataComputed },
  mounted() {
    console.log('formMakingData : ', this.formMakingData)
  },
  methods: {
    getTextData(textData) {
      this.submitData = { ...this.submitData, ...textData }
    },
    getNumData(numData) {
      this.submitData = { ...this.submitData, ...numData }
    },
    getChkData(chkData) {
      this.submitData = { ...this.submitData, ...chkData }
    },
    getSelData(selData) {
      console.log(selData)
      this.submitData = { ...this.submitData, ...selData }
    },
    submitForm() {
      alert(JSON.stringify(this.submitData))
      console.log('submitData', this.submitData)
    },
  },
}
</script>

<style scoped></style>
