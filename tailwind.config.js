/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    'templates/**/*.templ',
    'components/**/*.templ',
  ],
  theme: {
    extend: {},
  },
  plugins: [require("daisyui")],
  daisyui: {
    styled: true,
    themes: true,
    base: true,
    utils: true,
    logs: true,
    rtl: false,
    themes: [
      'dim',
      'light',
    ],
  },
}

