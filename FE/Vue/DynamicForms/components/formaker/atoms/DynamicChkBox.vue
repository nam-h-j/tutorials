<template>
  <div>
    <div>{{ formData.name }}</div>
    <div v-for="(option, index) in formData.options" :key="index">
      <label>
        <input type="checkbox" :value="option" @click="inputChange" />{{
          option
        }}
      </label>
    </div>
  </div>
</template>

<script>
export default {
  name: 'DynamicChkBox',
  props: {
    formData: {
      type: [Object],
      default: () => {},
    },
  },
  data() {
    return {
      chkData: [],
    }
  },
  methods: {
    inputChange(e) {
      if (e.target.checked) {
        this.chkData.push(e.target.value)
      } else {
        this.chkData = this.chkData.filter((item) => item !== e.target.value)
      }
      this.$emit('updateChkData', this.chkData)
    },
  },
}
</script>
