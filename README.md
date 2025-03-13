Branches organize learning projects and exercise platforms (like Exercism, HackerRank, etc.). After learning projects are completed, they will be merged into the main branch, while exercise platforms will remain as independent branches.


* **React**
  * [**Recommendations for Creating a React App** (Recommended Full-Stack Frameworks)](https://react.dev/learn/creating-a-react-app)
      * [**React App from Scratch**](https://react.dev/learn/build-a-react-app-from-scratch)
      * **Next.js**
        * [create-next-app](https://nextjs.org/docs/app/api-reference/cli/create-next-app)
        * Fork [next-learn](github.com/vercel/next-learn)
      * [**create-react-app - Deprecated!**](https://github.com/facebook/create-react-app)  
        * [**Docs**](https://create-react-app.dev/docs/getting-started)
* **React Native**
  * [**React Native Tutorial**](https://reactnative.dev/docs/tutorial)  
      * [React Native Documentation](https://reactnative.dev/docs/intro-react)
  * [**create-react-native-app**](https://github.com/expo/create-react-native-app)

## React.js  Sandbox
```html
<!DOCTYPE html>
<html>
  <head>
    <meta charset="UTF-8" />
    <title>Hello World</title>
    <script src="https://unpkg.com/react@18/umd/react.development.js"></script>
    <script src="https://unpkg.com/react-dom@18/umd/react-dom.development.js"></script>

    <!-- Don't use this in production: this page is a great way to try React but it's not suitable for production.
      It slowly compiles JSX with Babel in the browser and uses a large development build of React.-->
    <script src="https://unpkg.com/@babel/standalone/babel.min.js"></script>
  </head>
  <body>
    <div id="root"></div>
    <script type="text/babel">
    
      function MyApp() {
        return <h1>Hello, world!</h1>;
      }

      const container = document.getElementById('root');
      const root = ReactDOM.createRoot(container);
      root.render(<MyApp />);

    </script>
  </body>
</html>
```