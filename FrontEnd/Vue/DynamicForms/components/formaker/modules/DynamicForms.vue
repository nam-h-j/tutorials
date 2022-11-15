<template>
  <div class="DynamicForms">
    <div v-if="formData.type === 'text'">
      <DynamicTextInput
        :form-data="formData"
        @updateTextData="updateTextData"
      />
    </div>
    <div v-if="formData.type === 'number'">
      <DynamicNumberInput
        :form-data="formData"
        @updateNumData="updateNumData"
      />
    </div>
    <div v-if="formData.type === 'check'">
      <DynamicChkBox :form-data="formData" @updateChkData="updateChkData" />
    </div>
    <div v-if="formData.type === 'select'">
      <DynamicSelBox :form-data="formData" @updateSelData="updateSelData" />
    </div>
  </div>
</template>

<script>
import DynamicTextInput from '~/components/formaker/atoms/DynamicTextInput'
import DynamicNumberInput from '~/components/formaker/atoms/DynamicNumberInput'
import DynamicChkBox from '~/components/formaker/atoms/DynamicChkBox'
import DynamicSelBox from '~/components/formaker/atoms/DynamicSelBox'

export default {
  name: 'DynamicForms',
  components: {
    DynamicSelBox,
    DynamicTextInput,
    DynamicNumberInput,
    DynamicChkBox,
  },
  props: {
    formData: {
      type: [Object],
      default: () => {},
    },
  },
  data: () => ({
    textData: {},
    numData: {},
    chkData: [],
  }),
  mounted() {
    console.log('formData : ', this.formData.type)
  },
  methods: {
    updateTextData(textData) {
      this.textData = { ...this.textData, [this.formData.name]: textData }
      this.$emit('getTextData', this.textData)
    },
    updateNumData(numData) {
      this.numData = { ...this.numData, [this.formData.name]: numData }
      this.$emit('getNumData', this.numData)
    },
    updateChkData(chkData) {
      this.chkData = { ...this.chkData, [this.formData.name]: chkData }
      this.$emit('getChkData', this.chkData)
    },
    updateSelData(selData) {
      this.selData = { ...this.selData, [this.formData.name]: selData }
      this.$emit('getSelData', this.selData)
    },
  },
}
</script>

<style scoped></style>
