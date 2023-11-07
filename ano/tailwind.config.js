/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    "./index.html",
    "./src/**/*.{js,ts,jsx,tsx,vue}",
  ],
  theme: {
    extend: {},
    screens: {
      'pc': {'max': '1300px'},
      'laptop':{'max':'935px'},
      'tablet':{'max':'760px'},
      'mobile':{'max':'520px'},
    }
  },
  plugins: [
  require('@tailwindcss/forms'),
  ],
}

