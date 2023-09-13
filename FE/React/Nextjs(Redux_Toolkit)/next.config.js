/** @type {import('next').NextConfig} */

module.exports = {
  reactStrictMode: true,
  trailingSlash: true,
  assetPrefix: '',
  env: {
    ENVIRONMENT: process.env.ENVIRONMENT,
    DEBUGMODE: process.env.DEBUGMODE,
    BUILDMODE: process.env.BUILDMODE,
  },
  webpack: (config) => {
    console.log('[next.config.js] webpack')
    config.module.rules.push({
      test: /\.svg$/,
      issuer: /\.[jt]sx?$/,
      use: ['@svgr/webpack'],
    })
    console.log('webpack config: ', config)
    return config
  },
}
