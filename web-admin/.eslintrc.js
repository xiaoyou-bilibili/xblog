module.exports = {
  root: true,
  env: {
    browser: true,
    node: true
  },
  parserOptions: {
    parser: 'babel-eslint'
  },
  extends: [
    '@nuxtjs',
    'plugin:nuxt/recommended'
  ],
  plugins: [
  ],
  // add your custom rules here
  rules: {
    'vue/no-v-html': 'off', // 关闭v-html 检查
    'vue/require-prop-types': 'off', // 关闭prop提示'
    'no-unused-vars': 'off'
  }
}
