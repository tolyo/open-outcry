import Todo from "./todo";

export default class ToDoList {
  constructor() {
    /** @type {Todo[]} */
    this.tasks = [
      new Todo("Learn AngularTS"),
      new Todo("Build an AngularTS app"),
    ];
  }

  /**
   * @param {String} task
   * @return {void}
   */
  add(task) {
    this.tasks.push(new Todo(task));
  }

  /**
   * Delete all finished tasks
   * @return {void}
   */
  archive() {
    this.tasks = this.tasks.filter((task) => !task.done);
  }
}
