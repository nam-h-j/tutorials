/* eslint-disable @typescript-eslint/explicit-module-boundary-types */
/* eslint-disable react-hooks/rules-of-hooks */
import { Action, ThunkAction, combineReducers, configureStore } from '@reduxjs/toolkit'
import { useDispatch } from 'react-redux'

// import slice here
import commonReducer from './slice/Common'

const rootReducer = combineReducers({
  Common: commonReducer,
})

export type RootState = ReturnType<typeof rootReducer>
export type AppThunk = ThunkAction<void, RootState, unknown, Action<string>>
export type AppDispatch = typeof store.dispatch

export const userAppDispatch = () => useDispatch<AppDispatch>()

const store = configureStore({
  reducer: rootReducer,
  devTools: process.env.NODE_ENV !== 'production',
})

export default store
