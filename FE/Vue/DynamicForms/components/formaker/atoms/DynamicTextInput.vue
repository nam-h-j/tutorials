<template>
  <div>
    <label>{{ formData.name }}</label>
    <input type="text" @keyup="inputChange" />
  </div>
</template>

<script>
export default {
  name: 'DynamicTextInput',
  props: {
    formData: {
      type: [Object],
      default: () => {},
    },
  },
  data() {
    return {
      textData: '',
    }
  },
  methods: {
    inputChange(e) {
      this.textData = e.target.value
      // 부모컴포넌트로 값을 넘겨준다.
      this.$emit('updateTextData', this.textData)
    },
  },
}
</script>

<style scoped></style>
