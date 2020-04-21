import { notifications } from '../stores';

function notify(message, color = null, timeout = 5000) {
  notifications.update((ns) => [...ns, { message, color }]);
  const timeoutID = setTimeout(() => {
    // Remove the first notification in the array
    notifications.update((ns) => ns.slice(1));
    clearTimeout(timeoutID);
  }, timeout);
}

export default notify;
