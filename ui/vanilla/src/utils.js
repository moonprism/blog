export function debounce(func, delay) {
  let timeoutID;
  return function (...args) {
    clearTimeout(timeoutID);
    timeoutID = setTimeout(function () {
      func.apply(this, args);
    }, delay);
  };
}

export function timeAgoStr(timeStamp) {
  const minute = 1000 * 60;
  const hour = minute * 60;
  const day = hour * 24;
  const month = day * 30;
  const now = new Date().getTime();
  const diffValue = now - timeStamp*1000;
  if (diffValue < 0) {
    return;
  }
  const monthC = diffValue / month;
  const weekC = diffValue / (7 * day);
  const dayC = diffValue / day;
  const hourC = diffValue / hour;
  const minC = diffValue / minute;
  const rtf = new Intl.RelativeTimeFormat("zh", {
    numeric: "auto",
  });
  if (monthC > 12) {
    return rtf.format(-parseInt(monthC / 12), "year");
  } else if (monthC >= 1 && monthC <= 12) {
    return rtf.format(-parseInt(monthC), "month");
  } else if (weekC >= 1) {
    return rtf.format(-parseInt(weekC), "week");
  } else if (dayC >= 1) {
    return rtf.format(-parseInt(dayC), "day");
  } else if (hourC >= 1) {
    return rtf.format(-parseInt(hourC), "hour");
  } else if (minC >= 1) {
    return rtf.format(-parseInt(minC), "minute");
  } else {
    return "刚刚";
  }
}
