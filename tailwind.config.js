const defaultTheme = require('tailwindcss/defaultTheme')
const colors = require('tailwindcss/colors')

module.exports = {
  purge: [],
  darkMode: false, // or 'media' or 'class'
  theme: {
    extend: {},
    colors: {
        gray: colors.coolGray,
        indigo: colors.indigo,
        red: colors.red,
        white: colors.white,
        'axo-green':'#44ABA3',
        'axo-green-100':'#8ECCC7',
        'axo-green-200':'#7CC4BE',
        'axo-green-300':'#69BBB5',
        'axo-green-400':'#56B3AC',
        'axo-green-500':'#56B3AC',
        'axo-green-600':'#368882',
        'axo-green-700':'#2F7772',
    },
    fontFamily: {
        sans: ['Inter var', ...defaultTheme.fontFamily.sans],
    },
  },
  variants: {
    extend: {},
  },
  plugins: [
    require('@tailwindcss/aspect-ratio'),
    require('@tailwindcss/forms'),
    require('@tailwindcss/typography'),
  ],
}
