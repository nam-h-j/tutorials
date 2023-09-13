import { createSlice } from '@reduxjs/toolkit'
import { commonStateType } from '../stateTypes'
import { AppThunk } from '../createStore'

const initialState: commonStateType = {
  isLogin: false,
  isDark: false,
  isMente: false,
  layoutType: 'launcher',
  themeType: 'default',
  isSm: false,
  isMd: false,
  newNotice: false,
}

export const commonSlice = createSlice({
  name: 'Common',
  initialState: initialState,
  reducers: {
    resetStateReducer: () => initialState,

    isSampleReducer: (state, action) => {
      state.isLogin = action.payload
    },
  },
})

export const {
  isSampleReducer,
} = commonSlice.actions

export const updateSample = () => {
  return (dispatch: any): void => {
    dispatch(isSampleReducer())
  }
}

// login, notice cookie check relation

export default commonSlice.reducer
