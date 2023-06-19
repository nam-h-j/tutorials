/*
 * modalBSPayReq = 바텀시트, 결제페이지 요청사항 옵션
 * modalCouponForLikeAndAcc = 마이마켓 찜 + 계정연동 쿠폰
 * */
export const state = () => ({
  formMakingData: { name: '', type: 'text' },
  formPublishData: [],
})

export const mutations = {
  formMakingDataMut(state, payload) {
    state.formMakingData = { ...state.formMakingData, ...payload }
  },
  formPublishDataMut(state, payload) {
    state.formPublishData.push(payload)
  },
}

export const actions = {
  formMakingDataAct({ commit }, { payload }) {
    commit('formMakingDataMut', payload)
  },
  formPublishDataAct({ commit }, { payload }) {
    commit('formPublishDataMut', payload)
  },
}
