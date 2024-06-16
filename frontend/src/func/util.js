function paddingZero(num) {
  return ('00' + num).slice(-2);
}

// second to mm:ss
function secToMinSec(sec) {
  const min = Math.floor(sec / 60);
  const s = sec % 60;
  return `${paddingZero(min)}:${paddingZero(s)}`;
}

// 2007-05-26T21:57:18+09:00 -> 2007/05/26 21:57
function formatDateTime(dt) {
  const d = new Date(dt);
  const year = d.getFullYear();
  const month = paddingZero(d.getMonth() + 1);
  const day = paddingZero(d.getDate());
  const hour = paddingZero(d.getHours());
  const min = paddingZero(d.getMinutes());
  return `${year}/${month}/${day} ${hour}:${min}`;
}

export default { secToMinSec, formatDateTime };
