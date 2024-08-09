import "@angular-wave/angular.ts";
import TodoController from "./todo-ctl";

/**
 * AngularTS is a multi-paradigm framework that can be adjusted to meet any style.
 * This app features a single controller in a style of Stimilus.js (https://stimulus.hotwired.dev/).
 * Its a great way to inject custom DOM behaviour into your server-rendered views.
 */

window.angular.module("todo", []).controller("TodoCtrl", TodoController);
