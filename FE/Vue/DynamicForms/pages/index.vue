<template>
  <!--  <FormMakerPage />-->
  <div>
    <h1>Form MAKER</h1>
    <div>
      <label>폼 이름</label>
      <input ref="formName" type="text" @change="formNameChanged" />
    </div>
    <div>
      <label>폼 타입</label>
      <select ref="formType" @change="formTypeChanged">
        <option value="text">텍스트 타입</option>
        <option value="number">숫자 타입</option>
        <option value="check">체크 박스</option>
        <option value="select">셀렉트 박스</option>
      </select>
    </div>
    <ChkBoxMaker v-if="formType === 'check'" />
    <SelBoxMaker v-if="formType === 'select'" />
    <button v-if="formType !== 'check'" @click="makeForm">생성</button>
  </div>
</template>

<script>
import { formDataComputed, formDataMethods } from '~/store/formDataHelper'
import ChkBoxMaker from '~/components/formaker/atoms/ChkBoxMaker.vue'
import SelBoxMaker from '~/components/formaker/atoms/SelBoxMaker.vue'

export default {
  name: 'IndexPage',
  components: { ChkBoxMaker, SelBoxMaker },
  data: () => ({
    formName: '',
    formType: '',
  }),
  computed: { ...formDataComputed },
  methods: {
    ...formDataMethods,
    async formNameChanged(e) {
      await this.formMakingDataAct({
        payload: {
          name: e.target.value,
        },
      })
    },
    async formTypeChanged(e) {
      this.formType = e.target.value
      await this.formMakingDataAct({
        payload: {
          type: e.target.value,
        },
      })
    },
    async makeForm() {
      if (!this.formMakingData.name) {
        return alert('폼 이름을 입력해주세요.')
      }
      await this.formPublishDataAct({
        payload: { ...this.formMakingData },
      })
      console.log(this.formPublishData)
      this.$router.push('/makeformresult')
    },
  },
}
</script>
