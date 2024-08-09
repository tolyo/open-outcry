export default class Todo {
  /**
   * @param {String} task - to be done
   */
  constructor(task) {
    /** @type {String} */
    this.task = task;
    /** @type {boolean} */
    this.done = false;
  }
}
