<template>
  <div>
    <div>{{ formData.name }}</div>
    <select @click="inputChange">
      <option
        v-for="(option, index) in formData.options"
        :key="index"
        :value="option"
      >
        {{ option }}
      </option>
    </select>
  </div>
</template>

<script>
export default {
  name: 'DynamicSelBox',
  props: {
    formData: {
      type: [Object],
      default: () => {},
    },
  },
  data() {
    return {
      selData: '',
    }
  },
  methods: {
    inputChange(e) {
      if (e.target.value) {
        this.selData = e.target.value
      }
      this.$emit('updateSelData', this.selData)
    },
  },
}
</script>
