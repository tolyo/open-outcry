import ToDoList from "./todo-list";

export default class TodoController {
  constructor() {
    /** @type {ToDoList} */
    this.list = new ToDoList();
  }

  /**
   * Create a new task
   * @param {String} task
   * @return {void}
   */
  add(task) {
    this.list.add(task);
  }

  /**
   * Delete all finished tasks
   * @return {void}
   */
  archive() {
    this.list.archive();
  }
}
