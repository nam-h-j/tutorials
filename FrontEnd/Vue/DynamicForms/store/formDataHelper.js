import { mapActions, mapState } from 'vuex'

export const formDataComputed = {
  ...mapState('formData', {
    formMakingData: (state) => state.formMakingData,
    formPublishData: (state) => state.formPublishData,
  }),
}
export const formDataMethods = mapActions('formData', [
  'formMakingDataAct',
  'formPublishDataAct',
])
