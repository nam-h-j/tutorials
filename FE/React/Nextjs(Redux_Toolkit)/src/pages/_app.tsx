/* eslint-disable @typescript-eslint/naming-convention */
import { AppProps } from 'next/app'
import '../scss/_reset.scss'
import '../scss/_base.scss'

// redux
import { Provider } from 'react-redux'
import store from '../store/createStore'



const MyApp = ({ Component, pageProps }: AppProps): JSX.Element => {
  return (
    <Provider store={store}>
      {/* global layouy here */}
      <div>
        <Component {...pageProps} />
      </div>
    </Provider>
  )
}
export default MyApp
