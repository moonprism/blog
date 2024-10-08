export function debounce(func, delay) {
  let timeoutID;
  return function (...args) {
    clearTimeout(timeoutID);
    timeoutID = setTimeout(function () {
      func.apply(this, args);
    }, delay);
  };
}
