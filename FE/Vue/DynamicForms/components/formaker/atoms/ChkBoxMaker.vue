<template>
  <div>
    <label>체크박스 항목 생성</label>
    <div v-for="(checkItemName, index) in checkItemNames" :key="index">
      항목{{ index + 1
      }}<input
        :id="`checkItem${index + 1}`"
        type="text"
        @change="makeChkBoxOpts(index)"
      />
      <button @click="deleteCheckBoxItem(index)">항목빼기</button>
    </div>
    <button @click="addCheckBoxItem">항목추가</button>
    <button @click="makeChkBox">생성</button>
  </div>
</template>

<script>
import { formDataComputed, formDataMethods } from '~/store/formDataHelper'

export default {
  name: 'ChkBoxMaker',
  data: () => ({
    checkItemNames: [],
  }),
  computed: { ...formDataComputed },
  methods: {
    ...formDataMethods,
    addCheckBoxItem() {
      this.checkItemNames.push('')
    },
    deleteCheckBoxItem(index) {
      this.checkItemNames.splice(index, 1) // 값빼기
      this.reBindAllValues()
    },
    // 값을 빼고나서 순서가 바뀐 값으로 폼에 다시 바인딩
    reBindAllValues() {
      this.checkItemNames.forEach((item, index) => {
        document.getElementById(`checkItem${index + 1}`).value = item
      })
    },
    makeChkBoxOpts(index) {
      this.checkItemNames[index] = document.getElementById(
        `checkItem${index + 1}`
      ).value
    },
    async makeChkBox() {
      if (!this.formMakingData.name) {
        return alert('폼 이름을 입력해주세요.')
      }
      this.checkItemNames = this.checkItemNames.filter((item) => item !== '')
      await this.formMakingDataAct({
        payload: {
          options: this.checkItemNames,
        },
      })
      await this.formPublishDataAct({
        payload: {
          ...this.formMakingData,
        },
      })
      console.log('formMakingData', this.formMakingData)
      console.log('formPublishData', this.formPublishData)
      this.$router.push('/makeformresult')
    },
  },
}
</script>

<style scoped></style>
