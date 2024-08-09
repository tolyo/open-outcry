# `angular-seed` — the seed for AngularTS apps

This project is an application skeleton for a typical [AngularTS][angularts] web app. You can use it to quickly bootstrap your next project.
It contains a sample AngularTS application and is preconfigured with all the necessary tools for developing, testing and deployment.
In order to leave your options open, it attempts to be as minimalist as possible. None of the tools are specific to or even necessary
for AngularTS development and would apply equally well to any generic web project.

For inspiration, the seed comes bundled with the following apps under `apps` folder:

- A todo list

### Install Dependencies

```
make setup
```

### Run the Application

This project preconfigured the project with [Web Dev Server](https://modern-web.dev/docs/dev-server/overview/), an optimal solution for lightweight and buildless workflows.
The simplest way to start the server is:

```
make start
```

Now open your browser to `localhost:4000/`

[angularts]: https://github.com/Angular-Wave/angular.ts

### Build the App for Production

AngularTS apps consist of static HTML, CSS, and JavaScript files that need to be hosted on a server accessible to browsers.
To generate a production-ready bundle with minified HTML, CSS, and JavaScript, run:

```
make build
```

This will execute the following tasks:

- Minified HTML with [Rollup Plugin HTML](https://modern-web.dev/docs/building/rollup-plugin-html/)
- Bundled JS with [Rollup](https://rollupjs.org/) and minified it with terser [Terser](https://terser.org/)
- Bundled and minified CSS with [Lightning CSS](https://lightningcss.dev/)

Your app should be available in `/dist` folder and can then be uploaded to a static server.
