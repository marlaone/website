/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    "./index.html",
    "./views/**/*.html",
    "./src/**/*.{vue,js,ts,jsx,tsx}",
  ],
  theme: {
    extend: {},
  },
  plugins: [require('daisyui')],
}
