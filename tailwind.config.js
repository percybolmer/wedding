/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    'templates/**/*.templ',
    'views/**/*.templ',
  ],
  darkMode: 'class',
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
      {
        mytheme: {
          "primary": "#dc2626",

          "secondary": "#374151",

          "accent": "#4f46e5",

          "neutral": "#ff00ff",

          "base-100": "#374151",

          "info": "#0000ff",

          "success": "#00ff00",

          "warning": "#00ff00",

          "error": "#ff0000",
        },
      },
    ],
  },
}

