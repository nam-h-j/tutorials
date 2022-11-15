<template>
  <div>
    <label>{{ formData.name }}</label>
    <input type="text" @keyup="inputChange" />
  </div>
</template>

<script>
export default {
  name: 'DynamicNumberInput',
  props: {
    formData: {
      type: [Object],
      default: () => {},
    },
  },
  data() {
    return {
      numData: '',
    }
  },
  methods: {
    inputChange(e) {
      this.numData = e.target.value
      // 부모컴포넌트로 값을 넘겨준다.
      this.$emit('updateNumData', this.numData)
    },
  },
}
</script>

<style scoped></style>
