import Promise from 'promise-polyfill'
if (!window.Promise) {
  window.Promise = Promise
}
import 'whatwg-fetch'

function sendBeacon(url, data) {
  if (url !== '/log' && document.visibilityState !== 'visible') {
    sendBeacon('/log', document.visibilityState)
    var lis = document.addEventListener("visibilitychange", function() {
      if (document.visibilityState === 'visible') {
        sendBeacon(url, data)
        document.removeEventListener("visibilitychange", lis)
      }
    });
    return
  }

  if (navigator.sendBeacon) {
    navigator.sendBeacon(url, data)
    return
  }

  if (window.fetch) {
    fetch(url, {
      method: "POST",
      body: data
    })
    return
  }

  console.log("cannot send beacon")
}

window.sendBeacon = sendBeacon
