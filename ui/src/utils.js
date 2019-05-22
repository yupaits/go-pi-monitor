const minute = 60;
const hour = 60 * minute;
const day = 24 * hour;

const per = 1024;
const sizes = ['B', 'KB', 'MB', 'GB', 'TB', 'PB', 'EB', 'ZB', 'YB'];

export default {
  relativeTime(duration) {
    let result = '';
    if (!duration) {
      return result;
    }
    if (duration > day) {
      result += `${Math.floor(duration / day)}天 `;
      duration = duration % day;
    }
    if (duration > hour) {
      result += `${Math.floor(duration / hour)}小时 `;
      duration = duration % hour;
    }
    if (duration > minute) {
      result += `${Math.floor(duration / minute)}分钟 `;
      duration = duration % minute;
    }
    result += `${duration}秒`;
    return result;
  },
  percentStatus(percent) {
    if (!percent) {
      return 'active';
    }
    if (percent >= 99) {
      return 'exception';
    } else if (percent >= 80) {
      return 'normal';
    } else {
      return 'active'; 
    }
  },
  bytesToSize(bytes) {
    if (!bytes) {
      return '0 B';
    }
    const i = Math.floor(Math.log(bytes) / Math.log(per));
    return `${(bytes / Math.pow(per, i)).toFixed(1)} ${sizes[i]}`;
  }
}